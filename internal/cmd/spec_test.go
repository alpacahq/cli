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

	// Known divergences: CLI intentionally uses a different set than the spec.
	knownDivergences := map[string]string{
		"class":      "CLI includes fixed_income which the spec omits",
		"exchange":   "CLI uses a broader set of exchanges",
		"feed":       "stock data feed differs from spec (otc, delayed_sip)",
		"timeframe":  "no spec enum; free-form in API",
		"period":     "no spec enum; free-form in API",
		"adjustment": "no spec enum",
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

			if _, ok := knownDivergences[flagName]; ok {
				continue
			}

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

// TestMutatingCommandsHaveWarnLive verifies that every command calling a
// mutating API method also calls warnLive() or requireConfirmation().
// Exemptions must be listed explicitly with a reason.
func TestMutatingCommandsHaveWarnLive(t *testing.T) {
	exempted := map[string]string{
		"PatchAccountConfig":                "config update, low-risk",
		"PostWatchlist":                     "watchlist ops are non-financial",
		"UpdateWatchlistByID":              "watchlist ops are non-financial",
		"AddAssetToWatchlist":               "watchlist ops are non-financial",
		"DeleteWatchlistByID":              "watchlist ops are non-financial",
		"RemoveAssetFromWatchlist":          "watchlist ops are non-financial",
		"UpdateWatchlistByName":             "watchlist ops are non-financial",
		"AddAssetToWatchlistByName":         "watchlist ops are non-financial",
		"DeleteWatchlistByName":             "watchlist ops are non-financial",
		"CreateWhitelistedAddress":          "address management, not a trade",
		"DeleteWhitelistedAddress":          "address management, not a trade",
		"CreateWhitelistedPerpAddress":      "address management, not a trade",
		"DeleteWhitelistedPerpAddress":      "address management, not a trade",
		"SetCryptoPerpAccountLeverage":      "config setting, not a trade",
		"CreateCryptoPerpTransferForAccount": "internal wallet transfer",
	}

	dir := cmdDir()
	entries, err := os.ReadDir(dir)
	if err != nil {
		t.Fatal(err)
	}

	methodCall := regexp.MustCompile(`(?:tradingClient|dataClient)\.\s*(\w+)\s*\(`)

	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".go") || strings.HasSuffix(e.Name(), "_test.go") {
			continue
		}
		data, err := os.ReadFile(filepath.Join(dir, e.Name()))
		if err != nil {
			t.Fatal(err)
		}
		src := string(data)

		matches := methodCall.FindAllStringSubmatch(src, -1)
		for _, m := range matches {
			method := m[1]
			if !api.TradingMutatingMethods[method] {
				continue
			}
			if _, ok := exempted[method]; ok {
				continue
			}
			if !strings.Contains(src, "warnLive()") && !strings.Contains(src, "requireConfirmation(") {
				t.Errorf("%s calls mutating method %s but has no warnLive() or requireConfirmation()", e.Name(), method)
			}
		}
	}
}

// TestValidateMethods verifies generated Validate() methods work correctly.
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

	t.Run("UpdateWatchlistRequest/empty", func(t *testing.T) {
		r := &api.UpdateWatchlistRequest{}
		if err := r.Validate(); err == nil {
			t.Error("expected validation error for empty name")
		}
	})

	t.Run("UpdateWatchlistRequest/valid", func(t *testing.T) {
		r := &api.UpdateWatchlistRequest{Name: "my-list"}
		if err := r.Validate(); err != nil {
			t.Errorf("unexpected error: %s", err)
		}
	})
}
