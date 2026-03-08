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

	time.Sleep(2 * time.Second)

	// Verify position is gone
	_, _, code := alpacaWithStderr(t, "position", "get", symbol, "--json")
	if code == 0 {
		t.Error("position should be gone after close")
	}
}

func TestPositionCloseAll(t *testing.T) {
	_ = submitCryptoFill(t)

	alpaca(t, "position", "close-all")
	time.Sleep(2 * time.Second)

	out := alpaca(t, "position", "list", "--json")
	positions := parseJSONArray(t, out)
	if len(positions) > 0 {
		t.Errorf("expected 0 positions after close-all, got %d", len(positions))
	}
}
