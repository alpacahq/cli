//go:build integration

package integration

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"os/exec"
	"strings"
	"testing"
	"time"
)

func TestEventsActivities_Historical(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	start := daysAgo(30)
	end := daysAgo(1)

	cmd := exec.CommandContext(ctx, cliBinary,
		"events", "activities",
		"--since", start,
		"--until", end,
		"--quiet",
	)
	cmd.Env = cliEnv()

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			t.Skipf("SSE endpoint timed out (may not be available on paper): %s", stderr.String())
		}
		if strings.Contains(stderr.String(), "403") || strings.Contains(stderr.String(), "404") {
			t.Skipf("SSE endpoint unavailable on paper: %s", stderr.String())
		}
		t.Fatalf("events activities failed: %v\nstderr: %s", err, stderr.String())
	}

	output := strings.TrimSpace(stdout.String())
	if output == "" {
		t.Skip("no historical events in the date range")
	}

	scanner := bufio.NewScanner(strings.NewReader(output))
	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		var event map[string]any
		if err := json.Unmarshal([]byte(line), &event); err != nil {
			t.Fatalf("line %d is not valid JSON: %q\nerror: %v", count+1, line, err)
		}
		count++
		if count == 1 {
			if _, ok := event["activity_type"]; !ok {
				t.Errorf("first event missing activity_type field: %v", event)
			}
		}
	}
	t.Logf("received %d events from %s to %s", count, start, end)
}
