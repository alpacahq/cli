//go:build integration

package integration

import (
	"testing"
)

func TestDataBars(t *testing.T) {
	out := alpaca(t, "data", "bars",
		"--symbol", "AAPL",
		"--start", "2025-01-02",
		"--end", "2025-01-10",
		"--timeframe", "1Day",
		"-o", "json",
	)

	bars := parseJSONArray(t, out)
	if len(bars) == 0 {
		t.Fatal("expected at least one bar")
	}

	bar := bars[0]
	for _, field := range []string{"t", "o", "h", "l", "c", "v"} {
		if _, ok := bar[field]; !ok {
			t.Errorf("bar missing field: %s", field)
		}
	}
}

func TestDataLatestTrade(t *testing.T) {
	out := alpaca(t, "data", "latest", "trade", "AAPL", "-o", "json")
	data := parseJSONMap(t, out)

	// Response has "trade" key for single stock
	if _, ok := data["trade"]; !ok {
		t.Error("latest trade response missing 'trade' key")
	}
}

func TestDataLatestQuote(t *testing.T) {
	out := alpaca(t, "data", "latest", "quote", "AAPL", "-o", "json")
	data := parseJSONMap(t, out)

	if _, ok := data["quote"]; !ok {
		t.Error("latest quote response missing 'quote' key")
	}
}

func TestDataSnapshot(t *testing.T) {
	out := alpaca(t, "data", "snapshot", "AAPL", "-o", "json")
	data := parseJSONMap(t, out)

	for _, key := range []string{"latestTrade", "latestQuote", "minuteBar", "dailyBar"} {
		if _, ok := data[key]; !ok {
			t.Errorf("snapshot missing key: %s", key)
		}
	}
}

func TestDataNews(t *testing.T) {
	out := alpaca(t, "news", "--symbols", "AAPL", "--limit", "5", "-o", "json")
	news := parseJSONArray(t, out)

	if len(news) == 0 {
		t.Fatal("expected at least one news article")
	}
	if _, ok := news[0]["headline"]; !ok {
		t.Error("news article missing 'headline'")
	}
}
