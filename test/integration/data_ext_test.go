//go:build integration

package integration

import (
	"testing"
)

func TestDataCryptoOrderbook(t *testing.T) {
	t.Parallel()
	out := alpaca(t, "data", "crypto-orderbook", "--symbols", "BTC/USD")
	data := parseJSONMap(t, out)
	if len(data) == 0 {
		t.Error("expected non-empty crypto orderbook response")
	}
}

func TestDataAuctions(t *testing.T) {
	t.Parallel()
	out := alpaca(t, "data", "auctions", "--symbols", "AAPL", "--start", daysAgo(100), "--end", daysAgo(93))
	data := parseJSONMap(t, out)
	if len(data) == 0 {
		t.Error("expected non-empty auctions response")
	}
}

func TestDataCorporateActions(t *testing.T) {
	t.Parallel()
	out := alpaca(t, "data", "corporate-actions", "--symbols", "AAPL")
	data := parseJSONMap(t, out)
	if len(data) == 0 {
		t.Error("expected non-empty corporate actions response")
	}
}

func TestDataMetaExchanges(t *testing.T) {
	t.Parallel()
	out := alpaca(t, "data", "meta", "exchanges")
	data := parseJSONMap(t, out)
	if len(data) == 0 {
		t.Error("expected non-empty exchanges map")
	}
}

func TestDataMetaConditions(t *testing.T) {
	t.Parallel()
	out := alpaca(t, "data", "meta", "conditions", "--ticktype", "trade", "--tape", "A")
	data := parseJSONMap(t, out)
	if len(data) == 0 {
		t.Error("expected non-empty conditions map")
	}
}

func TestDataScreenerMovers(t *testing.T) {
	t.Parallel()
	out := alpaca(t, "data", "screener", "movers")
	data := parseJSONMap(t, out)
	if len(data) == 0 {
		t.Error("expected non-empty movers response")
	}
}
