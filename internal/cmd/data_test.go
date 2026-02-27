package cmd

import (
	"encoding/json"
	"testing"
)

func TestExtractBars_SingleSymbol(t *testing.T) {
	raw := json.RawMessage(`{"bars":[{"t":"2025-01-02","o":180,"h":185,"l":179,"c":184,"v":1000}]}`)
	result := extractBars(raw, "AAPL")
	var arr []map[string]any
	if err := json.Unmarshal(result, &arr); err != nil {
		t.Fatalf("expected array, got: %s", string(result))
	}
	if len(arr) != 1 {
		t.Errorf("expected 1 bar, got %d", len(arr))
	}
}

func TestExtractBars_MultiSymbolMap(t *testing.T) {
	raw := json.RawMessage(`{"bars":{"AAPL":[{"t":"2025-01-02","o":180}],"MSFT":[{"t":"2025-01-02","o":400}]}}`)
	result := extractBars(raw, "AAPL")
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

func TestExtractBars_InvalidJSON(t *testing.T) {
	raw := json.RawMessage(`not json`)
	result := extractBars(raw, "AAPL")
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

func TestExtractTrade_StockFormat(t *testing.T) {
	m := map[string]any{"trade": map[string]any{"p": 150.5, "s": 100.0}}
	trade := extractTrade(m, "AAPL")
	if trade["p"].(float64) != 150.5 {
		t.Errorf("expected price=150.5, got %v", trade["p"])
	}
}

func TestExtractTrade_CryptoFormat(t *testing.T) {
	m := map[string]any{
		"trades": map[string]any{
			"BTC/USD": map[string]any{"p": 50000.0, "s": 0.5},
		},
	}
	trade := extractTrade(m, "BTC/USD")
	if trade["p"].(float64) != 50000.0 {
		t.Errorf("expected price=50000, got %v", trade["p"])
	}
}

func TestExtractQuote_StockFormat(t *testing.T) {
	m := map[string]any{"quote": map[string]any{"bp": 149.9, "ap": 150.1}}
	quote := extractQuote(m, "AAPL")
	if quote["bp"].(float64) != 149.9 {
		t.Errorf("expected bid=149.9, got %v", quote["bp"])
	}
}

func TestExtractQuote_CryptoFormat(t *testing.T) {
	m := map[string]any{
		"quotes": map[string]any{
			"ETH/USD": map[string]any{"bp": 3000.0, "ap": 3001.0},
		},
	}
	quote := extractQuote(m, "ETH/USD")
	if quote["ap"].(float64) != 3001.0 {
		t.Errorf("expected ask=3001, got %v", quote["ap"])
	}
}

func TestIsCrypto(t *testing.T) {
	cases := []struct {
		symbol string
		want   bool
	}{
		{"AAPL", false},
		{"MSFT", false},
		{"BTC/USD", true},
		{"ETH/USD", true},
	}
	for _, tc := range cases {
		if got := isCrypto(tc.symbol); got != tc.want {
			t.Errorf("isCrypto(%q) = %v, want %v", tc.symbol, got, tc.want)
		}
	}
}
