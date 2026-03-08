//go:build integration

package integration

import (
	"testing"
	"time"
)

func TestWatchlistLifecycle(t *testing.T) {
	out := alpaca(t, "watchlist", "create",
		"cli-integration-test",
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

	t.Run("create_fields", func(t *testing.T) {
		if wl["name"] != "cli-integration-test" {
			t.Errorf("expected name 'cli-integration-test', got %v", wl["name"])
		}
	})

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

	t.Run("add_remove_symbols", func(t *testing.T) {
		alpaca(t, "watchlist", "add", wlID, "GOOG")
		time.Sleep(300 * time.Millisecond)

		alpaca(t, "watchlist", "remove", wlID, "MSFT")
		time.Sleep(300 * time.Millisecond)

		out := alpaca(t, "watchlist", "get", wlID, "--json")
		updated := parseJSONMap(t, out)
		assets, _ := updated["assets"].([]any)
		symbols := make([]string, 0, len(assets))
		for _, a := range assets {
			if m, ok := a.(map[string]any); ok {
				if s, ok := m["symbol"].(string); ok {
					symbols = append(symbols, s)
				}
			}
		}
		hasAAPL, hasGOOG, hasMSFT := false, false, false
		for _, s := range symbols {
			switch s {
			case "AAPL":
				hasAAPL = true
			case "GOOG":
				hasGOOG = true
			case "MSFT":
				hasMSFT = true
			}
		}
		if !hasAAPL {
			t.Error("watchlist should still contain AAPL")
		}
		if !hasGOOG {
			t.Error("watchlist should contain GOOG after add")
		}
		if hasMSFT {
			t.Error("watchlist should not contain MSFT after remove")
		}
	})
}

func TestWatchlistByNameOps(t *testing.T) {
	name := "cli-byname-test"

	out := alpaca(t, "watchlist", "create", name,
		"--symbols", "AAPL",
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

	time.Sleep(300 * time.Millisecond)

	t.Run("get_by_name", func(t *testing.T) {
		out := alpaca(t, "watchlist", "get-by-name", name, "--json")
		fetched := parseJSONMap(t, out)
		if fetched["name"] != name {
			t.Errorf("get-by-name returned wrong name: %v", fetched["name"])
		}
	})

	t.Run("add_by_name", func(t *testing.T) {
		out := alpaca(t, "watchlist", "add-by-name", name, "MSFT", "--json")
		added := parseJSONMap(t, out)
		if added["name"] != name {
			t.Errorf("add-by-name returned wrong name: %v", added["name"])
		}
	})

	time.Sleep(300 * time.Millisecond)

	t.Run("update_by_name", func(t *testing.T) {
		out := alpaca(t, "watchlist", "update-by-name", name,
			"--symbols", "GOOG,TSLA",
			"--json",
		)
		updated := parseJSONMap(t, out)
		if updated["name"] != name {
			t.Errorf("update-by-name returned wrong name: %v", updated["name"])
		}
	})

	time.Sleep(300 * time.Millisecond)

	t.Run("delete_by_name", func(t *testing.T) {
		alpaca(t, "watchlist", "delete-by-name", name)
	})
}
