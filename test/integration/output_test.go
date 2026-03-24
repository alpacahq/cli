//go:build integration

package integration

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestOutput_AccountGet(t *testing.T) {
	t.Parallel()

	t.Run("json_default", func(t *testing.T) {
		out := alpaca(t, "account", "get")
		acct := parseJSONMap(t, out)
		requireFields(t, acct, "id", "status")
	})

	t.Run("csv", func(t *testing.T) {
		out := alpaca(t, "account", "get", "--csv")
		s := strings.TrimSpace(string(out))
		if !strings.Contains(s, ",") {
			t.Error("CSV output should contain commas")
		}
		if !strings.Contains(s, "id") {
			t.Error("CSV output should contain 'id' column")
		}
	})
}

func TestOutput_OrderList(t *testing.T) {
	t.Parallel()

	t.Run("json_default", func(t *testing.T) {
		out := alpaca(t, "order", "list", "--status", "all", "--limit", "1")
		_ = parseJSONArray(t, out)
	})

	t.Run("csv", func(t *testing.T) {
		out := alpaca(t, "order", "list", "--status", "all", "--limit", "1", "--csv")
		lines := strings.Split(strings.TrimSpace(string(out)), "\n")
		if len(lines) < 1 {
			t.Fatal("expected at least a header row")
		}
		if !strings.Contains(lines[0], ",") {
			t.Error("CSV header should contain commas")
		}
	})
}

func TestOutput_AssetGet(t *testing.T) {
	t.Parallel()
	out := alpaca(t, "asset", "get", "--symbol-or-asset-id", "AAPL")
	acct := parseJSONMap(t, out)
	if acct["symbol"] != "AAPL" {
		t.Error("JSON output for asset get should contain AAPL")
	}
}

func TestOutput_AssetList_CSV(t *testing.T) {
	t.Parallel()
	out := alpaca(t, "asset", "list", "--asset-class", "us_equity", "--status", "active", "--csv")
	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	if len(lines) < 2 {
		t.Fatal("expected header + at least one data row")
	}
	if !strings.Contains(lines[0], "symbol") {
		t.Error("CSV header should contain 'symbol' column")
	}
}

func TestOutput_PositionList_CSV(t *testing.T) {
	t.Parallel()
	stdout, _, code := alpacaWithStderr(t, "position", "list", "--csv")
	if code != 0 {
		t.Fatalf("position list --csv exited %d", code)
	}
	s := strings.TrimSpace(string(stdout))
	if s == "" {
		t.Fatal("position list --csv should emit at least a header row")
	}
	lines := strings.Split(s, "\n")
	if !strings.Contains(lines[0], ",") {
		t.Error("CSV header should contain commas")
	}
}

func TestOutput_DataBars_CSV(t *testing.T) {
	t.Parallel()
	out := alpaca(t, "data", "bars", "--symbol", "AAPL",
		"--start", daysAgo(100),
		"--end", daysAgo(93),
		"--timeframe", "1Day",
		"--csv",
	)
	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	if len(lines) < 1 {
		t.Fatal("expected at least a header row")
	}
	if !strings.Contains(lines[0], ",") {
		t.Error("CSV header should contain commas")
	}
}

func TestOutput_Clock_CSV(t *testing.T) {
	t.Parallel()
	out := alpaca(t, "clock", "--csv")
	s := strings.TrimSpace(string(out))
	if !strings.Contains(s, ",") {
		t.Error("CSV output should contain commas")
	}
	if !strings.Contains(s, "is_open") {
		t.Error("CSV output should contain 'is_open' column")
	}
}

// --- --jq flag ---

func TestOutput_JQ_ExtractField(t *testing.T) {
	t.Parallel()
	out := alpaca(t, "account", "get", "--jq", ".id")
	var id string
	if err := json.Unmarshal(out, &id); err != nil {
		t.Fatalf("--jq '.id' should return a JSON string, got: %s", out)
	}
	if id == "" {
		t.Error("expected non-empty account id")
	}
}

func TestOutput_JQ_ArrayElement(t *testing.T) {
	t.Parallel()
	out := alpaca(t, "order", "list", "--status", "all", "--limit", "5", "--jq", ".[0]")
	s := strings.TrimSpace(string(out))
	if s == "null" {
		t.Skip("no orders to test --jq array indexing")
	}
	order := parseJSONMap(t, out)
	requireFields(t, order, "id", "symbol")
}

func TestOutput_JQ_MapArray(t *testing.T) {
	t.Parallel()
	out := alpaca(t, "order", "list", "--status", "all", "--limit", "3", "--jq", "[.[].symbol]")
	var symbols []string
	if err := json.Unmarshal(out, &symbols); err != nil {
		t.Fatalf("--jq '[.[].symbol]' should return array of strings, got: %s", out)
	}
}

func TestOutput_JQ_WithCSV(t *testing.T) {
	t.Parallel()
	out := alpaca(t, "account", "get", "--jq", "{id, status}", "--csv")
	s := strings.TrimSpace(string(out))
	if !strings.Contains(s, ",") {
		t.Error("--jq + --csv should produce comma-separated output")
	}
}

func TestOutput_JQ_InvalidExpression(t *testing.T) {
	t.Parallel()
	_, _, code := alpacaFail(t, "account", "get", "--jq", ".[invalid")
	if code == 0 {
		t.Fatal("invalid jq expression should fail")
	}
}
