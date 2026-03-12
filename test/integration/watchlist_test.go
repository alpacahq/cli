//go:build integration

package integration

import (
	"testing"
	"time"
)

func TestWatchlist(t *testing.T) {
	name := "cli-integration-test"

	out := alpaca(t, "watchlist", "create", name,
		"--symbols", "AAPL,MSFT",
		"--json",
	)
	wl := parseJSONMap(t, out)
	wlID, ok := wl["id"].(string)
	if !ok || wlID == "" {
		t.Fatal("watchlist missing id")
	}
	t.Cleanup(func() {
		_ = makeCmd("watchlist", "delete", wlID).Run()
	})

	if wl["name"] != name {
		t.Errorf("expected name %q, got %v", name, wl["name"])
	}

	time.Sleep(300 * time.Millisecond)

	t.Run("get", func(t *testing.T) {
		out := alpaca(t, "watchlist", "get", wlID, "--json")
		fetched := parseJSONMap(t, out)
		if fetched["id"] != wlID {
			t.Errorf("get returned wrong watchlist: %v", fetched["id"])
		}
	})

	t.Run("list", func(t *testing.T) {
		out := alpaca(t, "watchlist", "list", "--json")
		lists := parseJSONArray(t, out)
		if !containsID(lists, wlID) {
			t.Error("created watchlist not found in list")
		}
	})

	t.Run("add_remove", func(t *testing.T) {
		alpaca(t, "watchlist", "add", wlID, "GOOG")
		time.Sleep(300 * time.Millisecond)

		alpaca(t, "watchlist", "remove", wlID, "MSFT")
		time.Sleep(300 * time.Millisecond)

		out := alpaca(t, "watchlist", "get", wlID, "--json")
		updated := parseJSONMap(t, out)
		assets, _ := updated["assets"].([]any)
		symbols := extractSymbols(assets)
		assertContains(t, symbols, "AAPL", true, "should still contain AAPL")
		assertContains(t, symbols, "GOOG", true, "should contain GOOG after add")
		assertContains(t, symbols, "MSFT", false, "should not contain MSFT after remove")
	})

	t.Run("by_name", func(t *testing.T) {
		out := alpaca(t, "watchlist", "get-by-name", name, "--json")
		fetched := parseJSONMap(t, out)
		if fetched["name"] != name {
			t.Errorf("get-by-name returned wrong name: %v", fetched["name"])
		}

		out = alpaca(t, "watchlist", "add-by-name", name, "TSLA", "--json")
		added := parseJSONMap(t, out)
		if added["name"] != name {
			t.Errorf("add-by-name returned wrong name: %v", added["name"])
		}

		time.Sleep(300 * time.Millisecond)

		out = alpaca(t, "watchlist", "update-by-name", name,
			"--symbols", "AAPL",
			"--json",
		)
		updated := parseJSONMap(t, out)
		if updated["name"] != name {
			t.Errorf("update-by-name returned wrong name: %v", updated["name"])
		}
	})
}

func extractSymbols(assets []any) map[string]bool {
	out := make(map[string]bool, len(assets))
	for _, a := range assets {
		if m, ok := a.(map[string]any); ok {
			if s, ok := m["symbol"].(string); ok {
				out[s] = true
			}
		}
	}
	return out
}

func assertContains(t *testing.T, symbols map[string]bool, sym string, want bool, msg string) {
	t.Helper()
	if symbols[sym] != want {
		t.Error(msg)
	}
}
