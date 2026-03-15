//go:build integration

package integration

import (
	"testing"
)

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

func TestDataCryptoOrderbook(t *testing.T) {
	t.Parallel()
	out := alpaca(t, "data", "crypto-orderbook", "--symbols", "BTC/USD")
	data := parseJSONMap(t, out)
	if len(data) == 0 {
		t.Error("expected non-empty crypto orderbook response")
	}
}
