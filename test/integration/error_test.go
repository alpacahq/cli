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

func TestAPIError_NonExistentPosition(t *testing.T) {
	_, stderr, code := alpacaFail(t, "position", "get", "ZZZZZZZZZ", "--json")
	if code == 0 {
		t.Fatal("expected non-zero exit code for non-existent position")
	}
	errMap := parseJSONMap(t, stderr)
	if errMap["error"] == nil || errMap["error"] == "" {
		t.Error("expected 'error' field in JSON error")
	}
}

func TestAPIError_NonExistentOrderCancel(t *testing.T) {
	_, stderr, code := alpacaFail(t,
		"order", "cancel", "00000000-0000-0000-0000-000000000001",
		"--json",
	)
	if code == 0 {
		t.Fatal("expected non-zero exit code for canceling non-existent order")
	}
	errMap := parseJSONMap(t, stderr)
	if errMap["error"] == nil || errMap["error"] == "" {
		t.Error("expected 'error' field in JSON error")
	}
}

func TestAPIError_NonExistentAsset(t *testing.T) {
	_, stderr, code := alpacaFail(t, "asset", "get", "ZZZZZZZZZ", "--json")
	if code == 0 {
		t.Fatal("expected non-zero exit code for non-existent asset")
	}
	errMap := parseJSONMap(t, stderr)
	if errMsg, ok := errMap["error"].(string); !ok || errMsg == "" {
		t.Error("expected non-empty error message")
	}
}

func TestExitCode_APIErrorIs1(t *testing.T) {
	_, _, code := alpacaFail(t,
		"order", "get", "00000000-0000-0000-0000-000000000000",
	)
	if code != 1 {
		t.Errorf("expected exit code 1 for 404 API error, got %d", code)
	}
}

func TestAPIError_JSONErrorStructure(t *testing.T) {
	_, stderr, code := alpacaFail(t,
		"order", "get", "00000000-0000-0000-0000-000000000000",
		"--json",
	)
	if code == 0 {
		t.Fatal("expected non-zero exit code")
	}

	errMap := parseJSONMap(t, stderr)

	for _, field := range []string{"error", "status"} {
		if _, ok := errMap[field]; !ok {
			t.Errorf("JSON error missing required field %q", field)
		}
	}

	if errMsg, ok := errMap["error"].(string); !ok || errMsg == "" {
		t.Errorf("expected non-empty error message, got %v", errMap["error"])
	}
	if status, ok := errMap["status"].(float64); !ok || status < 400 {
		t.Errorf("expected HTTP status >= 400, got %v", errMap["status"])
	}
}
