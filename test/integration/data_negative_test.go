//go:build integration

package integration

import (
	"encoding/json"
	"testing"
)

func TestDataError_MissingRequiredSymbol(t *testing.T) {
	t.Parallel()
	_, _, code := alpacaFail(t, "data", "bars",
		"--start", daysAgo(100),
		"--end", daysAgo(93),
		"--timeframe", "1Day",
	)
	if code == 0 {
		t.Fatal("expected non-zero exit for missing required --symbol")
	}
}

func TestDataError_MissingRequiredTimeframe(t *testing.T) {
	t.Parallel()
	_, _, code := alpacaFail(t, "data", "bars",
		"--symbol", "AAPL",
		"--start", daysAgo(100),
		"--end", daysAgo(93),
	)
	if code == 0 {
		t.Fatal("expected non-zero exit for missing required --timeframe")
	}
}

func TestDataError_InvalidTimeframe(t *testing.T) {
	t.Parallel()
	_, stderr, code := alpacaFail(t, "data", "bars",
		"--symbol", "AAPL",
		"--start", daysAgo(100),
		"--end", daysAgo(93),
		"--timeframe", "bogus",
	)
	if code == 0 {
		t.Fatal("expected non-zero exit for invalid timeframe")
	}
	errMap := map[string]any{}
	if err := json.Unmarshal(stderr, &errMap); err == nil {
		if errMap["error"] == nil || errMap["error"] == "" {
			t.Error("expected error message in JSON error response")
		}
	}
}

func TestDataError_StartAfterEnd(t *testing.T) {
	t.Parallel()
	_, stderr, code := alpacaFail(t, "data", "bars",
		"--symbol", "AAPL",
		"--start", daysAgo(5),
		"--end", daysAgo(10),
		"--timeframe", "1Day",
	)
	if code == 0 {
		t.Fatal("expected non-zero exit when start is after end")
	}
	errMap := map[string]any{}
	if err := json.Unmarshal(stderr, &errMap); err == nil {
		if errMap["error"] == nil || errMap["error"] == "" {
			t.Error("expected error message in JSON error response")
		}
	}
}

func TestDataError_MissingRequiredSymbols_Multi(t *testing.T) {
	t.Parallel()
	_, _, code := alpacaFail(t, "data", "multi-bars",
		"--timeframe", "1Day",
		"--start", daysAgo(100),
		"--end", daysAgo(93),
	)
	if code == 0 {
		t.Fatal("expected non-zero exit for missing required --symbols")
	}
}
