//go:build integration

package integration

import (
	"testing"
	"time"
)

func TestPositionList(t *testing.T) {
	t.Parallel()
	out := alpaca(t, "position", "list")
	_ = parseJSONArray(t, out)
}

func TestPositionGetNotFound(t *testing.T) {
	t.Parallel()
	_, stderr, code := alpacaFail(t, "position", "get", "--symbol-or-asset-id", "ZZZZZZ")
	if code == 0 {
		t.Fatal("expected non-zero exit code for invalid symbol")
	}

	errMap := parseJSONMap(t, stderr)
	if errMap["error"] == nil || errMap["error"] == "" {
		t.Error("expected error message in JSON error output")
	}
}

func TestPositionOps(t *testing.T) {
	symbol := submitCryptoFill(t, "BTC/USD")

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

func TestPositionClose_Percentage(t *testing.T) {
	symbol := submitCryptoFill(t, "LTC/USD")

	out := alpaca(t, "position", "close", "--symbol-or-asset-id", symbol, "--percentage", "50")
	closed := parseJSONMap(t, out)
	requireFields(t, closed, "id", "symbol", "status")

	pollFor(t, 10*time.Second, "position to still exist after partial close", func() bool {
		stdout, _, code := alpacaWithStderr(t, "position", "get", "--symbol-or-asset-id", symbol)
		if code != 0 {
			return false
		}
		pos := parseJSONMap(t, stdout)
		return pos["symbol"] != nil
	})
}

func TestPositionCloseAll(t *testing.T) {
	_ = submitCryptoFill(t, "BCH/USD")

	alpaca(t, "position", "close-all")

	pollFor(t, 10*time.Second, "all positions to be closed", func() bool {
		out := alpaca(t, "position", "list")
		return len(parseJSONArray(t, out)) == 0
	})
}

func TestPositionCloseAll_CancelOrders(t *testing.T) {
	_ = submitCryptoFill(t, "DOGE/USD")
	_ = submitTestOrder(t, "COST")

	alpaca(t, "position", "close-all", "--cancel-orders")

	pollFor(t, 10*time.Second, "all positions to be closed", func() bool {
		out := alpaca(t, "position", "list")
		return len(parseJSONArray(t, out)) == 0
	})
	pollFor(t, 10*time.Second, "all orders to be canceled", func() bool {
		out := alpaca(t, "order", "list", "--status", "open")
		return len(parseJSONArray(t, out)) == 0
	})
}
