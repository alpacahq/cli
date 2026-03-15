//go:build integration

package integration

import (
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
	if s != "" && !strings.Contains(s, ",") {
		t.Error("non-empty CSV output should contain commas")
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
