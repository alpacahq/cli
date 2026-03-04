package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"math/rand/v2"
	"net/http"
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
	HTTP      *http.Client
	BaseURL   string
	DataURL   string
	APIKey    string
	Secret    string
	UserAgent string
	Verbose   bool
	Quiet     bool
	Timeout   time.Duration
}

type APIError struct {
	StatusCode int    `json:"status"`
	Code       int    `json:"code"`
	Message    string `json:"message"`
	Method     string `json:"method,omitempty"`
	Path       string `json:"path,omitempty"`
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
	if e.StatusCode == 401 || e.StatusCode == 403 {
		return ExitAuthError
	}
	return ExitAPIError
}

func (e *APIError) Hint() string {
	if e.hint != "" {
		return e.hint
	}
	switch e.StatusCode {
	case 429:
		return "Rate limited. Reduce request frequency or add delays between calls."
	case 401:
		return "Invalid credentials. Run `alpaca profile login` to re-authenticate."
	case 403:
		return "Access denied. Check your API key permissions or account status."
	}
	return ""
}

var Version = "dev"

func New(cfg *config.Resolved) *Client {
	return &Client{
		HTTP:      &http.Client{Timeout: 30 * time.Second},
		BaseURL:   strings.TrimRight(cfg.BaseURL, "/"),
		DataURL:   strings.TrimRight(cfg.DataURL, "/"),
		APIKey:    cfg.APIKey,
		Secret:    cfg.SecretKey,
		UserAgent: "alpaca-cli/" + Version,
		Timeout:   30 * time.Second,
	}
}

func (c *Client) SetTimeout(d time.Duration) {
	c.Timeout = d
	c.HTTP.Timeout = d
}

func (c *Client) Get(path string, params url.Values) (json.RawMessage, error) {
	u := c.tradingURL(path, params)
	return c.doWithRetry("GET", u, nil)
}

func (c *Client) Post(path string, body any) (json.RawMessage, error) {
	u := c.tradingURL(path, nil)
	return c.doWithRetry("POST", u, body)
}

func (c *Client) Put(path string, body any) (json.RawMessage, error) {
	u := c.tradingURL(path, nil)
	return c.doWithRetry("PUT", u, body)
}

func (c *Client) Patch(path string, body any) (json.RawMessage, error) {
	u := c.tradingURL(path, nil)
	return c.doWithRetry("PATCH", u, body)
}

func (c *Client) Delete(path string, params url.Values) (json.RawMessage, error) {
	u := c.tradingURL(path, params)
	return c.doWithRetry("DELETE", u, nil)
}

func (c *Client) GetData(path string, params url.Values) (json.RawMessage, error) {
	u := c.dataURL(path, params)
	return c.doWithRetry("GET", u, nil)
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

	var bodyReader io.Reader
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("encoding request body: %w", err)
		}
		bodyReader = bytes.NewReader(data)
	}

	req, err := http.NewRequest(method, reqURL, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	req.Header.Set("APCA-API-KEY-ID", c.APIKey)
	req.Header.Set("APCA-API-SECRET-KEY", c.Secret)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("User-Agent", c.UserAgent)

	resp, err := c.HTTP.Do(req)
	if err != nil {
		return nil, &APIError{
			Message: fmt.Sprintf("could not reach %s: %v", c.scrub(reqURL), c.scrubErr(err)),
			hint:    "check your internet connection and base URL. Run `alpaca profile status` to verify configuration",
		}
	}
	defer func() { _ = resp.Body.Close() }()

	elapsed := time.Since(start)
	if c.Verbose {
		fmt.Fprintf(os.Stderr, "%s %s → %d (%dms)\n", method, c.scrub(reqURL), resp.StatusCode, elapsed.Milliseconds())
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response: %w", err)
	}

	if resp.StatusCode >= 400 {
		apiErr := &APIError{StatusCode: resp.StatusCode, Method: method, Path: reqURL}
		if json.Unmarshal(respBody, apiErr) != nil || apiErr.Message == "" {
			apiErr.Message = strings.TrimSpace(string(respBody))
			if apiErr.Message == "" {
				apiErr.Message = http.StatusText(resp.StatusCode)
			}
		}
		if resp.StatusCode == 401 {
			var probe map[string]any
			if json.Unmarshal(respBody, &probe) != nil {
				apiErr.hint = "Received a non-API response (possible proxy or wrong URL). Verify your base URL with `alpaca profile status`."
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

func (c *Client) scrub(s string) string {
	if c.APIKey != "" {
		s = strings.ReplaceAll(s, c.APIKey, "[REDACTED]")
	}
	if c.Secret != "" {
		s = strings.ReplaceAll(s, c.Secret, "[REDACTED]")
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

func (c *Client) tradingURL(path string, params url.Values) string {
	u := c.BaseURL + path
	if len(params) > 0 {
		u += "?" + params.Encode()
	}
	return u
}

func (c *Client) dataURL(path string, params url.Values) string {
	u := c.DataURL + path
	if len(params) > 0 {
		u += "?" + params.Encode()
	}
	return u
}
