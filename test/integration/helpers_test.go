//go:build integration

package integration

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

var cliBinary string

func TestMain(m *testing.M) {
	if os.Getenv("ALPACA_TEST_API_KEY") == "" || os.Getenv("ALPACA_TEST_SECRET_KEY") == "" {
		fmt.Fprintln(os.Stderr, "Skipping integration tests: ALPACA_TEST_API_KEY and ALPACA_TEST_SECRET_KEY must be set")
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

// alpacaExpectFail runs the CLI and expects it to fail.
func alpacaExpectFail(t *testing.T, args ...string) (stdout, stderr []byte, exitCode int) {
	t.Helper()
	cmd := exec.Command(cliBinary, args...)
	cmd.Env = cliEnv()

	out, err := cmd.Output()
	if err == nil {
		t.Fatalf("expected alpaca %s to fail, but it succeeded:\n%s", strings.Join(args, " "), string(out))
	}

	var exitErr *exec.ExitError
	if errors.As(err, &exitErr) {
		return out, exitErr.Stderr, exitErr.ExitCode()
	}
	t.Fatalf("unexpected error type: %v", err)
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

func cliEnv() []string {
	baseURL := os.Getenv("ALPACA_TEST_BASE_URL")
	if baseURL == "" {
		baseURL = "https://paper-api.alpaca.markets"
	}
	dataURL := os.Getenv("ALPACA_TEST_DATA_URL")
	if dataURL == "" {
		dataURL = "https://data.alpaca.markets"
	}

	return append(os.Environ(),
		"ALPACA_API_KEY="+os.Getenv("ALPACA_TEST_API_KEY"),
		"ALPACA_SECRET_KEY="+os.Getenv("ALPACA_TEST_SECRET_KEY"),
		"ALPACA_BASE_URL="+baseURL,
		"ALPACA_DATA_URL="+dataURL,
		"ALPACA_CONFIG_DIR="+os.TempDir()+"/alpaca-cli-test-config",
	)
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

func containsID(items []map[string]any, id string) bool {
	for _, item := range items {
		if item["id"] == id {
			return true
		}
	}
	return false
}
