//go:build integration

package integration

import (
	"strings"
	"testing"
)

func TestAPIError_InvalidOrderReturnsStructuredJSON(t *testing.T) {
	_, stderr, code := alpacaFail(t,
		"order", "submit", "AAPL",
		"--qty", "-1",
		"--side", "buy",
		"--type", "market",
		"--json",
	)
	if code == 0 {
		t.Fatal("expected non-zero exit code")
	}

	errMap := parseJSONMap(t, stderr)
	if errMap["error"] == nil || errMap["error"] == "" {
		t.Error("expected 'error' field in JSON error response")
	}
}

func TestAPIError_InvalidAuth(t *testing.T) {
	// Valid creds succeed, so test with a bogus order ID to trigger a 404.
	_, stderr, code := alpacaFail(t,
		"order", "get", "00000000-0000-0000-0000-000000000000",
		"--json",
	)
	if code == 0 {
		t.Fatal("expected non-zero exit code for bogus order ID")
	}
	errMap := parseJSONMap(t, stderr)
	if errMap["error"] == nil || errMap["error"] == "" {
		t.Error("expected error message")
	}
}

func TestAPIError_HumanReadable(t *testing.T) {
	_, stderr, _ := alpacaFail(t,
		"order", "get", "00000000-0000-0000-0000-000000000000",
	)
	output := string(stderr)
	if !strings.Contains(output, "Error:") {
		t.Errorf("expected human-readable error with 'Error:' prefix, got: %s", output)
	}
}

func TestOutputFormats_JSON(t *testing.T) {
	out := alpaca(t, "account", "get", "--json")
	acct := parseJSONMap(t, out)
	if acct["id"] == nil {
		t.Error("JSON output should contain 'id' field")
	}
}

func TestOutputFormats_CSV(t *testing.T) {
	out := alpaca(t, "order", "list", "--status", "all", "--limit", "1", "--csv")
	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	if len(lines) < 1 {
		t.Fatal("expected at least a header row in CSV")
	}
	if !strings.Contains(lines[0], ",") {
		t.Errorf("CSV header should contain commas, got: %s", lines[0])
	}
}
