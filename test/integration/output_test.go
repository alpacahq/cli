//go:build integration

package integration

import (
	"strings"
	"testing"
)

func TestOutput_AccountGet(t *testing.T) {
	t.Parallel()

	t.Run("table", func(t *testing.T) {
		out := alpaca(t, "account", "get")
		lines := strings.Split(strings.TrimSpace(string(out)), "\n")
		if len(lines) < 2 {
			t.Fatal("table output should have at least a header and data row")
		}
	})

	t.Run("json", func(t *testing.T) {
		out := alpaca(t, "account", "get", "--json")
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

	t.Run("table", func(t *testing.T) {
		_ = string(alpaca(t, "order", "list", "--status", "all", "--limit", "1"))
	})

	t.Run("json", func(t *testing.T) {
		out := alpaca(t, "order", "list", "--status", "all", "--limit", "1", "--json")
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
	out := alpaca(t, "asset", "get", "AAPL")
	s := string(out)
	if !strings.Contains(s, "AAPL") {
		t.Error("table output for asset get should contain AAPL")
	}
}
