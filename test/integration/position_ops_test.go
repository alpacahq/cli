//go:build integration

package integration

import (
	"testing"
	"time"
)

func TestPositionGet_Success(t *testing.T) {
	symbol := submitCryptoFill(t)

	out := alpaca(t, "position", "get", symbol, "--json")
	pos := parseJSONMap(t, out)
	requireFields(t, pos, "symbol", "qty", "market_value", "avg_entry_price")
}

func TestPositionClose(t *testing.T) {
	symbol := submitCryptoFill(t)

	out := alpaca(t, "position", "close", symbol, "--json")
	closed := parseJSONMap(t, out)
	requireFields(t, closed, "id", "symbol", "status")

	pollFor(t, 10*time.Second, "position to be closed", func() bool {
		_, _, code := alpacaWithStderr(t, "position", "get", symbol, "--json")
		return code != 0
	})
}

func TestPositionCloseAll(t *testing.T) {
	_ = submitCryptoFill(t)

	alpaca(t, "position", "close-all")

	pollFor(t, 10*time.Second, "all positions to be closed", func() bool {
		out := alpaca(t, "position", "list", "--json")
		return len(parseJSONArray(t, out)) == 0
	})
}
