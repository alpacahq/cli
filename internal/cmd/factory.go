package cmd

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/alpacahq/cli/internal/api"
	"github.com/alpacahq/cli/internal/cmdutil"
	"github.com/spf13/cobra"
)

// Per-command overrides set via configure closures.
var cmdFlagOpts = map[*cobra.Command]*cmdutil.FlagOpts{}

// flagOpts sets custom FlagOpts for OAS flag registration. Use to override
// defaults shown in --help.
func flagOpts(opts *cmdutil.FlagOpts) func(*cobra.Command) {
	return func(c *cobra.Command) {
		cmdFlagOpts[c] = opts
	}
}

// fetchCmd creates a command that fetches data and renders it.
// All OAS flags (path, query, body) are auto-registered. Path params become
// required flags enforced via RequiredFlags(). Configure closures should NOT
// call RegisterFlags for OAS flags — use flagOpts instead.
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
		return renderData(cmd.OutOrStdout(), data)
	}
	for _, fn := range configure {
		fn(cmd)
	}
	cmdutil.RegisterFlags(cmd, op.Flags(), cmdFlagOpts[cmd])
	return cmd
}

// attachCmd mirrors fetchCmd but operates on an existing command. Used for
// self: true commands where a parent group is also directly runnable.
func attachCmd(cmd *cobra.Command, op api.Op, fetch func(cmd *cobra.Command, args []string) (any, error), configure ...func(*cobra.Command)) {
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
		return renderData(cmd.OutOrStdout(), data)
	}
	for _, fn := range configure {
		fn(cmd)
	}
	cmdutil.RegisterFlags(cmd, op.Flags(), cmdFlagOpts[cmd])
}

// queryFromFlags builds url.Values from cobra flags using FlagDef metadata,
// replacing per-endpoint *ParamsFromFlags boilerplate with a single runtime helper.
func queryFromFlags(cmd *cobra.Command, op api.Op) url.Values {
	v := url.Values{}
	for _, f := range op.Flags() {
		if f.Source != "query" || !cmd.Flags().Changed(f.Name) {
			continue
		}
		switch f.Type {
		case "string":
			v.Set(f.OASName, cmdutil.Str(cmd, f.Name))
		case "int":
			v.Set(f.OASName, fmt.Sprint(cmdutil.Int(cmd, f.Name)))
		case "bool":
			if cmdutil.Bool(cmd, f.Name) {
				v.Set(f.OASName, "true")
			}
		}
	}
	return v
}

// normalizePathParam strips slashes from path param values (e.g. BTC/USD → BTCUSD).
func normalizePathParam(s string) string {
	return strings.ReplaceAll(s, "/", "")
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
// actually registered on the command.
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
