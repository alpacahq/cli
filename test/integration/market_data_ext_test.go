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
		"--json",
	)
	quotes := requireArrayNonEmpty(t, out)
	requireFields(t, quotes[0], "bp", "ap", "t")
}

func TestDataTrades(t *testing.T) {
	t.Parallel()
	out := alpaca(t, "data", "trades", "AAPL",
		"--start", daysAgo(95),
		"--end", daysAgo(94),
		"--limit", "5",
		"--json",
	)
	trades := requireArrayNonEmpty(t, out)
	requireFields(t, trades[0], "p", "s", "t")
}

func TestDataLatestBar(t *testing.T) {
	t.Parallel()
	out := alpaca(t, "data", "latest", "bar", "AAPL", "--json")
	bar := parseJSONMap(t, out)
	requireFields(t, bar, "o", "h", "l", "c", "v")
}

func TestDataBars_Timeframe(t *testing.T) {
	t.Parallel()
	out := alpaca(t, "data", "bars", "AAPL",
		"--start", daysAgo(95),
		"--end", daysAgo(94),
		"--timeframe", "1Hour",
		"--json",
	)
	bars := requireArrayNonEmpty(t, out)
	requireFields(t, bars[0], "o", "h", "l", "c", "v", "t")
}

func TestDataBars_Adjustment(t *testing.T) {
	t.Parallel()
	out := alpaca(t, "data", "bars", "AAPL",
		"--start", daysAgo(100),
		"--end", daysAgo(93),
		"--timeframe", "1Day",
		"--adjustment", "split",
		"--json",
	)
	bars := requireArrayNonEmpty(t, out)
	requireFields(t, bars[0], "o", "h", "l", "c", "v")
}
