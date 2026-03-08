//go:build integration

package integration

import (
	"testing"
)

func TestAccountConfigSet(t *testing.T) {
	// Read current config
	out := alpaca(t, "account", "config", "get", "--json")
	original := parseJSONMap(t, out)
	origVal, _ := original["trade_confirm_email"].(string)

	// Pick a different value
	newVal := "none"
	if origVal == "none" {
		newVal = "all"
	}

	// Set it
	alpaca(t, "account", "config", "set", "--trade-confirm-email", newVal, "--json")

	// Restore original on cleanup
	t.Cleanup(func() {
		_ = makeCmd("account", "config", "set", "--trade-confirm-email", origVal).Run()
	})

	// Verify it changed
	out = alpaca(t, "account", "config", "get", "--json")
	updated := parseJSONMap(t, out)
	if updated["trade_confirm_email"] != newVal {
		t.Errorf("expected trade_confirm_email %q, got %v", newVal, updated["trade_confirm_email"])
	}
}

func TestAccountActivityList_WithType(t *testing.T) {
	out := alpaca(t, "account", "activity", "list", "--activity-types", "FILL", "--page-size", "5", "--json")
	_ = parseJSONArray(t, out)
}

func TestAccountActivityList_Pagination(t *testing.T) {
	out := alpaca(t, "account", "activity", "list", "--page-size", "2", "--direction", "asc", "--json")
	_ = parseJSONArray(t, out)
}

func TestAccountPortfolio_WithParams(t *testing.T) {
	out := alpaca(t, "account", "portfolio", "--period", "1W", "--timeframe", "1D", "--json")
	data := parseJSONMap(t, out)
	requireFields(t, data, "equity", "timestamp")
}
