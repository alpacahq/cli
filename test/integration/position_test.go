//go:build integration

package integration

import (
	"testing"
)

func TestPositionList(t *testing.T) {
	t.Parallel()
	out := alpaca(t, "position", "list")
	_ = parseJSONArray(t, out)
}

func TestPositionGetNotFound(t *testing.T) {
	t.Parallel()
	_, stderr, code := alpacaFail(t, "position", "get", "ZZZZZZ")
	if code == 0 {
		t.Fatal("expected non-zero exit code for invalid symbol")
	}

	errMap := parseJSONMap(t, stderr)
	if errMap["error"] == nil || errMap["error"] == "" {
		t.Error("expected error message in JSON error output")
	}
}
