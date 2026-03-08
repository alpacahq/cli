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

	t.Run("get", func(t *testing.T) {
		out := alpaca(t, "order", "get", orderID, "--json")
		fetched := parseJSONMap(t, out)
		if fetched["id"] != orderID {
			t.Errorf("get returned wrong order: %v", fetched["id"])
		}
	})

	t.Run("list_open", func(t *testing.T) {
		pollFor(t, 5*time.Second, "order to appear in open list", func() bool {
			out := alpaca(t, "order", "list", "--status", "open", "--json")
			return containsID(parseJSONArray(t, out), orderID)
		})
	})

	t.Run("cancel", func(t *testing.T) {
		alpaca(t, "order", "cancel", orderID)
		pollFor(t, 5*time.Second, "order to be canceled", func() bool {
			out := alpaca(t, "order", "get", orderID, "--json")
			status, _ := parseJSONMap(t, out)["status"].(string)
			return status == "canceled" || status == "cancelled" || status == "pending_cancel"
		})
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

	pollFor(t, 5*time.Second, "orders to appear", func() bool {
		out := alpaca(t, "order", "list", "--status", "open", "--json")
		return len(parseJSONArray(t, out)) >= 2
	})

	alpaca(t, "order", "cancel-all")

	pollFor(t, 10*time.Second, "all orders to be canceled", func() bool {
		out := alpaca(t, "order", "list", "--status", "open", "--json")
		return len(parseJSONArray(t, out)) == 0
	})
}
