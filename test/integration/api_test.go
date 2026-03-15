//go:build integration

package integration

import (
	"testing"
)

func TestAPI_Get(t *testing.T) {
	t.Parallel()
	out := alpaca(t, "api", "GET", "/v2/clock")
	clock := parseJSONMap(t, out)
	requireFields(t, clock, "is_open")
}

func TestAPI_Post(t *testing.T) {
	out := alpaca(t, "api", "POST", "/v2/watchlists",
		"--body", `{"name":"api-test-post","symbols":[]}`,
	)
	wl := parseJSONMap(t, out)
	id, _ := wl["id"].(string)
	if id == "" {
		t.Fatal("api POST did not return watchlist with id")
	}
	t.Cleanup(func() {
		_ = makeCmd("api", "DELETE", "/v2/watchlists/"+id).Run()
	})

	if wl["name"] != "api-test-post" {
		t.Errorf("expected name api-test-post, got %v", wl["name"])
	}
}

func TestAPI_Patch(t *testing.T) {
	out := alpaca(t, "api", "GET", "/v2/account/configurations")
	original := parseJSONMap(t, out)
	origVal, _ := original["trade_confirm_email"].(string)

	newVal := "none"
	if origVal == "none" {
		newVal = "all"
	}

	out = alpaca(t, "api", "PATCH", "/v2/account/configurations",
		"--body", `{"trade_confirm_email":"`+newVal+`"}`,
	)
	updated := parseJSONMap(t, out)
	if updated["trade_confirm_email"] != newVal {
		t.Errorf("expected trade_confirm_email %q, got %v", newVal, updated["trade_confirm_email"])
	}

	t.Cleanup(func() {
		_ = makeCmd("api", "PATCH", "/v2/account/configurations",
			"--body", `{"trade_confirm_email":"`+origVal+`"}`).Run()
	})
}

func TestAPI_Delete(t *testing.T) {
	out := alpaca(t, "api", "POST", "/v2/watchlists",
		"--body", `{"name":"api-test-delete","symbols":[]}`,
	)
	wl := parseJSONMap(t, out)
	id := wl["id"].(string)

	alpaca(t, "api", "DELETE", "/v2/watchlists/"+id)

	_, _, code := alpacaWithStderr(t, "api", "GET", "/v2/watchlists/"+id)
	if code == 0 {
		t.Error("watchlist should be deleted")
	}
}

func TestAPI_UseDataAPI(t *testing.T) {
	t.Parallel()
	out := alpaca(t, "api", "GET", "/v2/stocks/AAPL/trades/latest", "--use-data-api")
	trade := parseJSONMap(t, out)
	requireFields(t, trade, "trade")
}

func TestAPI_QueryFlag(t *testing.T) {
	t.Parallel()
	out := alpaca(t, "api", "GET", "/v2/orders", "--query", "status=all&limit=1")
	orders := parseJSONArray(t, out)
	if len(orders) > 1 {
		t.Errorf("--query limit=1 returned %d orders", len(orders))
	}
}
