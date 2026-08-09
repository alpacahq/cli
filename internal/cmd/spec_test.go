package cmd

import (
	"encoding/json"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"sort"
	"strings"
	"testing"

	"github.com/alpacahq/cli/internal/api"
)

func projectRoot() string {
	_, file, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(file), "..", "..")
}

func cmdDir() string {
	_, file, _, _ := runtime.Caller(0)
	return filepath.Dir(file)
}

func loadTestSpec(t *testing.T, name string) map[string]any {
	t.Helper()
	data, err := os.ReadFile(filepath.Join(projectRoot(), "api", "specs", name))
	if err != nil {
		t.Fatal(err)
	}
	var spec map[string]any
	if err := json.Unmarshal(data, &spec); err != nil {
		t.Fatal(err)
	}
	return spec
}

// allSpecEnums extracts every enum from a spec: schema-level, inline property,
// and query parameter enums. Returns map[enumKey][]string where enumKey is a
// descriptive identifier like "schemas.OrderSide" or "params.status.GET./v2/orders".
func allSpecEnums(spec map[string]any) map[string][]string {
	result := map[string][]string{}

	schemas, _ := spec["components"].(map[string]any)
	if schemas != nil {
		schemas, _ = schemas["schemas"].(map[string]any)
	}
	for name, raw := range schemas {
		schema, ok := raw.(map[string]any)
		if !ok {
			continue
		}
		if enumVals := extractEnum(schema); len(enumVals) > 0 {
			result["schema:"+name] = enumVals
		}
		props, _ := schema["properties"].(map[string]any)
		for propName, propRaw := range props {
			propSchema, ok := propRaw.(map[string]any)
			if !ok {
				continue
			}
			if enumVals := extractEnum(propSchema); len(enumVals) > 0 {
				result["prop:"+name+"."+propName] = enumVals
			}
		}
	}

	paths, _ := spec["paths"].(map[string]any)
	for path, pathRaw := range paths {
		pathObj, ok := pathRaw.(map[string]any)
		if !ok {
			continue
		}
		for method, opRaw := range pathObj {
			op, ok := opRaw.(map[string]any)
			if !ok {
				continue
			}
			params, _ := op["parameters"].([]any)
			for _, pRaw := range params {
				p, ok := pRaw.(map[string]any)
				if !ok {
					continue
				}
				pName, _ := p["name"].(string)
				pSchema, _ := p["schema"].(map[string]any)
				if pSchema == nil {
					continue
				}
				if ref, ok := pSchema["$ref"].(string); ok {
					refName := ref[strings.LastIndex(ref, "/")+1:]
					if s, ok := schemas[refName].(map[string]any); ok {
						pSchema = s
					}
				}
				if enumVals := extractEnum(pSchema); len(enumVals) > 0 {
					result["param:"+pName+"."+method+"."+path] = enumVals
				}
			}
		}
	}

	return result
}

func extractEnum(schema map[string]any) []string {
	raw, ok := schema["enum"].([]any)
	if !ok || len(raw) == 0 {
		return nil
	}
	var vals []string
	for _, v := range raw {
		if s := strings.TrimSpace(stringOf(v)); s != "" {
			vals = append(vals, s)
		}
	}
	sort.Strings(vals)
	return vals
}

func stringOf(v any) string {
	if v == nil {
		return ""
	}
	if s, ok := v.(string); ok {
		return s
	}
	return ""
}

// TestHardcodedCompletionsMatchSpec parses the spec and every
// cobra.FixedCompletions([]string{...}, ...) call in cmd source. For each
// hardcoded set, it checks whether a matching spec enum exists. If one does
// and the values don't match, the test fails — that's real drift.
func TestHardcodedCompletionsMatchSpec(t *testing.T) {
	tradingSpec := loadTestSpec(t, "trading-api.json")
	marketSpec := loadTestSpec(t, "market-data-api.json")

	tradingEnums := allSpecEnums(tradingSpec)
	marketEnums := allSpecEnums(marketSpec)

	allEnums := map[string][]string{}
	for k, v := range tradingEnums {
		allEnums[k] = v
	}
	for k, v := range marketEnums {
		allEnums[k] = v
	}

	// Build a lookup: normalized enum values → spec key, for matching.
	type specEnum struct {
		key    string
		values []string
	}
	byValueSet := map[string]specEnum{}
	for key, vals := range allEnums {
		norm := strings.Join(vals, ",")
		byValueSet[norm] = specEnum{key: key, values: vals}
	}

	hardcoded := regexp.MustCompile(
		`RegisterFlagCompletionFunc\("([^"]+)",\s*cobra\.FixedCompletions\(\[\]string\{([^}]+)\}`,
	)

	dir := cmdDir()
	entries, err := os.ReadDir(dir)
	if err != nil {
		t.Fatal(err)
	}

	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".go") || strings.HasSuffix(e.Name(), "_test.go") {
			continue
		}
		data, err := os.ReadFile(filepath.Join(dir, e.Name()))
		if err != nil {
			t.Fatal(err)
		}
		src := string(data)

		matches := hardcoded.FindAllStringSubmatch(src, -1)
		for _, m := range matches {
			flagName := m[1]
			rawValues := m[2]

			var cliValues []string
			for _, v := range strings.Split(rawValues, ",") {
				v = strings.TrimSpace(v)
				v = strings.Trim(v, `"`)
				if v != "" {
					cliValues = append(cliValues, v)
				}
			}
			sort.Strings(cliValues)

			for enumKey, specValues := range allEnums {
				if !isParamOrPropMatch(enumKey, flagName) {
					continue
				}

				cliStr := strings.Join(cliValues, ",")
				specStr := strings.Join(specValues, ",")
				if cliStr != specStr {
					t.Errorf(
						"%s: hardcoded completions for --%s drifted from spec enum %s\n  cli:  %v\n  spec: %v",
						e.Name(), flagName, enumKey, cliValues, specValues,
					)
				}
			}
		}
	}
}

// isParamOrPropMatch checks if a spec enum key plausibly corresponds to a CLI flag name.
func isParamOrPropMatch(enumKey, flagName string) bool {
	specName := enumKey
	if i := strings.LastIndex(specName, "."); i >= 0 {
		specName = specName[strings.Index(specName, ":")+1 : i]
		if j := strings.LastIndex(specName, "."); j >= 0 {
			specName = specName[j+1:]
		}
	}
	flagNorm := strings.ReplaceAll(flagName, "-", "_")
	return specName == flagNorm
}

// TestValidateMethods verifies generated required-field validation works correctly.
func TestValidateMethods(t *testing.T) {
	t.Run("CreateCryptoTransferRequest/empty", func(t *testing.T) {
		r := &api.CreateCryptoTransferRequest{}
		err := r.Validate()
		if err == nil {
			t.Fatal("expected validation error for empty request")
		}
		for _, field := range []string{"address", "amount", "asset"} {
			if !strings.Contains(err.Error(), field) {
				t.Errorf("expected %q in error, got: %s", field, err)
			}
		}
	})

	t.Run("CreateCryptoTransferRequest/valid", func(t *testing.T) {
		r := &api.CreateCryptoTransferRequest{Address: "0x1", Amount: "100", Asset: "BTC"}
		if err := r.Validate(); err != nil {
			t.Errorf("unexpected error: %s", err)
		}
	})
}

// TestAllOpsValid iterates every generated Op via api.AllOps and validates
// summaries, flag descriptions, and kebab-case names. Replaces the old
// TestNoEmptyFlagDescriptions (which matched a removed pattern) and
// TestTypedDescriptionsCompile (which required a hand-curated map).
func TestAllOpsValid(t *testing.T) {
	if len(api.AllOps) == 0 {
		t.Fatal("api.AllOps is empty — generator may not have run")
	}
	for _, op := range api.AllOps {
		if op.Summary == "" {
			t.Errorf("op %q has empty Summary", op.Name)
			continue
		}
		t.Run(op.Summary, func(t *testing.T) {
			for _, f := range op.Flags {
				if f.Name == "" {
					t.Error("FlagDef has empty Name")
				}
				if f.Description == "" {
					t.Errorf("flag %q has empty Description", f.Name)
				}
				if strings.Contains(f.Name, "_") {
					t.Errorf("flag %q contains underscore (should be kebab-case)", f.Name)
				}
			}
		})
	}
}

// TestGeneratedMutatingMethodsPassParams scans the generated client files and
// verifies that every POST/PUT/PATCH method accepting a params url.Values argument
// actually passes it in its HTTP call.
func TestGeneratedMutatingMethodsPassParams(t *testing.T) {
	root := projectRoot()
	clientFiles := []string{
		filepath.Join(root, "internal", "api", "trading_client.gen.go"),
		filepath.Join(root, "internal", "api", "marketdata_client.gen.go"),
	}

	funcSig := regexp.MustCompile(`^func \(c \*\w+Client\) (\w+)\(.* params url\.Values`)
	callLine := regexp.MustCompile(`c\.Raw\.(Post|Put|Patch)\(`)
	paramsUsed := regexp.MustCompile(`\bparams\b`)

	for _, path := range clientFiles {
		data, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("reading %s: %v", path, err)
		}
		lines := strings.Split(string(data), "\n")

		for i, line := range lines {
			if m := funcSig.FindStringSubmatch(line); m != nil {
				funcName := m[1]
				for j := i + 1; j < len(lines); j++ {
					body := lines[j]
					if strings.TrimSpace(body) == "}" && !strings.HasPrefix(body, "\t\t") {
						break
					}
					if callLine.MatchString(body) && !paramsUsed.MatchString(body) {
						t.Errorf(
							"%s:%d: %s calls Post/Put/Patch without passing params — query params would be dropped",
							filepath.Base(path), j+1, funcName,
						)
					}
				}
			}
		}
	}
}
