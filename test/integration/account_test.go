//go:build integration

package integration

import (
	"testing"
)

func TestAccountConfigSet(t *testing.T) {
	// Read current config
	out := alpaca(t, "account", "config", "get")
	original := parseJSONMap(t, out)
	origVal, _ := original["trade_confirm_email"].(string)

	// Pick a different value
	newVal := "none"
	if origVal == "none" {
		newVal = "all"
	}

	// Set it
	alpaca(t, "account", "config", "set", "--trade-confirm-email", newVal)

	// Restore original on cleanup
	t.Cleanup(func() {
		_ = makeCmd("account", "config", "set", "--trade-confirm-email", origVal).Run()
	})

	// Verify it changed
	out = alpaca(t, "account", "config", "get")
	updated := parseJSONMap(t, out)
	if updated["trade_confirm_email"] != newVal {
		t.Errorf("expected trade_confirm_email %q, got %v", newVal, updated["trade_confirm_email"])
	}
}

func TestAccountActivityList_WithType(t *testing.T) {
	t.Parallel()
	out := alpaca(t, "account", "activity", "list", "--activity-types", "FILL", "--page-size", "5")
	_ = parseJSONArray(t, out)
}

func TestAccountActivityListByType(t *testing.T) {
	t.Parallel()
	out := alpaca(t, "account", "activity", "list-by-type", "--activity-type", "FILL", "--page-size", "5")
	_ = parseJSONArray(t, out)
}

func TestAccountActivityList_Pagination(t *testing.T) {
	t.Parallel()
	out := alpaca(t, "account", "activity", "list", "--page-size", "2", "--direction", "asc")
	_ = parseJSONArray(t, out)
}

func TestAccountActivityList_PageToken(t *testing.T) {
	t.Parallel()
	out := alpaca(t, "account", "activity", "list", "--page-size", "1", "--direction", "asc")
	page1 := parseJSONArray(t, out)
	if len(page1) == 0 {
		t.Skip("no activities to paginate")
	}

	token, ok := page1[len(page1)-1]["id"].(string)
	if !ok || token == "" {
		t.Skip("activity missing id for page-token")
	}

	out = alpaca(t, "account", "activity", "list", "--page-size", "1", "--direction", "asc", "--page-token", token)
	page2 := parseJSONArray(t, out)
	if len(page2) == 0 {
		t.Skip("only one page of activities")
	}

	if page2[0]["id"] == page1[0]["id"] {
		t.Error("page-token pagination returned the same activity as page 1")
	}
}

func TestAccountPortfolio_WithParams(t *testing.T) {
	t.Parallel()
	out := alpaca(t, "account", "portfolio", "--period", "1W", "--timeframe", "1D")
	data := parseJSONMap(t, out)
	requireFields(t, data, "equity", "timestamp")
}
