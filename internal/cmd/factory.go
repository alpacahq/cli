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
		if req := op.RequiredFlags(); len(req) > 0 {
			if err := cmdutil.RequireAll(cmd, req...); err != nil {
				return err
			}
		}
		data, err := fetch(cmd, args)
		if err != nil {
			return err
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

// autoRegisterFlags registers op flags with nil FlagOpts unless a configure
// closure already registered them (detected by checking for any known flag).
func autoRegisterFlags(cmd *cobra.Command, op api.Op) {
	flags := op.Flags()
	if len(flags) == 0 {
		return
	}
	for _, f := range flags {
		if cmd.Flags().Lookup(f.Name) != nil {
			return // configure already handled registration
		}
	}
	cmdutil.RegisterFlags(cmd, flags, nil)
}

// jsonCmd creates a command that fetches data and always renders it as JSON.
// Use for responses with complex nested or map-of-symbols structures where
// tabular rendering doesn't make sense.
func jsonCmd(use string, op api.Op, fetch func(cmd *cobra.Command, args []string) (any, error), configure ...func(*cobra.Command)) *cobra.Command {
	cmd := &cobra.Command{
		Use:   use,
		Short: op.Summary(),
	}
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		if req := op.RequiredFlags(); len(req) > 0 {
			if err := cmdutil.RequireAll(cmd, req...); err != nil {
				return err
			}
		}
		data, err := fetch(cmd, args)
		if err != nil {
			return err
		}
		return output.JSON(cmd.OutOrStdout(), data)
	}
	for _, fn := range configure {
		fn(cmd)
	}
	autoRegisterFlags(cmd, op)
	return cmd
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
