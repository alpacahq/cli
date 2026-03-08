//go:build integration

package integration

import (
	"strings"
	"testing"
)

func TestAPIError_InvalidOrderReturnsStructuredJSON(t *testing.T) {
	t.Parallel()
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
	t.Parallel()
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
	t.Parallel()
	_, stderr, _ := alpacaFail(t,
		"order", "get", "00000000-0000-0000-0000-000000000000",
	)
	output := string(stderr)
	if !strings.Contains(output, "Error:") {
		t.Errorf("expected human-readable error with 'Error:' prefix, got: %s", output)
	}
}

func TestAPIError_NonExistentPosition(t *testing.T) {
	t.Parallel()
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
	t.Parallel()
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
	t.Parallel()
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
	t.Parallel()
	_, _, code := alpacaFail(t,
		"order", "get", "00000000-0000-0000-0000-000000000000",
	)
	if code != 1 {
		t.Errorf("expected exit code 1 for 404 API error, got %d", code)
	}
}

func TestAPIError_JSONErrorStructure(t *testing.T) {
	t.Parallel()
	_, stderr, code := alpacaFail(t,
		"order", "get", "00000000-0000-0000-0000-000000000000",
		"--json",
	)
	if code == 0 {
		t.Fatal("expected non-zero exit code")
	}

	errMap := parseJSONMap(t, stderr)
	requireFields(t, errMap, "error", "status")

	if errMsg, ok := errMap["error"].(string); !ok || errMsg == "" {
		t.Errorf("expected non-empty error message, got %v", errMap["error"])
	}
	if status, ok := errMap["status"].(float64); !ok || status < 400 {
		t.Errorf("expected HTTP status >= 400, got %v", errMap["status"])
	}
}
