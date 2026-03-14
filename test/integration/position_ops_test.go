//go:build integration

package integration

import (
	"testing"
	"time"
)

func TestPositionOps(t *testing.T) {
	symbol := submitCryptoFill(t)

	t.Run("get", func(t *testing.T) {
		out := alpaca(t, "position", "get", "--symbol-or-asset-id", symbol)
		pos := parseJSONMap(t, out)
		requireFields(t, pos, "symbol", "qty", "market_value", "avg_entry_price")
	})

	t.Run("close", func(t *testing.T) {
		out := alpaca(t, "position", "close", "--symbol-or-asset-id", symbol)
		closed := parseJSONMap(t, out)
		requireFields(t, closed, "id", "symbol", "status")

		pollFor(t, 10*time.Second, "position to be closed", func() bool {
			_, _, code := alpacaWithStderr(t, "position", "get", "--symbol-or-asset-id", symbol)
			return code != 0
		})
	})
}

func TestPositionCloseAll(t *testing.T) {
	_ = submitCryptoFill(t)

	alpaca(t, "position", "close-all")

	pollFor(t, 10*time.Second, "all positions to be closed", func() bool {
		out := alpaca(t, "position", "list")
		return len(parseJSONArray(t, out)) == 0
	})
}
