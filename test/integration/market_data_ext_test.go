//go:build integration

package integration

import (
	"testing"
)

func TestDataQuotes(t *testing.T) {
	t.Parallel()
	out := alpaca(t, "data", "quotes", "AAPL",
		"--start", daysAgo(95),
		"--end", daysAgo(94),
		"--limit", "5",
	)
	data := parseJSONMap(t, out)
	requireFields(t, data, "quotes")
}

func TestDataTrades(t *testing.T) {
	t.Parallel()
	out := alpaca(t, "data", "trades", "AAPL",
		"--start", daysAgo(95),
		"--end", daysAgo(94),
		"--limit", "5",
	)
	data := parseJSONMap(t, out)
	requireFields(t, data, "trades")
}

func TestDataLatestBar(t *testing.T) {
	t.Parallel()
	out := alpaca(t, "data", "latest", "bar", "AAPL")
	data := parseJSONMap(t, out)
	requireFields(t, data, "bar")
}

func TestDataBars_Timeframe(t *testing.T) {
	t.Parallel()
	out := alpaca(t, "data", "bars", "AAPL",
		"--start", daysAgo(95),
		"--end", daysAgo(94),
		"--timeframe", "1Hour",
	)
	data := parseJSONMap(t, out)
	requireFields(t, data, "bars")
}

func TestDataBars_Adjustment(t *testing.T) {
	t.Parallel()
	out := alpaca(t, "data", "bars", "AAPL",
		"--start", daysAgo(100),
		"--end", daysAgo(93),
		"--timeframe", "1Day",
		"--adjustment", "split",
	)
	data := parseJSONMap(t, out)
	requireFields(t, data, "bars")
}
