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

func main() {
	root := findProjectRoot()
	outDir := filepath.Join(root, "internal", "api")
	if err := os.MkdirAll(outDir, 0755); err != nil {
		log.Fatalf("creating output dir: %v", err)
	}

	tSpec := loadSpec(filepath.Join(root, "api", "specs", "trading-api.json"))
	mSpec := loadSpec(filepath.Join(root, "api", "specs", "market-data-api.json"))

	tSchemas := extractSchemas(tSpec)
	mSchemas := extractSchemas(mSpec)
	tEndpoints := extractEndpoints(tSpec)
	mEndpoints := extractEndpoints(mSpec)

	dedup(tSchemas)
	dedup(mSchemas)

	writeGo(filepath.Join(outDir, "trading_types.go"), genTypes("trading-api.json", tSchemas))
	writeGo(filepath.Join(outDir, "trading_client.go"), genClient("Trading", "trading-api.json", tSchemas, tEndpoints, false))
	writeGo(filepath.Join(outDir, "marketdata_types.go"), genTypes("market-data-api.json", mSchemas))
	writeGo(filepath.Join(outDir, "marketdata_client.go"), genClient("MarketData", "market-data-api.json", mSchemas, mEndpoints, true))

	fmt.Printf("Generated %d trading types, %d market data types\n", len(tSchemas), len(mSchemas))
	fmt.Printf("Generated %d trading endpoints, %d market data endpoints\n", len(tEndpoints), len(mEndpoints))
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
		props, hasProps := s["properties"].(map[string]any)
		enumVals, hasEnum := s["enum"].([]any)

		if hasEnum {
			info.kind = "enum"
			for _, v := range enumVals {
				if v == nil {
					continue
				}
				sv := fmt.Sprint(v)
				if sv == "" {
					continue
				}
				info.enumValues = append(info.enumValues, sv)
			}
			sort.Strings(info.enumValues)
		} else if hasProps {
			info.kind = "struct"
			info.props = make(map[string]map[string]any)
			for fn, fs := range props {
				if m, ok := fs.(map[string]any); ok {
					info.props[fn] = m
				}
			}
			info.required = make(map[string]bool)
			if reqList, ok := s["required"].([]any); ok {
				for _, r := range reqList {
					if str, ok := r.(string); ok {
						info.required[str] = true
					}
				}
			}
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
	for _, s := range schemas {
		if idx, ok := seen[s.goName]; ok {
			prev := schemas[idx]
			if unicode.IsUpper(rune(prev.name[0])) {
				s.goName = s.goName + "V3"
			} else {
				prev.goName = prev.goName + "V3"
			}
		}
		seen[s.goName] = indexOf(schemas, s)
	}
}

func indexOf(schemas []*schemaInfo, target *schemaInfo) int {
	for i, s := range schemas {
		if s == target {
			return i
		}
	}
	return -1
}

// --- Endpoint extraction ---

type endpointInfo struct {
	method       string
	path         string
	operationID  string
	summary      string
	goName       string
	pathParams   []paramInfo
	queryParams  []paramInfo
	bodyRef      string
	bodyInline   map[string]any
	responseRef  string
	returnsArray bool
}

type paramInfo struct {
	name       string
	goName     string
	goType     string
	required   bool
	enumValues []string
	defaultVal string
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

			ep := &endpointInfo{
				method:      strings.ToUpper(method),
				path:        path,
				operationID: opID,
				summary:     summary,
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
						continue
					}
				}
				name, _ := p["name"].(string)
				in, _ := p["in"].(string)
				req, _ := p["required"].(bool)
				pSchema, _ := p["schema"].(map[string]any)
				goType := "string"
				if pSchema != nil {
					switch t, _ := pSchema["type"].(string); t {
					case "integer":
						goType = "int"
					case "boolean":
						goType = "bool"
					case "number":
						goType = "float64"
					}
				}
			pi := paramInfo{name: name, goName: toGoFieldName(name), goType: goType, required: req}
			if pSchema != nil {
				if ev, ok := pSchema["enum"].([]any); ok {
					for _, v := range ev {
						if v == nil {
							continue
						}
						sv := fmt.Sprint(v)
						if sv != "" {
							pi.enumValues = append(pi.enumValues, sv)
						}
					}
					sort.Strings(pi.enumValues)
				}
				if dv, ok := pSchema["default"]; ok && dv != nil {
					pi.defaultVal = fmt.Sprint(dv)
				}
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

			if responses, ok := op["responses"].(map[string]any); ok {
				for code, respRaw := range responses {
					if len(code) == 0 || code[0] != '2' {
						continue
					}
					resp, ok := respRaw.(map[string]any)
					if !ok {
						continue
					}
					if content, ok := resp["content"].(map[string]any); ok {
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
			genEnum(&buf, s)
		}
	}
	for _, s := range schemas {
		if s.kind == "struct" {
			genStruct(&buf, s, schemas)
		}
	}
	for _, s := range schemas {
		if s.kind == "struct" {
			genInlineEnumValues(&buf, s)
		}
	}
	return buf.String()
}

func genInlineEnumValues(buf *bytes.Buffer, s *schemaInfo) {
	var fields []string
	for f := range s.props {
		fields = append(fields, f)
	}
	sort.Strings(fields)

	for _, fieldName := range fields {
		fieldSchema := s.props[fieldName]
		enumVals, ok := fieldSchema["enum"].([]any)
		if !ok || len(enumVals) == 0 {
			continue
		}

		var vals []string
		for _, v := range enumVals {
			if v == nil {
				continue
			}
			sv := fmt.Sprint(v)
			if sv == "" {
				continue
			}
			vals = append(vals, sv)
		}
		if len(vals) == 0 {
			continue
		}
		sort.Strings(vals)

		varName := s.goName + toGoFieldName(fieldName)
		genValuesSlice(buf, varName, vals)
	}
}

func genEnum(buf *bytes.Buffer, s *schemaInfo) {
	fmt.Fprintf(buf, "type %s string\n\nconst (\n", s.goName)
	for _, val := range s.enumValues {
		constName := s.goName + sanitizeConstName(val)
		fmt.Fprintf(buf, "\t%s %s = %q\n", constName, s.goName, val)
	}
	fmt.Fprintf(buf, ")\n\n")

	genValuesSlice(buf, s.goName, s.enumValues)
}

func genValuesSlice(buf *bytes.Buffer, name string, vals []string) {
	fmt.Fprintf(buf, "var %sValues = []string{", name)
	for i, val := range vals {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(buf, "%q", val)
	}
	buf.WriteString("}\n\n")
}

func sanitizeConstName(val string) string {
	cleaned := strings.NewReplacer("/", "x", ".", "dot", "+", "Plus", " ", "").Replace(val)
	name := toGoName(cleaned)
	if name == "" {
		return "Empty"
	}
	if unicode.IsDigit(rune(name[0])) {
		name = "V" + name
	}
	return name
}

func genStruct(buf *bytes.Buffer, s *schemaInfo, allSchemas []*schemaInfo) {
	fmt.Fprintf(buf, "type %s struct {\n", s.goName)
	var fields []string
	for f := range s.props {
		fields = append(fields, f)
	}
	sort.Strings(fields)

	for _, fieldName := range fields {
		fieldSchema := s.props[fieldName]
		goField := toGoFieldName(fieldName)
		goType := resolveGoType(fieldSchema, allSchemas)
		tag := fieldName
		if !s.required[fieldName] {
			tag += ",omitempty"
		}
		fmt.Fprintf(buf, "\t%s %s `json:%q`\n", goField, goType, tag)
	}
	fmt.Fprintf(buf, "}\n\n")
}

func resolveGoType(schema map[string]any, allSchemas []*schemaInfo) string {
	if ref, ok := schema["$ref"].(string); ok {
		rn := refBaseName(ref)
		for _, s := range allSchemas {
			if s.name == rn {
				if s.kind == "alias" {
					return resolveGoType(s.raw, allSchemas)
				}
				return s.goName
			}
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
			return "[]" + resolveGoType(items, allSchemas)
		}
		return "[]any"
	case "object":
		if addProps, ok := schema["additionalProperties"].(map[string]any); ok {
			return "map[string]" + resolveGoType(addProps, allSchemas)
		}
		return "map[string]any"
	default:
		return "any"
	}
}

// --- Client generation ---

func genClient(clientName, specFile string, schemas []*schemaInfo, endpoints []*endpointInfo, isData bool) string {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "// Code generated from api/specs/%s; DO NOT EDIT.\n\n", specFile)
	fmt.Fprintf(&buf, "package api\n\n")
	needsStrings := isData
	if !needsStrings {
		for _, ep := range endpoints {
			if ep.bodyRef != "" || ep.bodyInline != nil {
				needsStrings = true
				break
			}
		}
	}

	fmt.Fprintf(&buf, "import (\n")
	fmt.Fprintf(&buf, "\t\"encoding/json\"\n")
	fmt.Fprintf(&buf, "\t\"fmt\"\n")
	fmt.Fprintf(&buf, "\t\"net/url\"\n")
	if needsStrings {
		fmt.Fprintf(&buf, "\t\"strings\"\n")
	}
	fmt.Fprintf(&buf, "\n\t\"github.com/alpacahq/cli/internal/client\"\n")
	fmt.Fprintf(&buf, ")\n\n")

	fmt.Fprintf(&buf, "// %sClient provides typed methods for the %s API.\n", clientName, clientName)
	fmt.Fprintf(&buf, "type %sClient struct {\n\tRaw *client.Client\n}\n\n", clientName)
	fmt.Fprintf(&buf, "func New%sClient(raw *client.Client) *%sClient {\n", clientName, clientName)
	fmt.Fprintf(&buf, "\treturn &%sClient{Raw: raw}\n}\n\n", clientName)

	getMethod := "Raw.Get"
	if isData {
		getMethod = "Raw.GetData"
	}

	for _, ep := range endpoints {
		genEndpointMethod(&buf, ep, clientName, getMethod, schemas, isData)
	}

	validated := map[string]bool{}
	for _, ep := range endpoints {
		if ep.bodyRef == "" || validated[ep.bodyRef] {
			continue
		}
		for _, s := range schemas {
			if s.goName == ep.bodyRef && s.kind == "struct" && len(s.required) > 0 {
				var reqFields []string
				for fn, isReq := range s.required {
					if !isReq {
						continue
					}
					fs, ok := s.props[fn]
					if !ok {
						continue
					}
					if resolveGoType(fs, schemas) == "string" {
						reqFields = append(reqFields, fn)
					}
				}
				sort.Strings(reqFields)
				genValidateMethod(&buf, s.goName, reqFields)
				validated[ep.bodyRef] = true
				break
			}
		}
	}

	if isData {
		genUnifiedMethods(&buf, clientName)
	}

	var mutating []string
	for _, ep := range endpoints {
		switch ep.method {
		case "POST", "PUT", "PATCH", "DELETE":
			mutating = append(mutating, ep.goName)
		}
	}
	if len(mutating) > 0 {
		fmt.Fprintf(&buf, "var %sMutatingMethods = map[string]bool{\n", clientName)
		for _, m := range mutating {
			fmt.Fprintf(&buf, "\t%q: true,\n", m)
		}
		fmt.Fprintf(&buf, "}\n\n")
	}

	return buf.String()
}

func genEndpointMethod(buf *bytes.Buffer, ep *endpointInfo, clientName, getMethod string, schemas []*schemaInfo, isData bool) {
	hasParams := len(ep.queryParams) > 0
	paramsTypeName := ep.goName + "Params"

	// If response type references an alias schema, clear it to use json.RawMessage
	if ep.responseRef != "" && !schemaHasType(schemas, ep.responseRef) {
		ep.responseRef = ""
	}

	if hasParams {
		fmt.Fprintf(buf, "type %s struct {\n", paramsTypeName)
		for _, p := range ep.queryParams {
			if p.defaultVal != "" {
				fmt.Fprintf(buf, "\t%s %s // default: %s\n", p.goName, p.goType, p.defaultVal)
			} else {
				fmt.Fprintf(buf, "\t%s %s\n", p.goName, p.goType)
			}
		}
		fmt.Fprintf(buf, "}\n\n")

		fmt.Fprintf(buf, "func (p *%s) Values() url.Values {\n", paramsTypeName)
		fmt.Fprintf(buf, "\tif p == nil { return nil }\n")
		fmt.Fprintf(buf, "\tv := url.Values{}\n")
		for _, p := range ep.queryParams {
			switch p.goType {
			case "string":
				fmt.Fprintf(buf, "\tif p.%s != \"\" { v.Set(%q, p.%s) }\n", p.goName, p.name, p.goName)
			case "int":
				fmt.Fprintf(buf, "\tif p.%s != 0 { v.Set(%q, fmt.Sprint(p.%s)) }\n", p.goName, p.name, p.goName)
			case "bool":
				fmt.Fprintf(buf, "\tif p.%s { v.Set(%q, \"true\") }\n", p.goName, p.name)
			case "float64":
				fmt.Fprintf(buf, "\tif p.%s != 0 { v.Set(%q, fmt.Sprintf(\"%%g\", p.%s)) }\n", p.goName, p.name, p.goName)
			}
		}
		fmt.Fprintf(buf, "\treturn v\n}\n\n")

		for _, p := range ep.queryParams {
			if len(p.enumValues) > 0 {
				genValuesSlice(buf, paramsTypeName+p.goName, p.enumValues)
			}
		}
	}

	bodyType := ep.bodyRef
	if bodyType == "" && ep.bodyInline != nil {
		bodyType = ep.goName + "Request"
		if !schemaExists(schemas, bodyType) {
			genInlineRequestStruct(buf, bodyType, ep.bodyInline, schemas)
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
		args = append(args, "params *"+paramsTypeName)
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
	fmt.Fprintf(buf, "\tpath := %s\n", pathExpr)

	switch ep.method {
	case "GET":
		if hasParams {
			fmt.Fprintf(buf, "\tdata, err := c.%s(path, params.Values())\n", getMethod)
		} else {
			fmt.Fprintf(buf, "\tdata, err := c.%s(path, nil)\n", getMethod)
		}
	case "DELETE":
		if hasParams {
			fmt.Fprintf(buf, "\tdata, err := c.Raw.Delete(path, params.Values())\n")
		} else {
			fmt.Fprintf(buf, "\tdata, err := c.Raw.Delete(path, nil)\n")
		}
	case "POST":
		if bodyType != "" {
			fmt.Fprintf(buf, "\tdata, err := c.Raw.Post(path, body)\n")
		} else {
			fmt.Fprintf(buf, "\tdata, err := c.Raw.Post(path, nil)\n")
		}
	case "PUT":
		if bodyType != "" {
			fmt.Fprintf(buf, "\tdata, err := c.Raw.Put(path, body)\n")
		} else {
			fmt.Fprintf(buf, "\tdata, err := c.Raw.Put(path, nil)\n")
		}
	case "PATCH":
		if bodyType != "" {
			fmt.Fprintf(buf, "\tdata, err := c.Raw.Patch(path, body)\n")
		} else {
			fmt.Fprintf(buf, "\tdata, err := c.Raw.Patch(path, nil)\n")
		}
	}

	fmt.Fprintf(buf, "\tif err != nil { return nil, err }\n")

	if respType != "" {
		if ep.returnsArray {
			fmt.Fprintf(buf, "\tvar result []%s\n", respType)
		} else {
			fmt.Fprintf(buf, "\tvar result %s\n", respType)
		}
		fmt.Fprintf(buf, "\treturn ")
		if !ep.returnsArray {
			fmt.Fprintf(buf, "&")
		}
		fmt.Fprintf(buf, "result, json.Unmarshal(data, &result)\n")
	} else {
		fmt.Fprintf(buf, "\treturn data, nil\n")
	}

	fmt.Fprintf(buf, "}\n\n")
}

func genInlineRequestStruct(buf *bytes.Buffer, name string, schema map[string]any, allSchemas []*schemaInfo) {
	props, _ := schema["properties"].(map[string]any)
	if props == nil {
		return
	}
	reqMap := map[string]bool{}
	if reqList, ok := schema["required"].([]any); ok {
		for _, r := range reqList {
			if s, ok := r.(string); ok {
				reqMap[s] = true
			}
		}
	}

	fmt.Fprintf(buf, "type %s struct {\n", name)
	var fields []string
	for f := range props {
		fields = append(fields, f)
	}
	sort.Strings(fields)
	for _, fn := range fields {
		fs, ok := props[fn].(map[string]any)
		if !ok {
			continue
		}
		goField := toGoFieldName(fn)
		goType := resolveGoType(fs, allSchemas)
		tag := fn
		if !reqMap[fn] {
			tag += ",omitempty"
		}
		fmt.Fprintf(buf, "\t%s %s `json:%q`\n", goField, goType, tag)
	}
	fmt.Fprintf(buf, "}\n\n")

	var reqFields []string
	for fn := range reqMap {
		p, ok := props[fn]
		if !ok {
			continue
		}
		fs, ok := p.(map[string]any)
		if !ok {
			continue
		}
		if t, _ := fs["type"].(string); t == "string" {
			reqFields = append(reqFields, fn)
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
		fmt.Fprintf(buf, "\tif r.%s == \"\" { missing = append(missing, %q) }\n", toGoFieldName(fn), fn)
	}
	fmt.Fprintf(buf, "\tif len(missing) > 0 {\n")
	fmt.Fprintf(buf, "\t\treturn fmt.Errorf(\"missing required fields: %%s\", strings.Join(missing, \", \"))\n")
	fmt.Fprintf(buf, "\t}\n")
	fmt.Fprintf(buf, "\treturn nil\n")
	fmt.Fprintf(buf, "}\n\n")
}

func schemaExists(schemas []*schemaInfo, goName string) bool {
	for _, s := range schemas {
		if s.goName == goName {
			return true
		}
	}
	return false
}

func schemaHasType(schemas []*schemaInfo, goName string) bool {
	for _, s := range schemas {
		if s.goName == goName && (s.kind == "struct" || s.kind == "enum") {
			return true
		}
	}
	return false
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
			fmtArgs = append(fmtArgs, p.goName)
		}
	}
	if len(fmtArgs) > 0 {
		return fmt.Sprintf("fmt.Sprintf(%q, %s)", expr, strings.Join(fmtArgs, ", "))
	}
	return fmt.Sprintf("%q", path)
}

// --- Unified stock/crypto convenience methods ---

type unifiedMethod struct {
	name       string
	stockPath  string // uses %s for symbol
	cryptoPath string // symbol goes in query params
}

var unifiedMethods = []unifiedMethod{
	{"Bars", "/v2/stocks/%s/bars", "/v1beta3/crypto/us/bars"},
	{"Quotes", "/v2/stocks/%s/quotes", "/v1beta3/crypto/us/quotes"},
	{"Trades", "/v2/stocks/%s/trades", "/v1beta3/crypto/us/trades"},
	{"Snapshot", "/v2/stocks/%s/snapshot", "/v1beta3/crypto/us/snapshots"},
	{"LatestBar", "/v2/stocks/%s/bars/latest", "/v1beta3/crypto/us/latest/bars"},
	{"LatestQuote", "/v2/stocks/%s/quotes/latest", "/v1beta3/crypto/us/latest/quotes"},
	{"LatestTrade", "/v2/stocks/%s/trades/latest", "/v1beta3/crypto/us/latest/trades"},
}

func genUnifiedMethods(buf *bytes.Buffer, clientName string) {
	for _, um := range unifiedMethods {
		fmt.Fprintf(buf, "// %s routes to the stock or crypto endpoint based on symbol format.\n", um.name)
		fmt.Fprintf(buf, "func (c *%sClient) %s(symbol string, params url.Values) (json.RawMessage, error) {\n", clientName, um.name)
		fmt.Fprintf(buf, "\tif params == nil { params = url.Values{} }\n")
		fmt.Fprintf(buf, "\tif strings.Contains(symbol, \"/\") {\n")
		fmt.Fprintf(buf, "\t\tparams.Set(\"symbols\", symbol)\n")
		fmt.Fprintf(buf, "\t\treturn c.Raw.GetData(%q, params)\n", um.cryptoPath)
		fmt.Fprintf(buf, "\t}\n")
		fmt.Fprintf(buf, "\treturn c.Raw.GetData(fmt.Sprintf(%q, symbol), params)\n", um.stockPath)
		fmt.Fprintf(buf, "}\n\n")
	}
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

func toGoFieldName(s string) string {
	return toGoName(s)
}

func splitIdent(s string) []string {
	if strings.ContainsAny(s, "_-") {
		raw := strings.FieldsFunc(s, func(r rune) bool {
			return r == '_' || r == '-'
		})
		var parts []string
		for _, p := range raw {
			if p != "" {
				parts = append(parts, p)
			}
		}
		return parts
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

// --- File writing ---

func writeGo(path, content string) {
	formatted, err := format.Source([]byte(content))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Warning: gofmt failed for %s: %v\nWriting unformatted.\n", filepath.Base(path), err)
		formatted = []byte(content)
	}
	if err := os.WriteFile(path, formatted, 0644); err != nil {
		log.Fatalf("writing %s: %v", path, err)
	}
}
