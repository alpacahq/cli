package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/format"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"unicode"
)

var (
	schemaByGoName  map[string]*schemaInfo
	schemaByOASName map[string]*schemaInfo
)

type specConfig struct {
	specFile     string
	prefix       string // output file prefix: "trading" or "marketdata"
	clientName   string
	clientVar    string // package-level var in root.go: "tradingClient" or "dataClient"
	baseURLField string // field on client.Client: "BaseURL" or "DataURL"
}

type specResult struct {
	config    specConfig
	spec      map[string]any
	schemas   []*schemaInfo
	endpoints []*endpointInfo
}

var specs = []specConfig{
	{"trading-api.json", "trading", "Trading", "tradingClient", "BaseURL"},
	{"market-data-api.json", "marketdata", "MarketData", "dataClient", "DataURL"},
}

func main() {
	root := findProjectRoot()
	outDir := filepath.Join(root, "internal", "api")
	if err := os.MkdirAll(outDir, 0o755); err != nil {
		log.Fatalf("creating output dir: %v", err)
	}

	// Phase 1: load specs, extract schemas and endpoints.
	results := make([]specResult, len(specs))
	var allSchemas []*schemaInfo
	for i, sc := range specs {
		spec := loadSpec(filepath.Join(root, "api", "specs", sc.specFile))
		schemas := extractSchemas(spec)
		endpoints := extractEndpoints(spec)
		dedup(schemas)
		allSchemas = append(allSchemas, schemas...)
		results[i] = specResult{sc, spec, schemas, endpoints}
	}

	schemaByGoName = make(map[string]*schemaInfo, len(allSchemas))
	schemaByOASName = make(map[string]*schemaInfo, len(allSchemas))
	for _, s := range allSchemas {
		schemaByGoName[s.goName] = s
		schemaByOASName[s.name] = s
	}

	// Phase 2: generate types, clients, descriptions, and commands.
	var ops []*opDesc
	var allEndpoints []*endpointInfo
	for _, r := range results {
		writeGo(filepath.Join(outDir, r.config.prefix+"_types.gen.go"), genTypes(r.config.specFile, r.schemas))
		writeGo(filepath.Join(outDir, r.config.prefix+"_client.gen.go"), genClient(r.config.clientName, r.config.specFile, r.endpoints, r.config.baseURLField))
		ops = append(ops, collectDescriptions(r.endpoints, r.spec)...)
		for _, ep := range r.endpoints {
			ep.clientVar = r.config.clientVar
		}
		allEndpoints = append(allEndpoints, r.endpoints...)
	}

	sort.Slice(ops, func(i, j int) bool {
		return ops[i].operationID < ops[j].operationID
	})
	writeGo(filepath.Join(outDir, "descriptions.gen.go"), writeTypedDescriptionsFile(ops))

	epByOp := make(map[string]*endpointInfo, len(allEndpoints))
	for _, ep := range allEndpoints {
		epByOp[ep.goName] = ep
	}

	cmdDir := filepath.Join(root, "internal", "cmd")
	checkExhaustive(epByOp)
	writeGo(filepath.Join(cmdDir, "commands.gen.go"), genCommands(epByOp))

	fmt.Printf("Generated %d types, %d endpoints, %d operations\n",
		len(allSchemas), len(allEndpoints), len(ops))
}

// --- Spec loading ---

func loadSpec(path string) map[string]any {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("reading spec: %v", err)
	}
	var spec map[string]any
	if err := json.Unmarshal(data, &spec); err != nil {
		log.Fatalf("parsing spec: %v", err)
	}
	return spec
}

func findProjectRoot() string {
	dir, _ := os.Getwd()
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			log.Fatal("could not find project root (no go.mod)")
		}
		dir = parent
	}
}

// --- Schema extraction ---

type schemaInfo struct {
	name       string
	goName     string
	kind       string // "struct", "enum", "alias"
	props      map[string]map[string]any
	required   map[string]bool
	enumValues []string
	raw        map[string]any
}

func extractSchemas(spec map[string]any) []*schemaInfo {
	components := mapGet(spec, "components")
	schemas := mapGet(components, "schemas")

	var result []*schemaInfo
	for name, raw := range schemas {
		s, ok := raw.(map[string]any)
		if !ok {
			continue
		}
		info := &schemaInfo{name: name, goName: toGoName(name), raw: s}
		_, hasProps := s["properties"].(map[string]any)
		_, hasEnum := s["enum"].([]any)

		if hasEnum {
			info.kind = "enum"
			info.enumValues = extractEnumValues(s)
		} else if hasProps {
			info.kind = "struct"
			info.props = extractProps(s)
			info.required = extractRequired(s)
		} else {
			info.kind = "alias"
		}
		result = append(result, info)
	}

	sort.Slice(result, func(i, j int) bool { return result[i].name < result[j].name })
	return result
}

func dedup(schemas []*schemaInfo) {
	seen := map[string]int{}
	for i, s := range schemas {
		if idx, ok := seen[s.goName]; ok {
			prev := schemas[idx]
			if unicode.IsUpper(rune(prev.name[0])) {
				s.goName = s.goName + "V3"
			} else {
				delete(seen, prev.goName)
				prev.goName = prev.goName + "V3"
				seen[prev.goName] = idx
			}
		}
		seen[s.goName] = i
	}
}

// --- Endpoint extraction ---

type endpointInfo struct {
	method        string
	path          string
	operationID   string
	summary       string
	description   string // OAS operation description (longer than summary)
	goName        string
	pathParams    []paramInfo
	queryParams   []paramInfo
	bodyRef       string
	bodyInline    map[string]any
	responseRef   string
	returnsArray  bool
	responseEmpty bool   // true if 2xx response has no content body (204 etc.)
	clientVar     string // package-level var name for generated commands
}

type paramInfo struct {
	name        string
	goName      string
	goType      string
	required    bool
	enumValues  []string
	defaultVal  string
	description string
}

func extractEndpoints(spec map[string]any) []*endpointInfo {
	paths, _ := spec["paths"].(map[string]any)
	compParams := map[string]map[string]any{}
	if cp := mapGet(mapGet(spec, "components"), "parameters"); cp != nil {
		for k, v := range cp {
			if m, ok := v.(map[string]any); ok {
				compParams[k] = m
			}
		}
	}
	compSchemas := mapGet(mapGet(spec, "components"), "schemas")

	var result []*endpointInfo
	for path, methods := range paths {
		methodMap, ok := methods.(map[string]any)
		if !ok {
			continue
		}
		// Path-level parameters (shared by all methods on this path)
		pathLevelParams, _ := methodMap["parameters"].([]any)

		for method, opRaw := range methodMap {
			if method == "parameters" {
				continue
			}
			op, ok := opRaw.(map[string]any)
			if !ok {
				continue
			}
			opID, _ := op["operationId"].(string)
			if opID == "" {
				continue
			}
			summary, _ := op["summary"].(string)
			description, _ := op["description"].(string)

			ep := &endpointInfo{
				method:      strings.ToUpper(method),
				path:        path,
				operationID: opID,
				summary:     summary,
				description: description,
				goName:      toGoName(opID),
			}

			// Merge path-level + operation-level parameters
			var allParams []any
			allParams = append(allParams, pathLevelParams...)
			opParams, _ := op["parameters"].([]any)
			allParams = append(allParams, opParams...)

			for _, pRaw := range allParams {
				p, ok := pRaw.(map[string]any)
				if !ok {
					continue
				}
				if ref, ok := p["$ref"].(string); ok {
					refName := refBaseName(ref)
					if resolved, ok := compParams[refName]; ok {
						p = resolved
					} else {
						fmt.Fprintf(os.Stderr, "Warning: unresolved parameter $ref %q in %s %s\n", ref, method, path)
						continue
					}
				}
				name, _ := p["name"].(string)
				in, _ := p["in"].(string)
				req, _ := p["required"].(bool)
				pSchema, _ := p["schema"].(map[string]any)
				goType := "string"
				if pSchema != nil {
					pSchema = resolveRef(pSchema, compSchemas)
					goType = schemaToFlagType(pSchema)
				}
				pi := paramInfo{name: name, goName: toGoName(name), goType: goType, required: req}
				pi.description, _ = p["description"].(string)
				if pi.description == "" && pSchema != nil {
					pi.description = schemaDesc(pSchema, compSchemas)
				}
				if pSchema != nil {
					pi.enumValues = extractEnumValues(pSchema)
					pi.defaultVal = extractDefault(pSchema)
				}
				switch in {
				case "path":
					ep.pathParams = append(ep.pathParams, pi)
				case "query":
					ep.queryParams = append(ep.queryParams, pi)
				}
			}

			if rb, ok := op["requestBody"].(map[string]any); ok {
				if content, ok := rb["content"].(map[string]any); ok {
					if jc, ok := content["application/json"].(map[string]any); ok {
						if schema, ok := jc["schema"].(map[string]any); ok {
							if ref, ok := schema["$ref"].(string); ok {
								ep.bodyRef = toGoName(refBaseName(ref))
							} else if _, ok := schema["properties"]; ok {
								ep.bodyInline = schema
							}
						}
					}
				}
			}

			ep.responseEmpty = true
			if responses, ok := op["responses"].(map[string]any); ok {
				var codes []string
				for code := range responses {
					if len(code) > 0 && code[0] == '2' {
						codes = append(codes, code)
					}
				}
				sort.Strings(codes)
				for _, code := range codes {
					resp, ok := responses[code].(map[string]any)
					if !ok {
						continue
					}
					if content, ok := resp["content"].(map[string]any); ok {
						ep.responseEmpty = false
						if jc, ok := content["application/json"].(map[string]any); ok {
							if schema, ok := jc["schema"].(map[string]any); ok {
								if ref, ok := schema["$ref"].(string); ok {
									ep.responseRef = toGoName(refBaseName(ref))
								} else if t, _ := schema["type"].(string); t == "array" {
									if items, ok := schema["items"].(map[string]any); ok {
										if ref, ok := items["$ref"].(string); ok {
											ep.responseRef = toGoName(refBaseName(ref))
											ep.returnsArray = true
										}
									}
								}
							}
						}
					}
					break
				}
			}

			result = append(result, ep)
		}
	}

	sort.Slice(result, func(i, j int) bool {
		if result[i].path == result[j].path {
			return methodOrder(result[i].method) < methodOrder(result[j].method)
		}
		return result[i].path < result[j].path
	})
	return result
}

func methodOrder(m string) int {
	switch m {
	case "GET":
		return 0
	case "POST":
		return 1
	case "PUT":
		return 2
	case "PATCH":
		return 3
	case "DELETE":
		return 4
	default:
		return 5
	}
}

// --- Type generation ---

func genTypes(specFile string, schemas []*schemaInfo) string {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "// Code generated from api/specs/%s; DO NOT EDIT.\n\n", specFile)
	fmt.Fprintf(&buf, "package api\n\n")

	for _, s := range schemas {
		if s.kind == "enum" {
			fmt.Fprintf(&buf, "type %s string\n\n", s.goName)
		}
	}
	for _, s := range schemas {
		if s.kind == "struct" {
			genStruct(&buf, s)
		}
	}
	return buf.String()
}

func genStruct(buf *bytes.Buffer, s *schemaInfo) {
	fmt.Fprintf(buf, "type %s struct {\n", s.goName)
	fields := sortedKeys(s.props)

	for _, fieldName := range fields {
		fieldSchema := s.props[fieldName]
		goField := toGoName(fieldName)
		goType := resolveGoType(fieldSchema)
		tag := fieldName
		if !s.required[fieldName] {
			tag += ",omitempty"
		}
		fmt.Fprintf(buf, "\t%s %s `json:%q`\n", goField, goType, tag)
	}
	fmt.Fprintf(buf, "}\n\n")
}

func resolveGoType(schema map[string]any) string {
	if ref, ok := schema["$ref"].(string); ok {
		rn := refBaseName(ref)
		if s, ok := schemaByOASName[rn]; ok {
			if s.kind == "alias" {
				return resolveGoType(s.raw)
			}
			return s.goName
		}
		return toGoName(rn)
	}

	typ, _ := schema["type"].(string)
	switch typ {
	case "string":
		return "string"
	case "integer":
		return "int"
	case "number":
		return "float64"
	case "boolean":
		return "bool"
	case "array":
		if items, ok := schema["items"].(map[string]any); ok {
			return "[]" + resolveGoType(items)
		}
		return "[]any"
	case "object":
		if addProps, ok := schema["additionalProperties"].(map[string]any); ok {
			return "map[string]" + resolveGoType(addProps)
		}
		return "map[string]any"
	default:
		return "any"
	}
}

// --- Client generation ---

func genClient(clientName, specFile string, endpoints []*endpointInfo, baseURLField string) string {
	var body bytes.Buffer

	fmt.Fprintf(&body, "// %sClient provides typed methods for the %s API.\n", clientName, clientName)
	fmt.Fprintf(&body, "type %sClient struct {\n\tRaw *client.Client\n\tbaseURL string\n}\n\n", clientName)
	fmt.Fprintf(&body, "func New%sClient(raw *client.Client) *%sClient {\n", clientName, clientName)
	fmt.Fprintf(&body, "\treturn &%sClient{Raw: raw, baseURL: raw.%s}\n}\n\n", clientName, baseURLField)

	for _, ep := range endpoints {
		genEndpointMethod(&body, ep, clientName)
	}

	validated := map[string]bool{}
	for _, ep := range endpoints {
		if ep.bodyRef == "" || validated[ep.bodyRef] {
			continue
		}
		s := findSchema(ep.bodyRef)
		if s != nil && s.kind == "struct" && len(s.required) > 0 {
			var reqFields []string
			for fn, isReq := range s.required {
				if !isReq {
					continue
				}
				fs, ok := s.props[fn]
				if !ok {
					continue
				}
				if resolveGoType(fs) == "string" {
					reqFields = append(reqFields, fn)
				}
			}
			sort.Strings(reqFields)
			genValidateMethod(&body, s.goName, reqFields)
			validated[ep.bodyRef] = true
		}
	}

	// Assemble with header — detect imports from generated body
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "// Code generated from api/specs/%s; DO NOT EDIT.\n\n", specFile)
	fmt.Fprintf(&buf, "package api\n\n")

	bodyStr := body.String()
	fmt.Fprintf(&buf, "import (\n")
	if strings.Contains(bodyStr, "json.RawMessage") {
		fmt.Fprintf(&buf, "\t\"encoding/json\"\n")
	}
	if strings.Contains(bodyStr, "fmt.") {
		fmt.Fprintf(&buf, "\t\"fmt\"\n")
	}
	if strings.Contains(bodyStr, "url.") {
		fmt.Fprintf(&buf, "\t\"net/url\"\n")
	}
	if strings.Contains(bodyStr, "strings.") {
		fmt.Fprintf(&buf, "\t\"strings\"\n")
	}
	fmt.Fprintf(&buf, "\n\t\"github.com/alpacahq/cli/internal/client\"\n")
	fmt.Fprintf(&buf, ")\n\n")
	buf.Write(body.Bytes())

	return buf.String()
}

func genEndpointMethod(buf *bytes.Buffer, ep *endpointInfo, clientName string) {
	hasParams := len(ep.queryParams) > 0

	if ep.responseRef != "" && !schemaHasType(ep.responseRef) {
		ep.responseRef = ""
	}

	bodyType := ep.bodyRef
	if bodyType == "" && ep.bodyInline != nil {
		bodyType = ep.goName + "Request"
		if findSchema(bodyType) == nil {
			genInlineRequestStruct(buf, bodyType, ep.bodyInline)
		}
	}

	respType := ep.responseRef

	var sig strings.Builder
	fmt.Fprintf(&sig, "func (c *%sClient) %s(", clientName, ep.goName)

	var args []string
	for _, pp := range ep.pathParams {
		args = append(args, pp.goName+" string")
	}
	if hasParams {
		args = append(args, "params url.Values")
	}
	if bodyType != "" {
		args = append(args, "body *"+bodyType)
	}
	sig.WriteString(strings.Join(args, ", "))
	sig.WriteString(")")

	if respType != "" {
		if ep.returnsArray {
			fmt.Fprintf(&sig, " ([]%s, error)", respType)
		} else {
			fmt.Fprintf(&sig, " (*%s, error)", respType)
		}
	} else {
		sig.WriteString(" (json.RawMessage, error)")
	}

	if ep.summary != "" {
		fmt.Fprintf(buf, "// %s — %s\n", ep.goName, ep.summary)
	}
	fmt.Fprintf(buf, "%s {\n", sig.String())

	pathExpr := buildPathExpr(ep.path, ep.pathParams)
	paramsExpr := "nil"
	if hasParams {
		paramsExpr = "params"
	}
	bodyExpr := "nil"
	if bodyType != "" {
		bodyExpr = "body"
	}

	callExpr := fmt.Sprintf("c.Raw.Do(%q, c.baseURL, %s, %s, %s)", ep.method, pathExpr, paramsExpr, bodyExpr)

	if respType != "" {
		if ep.returnsArray {
			fmt.Fprintf(buf, "\treturn unmarshalSlice[%s](%s)\n", respType, callExpr)
		} else {
			fmt.Fprintf(buf, "\treturn unmarshal[%s](%s)\n", respType, callExpr)
		}
	} else {
		fmt.Fprintf(buf, "\treturn %s\n", callExpr)
	}

	fmt.Fprintf(buf, "}\n\n")
}

func genInlineRequestStruct(buf *bytes.Buffer, name string, schema map[string]any) {
	props := extractProps(schema)
	if props == nil {
		return
	}
	reqMap := extractRequired(schema)
	genStruct(buf, &schemaInfo{goName: name, props: props, required: reqMap})

	var reqFields []string
	for fn := range reqMap {
		if p, ok := props[fn]; ok {
			if t, _ := p["type"].(string); t == "string" {
				reqFields = append(reqFields, fn)
			}
		}
	}
	sort.Strings(reqFields)
	genValidateMethod(buf, name, reqFields)
}

func genValidateMethod(buf *bytes.Buffer, typeName string, reqFields []string) {
	if len(reqFields) == 0 {
		return
	}
	fmt.Fprintf(buf, "func (r *%s) Validate() error {\n", typeName)
	fmt.Fprintf(buf, "\tvar missing []string\n")
	for _, fn := range reqFields {
		fmt.Fprintf(buf, "\tif r.%s == \"\" { missing = append(missing, %q) }\n", toGoName(fn), fn)
	}
	fmt.Fprintf(buf, "\tif len(missing) > 0 {\n")
	fmt.Fprintf(buf, "\t\treturn fmt.Errorf(\"missing required fields: %%s\", strings.Join(missing, \", \"))\n")
	fmt.Fprintf(buf, "\t}\n")
	fmt.Fprintf(buf, "\treturn nil\n")
	fmt.Fprintf(buf, "}\n\n")
}

func schemaHasType(goName string) bool {
	s := findSchema(goName)
	return s != nil && (s.kind == "struct" || s.kind == "enum")
}

func buildPathExpr(path string, pathParams []paramInfo) string {
	if len(pathParams) == 0 {
		return fmt.Sprintf("%q", path)
	}
	expr := path
	var fmtArgs []string
	for _, p := range pathParams {
		placeholder := "{" + p.name + "}"
		if strings.Contains(expr, placeholder) {
			expr = strings.Replace(expr, placeholder, "%s", 1)
			fmtArgs = append(fmtArgs, "url.PathEscape("+p.goName+")")
		}
	}
	if len(fmtArgs) > 0 {
		return fmt.Sprintf("fmt.Sprintf(%q, %s)", expr, strings.Join(fmtArgs, ", "))
	}
	return fmt.Sprintf("%q", path)
}

// --- Naming utilities ---

var abbreviations = map[string]string{
	"id": "ID", "url": "URL", "api": "API", "http": "HTTP",
	"json": "JSON", "ip": "IP", "uuid": "UUID", "sql": "SQL",
	"html": "HTML", "css": "CSS", "uri": "URI", "tls": "TLS",
	"ssl": "SSL", "ssh": "SSH", "tcp": "TCP", "udp": "UDP",
	"cpu": "CPU", "gpu": "GPU", "os": "OS", "ui": "UI",
	"pl": "PL", "pnl": "PNL", "qty": "Qty", "dtbp": "DTBP",
	"pdt": "PDT", "sma": "SMA",
}

func toGoName(s string) string {
	if s == "" {
		return s
	}
	parts := splitIdent(s)
	var result strings.Builder
	for _, part := range parts {
		lower := strings.ToLower(part)
		if abbr, ok := abbreviations[lower]; ok {
			result.WriteString(abbr)
		} else {
			result.WriteString(strings.ToUpper(part[:1]) + part[1:])
		}
	}
	return result.String()
}

func splitIdent(s string) []string {
	if strings.ContainsAny(s, "_-") {
		return strings.FieldsFunc(s, func(r rune) bool {
			return r == '_' || r == '-'
		})
	}
	// camelCase split
	var parts []string
	var current strings.Builder
	for i, r := range s {
		if unicode.IsUpper(r) && i > 0 {
			if current.Len() > 0 {
				parts = append(parts, current.String())
				current.Reset()
			}
		}
		current.WriteRune(r)
	}
	if current.Len() > 0 {
		parts = append(parts, current.String())
	}
	return parts
}

func refBaseName(ref string) string {
	parts := strings.Split(ref, "/")
	return parts[len(parts)-1]
}

func mapGet(m map[string]any, key string) map[string]any {
	if m == nil {
		return nil
	}
	v, _ := m[key].(map[string]any)
	return v
}

func sortedKeys[M ~map[string]V, V any](m M) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

// --- Description generation ---

type opDesc struct {
	operationID    string
	goName         string
	summary        string
	description    string // OAS operation description, cleaned up for Long help
	params         []*flagDesc
	responseFields []*responseFieldDesc
	responseRef    string
	returnsArray   bool
}

type responseFieldDesc struct {
	name       string
	jsonType   string
	desc       string
	enumValues []string
}

type flagDesc struct {
	oasName     string
	goFieldName string
	flagName    string
	flagType    string
	defaultVal  string
	description string
	enumValues  []string
	source      string // "path", "query", or "body"
	required    bool
}

func paramToFlag(p paramInfo, source string) *flagDesc {
	desc := p.description
	if desc == "" {
		desc = humanize(p.name, p.enumValues)
	}
	fd := &flagDesc{
		oasName:     p.name,
		goFieldName: toGoName(p.name),
		flagName:    strings.ReplaceAll(p.name, "_", "-"),
		flagType:    p.goType,
		description: normalizeDesc(desc),
		enumValues:  p.enumValues,
		source:      source,
	}
	switch source {
	case "path":
		fd.required = true
	case "query":
		fd.required = p.required
		fd.defaultVal = p.defaultVal
	}
	return fd
}

func collectDescriptions(endpoints []*endpointInfo, spec map[string]any) []*opDesc {
	compSchemas := mapGet(mapGet(spec, "components"), "schemas")
	var ops []*opDesc

	for _, ep := range endpoints {
		op := &opDesc{
			operationID: ep.operationID,
			goName:      toGoName(ep.operationID),
			summary:     normalizeSummary(ep.method, ep.summary, ep.returnsArray),
			description: normalizeOASDescription(ep.description),
		}

		for _, p := range ep.queryParams {
			op.params = append(op.params, paramToFlag(p, "query"))
		}
		for _, p := range ep.pathParams {
			op.params = append(op.params, paramToFlag(p, "path"))
		}

		bodyProps := getBodyProps(ep)
		for propName, propSchema := range bodyProps {
			enums := propertyEnums(propSchema, compSchemas)
			desc := schemaDesc(propSchema, compSchemas)
			if desc == "" {
				desc = humanize(propName, enums)
			}
			op.params = append(op.params, &flagDesc{
				oasName:     propName,
				goFieldName: toGoName(propName),
				flagName:    strings.ReplaceAll(propName, "_", "-"),
				flagType:    bodyPropFlagType(propSchema, compSchemas),
				description: normalizeDesc(desc),
				enumValues:  enums,
				source:      "body",
			})
		}

		sort.Slice(op.params, func(i, j int) bool {
			return op.params[i].oasName < op.params[j].oasName
		})

		if ep.responseRef != "" {
			if s := findSchema(ep.responseRef); s != nil && s.kind == "struct" {
				fieldNames := sortedKeys(s.props)
				for _, name := range fieldNames {
					propSchema := s.props[name]
					desc := schemaDesc(propSchema, compSchemas)
					if desc == "" {
						desc = humanize(name, nil)
					}
					enums := propertyEnums(propSchema, compSchemas)
					if len(enums) == 0 {
						if items, ok := propSchema["items"].(map[string]any); ok {
							enums = propertyEnums(items, compSchemas)
						}
					}
					op.responseFields = append(op.responseFields, &responseFieldDesc{
						name:       name,
						jsonType:   oasFieldType(propSchema, compSchemas),
						desc:       normalizeDesc(desc),
						enumValues: enums,
					})
				}
			}
		}
		op.responseRef = ep.responseRef
		op.returnsArray = ep.returnsArray

		ops = append(ops, op)
	}

	return ops
}

func resolveRef(schema map[string]any, compSchemas map[string]any) map[string]any {
	if ref, ok := schema["$ref"].(string); ok && compSchemas != nil {
		if target, ok := compSchemas[refBaseName(ref)].(map[string]any); ok {
			return target
		}
	}
	return schema
}

func bodyPropFlagType(prop map[string]any, compSchemas map[string]any) string {
	return schemaToFlagType(resolveRef(prop, compSchemas))
}

func schemaToFlagType(s map[string]any) string {
	switch t, _ := s["type"].(string); t {
	case "integer":
		return "int"
	case "boolean":
		return "bool"
	default:
		return "string"
	}
}

func getBodyProps(ep *endpointInfo) map[string]map[string]any {
	if ep.bodyInline != nil {
		return extractProps(ep.bodyInline)
	}
	if ep.bodyRef != "" {
		if s := findSchema(ep.bodyRef); s != nil {
			return s.props
		}
	}
	return nil
}

func schemaDesc(schema map[string]any, compSchemas map[string]any) string {
	if desc, ok := schema["description"].(string); ok {
		return desc
	}
	desc, _ := resolveRef(schema, compSchemas)["description"].(string)
	return desc
}

func propertyEnums(prop map[string]any, compSchemas map[string]any) []string {
	return extractEnumValues(resolveRef(prop, compSchemas))
}

func extractRequired(schema map[string]any) map[string]bool {
	reqMap := map[string]bool{}
	if reqList, ok := schema["required"].([]any); ok {
		for _, r := range reqList {
			if s, ok := r.(string); ok {
				reqMap[s] = true
			}
		}
	}
	return reqMap
}

func extractDefault(schema map[string]any) string {
	dv, ok := schema["default"]
	if !ok || dv == nil {
		return ""
	}
	if arr, isArr := dv.([]any); isArr {
		strs := make([]string, 0, len(arr))
		for _, v := range arr {
			strs = append(strs, fmt.Sprint(v))
		}
		return strings.Join(strs, ",")
	}
	return fmt.Sprint(dv)
}

func extractEnumValues(schema map[string]any) []string {
	raw, ok := schema["enum"].([]any)
	if !ok {
		return nil
	}
	var vals []string
	for _, v := range raw {
		if v == nil {
			continue
		}
		if s := fmt.Sprint(v); s != "" {
			vals = append(vals, s)
		}
	}
	sort.Strings(vals)
	return vals
}

func oasFieldType(schema map[string]any, compSchemas map[string]any) string {
	if ref, ok := schema["$ref"].(string); ok {
		name := refBaseName(ref)
		if compSchemas != nil {
			if target, ok := compSchemas[name].(map[string]any); ok {
				return oasFieldType(target, nil)
			}
		}
		return name
	}
	typ, _ := schema["type"].(string)
	switch typ {
	case "string":
		if _, hasEnum := schema["enum"]; hasEnum {
			return "enum"
		}
		return "string"
	case "integer":
		return "integer"
	case "number":
		return "number"
	case "boolean":
		return "boolean"
	case "array":
		if items, ok := schema["items"].(map[string]any); ok {
			return "[]" + oasFieldType(items, compSchemas)
		}
		return "array"
	case "object":
		if ap, ok := schema["additionalProperties"].(map[string]any); ok {
			return "map[string]" + oasFieldType(ap, compSchemas)
		}
		return "object"
	default:
		return "any"
	}
}

func normalizeDesc(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return s
	}
	s = strings.TrimRight(s, ".")
	s = strings.ReplaceAll(s, "`", "")
	s = strings.TrimPrefix(s, "The ")
	if len(s) > 120 {
		s = firstSentence(s)
	}
	if len(s) > 120 {
		s = s[:117] + "..."
	}
	if len(s) > 1 && unicode.IsUpper(rune(s[0])) {
		if unicode.IsLower(rune(s[1])) {
			s = strings.ToLower(s[:1]) + s[1:]
		}
	}
	return s
}

func firstSentence(s string) string {
	if i := strings.Index(s, ". "); i > 0 {
		return s[:i]
	}
	if i := strings.Index(s, "\n"); i > 0 {
		return s[:i]
	}
	return s
}

var imperativeVerbs = map[string]bool{
	"get": true, "list": true, "create": true, "delete": true,
	"update": true, "close": true, "cancel": true, "submit": true,
	"replace": true, "remove": true, "add": true, "set": true,
	"retrieve": true, "fetch": true, "request": true, "do": true,
	"exercise": true, "return": true, "returns": true, "check": true,
	"show": true, "estimate": true, "mark": true,
}

func normalizeSummary(method, summary string, returnsArray bool) string {
	s := strings.TrimRight(strings.TrimSpace(summary), ".")
	if s == "" {
		return s
	}
	s = titleToSentence(s)

	words := strings.Fields(s)
	first := strings.ToLower(words[0])
	if !imperativeVerbs[first] {
		verb := httpMethodVerb(method, returnsArray)
		if len(s) > 1 && unicode.IsUpper(rune(s[0])) && unicode.IsUpper(rune(s[1])) {
			s = verb + " " + s
		} else {
			s = verb + " " + strings.ToLower(s[:1]) + s[1:]
		}
	}
	return s
}

// normalizeOASDescription cleans an OAS operation description for use as
// Cobra Long help text. Takes the first paragraph, strips markdown artifacts,
// and trims to a reasonable length at a sentence boundary.
func normalizeOASDescription(raw string) string {
	s := strings.TrimSpace(raw)
	if s == "" {
		return ""
	}

	// Take only the first paragraph (split on blank line).
	if i := strings.Index(s, "\n\n"); i >= 0 {
		s = s[:i]
	}

	// Strip markdown blockquote markers.
	lines := strings.Split(s, "\n")
	var cleaned []string
	for _, line := range lines {
		line = strings.TrimPrefix(line, "> ")
		line = strings.TrimPrefix(line, ">")
		if trimmed := strings.TrimSpace(line); trimmed != "" {
			cleaned = append(cleaned, trimmed)
		}
	}
	s = strings.Join(cleaned, " ")

	// Convert markdown links [text](url) to just text.
	s = stripMarkdownLinks(s)

	// Trim trailing period for consistency with Short.
	s = strings.TrimRight(s, ".")

	// Cap at ~300 chars at a sentence boundary.
	if len(s) > 300 {
		cut := 300
		if i := strings.LastIndex(s[:cut], ". "); i > 100 {
			s = s[:i]
		} else {
			s = s[:cut]
			if i := strings.LastIndex(s, " "); i > 0 {
				s = strings.TrimRight(s[:i], " .,")
			}
		}
	}

	return s
}

// stripMarkdownLinks replaces [text](url) with text.
func stripMarkdownLinks(s string) string {
	var b strings.Builder
	b.Grow(len(s))
	for i := 0; i < len(s); {
		if s[i] == '[' {
			close := strings.Index(s[i:], "](")
			if close < 0 {
				b.WriteByte(s[i])
				i++
				continue
			}
			linkText := s[i+1 : i+close]
			urlStart := i + close + 2
			urlEnd := strings.IndexByte(s[urlStart:], ')')
			if urlEnd < 0 {
				b.WriteByte(s[i])
				i++
				continue
			}
			b.WriteString(linkText)
			i = urlStart + urlEnd + 1
		} else {
			b.WriteByte(s[i])
			i++
		}
	}
	return b.String()
}

func titleToSentence(s string) string {
	words := strings.Fields(s)
	for i := 1; i < len(words); i++ {
		w := words[i]
		if strings.ToUpper(w) == w && len(w) > 1 {
			continue
		}
		words[i] = strings.ToLower(w)
	}
	return strings.Join(words, " ")
}

func httpMethodVerb(method string, returnsArray bool) string {
	switch method {
	case "GET":
		if returnsArray {
			return "List"
		}
		return "Get"
	case "POST":
		return "Create"
	case "DELETE":
		return "Delete"
	case "PATCH":
		return "Update"
	case "PUT":
		return "Update"
	default:
		return "Get"
	}
}

func humanize(name string, enumValues []string) string {
	if len(enumValues) > 0 {
		return strings.Join(enumValues, ", ")
	}
	s := strings.ReplaceAll(name, "_", " ")
	if len(s) > 0 {
		s = strings.ToUpper(s[:1]) + s[1:]
	}
	return s
}

func writeTypedDescriptionsFile(ops []*opDesc) string {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "// Code generated from api/specs; DO NOT EDIT.\n\n")
	fmt.Fprintf(&buf, "package api\n\n")
	fmt.Fprintf(&buf, "import \"sync\"\n\n")

	// Per-op vars
	for _, op := range ops {
		fmt.Fprintf(&buf, "var %sOp = Op{\n", op.goName)
		fmt.Fprintf(&buf, "\tName: %q, Summary: %q", op.goName, op.summary)
		if op.returnsArray {
			fmt.Fprintf(&buf, ", ReturnsArray: true")
		}
		fmt.Fprintf(&buf, ",\n")

		def := cmdRegistry[op.goName]
		if def.long != "" {
			fmt.Fprintf(&buf, "\tLong: %q,\n", def.long)
		} else if op.description != "" && op.description != op.summary {
			fmt.Fprintf(&buf, "\tLong: %q,\n", op.description)
		}
		if def.examples != "" {
			fmt.Fprintf(&buf, "\tExample: %s,\n", backtickQuote(def.examples))
		}

		if len(op.params) > 0 {
			seen := map[string]bool{}
			fmt.Fprintf(&buf, "\tFlags: []FlagDef{\n")
			for _, p := range op.params {
				if seen[p.oasName] {
					continue
				}
				seen[p.oasName] = true
				fmt.Fprintf(&buf, "\t\t{Name: %q, OASName: %q, Type: %q", p.flagName, p.oasName, p.flagType)
				defaultVal := p.defaultVal
				if override, ok := def.defaults[p.oasName]; ok {
					defaultVal = override
				}
				if defaultVal != "" {
					fmt.Fprintf(&buf, ", Default: %q", defaultVal)
				}
				fmt.Fprintf(&buf, ", Description: %q", p.description)
				if len(p.enumValues) > 0 {
					fmt.Fprintf(&buf, ", Completions: []string{")
					for i, v := range p.enumValues {
						if i > 0 {
							buf.WriteString(", ")
						}
						fmt.Fprintf(&buf, "%q", v)
					}
					buf.WriteString("}")
				}
				if p.required {
					fmt.Fprintf(&buf, ", Required: true")
				}
				fmt.Fprintf(&buf, ", Source: %q", p.source)
				fmt.Fprintf(&buf, "},\n")
			}
			fmt.Fprintf(&buf, "\t},\n")
		}

		fmt.Fprintf(&buf, "}\n\n")
	}

	// Deduplicate response schemas: group by responseRef
	refToOps := map[string][]*opDesc{}
	for _, op := range ops {
		if op.responseRef != "" && len(op.responseFields) > 0 {
			refToOps[op.responseRef] = append(refToOps[op.responseRef], op)
		}
	}
	sharedRefs := map[string]string{}
	for _, ref := range sortedKeys(refToOps) {
		if len(refToOps[ref]) < 2 {
			continue
		}
		varName := lcFirst(ref) + "ResponseFields"
		sharedRefs[ref] = varName
		fmt.Fprintf(&buf, "var %s = []ResponseField{\n", varName)
		writeResponseFields(&buf, refToOps[ref][0].responseFields, "\t")
		fmt.Fprintf(&buf, "}\n\n")
	}

	// Lazy-loaded ResponseSchemas
	fmt.Fprintf(&buf, "var (\n")
	fmt.Fprintf(&buf, "\tresponseSchemas     map[string][]ResponseField\n")
	fmt.Fprintf(&buf, "\tresponseSchemasOnce sync.Once\n")
	fmt.Fprintf(&buf, ")\n\n")

	fmt.Fprintf(&buf, "// ResponseSchema returns response fields for an operation (lazy-loaded).\n")
	fmt.Fprintf(&buf, "func ResponseSchema(opName string) ([]ResponseField, bool) {\n")
	fmt.Fprintf(&buf, "\tresponseSchemasOnce.Do(func() {\n")
	fmt.Fprintf(&buf, "\t\tresponseSchemas = map[string][]ResponseField{\n")
	for _, op := range ops {
		if len(op.responseFields) == 0 {
			continue
		}
		if varName, ok := sharedRefs[op.responseRef]; ok {
			fmt.Fprintf(&buf, "\t\t\t%q: %s,\n", op.goName, varName)
		} else {
			fmt.Fprintf(&buf, "\t\t\t%q: {\n", op.goName)
			writeResponseFields(&buf, op.responseFields, "\t\t\t\t")
			fmt.Fprintf(&buf, "\t\t\t},\n")
		}
	}
	fmt.Fprintf(&buf, "\t\t}\n")
	fmt.Fprintf(&buf, "\t})\n")
	fmt.Fprintf(&buf, "\tfields, ok := responseSchemas[opName]\n")
	fmt.Fprintf(&buf, "\treturn fields, ok\n")
	fmt.Fprintf(&buf, "}\n\n")

	fmt.Fprintf(&buf, "// AllOps lists every generated Op for iteration in tests and tooling.\n")
	fmt.Fprintf(&buf, "var AllOps = []Op{\n")
	for _, op := range ops {
		fmt.Fprintf(&buf, "\t%sOp,\n", op.goName)
	}
	fmt.Fprintf(&buf, "}\n\n")

	fmt.Fprintf(&buf, "var opByName map[string]Op\n\n")
	fmt.Fprintf(&buf, "func init() {\n")
	fmt.Fprintf(&buf, "\topByName = make(map[string]Op, len(AllOps))\n")
	fmt.Fprintf(&buf, "\tfor _, op := range AllOps {\n")
	fmt.Fprintf(&buf, "\t\topByName[op.Name] = op\n")
	fmt.Fprintf(&buf, "\t}\n")
	fmt.Fprintf(&buf, "}\n\n")
	fmt.Fprintf(&buf, "// OpByName returns the Op with the given name, if any.\n")
	fmt.Fprintf(&buf, "func OpByName(name string) (Op, bool) {\n")
	fmt.Fprintf(&buf, "\top, ok := opByName[name]\n")
	fmt.Fprintf(&buf, "\treturn op, ok\n")
	fmt.Fprintf(&buf, "}\n")

	return buf.String()
}

func writeResponseFields(buf *bytes.Buffer, fields []*responseFieldDesc, indent string) {
	for _, f := range fields {
		fmt.Fprintf(buf, "%s{Name: %q, Type: %q, Description: %q", indent, f.name, f.jsonType, f.desc)
		if len(f.enumValues) > 0 {
			fmt.Fprintf(buf, ", EnumValues: []string{")
			for i, v := range f.enumValues {
				if i > 0 {
					buf.WriteString(", ")
				}
				fmt.Fprintf(buf, "%q", v)
			}
			buf.WriteString("}")
		}
		fmt.Fprintf(buf, "},\n")
	}
}

func lcFirst(s string) string {
	if s == "" {
		return s
	}
	r := []rune(s)
	r[0] = unicode.ToLower(r[0])
	return string(r)
}

// --- File writing ---

func writeGo(path, content string) {
	formatted, err := format.Source([]byte(content))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Warning: gofmt failed for %s: %v\nWriting unformatted.\n", filepath.Base(path), err)
		formatted = []byte(content)
	}
	if err := os.WriteFile(path, formatted, 0o644); err != nil {
		log.Fatalf("writing %s: %v", path, err)
	}
}
