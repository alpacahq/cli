package cmd

import (
	"bufio"
	"bytes"
	"context"
	"net/http"
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

func runSSE(t *testing.T, input string) string {
	t.Helper()
	cmd := &cobra.Command{}
	var buf bytes.Buffer
	cmd.SetOut(&buf)
	if err := streamSSE(context.Background(), cmd, bufio.NewScanner(strings.NewReader(input))); err != nil {
		t.Fatalf("streamSSE: %v", err)
	}
	return buf.String()
}

func TestStreamSSE(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantSub []string
	}{
		{"two events", "data: {\"id\":\"1\"}\n\ndata: {\"id\":\"2\"}\n\n", []string{`"id":"1"`, `"id":"2"`}},
		{"multi-line data", "data: {\"a\":1,\ndata: \"b\":2}\n\n", []string{`"a":1`, `"b":2`}},
		{"ignores non-data fields", "event: x\nid: 42\ndata: {\"ok\":true}\n\n", []string{`"ok":true`}},
		{"empty stream", "", nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := runSSE(t, tt.input)
			for _, sub := range tt.wantSub {
				if !strings.Contains(out, sub) {
					t.Errorf("output missing %q: %s", sub, out)
				}
			}
			if tt.wantSub == nil && out != "" {
				t.Errorf("expected empty output, got %q", out)
			}
		})
	}
}

func TestStreamSSE_ContextCancelled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	cmd := &cobra.Command{}
	cmd.SetOut(&bytes.Buffer{})
	err := streamSSE(ctx, cmd, bufio.NewScanner(strings.NewReader("data: {}\n\n")))
	if err != nil {
		t.Fatalf("expected nil error on canceled ctx, got: %v", err)
	}
}

func TestStreamSSE_JQFilter(t *testing.T) {
	oldJQ := jqFlag
	jqFlag = `select(.type == "FILL")`
	defer func() { jqFlag = oldJQ }()

	input := "data: {\"id\":\"1\",\"type\":\"FILL\"}\n\ndata: {\"id\":\"2\",\"type\":\"DIV\"}\n\ndata: {\"id\":\"3\",\"type\":\"FILL\"}\n\n"
	out := runSSE(t, input)

	lines := strings.Split(strings.TrimSpace(out), "\n")
	if len(lines) != 2 {
		t.Fatalf("expected 2 lines after filter, got %d: %q", len(lines), out)
	}
	if !strings.Contains(lines[0], `"id":"1"`) || !strings.Contains(lines[1], `"id":"3"`) {
		t.Errorf("wrong events: %s", out)
	}
}

func TestStreamCmd_E2E(t *testing.T) {
	cleanup := setupMockClients(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		_, _ = w.Write([]byte("data: {\"id\":\"ev1\"}\n\ndata: {\"id\":\"ev2\"}\n\n"))
	})
	defer cleanup()

	root := Root()
	var stdout bytes.Buffer
	root.SetOut(&stdout)
	root.SetErr(&bytes.Buffer{})
	root.SetArgs([]string{"events", "activities"})

	if err := root.Execute(); err != nil {
		t.Fatalf("command failed: %v", err)
	}

	lines := strings.Split(strings.TrimSpace(stdout.String()), "\n")
	if len(lines) != 2 {
		t.Fatalf("expected 2 events, got %d: %q", len(lines), stdout.String())
	}
}
