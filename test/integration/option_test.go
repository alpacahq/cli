//go:build integration

package integration

import (
	"testing"
)

func TestOptionContracts(t *testing.T) {
	t.Parallel()
	out := alpaca(t, "option", "contracts", "--underlying-symbols", "AAPL")
	wrapper := parseJSONMap(t, out)
	contracts, ok := wrapper["option_contracts"].([]any)
	if !ok || len(contracts) == 0 {
		t.Fatal("expected non-empty option_contracts array")
	}
	first, _ := contracts[0].(map[string]any)
	requireFields(t, first, "symbol", "underlying_symbol", "expiration_date")

	t.Run("filter", func(t *testing.T) {
		out := alpaca(t, "option", "contracts", "--underlying-symbols", "AAPL", "--type", "call")
		wrapper := parseJSONMap(t, out)
		filtered, ok := wrapper["option_contracts"].([]any)
		if !ok || len(filtered) == 0 {
			t.Fatal("expected non-empty option_contracts array")
		}
		for _, c := range filtered[:min(len(filtered), 3)] {
			m, _ := c.(map[string]any)
			if m["type"] != "call" {
				t.Errorf("expected type call, got %v", m["type"])
			}
		}
	})

	t.Run("expiration_filter", func(t *testing.T) {
		expDate, _ := first["expiration_date"].(string)
		if expDate == "" {
			t.Skip("no expiration_date on first contract")
		}
		out := alpaca(t, "option", "contracts",
			"--underlying-symbols", "AAPL",
			"--expiration-date-gte", expDate,
			"--expiration-date-lte", expDate,
		)
		wrapper := parseJSONMap(t, out)
		filtered, ok := wrapper["option_contracts"].([]any)
		if !ok || len(filtered) == 0 {
			t.Fatal("expected non-empty option_contracts for expiration filter")
		}
		for _, c := range filtered[:min(len(filtered), 3)] {
			m, _ := c.(map[string]any)
			if m["expiration_date"] != expDate {
				t.Errorf("expected expiration_date %q, got %v", expDate, m["expiration_date"])
			}
		}
	})

	t.Run("strike_filter", func(t *testing.T) {
		strike, _ := first["strike_price"].(string)
		if strike == "" {
			t.Skip("no strike_price on first contract")
		}
		out := alpaca(t, "option", "contracts",
			"--underlying-symbols", "AAPL",
			"--strike-price-gte", strike,
			"--strike-price-lte", strike,
		)
		wrapper := parseJSONMap(t, out)
		filtered, ok := wrapper["option_contracts"].([]any)
		if !ok || len(filtered) == 0 {
			t.Fatal("expected non-empty option_contracts for strike filter")
		}
	})

	t.Run("status_filter", func(t *testing.T) {
		out := alpaca(t, "option", "contracts",
			"--underlying-symbols", "AAPL",
			"--status", "active",
		)
		wrapper := parseJSONMap(t, out)
		filtered, ok := wrapper["option_contracts"].([]any)
		if !ok || len(filtered) == 0 {
			t.Fatal("expected non-empty option_contracts for status=active")
		}
		for _, c := range filtered[:min(len(filtered), 3)] {
			m, _ := c.(map[string]any)
			if m["status"] != "active" {
				t.Errorf("expected status active, got %v", m["status"])
			}
		}
	})

	t.Run("get", func(t *testing.T) {
		sym := first["symbol"].(string)
		out := alpaca(t, "option", "get", "--symbol-or-id", sym)
		contract := parseJSONMap(t, out)
		requireFields(t, contract, "symbol", "underlying_symbol", "expiration_date", "strike_price")
		if contract["symbol"] != sym {
			t.Errorf("expected symbol %q, got %v", sym, contract["symbol"])
		}
	})
}
