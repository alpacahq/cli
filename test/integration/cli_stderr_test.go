package integration

import (
	"encoding/json"
	"strings"
	"testing"
)

// stripCLIStderr removes debug and retry status lines so the remaining
// stderr is the JSON error object. A 429 before a final 403/404 writes
// "Rate limited, retrying..." (or the verbose "retrying in ..." line)
// alongside --debug arrow lines; leaving those in breaks JSON parsing.
func stripCLIStderr(stderr string) string {
	var b strings.Builder
	for _, line := range strings.Split(stderr, "\n") {
		if strings.HasPrefix(line, "→ ") || strings.HasPrefix(line, "← ") {
			continue
		}
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "Rate limited, retrying") || strings.HasPrefix(trimmed, "retrying in ") {
			continue
		}
		b.WriteString(line)
		b.WriteByte('\n')
	}
	return b.String()
}

func TestStripCLIStderrDropsRetryStatus(t *testing.T) {
	stderr := strings.Join([]string{
		"→ GET https://paper-api.alpaca.markets/v2/wallets/travel-rule/vasps?q=Coinbase",
		"← {\"message\":\"rate limited\"}",
		"Rate limited, retrying in 1.5s...",
		"→ GET https://paper-api.alpaca.markets/v2/wallets/travel-rule/vasps?q=Coinbase",
		"← {\"message\":\"forbidden\"}",
		"{",
		`  "error": "forbidden",`,
		`  "status": 403`,
		"}",
	}, "\n")

	got := stripCLIStderr(stderr)
	var errObj map[string]any
	if err := json.Unmarshal([]byte(got), &errObj); err != nil {
		t.Fatalf("stderr is not JSON after stripping retry status: %v\n%s", err, got)
	}
	if errObj["status"] != float64(403) {
		t.Fatalf("status = %v, want 403", errObj["status"])
	}
}

func TestStripCLIStderrDropsVerboseRetryStatus(t *testing.T) {
	stderr := "  retrying in 500ms (attempt 1/3)\n{\n  \"status\": 404\n}\n"
	got := stripCLIStderr(stderr)
	var errObj map[string]any
	if err := json.Unmarshal([]byte(got), &errObj); err != nil {
		t.Fatalf("stderr is not JSON after stripping verbose retry status: %v\n%s", err, got)
	}
	if errObj["status"] != float64(404) {
		t.Fatalf("status = %v, want 404", errObj["status"])
	}
}
