package cmd

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/alpacahq/cli/internal/api"
	"github.com/alpacahq/cli/internal/cmdutil"
	"github.com/alpacahq/cli/internal/output"
	"github.com/spf13/cobra"
)

// cmdColumns allows per-command column overrides set via configure closures
// for commands that need custom columns (e.g. positions with P/L coloring).
var cmdColumns = map[*cobra.Command][]output.Column{}

// cmdJSON marks commands that should always render JSON regardless of format flags.
var cmdJSON = map[*cobra.Command]bool{}

// withJSON is a configure closure that marks a command as JSON-only output.
// Use for responses with complex nested or map-of-symbols structures where
// tabular rendering doesn't make sense.
func withJSON(c *cobra.Command) {
	cmdJSON[c] = true
}

// fetchCmd creates a command that fetches data and renders it.
// Slices render as tables; single objects render as key-value pairs.
// RequiredFlags from the op are enforced automatically.
//
// Op flags are auto-registered with nil FlagOpts. If a configure closure
// registers flags first (e.g. with custom FlagOpts), auto-registration
// is skipped.
func fetchCmd(use string, op api.Op, fetch func(cmd *cobra.Command, args []string) (any, error), configure ...func(*cobra.Command)) *cobra.Command {
	cmd := &cobra.Command{
		Use:   use,
		Short: op.Summary(),
	}
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		if req := registeredRequired(cmd, op); len(req) > 0 {
			if err := cmdutil.RequireAll(cmd, req...); err != nil {
				return err
			}
		}
		data, err := fetch(cmd, args)
		if err != nil {
			return err
		}
		if cmdJSON[cmd] {
			return output.JSON(cmd.OutOrStdout(), data)
		}
		cols := cmdColumns[cmd]
		if cols == nil {
			cols = columnsForOp(op)
		}
		w := cmd.OutOrStdout()
		format := getOutput()
		if isSliceResult(data) {
			return output.Render(w, format, cols, data)
		}
		return output.PrintSingle(w, format, cols, data)
	}
	for _, fn := range configure {
		fn(cmd)
	}
	autoRegisterFlags(cmd, op)
	return cmd
}

// actionCmd creates a command that performs a side effect and prints a message.
// If msg is empty, nothing is printed (the do func handles its own output).
func actionCmd(use string, op api.Op, msg string, do func(cmd *cobra.Command, args []string) error, configure ...func(*cobra.Command)) *cobra.Command {
	cmd := &cobra.Command{
		Use:   use,
		Short: op.Summary(),
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := do(cmd, args); err != nil {
				return err
			}
			if msg != "" {
				_, _ = fmt.Fprintln(cmd.OutOrStdout(), msg)
			}
			return nil
		},
	}
	for _, fn := range configure {
		fn(cmd)
	}
	autoRegisterFlags(cmd, op)
	return cmd
}

// registeredRequired returns only those RequiredFlags from the op that are
// actually registered on the command. Flags excluded via FlagOpts.Exclude
// (because they're positional args) are silently skipped.
func registeredRequired(cmd *cobra.Command, op api.Op) []string {
	all := op.RequiredFlags()
	if len(all) == 0 {
		return nil
	}
	var present []string
	for _, name := range all {
		if cmd.Flags().Lookup(name) != nil {
			present = append(present, name)
		}
	}
	return present
}

// autoRegisterFlags registers op flags with nil FlagOpts unless a configure
// closure already called RegisterFlags (detected via the "op" annotation that
// RegisterFlags always sets). This correctly handles the case where a configure
// closure excludes ALL flags — the annotation is still set.
func autoRegisterFlags(cmd *cobra.Command, op api.Op) {
	if cmd.Annotations != nil && cmd.Annotations["op"] != "" {
		return
	}
	flags := op.Flags()
	if len(flags) == 0 {
		return
	}
	cmdutil.RegisterFlags(cmd, flags, nil)
}

// isSliceResult reports whether data should be rendered as a list (table)
// rather than a single object (key-value).
func isSliceResult(data any) bool {
	switch v := data.(type) {
	case json.RawMessage:
		for _, b := range v {
			if b == ' ' || b == '\t' || b == '\n' || b == '\r' {
				continue
			}
			return b == '['
		}
		return false
	default:
		rv := reflect.ValueOf(data)
		for rv.Kind() == reflect.Ptr {
			rv = rv.Elem()
		}
		return rv.Kind() == reflect.Slice
	}
}
