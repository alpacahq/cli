//go:build integration

package integration

import (
	"strings"
	"testing"
)

func TestActivity(t *testing.T) {
	out := alpaca(t, "account", "activity", "list", "--page-size", "5", "--json")
	_ = parseJSONArray(t, out)
}

func TestPortfolioHistory(t *testing.T) {
	out := alpaca(t, "account", "portfolio", "--period", "1W", "--timeframe", "1D", "--json")
	data := parseJSONMap(t, out)
	if data["equity"] == nil && data["timestamp"] == nil {
		t.Error("portfolio history missing expected fields")
	}
}

func TestAssetList(t *testing.T) {
	out := alpaca(t, "asset", "list", "--status", "active", "--asset-class", "us_equity", "--json")
	assets := parseJSONArray(t, out)
	if len(assets) == 0 {
		t.Fatal("expected at least one asset")
	}
}

func TestScreenerMostActives(t *testing.T) {
	out := alpaca(t, "data", "screener", "most-actives", "--top", "5", "--json")
	actives := parseJSONArray(t, out)
	if len(actives) == 0 {
		t.Error("screener response returned no results")
	}
}

func TestAPIPassthrough(t *testing.T) {
	out := alpaca(t, "api", "get", "/v2/clock", "--json")
	clock := parseJSONMap(t, out)
	if clock["is_open"] == nil {
		t.Error("api get /v2/clock missing 'is_open'")
	}
}

func TestDoctor(t *testing.T) {
	// Doctor may report "some checks failed" in isolated config dir,
	// but should still run and produce output.
	stdout, stderr, _ := alpacaWithStderr(t, "doctor")
	combined := string(stdout) + string(stderr)
	if !strings.Contains(combined, "trading") && !strings.Contains(combined, "Trading") &&
		!strings.Contains(combined, "API") && !strings.Contains(combined, "check") {
		t.Error("doctor output should mention API or trading checks")
	}
}

func TestProfileStatus(t *testing.T) {
	out := alpaca(t, "profile", "status")
	s := string(out)
	if s == "" {
		t.Error("profile status should produce output")
	}
}

func TestProfileList(t *testing.T) {
	out := alpaca(t, "profile", "list")
	_ = string(out) // just ensure exit 0
}

