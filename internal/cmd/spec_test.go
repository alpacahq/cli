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
		"asset-class": "CLI includes fixed_income which the spec omits",
		"exchange":    "CLI uses a broader set of exchanges",
		"feed":        "stock data feed differs from spec (otc, delayed_sip)",
		"timeframe":   "no spec enum; free-form in API",
		"period":      "no spec enum; free-form in API",
		"adjustment":  "no spec enum",
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

// TestNoEmptyFlagDescriptions scans the generated descriptions.go for all
// FlagDef variables, then verifies every flag used by command source has
// a non-empty Description and valid Name. No manual map needed — new flag
// sets are discovered automatically.
func TestNoEmptyFlagDescriptions(t *testing.T) {
	root := projectRoot()
	descData, err := os.ReadFile(filepath.Join(root, "internal", "api", "descriptions.go"))
	if err != nil {
		t.Fatal(err)
	}

	// Discover all generated flag set variables: var XxxFlags = []FlagDef{
	flagSetDecl := regexp.MustCompile(`var (\w+Flags) = \[\]FlagDef\{`)
	generatedSets := map[string]bool{}
	for _, m := range flagSetDecl.FindAllStringSubmatch(string(descData), -1) {
		generatedSets[m[1]] = true
	}

	// Find which flag sets are referenced in command source
	flagRef := regexp.MustCompile(`api\.(\w+Flags)`)
	dir := cmdDir()
	entries, err := os.ReadDir(dir)
	if err != nil {
		t.Fatal(err)
	}
	referenced := map[string]bool{}
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".go") || strings.HasSuffix(e.Name(), "_test.go") {
			continue
		}
		data, err := os.ReadFile(filepath.Join(dir, e.Name()))
		if err != nil {
			t.Fatal(err)
		}
		for _, m := range flagRef.FindAllStringSubmatch(string(data), -1) {
			referenced[m[1]] = true
		}
	}

	// Verify every referenced flag set exists in generated code
	for name := range referenced {
		if !generatedSets[name] {
			t.Errorf("command source references api.%s but it doesn't exist in descriptions.go", name)
		}
	}

	// Structurally scan each FlagDef in descriptions.go for empty descriptions/names
	flagEntry := regexp.MustCompile(`\{Name: "([^"]*)"[^}]*Description: "([^"]*)"`)
	src := string(descData)
	lines := strings.Split(src, "\n")
	var currentSet string
	setDecl := regexp.MustCompile(`^var (\w+Flags) = \[\]FlagDef\{`)
	for _, line := range lines {
		if m := setDecl.FindStringSubmatch(line); m != nil {
			currentSet = m[1]
			continue
		}
		if currentSet == "" || !referenced[currentSet] {
			continue
		}
		if m := flagEntry.FindStringSubmatch(line); m != nil {
			name, desc := m[1], m[2]
			if name == "" {
				t.Errorf("%s: FlagDef has empty Name", currentSet)
			}
			if desc == "" {
				t.Errorf("%s: flag %q has empty description", currentSet, name)
			}
			if strings.Contains(name, "_") {
				t.Errorf("%s: flag %q contains underscore (should be kebab-case)", currentSet, name)
			}
		}
	}
}

// TestRenderCallsUseColumnDefinitions scans all command source files and
// verifies that every output.Render and output.PrintSingle call includes
// column definitions (a *Columns() function or a `cols` parameter). This
// catches the antipattern of passing raw data through the table/CSV renderer
// without defining columns — which produces blank or garbled output.
func TestRenderCallsUseColumnDefinitions(t *testing.T) {
	renderCall := regexp.MustCompile(`output\.(Render|PrintSingle)\(`)
	columnsRef := regexp.MustCompile(`\w+Columns\(\)|\bcols\b`)

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

		lines := strings.Split(string(data), "\n")
		for i, line := range lines {
			if !renderCall.MatchString(line) {
				continue
			}
			if !columnsRef.MatchString(line) {
				t.Errorf("%s:%d: output.Render/PrintSingle call missing column definitions:\n  %s",
					e.Name(), i+1, strings.TrimSpace(line))
			}
		}
	}
}

// TestTypedDescriptionsCompile verifies that the typed description structs
// referenced in command source code have non-empty Summary fields.
// Field references are caught at compile time; this test catches empty values.
func TestTypedDescriptionsCompile(t *testing.T) {
	opRef := regexp.MustCompile(`api\.(\w+Op)\.Summary`)

	dir := cmdDir()
	entries, err := os.ReadDir(dir)
	if err != nil {
		t.Fatal(err)
	}

	summaries := map[string]string{
		"PostOrderOp":                            api.PostOrderOp.Summary,
		"GetAllOrdersOp":                         api.GetAllOrdersOp.Summary,
		"GetOrderByOrderIDOp":                    api.GetOrderByOrderIDOp.Summary,
		"DeleteOrderByOrderIDOp":                 api.DeleteOrderByOrderIDOp.Summary,
		"DeleteAllOrdersOp":                      api.DeleteAllOrdersOp.Summary,
		"PatchOrderByOrderIDOp":                  api.PatchOrderByOrderIDOp.Summary,
		"GetAllOpenPositionsOp":                  api.GetAllOpenPositionsOp.Summary,
		"GetOpenPositionOp":                      api.GetOpenPositionOp.Summary,
		"DeleteOpenPositionOp":                   api.DeleteOpenPositionOp.Summary,
		"DeleteAllOpenPositionsOp":               api.DeleteAllOpenPositionsOp.Summary,
		"GetAccountOp":                           api.GetAccountOp.Summary,
		"GetAccountConfigOp":                     api.GetAccountConfigOp.Summary,
		"PatchAccountConfigOp":                   api.PatchAccountConfigOp.Summary,
		"LegacyClockOp":                          api.LegacyClockOp.Summary,
		"LegacyCalendarOp":                       api.LegacyCalendarOp.Summary,
		"GetAccountPortfolioHistoryOp":           api.GetAccountPortfolioHistoryOp.Summary,
		"NewsOp":                                 api.NewsOp.Summary,
		"GetAccountActivitiesOp":                 api.GetAccountActivitiesOp.Summary,
		"MostActivesOp":                          api.MostActivesOp.Summary,
		"MoversOp":                               api.MoversOp.Summary,
		"GetV2AssetsOp":                          api.GetV2AssetsOp.Summary,
		"GetV2AssetsSymbolOrAssetIDOp":           api.GetV2AssetsSymbolOrAssetIDOp.Summary,
		"UsTreasuriesOp":                         api.UsTreasuriesOp.Summary,
		"UsCorporatesOp":                         api.UsCorporatesOp.Summary,
		"GetV2CorporateActionsAnnouncementsOp":   api.GetV2CorporateActionsAnnouncementsOp.Summary,
		"GetV2CorporateActionsAnnouncementsIDOp": api.GetV2CorporateActionsAnnouncementsIDOp.Summary,
		"GetOptionsContractsOp":                  api.GetOptionsContractsOp.Summary,
		"GetOptionContractSymbolOrIDOp":          api.GetOptionContractSymbolOrIDOp.Summary,
		"OptionExerciseOp":                       api.OptionExerciseOp.Summary,
		"OptionDoNotExerciseOp":                  api.OptionDoNotExerciseOp.Summary,
		"ListCryptoFundingWalletsOp":             api.ListCryptoFundingWalletsOp.Summary,
		"ListCryptoFundingTransfersOp":           api.ListCryptoFundingTransfersOp.Summary,
		"GetCryptoFundingTransferOp":             api.GetCryptoFundingTransferOp.Summary,
		"CreateCryptoTransferForAccountOp":       api.CreateCryptoTransferForAccountOp.Summary,
		"ListWhitelistedAddressOp":               api.ListWhitelistedAddressOp.Summary,
		"CreateWhitelistedAddressOp":             api.CreateWhitelistedAddressOp.Summary,
		"DeleteWhitelistedAddressOp":             api.DeleteWhitelistedAddressOp.Summary,
		"RatesOp":                                api.RatesOp.Summary,
		"LatestRatesOp":                          api.LatestRatesOp.Summary,
		"CryptoLatestOrderbooksOp":               api.CryptoLatestOrderbooksOp.Summary,
		"StockAuctionsOp":                        api.StockAuctionsOp.Summary,
		"CorporateActionsOp":                     api.CorporateActionsOp.Summary,
		"FixedIncomeLatestPricesOp":              api.FixedIncomeLatestPricesOp.Summary,
		"OptionBarsOp":                           api.OptionBarsOp.Summary,
		"OptionTradesOp":                         api.OptionTradesOp.Summary,
		"OptionSnapshotsOp":                      api.OptionSnapshotsOp.Summary,
		"OptionChainOp":                          api.OptionChainOp.Summary,
		"OptionLatestQuotesOp":                   api.OptionLatestQuotesOp.Summary,
		"OptionLatestTradesOp":                   api.OptionLatestTradesOp.Summary,
		"GetAccountActivitiesByActivityTypeOp":   api.GetAccountActivitiesByActivityTypeOp.Summary,
	}

	referenced := map[string]bool{}
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".go") || strings.HasSuffix(e.Name(), "_test.go") {
			continue
		}
		data, err := os.ReadFile(filepath.Join(dir, e.Name()))
		if err != nil {
			t.Fatal(err)
		}
		for _, m := range opRef.FindAllStringSubmatch(string(data), -1) {
			referenced[m[1]] = true
		}
	}

	for name, summary := range summaries {
		if !referenced[name] {
			continue
		}
		if summary == "" {
			t.Errorf("%s.Summary is empty", name)
		}
	}
}

// TestGeneratedPostPutPatchPassQueryParams scans the generated trading and market
// data client files and verifies that every method accepting a *Params argument
// for POST/PUT/PATCH also passes params.Values() in its body.
func TestGeneratedPostPutPatchPassQueryParams(t *testing.T) {
	root := projectRoot()
	clientFiles := []string{
		filepath.Join(root, "internal", "api", "trading_client.go"),
		filepath.Join(root, "internal", "api", "marketdata_client.go"),
	}

	funcSig := regexp.MustCompile(`^func \(c \*\w+Client\) (\w+)\(params \*\w+Params`)
	callLine := regexp.MustCompile(`c\.Raw\.(Post|Put|Patch)\(`)
	paramsUsed := regexp.MustCompile(`params\.Values\(\)`)

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
							"%s:%d: %s calls Post/Put/Patch without params.Values() — query params would be dropped",
							filepath.Base(path), j+1, funcName,
						)
					}
				}
			}
		}
	}
}
