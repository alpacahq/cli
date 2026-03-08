//go:build integration

package integration

import (
	"testing"
	"time"
)

func TestOrderLifecycle(t *testing.T) {
	out := alpaca(t, "order", "submit", "AAPL",
		"--qty", "1",
		"--side", "buy",
		"--type", "limit",
		"--limit-price", "1.00",
		"--time-in-force", "gtc",
		"--json",
	)

	order := parseJSONMap(t, out)
	orderID, ok := order["id"].(string)
	if !ok || orderID == "" {
		t.Fatal("order missing id")
	}
	t.Cleanup(func() {
		_ = makeCmd("order", "cancel", orderID).Run()
	})

	t.Run("submit_fields", func(t *testing.T) {
		if order["symbol"] != "AAPL" {
			t.Errorf("expected symbol AAPL, got %v", order["symbol"])
		}
		if order["side"] != "buy" {
			t.Errorf("expected side buy, got %v", order["side"])
		}
	})

	time.Sleep(500 * time.Millisecond)

	t.Run("get", func(t *testing.T) {
		out := alpaca(t, "order", "get", orderID, "--json")
		fetched := parseJSONMap(t, out)
		if fetched["id"] != orderID {
			t.Errorf("get returned wrong order: %v", fetched["id"])
		}
	})

	t.Run("list_open", func(t *testing.T) {
		out := alpaca(t, "order", "list", "--status", "open", "--json")
		orders := parseJSONArray(t, out)
		if !containsID(orders, orderID) {
			t.Error("open orders list does not contain our order")
		}
	})

	t.Run("cancel", func(t *testing.T) {
		alpaca(t, "order", "cancel", orderID)
		time.Sleep(500 * time.Millisecond)

		out := alpaca(t, "order", "get", orderID, "--json")
		cancelled := parseJSONMap(t, out)
		status, _ := cancelled["status"].(string)
		if status != "canceled" && status != "cancelled" && status != "pending_cancel" {
			t.Errorf("expected order to be canceled, got %v", status)
		}
	})
}

func TestOrderCancelAll(t *testing.T) {
	// Submit two orders
	for range 2 {
		alpaca(t, "order", "submit", "AAPL",
			"--qty", "1",
			"--side", "buy",
			"--type", "limit",
			"--limit-price", "1.00",
			"--time-in-force", "gtc",
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
