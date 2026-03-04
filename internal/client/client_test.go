package client

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/alpacahq/cli/internal/config"
)

func testClient(t *testing.T, handler http.HandlerFunc) *Client {
	t.Helper()
	srv := httptest.NewServer(handler)
	t.Cleanup(srv.Close)
	return &Client{
		HTTP:      &http.Client{Timeout: 5 * time.Second},
		BaseURL:   srv.URL,
		DataURL:   srv.URL,
		APIKey:    "PKTEST123456",
		Secret:    "secret_test_abc",
		UserAgent: "alpaca-cli/test",
	}
}

func TestAuthHeaders(t *testing.T) {
	var gotKeyID, gotSecret, gotUA string
	c := testClient(t, func(w http.ResponseWriter, r *http.Request) {
		gotKeyID = r.Header.Get("APCA-API-KEY-ID")
		gotSecret = r.Header.Get("APCA-API-SECRET-KEY")
		gotUA = r.Header.Get("User-Agent")
		w.WriteHeader(200)
		_, _ = w.Write([]byte(`{"ok":true}`))
	})

	_, err := c.Get("/v2/account", nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotKeyID != "PKTEST123456" {
		t.Errorf("expected API key header PKTEST123456, got %q", gotKeyID)
	}
	if gotSecret != "secret_test_abc" {
		t.Errorf("expected secret header, got %q", gotSecret)
	}
	if gotUA != "alpaca-cli/test" {
		t.Errorf("expected User-Agent alpaca-cli/test, got %q", gotUA)
	}
}

func TestGet200(t *testing.T) {
	c := testClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("expected GET, got %s", r.Method)
		}
		w.WriteHeader(200)
		_, _ = w.Write([]byte(`{"id":"abc","equity":"10000"}`))
	})

	data, err := c.Get("/v2/account", nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var m map[string]any
	if err := json.Unmarshal(data, &m); err != nil {
		t.Fatalf("failed to parse response: %v", err)
	}
	if m["id"] != "abc" {
		t.Errorf("expected id=abc, got %v", m["id"])
	}
}

func TestPostSendsBody(t *testing.T) {
	var gotBody string
	var gotContentType string
	c := testClient(t, func(w http.ResponseWriter, r *http.Request) {
		gotContentType = r.Header.Get("Content-Type")
		b := make([]byte, r.ContentLength)
		_, _ = r.Body.Read(b)
		gotBody = string(b)
		w.WriteHeader(200)
		_, _ = w.Write([]byte(`{"id":"order1"}`))
	})

	body := map[string]string{"symbol": "AAPL", "qty": "10"}
	_, err := c.Post("/v2/orders", nil, body)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotContentType != "application/json" {
		t.Errorf("expected Content-Type application/json, got %q", gotContentType)
	}
	if !strings.Contains(gotBody, `"symbol":"AAPL"`) {
		t.Errorf("body missing symbol: %s", gotBody)
	}
}

func TestError400(t *testing.T) {
	c := testClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(400)
		_, _ = w.Write([]byte(`{"code":40010001,"message":"qty is required"}`))
	})

	_, err := c.Get("/v2/orders", nil)
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	apiErr, ok := err.(*APIError)
	if !ok {
		t.Fatalf("expected *APIError, got %T", err)
	}
	if apiErr.StatusCode != 400 {
		t.Errorf("expected status 400, got %d", apiErr.StatusCode)
	}
	if apiErr.Code != 40010001 {
		t.Errorf("expected code 40010001, got %d", apiErr.Code)
	}
	if apiErr.Message != "qty is required" {
		t.Errorf("expected message 'qty is required', got %q", apiErr.Message)
	}
	if apiErr.ExitCode() != ExitAPIError {
		t.Errorf("expected exit code %d, got %d", ExitAPIError, apiErr.ExitCode())
	}
}

func TestError401(t *testing.T) {
	c := testClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(401)
		_, _ = w.Write([]byte(`{"message":"unauthorized"}`))
	})

	_, err := c.Get("/v2/account", nil)
	apiErr, ok := err.(*APIError)
	if !ok {
		t.Fatalf("expected *APIError, got %T", err)
	}
	if apiErr.ExitCode() != ExitAuthError {
		t.Errorf("expected exit code %d for 401, got %d", ExitAuthError, apiErr.ExitCode())
	}
	if apiErr.Hint() == "" {
		t.Error("expected hint for 401 error")
	}
}

func TestError403(t *testing.T) {
	c := testClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(403)
		_, _ = w.Write([]byte(`{"message":"forbidden"}`))
	})

	_, err := c.Get("/v2/account", nil)
	apiErr, ok := err.(*APIError)
	if !ok {
		t.Fatalf("expected *APIError, got %T", err)
	}
	if apiErr.ExitCode() != ExitAuthError {
		t.Errorf("expected exit code %d for 403, got %d", ExitAuthError, apiErr.ExitCode())
	}
}

func TestError422MalformedJSON(t *testing.T) {
	c := testClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(422)
		_, _ = w.Write([]byte(`not json at all`))
	})

	_, err := c.Get("/v2/orders", nil)
	apiErr, ok := err.(*APIError)
	if !ok {
		t.Fatalf("expected *APIError, got %T", err)
	}
	if apiErr.StatusCode != 422 {
		t.Errorf("expected status 422, got %d", apiErr.StatusCode)
	}
	if !strings.Contains(apiErr.Message, "not json at all") {
		t.Errorf("expected raw body in message, got %q", apiErr.Message)
	}
}

func TestError429Hint(t *testing.T) {
	c := testClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Retry-After", "5")
		w.WriteHeader(429)
		_, _ = w.Write([]byte(`{"message":"rate limit exceeded"}`))
	})

	_, err := c.Get("/v2/orders", nil)
	apiErr, ok := err.(*APIError)
	if !ok {
		t.Fatalf("expected *APIError, got %T", err)
	}
	if apiErr.StatusCode != 429 {
		t.Errorf("expected status 429, got %d", apiErr.StatusCode)
	}
	hint := apiErr.Hint()
	if !strings.Contains(strings.ToLower(hint), "rate limit") {
		t.Errorf("expected rate limit hint, got %q", hint)
	}
}

func TestEmpty204Response(t *testing.T) {
	c := testClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(204)
	})

	data, err := c.Delete("/v2/orders/abc", nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if string(data) != "null" {
		t.Errorf("expected null for empty body, got %s", string(data))
	}
}

func TestRetryOn500(t *testing.T) {
	attempts := 0
	c := testClient(t, func(w http.ResponseWriter, r *http.Request) {
		attempts++
		if attempts < 3 {
			w.WriteHeader(500)
			_, _ = w.Write([]byte(`{"message":"internal server error"}`))
			return
		}
		w.WriteHeader(200)
		_, _ = w.Write([]byte(`{"ok":true}`))
	})

	data, err := c.Get("/v2/account", nil)
	if err != nil {
		t.Fatalf("expected success after retries, got: %v", err)
	}
	if attempts != 3 {
		t.Errorf("expected 3 attempts, got %d", attempts)
	}
	var m map[string]any
	_ = json.Unmarshal(data, &m)
	if m["ok"] != true {
		t.Errorf("expected ok=true, got %v", m["ok"])
	}
}

func TestNoRetryOn400(t *testing.T) {
	attempts := 0
	c := testClient(t, func(w http.ResponseWriter, r *http.Request) {
		attempts++
		w.WriteHeader(400)
		_, _ = w.Write([]byte(`{"message":"bad request"}`))
	})

	_, err := c.Get("/v2/orders", nil)
	if err == nil {
		t.Fatal("expected error")
	}
	if attempts != 1 {
		t.Errorf("expected 1 attempt (no retry on 400), got %d", attempts)
	}
}

func TestScrubSecrets(t *testing.T) {
	c := &Client{
		APIKey: "PKTEST123456",
		Secret: "secretXYZ789",
	}

	input := "failed to reach https://api.alpaca.markets with key PKTEST123456 and secret secretXYZ789"
	result := c.scrub(input)

	if strings.Contains(result, "PKTEST123456") {
		t.Errorf("API key not scrubbed: %s", result)
	}
	if strings.Contains(result, "secretXYZ789") {
		t.Errorf("Secret not scrubbed: %s", result)
	}
	if !strings.Contains(result, "[REDACTED]") {
		t.Errorf("expected [REDACTED] in output: %s", result)
	}
}

func TestScrubEmptyCredentials(t *testing.T) {
	c := &Client{APIKey: "", Secret: ""}
	result := c.scrub("some text with no credentials")
	if result != "some text with no credentials" {
		t.Errorf("empty credentials should not alter text: %s", result)
	}
}

func TestNewClient(t *testing.T) {
	cfg := &config.Resolved{
		BaseURL:   "https://paper-api.alpaca.markets/",
		DataURL:   "https://data.alpaca.markets/",
		APIKey:    "PK123",
		SecretKey: "SK456",
	}
	c := New(cfg)
	if c.BaseURL != "https://paper-api.alpaca.markets" {
		t.Errorf("trailing slash not trimmed from BaseURL: %s", c.BaseURL)
	}
	if c.DataURL != "https://data.alpaca.markets" {
		t.Errorf("trailing slash not trimmed from DataURL: %s", c.DataURL)
	}
}

func TestDataURL(t *testing.T) {
	var gotPath string
	c := testClient(t, func(w http.ResponseWriter, r *http.Request) {
		gotPath = r.URL.Path
		w.WriteHeader(200)
		_, _ = w.Write([]byte(`{"bars":[]}`))
	})

	_, err := c.GetData("/v2/stocks/AAPL/bars", nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotPath != "/v2/stocks/AAPL/bars" {
		t.Errorf("expected path /v2/stocks/AAPL/bars, got %s", gotPath)
	}
}

func TestHTTPMethods(t *testing.T) {
	tests := []struct {
		name   string
		call   func(c *Client) (json.RawMessage, error)
		method string
	}{
		{"Get", func(c *Client) (json.RawMessage, error) { return c.Get("/path", nil) }, "GET"},
		{"Post", func(c *Client) (json.RawMessage, error) { return c.Post("/path", nil, nil) }, "POST"},
		{"Put", func(c *Client) (json.RawMessage, error) { return c.Put("/path", nil, nil) }, "PUT"},
		{"Patch", func(c *Client) (json.RawMessage, error) { return c.Patch("/path", nil, nil) }, "PATCH"},
		{"Delete", func(c *Client) (json.RawMessage, error) { return c.Delete("/path", nil) }, "DELETE"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var gotMethod string
			c := testClient(t, func(w http.ResponseWriter, r *http.Request) {
				gotMethod = r.Method
				w.WriteHeader(200)
				_, _ = w.Write([]byte(`{}`))
			})
			_, err := tt.call(c)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if gotMethod != tt.method {
				t.Errorf("expected %s, got %s", tt.method, gotMethod)
			}
		})
	}
}

func TestErrorString(t *testing.T) {
	tests := []struct {
		name string
		err  APIError
		want string
	}{
		{"empty message", APIError{StatusCode: 500}, "API error (HTTP 500)"},
		{"with code", APIError{StatusCode: 400, Code: 40010001, Message: "qty required"}, "qty required [40010001] (HTTP 400)"},
		{"message only", APIError{StatusCode: 422, Message: "invalid"}, "invalid (HTTP 422)"},
		{"no status", APIError{Message: "something"}, "something"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.err.Error()
			if got != tt.want {
				t.Errorf("Error() = %q, want %q", got, tt.want)
			}
		})
	}
}

// TestErrorPathInfrastructure verifies that HTTP error codes produce the
// correct exit code, hint, message, and JSON structure through the full
// client → APIError path. Only non-retryable codes are tested via HTTP to
// avoid retry delays; retryable codes (429, 500) are tested in struct tests.
func TestErrorPathInfrastructure(t *testing.T) {
	tests := []struct {
		name     string
		status   int
		body     string
		wantExit int
		wantHint string // lowercase substring; empty = expect no hint
		wantMsg  string // lowercase substring in message
	}{
		{
			name:     "401_unauthorized",
			status:   401,
			body:     `{"message":"unauthorized"}`,
			wantExit: ExitAuthError,
			wantHint: "credentials",
			wantMsg:  "unauthorized",
		},
		{
			name:     "403_forbidden",
			status:   403,
			body:     `{"message":"forbidden"}`,
			wantExit: ExitAuthError,
			wantHint: "denied",
			wantMsg:  "forbidden",
		},
		{
			name:     "404_not_found",
			status:   404,
			body:     `{"message":"order not found"}`,
			wantExit: ExitAPIError,
			wantMsg:  "not found",
		},
		{
			name:     "422_validation",
			status:   422,
			body:     `{"code":42210001,"message":"qty must be positive"}`,
			wantExit: ExitAPIError,
			wantMsg:  "qty must be positive",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := testClient(t, func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tt.status)
				_, _ = w.Write([]byte(tt.body))
			})

			_, err := c.Get("/v2/test", nil)
			if err == nil {
				t.Fatal("expected error")
			}

			apiErr, ok := err.(*APIError)
			if !ok {
				t.Fatalf("expected *APIError, got %T", err)
			}

			if apiErr.ExitCode() != tt.wantExit {
				t.Errorf("ExitCode() = %d, want %d", apiErr.ExitCode(), tt.wantExit)
			}

			hint := apiErr.Hint()
			if tt.wantHint != "" && !strings.Contains(strings.ToLower(hint), tt.wantHint) {
				t.Errorf("Hint() = %q, want substring %q", hint, tt.wantHint)
			}
			if tt.wantHint == "" && hint != "" {
				t.Errorf("expected no hint for %d, got %q", tt.status, hint)
			}

			if !strings.Contains(strings.ToLower(apiErr.Message), strings.ToLower(tt.wantMsg)) {
				t.Errorf("Message = %q, want substring %q", apiErr.Message, tt.wantMsg)
			}

			data, marshalErr := json.Marshal(apiErr)
			if marshalErr != nil {
				t.Fatalf("json.Marshal: %v", marshalErr)
			}
			var m map[string]any
			if err := json.Unmarshal(data, &m); err != nil {
				t.Fatalf("json.Unmarshal: %v", err)
			}
			if int(m["status"].(float64)) != tt.status {
				t.Errorf("JSON status = %v, want %d", m["status"], tt.status)
			}
			if m["message"] == nil || m["message"] == "" {
				t.Error("JSON missing 'message' field")
			}
		})
	}
}

// TestAPIErrorMethodAndPath verifies that the Method and Path fields are
// populated correctly in API errors and round-trip through JSON.
func TestAPIErrorMethodAndPath(t *testing.T) {
	c := testClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(400)
		_, _ = w.Write([]byte(`{"message":"bad request"}`))
	})

	_, err := c.Post("/v2/orders", nil, map[string]string{"qty": "-1"})
	apiErr, ok := err.(*APIError)
	if !ok {
		t.Fatalf("expected *APIError, got %T", err)
	}

	if apiErr.Method != "POST" {
		t.Errorf("Method = %q, want POST", apiErr.Method)
	}
	if !strings.Contains(apiErr.Path, "/v2/orders") {
		t.Errorf("Path = %q, want to contain /v2/orders", apiErr.Path)
	}

	data, _ := json.Marshal(apiErr)
	var m map[string]any
	_ = json.Unmarshal(data, &m)
	if m["method"] != "POST" {
		t.Errorf("JSON method = %v, want POST", m["method"])
	}
	if path, ok := m["path"].(string); !ok || !strings.Contains(path, "/v2/orders") {
		t.Errorf("JSON path = %v, want to contain /v2/orders", m["path"])
	}
}

// TestPostPutPatchPassQueryParams verifies that Post, Put, and Patch include
// query parameters in the request URL.
func TestPostPutPatchPassQueryParams(t *testing.T) {
	tests := []struct {
		name   string
		call   func(c *Client, params url.Values) (json.RawMessage, error)
		method string
	}{
		{"Post", func(c *Client, p url.Values) (json.RawMessage, error) { return c.Post("/path", p, nil) }, "POST"},
		{"Put", func(c *Client, p url.Values) (json.RawMessage, error) { return c.Put("/path", p, nil) }, "PUT"},
		{"Patch", func(c *Client, p url.Values) (json.RawMessage, error) { return c.Patch("/path", p, nil) }, "PATCH"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var gotQuery string
			c := testClient(t, func(w http.ResponseWriter, r *http.Request) {
				gotQuery = r.URL.RawQuery
				w.WriteHeader(200)
				_, _ = w.Write([]byte(`{}`))
			})

			params := url.Values{}
			params.Set("name", "Tech Stocks")
			params.Set("extra", "val")

			_, err := tt.call(c, params)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !strings.Contains(gotQuery, "name=Tech+Stocks") {
				t.Errorf("%s did not pass query params; got query %q", tt.method, gotQuery)
			}
			if !strings.Contains(gotQuery, "extra=val") {
				t.Errorf("%s missing extra param; got query %q", tt.method, gotQuery)
			}
		})
	}
}

// TestPostPutPatchNilParams verifies that nil params produce no query string.
func TestPostPutPatchNilParams(t *testing.T) {
	var gotQuery string
	c := testClient(t, func(w http.ResponseWriter, r *http.Request) {
		gotQuery = r.URL.RawQuery
		w.WriteHeader(200)
		_, _ = w.Write([]byte(`{}`))
	})

	_, err := c.Post("/path", nil, map[string]string{"key": "val"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotQuery != "" {
		t.Errorf("expected empty query string for nil params, got %q", gotQuery)
	}
}

// TestNetworkErrorHasHint verifies that connection failures produce an
// APIError with a helpful hint about checking connectivity.
func TestNetworkErrorHasHint(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	}))
	serverURL := srv.URL
	srv.Close()

	c := &Client{
		HTTP:      &http.Client{Timeout: 2 * time.Second},
		BaseURL:   serverURL,
		APIKey:    "test",
		Secret:    "test",
		UserAgent: "alpaca-cli/test",
	}

	_, err := c.Get("/v2/account", nil)
	if err == nil {
		t.Fatal("expected connection error")
	}

	apiErr, ok := err.(*APIError)
	if !ok {
		t.Fatalf("expected *APIError, got %T: %v", err, err)
	}

	if !strings.Contains(strings.ToLower(apiErr.Hint()), "internet connection") {
		t.Errorf("expected connection hint, got: %q", apiErr.Hint())
	}
}

// TestVerboseOutputScrubsSecrets verifies that verbose mode never leaks API
// keys or secrets to stderr, even when they appear in request URLs. This is
// a regression test: if someone adds header or body logging without scrubbing,
// this test will catch it.
func TestVerboseOutputScrubsSecrets(t *testing.T) {
	const apiKey = "PKTESTKEY123456789"
	const secret = "SECRETTESTVALUE987654"

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		_, _ = w.Write([]byte(`{"ok":true}`))
	}))
	t.Cleanup(srv.Close)

	c := &Client{
		HTTP:      &http.Client{Timeout: 5 * time.Second},
		BaseURL:   srv.URL,
		DataURL:   srv.URL,
		APIKey:    apiKey,
		Secret:    secret,
		UserAgent: "alpaca-cli/test",
		Verbose:   true,
	}

	origStderr := os.Stderr
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	os.Stderr = w

	// Embed the API key in the URL to stress-test scrubbing
	_, _ = c.Get("/v2/account?token="+apiKey, nil)

	_ = w.Close()
	os.Stderr = origStderr

	captured, _ := io.ReadAll(r)
	output := string(captured)

	if output == "" {
		t.Fatal("expected verbose output on stderr, got nothing")
	}
	if strings.Contains(output, apiKey) {
		t.Errorf("verbose output leaks API key:\n%s", output)
	}
	if strings.Contains(output, secret) {
		t.Errorf("verbose output leaks secret:\n%s", output)
	}
}
