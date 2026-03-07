//go:build integration

package integration

import (
	"testing"
)

func TestDataBars(t *testing.T) {
	out := alpaca(t, "data", "bars", "AAPL",
		"--start", "2025-01-02",
		"--end", "2025-01-10",
		"--timeframe", "1Day",
		"--json",
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
	out := alpaca(t, "data", "latest", "trade", "AAPL", "--json")
	data := parseJSONMap(t, out)

	for _, field := range []string{"t", "p", "s"} {
		if _, ok := data[field]; !ok {
			t.Errorf("latest trade missing field: %s", field)
		}
	}
}

func TestDataLatestQuote(t *testing.T) {
	out := alpaca(t, "data", "latest", "quote", "AAPL", "--json")
	data := parseJSONMap(t, out)

	for _, field := range []string{"bp", "ap", "bs", "as"} {
		if _, ok := data[field]; !ok {
			t.Errorf("latest quote missing field: %s", field)
		}
	}
}

func TestDataSnapshot(t *testing.T) {
	out := alpaca(t, "data", "snapshot", "AAPL", "--json")
	data := parseJSONMap(t, out)

	for _, key := range []string{"latestTrade", "latestQuote", "minuteBar", "dailyBar"} {
		if _, ok := data[key]; !ok {
			t.Errorf("snapshot missing key: %s", key)
		}
	}
}

func TestDataNews(t *testing.T) {
	out := alpaca(t, "data", "news", "--symbols", "AAPL", "--limit", "5", "--json")
	news := parseJSONArray(t, out)

	if len(news) == 0 {
		t.Fatal("expected at least one news article")
	}
	if _, ok := news[0]["headline"]; !ok {
		t.Error("news article missing 'headline'")
	}
}
