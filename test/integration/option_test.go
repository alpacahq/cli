//go:build integration

package integration

import (
	"testing"
)

func TestOptionContracts(t *testing.T) {
	out := alpaca(t, "option", "contracts", "AAPL", "--json")
	wrapper := parseJSONMap(t, out)
	contracts, ok := wrapper["option_contracts"].([]any)
	if !ok || len(contracts) == 0 {
		t.Fatal("expected non-empty option_contracts array")
	}
	first, _ := contracts[0].(map[string]any)
	requireFields(t, first, "symbol", "underlying_symbol", "expiration_date")
}

func TestOptionContracts_Filter(t *testing.T) {
	out := alpaca(t, "option", "contracts", "AAPL", "--type", "call", "--json")
	wrapper := parseJSONMap(t, out)
	contracts, ok := wrapper["option_contracts"].([]any)
	if !ok || len(contracts) == 0 {
		t.Fatal("expected non-empty option_contracts array")
	}
	for _, c := range contracts[:min(len(contracts), 3)] {
		m, _ := c.(map[string]any)
		if m["type"] != "call" {
			t.Errorf("expected type call, got %v", m["type"])
		}
	}
}

func TestOptionGet(t *testing.T) {
	out := alpaca(t, "option", "contracts", "AAPL", "--json")
	wrapper := parseJSONMap(t, out)
	contracts, ok := wrapper["option_contracts"].([]any)
	if !ok || len(contracts) == 0 {
		t.Fatal("no contracts to test get")
	}
	first, _ := contracts[0].(map[string]any)
	sym := first["symbol"].(string)

	out = alpaca(t, "option", "get", sym, "--json")
	contract := parseJSONMap(t, out)
	requireFields(t, contract, "symbol", "underlying_symbol", "expiration_date", "strike_price")
	if contract["symbol"] != sym {
		t.Errorf("expected symbol %q, got %v", sym, contract["symbol"])
	}
}
