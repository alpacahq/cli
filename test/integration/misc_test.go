//go:build integration

package integration

import (
	"testing"
)

func TestActivity(t *testing.T) {
	out := alpaca(t, "activity", "list", "--limit", "5", "--json")
	_ = parseJSONArray(t, out)
}

func TestPortfolioHistory(t *testing.T) {
	out := alpaca(t, "portfolio", "history", "--period", "1W", "--timeframe", "1D", "--json")
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
	out := alpaca(t, "screener", "most-actives", "--top", "5", "--json")
	data := parseJSONMap(t, out)
	if data["most_actives"] == nil {
		t.Error("screener response missing 'most_actives'")
	}
}

func TestAPIPassthrough(t *testing.T) {
	out := alpaca(t, "api", "get", "/v2/clock", "--json")
	clock := parseJSONMap(t, out)
	if clock["is_open"] == nil {
		t.Error("api get /v2/clock missing 'is_open'")
	}
}

func TestPriceShortcut(t *testing.T) {
	out := alpaca(t, "price", "AAPL", "--json")
	data := parseJSONMap(t, out)
	if data["latestTrade"] == nil {
		t.Error("price snapshot missing 'latestTrade'")
	}
}

func TestOrdersShortcut(t *testing.T) {
	out := alpaca(t, "orders", "--json")
	_ = parseJSONArray(t, out)
}
