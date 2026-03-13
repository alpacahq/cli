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
	out := alpaca(t, "account", "get")
	acct := parseJSONMap(t, out)
	requireFields(t, acct, "id", "status", "equity", "buying_power", "cash")
}

func TestSmoke_Clock(t *testing.T) {
	t.Parallel()
	out := alpaca(t, "clock")
	clock := parseJSONMap(t, out)
	requireFields(t, clock, "is_open", "next_open", "next_close")
}

func TestSmoke_Calendar(t *testing.T) {
	t.Parallel()
	calStart, calEnd := monthRange(3)
	out := alpaca(t, "calendar", "--start", calStart, "--end", calEnd)
	days := requireArrayNonEmpty(t, out)
	requireFields(t, days[0], "date")
}

func TestSmoke_AccountConfig(t *testing.T) {
	t.Parallel()
	out := alpaca(t, "account", "config", "get")
	cfg := parseJSONMap(t, out)
	requireFields(t, cfg, "dtbp_check")
}

func TestSmoke_Assets(t *testing.T) {
	t.Parallel()
	out := alpaca(t, "asset", "get", "AAPL")
	asset := parseJSONMap(t, out)

	if asset["symbol"] != "AAPL" {
		t.Errorf("expected symbol AAPL, got %v", asset["symbol"])
	}
	if asset["tradable"] != true {
		t.Errorf("expected AAPL to be tradable")
	}
}

func TestSmoke_Doctor(t *testing.T) {
	t.Parallel()
	stdout, stderr, _ := alpacaWithStderr(t, "doctor")
	combined := string(stdout) + string(stderr)
	if !strings.Contains(combined, "trading") && !strings.Contains(combined, "Trading") &&
		!strings.Contains(combined, "API") && !strings.Contains(combined, "check") {
		t.Error("doctor output should mention API or trading checks")
	}
}

func TestSmoke_ProfileStatus(t *testing.T) {
	t.Parallel()
	out := alpaca(t, "profile", "status")
	if string(out) == "" {
		t.Error("profile status should produce output")
	}
}

func TestSmoke_ProfileList(t *testing.T) {
	t.Parallel()
	_ = string(alpaca(t, "profile", "list"))
}

func TestSmoke_ScreenerMostActives(t *testing.T) {
	t.Parallel()
	out := alpaca(t, "data", "screener", "most-actives", "--top", "5")
	actives := parseJSONArray(t, out)
	if len(actives) == 0 {
		t.Error("screener response returned no results")
	}
}
