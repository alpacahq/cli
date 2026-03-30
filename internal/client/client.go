package client

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"math/rand/v2"
	"net/http"
	"net/http/httptrace"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/alpacahq/cli/internal/config"
)

const (
	ExitAPIError  = 1
	ExitAuthError = 2
)

const maxRetries = 3

type Client struct {
	HTTP        *http.Client
	BaseURL     string
	DataURL     string
	APIKey      string
	Secret      string
	AccessToken string
	UserAgent   string
	Verbose     bool
	Debug       bool
	Quiet       bool
	Trace       bool
	Timeout     time.Duration
}

type APIError struct {
	StatusCode int    `json:"status"`
	Code       int    `json:"code"`
	Message    string `json:"message"`
	Method     string `json:"method,omitempty"`
	Path       string `json:"path,omitempty"`
	RequestID  string `json:"request_id,omitempty"`
	hint       string
	retryAfter time.Duration
}

func (e *APIError) Error() string {
	if e.Message == "" {
		return fmt.Sprintf("API error (HTTP %d)", e.StatusCode)
	}
	if e.Code > 0 {
		return fmt.Sprintf("%s [%d] (HTTP %d)", e.Message, e.Code, e.StatusCode)
	}
	if e.StatusCode > 0 {
		return fmt.Sprintf("%s (HTTP %d)", e.Message, e.StatusCode)
	}
	return e.Message
}

func (e *APIError) ExitCode() int {
	if e.StatusCode == 401 {
		return ExitAuthError
	}
	return ExitAPIError
}

func (e *APIError) Hint() string {
	if e.hint != "" {
		return e.hint
	}
	switch e.StatusCode {
	case 422:
		return "Validation error. Check parameter values — common issues: invalid qty, bad price, unknown symbol, or missing required fields."
	case 429:
		return "Rate limited. Reduce request frequency or add delays between calls."
	case 401:
		return "Invalid credentials. Run `alpaca profile login` to re-authenticate."
	case 403:
		return "Forbidden. Check your permissions, account status, minimum order size, or feature availability."
	}
	return ""
}

var Version = "dev"

func New(cfg *config.Resolved) *Client {
	return &Client{
		HTTP:        &http.Client{Timeout: 30 * time.Second},
		BaseURL:     strings.TrimRight(cfg.BaseURL, "/"),
		DataURL:     strings.TrimRight(cfg.DataURL, "/"),
		APIKey:      cfg.APIKey,
		Secret:      cfg.SecretKey,
		AccessToken: cfg.AccessToken,
		UserAgent:   "alpaca-cli/" + Version,
		Timeout:     30 * time.Second,
	}
}

func (c *Client) SetTimeout(d time.Duration) {
	c.Timeout = d
	c.HTTP.Timeout = d
}

// Do sends an HTTP request against baseURL + path with optional query params
// and body. All other Client methods are thin wrappers around Do.
func (c *Client) Do(method, baseURL, path string, params url.Values, body any) (json.RawMessage, error) {
	u := baseURL + path
	if len(params) > 0 {
		u += "?" + params.Encode()
	}
	return c.doWithRetry(method, u, body)
}

func (c *Client) Get(path string, params url.Values) (json.RawMessage, error) {
	return c.Do("GET", c.BaseURL, path, params, nil)
}

func (c *Client) Post(path string, params url.Values, body any) (json.RawMessage, error) {
	return c.Do("POST", c.BaseURL, path, params, body)
}

func (c *Client) Put(path string, params url.Values, body any) (json.RawMessage, error) {
	return c.Do("PUT", c.BaseURL, path, params, body)
}

func (c *Client) Patch(path string, params url.Values, body any) (json.RawMessage, error) {
	return c.Do("PATCH", c.BaseURL, path, params, body)
}

func (c *Client) Delete(path string, params url.Values) (json.RawMessage, error) {
	return c.Do("DELETE", c.BaseURL, path, params, nil)
}

func (c *Client) GetData(path string, params url.Values) (json.RawMessage, error) {
	return c.Do("GET", c.DataURL, path, params, nil)
}

func (c *Client) RawRequest(method, fullURL string, body any) (json.RawMessage, error) {
	return c.doWithRetry(method, fullURL, body)
}

func (c *Client) doWithRetry(method, reqURL string, body any) (json.RawMessage, error) {
	var lastErr error
	for attempt := range maxRetries {
		result, err := c.do(method, reqURL, body)
		if err == nil {
			return result, nil
		}
		lastErr = err

		apiErr, ok := err.(*APIError)
		if !ok || !isRetryable(apiErr.StatusCode) {
			return nil, err
		}

		delay := retryDelay(apiErr, attempt)
		if c.Verbose {
			fmt.Fprintf(os.Stderr, "  retrying in %s (attempt %d/%d)\n", delay, attempt+1, maxRetries)
		} else if !c.Quiet && apiErr.StatusCode == 429 {
			fmt.Fprintf(os.Stderr, "Rate limited, retrying in %s...\n", delay)
		}
		time.Sleep(delay)
	}
	return nil, lastErr
}

func isRetryable(status int) bool {
	return status == 429 || status == 500 || status == 502 || status == 503 || status == 504
}

func retryDelay(apiErr *APIError, attempt int) time.Duration {
	if apiErr.StatusCode == 429 && apiErr.retryAfter > 0 {
		return apiErr.retryAfter
	}
	base := time.Duration(math.Pow(2, float64(attempt))) * 500 * time.Millisecond
	jitter := time.Duration(rand.Int64N(int64(base/2) + 1))
	return base + jitter
}

func (c *Client) do(method, reqURL string, body any) (json.RawMessage, error) {
	start := time.Now()

	var reqData []byte
	var bodyReader io.Reader
	if body != nil {
		var err error
		reqData, err = json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("encoding request body: %w", err)
		}
		bodyReader = bytes.NewReader(reqData)
	}

	req, err := http.NewRequest(method, reqURL, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	if c.AccessToken != "" {
		req.Header.Set("Authorization", "Bearer "+c.AccessToken)
	} else {
		req.Header.Set("APCA-API-KEY-ID", c.APIKey)
		req.Header.Set("APCA-API-SECRET-KEY", c.Secret)
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("User-Agent", c.UserAgent)

	if c.Debug {
		fmt.Fprintf(os.Stderr, "→ %s %s\n", method, c.scrub(reqURL))
		c.debugHeaders("→", req.Header)
		if len(reqData) > 0 {
			fmt.Fprintf(os.Stderr, "→ %s\n", string(reqData))
		}
	}

	var tt *traceTimings
	if c.Trace {
		tt = newTrace(os.Stderr, method, c.scrub(reqURL))
		req = req.WithContext(httptrace.WithClientTrace(req.Context(), tt.clientTrace()))
	}

	resp, err := c.HTTP.Do(req)
	if err != nil {
		return nil, &APIError{
			Message: fmt.Sprintf("could not reach %s: %v", c.scrub(reqURL), c.scrubErr(err)),
			hint:    "check your internet connection and base URL. Run `alpaca doctor` to verify configuration",
		}
	}
	defer func() { _ = resp.Body.Close() }()

	elapsed := time.Since(start)
	if tt != nil {
		tt.finish(os.Stderr, resp.StatusCode, elapsed)
	} else if c.Verbose {
		fmt.Fprintf(os.Stderr, "%s %s → %d (%dms)\n", method, c.scrub(reqURL), resp.StatusCode, elapsed.Milliseconds())
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response: %w", err)
	}

	if c.Debug {
		c.debugHeaders("←", resp.Header)
		if len(respBody) > 0 {
			fmt.Fprintf(os.Stderr, "← %s\n", c.scrub(string(respBody)))
		}
	}

	if resp.StatusCode >= 400 {
		apiErr := &APIError{StatusCode: resp.StatusCode, Method: method, Path: reqURL, RequestID: resp.Header.Get("X-Request-Id")}
		if json.Unmarshal(respBody, apiErr) != nil || apiErr.Message == "" {
			apiErr.Message = strings.TrimSpace(string(respBody))
			if apiErr.Message == "" {
				apiErr.Message = http.StatusText(resp.StatusCode)
			}
		}
		if resp.StatusCode == 401 {
			var probe map[string]any
			if json.Unmarshal(respBody, &probe) != nil {
				apiErr.hint = "Received a non-API response (possible proxy or wrong URL). Verify your base URL with `alpaca doctor`."
			}
		}
		if resp.StatusCode == 429 {
			if ra := resp.Header.Get("Retry-After"); ra != "" {
				if secs, err := strconv.Atoi(ra); err == nil {
					apiErr.retryAfter = time.Duration(secs) * time.Second
				}
			}
		}
		return nil, apiErr
	}

	if len(respBody) == 0 {
		return json.RawMessage("null"), nil
	}

	return json.RawMessage(respBody), nil
}

func (c *Client) debugHeaders(prefix string, h http.Header) {
	for name, vals := range h {
		for _, v := range vals {
			fmt.Fprintf(os.Stderr, "%s %s: %s\n", prefix, name, c.scrub(v))
		}
	}
}

func (c *Client) scrub(s string) string {
	if c.APIKey != "" {
		s = strings.ReplaceAll(s, c.APIKey, "[REDACTED]")
	}
	if c.Secret != "" {
		s = strings.ReplaceAll(s, c.Secret, "[REDACTED]")
	}
	if c.AccessToken != "" {
		s = strings.ReplaceAll(s, c.AccessToken, "[REDACTED]")
	}
	return s
}

func (c *Client) scrubErr(err error) error {
	if err == nil {
		return nil
	}
	msg := c.scrub(err.Error())
	if msg == err.Error() {
		return err
	}
	return fmt.Errorf("%s", msg)
}

// traceTimings prints each HTTP phase to stderr as it completes.
type traceTimings struct {
	w                      io.Writer
	dnsStart, connectStart time.Time
	tlsStart, wroteRequest time.Time
	gotFirstByte           time.Time
}

func newTrace(w io.Writer, method, url string) *traceTimings {
	fmt.Fprintf(w, "trace: %s %s\n", method, url)
	return &traceTimings{w: w}
}

func (t *traceTimings) clientTrace() *httptrace.ClientTrace {
	return &httptrace.ClientTrace{
		DNSStart: func(_ httptrace.DNSStartInfo) { t.dnsStart = time.Now() },
		DNSDone: func(_ httptrace.DNSDoneInfo) {
			fmt.Fprintf(t.w, "  dns:     %dms\n", time.Since(t.dnsStart).Milliseconds())
		},
		ConnectStart: func(_, _ string) { t.connectStart = time.Now() },
		ConnectDone: func(_, addr string, _ error) {
			fmt.Fprintf(t.w, "  tcp:     %dms  (%s)\n", time.Since(t.connectStart).Milliseconds(), addr)
		},
		TLSHandshakeStart: func() { t.tlsStart = time.Now() },
		TLSHandshakeDone: func(_ tls.ConnectionState, _ error) {
			fmt.Fprintf(t.w, "  tls:     %dms\n", time.Since(t.tlsStart).Milliseconds())
		},
		WroteRequest:         func(_ httptrace.WroteRequestInfo) { t.wroteRequest = time.Now() },
		GotFirstResponseByte: func() { t.gotFirstByte = time.Now() },
		GotConn: func(info httptrace.GotConnInfo) {
			if info.Reused {
				fmt.Fprintf(t.w, "  conn:    reused (%s)\n", info.Conn.RemoteAddr())
			}
		},
	}
}

func (t *traceTimings) finish(w io.Writer, status int, elapsed time.Duration) {
	if !t.wroteRequest.IsZero() && !t.gotFirstByte.IsZero() {
		fmt.Fprintf(w, "  ttfb:    %dms\n", t.gotFirstByte.Sub(t.wroteRequest).Milliseconds())
	}
	fmt.Fprintf(w, "  total:   %dms → %d\n", elapsed.Milliseconds(), status)
}
