//go:build integration

package integration

import (
	"strings"
	"testing"
	"time"
)

func TestOrderReplace(t *testing.T) {
	// Use crypto limit order — crypto is 24/7, so the order transitions
	// out of "accepted" immediately unlike equity orders outside market hours.
	// Crypto min cost basis is $10; qty 10 * $1.00 = $10
	out := alpaca(t, "order", "submit", "BTC/USD",
		"--qty", "10",
		"--side", "buy",
		"--type", "limit",
		"--limit-price", "1.00",
		"--time-in-force", "gtc",
		"--json",
	)
	order := parseJSONMap(t, out)
	id := order["id"].(string)
	t.Cleanup(func() {
		_ = makeCmd("order", "cancel", id).Run()
	})

	// Poll until order leaves "accepted" status (up to 10s)
	for i := 0; i < 20; i++ {
		time.Sleep(500 * time.Millisecond)
		out = alpaca(t, "order", "get", id, "--json")
		o := parseJSONMap(t, out)
		if o["status"] != "accepted" && o["status"] != "pending_new" {
			break
		}
	}

	out = alpaca(t, "order", "replace", id,
		"--qty", "11",
		"--limit-price", "1.50",
		"--json",
	)
	replaced := parseJSONMap(t, out)

	newID, _ := replaced["id"].(string)
	if newID == "" {
		t.Fatal("replace returned no order id")
	}

	time.Sleep(1 * time.Second)

	out = alpaca(t, "order", "get", newID, "--json")
	fetched := parseJSONMap(t, out)
	if fetched["qty"] != "11" {
		t.Errorf("expected qty 11, got %v", fetched["qty"])
	}
}

func TestOrderSubmit_LimitOrder(t *testing.T) {
	out := alpaca(t, "order", "submit", "AAPL",
		"--qty", "1",
		"--side", "buy",
		"--type", "limit",
		"--limit-price", "1.00",
		"--time-in-force", "gtc",
		"--json",
	)
	order := parseJSONMap(t, out)
	id := order["id"].(string)
	t.Cleanup(func() {
		alpacaWithStderr(t, "order", "cancel", id)
	})

	if order["type"] != "limit" {
		t.Errorf("expected type limit, got %v", order["type"])
	}
	requireFields(t, order, "limit_price")
}

func TestOrderSubmit_StopOrder(t *testing.T) {
	out := alpaca(t, "order", "submit", "AAPL",
		"--qty", "1",
		"--side", "sell",
		"--type", "stop",
		"--stop-price", "1.00",
		"--time-in-force", "gtc",
		"--json",
	)
	order := parseJSONMap(t, out)
	id := order["id"].(string)
	t.Cleanup(func() {
		alpacaWithStderr(t, "order", "cancel", id)
	})

	if order["type"] != "stop" {
		t.Errorf("expected type stop, got %v", order["type"])
	}
	requireFields(t, order, "stop_price")
}

func TestOrderSubmit_StopLimitOrder(t *testing.T) {
	out := alpaca(t, "order", "submit", "AAPL",
		"--qty", "1",
		"--side", "sell",
		"--type", "stop_limit",
		"--stop-price", "1.00",
		"--limit-price", "0.90",
		"--time-in-force", "gtc",
		"--json",
	)
	order := parseJSONMap(t, out)
	id := order["id"].(string)
	t.Cleanup(func() {
		alpacaWithStderr(t, "order", "cancel", id)
	})

	if order["type"] != "stop_limit" {
		t.Errorf("expected type stop_limit, got %v", order["type"])
	}
	requireFields(t, order, "stop_price", "limit_price")
}

func TestOrderSubmit_TrailingStopOrder(t *testing.T) {
	out := alpaca(t, "order", "submit", "AAPL",
		"--qty", "1",
		"--side", "sell",
		"--type", "trailing_stop",
		"--trail-percent", "5",
		"--time-in-force", "gtc",
		"--json",
	)
	order := parseJSONMap(t, out)
	id := order["id"].(string)
	t.Cleanup(func() {
		alpacaWithStderr(t, "order", "cancel", id)
	})

	if order["type"] != "trailing_stop" {
		t.Errorf("expected type trailing_stop, got %v", order["type"])
	}
	requireFields(t, order, "trail_percent")
}

func TestOrderSubmit_Notional(t *testing.T) {
	out := alpaca(t, "order", "submit", "BTC/USD",
		"--notional", "10",
		"--side", "buy",
		"--type", "market",
		"--time-in-force", "gtc",
		"--json",
	)
	order := parseJSONMap(t, out)
	requireFields(t, order, "id", "notional")

	t.Cleanup(func() {
		time.Sleep(2 * time.Second)
		_ = makeCmd("position", "close", "BTC/USD").Run()
	})
}

func TestOrderSubmit_ClientOrderID(t *testing.T) {
	clientID := "integ-test-" + time.Now().Format("20060102150405")
	out := alpaca(t, "order", "submit", "AAPL",
		"--qty", "1",
		"--side", "buy",
		"--type", "limit",
		"--limit-price", "1.00",
		"--time-in-force", "gtc",
		"--client-order-id", clientID,
		"--json",
	)
	order := parseJSONMap(t, out)
	id := order["id"].(string)
	t.Cleanup(func() {
		alpacaWithStderr(t, "order", "cancel", id)
	})

	if order["client_order_id"] != clientID {
		t.Errorf("expected client_order_id %q, got %v", clientID, order["client_order_id"])
	}

	time.Sleep(500 * time.Millisecond)

	out = alpaca(t, "order", "get", "--client-id", clientID, "--json")
	fetched := parseJSONMap(t, out)
	if fetched["id"] != id {
		t.Errorf("get by client-id returned wrong order: got %v, want %s", fetched["id"], id)
	}
}

func TestOrderSubmit_DryRun(t *testing.T) {
	out := alpaca(t, "order", "submit", "AAPL",
		"--qty", "1",
		"--side", "buy",
		"--type", "market",
		"--dry-run",
		"--json",
	)
	body := parseJSONMap(t, out)
	if body["symbol"] != "AAPL" {
		t.Errorf("dry-run body should have symbol AAPL, got %v", body["symbol"])
	}
	if body["side"] != "buy" {
		t.Errorf("dry-run body should have side buy, got %v", body["side"])
	}
}

func TestOrderList_StatusFilter(t *testing.T) {
	out := alpaca(t, "order", "list", "--status", "all", "--limit", "5", "--json")
	_ = parseJSONArray(t, out)

	out = alpaca(t, "order", "list", "--status", "closed", "--limit", "5", "--json")
	_ = parseJSONArray(t, out)
}

func TestOrderList_SymbolFilter(t *testing.T) {
	_ = submitTestOrder(t)

	out := alpaca(t, "order", "list", "--status", "open", "--symbols", "AAPL", "--json")
	orders := parseJSONArray(t, out)
	for _, o := range orders {
		if o["symbol"] != "AAPL" {
			t.Errorf("symbol filter returned non-AAPL order: %v", o["symbol"])
		}
	}
}

func TestOrderList_Limit(t *testing.T) {
	// Ensure at least 2 orders exist
	_ = submitTestOrder(t)
	_ = submitTestOrder(t)

	out := alpaca(t, "order", "list", "--status", "open", "--limit", "1", "--json")
	orders := parseJSONArray(t, out)
	if len(orders) > 1 {
		t.Errorf("--limit 1 returned %d orders", len(orders))
	}
}

func TestOrderList_Direction(t *testing.T) {
	out := alpaca(t, "order", "list", "--status", "all", "--limit", "5", "--direction", "asc", "--json")
	orders := parseJSONArray(t, out)

	if len(orders) >= 2 {
		first, _ := orders[0]["created_at"].(string)
		last, _ := orders[len(orders)-1]["created_at"].(string)
		if strings.Compare(first, last) > 0 {
			t.Error("--direction asc: first order should have earlier created_at than last")
		}
	}
}
