//go:build integration

package integration

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"
)

var (
	cliBinary     string
	testConfigDir string
	testProfile   string
)

// testProfileName is used when routing an OAuth token via a profile file.
// ALPACA_ACCESS_TOKEN is no longer a supported env var; the token must live
// in a profile YAML and be selected with ALPACA_PROFILE.
const testProfileName = "integration-test"

func TestMain(m *testing.M) {
	hasAPIKey := os.Getenv("ALPACA_TEST_API_KEY") != "" && os.Getenv("ALPACA_TEST_SECRET_KEY") != ""
	hasToken := os.Getenv("ALPACA_TEST_ACCESS_TOKEN") != ""
	if !hasAPIKey && !hasToken {
		fmt.Fprintln(os.Stderr, "Integration tests require credentials: set ALPACA_TEST_API_KEY+ALPACA_TEST_SECRET_KEY or ALPACA_TEST_ACCESS_TOKEN")
		os.Exit(1)
	}

	dir, err := os.MkdirTemp("", "alpaca-cli-test")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create temp dir: %v\n", err)
		os.Exit(1)
	}
	defer os.RemoveAll(dir)

	binary := filepath.Join(dir, "alpaca")
	if runtime.GOOS == "windows" {
		binary += ".exe"
	}

	fmt.Println("Building CLI binary...")
	build := exec.Command("go", "build", "-o", binary, "./cmd/alpaca")
	build.Dir = projectRoot()
	build.Stdout = os.Stdout
	build.Stderr = os.Stderr
	if err := build.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to build: %v\n", err)
		os.Exit(1)
	}

	cliBinary = binary

	testConfigDir = filepath.Join(os.TempDir(), "alpaca-cli-test-config")
	if err := writeTestProfile(hasToken); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to write test profile: %v\n", err)
		os.Exit(1)
	}

	os.Exit(m.Run())
}

// writeTestProfile stashes the access token in a profile YAML when tests run
// with ALPACA_TEST_ACCESS_TOKEN. For API-key runs no profile is written; the
// env vars flow through atomic-bundle resolution directly.
func writeTestProfile(hasToken bool) error {
	if !hasToken {
		return nil
	}

	profilesDir := filepath.Join(testConfigDir, "profiles")
	if err := os.MkdirAll(profilesDir, 0o700); err != nil {
		return err
	}

	yaml := fmt.Sprintf(
		"api_key: \"\"\nsecret_key: \"\"\naccess_token: %q\n",
		os.Getenv("ALPACA_TEST_ACCESS_TOKEN"),
	)
	if err := os.WriteFile(filepath.Join(profilesDir, testProfileName+".yaml"), []byte(yaml), 0o600); err != nil {
		return err
	}

	gcfg := fmt.Sprintf("default_profile: %s\noutput: json\n", testProfileName)
	if err := os.WriteFile(filepath.Join(testConfigDir, "config.yaml"), []byte(gcfg), 0o600); err != nil {
		return err
	}

	testProfile = testProfileName
	return nil
}

// alpaca runs the CLI and returns stdout. Fatals on non-zero exit.
func alpaca(t *testing.T, args ...string) []byte {
	t.Helper()
	cmd := exec.Command(cliBinary, args...)
	cmd.Env = cliEnv()

	out, err := cmd.Output()
	if err != nil {
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) {
			t.Fatalf("alpaca %s failed (exit %d):\nstdout: %s\nstderr: %s",
				strings.Join(args, " "), exitErr.ExitCode(), string(out), string(exitErr.Stderr))
		}
		t.Fatalf("alpaca %s failed: %v", strings.Join(args, " "), err)
	}
	return out
}

// alpacaRetry is like alpaca but retries up to 3 times on transient network
// errors (EOF, connection reset). Only use for idempotent read-only calls.
func alpacaRetry(t *testing.T, args ...string) []byte {
	t.Helper()
	const maxAttempts = 3
	for attempt := range maxAttempts {
		cmd := exec.Command(cliBinary, append(args, "--timeout", "5")...)
		cmd.Env = cliEnv()

		out, err := cmd.Output()
		if err != nil {
			var exitErr *exec.ExitError
			if errors.As(err, &exitErr) {
				stderr := string(exitErr.Stderr)
				if attempt < maxAttempts-1 && isTransientError(stderr) {
					t.Logf("transient error on attempt %d, retrying: %s", attempt+1, stderr)
					time.Sleep(time.Duration(attempt+1) * 500 * time.Millisecond)
					continue
				}
				t.Fatalf("alpaca %s failed (exit %d):\nstdout: %s\nstderr: %s",
					strings.Join(args, " "), exitErr.ExitCode(), string(out), stderr)
			}
			t.Fatalf("alpaca %s failed: %v", strings.Join(args, " "), err)
		}
		return out
	}
	panic("unreachable")
}

func isTransientError(stderr string) bool {
	for _, sig := range []string{": EOF", "connection reset", "connection refused", "TLS handshake timeout", "context deadline exceeded"} {
		if strings.Contains(stderr, sig) {
			return true
		}
	}
	return false
}

// alpacaWithStderr runs the CLI and returns both stdout and stderr.
// Does NOT fatal on non-zero exit — caller must check.
func alpacaWithStderr(t *testing.T, args ...string) (stdout, stderr []byte, exitCode int) {
	t.Helper()
	cmd := exec.Command(cliBinary, args...)
	cmd.Env = cliEnv()
	var stderrBuf bytes.Buffer
	cmd.Stderr = &stderrBuf
	stdout, err := cmd.Output()
	if err != nil {
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) {
			return stdout, stderrBuf.Bytes(), exitErr.ExitCode()
		}
		t.Fatalf("alpaca %s failed unexpectedly: %v", strings.Join(args, " "), err)
	}
	return stdout, stderrBuf.Bytes(), 0
}

func alpacaJSONOrStructuredError(t *testing.T, args ...string) (map[string]any, bool) {
	t.Helper()
	stdout, stderr, exitCode := alpacaWithStderr(t, args...)
	if exitCode == 0 {
		return parseJSONMap(t, stdout), true
	}
	if len(stdout) != 0 {
		t.Fatalf("alpaca %s wrote stdout on error: %s", strings.Join(args, " "), string(stdout))
	}
	data := parseJSONMap(t, stderr)
	requireFields(t, data, "error", "status")
	return data, false
}

// alpacaFail runs the CLI and expects non-zero exit. Fatals if it succeeds.
func alpacaFail(t *testing.T, args ...string) (stdout, stderr []byte, exitCode int) {
	t.Helper()
	cmd := exec.Command(cliBinary, args...)
	cmd.Env = cliEnv()
	stdout, err := cmd.Output()
	if err != nil {
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) {
			return stdout, exitErr.Stderr, exitErr.ExitCode()
		}
		t.Fatalf("alpaca %s failed unexpectedly: %v", strings.Join(args, " "), err)
	}
	t.Fatalf("alpaca %s succeeded but expected failure", strings.Join(args, " "))
	return nil, nil, 0
}

func parseJSON[T any](t *testing.T, data []byte) T {
	t.Helper()
	var v T
	if err := json.Unmarshal(data, &v); err != nil {
		t.Fatalf("Failed to parse JSON:\n%s\nError: %v", string(data), err)
	}
	return v
}

func parseJSONMap(t *testing.T, data []byte) map[string]any {
	return parseJSON[map[string]any](t, data)
}

func parseJSONArray(t *testing.T, data []byte) []map[string]any {
	return parseJSON[[]map[string]any](t, data)
}

// requireFields asserts that the JSON map has all listed fields.
func requireFields(t *testing.T, m map[string]any, fields ...string) {
	t.Helper()
	for _, f := range fields {
		if _, ok := m[f]; !ok {
			t.Errorf("missing required field %q", f)
		}
	}
}

// requireArrayNonEmpty parses a JSON array and fails if empty.
func requireArrayNonEmpty(t *testing.T, data []byte) []map[string]any {
	t.Helper()
	arr := parseJSONArray(t, data)
	if len(arr) == 0 {
		t.Fatal("expected non-empty array")
	}
	return arr
}

func cliEnv() []string {
	// Integration tests always run against the real paper API. Unsetting
	// ALPACA_LIVE_TRADE (the default) is all the routing we need - live is
	// strictly opt-in.
	env := append(os.Environ(),
		"ALPACA_CONFIG_DIR="+testConfigDir,
		"ALPACA_LIVE_TRADE=",
	)

	if testProfile != "" {
		// OAuth runs: env API keys would shadow the profile token under
		// atomic-bundle resolution, so blank them out and point at the
		// profile written in TestMain.
		env = append(env,
			"ALPACA_PROFILE="+testProfile,
			"ALPACA_API_KEY=",
			"ALPACA_SECRET_KEY=",
		)
	} else {
		env = append(env,
			"ALPACA_API_KEY="+os.Getenv("ALPACA_TEST_API_KEY"),
			"ALPACA_SECRET_KEY="+os.Getenv("ALPACA_TEST_SECRET_KEY"),
		)
	}
	return env
}

func projectRoot() string {
	wd, _ := os.Getwd()
	for {
		if _, err := os.Stat(filepath.Join(wd, "go.mod")); err == nil {
			return wd
		}
		parent := filepath.Dir(wd)
		if parent == wd {
			break
		}
		wd = parent
	}
	return "."
}

// makeCmd creates an exec.Command for the CLI with the test environment.
func makeCmd(args ...string) *exec.Cmd {
	cmd := exec.Command(cliBinary, args...)
	cmd.Env = cliEnv()
	return cmd
}

func containsID(items []map[string]any, id string) bool {
	for _, item := range items {
		if item["id"] == id {
			return true
		}
	}
	return false
}

func daysAgo(n int) string {
	return time.Now().AddDate(0, 0, -n).Format("2006-01-02")
}

func monthRange(monthsAgo int) (start, end string) {
	now := time.Now()
	y, m, _ := now.Date()
	first := time.Date(y, m-time.Month(monthsAgo), 1, 0, 0, 0, 0, time.UTC)
	last := first.AddDate(0, 1, -1)
	return first.Format("2006-01-02"), last.Format("2006-01-02")
}

func pollFor(t *testing.T, timeout time.Duration, desc string, fn func() bool) {
	t.Helper()
	deadline := time.Now().Add(timeout)
	for {
		if fn() {
			return
		}
		if time.Now().After(deadline) {
			t.Fatalf("timed out after %v waiting for %s", timeout, desc)
		}
		time.Sleep(200 * time.Millisecond)
	}
}

// submitTestOrder places a safe GTC limit buy at $1.00 (will never fill).
// Returns the order ID. Registers t.Cleanup to cancel.
//
// Each parallel test must pass a unique symbol to avoid wash-trade conflicts.
// Bracket orders create sell-side child legs; if another test has a buy order
// on the same symbol, the API rejects it as a potential wash trade.
func submitTestOrder(t *testing.T, symbol string) string {
	t.Helper()
	out := alpaca(t, "order", "submit",
		"--symbol", symbol,
		"--qty", "1",
		"--side", "buy",
		"--type", "limit",
		"--limit-price", "1.00",
		"--time-in-force", "gtc",
	)
	order := parseJSONMap(t, out)
	id, ok := order["id"].(string)
	if !ok || id == "" {
		t.Fatal("order missing id")
	}
	t.Cleanup(func() {
		cmd := exec.Command(cliBinary, "order", "cancel", "--order-id", id)
		cmd.Env = cliEnv()
		_ = cmd.Run()
	})
	pollFor(t, 5*time.Second, "order to be retrievable", func() bool {
		_, _, code := alpacaWithStderr(t, "order", "get", "--order-id", id)
		return code == 0
	})
	return id
}

// submitCryptoFill places a crypto market buy for $10 notional and polls
// until a position appears. Crypto trades 24/7 so this works regardless of
// equity market hours. Registers t.Cleanup to close the position.
//
// Each test must pass a unique crypto symbol to avoid wash-trade conflicts.
// Closing a position creates a sell order; if the next test tries to buy
// the same symbol before that sell fills, the API rejects it.
func submitCryptoFill(t *testing.T, symbol string) string {
	t.Helper()
	alpaca(t, "order", "submit",
		"--symbol", symbol,
		"--notional", "10",
		"--side", "buy",
		"--type", "market",
		"--time-in-force", "gtc",
	)

	t.Cleanup(func() {
		cmd := exec.Command(cliBinary, "position", "close", "--symbol-or-asset-id", symbol)
		cmd.Env = cliEnv()
		_ = cmd.Run()
	})
	pollFor(t, 15*time.Second, symbol+" position to appear", func() bool {
		_, _, code := alpacaWithStderr(t, "position", "get", "--symbol-or-asset-id", symbol)
		return code == 0
	})
	return symbol
}
