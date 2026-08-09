package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/alpacahq/cli/internal/api"
	"github.com/alpacahq/cli/internal/cmdutil"
	"github.com/spf13/cobra"
)

// fetchCmd creates a command that fetches data and renders it.
// All OAS flags (path, query, header, body) are auto-registered from the Op.
// Defaults are baked into FlagDef.Default at generation time.
func fetchCmd(use string, op api.Op, fetch func(cmd *cobra.Command, args []string) (any, error), configure ...func(*cobra.Command)) *cobra.Command {
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
		data, err := fetch(cmd, args)
		if err != nil {
			return err
		}
		return renderData(cmd, data)
	}
	for _, fn := range configure {
		fn(cmd)
	}
	cmdutil.RegisterFlags(cmd, op.Flags, op.Name, nil)
	return cmd
}

// attachCmd mirrors fetchCmd but operates on an existing command. Used for
// self: true commands where a parent group is also directly runnable.
func attachCmd(cmd *cobra.Command, op api.Op, fetch func(cmd *cobra.Command, args []string) (any, error), configure ...func(*cobra.Command)) {
	if op.Long != "" {
		cmd.Long = op.Long
	}
	if op.Example != "" {
		cmd.Example = op.Example
	}
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		if err := requireFlags(cmd, op); err != nil {
			return err
		}
		data, err := fetch(cmd, args)
		if err != nil {
			return err
		}
		return renderData(cmd, data)
	}
	for _, fn := range configure {
		fn(cmd)
	}
	cmdutil.RegisterFlags(cmd, op.Flags, op.Name, nil)
}

// queryFromFlags builds url.Values from cobra flags using FlagDef metadata,
// replacing per-endpoint *ParamsFromFlags boilerplate with a single runtime helper.
func queryFromFlags(cmd *cobra.Command, op api.Op) url.Values {
	v := url.Values{}
	for _, f := range op.Flags {
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

// headersFromFlags builds request headers from explicitly set OAS header flags.
func headersFromFlags(cmd *cobra.Command, op api.Op) http.Header {
	headers := http.Header{}
	for _, f := range op.Flags {
		if f.Source != "header" || !cmd.Flags().Changed(f.Name) {
			continue
		}
		switch f.Type {
		case "string":
			headers.Set(f.OASName, cmdutil.Str(cmd, f.Name))
		case "int":
			headers.Set(f.OASName, fmt.Sprint(cmdutil.Int(cmd, f.Name)))
		case "bool":
			headers.Set(f.OASName, fmt.Sprint(cmdutil.Bool(cmd, f.Name)))
		}
	}
	return headers
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

// requireFlags validates that all OAS-required flags have been provided.
func requireFlags(cmd *cobra.Command, op api.Op) error {
	var req []string
	for _, f := range op.Flags {
		if f.Required {
			req = append(req, f.Name)
		}
	}
	if len(req) == 0 {
		return nil
	}
	return cmdutil.RequireAll(cmd, req...)
}
