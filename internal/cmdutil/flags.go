package cmdutil

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/alpacahq/cli/internal/api"
	"github.com/spf13/cobra"
)

// FlagOpts configures RegisterFlags behavior. All map keys use OAS names.
type FlagOpts struct {
	Exclude  map[string]bool   // OAS param names to skip
	Aliases  map[string]string // OAS name → CLI flag name override
	Defaults map[string]string // OAS name → default value override
}

// RegisterFlags registers CLI flags from generated FlagDef definitions.
func RegisterFlags(cmd *cobra.Command, defs []api.FlagDef, opts *FlagOpts) {
	for _, d := range defs {
		if opts != nil && opts.Exclude[d.OASName] {
			continue
		}

		name := d.Name
		defaultVal := d.Default
		if opts != nil {
			if alias, ok := opts.Aliases[d.OASName]; ok {
				name = alias
			}
			if def, ok := opts.Defaults[d.OASName]; ok {
				defaultVal = def
			}
		}

		switch d.Type {
		case "bool":
			cmd.Flags().Bool(name, defaultVal == "true", d.Description)
		case "int":
			defInt, _ := strconv.Atoi(defaultVal)
			cmd.Flags().Int(name, defInt, d.Description)
		default:
			cmd.Flags().String(name, defaultVal, d.Description)
		}

		if len(d.Completions) > 0 {
			_ = cmd.RegisterFlagCompletionFunc(name, cobra.FixedCompletions(d.Completions, cobra.ShellCompDirectiveNoFileComp))
		}
	}
}

func Str(cmd *cobra.Command, name string) string {
	v, _ := cmd.Flags().GetString(name)
	return v
}

func RequireStr(cmd *cobra.Command, name string) (string, error) {
	v := Str(cmd, name)
	if v == "" {
		return "", fmt.Errorf("--%s is required", name)
	}
	return v, nil
}

func RequireAll(cmd *cobra.Command, names ...string) error {
	var missing []string
	for _, n := range names {
		if Str(cmd, n) == "" {
			missing = append(missing, "--"+n)
		}
	}
	if len(missing) > 0 {
		return fmt.Errorf("%s required", strings.Join(missing, ", "))
	}
	return nil
}

func Bool(cmd *cobra.Command, name string) bool {
	v, _ := cmd.Flags().GetBool(name)
	return v
}

func Int(cmd *cobra.Command, name string) int {
	v, _ := cmd.Flags().GetInt(name)
	return v
}

func Changed(cmd *cobra.Command, name string) bool {
	return cmd.Flags().Changed(name)
}

func AddSortFlag(cmd *cobra.Command, completions []string) {
	cmd.Flags().String("sort", "", "Sort order: asc or desc")
	_ = cmd.RegisterFlagCompletionFunc("sort",
		cobra.FixedCompletions(completions, cobra.ShellCompDirectiveNoFileComp))
}

func AddDateRangeFlags(cmd *cobra.Command) {
	cmd.Flags().String("start", "", "Start date (YYYY-MM-DD or RFC3339)")
	cmd.Flags().String("end", "", "End date (YYYY-MM-DD or RFC3339)")
}

func AddLimitFlag(cmd *cobra.Command) {
	cmd.Flags().Int("limit", 0, "Max number of results")
}
