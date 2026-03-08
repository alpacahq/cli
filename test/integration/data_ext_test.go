//go:build integration

package integration

import (
	"testing"
)

func TestDataCryptoOrderbook(t *testing.T) {
	out := alpaca(t, "data", "crypto-orderbook", "--symbols", "BTC/USD", "--json")
	data := parseJSONMap(t, out)
	if len(data) == 0 {
		t.Error("expected non-empty crypto orderbook response")
	}
}

func TestDataAuctions(t *testing.T) {
	out := alpaca(t, "data", "auctions", "--symbols", "AAPL", "--start", daysAgo(100), "--end", daysAgo(93), "--json")
	data := parseJSONMap(t, out)
	if len(data) == 0 {
		t.Error("expected non-empty auctions response")
	}
}

func TestDataCorporateActions(t *testing.T) {
	out := alpaca(t, "data", "corporate-actions", "--symbols", "AAPL", "--json")
	data := parseJSONMap(t, out)
	if len(data) == 0 {
		t.Error("expected non-empty corporate actions response")
	}
}

func TestDataMetaExchanges(t *testing.T) {
	out := alpaca(t, "data", "meta", "exchanges", "--json")
	data := parseJSONMap(t, out)
	if len(data) == 0 {
		t.Error("expected non-empty exchanges map")
	}
}

func TestDataMetaConditions(t *testing.T) {
	out := alpaca(t, "data", "meta", "conditions", "trade", "--tape", "A", "--json")
	data := parseJSONMap(t, out)
	if len(data) == 0 {
		t.Error("expected non-empty conditions map")
	}
}

func TestDataScreenerMovers(t *testing.T) {
	out := alpaca(t, "data", "screener", "movers", "--json")
	data := parseJSONMap(t, out)
	if len(data) == 0 {
		t.Error("expected non-empty movers response")
	}
}
