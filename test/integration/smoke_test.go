//go:build integration

package integration

import (
	"strings"
	"testing"
)

func TestSmoke_Version(t *testing.T) {
	out := alpaca(t, "version")
	if !strings.Contains(string(out), "alpaca version") {
		t.Fatalf("unexpected version output: %s", out)
	}
}

func TestSmoke_Account(t *testing.T) {
	out := alpaca(t, "account", "get", "-o", "json")
	acct := parseJSONMap(t, out)

	for _, field := range []string{"id", "status", "equity", "buying_power", "cash"} {
		if _, ok := acct[field]; !ok {
			t.Errorf("account missing field: %s", field)
		}
	}
}

func TestSmoke_Clock(t *testing.T) {
	out := alpaca(t, "clock", "-o", "json")
	clock := parseJSONMap(t, out)

	for _, field := range []string{"is_open", "next_open", "next_close"} {
		if _, ok := clock[field]; !ok {
			t.Errorf("clock missing field: %s", field)
		}
	}
}

func TestSmoke_Calendar(t *testing.T) {
	out := alpaca(t, "calendar", "--start", "2025-01-01", "--end", "2025-01-31", "-o", "json")
	days := parseJSONArray(t, out)

	if len(days) == 0 {
		t.Fatal("expected at least one calendar day")
	}
	if _, ok := days[0]["date"]; !ok {
		t.Error("calendar entry missing 'date' field")
	}
}

func TestSmoke_AccountConfig(t *testing.T) {
	out := alpaca(t, "account", "config", "get", "-o", "json")
	cfg := parseJSONMap(t, out)

	if _, ok := cfg["dtbp_check"]; !ok {
		t.Error("account config missing 'dtbp_check'")
	}
}

func TestSmoke_Assets(t *testing.T) {
	out := alpaca(t, "asset", "get", "AAPL", "-o", "json")
	asset := parseJSONMap(t, out)

	if asset["symbol"] != "AAPL" {
		t.Errorf("expected symbol AAPL, got %v", asset["symbol"])
	}
	if asset["tradable"] != true {
		t.Errorf("expected AAPL to be tradable")
	}
}
