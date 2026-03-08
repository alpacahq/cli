//go:build integration

package integration

import (
	"strings"
	"testing"
)

func TestSmoke_Version(t *testing.T) {
	t.Parallel()
	out := alpaca(t, "version")
	if !strings.Contains(string(out), "alpaca version") {
		t.Fatalf("unexpected version output: %s", out)
	}
}

func TestSmoke_Account(t *testing.T) {
	t.Parallel()
	out := alpaca(t, "account", "get", "--json")
	acct := parseJSONMap(t, out)
	requireFields(t, acct, "id", "status", "equity", "buying_power", "cash")
}

func TestSmoke_Clock(t *testing.T) {
	t.Parallel()
	out := alpaca(t, "clock", "--json")
	clock := parseJSONMap(t, out)
	requireFields(t, clock, "is_open", "next_open", "next_close")
}

func TestSmoke_Calendar(t *testing.T) {
	t.Parallel()
	calStart, calEnd := monthRange(3)
	out := alpaca(t, "calendar", "--start", calStart, "--end", calEnd, "--json")
	days := requireArrayNonEmpty(t, out)
	requireFields(t, days[0], "date")
}

func TestSmoke_AccountConfig(t *testing.T) {
	t.Parallel()
	out := alpaca(t, "account", "config", "get", "--json")
	cfg := parseJSONMap(t, out)
	requireFields(t, cfg, "dtbp_check")
}

func TestSmoke_Assets(t *testing.T) {
	t.Parallel()
	out := alpaca(t, "asset", "get", "AAPL", "--json")
	asset := parseJSONMap(t, out)

	if asset["symbol"] != "AAPL" {
		t.Errorf("expected symbol AAPL, got %v", asset["symbol"])
	}
	if asset["tradable"] != true {
		t.Errorf("expected AAPL to be tradable")
	}
}
