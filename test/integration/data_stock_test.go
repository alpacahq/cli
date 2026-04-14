//go:build integration

package integration

import (
	"testing"
)

// --- Single-symbol ---

func TestDataBars(t *testing.T) {
	t.Parallel()
	out := alpacaRetry(t, "data", "bars", "--symbol", "AAPL",
		"--start", daysAgo(100),
		"--end", daysAgo(93),
		"--timeframe", "1Day",
	)
	data := parseJSONMap(t, out)
	requireFields(t, data, "bars")
}

func TestDataBars_Timeframe(t *testing.T) {
	t.Parallel()
	out := alpacaRetry(t, "data", "bars", "--symbol", "AAPL",
		"--start", daysAgo(95),
		"--end", daysAgo(94),
		"--timeframe", "1Hour",
	)
	data := parseJSONMap(t, out)
	requireFields(t, data, "bars")
}

func TestDataBars_Adjustment(t *testing.T) {
	t.Parallel()
	out := alpacaRetry(t, "data", "bars", "--symbol", "AAPL",
		"--start", daysAgo(100),
		"--end", daysAgo(93),
		"--timeframe", "1Day",
		"--adjustment", "split",
	)
	data := parseJSONMap(t, out)
	requireFields(t, data, "bars")
}

func TestDataBars_Feed(t *testing.T) {
	t.Parallel()
	out := alpacaRetry(t, "data", "bars", "--symbol", "AAPL",
		"--start", daysAgo(100),
		"--end", daysAgo(93),
		"--timeframe", "1Day",
		"--feed", "iex",
	)
	data := parseJSONMap(t, out)
	requireFields(t, data, "bars")
}

func TestDataBars_Currency(t *testing.T) {
	t.Parallel()
	out := alpacaRetry(t, "data", "bars", "--symbol", "AAPL",
		"--start", daysAgo(100),
		"--end", daysAgo(93),
		"--timeframe", "1Day",
		"--currency", "USD",
	)
	data := parseJSONMap(t, out)
	requireFields(t, data, "bars")
}

func TestDataBars_Sort(t *testing.T) {
	t.Parallel()
	outAsc := alpacaRetry(t, "data", "bars", "--symbol", "AAPL",
		"--start", daysAgo(100),
		"--end", daysAgo(90),
		"--timeframe", "1Day",
		"--sort", "asc",
	)
	outDesc := alpacaRetry(t, "data", "bars", "--symbol", "AAPL",
		"--start", daysAgo(100),
		"--end", daysAgo(90),
		"--timeframe", "1Day",
		"--sort", "desc",
	)
	asc := parseJSONMap(t, outAsc)
	desc := parseJSONMap(t, outDesc)
	requireFields(t, asc, "bars")
	requireFields(t, desc, "bars")

	ascBars, _ := asc["bars"].([]any)
	descBars, _ := desc["bars"].([]any)
	if len(ascBars) < 2 || len(descBars) < 2 {
		t.Skip("not enough bars to verify sort order")
	}
	firstAsc, _ := ascBars[0].(map[string]any)
	firstDesc, _ := descBars[0].(map[string]any)
	if firstAsc["t"] == firstDesc["t"] {
		t.Error("expected different first timestamps for asc vs desc sort")
	}
}

func TestDataQuotes(t *testing.T) {
	t.Parallel()
	out := alpacaRetry(t, "data", "quotes", "--symbol", "AAPL",
		"--start", daysAgo(95),
		"--end", daysAgo(94),
		"--limit", "5",
	)
	data := parseJSONMap(t, out)
	requireFields(t, data, "quotes")
}

func TestDataTrades(t *testing.T) {
	t.Parallel()
	out := alpacaRetry(t, "data", "trades", "--symbol", "AAPL",
		"--start", daysAgo(95),
		"--end", daysAgo(94),
		"--limit", "5",
	)
	data := parseJSONMap(t, out)
	requireFields(t, data, "trades")
}

func TestDataLatestTrade(t *testing.T) {
	t.Parallel()
	out := alpacaRetry(t, "data", "latest-trade", "--symbol", "AAPL")
	data := parseJSONMap(t, out)
	requireFields(t, data, "trade")
}

func TestDataLatestQuote(t *testing.T) {
	t.Parallel()
	out := alpacaRetry(t, "data", "latest-quote", "--symbol", "AAPL")
	data := parseJSONMap(t, out)
	requireFields(t, data, "quote")
}

func TestDataLatestBar(t *testing.T) {
	t.Parallel()
	out := alpacaRetry(t, "data", "latest-bar", "--symbol", "AAPL")
	data := parseJSONMap(t, out)
	requireFields(t, data, "bar")
}

func TestDataSnapshot(t *testing.T) {
	t.Parallel()
	out := alpacaRetry(t, "data", "snapshot", "--symbol", "AAPL")
	data := parseJSONMap(t, out)
	requireFields(t, data, "latestTrade", "latestQuote", "minuteBar", "dailyBar")
}

func TestDataAuction(t *testing.T) {
	t.Parallel()
	out := alpacaRetry(t, "data", "auction", "--symbol", "AAPL",
		"--start", daysAgo(100),
		"--end", daysAgo(93),
	)
	data := parseJSONMap(t, out)
	requireFields(t, data, "auctions")
}

func TestDataAuctions(t *testing.T) {
	t.Parallel()
	out := alpacaRetry(t, "data", "auctions", "--symbols", "AAPL", "--start", daysAgo(100), "--end", daysAgo(93))
	data := parseJSONMap(t, out)
	if len(data) == 0 {
		t.Error("expected non-empty auctions response")
	}
}

// --- Multi-symbol ---

func TestDataMultiBars(t *testing.T) {
	t.Parallel()
	out := alpacaRetry(t, "data", "multi-bars",
		"--symbols", "AAPL,MSFT",
		"--timeframe", "1Day",
		"--start", daysAgo(100),
		"--end", daysAgo(93),
	)
	data := parseJSONMap(t, out)
	requireFields(t, data, "bars")
	bars := data["bars"].(map[string]any)
	if _, ok := bars["AAPL"]; !ok {
		t.Error("expected bars to contain AAPL")
	}
	if _, ok := bars["MSFT"]; !ok {
		t.Error("expected bars to contain MSFT")
	}
}

func TestDataMultiQuotes(t *testing.T) {
	t.Parallel()
	out := alpacaRetry(t, "data", "multi-quotes",
		"--symbols", "AAPL,MSFT",
		"--start", daysAgo(95),
		"--end", daysAgo(94),
		"--limit", "2",
	)
	data := parseJSONMap(t, out)
	requireFields(t, data, "quotes")
}

func TestDataMultiTrades(t *testing.T) {
	t.Parallel()
	out := alpacaRetry(t, "data", "multi-trades",
		"--symbols", "AAPL,MSFT",
		"--start", daysAgo(95),
		"--end", daysAgo(94),
		"--limit", "2",
	)
	data := parseJSONMap(t, out)
	requireFields(t, data, "trades")
}

func TestDataMultiSnapshots(t *testing.T) {
	t.Parallel()
	out := alpacaRetry(t, "data", "multi-snapshots", "--symbols", "AAPL,MSFT")
	data := parseJSONMap(t, out)
	if _, ok := data["AAPL"]; !ok {
		t.Error("expected snapshot to contain AAPL")
	}
	if _, ok := data["MSFT"]; !ok {
		t.Error("expected snapshot to contain MSFT")
	}
}

func TestDataLatestBarsMulti(t *testing.T) {
	t.Parallel()
	out := alpacaRetry(t, "data", "latest-bars", "--symbols", "AAPL,MSFT")
	data := parseJSONMap(t, out)
	requireFields(t, data, "bars")
	bars := data["bars"].(map[string]any)
	if _, ok := bars["AAPL"]; !ok {
		t.Error("expected latest bars to contain AAPL")
	}
}

func TestDataLatestQuotesMulti(t *testing.T) {
	t.Parallel()
	out := alpacaRetry(t, "data", "latest-quotes", "--symbols", "AAPL,MSFT")
	data := parseJSONMap(t, out)
	requireFields(t, data, "quotes")
	quotes := data["quotes"].(map[string]any)
	if _, ok := quotes["AAPL"]; !ok {
		t.Error("expected latest quotes to contain AAPL")
	}
}

func TestDataLatestTradesMulti(t *testing.T) {
	t.Parallel()
	out := alpacaRetry(t, "data", "latest-trades", "--symbols", "AAPL,MSFT")
	data := parseJSONMap(t, out)
	requireFields(t, data, "trades")
	trades := data["trades"].(map[string]any)
	if _, ok := trades["AAPL"]; !ok {
		t.Error("expected latest trades to contain AAPL")
	}
}

// --- Pagination ---

func TestDataMultiQuotes_PageToken(t *testing.T) {
	t.Parallel()
	out := alpacaRetry(t, "data", "multi-quotes",
		"--symbols", "AAPL",
		"--start", daysAgo(95),
		"--end", daysAgo(94),
		"--limit", "2",
	)
	page1 := parseJSONMap(t, out)
	token, ok := page1["next_page_token"].(string)
	if !ok || token == "" {
		t.Skip("no next_page_token — not enough data to test pagination")
	}

	out = alpacaRetry(t, "data", "multi-quotes",
		"--symbols", "AAPL",
		"--start", daysAgo(95),
		"--end", daysAgo(94),
		"--limit", "2",
		"--page-token", token,
	)
	page2 := parseJSONMap(t, out)
	requireFields(t, page2, "quotes")
}
