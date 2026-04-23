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
	bars, ok := data["bars"].([]any)
	if !ok || len(bars) == 0 {
		t.Fatal("expected non-empty bars array")
	}
	first, _ := bars[0].(map[string]any)
	requireFields(t, first, "t", "o", "h", "l", "c", "v")
}

func TestDataBars_Timeframe(t *testing.T) {
	t.Parallel()
	out := alpacaRetry(t, "data", "bars", "--symbol", "AAPL",
		"--start", daysAgo(100),
		"--end", daysAgo(93),
		"--timeframe", "1Hour",
	)
	data := parseJSONMap(t, out)
	bars, ok := data["bars"].([]any)
	if !ok || len(bars) == 0 {
		t.Fatal("expected non-empty bars array for 1Hour timeframe")
	}
	first, _ := bars[0].(map[string]any)
	requireFields(t, first, "t", "o", "h", "l", "c", "v")
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
	bars, ok := data["bars"].([]any)
	if !ok || len(bars) == 0 {
		t.Fatal("expected non-empty bars array with split adjustment")
	}
	first, _ := bars[0].(map[string]any)
	requireFields(t, first, "t", "o", "h", "l", "c", "v")
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
	bars, ok := data["bars"].([]any)
	if !ok || len(bars) == 0 {
		t.Fatal("expected non-empty bars array with iex feed")
	}
	first, _ := bars[0].(map[string]any)
	requireFields(t, first, "t", "o", "h", "l", "c", "v")
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
	bars, ok := data["bars"].([]any)
	if !ok || len(bars) == 0 {
		t.Fatal("expected non-empty bars array with USD currency")
	}
	first, _ := bars[0].(map[string]any)
	requireFields(t, first, "t", "o", "h", "l", "c", "v")
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
		"--start", daysAgo(100),
		"--end", daysAgo(93),
		"--limit", "5",
	)
	data := parseJSONMap(t, out)
	quotes, ok := data["quotes"].([]any)
	if !ok || len(quotes) == 0 {
		t.Fatal("expected non-empty quotes array")
	}
	first, _ := quotes[0].(map[string]any)
	requireFields(t, first, "t", "bp", "ap", "bs", "as")
}

func TestDataTrades(t *testing.T) {
	t.Parallel()
	out := alpacaRetry(t, "data", "trades", "--symbol", "AAPL",
		"--start", daysAgo(100),
		"--end", daysAgo(93),
		"--limit", "5",
	)
	data := parseJSONMap(t, out)
	trades, ok := data["trades"].([]any)
	if !ok || len(trades) == 0 {
		t.Fatal("expected non-empty trades array")
	}
	first, _ := trades[0].(map[string]any)
	requireFields(t, first, "t", "p", "s")
}

func TestDataLatestTrade(t *testing.T) {
	t.Parallel()
	out := alpacaRetry(t, "data", "latest-trade", "--symbol", "AAPL")
	data := parseJSONMap(t, out)
	trade, ok := data["trade"].(map[string]any)
	if !ok {
		t.Fatal("expected trade object")
	}
	requireFields(t, trade, "t", "p", "s")
}

func TestDataLatestQuote(t *testing.T) {
	t.Parallel()
	out := alpacaRetry(t, "data", "latest-quote", "--symbol", "AAPL")
	data := parseJSONMap(t, out)
	quote, ok := data["quote"].(map[string]any)
	if !ok {
		t.Fatal("expected quote object")
	}
	requireFields(t, quote, "t", "bp", "ap", "bs", "as")
}

func TestDataLatestBar(t *testing.T) {
	t.Parallel()
	out := alpacaRetry(t, "data", "latest-bar", "--symbol", "AAPL")
	data := parseJSONMap(t, out)
	bar, ok := data["bar"].(map[string]any)
	if !ok {
		t.Fatal("expected bar object")
	}
	requireFields(t, bar, "t", "o", "h", "l", "c", "v")
}

func TestDataSnapshot(t *testing.T) {
	t.Parallel()
	out := alpacaRetry(t, "data", "snapshot", "--symbol", "AAPL")
	data := parseJSONMap(t, out)
	requireFields(t, data, "latestTrade", "latestQuote", "minuteBar", "dailyBar")

	trade, _ := data["latestTrade"].(map[string]any)
	requireFields(t, trade, "t", "p", "s")
	quote, _ := data["latestQuote"].(map[string]any)
	requireFields(t, quote, "t", "bp", "ap")
	bar, _ := data["dailyBar"].(map[string]any)
	requireFields(t, bar, "t", "o", "h", "l", "c", "v")
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
		"--start", daysAgo(100),
		"--end", daysAgo(93),
		"--limit", "5",
	)
	data := parseJSONMap(t, out)
	requireFields(t, data, "quotes")
	quotes := data["quotes"].(map[string]any)
	if len(quotes) == 0 {
		t.Fatal("expected at least one symbol in quotes")
	}
	for sym, v := range quotes {
		items, ok := v.([]any)
		if !ok || len(items) == 0 {
			t.Errorf("expected non-empty quotes for %s", sym)
			continue
		}
		first, _ := items[0].(map[string]any)
		requireFields(t, first, "t", "bp", "ap", "bs", "as")
		break
	}
}

func TestDataMultiTrades(t *testing.T) {
	t.Parallel()
	out := alpacaRetry(t, "data", "multi-trades",
		"--symbols", "AAPL,MSFT",
		"--start", daysAgo(100),
		"--end", daysAgo(93),
		"--limit", "5",
	)
	data := parseJSONMap(t, out)
	requireFields(t, data, "trades")
	trades := data["trades"].(map[string]any)
	if len(trades) == 0 {
		t.Fatal("expected at least one symbol in trades")
	}
	for sym, v := range trades {
		items, ok := v.([]any)
		if !ok || len(items) == 0 {
			t.Errorf("expected non-empty trades for %s", sym)
			continue
		}
		first, _ := items[0].(map[string]any)
		requireFields(t, first, "t", "p", "s")
		break
	}
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
		"--start", daysAgo(100),
		"--end", daysAgo(93),
		"--limit", "2",
	)
	page1 := parseJSONMap(t, out)
	token, ok := page1["next_page_token"].(string)
	if !ok || token == "" {
		t.Skip("no next_page_token - not enough data to test pagination")
	}

	out = alpacaRetry(t, "data", "multi-quotes",
		"--symbols", "AAPL",
		"--start", daysAgo(100),
		"--end", daysAgo(93),
		"--limit", "2",
		"--page-token", token,
	)
	page2 := parseJSONMap(t, out)
	requireFields(t, page2, "quotes")
}
