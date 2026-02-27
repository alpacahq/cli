package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/alpacahq/cli/internal/config"
)

const ExitAPIError = 1
const ExitAuthError = 2
const ExitValidationError = 3
const ExitNetworkError = 4

type Client struct {
	HTTP      *http.Client
	BaseURL   string
	DataURL   string
	APIKey    string
	Secret    string
	UserAgent string
}

type APIError struct {
	StatusCode int    `json:"-"`
	Code       int    `json:"code"`
	Message    string `json:"message"`
}

func (e *APIError) Error() string {
	if e.Message != "" {
		return e.Message
	}
	return fmt.Sprintf("API error (HTTP %d)", e.StatusCode)
}

func (e *APIError) ExitCode() int {
	if e.StatusCode == 401 || e.StatusCode == 403 {
		return ExitAuthError
	}
	return ExitAPIError
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
	}
}

func (c *Client) Get(path string, params url.Values) (json.RawMessage, error) {
	u := c.tradingURL(path, params)
	return c.do("GET", u, nil)
}

func (c *Client) Post(path string, body any) (json.RawMessage, error) {
	u := c.tradingURL(path, nil)
	return c.do("POST", u, body)
}

func (c *Client) Put(path string, body any) (json.RawMessage, error) {
	u := c.tradingURL(path, nil)
	return c.do("PUT", u, body)
}

func (c *Client) Patch(path string, body any) (json.RawMessage, error) {
	u := c.tradingURL(path, nil)
	return c.do("PATCH", u, body)
}

func (c *Client) Delete(path string, params url.Values) (json.RawMessage, error) {
	u := c.tradingURL(path, params)
	return c.do("DELETE", u, nil)
}

func (c *Client) GetData(path string, params url.Values) (json.RawMessage, error) {
	u := c.dataURL(path, params)
	return c.do("GET", u, nil)
}

func (c *Client) RawRequest(method, fullURL string, body any) (json.RawMessage, error) {
	return c.do(method, fullURL, body)
}

func (c *Client) do(method, url string, body any) (json.RawMessage, error) {
	var bodyReader io.Reader
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("encoding request body: %w", err)
		}
		bodyReader = bytes.NewReader(data)
	}

	req, err := http.NewRequest(method, url, bodyReader)
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
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response: %w", err)
	}

	if resp.StatusCode >= 400 {
		apiErr := &APIError{StatusCode: resp.StatusCode}
		if json.Unmarshal(respBody, apiErr) != nil || apiErr.Message == "" {
			apiErr.Message = fmt.Sprintf("HTTP %d: %s", resp.StatusCode, string(respBody))
		}
		return nil, apiErr
	}

	if len(respBody) == 0 {
		return json.RawMessage("null"), nil
	}

	return json.RawMessage(respBody), nil
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
