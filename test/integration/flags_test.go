//go:build integration

package integration

import (
	"bytes"
	"errors"
	"os/exec"
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

func TestQuietEnvVar(t *testing.T) {
	t.Parallel()
	cmd := makeCmd("order", "list")
	cmd.Env = append(cmd.Env, "ALPACA_QUIET=1")
	var stderrBuf bytes.Buffer
	cmd.Stderr = &stderrBuf
	stdout, err := cmd.Output()
	if err != nil {
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) {
			t.Fatalf("expected exit 0, got %d (stderr: %s)", exitErr.ExitCode(), stderrBuf.Bytes())
		}
		t.Fatalf("unexpected error: %v", err)
	}
	if len(stdout) == 0 {
		t.Fatal("expected JSON output on stdout")
	}
	se := stderrBuf.String()
	if strings.Contains(se, "Hint") || strings.Contains(se, "hint") {
		t.Error("ALPACA_QUIET should suppress hint output")
	}
}

func TestHelpAll(t *testing.T) {
	t.Parallel()
	out := alpaca(t, "--help-all")
	s := string(out)
	for _, cmd := range []string{"order", "position", "account", "data", "watchlist", "tokenization"} {
		if !strings.Contains(s, cmd) {
			t.Errorf("--help-all should contain %q command", cmd)
		}
	}
}

func TestHelpHasNoUpdateNoticeOnStderr(t *testing.T) {
	t.Parallel()
	_, stderr, code := alpacaWithStderr(t, "--help")
	if code != 0 {
		t.Fatalf("expected exit 0, got %d", code)
	}
	if len(stderr) != 0 {
		t.Fatalf("--help should not write to stderr, got: %s", stderr)
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
