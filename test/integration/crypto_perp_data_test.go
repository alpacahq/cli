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
	bars, _ := data["bars"].(map[string]any)
	if btc, ok := bars["BTC.P"].(map[string]any); ok {
		requireFields(t, btc, "t", "o", "h", "l", "c", "v")
	}
}

func TestCryptoPerpDataLatestTrades(t *testing.T) {
	t.Parallel()
	out := alpacaRetry(t, "crypto-perp", "data", "latest-trades", "--symbols", "BTC.P", "--loc", "global")
	data := parseJSONMap(t, out)
	requireFields(t, data, "trades")
	trades, _ := data["trades"].(map[string]any)
	if btc, ok := trades["BTC.P"].(map[string]any); ok {
		requireFields(t, btc, "t", "p", "s")
	}
}

func TestCryptoPerpDataLatestQuotes(t *testing.T) {
	t.Parallel()
	out := alpacaRetry(t, "crypto-perp", "data", "latest-quotes", "--symbols", "BTC.P", "--loc", "global")
	data := parseJSONMap(t, out)
	requireFields(t, data, "quotes")
	quotes, _ := data["quotes"].(map[string]any)
	if btc, ok := quotes["BTC.P"].(map[string]any); ok {
		requireFields(t, btc, "t", "bp", "ap", "bs", "as")
	}
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
