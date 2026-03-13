//go:build integration

package integration

import (
	"testing"
)

func TestDataBars(t *testing.T) {
	t.Parallel()
	out := alpaca(t, "data", "bars", "AAPL",
		"--start", daysAgo(100),
		"--end", daysAgo(93),
		"--timeframe", "1Day",
	)

	bars := requireArrayNonEmpty(t, out)
	requireFields(t, bars[0], "t", "o", "h", "l", "c", "v")
}

func TestDataLatestTrade(t *testing.T) {
	t.Parallel()
	out := alpaca(t, "data", "latest", "trade", "AAPL")
	data := parseJSONMap(t, out)
	requireFields(t, data, "t", "p", "s")
}

func TestDataLatestQuote(t *testing.T) {
	t.Parallel()
	out := alpaca(t, "data", "latest", "quote", "AAPL")
	data := parseJSONMap(t, out)
	requireFields(t, data, "bp", "ap", "bs", "as")
}

func TestDataSnapshot(t *testing.T) {
	t.Parallel()
	out := alpaca(t, "data", "snapshot", "AAPL")
	data := parseJSONMap(t, out)
	requireFields(t, data, "latestTrade", "latestQuote", "minuteBar", "dailyBar")
}

func TestDataNews(t *testing.T) {
	t.Parallel()
	out := alpaca(t, "data", "news", "--symbols", "AAPL", "--limit", "5")
	news := requireArrayNonEmpty(t, out)
	requireFields(t, news[0], "headline")
}
