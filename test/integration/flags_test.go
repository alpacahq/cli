//go:build integration

package integration

import (
	"strings"
	"testing"
)

func TestVerboseFlag(t *testing.T) {
	t.Parallel()
	stdout, stderr, code := alpacaWithStderr(t, "account", "get", "--verbose")
	if code != 0 {
		t.Fatalf("expected exit 0, got %d (stderr: %s)", code, stderr)
	}
	_ = parseJSONMap(t, stdout)
	se := string(stderr)
	if !strings.Contains(se, "GET") && !strings.Contains(se, "get") {
		t.Errorf("--verbose stderr should contain HTTP method, got: %s", se)
	}
}

func TestDebugFlag(t *testing.T) {
	t.Parallel()
	stdout, stderr, code := alpacaWithStderr(t, "account", "get", "--debug")
	if code != 0 {
		t.Fatalf("expected exit 0, got %d (stderr: %s)", code, stderr)
	}
	_ = parseJSONMap(t, stdout)
	se := string(stderr)
	if !strings.Contains(se, "→") && !strings.Contains(se, "GET") {
		t.Errorf("--debug stderr should contain request details, got: %s", se)
	}
}

func TestTimeoutFlag(t *testing.T) {
	t.Parallel()
	out := alpaca(t, "account", "get", "--timeout", "10")
	acct := parseJSONMap(t, out)
	requireFields(t, acct, "id")
}

func TestQuietFlag(t *testing.T) {
	t.Parallel()
	_, stderr, code := alpacaWithStderr(t, "order", "list", "--quiet")
	if code != 0 {
		t.Fatalf("expected exit 0, got %d", code)
	}
	if len(stderr) > 0 {
		se := string(stderr)
		if strings.Contains(se, "Hint") || strings.Contains(se, "hint") {
			t.Error("--quiet should suppress hint output")
		}
	}
}

func TestHelpAll(t *testing.T) {
	t.Parallel()
	out := alpaca(t, "--help-all")
	s := string(out)
	for _, cmd := range []string{"order", "position", "account", "data", "watchlist"} {
		if !strings.Contains(s, cmd) {
			t.Errorf("--help-all should contain %q command", cmd)
		}
	}
}

func TestUnknownCommand(t *testing.T) {
	t.Parallel()
	_, stderr, code := alpacaWithStderr(t, "notacommand")
	if code == 0 {
		t.Fatal("unknown command should return non-zero exit code")
	}
	if len(stderr) == 0 {
		t.Error("unknown command should produce stderr output")
	}
}
