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

var cliBinary string

func TestMain(m *testing.M) {
	hasAPIKey := os.Getenv("ALPACA_TEST_API_KEY") != "" && os.Getenv("ALPACA_TEST_SECRET_KEY") != ""
	hasToken := os.Getenv("ALPACA_TEST_ACCESS_TOKEN") != ""
	if !hasAPIKey && !hasToken {
		fmt.Fprintln(os.Stderr, "Skipping integration tests: set ALPACA_TEST_API_KEY+ALPACA_TEST_SECRET_KEY or ALPACA_TEST_ACCESS_TOKEN")
		os.Exit(0)
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
	os.Exit(m.Run())
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
	baseURL := os.Getenv("ALPACA_TEST_BASE_URL")
	if baseURL == "" {
		baseURL = "https://paper-api.alpaca.markets"
	}
	dataURL := os.Getenv("ALPACA_TEST_DATA_URL")
	if dataURL == "" {
		dataURL = "https://data.alpaca.markets"
	}

	env := append(os.Environ(),
		"ALPACA_BASE_URL="+baseURL,
		"ALPACA_DATA_URL="+dataURL,
		"ALPACA_CONFIG_DIR="+os.TempDir()+"/alpaca-cli-test-config",
	)

	if token := os.Getenv("ALPACA_TEST_ACCESS_TOKEN"); token != "" {
		env = append(env, "ALPACA_ACCESS_TOKEN="+token)
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

// submitTestOrder places a safe GTC limit buy of AAPL at $1.00 (will never fill).
// Returns the order ID. Registers t.Cleanup to cancel.
func submitTestOrder(t *testing.T) string {
	t.Helper()
	out := alpaca(t, "order", "submit",
		"--symbol", "AAPL",
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
		cmd := exec.Command(cliBinary, "order", "cancel", id)
		cmd.Env = cliEnv()
		_ = cmd.Run()
	})
	pollFor(t, 5*time.Second, "order to be retrievable", func() bool {
		_, _, code := alpacaWithStderr(t, "order", "get", id)
		return code == 0
	})
	return id
}

// submitCryptoFill places a BTC/USD market buy for $1 notional and polls
// until a position appears. Crypto trades 24/7 so this works regardless of
// equity market hours. Registers t.Cleanup to close the position.
func submitCryptoFill(t *testing.T) string {
	t.Helper()
	alpaca(t, "order", "submit",
		"--symbol", "BTC/USD",
		"--notional", "10",
		"--side", "buy",
		"--type", "market",
		"--time-in-force", "gtc",
	)

	symbol := "BTC/USD"
	t.Cleanup(func() {
		cmd := exec.Command(cliBinary, "position", "close", symbol)
		cmd.Env = cliEnv()
		_ = cmd.Run()
	})
	pollFor(t, 15*time.Second, "BTC/USD position to appear", func() bool {
		_, _, code := alpacaWithStderr(t, "position", "get", symbol)
		return code == 0
	})
	return symbol
}
