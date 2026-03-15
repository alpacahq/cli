package output

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func TestApplyJQ_Identity(t *testing.T) {
	data := json.RawMessage(`{"symbol":"AAPL","qty":"10"}`)
	got, err := ApplyJQ(data, ".")
	if err != nil {
		t.Fatal(err)
	}
	m, ok := got.(map[string]any)
	if !ok {
		t.Fatalf("expected map, got %T", got)
	}
	if m["symbol"] != "AAPL" {
		t.Errorf("expected AAPL, got %v", m["symbol"])
	}
}

func TestApplyJQ_FieldSelection(t *testing.T) {
	data := json.RawMessage(`{"symbol":"AAPL","qty":"10","side":"buy","type":"market"}`)
	got, err := ApplyJQ(data, "{symbol, qty}")
	if err != nil {
		t.Fatal(err)
	}
	m, ok := got.(map[string]any)
	if !ok {
		t.Fatalf("expected map, got %T", got)
	}
	if len(m) != 2 {
		t.Errorf("expected 2 keys, got %d: %v", len(m), m)
	}
	if m["symbol"] != "AAPL" || m["qty"] != "10" {
		t.Errorf("unexpected values: %v", m)
	}
}

func TestApplyJQ_ArrayFieldSelection(t *testing.T) {
	data := json.RawMessage(`[
		{"symbol":"AAPL","qty":"10","side":"buy"},
		{"symbol":"MSFT","qty":"5","side":"sell"}
	]`)
	got, err := ApplyJQ(data, "[.[] | {symbol, qty}]")
	if err != nil {
		t.Fatal(err)
	}
	arr, ok := got.([]any)
	if !ok {
		t.Fatalf("expected []any, got %T", got)
	}
	if len(arr) != 2 {
		t.Errorf("expected 2 items, got %d", len(arr))
	}
	first := arr[0].(map[string]any)
	if first["symbol"] != "AAPL" {
		t.Errorf("expected AAPL, got %v", first["symbol"])
	}
	if _, hasSide := first["side"]; hasSide {
		t.Error("side should have been filtered out")
	}
}

func TestApplyJQ_ScalarResult(t *testing.T) {
	data := json.RawMessage(`{"symbol":"AAPL","qty":"10"}`)
	got, err := ApplyJQ(data, ".symbol")
	if err != nil {
		t.Fatal(err)
	}
	if got != "AAPL" {
		t.Errorf("expected AAPL, got %v", got)
	}
}

func TestApplyJQ_MultipleResults(t *testing.T) {
	data := json.RawMessage(`[{"s":"A"},{"s":"B"},{"s":"C"}]`)
	got, err := ApplyJQ(data, ".[] | .s")
	if err != nil {
		t.Fatal(err)
	}
	arr, ok := got.([]any)
	if !ok {
		t.Fatalf("expected []any, got %T", got)
	}
	if len(arr) != 3 {
		t.Errorf("expected 3 results, got %d", len(arr))
	}
}

func TestApplyJQ_InvalidExpression(t *testing.T) {
	data := json.RawMessage(`{}`)
	_, err := ApplyJQ(data, ".[invalid")
	if err == nil {
		t.Fatal("expected error for invalid jq expression")
	}
	if !strings.Contains(err.Error(), "--jq") {
		t.Errorf("error should mention --jq: %v", err)
	}
}

func TestApplyJQ_GoStruct(t *testing.T) {
	type pos struct {
		Symbol string `json:"symbol"`
		Qty    int    `json:"qty"`
	}
	data := pos{Symbol: "AAPL", Qty: 10}
	got, err := ApplyJQ(data, ".symbol")
	if err != nil {
		t.Fatal(err)
	}
	if got != "AAPL" {
		t.Errorf("expected AAPL, got %v", got)
	}
}

func TestApplyJQ_WithCSV(t *testing.T) {
	data := json.RawMessage(`[
		{"symbol":"AAPL","qty":"10","side":"buy","market_value":"1850"},
		{"symbol":"MSFT","qty":"5","side":"sell","market_value":"2100"}
	]`)

	filtered, err := ApplyJQ(data, "[.[] | {symbol, qty}]")
	if err != nil {
		t.Fatal(err)
	}

	var buf bytes.Buffer
	if err := Render(&buf, FormatCSV, filtered); err != nil {
		t.Fatal(err)
	}

	out := buf.String()
	lines := strings.Split(strings.TrimSpace(out), "\n")
	if len(lines) != 3 {
		t.Fatalf("expected 3 lines (header + 2 rows), got %d: %s", len(lines), out)
	}
	if lines[0] != "qty,symbol" {
		t.Errorf("expected header 'qty,symbol', got: %s", lines[0])
	}
	if !strings.Contains(lines[1], "AAPL") {
		t.Errorf("expected AAPL in first row, got: %s", lines[1])
	}
}

func TestApplyJQ_NilResult(t *testing.T) {
	data := json.RawMessage(`{"a":1}`)
	got, err := ApplyJQ(data, ".b")
	if err != nil {
		t.Fatal(err)
	}
	if got != nil {
		t.Errorf("expected nil, got %v", got)
	}
}
