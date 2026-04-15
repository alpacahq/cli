//go:build integration

package integration

import (
	"testing"
)

func TestDataCryptoBars(t *testing.T) {
	t.Parallel()
	out := alpacaRetry(t, "data", "crypto", "bars",
		"--symbols", "BTC/USD",
		"--timeframe", "1Day",
		"--start", daysAgo(10),
		"--end", daysAgo(3),
	)
	data := parseJSONMap(t, out)
	requireFields(t, data, "bars")
	bars, _ := data["bars"].(map[string]any)
	btcBars, ok := bars["BTC/USD"].([]any)
	if !ok || len(btcBars) == 0 {
		t.Fatal("expected non-empty BTC/USD bars")
	}
	first, _ := btcBars[0].(map[string]any)
	requireFields(t, first, "t", "o", "h", "l", "c", "v")
}

func TestDataCryptoTrades(t *testing.T) {
	t.Parallel()
	out := alpacaRetry(t, "data", "crypto", "trades",
		"--symbols", "BTC/USD",
		"--start", daysAgo(5),
		"--end", daysAgo(4),
		"--limit", "5",
	)
	data := parseJSONMap(t, out)
	requireFields(t, data, "trades")
	trades, _ := data["trades"].(map[string]any)
	btcTrades, ok := trades["BTC/USD"].([]any)
	if !ok || len(btcTrades) == 0 {
		t.Fatal("expected non-empty BTC/USD trades")
	}
	first, _ := btcTrades[0].(map[string]any)
	requireFields(t, first, "t", "p", "s")
}

func TestDataCryptoQuotes(t *testing.T) {
	t.Parallel()
	out := alpacaRetry(t, "data", "crypto", "quotes",
		"--symbols", "BTC/USD",
		"--start", daysAgo(5),
		"--end", daysAgo(4),
		"--limit", "5",
	)
	data := parseJSONMap(t, out)
	requireFields(t, data, "quotes")
	quotes, _ := data["quotes"].(map[string]any)
	btcQuotes, ok := quotes["BTC/USD"].([]any)
	if !ok || len(btcQuotes) == 0 {
		t.Fatal("expected non-empty BTC/USD quotes")
	}
	first, _ := btcQuotes[0].(map[string]any)
	requireFields(t, first, "t", "bp", "ap", "bs", "as")
}

func TestDataCryptoSnapshots(t *testing.T) {
	t.Parallel()
	out := alpacaRetry(t, "data", "crypto", "snapshots", "--symbols", "BTC/USD")
	data := parseJSONMap(t, out)
	requireFields(t, data, "snapshots")
	snapshots, _ := data["snapshots"].(map[string]any)
	btc, ok := snapshots["BTC/USD"].(map[string]any)
	if !ok {
		t.Fatal("expected BTC/USD snapshot")
	}
	requireFields(t, btc, "latestTrade", "latestQuote", "dailyBar", "minuteBar")
}

func TestDataCryptoLatestTrades(t *testing.T) {
	t.Parallel()
	out := alpacaRetry(t, "data", "crypto", "latest-trades", "--symbols", "BTC/USD")
	data := parseJSONMap(t, out)
	requireFields(t, data, "trades")
	trades, _ := data["trades"].(map[string]any)
	btc, ok := trades["BTC/USD"].(map[string]any)
	if !ok {
		t.Fatal("expected BTC/USD in latest trades")
	}
	requireFields(t, btc, "t", "p", "s")
}

func TestDataCryptoLatestQuotes(t *testing.T) {
	t.Parallel()
	out := alpacaRetry(t, "data", "crypto", "latest-quotes", "--symbols", "BTC/USD")
	data := parseJSONMap(t, out)
	requireFields(t, data, "quotes")
	quotes, _ := data["quotes"].(map[string]any)
	btc, ok := quotes["BTC/USD"].(map[string]any)
	if !ok {
		t.Fatal("expected BTC/USD in latest quotes")
	}
	requireFields(t, btc, "t", "bp", "ap", "bs", "as")
}

func TestDataCryptoLatestBars(t *testing.T) {
	t.Parallel()
	out := alpacaRetry(t, "data", "crypto", "latest-bars", "--symbols", "BTC/USD")
	data := parseJSONMap(t, out)
	requireFields(t, data, "bars")
	bars, _ := data["bars"].(map[string]any)
	btc, ok := bars["BTC/USD"].(map[string]any)
	if !ok {
		t.Fatal("expected BTC/USD in latest bars")
	}
	requireFields(t, btc, "t", "o", "h", "l", "c", "v")
}

func TestDataCryptoOrderbook(t *testing.T) {
	t.Parallel()
	out := alpacaRetry(t, "data", "crypto-orderbook", "--symbols", "BTC/USD")
	data := parseJSONMap(t, out)
	requireFields(t, data, "orderbooks")
	orderbooks, _ := data["orderbooks"].(map[string]any)
	btc, ok := orderbooks["BTC/USD"].(map[string]any)
	if !ok {
		t.Fatal("expected BTC/USD in orderbooks")
	}
	requireFields(t, btc, "a", "b")
}
