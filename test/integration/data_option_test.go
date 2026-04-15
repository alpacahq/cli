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
		out, err := makeCmd("data", "option", "chain", "--underlying-symbol", "AAPL").Output()
		if err != nil {
			return
		}
		var resp struct {
			Snapshots map[string]json.RawMessage `json:"snapshots"`
		}
		if err := json.Unmarshal(out, &resp); err != nil {
			return
		}
		var fallback string
		for sym, raw := range resp.Snapshots {
			if fallback == "" {
				fallback = sym
			}
			var snap struct {
				LatestTrade *struct{} `json:"latestTrade"`
			}
			if json.Unmarshal(raw, &snap) == nil && snap.LatestTrade != nil {
				resolvedOptionSymbol = sym
				return
			}
		}
		resolvedOptionSymbol = fallback
	})
	if resolvedOptionSymbol == "" {
		t.Skip("no option contracts available for AAPL")
	}
	return resolvedOptionSymbol
}

func TestDataOptionChain(t *testing.T) {
	t.Parallel()
	out := alpacaRetry(t, "data", "option", "chain", "--underlying-symbol", "AAPL")
	chain := parseJSONMap(t, out)
	requireFields(t, chain, "snapshots")
	snapshots, _ := chain["snapshots"].(map[string]any)
	if len(snapshots) == 0 {
		t.Fatal("expected non-empty snapshots in option chain")
	}
}

func TestDataOptionSnapshot(t *testing.T) {
	t.Parallel()
	sym := discoverOptionSymbol(t)
	out := alpacaRetry(t, "data", "option", "snapshot", "--symbols", sym)
	data := parseJSONMap(t, out)
	if len(data) == 0 {
		t.Error("expected non-empty option snapshot")
	}
}

func TestDataOptionLatestQuotes(t *testing.T) {
	t.Parallel()
	sym := discoverOptionSymbol(t)
	out := alpacaRetry(t, "data", "option", "latest-quotes", "--symbols", sym)
	data := parseJSONMap(t, out)
	if len(data) == 0 {
		t.Error("expected non-empty option latest quotes")
	}
}

func TestDataOptionExchanges(t *testing.T) {
	t.Parallel()
	out := alpacaRetry(t, "data", "option", "exchanges")
	data := parseJSONMap(t, out)
	if len(data) == 0 {
		t.Error("expected non-empty option exchanges")
	}
}

func TestDataOptionConditions(t *testing.T) {
	t.Parallel()
	out := alpacaRetry(t, "data", "option", "conditions", "--ticktype", "trade")
	data := parseJSONMap(t, out)
	if len(data) == 0 {
		t.Error("expected non-empty option conditions")
	}
}

func TestDataOptionSnapshot_Fields(t *testing.T) {
	t.Parallel()
	sym := discoverOptionSymbol(t)
	out := alpacaRetry(t, "data", "option", "snapshot", "--symbols", sym)
	data := parseJSONMap(t, out)
	snapshots, _ := data["snapshots"].(map[string]any)
	if len(snapshots) == 0 {
		t.Fatal("expected non-empty snapshots")
	}
	snap, _ := snapshots[sym].(map[string]any)
	if snap == nil {
		t.Fatalf("expected snapshot for %s", sym)
	}

	if greeks, ok := snap["greeks"].(map[string]any); ok {
		for _, field := range []string{"delta", "gamma", "theta", "vega"} {
			if _, has := greeks[field]; !has {
				t.Errorf("greeks missing %q", field)
			}
		}
	}
	if quote, ok := snap["latestQuote"].(map[string]any); ok {
		requireFields(t, quote, "ap", "bp", "as", "bs", "t")
	}
	if trade, ok := snap["latestTrade"].(map[string]any); ok {
		requireFields(t, trade, "p", "s", "t", "x")
	}
}

func TestDataOptionTrades(t *testing.T) {
	t.Parallel()
	sym := discoverOptionSymbol(t)
	stdout, _, code := alpacaWithStderr(t, "data", "option", "trades", "--symbols", sym, "--start", daysAgo(90))
	if code != 0 {
		t.Skip("option trades not available for this contract")
	}
	data := parseJSONMap(t, stdout)
	requireFields(t, data, "trades")
}

func TestDataOptionLatestTrades(t *testing.T) {
	t.Parallel()
	sym := discoverOptionSymbol(t)
	stdout, _, code := alpacaWithStderr(t, "data", "option", "latest-trades", "--symbols", sym)
	if code != 0 {
		t.Skip("option latest trades not available for this contract")
	}
	data := parseJSONMap(t, stdout)
	requireFields(t, data, "trades")
}
