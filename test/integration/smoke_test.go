//go:build integration

package integration

import (
	"strings"
	"testing"
)

func TestSmoke_Version(t *testing.T) {
	t.Parallel()
	out := alpaca(t, "version")
	s := strings.TrimSpace(string(out))
	if s == "" {
		t.Fatal("version output should not be empty")
	}
	if strings.HasPrefix(s, "{") {
		t.Fatalf("version output should be plain text, got JSON: %s", s)
	}
}

func TestSmoke_RootVersionFlag(t *testing.T) {
	t.Parallel()
	out := alpaca(t, "--version")
	s := strings.TrimSpace(string(out))
	if s == "" {
		t.Fatal("--version output should not be empty")
	}
	if strings.HasPrefix(s, "{") {
		t.Fatalf("--version output should be plain text, got JSON: %s", s)
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
	out := alpaca(t, "asset", "get", "--symbol-or-asset-id", "AAPL")
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
	stdout, stderr, code := alpacaWithStderr(t, "doctor")
	if code != 0 {
		t.Fatalf("doctor should succeed for a healthy test setup, got exit %d\nstdout: %s\nstderr: %s", code, stdout, stderr)
	}
	combined := string(stdout) + string(stderr)
	if !strings.Contains(combined, "trading") && !strings.Contains(combined, "Trading") &&
		!strings.Contains(combined, "API") && !strings.Contains(combined, "check") {
		t.Error("doctor output should mention API or trading checks")
	}
}

func TestSmoke_ProfileList(t *testing.T) {
	t.Parallel()
	_ = string(alpaca(t, "profile", "list"))
}

func TestSmoke_ScreenerMostActives(t *testing.T) {
	t.Parallel()
	out := alpaca(t, "data", "screener", "most-actives", "--top", "5")
	data := parseJSONMap(t, out)
	requireFields(t, data, "most_actives", "last_updated")
}
