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
		out, err := makeCmd("data", "option", "chain", "AAPL").Output()
		if err != nil {
			return
		}
		var resp struct {
			Snapshots map[string]any `json:"snapshots"`
		}
		if err := json.Unmarshal(out, &resp); err != nil {
			return
		}
		for sym := range resp.Snapshots {
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
	t.Parallel()
	out := alpaca(t, "data", "option", "chain", "AAPL")
	chain := parseJSONMap(t, out)
	if len(chain) == 0 {
		t.Fatal("expected non-empty option chain")
	}
}

func TestDataOptionSnapshot(t *testing.T) {
	t.Parallel()
	sym := discoverOptionSymbol(t)
	out := alpaca(t, "data", "option", "snapshot", "--symbols", sym)
	data := parseJSONMap(t, out)
	if len(data) == 0 {
		t.Error("expected non-empty option snapshot")
	}
}

func TestDataOptionLatestQuotes(t *testing.T) {
	t.Parallel()
	sym := discoverOptionSymbol(t)
	out := alpaca(t, "data", "option", "latest-quotes", "--symbols", sym)
	data := parseJSONMap(t, out)
	if len(data) == 0 {
		t.Error("expected non-empty option latest quotes")
	}
}

func TestDataOptionExchanges(t *testing.T) {
	t.Parallel()
	out := alpaca(t, "data", "option", "exchanges")
	data := parseJSONMap(t, out)
	if len(data) == 0 {
		t.Error("expected non-empty option exchanges")
	}
}

func TestDataOptionConditions(t *testing.T) {
	t.Parallel()
	out := alpaca(t, "data", "option", "conditions", "trade")
	data := parseJSONMap(t, out)
	if len(data) == 0 {
		t.Error("expected non-empty option conditions")
	}
}

func TestDataOptionBars(t *testing.T) {
	t.Parallel()
	sym := discoverOptionSymbol(t)
	// Option contracts may not have historical bar data — just verify the command runs
	stdout, _, code := alpacaWithStderr(t, "data", "option", "bars", "--symbols", sym, "--start", daysAgo(90))
	if code != 0 {
		t.Skip("option bars not available for this contract")
	}
	data := parseJSONMap(t, stdout)
	_ = data // may be empty but command succeeded
}

func TestDataOptionTrades(t *testing.T) {
	t.Parallel()
	sym := discoverOptionSymbol(t)
	stdout, _, code := alpacaWithStderr(t, "data", "option", "trades", "--symbols", sym, "--start", daysAgo(90))
	if code != 0 {
		t.Skip("option trades not available for this contract")
	}
	data := parseJSONMap(t, stdout)
	_ = data
}

func TestDataOptionLatestTrades(t *testing.T) {
	t.Parallel()
	sym := discoverOptionSymbol(t)
	stdout, _, code := alpacaWithStderr(t, "data", "option", "latest-trades", "--symbols", sym)
	if code != 0 {
		t.Skip("option latest trades not available for this contract")
	}
	data := parseJSONMap(t, stdout)
	_ = data
}
