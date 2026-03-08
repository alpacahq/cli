//go:build integration

package integration

import (
	"testing"
)

func TestClock_V3_Markets(t *testing.T) {
	out := alpaca(t, "clock", "--markets", "XNYS", "--json")
	data := parseJSONMap(t, out)
	if len(data) == 0 {
		t.Error("expected non-empty clock v3 response")
	}
}

func TestCalendar_V3_Market(t *testing.T) {
	out := alpaca(t, "calendar", "--market", "XNYS",
		"--start", "2025-01-01", "--end", "2025-01-31", "--json")
	// V3 calendar returns {"calendar": [...], "market": {...}}
	data := parseJSONMap(t, out)
	cal, ok := data["calendar"].([]any)
	if !ok || len(cal) == 0 {
		t.Fatal("expected non-empty calendar array in v3 response")
	}
}
