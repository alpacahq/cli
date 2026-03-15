//go:build integration

package integration

import (
	"testing"
)

func TestDataNews(t *testing.T) {
	t.Parallel()
	out := alpaca(t, "data", "news", "--symbols", "AAPL", "--limit", "5")
	data := parseJSONMap(t, out)
	requireFields(t, data, "news")
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
