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

// Per-command overrides set via configure closures.
var (
	cmdColumns  = map[*cobra.Command][]output.Column{}
	cmdJSON     = map[*cobra.Command]bool{}
	cmdFlagOpts = map[*cobra.Command]*cmdutil.FlagOpts{}
)

// jsonOnly marks a command as JSON-only output. Use for responses with complex
// nested or map-of-symbols structures where tabular rendering doesn't make sense.
func jsonOnly(c *cobra.Command) {
	cmdJSON[c] = true
}

// flagOpts sets custom FlagOpts for OAS flag registration. Use to exclude
// params that are positional args, or to override defaults shown in --help.
func flagOpts(opts *cmdutil.FlagOpts) func(*cobra.Command) {
	return func(c *cobra.Command) {
		cmdFlagOpts[c] = opts
	}
}

// fetchCmd creates a command that fetches data and renders it.
// Slices render as tables; single objects render as key-value pairs.
// RequiredFlags from the op are enforced for flags actually registered
// on the command (excluded positional-arg params are skipped).
//
// OAS flags are registered after configure closures run, using FlagOpts
// from flagOpts (if any). Configure closures should NOT call RegisterFlags
// for OAS flags — use flagOpts instead.
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
	cmdutil.RegisterFlags(cmd, op.Flags(), cmdFlagOpts[cmd])
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
	cmdutil.RegisterFlags(cmd, op.Flags(), cmdFlagOpts[cmd])
	return cmd
}

// registeredRequired returns only those RequiredFlags from the op that are
// actually registered on the command. Params promoted to positional args
// (excluded via flagOpts) are silently skipped.
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
