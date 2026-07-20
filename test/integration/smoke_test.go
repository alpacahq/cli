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
	requireFields(t, acct, "id", "status", "equity", "buying_power", "cash", "currency", "account_number")

	status, _ := acct["status"].(string)
	if status != "ACTIVE" {
		t.Errorf("expected account status ACTIVE, got %q", status)
	}
	for _, field := range []string{"equity", "buying_power", "cash"} {
		val, ok := acct[field].(string)
		if !ok {
			t.Errorf("expected %s to be a string, got %T", field, acct[field])
			continue
		}
		if val == "" {
			t.Errorf("expected non-empty %s value", field)
		}
	}
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
	requireFields(t, cfg, "max_margin_multiplier")
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

func TestSmoke_ErrorContract(t *testing.T) {
	t.Parallel()
	_, stderr, code := alpacaFail(t,
		"order", "get", "--order-id", "00000000-0000-0000-0000-000000000000",
	)
	if code != 1 {
		t.Errorf("expected exit code 1 for API error, got %d", code)
	}
	errMap := parseJSONMap(t, stderr)
	requireFields(t, errMap, "error", "status")
	if errMsg, ok := errMap["error"].(string); !ok || errMsg == "" {
		t.Errorf("expected non-empty error message, got %v", errMap["error"])
	}
	if status, ok := errMap["status"].(float64); !ok || status < 400 {
		t.Errorf("expected HTTP status >= 400, got %v", errMap["status"])
	}
}
