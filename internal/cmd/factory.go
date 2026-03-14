package cmd

import (
	"encoding/json"

	"github.com/alpacahq/cli/internal/api"
	"github.com/alpacahq/cli/internal/cmdutil"
	"github.com/alpacahq/cli/internal/output"
	"github.com/spf13/cobra"
)

// Per-command overrides set via configure closures.
var (
	cmdJSON     = map[*cobra.Command]bool{}
	cmdFlagOpts = map[*cobra.Command]*cmdutil.FlagOpts{}
)

// jsonOnly marks a command as JSON-only output. Use for responses with complex
// nested or map-of-symbols structures where CSV rendering doesn't make sense.
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
		return output.Render(cmd.OutOrStdout(), getOutput(), data)
	}
	for _, fn := range configure {
		fn(cmd)
	}
	cmdutil.RegisterFlags(cmd, op.Flags(), cmdFlagOpts[cmd])
	return cmd
}

// voidResponse wraps a client call that may return empty data (204 no content),
// ensuring stdout is always valid JSON. Returns "{}" for empty responses.
func voidResponse(data json.RawMessage, err error) (any, error) {
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return json.RawMessage("{}"), nil
	}
	return data, nil
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
