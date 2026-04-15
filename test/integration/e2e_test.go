//go:build integration

package integration

import (
	"slices"
	"strings"
	"testing"
	"time"
)

// TestE2E_DataPipeline fetches market data and pipes it through --jq and --csv,
// verifying the full data-to-output pipeline that agents rely on:
// JSON -> --jq transform -> --csv -> --jq + --csv combined.
func TestE2E_DataPipeline(t *testing.T) {
	t.Parallel()
	baseArgs := []string{"data", "bars", "--symbol", "AAPL",
		"--start", daysAgo(100),
		"--end", daysAgo(93),
		"--timeframe", "1Day",
	}
	withFlags := func(extra ...string) []string {
		return append(slices.Clone(baseArgs), extra...)
	}

	out := alpacaRetry(t, baseArgs...)
	data := parseJSONMap(t, out)
	bars, ok := data["bars"].([]any)
	if !ok || len(bars) == 0 {
		t.Fatal("expected non-empty bars array")
	}
	barCount := len(bars)
	first, _ := bars[0].(map[string]any)
	requireFields(t, first, "t", "o", "h", "l", "c", "v")

	t.Run("jq_transform", func(t *testing.T) {
		jqOut := alpacaRetry(t, withFlags("--jq", "[.bars[] | {t, c}]")...)
		jqBars := parseJSONArray(t, jqOut)
		if len(jqBars) != barCount {
			t.Errorf("--jq returned %d items, JSON had %d bars", len(jqBars), barCount)
		}
		requireFields(t, jqBars[0], "t", "c")
		for _, dropped := range []string{"o", "h", "l", "v"} {
			if _, has := jqBars[0][dropped]; has {
				t.Errorf("--jq filtered output should not have %q", dropped)
			}
		}
	})

	t.Run("csv_format", func(t *testing.T) {
		// --csv alone on a wrapped response (object with nested array) flattens
		// the wrapper, not the inner bars. Use --jq first to extract the array.
		csvOut := alpacaRetry(t, withFlags("--jq", ".bars", "--csv")...)
		csvLines := strings.Split(strings.TrimSpace(string(csvOut)), "\n")
		if len(csvLines) < 2 {
			t.Fatal("--jq .bars --csv should have header + data rows")
		}
		colSet := csvColumnSet(csvLines[0])
		for _, expected := range []string{"t", "o", "h", "l", "c", "v"} {
			if !colSet[expected] {
				t.Errorf("CSV header missing %q, got: %s", expected, csvLines[0])
			}
		}
		csvDataRows := len(csvLines) - 1
		if csvDataRows != barCount {
			t.Errorf("CSV has %d data rows but JSON has %d bars", csvDataRows, barCount)
		}
	})

	t.Run("jq_plus_csv", func(t *testing.T) {
		jqCsvOut := alpacaRetry(t, withFlags("--jq", "[.bars[] | {t, c}]", "--csv")...)
		jqCsvLines := strings.Split(strings.TrimSpace(string(jqCsvOut)), "\n")
		if len(jqCsvLines) < 2 {
			t.Fatal("--jq + --csv should have header + data rows")
		}
		colSet := csvColumnSet(jqCsvLines[0])
		if !colSet["t"] || !colSet["c"] {
			t.Errorf("header should have t and c, got: %s", jqCsvLines[0])
		}
		if colSet["o"] || colSet["h"] || colSet["l"] || colSet["v"] {
			t.Errorf("header should not have OHLV after jq filtering, got: %s", jqCsvLines[0])
		}
		jqCsvDataRows := len(jqCsvLines) - 1
		if jqCsvDataRows != barCount {
			t.Errorf("--jq + --csv has %d rows but JSON has %d bars", jqCsvDataRows, barCount)
		}
	})
}

// TestE2E_CryptoTradeLifecycle exercises the full trading loop:
// submit market buy -> wait for fill -> verify position -> check activity feed -> close -> verify gone.
func TestE2E_CryptoTradeLifecycle(t *testing.T) {
	t.Parallel()
	symbol := "UNI/USD"

	out := alpaca(t, "order", "submit",
		"--symbol", symbol,
		"--notional", "10",
		"--side", "buy",
		"--type", "market",
		"--time-in-force", "gtc",
	)
	order := parseJSONMap(t, out)
	orderID := order["id"].(string)
	requireFields(t, order, "id", "symbol", "side", "type", "status")

	t.Cleanup(func() {
		_ = makeCmd("order", "cancel", "--order-id", orderID).Run()
		_ = makeCmd("position", "close", "--symbol-or-asset-id", symbol).Run()
	})

	t.Run("fill", func(t *testing.T) {
		pollFor(t, 15*time.Second, "order to fill", func() bool {
			out := alpaca(t, "order", "get", "--order-id", orderID)
			return parseJSONMap(t, out)["status"] == "filled"
		})

		out := alpaca(t, "order", "get", "--order-id", orderID)
		filled := parseJSONMap(t, out)
		requireFields(t, filled, "filled_at", "filled_qty", "filled_avg_price")
	})

	t.Run("position", func(t *testing.T) {
		out := alpaca(t, "position", "get", "--symbol-or-asset-id", symbol)
		pos := parseJSONMap(t, out)
		requireFields(t, pos, "symbol", "qty", "market_value", "avg_entry_price", "unrealized_pl", "side")
		// API normalizes crypto symbols: UNI/USD -> UNIUSD in positions
		posSymbol, _ := pos["symbol"].(string)
		if posSymbol == "" {
			t.Error("position symbol should not be empty")
		}
	})

	t.Run("activity_feed", func(t *testing.T) {
		pollFor(t, 10*time.Second, "fill activity to appear", func() bool {
			out := alpaca(t, "account", "activity", "list", "--activity-types", "FILL", "--page-size", "10")
			for _, a := range parseJSONArray(t, out) {
				if a["symbol"] == symbol {
					requireFields(t, a, "id", "activity_type", "qty", "price", "side")
					return true
				}
			}
			return false
		})
	})

	t.Run("close_and_verify", func(t *testing.T) {
		out := alpaca(t, "position", "close", "--symbol-or-asset-id", symbol)
		closeOrder := parseJSONMap(t, out)
		requireFields(t, closeOrder, "id", "symbol", "side", "status")
		if closeOrder["side"] != "sell" {
			t.Errorf("close order side: want sell, got %v", closeOrder["side"])
		}

		pollFor(t, 10*time.Second, "position to disappear", func() bool {
			_, _, code := alpacaWithStderr(t, "position", "get", "--symbol-or-asset-id", symbol)
			return code != 0
		})
	})
}

// TestE2E_OrderReplaceChain submits a limit order and replaces it twice,
// verifying each replace produces a new order ID and the final order
// reflects all accumulated changes.
func TestE2E_OrderReplaceChain(t *testing.T) {
	t.Parallel()

	out := alpaca(t, "order", "submit",
		"--symbol", "SOL/USD",
		"--qty", "10",
		"--side", "buy",
		"--type", "limit",
		"--limit-price", "1.00",
		"--time-in-force", "gtc",
	)
	order := parseJSONMap(t, out)
	id1 := order["id"].(string)

	var id2, id3 string
	t.Cleanup(func() {
		for _, id := range []string{id1, id2, id3} {
			if id != "" {
				_ = makeCmd("order", "cancel", "--order-id", id).Run()
			}
		}
	})

	waitReplaceable := func(t *testing.T, id string) {
		t.Helper()
		pollFor(t, 10*time.Second, "order to be replaceable", func() bool {
			out := alpaca(t, "order", "get", "--order-id", id)
			s, _ := parseJSONMap(t, out)["status"].(string)
			return s != "accepted" && s != "pending_new" && s != "pending_replace"
		})
	}

	waitReplaceable(t, id1)

	t.Run("replace_price", func(t *testing.T) {
		out := alpaca(t, "order", "replace", "--order-id", id1,
			"--limit-price", "1.50",
		)
		r := parseJSONMap(t, out)
		id2 = r["id"].(string)
		if id2 == id1 {
			t.Fatal("replace should produce a new order ID")
		}

		pollFor(t, 5*time.Second, "price replacement to settle", func() bool {
			out := alpaca(t, "order", "get", "--order-id", id2)
			return parseJSONMap(t, out)["limit_price"] == "1.5"
		})
	})

	t.Run("replace_qty", func(t *testing.T) {
		waitReplaceable(t, id2)
		out := alpaca(t, "order", "replace", "--order-id", id2,
			"--qty", "15",
		)
		r := parseJSONMap(t, out)
		id3 = r["id"].(string)
		if id3 == id2 || id3 == id1 {
			t.Fatal("second replace should produce a unique order ID")
		}

		pollFor(t, 5*time.Second, "qty replacement to settle", func() bool {
			out := alpaca(t, "order", "get", "--order-id", id3)
			return parseJSONMap(t, out)["qty"] == "15"
		})
	})

	t.Run("verify_final_state", func(t *testing.T) {
		out := alpaca(t, "order", "get", "--order-id", id3)
		final := parseJSONMap(t, out)
		if final["limit_price"] != "1.5" {
			t.Errorf("final price: want 1.5, got %v", final["limit_price"])
		}
		if final["qty"] != "15" {
			t.Errorf("final qty: want 15, got %v", final["qty"])
		}

		out = alpaca(t, "order", "get", "--order-id", id1)
		orig := parseJSONMap(t, out)
		if orig["status"] != "replaced" {
			t.Errorf("original order status: want replaced, got %v", orig["status"])
		}
	})
}

func csvColumnSet(header string) map[string]bool {
	cols := strings.Split(header, ",")
	set := make(map[string]bool, len(cols))
	for _, c := range cols {
		set[c] = true
	}
	return set
}
