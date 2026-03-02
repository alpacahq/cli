//go:build integration

package integration

import (
	"testing"
)

func TestPositionList(t *testing.T) {
	out := alpaca(t, "position", "list", "--json")
	_ = parseJSONArray(t, out)
}

func TestPositionGetNotFound(t *testing.T) {
	_, stderr, code := alpacaFail(t, "position", "get", "ZZZZZZ", "--json")
	if code == 0 {
		t.Fatal("expected non-zero exit code for invalid symbol")
	}

	errMap := parseJSONMap(t, stderr)
	if errMap["error"] == nil || errMap["error"] == "" {
		t.Error("expected error message in JSON error output")
	}
}

func TestPositionsShortcut(t *testing.T) {
	out := alpaca(t, "positions", "--json")
	_ = parseJSONArray(t, out)
}
