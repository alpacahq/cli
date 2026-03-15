//go:build integration

package integration

import (
	"testing"
)

func TestDataMultiBars(t *testing.T) {
	t.Parallel()
	out := alpaca(t, "data", "multi-bars",
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
	out := alpaca(t, "data", "multi-quotes",
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
	out := alpaca(t, "data", "multi-trades",
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
	out := alpaca(t, "data", "multi-snapshots", "--symbols", "AAPL,MSFT")
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
	out := alpaca(t, "data", "latest-bars", "--symbols", "AAPL,MSFT")
	data := parseJSONMap(t, out)
	requireFields(t, data, "bars")
	bars := data["bars"].(map[string]any)
	if _, ok := bars["AAPL"]; !ok {
		t.Error("expected latest bars to contain AAPL")
	}
}

func TestDataLatestQuotesMulti(t *testing.T) {
	t.Parallel()
	out := alpaca(t, "data", "latest-quotes", "--symbols", "AAPL,MSFT")
	data := parseJSONMap(t, out)
	requireFields(t, data, "quotes")
	quotes := data["quotes"].(map[string]any)
	if _, ok := quotes["AAPL"]; !ok {
		t.Error("expected latest quotes to contain AAPL")
	}
}

func TestDataLatestTradesMulti(t *testing.T) {
	t.Parallel()
	out := alpaca(t, "data", "latest-trades", "--symbols", "AAPL,MSFT")
	data := parseJSONMap(t, out)
	requireFields(t, data, "trades")
	trades := data["trades"].(map[string]any)
	if _, ok := trades["AAPL"]; !ok {
		t.Error("expected latest trades to contain AAPL")
	}
}
