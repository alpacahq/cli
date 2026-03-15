//go:build integration

package integration

import (
	"testing"
)

func TestDataBars(t *testing.T) {
	t.Parallel()
	out := alpaca(t, "data", "bars", "--symbol", "AAPL",
		"--start", daysAgo(100),
		"--end", daysAgo(93),
		"--timeframe", "1Day",
	)

	data := parseJSONMap(t, out)
	requireFields(t, data, "bars")
}

func TestDataLatestTrade(t *testing.T) {
	t.Parallel()
	out := alpaca(t, "data", "latest-trade", "--symbol", "AAPL")
	data := parseJSONMap(t, out)
	requireFields(t, data, "trade")
}

func TestDataLatestQuote(t *testing.T) {
	t.Parallel()
	out := alpaca(t, "data", "latest-quote", "--symbol", "AAPL")
	data := parseJSONMap(t, out)
	requireFields(t, data, "quote")
}

func TestDataCryptoLatestTrades(t *testing.T) {
	t.Parallel()
	out := alpaca(t, "data", "crypto", "latest-trades", "--symbols", "BTC/USD")
	data := parseJSONMap(t, out)
	requireFields(t, data, "trades")
}

func TestDataCryptoLatestQuotes(t *testing.T) {
	t.Parallel()
	out := alpaca(t, "data", "crypto", "latest-quotes", "--symbols", "BTC/USD")
	data := parseJSONMap(t, out)
	requireFields(t, data, "quotes")
}

func TestDataCryptoLatestBars(t *testing.T) {
	t.Parallel()
	out := alpaca(t, "data", "crypto", "latest-bars", "--symbols", "BTC/USD")
	data := parseJSONMap(t, out)
	requireFields(t, data, "bars")
}

func TestDataCryptoBars(t *testing.T) {
	t.Parallel()
	out := alpaca(t, "data", "crypto", "bars",
		"--symbols", "BTC/USD",
		"--timeframe", "1Day",
		"--start", daysAgo(10),
		"--end", daysAgo(3),
	)
	data := parseJSONMap(t, out)
	requireFields(t, data, "bars")
}

func TestDataCryptoTrades(t *testing.T) {
	t.Parallel()
	out := alpaca(t, "data", "crypto", "trades",
		"--symbols", "BTC/USD",
		"--start", daysAgo(5),
		"--end", daysAgo(4),
		"--limit", "5",
	)
	data := parseJSONMap(t, out)
	requireFields(t, data, "trades")
}

func TestDataCryptoQuotes(t *testing.T) {
	t.Parallel()
	out := alpaca(t, "data", "crypto", "quotes",
		"--symbols", "BTC/USD",
		"--start", daysAgo(5),
		"--end", daysAgo(4),
		"--limit", "5",
	)
	data := parseJSONMap(t, out)
	requireFields(t, data, "quotes")
}

func TestDataCryptoSnapshots(t *testing.T) {
	t.Parallel()
	out := alpaca(t, "data", "crypto", "snapshots", "--symbols", "BTC/USD")
	data := parseJSONMap(t, out)
	requireFields(t, data, "snapshots")
}

func TestDataSnapshot(t *testing.T) {
	t.Parallel()
	out := alpaca(t, "data", "snapshot", "--symbol", "AAPL")
	data := parseJSONMap(t, out)
	requireFields(t, data, "latestTrade", "latestQuote", "minuteBar", "dailyBar")
}

func TestDataNews(t *testing.T) {
	t.Parallel()
	out := alpaca(t, "data", "news", "--symbols", "AAPL", "--limit", "5")
	data := parseJSONMap(t, out)
	requireFields(t, data, "news")
}
