//go:build integration

package integration

import (
	"testing"
)

func TestCryptoPerpDataLatestBars(t *testing.T) {
	t.Parallel()
	out := alpacaRetry(t, "crypto-perp", "data", "latest-bars", "--symbols", "BTC.P", "--loc", "global")
	data := parseJSONMap(t, out)
	requireFields(t, data, "bars")
}

func TestCryptoPerpDataLatestTrades(t *testing.T) {
	t.Parallel()
	out := alpacaRetry(t, "crypto-perp", "data", "latest-trades", "--symbols", "BTC.P", "--loc", "global")
	data := parseJSONMap(t, out)
	requireFields(t, data, "trades")
}

func TestCryptoPerpDataLatestQuotes(t *testing.T) {
	t.Parallel()
	out := alpacaRetry(t, "crypto-perp", "data", "latest-quotes", "--symbols", "BTC.P", "--loc", "global")
	data := parseJSONMap(t, out)
	requireFields(t, data, "quotes")
}

func TestCryptoPerpDataLatestOrderbooks(t *testing.T) {
	t.Parallel()
	out := alpacaRetry(t, "crypto-perp", "data", "latest-orderbooks", "--symbols", "BTC.P", "--loc", "global")
	data := parseJSONMap(t, out)
	requireFields(t, data, "orderbooks")
}

func TestCryptoPerpDataLatestFuturesPricing(t *testing.T) {
	t.Parallel()
	out := alpacaRetry(t, "crypto-perp", "data", "latest-futures-pricing", "--symbols", "BTC.P", "--loc", "global")
	data := parseJSONMap(t, out)
	requireFields(t, data, "pricing")
}
