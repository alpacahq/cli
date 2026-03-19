//go:build integration

package integration

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
	"testing"
	"time"
)

func TestAPI_Get(t *testing.T) {
	t.Parallel()
	out := alpaca(t, "api", "GET", "/v2/clock")
	clock := parseJSONMap(t, out)
	requireFields(t, clock, "is_open")
}

func TestAPI_Post(t *testing.T) {
	out := alpaca(t, "api", "POST", "/v2/watchlists",
		"--body", `{"name":"api-test-post","symbols":[]}`,
	)
	wl := parseJSONMap(t, out)
	id, _ := wl["id"].(string)
	if id == "" {
		t.Fatal("api POST did not return watchlist with id")
	}
	t.Cleanup(func() {
		_ = makeCmd("api", "DELETE", "/v2/watchlists/"+id).Run()
	})

	if wl["name"] != "api-test-post" {
		t.Errorf("expected name api-test-post, got %v", wl["name"])
	}
}

func TestAPI_PostFromStdin(t *testing.T) {
	out := makeCmd("api", "POST", "/v2/watchlists")
	out.Stdin = bytes.NewBufferString(`{"name":"api-test-stdin","symbols":[]}`)

	stdout, err := out.Output()
	if err != nil {
		t.Fatalf("api POST via stdin failed: %v", err)
	}

	wl := parseJSONMap(t, stdout)
	id, _ := wl["id"].(string)
	if id == "" {
		t.Fatal("api POST via stdin did not return watchlist with id")
	}
	t.Cleanup(func() {
		_ = makeCmd("api", "DELETE", "/v2/watchlists/"+id).Run()
	})
}

func TestAPI_BodyFlagTakesPrecedenceOverStdin(t *testing.T) {
	cmd := makeCmd("api", "POST", "/v2/watchlists",
		"--body", `{"name":"api-test-body-wins","symbols":[]}`,
	)
	cmd.Stdin = bytes.NewBufferString(`{"name":"api-test-stdin-loses","symbols":[]}`)

	stdout, err := cmd.Output()
	if err != nil {
		t.Fatalf("api POST with --body and stdin failed: %v", err)
	}

	wl := parseJSONMap(t, stdout)
	id, _ := wl["id"].(string)
	if id == "" {
		t.Fatal("api POST did not return watchlist with id")
	}
	t.Cleanup(func() {
		_ = makeCmd("api", "DELETE", "/v2/watchlists/"+id).Run()
	})
	if wl["name"] != "api-test-body-wins" {
		t.Fatalf("expected --body to take precedence, got %v", wl["name"])
	}
}

func TestAPI_Patch(t *testing.T) {
	out := alpaca(t, "api", "GET", "/v2/account/configurations")
	original := parseJSONMap(t, out)
	origVal, _ := original["trade_confirm_email"].(string)

	newVal := "none"
	if origVal == "none" {
		newVal = "all"
	}

	out = alpaca(t, "api", "PATCH", "/v2/account/configurations",
		"--body", `{"trade_confirm_email":"`+newVal+`"}`,
	)
	updated := parseJSONMap(t, out)
	if updated["trade_confirm_email"] != newVal {
		t.Errorf("expected trade_confirm_email %q, got %v", newVal, updated["trade_confirm_email"])
	}

	t.Cleanup(func() {
		_ = makeCmd("api", "PATCH", "/v2/account/configurations",
			"--body", `{"trade_confirm_email":"`+origVal+`"}`).Run()
	})
}

func TestAPI_Delete(t *testing.T) {
	out := alpaca(t, "api", "POST", "/v2/watchlists",
		"--body", `{"name":"api-test-delete","symbols":[]}`,
	)
	wl := parseJSONMap(t, out)
	id := wl["id"].(string)

	alpaca(t, "api", "DELETE", "/v2/watchlists/"+id)

	_, _, code := alpacaWithStderr(t, "api", "GET", "/v2/watchlists/"+id)
	if code == 0 {
		t.Error("watchlist should be deleted")
	}
}

func TestAPI_UseDataAPI(t *testing.T) {
	t.Parallel()
	out := alpaca(t, "api", "GET", "/v2/stocks/AAPL/trades/latest", "--use-data-api")
	trade := parseJSONMap(t, out)
	requireFields(t, trade, "trade")
}

func TestAPI_QueryFlag(t *testing.T) {
	t.Parallel()
	out := alpaca(t, "api", "GET", "/v2/orders", "--query", "status=all&limit=1")
	orders := parseJSONArray(t, out)
	if len(orders) > 1 {
		t.Errorf("--query limit=1 returned %d orders", len(orders))
	}
}

func TestAPI_GetDoesNotReadOpenStdin(t *testing.T) {
	t.Parallel()

	reader, writer, err := os.Pipe()
	if err != nil {
		t.Fatalf("failed to create pipe: %v", err)
	}
	defer reader.Close()
	defer writer.Close()

	cmd := makeCmd("api", "GET", "/v2/clock")
	cmd.Stdin = reader

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if err := cmd.Start(); err != nil {
		t.Fatalf("failed to start api GET: %v", err)
	}

	done := make(chan error, 1)
	go func() {
		done <- cmd.Wait()
	}()

	select {
	case err := <-done:
		if err != nil {
			var exitErr *exec.ExitError
			if errors.As(err, &exitErr) {
				t.Fatalf("api GET failed with exit %d: %s", exitErr.ExitCode(), stderr.String())
			}
			t.Fatalf("api GET failed: %v", err)
		}
	case <-time.After(2 * time.Second):
		_ = cmd.Process.Kill()
		t.Fatal("api GET should not block on open stdin")
	}
}
