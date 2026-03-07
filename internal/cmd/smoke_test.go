package cmd

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/alpacahq/cli/internal/api"
	"github.com/alpacahq/cli/internal/client"
	"github.com/alpacahq/cli/internal/config"
)

func setupMockClients(t *testing.T, handler http.HandlerFunc) func() {
	t.Helper()
	srv := httptest.NewServer(handler)
	c := &client.Client{
		HTTP:      &http.Client{Timeout: 5 * time.Second},
		BaseURL:   srv.URL,
		DataURL:   srv.URL,
		APIKey:    "PKTEST",
		Secret:    "SKTEST",
		UserAgent: "alpaca-cli/test",
	}
	oldTrading := tradingClient
	oldData := dataClient
	oldCfg := cfg
	tradingClient = api.NewTradingClient(c)
	dataClient = api.NewMarketDataClient(c)
	cfg = &config.Resolved{Output: "json"}
	oldBaseURL := os.Getenv("ALPACA_BASE_URL")
	oldDataURL := os.Getenv("ALPACA_DATA_URL")
	oldAPIKey := os.Getenv("ALPACA_API_KEY")
	oldSecret := os.Getenv("ALPACA_SECRET_KEY")
	os.Setenv("ALPACA_BASE_URL", srv.URL)
	os.Setenv("ALPACA_DATA_URL", srv.URL)
	os.Setenv("ALPACA_API_KEY", "PKTEST")
	os.Setenv("ALPACA_SECRET_KEY", "SKTEST")
	return func() {
		tradingClient = oldTrading
		dataClient = oldData
		cfg = oldCfg
		srv.Close()
		os.Setenv("ALPACA_BASE_URL", oldBaseURL)
		os.Setenv("ALPACA_DATA_URL", oldDataURL)
		os.Setenv("ALPACA_API_KEY", oldAPIKey)
		os.Setenv("ALPACA_SECRET_KEY", oldSecret)
	}
}

func TestCommandSmoke(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		response string
		status   int
	}{
		{
			name:     "account get",
			args:     []string{"account", "get"},
			response: `{"id":"abc","account_number":"123","status":"ACTIVE","equity":"10000","cash":"5000","buying_power":"20000","portfolio_value":"10000","currency":"USD","pattern_day_trader":false,"trading_blocked":false}`,
			status:   200,
		},
		{
			name:     "position list",
			args:     []string{"position", "list"},
			response: `[]`,
			status:   200,
		},
		{
			name:     "order list",
			args:     []string{"order", "list"},
			response: `[]`,
			status:   200,
		},
		{
			name:     "account activity list trades",
			args:     []string{"account", "activity", "list", "--activity-types", "FILL"},
			response: `[]`,
			status:   200,
		},
		{
			name:     "clock",
			args:     []string{"clock"},
			response: `{"timestamp":"2025-01-15T14:30:00-05:00","is_open":true,"next_open":"2025-01-16T09:30:00-05:00","next_close":"2025-01-15T16:00:00-05:00"}`,
			status:   200,
		},
		{
			name:     "calendar",
			args:     []string{"calendar"},
			response: `[{"date":"2025-01-15","open":"09:30","close":"16:00","session_open":"04:00","session_close":"20:00"}]`,
			status:   200,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cleanup := setupMockClients(t, func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tt.status)
				_, _ = w.Write([]byte(tt.response))
			})
			defer cleanup()

			root := Root()
			buf := new(bytes.Buffer)
			root.SetOut(buf)
			root.SetErr(new(bytes.Buffer))
			root.SetArgs(append([]string{"--json"}, tt.args...))

			err := root.Execute()
			if err != nil {
				t.Errorf("command %v failed: %v", tt.args, err)
			}
			if buf.Len() == 0 {
				t.Errorf("command %v produced no output", tt.args)
			}
		})
	}
}
