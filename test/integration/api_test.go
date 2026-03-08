//go:build integration

package integration

import (
	"testing"
)

func TestAPI_Get(t *testing.T) {
	t.Parallel()
	out := alpaca(t, "api", "get", "/v2/clock")
	clock := parseJSONMap(t, out)
	requireFields(t, clock, "is_open")
}

func TestAPI_Post(t *testing.T) {
	out := alpaca(t, "api", "post", "/v2/watchlists",
		"--data", `{"name":"api-test-post","symbols":[]}`,
	)
	wl := parseJSONMap(t, out)
	id, _ := wl["id"].(string)
	if id == "" {
		t.Fatal("api post did not return watchlist with id")
	}
	t.Cleanup(func() {
		_ = makeCmd("api", "delete", "/v2/watchlists/"+id).Run()
	})

	if wl["name"] != "api-test-post" {
		t.Errorf("expected name api-test-post, got %v", wl["name"])
	}
}

func TestAPI_Patch(t *testing.T) {
	// Use account config for PATCH — watchlists use PUT
	out := alpaca(t, "api", "get", "/v2/account/configurations")
	original := parseJSONMap(t, out)
	origVal, _ := original["trade_confirm_email"].(string)

	newVal := "none"
	if origVal == "none" {
		newVal = "all"
	}

	out = alpaca(t, "api", "patch", "/v2/account/configurations",
		"--data", `{"trade_confirm_email":"`+newVal+`"}`,
	)
	updated := parseJSONMap(t, out)
	if updated["trade_confirm_email"] != newVal {
		t.Errorf("expected trade_confirm_email %q, got %v", newVal, updated["trade_confirm_email"])
	}

	// Restore
	t.Cleanup(func() {
		_ = makeCmd("api", "patch", "/v2/account/configurations",
			"--data", `{"trade_confirm_email":"`+origVal+`"}`).Run()
	})
}

func TestAPI_Delete(t *testing.T) {
	out := alpaca(t, "api", "post", "/v2/watchlists",
		"--data", `{"name":"api-test-delete","symbols":[]}`,
	)
	wl := parseJSONMap(t, out)
	id := wl["id"].(string)

	alpaca(t, "api", "delete", "/v2/watchlists/"+id)

	// Verify it's gone
	_, _, code := alpacaWithStderr(t, "api", "get", "/v2/watchlists/"+id)
	if code == 0 {
		t.Error("watchlist should be deleted")
	}
}

func TestAPI_UseDataAPI(t *testing.T) {
	t.Parallel()
	out := alpaca(t, "api", "get", "/v2/stocks/AAPL/trades/latest", "--use-data-api")
	trade := parseJSONMap(t, out)
	requireFields(t, trade, "trade")
}
