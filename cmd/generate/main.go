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
	if err := os.MkdirAll(outDir, 0o755); err != nil {
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

	var ops []*opDesc
	ops = append(ops, collectDescriptions(tEndpoints, tSchemas, tSpec)...)
	ops = append(ops, collectDescriptions(mEndpoints, mSchemas, mSpec)...)
	sort.Slice(ops, func(i, j int) bool {
		return ops[i].operationID < ops[j].operationID
	})
	writeGo(filepath.Join(outDir, "descriptions.go"), writeTypedDescriptionsFile(ops))

	var allEndpoints []*endpointInfo
	allEndpoints = append(allEndpoints, tEndpoints...)
	allEndpoints = append(allEndpoints, mEndpoints...)

	var allSchemas []*schemaInfo
	allSchemas = append(allSchemas, tSchemas...)
	allSchemas = append(allSchemas, mSchemas...)

	cmdDir := filepath.Join(root, "internal", "cmd")
	writeGo(filepath.Join(cmdDir, "params_generated.go"), genFromFlags(allEndpoints, allSchemas))

	fmt.Printf("Generated %d trading types, %d market data types\n", len(tSchemas), len(mSchemas))
	fmt.Printf("Generated %d trading endpoints, %d market data endpoints\n", len(tEndpoints), len(mEndpoints))
	nParams := 0
	for _, op := range ops {
		nParams += len(op.params)
	}
	nFromFlags := 0
	for _, ep := range allEndpoints {
		if len(ep.queryParams) > 0 {
			nFromFlags++
		}
	}
	fmt.Printf("Generated %d operations, %d param descriptions, %d FromFlags functions\n", len(ops), nParams, nFromFlags)
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
						fmt.Fprintf(os.Stderr, "Warning: unresolved parameter $ref %q in %s %s\n", ref, method, path)
						continue
					}
				}
				name, _ := p["name"].(string)
				in, _ := p["in"].(string)
				req, _ := p["required"].(bool)
				pSchema, _ := p["schema"].(map[string]any)
				if pSchema != nil {
					if ref, ok := pSchema["$ref"].(string); ok && compSchemas != nil {
						if resolved, ok := compSchemas[refBaseName(ref)].(map[string]any); ok {
							pSchema = resolved
						}
					}
				}
				goType := "string"
				if pSchema != nil {
					switch t, _ := pSchema["type"].(string); t {
					case "integer":
						goType = "int"
					case "boolean":
						goType = "bool"
					}
				}
				pi := paramInfo{name: name, goName: toGoFieldName(name), goType: goType, required: req}
				pi.description, _ = p["description"].(string)
				if pi.description == "" && pSchema != nil {
					pi.description = schemaDesc(pSchema, compSchemas)
				}
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
						if arr, isArr := dv.([]any); isArr {
							strs := make([]string, 0, len(arr))
							for _, v := range arr {
								strs = append(strs, fmt.Sprint(v))
							}
							pi.defaultVal = strings.Join(strs, ",")
						} else {
							pi.defaultVal = fmt.Sprint(dv)
						}
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
		genUnifiedMethods(&buf, clientName, endpoints)
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
		paramsArg := "nil"
		if hasParams {
			paramsArg = "params.Values()"
		}
		if bodyType != "" {
			fmt.Fprintf(buf, "\tdata, err := c.Raw.Post(path, %s, body)\n", paramsArg)
		} else {
			fmt.Fprintf(buf, "\tdata, err := c.Raw.Post(path, %s, nil)\n", paramsArg)
		}
	case "PUT":
		paramsArg := "nil"
		if hasParams {
			paramsArg = "params.Values()"
		}
		if bodyType != "" {
			fmt.Fprintf(buf, "\tdata, err := c.Raw.Put(path, %s, body)\n", paramsArg)
		} else {
			fmt.Fprintf(buf, "\tdata, err := c.Raw.Put(path, %s, nil)\n", paramsArg)
		}
	case "PATCH":
		paramsArg := "nil"
		if hasParams {
			paramsArg = "params.Values()"
		}
		if bodyType != "" {
			fmt.Fprintf(buf, "\tdata, err := c.Raw.Patch(path, %s, body)\n", paramsArg)
		} else {
			fmt.Fprintf(buf, "\tdata, err := c.Raw.Patch(path, %s, nil)\n", paramsArg)
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
			fmtArgs = append(fmtArgs, "url.PathEscape("+p.goName+")")
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

// deriveUnifiedMethods finds matching stock single-symbol GET endpoints
// (/v2/stocks/{symbol}/...) and crypto GET endpoints (/v1beta3/crypto/{loc}/...)
// to generate convenience methods that route based on symbol format.
func deriveUnifiedMethods(endpoints []*endpointInfo) []unifiedMethod {
	stockSingle := map[string]string{}
	cryptoEntities := map[string]string{}

	for _, ep := range endpoints {
		if ep.method != "GET" {
			continue
		}
		const stockPrefix = "/v2/stocks/{symbol}/"
		const cryptoPrefix = "/v1beta3/crypto/{loc}/"
		if strings.HasPrefix(ep.path, stockPrefix) {
			entity := ep.path[len(stockPrefix):]
			stockSingle[entity] = ep.path
		} else if strings.HasPrefix(ep.path, cryptoPrefix) {
			entity := ep.path[len(cryptoPrefix):]
			cryptoEntities[entity] = ep.path
		}
	}

	var methods []unifiedMethod
	for stockEntity, stockPath := range stockSingle {
		cryptoEntity := findCryptoMatch(stockEntity, cryptoEntities)
		if cryptoEntity == "" {
			continue
		}
		name := entityToMethodName(stockEntity)
		sp := strings.Replace(stockPath, "{symbol}", "%s", 1)
		cp := strings.Replace(cryptoEntities[cryptoEntity], "{loc}", "us", 1)
		methods = append(methods, unifiedMethod{name: name, stockPath: sp, cryptoPath: cp})
	}

	sort.Slice(methods, func(i, j int) bool { return methods[i].name < methods[j].name })
	return methods
}

// findCryptoMatch finds the crypto entity that corresponds to a stock entity.
// Stock "bars" matches crypto "bars"; stock "bars/latest" matches crypto "latest/bars";
// stock "snapshot" matches crypto "snapshots" (plural).
func findCryptoMatch(stockEntity string, cryptoEntities map[string]string) string {
	if _, ok := cryptoEntities[stockEntity]; ok {
		return stockEntity
	}
	if _, ok := cryptoEntities[stockEntity+"s"]; ok {
		return stockEntity + "s"
	}
	// Stock "X/latest" → crypto "latest/X"
	if strings.HasSuffix(stockEntity, "/latest") {
		base := strings.TrimSuffix(stockEntity, "/latest")
		reversed := "latest/" + base
		if _, ok := cryptoEntities[reversed]; ok {
			return reversed
		}
	}
	return ""
}

// entityToMethodName converts a path entity like "bars", "bars/latest", "snapshot"
// into a Go method name like "Bars", "LatestBar", "Snapshot".
func entityToMethodName(entity string) string {
	if strings.HasSuffix(entity, "/latest") {
		base := strings.TrimSuffix(entity, "/latest")
		base = strings.TrimSuffix(base, "s")
		return "Latest" + toGoName(base)
	}
	return toGoName(entity)
}

func genUnifiedMethods(buf *bytes.Buffer, clientName string, endpoints []*endpointInfo) {
	methods := deriveUnifiedMethods(endpoints)
	for _, um := range methods {
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

// --- Description generation ---

type opDesc struct {
	operationID    string
	goName         string
	summary        string
	params         []*flagDesc
	responseFields []*responseFieldDesc
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
	isPathParam bool
	required    bool
}

func collectDescriptions(endpoints []*endpointInfo, schemas []*schemaInfo, spec map[string]any) []*opDesc {
	compSchemas := mapGet(mapGet(spec, "components"), "schemas")
	var ops []*opDesc

	for _, ep := range endpoints {
		op := &opDesc{
			operationID: ep.operationID,
			goName:      toGoName(ep.operationID),
			summary:     normalizeSummary(ep.method, ep.summary, ep.returnsArray),
		}

		for _, p := range ep.queryParams {
			desc := p.description
			if desc == "" {
				desc = humanize(p.name, p.enumValues)
			}
			op.params = append(op.params, &flagDesc{
				oasName:     p.name,
				goFieldName: toGoFieldName(p.name),
				flagName:    strings.ReplaceAll(p.name, "_", "-"),
				flagType:    goTypeToFlagType(p.goType),
				defaultVal:  p.defaultVal,
				description: normalizeDesc(desc),
				enumValues:  p.enumValues,
				required:    p.required,
			})
		}

		for _, p := range ep.pathParams {
			desc := p.description
			if desc == "" {
				desc = humanize(p.name, p.enumValues)
			}
			op.params = append(op.params, &flagDesc{
				oasName:     p.name,
				goFieldName: toGoFieldName(p.name),
				flagName:    strings.ReplaceAll(p.name, "_", "-"),
				flagType:    goTypeToFlagType(p.goType),
				description: normalizeDesc(desc),
				isPathParam: true,
			})
		}

		bodyProps := getBodyProps(ep, schemas)
		for propName, propSchema := range bodyProps {
			desc := propertyDesc(propSchema, compSchemas)
			if desc == "" {
				desc = humanize(propName, propertyEnums(propSchema, compSchemas))
			}
			op.params = append(op.params, &flagDesc{
				oasName:     propName,
				goFieldName: toGoFieldName(propName),
				flagName:    strings.ReplaceAll(propName, "_", "-"),
				flagType:    bodyPropFlagType(propSchema, compSchemas),
				description: normalizeDesc(desc),
				enumValues:  propertyEnums(propSchema, compSchemas),
			})
		}

		sort.Slice(op.params, func(i, j int) bool {
			return op.params[i].oasName < op.params[j].oasName
		})

		if ep.responseRef != "" {
			for _, s := range schemas {
				if s.goName == ep.responseRef && s.kind == "struct" {
					var fieldNames []string
					for name := range s.props {
						fieldNames = append(fieldNames, name)
					}
					sort.Strings(fieldNames)
					for _, name := range fieldNames {
						propSchema := s.props[name]
						desc := propertyDesc(propSchema, compSchemas)
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
					break
				}
			}
		}
		op.returnsArray = ep.returnsArray

		ops = append(ops, op)
	}

	return ops
}

func goTypeToFlagType(goType string) string {
	switch goType {
	case "int":
		return "int"
	case "bool":
		return "bool"
	default:
		return "string"
	}
}

func bodyPropFlagType(prop map[string]any, compSchemas map[string]any) string {
	if ref, ok := prop["$ref"].(string); ok && compSchemas != nil {
		name := refBaseName(ref)
		if target, ok := compSchemas[name].(map[string]any); ok {
			if _, hasEnum := target["enum"].([]any); hasEnum {
				return "string"
			}
			return goTypeToFlagType(schemaToFlagType(target))
		}
	}
	return goTypeToFlagType(schemaToFlagType(prop))
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

func getBodyProps(ep *endpointInfo, schemas []*schemaInfo) map[string]map[string]any {
	if ep.bodyInline != nil {
		raw, _ := ep.bodyInline["properties"].(map[string]any)
		result := make(map[string]map[string]any, len(raw))
		for k, v := range raw {
			if m, ok := v.(map[string]any); ok {
				result[k] = m
			}
		}
		return result
	}
	if ep.bodyRef != "" {
		for _, s := range schemas {
			if s.goName == ep.bodyRef {
				return s.props
			}
		}
	}
	return nil
}

func schemaDesc(schema map[string]any, compSchemas map[string]any) string {
	if desc, ok := schema["description"].(string); ok {
		return desc
	}
	if ref, ok := schema["$ref"].(string); ok && compSchemas != nil {
		name := refBaseName(ref)
		if target, ok := compSchemas[name].(map[string]any); ok {
			if desc, ok := target["description"].(string); ok {
				return desc
			}
		}
	}
	return ""
}

func propertyDesc(prop map[string]any, compSchemas map[string]any) string {
	if desc, ok := prop["description"].(string); ok {
		return desc
	}
	if ref, ok := prop["$ref"].(string); ok && compSchemas != nil {
		name := refBaseName(ref)
		if target, ok := compSchemas[name].(map[string]any); ok {
			if desc, ok := target["description"].(string); ok {
				return desc
			}
		}
	}
	return ""
}

func propertyEnums(prop map[string]any, compSchemas map[string]any) []string {
	vals := extractEnumValues(prop)
	if len(vals) > 0 {
		return vals
	}
	if ref, ok := prop["$ref"].(string); ok && compSchemas != nil {
		name := refBaseName(ref)
		if target, ok := compSchemas[name].(map[string]any); ok {
			return extractEnumValues(target)
		}
	}
	return nil
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

	fmt.Fprintf(&buf, "// Op is satisfied by every generated operation variable (e.g. GetAccountOp).\n")
	fmt.Fprintf(&buf, "// Use it to pass operations type-safely instead of raw strings.\n")
	fmt.Fprintf(&buf, "type Op interface {\n")
	fmt.Fprintf(&buf, "\tSummary() string\n")
	fmt.Fprintf(&buf, "\tFlags() []FlagDef\n")
	fmt.Fprintf(&buf, "\tRequiredFlags() []string\n")
	fmt.Fprintf(&buf, "\tResponseFields() []ResponseField\n")
	fmt.Fprintf(&buf, "}\n\n")

	fmt.Fprintf(&buf, "// FlagDef describes a CLI flag derived from the OpenAPI spec.\n")
	fmt.Fprintf(&buf, "type FlagDef struct {\n")
	fmt.Fprintf(&buf, "\tName        string   // kebab-case CLI flag name\n")
	fmt.Fprintf(&buf, "\tOASName     string   // original OAS property/parameter name\n")
	fmt.Fprintf(&buf, "\tType        string   // \"string\", \"bool\", \"int\"\n")
	fmt.Fprintf(&buf, "\tDefault     string\n")
	fmt.Fprintf(&buf, "\tDescription string\n")
	fmt.Fprintf(&buf, "\tCompletions []string // enum values for shell completion\n")
	fmt.Fprintf(&buf, "\tOpName      string   // operation name for schema lookup\n")
	fmt.Fprintf(&buf, "\tRequired    bool     // true if OAS marks this parameter as required\n")
	fmt.Fprintf(&buf, "}\n\n")

	for _, op := range ops {
		typeName := lcFirst(op.goName) + "Op"

		fmt.Fprintf(&buf, "type %s struct{}\n\n", typeName)

		fmt.Fprintf(&buf, "var %sOp = %s{}\n\n", op.goName, typeName)

		fmt.Fprintf(&buf, "func (o %s) Summary() string {\n", typeName)
		fmt.Fprintf(&buf, "\treturn %q\n", op.summary)
		fmt.Fprintf(&buf, "}\n\n")

		fmt.Fprintf(&buf, "func (o %s) ResponseFields() []ResponseField {\n", typeName)
		fmt.Fprintf(&buf, "\treturn ResponseSchemas[%q]\n", op.goName)
		fmt.Fprintf(&buf, "}\n\n")

		var reqFlags []string
		reqSeen := map[string]bool{}
		for _, p := range op.params {
			if p.required && !p.isPathParam && !reqSeen[p.flagName] {
				reqFlags = append(reqFlags, p.flagName)
				reqSeen[p.flagName] = true
			}
		}
		fmt.Fprintf(&buf, "func (o %s) RequiredFlags() []string {\n", typeName)
		if len(reqFlags) > 0 {
			sort.Strings(reqFlags)
			fmt.Fprintf(&buf, "\treturn []string{")
			for i, name := range reqFlags {
				if i > 0 {
					buf.WriteString(", ")
				}
				fmt.Fprintf(&buf, "%q", name)
			}
			fmt.Fprintf(&buf, "}\n")
		} else {
			fmt.Fprintf(&buf, "\treturn nil\n")
		}
		fmt.Fprintf(&buf, "}\n\n")

		hasFlags := false
		for _, p := range op.params {
			if !p.isPathParam {
				hasFlags = true
				break
			}
		}
		fmt.Fprintf(&buf, "func (o %s) Flags() []FlagDef {\n", typeName)
		if hasFlags {
			seen := map[string]bool{}
			fmt.Fprintf(&buf, "\treturn []FlagDef{\n")
			for _, p := range op.params {
				if p.isPathParam || seen[p.oasName] {
					continue
				}
				seen[p.oasName] = true
				fmt.Fprintf(&buf, "\t\t{Name: %q, OASName: %q, Type: %q", p.flagName, p.oasName, p.flagType)
				if p.defaultVal != "" {
					fmt.Fprintf(&buf, ", Default: %q", p.defaultVal)
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
				fmt.Fprintf(&buf, ", OpName: %q", op.goName)
				if p.required {
					fmt.Fprintf(&buf, ", Required: true")
				}
				fmt.Fprintf(&buf, "},\n")
			}
			fmt.Fprintf(&buf, "\t}\n")
		} else {
			fmt.Fprintf(&buf, "\treturn nil\n")
		}
		fmt.Fprintf(&buf, "}\n\n")
	}

	fmt.Fprintf(&buf, "// ResponseField describes a field in an API response.\n")
	fmt.Fprintf(&buf, "type ResponseField struct {\n")
	fmt.Fprintf(&buf, "\tName        string\n")
	fmt.Fprintf(&buf, "\tType        string\n")
	fmt.Fprintf(&buf, "\tDescription string\n")
	fmt.Fprintf(&buf, "\tEnumValues  []string\n")
	fmt.Fprintf(&buf, "}\n\n")

	fmt.Fprintf(&buf, "// ResponseSchemas maps operation names to their response fields.\n")
	fmt.Fprintf(&buf, "var ResponseSchemas = map[string][]ResponseField{\n")
	for _, op := range ops {
		if len(op.responseFields) == 0 {
			continue
		}
		fmt.Fprintf(&buf, "\t%q: {\n", op.goName)
		for _, f := range op.responseFields {
			fmt.Fprintf(&buf, "\t\t{Name: %q, Type: %q, Description: %q", f.name, f.jsonType, f.desc)
			if len(f.enumValues) > 0 {
				fmt.Fprintf(&buf, ", EnumValues: []string{")
				for i, v := range f.enumValues {
					if i > 0 {
						buf.WriteString(", ")
					}
					fmt.Fprintf(&buf, "%q", v)
				}
				buf.WriteString("}")
			}
			fmt.Fprintf(&buf, "},\n")
		}
		fmt.Fprintf(&buf, "\t},\n")
	}
	fmt.Fprintf(&buf, "}\n\n")

	fmt.Fprintf(&buf, "// OperationSummaries maps operation names to their summaries.\n")
	fmt.Fprintf(&buf, "var OperationSummaries = map[string]string{\n")
	for _, op := range ops {
		if len(op.responseFields) == 0 {
			continue
		}
		fmt.Fprintf(&buf, "\t%q: %q,\n", op.goName, op.summary)
	}
	fmt.Fprintf(&buf, "}\n\n")

	fmt.Fprintf(&buf, "// ArrayResponses tracks which operations return arrays vs single objects.\n")
	fmt.Fprintf(&buf, "var ArrayResponses = map[string]bool{\n")
	for _, op := range ops {
		if len(op.responseFields) == 0 {
			continue
		}
		if op.returnsArray {
			fmt.Fprintf(&buf, "\t%q: true,\n", op.goName)
		}
	}
	fmt.Fprintf(&buf, "}\n\n")

	fmt.Fprintf(&buf, "// AllOps lists every generated Op for iteration in tests and tooling.\n")
	fmt.Fprintf(&buf, "var AllOps = []Op{\n")
	for _, op := range ops {
		fmt.Fprintf(&buf, "\t%sOp,\n", op.goName)
	}
	fmt.Fprintf(&buf, "}\n")

	return buf.String()
}

// --- FromFlags generation ---

func genFromFlags(endpoints []*endpointInfo, schemas []*schemaInfo) string {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "// Code generated by cmd/generate; DO NOT EDIT.\n\n")
	fmt.Fprintf(&buf, "package cmd\n\n")
	fmt.Fprintf(&buf, "import (\n")
	fmt.Fprintf(&buf, "\t\"github.com/alpacahq/cli/internal/api\"\n")
	fmt.Fprintf(&buf, "\t\"github.com/alpacahq/cli/internal/cmdutil\"\n")
	fmt.Fprintf(&buf, "\t\"github.com/spf13/cobra\"\n")
	fmt.Fprintf(&buf, ")\n\n")

	sorted := make([]*endpointInfo, len(endpoints))
	copy(sorted, endpoints)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].goName < sorted[j].goName
	})

	for _, ep := range sorted {
		if len(ep.queryParams) == 0 {
			continue
		}

		funcName := lcFirst(ep.goName) + "ParamsFromFlags"
		paramsType := ep.goName + "Params"

		qps := make([]paramInfo, len(ep.queryParams))
		copy(qps, ep.queryParams)
		sort.Slice(qps, func(i, j int) bool {
			return qps[i].name < qps[j].name
		})

		fmt.Fprintf(&buf, "func %s(cmd *cobra.Command) *api.%s {\n", funcName, paramsType)
		fmt.Fprintf(&buf, "\tp := &api.%s{}\n", paramsType)
		fmt.Fprintf(&buf, "\tflags := cmd.Flags()\n")

		for _, qp := range qps {
			flagName := strings.ReplaceAll(qp.name, "_", "-")
			helper := "cmdutil.Str"
			switch qp.goType {
			case "int":
				helper = "cmdutil.Int"
			case "bool":
				helper = "cmdutil.Bool"
			}
			fmt.Fprintf(&buf, "\tif flags.Changed(%q) {\n", flagName)
			fmt.Fprintf(&buf, "\t\tp.%s = %s(cmd, %q)\n", qp.goName, helper, flagName)
			fmt.Fprintf(&buf, "\t}\n")
		}

		fmt.Fprintf(&buf, "\treturn p\n")
		fmt.Fprintf(&buf, "}\n\n")
	}

	// GEN2: generate BodyFromFlags for PATCH/PUT endpoints with body refs
	generated := map[string]bool{}
	for _, ep := range sorted {
		if (ep.method == "PATCH" || ep.method == "PUT") && ep.bodyRef != "" && !generated[ep.bodyRef] {
			genBodyFromFlags(&buf, ep.bodyRef, schemas)
			generated[ep.bodyRef] = true
		}
	}

	return buf.String()
}

func genBodyFromFlags(buf *bytes.Buffer, bodyRef string, schemas []*schemaInfo) {
	var bodySchema *schemaInfo
	for _, s := range schemas {
		if s.goName == bodyRef && s.kind == "struct" {
			bodySchema = s
			break
		}
	}
	if bodySchema == nil {
		return
	}

	type fieldInfo struct {
		flagName    string
		goFieldName string
		kind        string // "Str", "Bool", "Int", "enum"
		enumGoType  string
	}

	var propNames []string
	for name := range bodySchema.props {
		propNames = append(propNames, name)
	}
	sort.Strings(propNames)

	var simpleFields, enumFields []fieldInfo
	for _, propName := range propNames {
		propSchema := bodySchema.props[propName]
		flagName := strings.ReplaceAll(propName, "_", "-")
		goField := toGoFieldName(propName)

		if ref, ok := propSchema["$ref"].(string); ok {
			refName := refBaseName(ref)
			for _, s := range schemas {
				if s.name == refName && s.kind == "enum" {
					enumFields = append(enumFields, fieldInfo{
						flagName: flagName, goFieldName: goField,
						kind: "enum", enumGoType: s.goName,
					})
				}
			}
			continue
		}

		switch typ, _ := propSchema["type"].(string); typ {
		case "string":
			simpleFields = append(simpleFields, fieldInfo{flagName: flagName, goFieldName: goField, kind: "Str"})
		case "boolean":
			simpleFields = append(simpleFields, fieldInfo{flagName: flagName, goFieldName: goField, kind: "Bool"})
		case "integer":
			simpleFields = append(simpleFields, fieldInfo{flagName: flagName, goFieldName: goField, kind: "Int"})
		}
	}

	if len(simpleFields) == 0 && len(enumFields) == 0 {
		return
	}

	funcName := lcFirst(bodyRef) + "BodyFromFlags"
	fmt.Fprintf(buf, "func %s(cmd *cobra.Command) (*api.%s, bool) {\n", funcName, bodyRef)
	fmt.Fprintf(buf, "\tbody := &api.%s{}\n", bodyRef)

	if len(simpleFields) > 0 {
		fmt.Fprintf(buf, "\tp := cmdutil.NewPatchHelper(cmd)\n")
		for _, f := range simpleFields {
			fmt.Fprintf(buf, "\tp.%s(%q, &body.%s)\n", f.kind, f.flagName, f.goFieldName)
		}
	}

	for _, f := range enumFields {
		fmt.Fprintf(buf, "\tif cmdutil.Changed(cmd, %q) {\n", f.flagName)
		fmt.Fprintf(buf, "\t\tbody.%s = api.%s(cmdutil.Str(cmd, %q))\n", f.goFieldName, f.enumGoType, f.flagName)
		fmt.Fprintf(buf, "\t}\n")
	}

	fmt.Fprintf(buf, "\treturn body, ")
	if len(simpleFields) > 0 {
		fmt.Fprintf(buf, "p.AnyChanged()")
	}
	for i, f := range enumFields {
		if i > 0 || len(simpleFields) > 0 {
			fmt.Fprintf(buf, " || ")
		}
		fmt.Fprintf(buf, "cmdutil.Changed(cmd, %q)", f.flagName)
	}
	fmt.Fprintf(buf, "\n}\n\n")
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
