package cmd

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/signal"
	"strings"

	"github.com/alpacahq/cli/internal/api"
	"github.com/alpacahq/cli/internal/cmdutil"
	"github.com/alpacahq/cli/internal/output"
	"github.com/spf13/cobra"
)

// streamCmd creates a command that opens an SSE stream and writes each event
// as a JSON line to stdout. It is the streaming counterpart to fetchCmd.
func streamCmd(use string, op api.Op, stream func(cmd *cobra.Command, args []string) (io.ReadCloser, error), configure ...func(*cobra.Command)) *cobra.Command {
	cmd := &cobra.Command{
		Use:     use,
		Short:   op.Summary,
		Long:    op.Long,
		Example: op.Example,
	}
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		if err := requireFlags(cmd, op); err != nil {
			return err
		}
		body, err := stream(cmd, args)
		if err != nil {
			return err
		}
		defer func() { _ = body.Close() }()

		ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
		defer stop()

		if !quietFlag {
			fmt.Fprintln(os.Stderr, "Streaming events... (Ctrl+C to stop)")
		}
		return streamSSE(ctx, cmd, bufio.NewScanner(body))
	}
	for _, fn := range configure {
		fn(cmd)
	}
	cmdutil.RegisterFlags(cmd, op.Flags, op.Name, nil)
	return cmd
}

// streamSSE reads SSE frames from the scanner, extracts data payloads,
// and writes each as a JSON line to stdout. Supports --jq per-event.
func streamSSE(ctx context.Context, cmd *cobra.Command, scanner *bufio.Scanner) error {
	var dataBuf strings.Builder
	w := cmd.OutOrStdout()

	for {
		select {
		case <-ctx.Done():
			return nil
		default:
		}

		if !scanner.Scan() {
			break
		}
		line := scanner.Text()

		if strings.HasPrefix(line, "data: ") {
			if dataBuf.Len() > 0 {
				dataBuf.WriteByte('\n')
			}
			dataBuf.WriteString(line[6:])
			continue
		}

		// Blank line = end of SSE frame; flush accumulated data.
		if line == "" && dataBuf.Len() > 0 {
			raw := json.RawMessage(dataBuf.String())
			dataBuf.Reset()

			if jqFlag != "" {
				filtered, err := output.ApplyJQ(raw, jqFlag)
				if err != nil {
					return err
				}
				if skipJQResult(filtered) {
					continue
				}
				if raw, err = json.Marshal(filtered); err != nil {
					return err
				}
			}

			if _, err := fmt.Fprintf(w, "%s\n", raw); err != nil {
				return nil
			}
		}
	}

	if err := scanner.Err(); err != nil {
		select {
		case <-ctx.Done():
			return nil
		default:
			return fmt.Errorf("reading event stream: %w", err)
		}
	}
	return nil
}

// skipJQResult returns true for jq outputs that should be suppressed:
// nil (no match), false, and Go values that marshal to JSON null.
func skipJQResult(v any) bool {
	if v == nil {
		return true
	}
	if b, ok := v.(bool); ok && !b {
		return true
	}
	raw, err := json.Marshal(v)
	return err == nil && string(raw) == "null"
}
