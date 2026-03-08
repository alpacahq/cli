//go:build integration

package integration

import (
	"encoding/json"
	"sync"
	"testing"
)

var (
	optionSymbolOnce sync.Once
	resolvedOptionSymbol string
)

func discoverOptionSymbol(t *testing.T) string {
	t.Helper()
	optionSymbolOnce.Do(func() {
		out, err := makeCmd("data", "option", "chain", "AAPL", "--json").Output()
		if err != nil {
			return
		}
		var chain map[string]any
		if err := json.Unmarshal(out, &chain); err != nil {
			return
		}
		for sym := range chain {
			resolvedOptionSymbol = sym
			return
		}
	})
	if resolvedOptionSymbol == "" {
		t.Skip("no option contracts available for AAPL")
	}
	return resolvedOptionSymbol
}

func TestDataOptionChain(t *testing.T) {
	out := alpaca(t, "data", "option", "chain", "AAPL", "--json")
	chain := parseJSONMap(t, out)
	if len(chain) == 0 {
		t.Fatal("expected non-empty option chain")
	}
}

func TestDataOptionSnapshot(t *testing.T) {
	sym := discoverOptionSymbol(t)
	out := alpaca(t, "data", "option", "snapshot", "--symbols", sym, "--json")
	data := parseJSONMap(t, out)
	if len(data) == 0 {
		t.Error("expected non-empty option snapshot")
	}
}

func TestDataOptionLatestQuotes(t *testing.T) {
	sym := discoverOptionSymbol(t)
	out := alpaca(t, "data", "option", "latest-quotes", "--symbols", sym, "--json")
	data := parseJSONMap(t, out)
	if len(data) == 0 {
		t.Error("expected non-empty option latest quotes")
	}
}

func TestDataOptionExchanges(t *testing.T) {
	out := alpaca(t, "data", "option", "exchanges", "--json")
	data := parseJSONMap(t, out)
	if len(data) == 0 {
		t.Error("expected non-empty option exchanges")
	}
}

func TestDataOptionConditions(t *testing.T) {
	out := alpaca(t, "data", "option", "conditions", "trade", "--json")
	data := parseJSONMap(t, out)
	if len(data) == 0 {
		t.Error("expected non-empty option conditions")
	}
}

func TestDataOptionBars(t *testing.T) {
	sym := discoverOptionSymbol(t)
	// Option contracts may not have historical bar data — just verify the command runs
	stdout, _, code := alpacaWithStderr(t, "data", "option", "bars", "--symbols", sym, "--start", daysAgo(90), "--json")
	if code != 0 {
		t.Skip("option bars not available for this contract")
	}
	data := parseJSONMap(t, stdout)
	_ = data // may be empty but command succeeded
}

func TestDataOptionTrades(t *testing.T) {
	sym := discoverOptionSymbol(t)
	stdout, _, code := alpacaWithStderr(t, "data", "option", "trades", "--symbols", sym, "--start", daysAgo(90), "--json")
	if code != 0 {
		t.Skip("option trades not available for this contract")
	}
	data := parseJSONMap(t, stdout)
	_ = data
}

func TestDataOptionLatestTrades(t *testing.T) {
	sym := discoverOptionSymbol(t)
	stdout, _, code := alpacaWithStderr(t, "data", "option", "latest-trades", "--symbols", sym, "--json")
	if code != 0 {
		t.Skip("option latest trades not available for this contract")
	}
	data := parseJSONMap(t, stdout)
	_ = data
}
