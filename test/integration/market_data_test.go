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

	data := parseJSONMap(t, out)
	requireFields(t, data, "bars")
}

func TestDataLatestTrade(t *testing.T) {
	t.Parallel()
	out := alpaca(t, "data", "latest", "trade", "AAPL")
	data := parseJSONMap(t, out)
	requireFields(t, data, "trade")
}

func TestDataLatestQuote(t *testing.T) {
	t.Parallel()
	out := alpaca(t, "data", "latest", "quote", "AAPL")
	data := parseJSONMap(t, out)
	requireFields(t, data, "quote")
}

func TestDataLatestTrade_Crypto(t *testing.T) {
	t.Parallel()
	out := alpaca(t, "data", "latest", "trade", "BTC/USD")
	data := parseJSONMap(t, out)
	requireFields(t, data, "trades")
}

func TestDataLatestQuote_Crypto(t *testing.T) {
	t.Parallel()
	out := alpaca(t, "data", "latest", "quote", "BTC/USD")
	data := parseJSONMap(t, out)
	requireFields(t, data, "quotes")
}

func TestDataLatestBar_Crypto(t *testing.T) {
	t.Parallel()
	out := alpaca(t, "data", "latest", "bar", "BTC/USD")
	data := parseJSONMap(t, out)
	requireFields(t, data, "bars")
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
	data := parseJSONMap(t, out)
	requireFields(t, data, "news")
}
