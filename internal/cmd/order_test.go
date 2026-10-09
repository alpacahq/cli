package cmd

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"
)

func TestOrderRequestOmitsUnsetAdvancedInstructions(t *testing.T) {
	const orderID = "61e69015-8549-4baf-b96f-9c4f3e8d0c35"
	const instructions = `{"algorithm":"TWAP","destination":"NYSE"}`

	tests := []struct {
		name      string
		args      []string
		method    string
		path      string
		wantKey   bool
		wantEmpty bool
	}{
		{
			name:   "submit",
			args:   []string{"order", "submit", "--symbol", "AAPL", "--qty", "1", "--side", "buy", "--type", "market"},
			method: http.MethodPost,
			path:   "/v2/orders",
		},
		{
			name:   "replace",
			args:   []string{"order", "replace", "--order-id", orderID, "--qty", "2"},
			method: http.MethodPatch,
			path:   "/v2/orders/" + orderID,
		},
		{
			name:    "submit with flag",
			args:    []string{"order", "submit", "--symbol", "AAPL", "--qty", "1", "--side", "buy", "--type", "market", "--advanced-instructions", instructions},
			method:  http.MethodPost,
			path:    "/v2/orders",
			wantKey: true,
		},
		{
			name:      "submit with empty object",
			args:      []string{"order", "submit", "--symbol", "AAPL", "--qty", "1", "--side", "buy", "--type", "market", "--advanced-instructions", "{}"},
			method:    http.MethodPost,
			path:      "/v2/orders",
			wantKey:   true,
			wantEmpty: true,
		},
		{
			name:    "replace with flag",
			args:    []string{"order", "replace", "--order-id", orderID, "--qty", "2", "--advanced-instructions", instructions},
			method:  http.MethodPatch,
			path:    "/v2/orders/" + orderID,
			wantKey: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var gotMethod, gotPath string
			var gotBody []byte
			cleanup := setupMockClients(t, func(w http.ResponseWriter, r *http.Request) {
				gotMethod = r.Method
				gotPath = r.URL.Path
				var err error
				gotBody, err = io.ReadAll(r.Body)
				if err != nil {
					t.Errorf("read body: %v", err)
				}
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				_, _ = w.Write([]byte(`{"id":"abc","symbol":"AAPL","status":"accepted"}`))
			})
			defer cleanup()

			root := Root()
			root.SetOut(new(bytes.Buffer))
			root.SetErr(new(bytes.Buffer))
			root.SetArgs(tt.args)
			if err := root.Execute(); err != nil {
				t.Fatalf("execute: %v", err)
			}
			if gotMethod != tt.method || gotPath != tt.path {
				t.Fatalf("request = %s %s, want %s %s", gotMethod, gotPath, tt.method, tt.path)
			}

			var sent map[string]any
			if err := json.Unmarshal(gotBody, &sent); err != nil {
				t.Fatalf("request body %s: %v", gotBody, err)
			}
			raw, present := sent["advanced_instructions"]
			if present != tt.wantKey {
				t.Fatalf("advanced_instructions present = %v, body %s", present, gotBody)
			}
			if !tt.wantKey {
				return
			}
			info, ok := raw.(map[string]any)
			if !ok {
				t.Fatalf("advanced_instructions = %#v", raw)
			}
			if tt.wantEmpty {
				if len(info) != 0 {
					t.Fatalf("advanced_instructions = %#v", info)
				}
				return
			}
			if info["algorithm"] != "TWAP" || info["destination"] != "NYSE" {
				t.Fatalf("advanced_instructions = %#v", raw)
			}
		})
	}
}
