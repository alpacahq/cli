//go:build integration

package integration

import (
	"strings"
	"testing"
)

func TestOutput_Table_AccountGet(t *testing.T) {
	out := alpaca(t, "account", "get")
	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	if len(lines) < 2 {
		t.Fatal("table output should have at least a header and data row")
	}
}

func TestOutput_JSON_AccountGet(t *testing.T) {
	out := alpaca(t, "account", "get", "--json")
	acct := parseJSONMap(t, out)
	requireFields(t, acct, "id", "status")
}

func TestOutput_CSV_AccountGet(t *testing.T) {
	out := alpaca(t, "account", "get", "--csv")
	s := strings.TrimSpace(string(out))
	if !strings.Contains(s, ",") {
		t.Error("CSV output should contain commas")
	}
	if !strings.Contains(s, "id") {
		t.Error("CSV output should contain 'id' column")
	}
}

func TestOutput_Table_OrderList(t *testing.T) {
	out := alpaca(t, "order", "list", "--status", "all", "--limit", "1")
	_ = string(out)
}

func TestOutput_JSON_OrderList(t *testing.T) {
	out := alpaca(t, "order", "list", "--status", "all", "--limit", "1", "--json")
	_ = parseJSONArray(t, out)
}

func TestOutput_CSV_OrderList(t *testing.T) {
	out := alpaca(t, "order", "list", "--status", "all", "--limit", "1", "--csv")
	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	if len(lines) < 1 {
		t.Fatal("expected at least a header row")
	}
	if !strings.Contains(lines[0], ",") {
		t.Error("CSV header should contain commas")
	}
}

func TestOutput_Table_AssetGet(t *testing.T) {
	out := alpaca(t, "asset", "get", "AAPL")
	s := string(out)
	if !strings.Contains(s, "AAPL") {
		t.Error("table output for asset get should contain AAPL")
	}
}
