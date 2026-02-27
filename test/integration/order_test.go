//go:build integration

package integration

import (
	"testing"
	"time"
)

func TestOrderLifecycle(t *testing.T) {
	// Submit a limit order at $1.00 — well below market, won't fill
	out := alpaca(t, "order", "submit",
		"--symbol", "AAPL",
		"--qty", "1",
		"--side", "buy",
		"--type", "limit",
		"--limit-price", "1.00",
		"--tif", "gtc",
		"--json",
	)

	order := parseJSONMap(t, out)
	orderID, ok := order["id"].(string)
	if !ok || orderID == "" {
		t.Fatal("order missing id")
	}
	t.Cleanup(func() {
		alpaca(t, "order", "cancel", orderID)
	})

	if order["symbol"] != "AAPL" {
		t.Errorf("expected symbol AAPL, got %v", order["symbol"])
	}
	if order["side"] != "buy" {
		t.Errorf("expected side buy, got %v", order["side"])
	}

	// Allow order to propagate
	time.Sleep(500 * time.Millisecond)

	// Get order by ID
	out = alpaca(t, "order", "get", orderID, "--json")
	fetched := parseJSONMap(t, out)
	if fetched["id"] != orderID {
		t.Errorf("get returned wrong order: %v", fetched["id"])
	}

	// List open orders — should contain our order
	out = alpaca(t, "order", "list", "--status", "open", "--json")
	orders := parseJSONArray(t, out)
	if !containsID(orders, orderID) {
		t.Error("open orders list does not contain our order")
	}

	// Cancel the order
	alpaca(t, "order", "cancel", orderID)
	time.Sleep(500 * time.Millisecond)

	// Verify cancelled
	out = alpaca(t, "order", "get", orderID, "--json")
	cancelled := parseJSONMap(t, out)
	status, _ := cancelled["status"].(string)
	if status != "canceled" && status != "cancelled" && status != "pending_cancel" {
		t.Errorf("expected order to be canceled, got %v", status)
	}
}

func TestBuyShortcut(t *testing.T) {
	out := alpaca(t, "buy", "AAPL", "1",
		"--limit", "1.00",
		"--tif", "gtc",
		"--json",
	)

	order := parseJSONMap(t, out)
	orderID, _ := order["id"].(string)
	t.Cleanup(func() {
		if orderID != "" {
			alpaca(t, "order", "cancel", orderID)
		}
	})

	if order["side"] != "buy" {
		t.Errorf("expected buy, got %v", order["side"])
	}
	if order["type"] != "limit" {
		t.Errorf("expected limit order, got %v", order["type"])
	}
}

func TestOrderCancelAll(t *testing.T) {
	// Submit two orders
	for range 2 {
		alpaca(t, "order", "submit",
			"--symbol", "AAPL",
			"--qty", "1",
			"--side", "buy",
			"--type", "limit",
			"--limit-price", "1.00",
			"--tif", "gtc",
			"--json",
		)
	}

	time.Sleep(500 * time.Millisecond)

	// Cancel all
	alpaca(t, "order", "cancel-all")

	time.Sleep(1 * time.Second)

	// Verify no open orders remain
	out := alpaca(t, "order", "list", "--status", "open", "--json")
	orders := parseJSONArray(t, out)
	if len(orders) > 0 {
		t.Errorf("expected 0 open orders after cancel-all, got %d", len(orders))
	}
}
