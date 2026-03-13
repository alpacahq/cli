package cmd

import (
	"encoding/json"
	"testing"
)

func TestExtractArray_Bars_SingleSymbol(t *testing.T) {
	raw := json.RawMessage(`{"bars":[{"t":"2025-01-02","o":180,"h":185,"l":179,"c":184,"v":1000}]}`)
	result := extractArray(raw, "AAPL", "bars")
	var arr []map[string]any
	if err := json.Unmarshal(result, &arr); err != nil {
		t.Fatalf("expected array, got: %s", string(result))
	}
	if len(arr) != 1 {
		t.Errorf("expected 1 bar, got %d", len(arr))
	}
}

func TestExtractArray_Bars_MultiSymbolMap(t *testing.T) {
	raw := json.RawMessage(`{"bars":{"AAPL":[{"t":"2025-01-02","o":180}],"MSFT":[{"t":"2025-01-02","o":400}]}}`)
	result := extractArray(raw, "AAPL", "bars")
	var arr []map[string]any
	if err := json.Unmarshal(result, &arr); err != nil {
		t.Fatalf("expected array, got: %s", string(result))
	}
	if len(arr) != 1 {
		t.Errorf("expected 1 bar for AAPL, got %d", len(arr))
	}
	if arr[0]["o"].(float64) != 180 {
		t.Errorf("expected open=180, got %v", arr[0]["o"])
	}
}

func TestExtractArray_Bars_InvalidJSON(t *testing.T) {
	raw := json.RawMessage(`not json`)
	result := extractArray(raw, "AAPL", "bars")
	if string(result) != "not json" {
		t.Errorf("expected passthrough, got: %s", string(result))
	}
}

func TestExtractArray_Quotes(t *testing.T) {
	raw := json.RawMessage(`{"quotes":[{"ap":150,"bp":149.9}]}`)
	result := extractArray(raw, "AAPL", "quotes")
	var arr []map[string]any
	if err := json.Unmarshal(result, &arr); err != nil {
		t.Fatalf("expected array, got: %s", string(result))
	}
	if len(arr) != 1 {
		t.Errorf("expected 1 quote, got %d", len(arr))
	}
}

func TestExtractArray_MultiSymbol(t *testing.T) {
	raw := json.RawMessage(`{"trades":{"BTC/USD":[{"p":50000}],"ETH/USD":[{"p":3000}]}}`)
	result := extractArray(raw, "BTC/USD", "trades")
	var arr []map[string]any
	if err := json.Unmarshal(result, &arr); err != nil {
		t.Fatalf("expected array, got: %s", string(result))
	}
	if arr[0]["p"].(float64) != 50000 {
		t.Errorf("expected price=50000, got %v", arr[0]["p"])
	}
}

func TestExtractSingle_StockTrade(t *testing.T) {
	m := map[string]any{"trade": map[string]any{"p": 150.5, "s": 100.0}}
	trade := extractSingle(m, "AAPL", "trade", "trades")
	if trade["p"].(float64) != 150.5 {
		t.Errorf("expected price=150.5, got %v", trade["p"])
	}
}

func TestExtractSingle_CryptoTrade(t *testing.T) {
	m := map[string]any{
		"trades": map[string]any{
			"BTC/USD": map[string]any{"p": 50000.0, "s": 0.5},
		},
	}
	trade := extractSingle(m, "BTC/USD", "trade", "trades")
	if trade["p"].(float64) != 50000.0 {
		t.Errorf("expected price=50000, got %v", trade["p"])
	}
}

func TestExtractSingle_StockQuote(t *testing.T) {
	m := map[string]any{"quote": map[string]any{"bp": 149.9, "ap": 150.1}}
	quote := extractSingle(m, "AAPL", "quote", "quotes")
	if quote["bp"].(float64) != 149.9 {
		t.Errorf("expected bid=149.9, got %v", quote["bp"])
	}
}

func TestExtractSingle_CryptoQuote(t *testing.T) {
	m := map[string]any{
		"quotes": map[string]any{
			"ETH/USD": map[string]any{"bp": 3000.0, "ap": 3001.0},
		},
	}
	quote := extractSingle(m, "ETH/USD", "quote", "quotes")
	if quote["ap"].(float64) != 3001.0 {
		t.Errorf("expected ask=3001, got %v", quote["ap"])
	}
}
