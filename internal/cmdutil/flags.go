package cmdutil

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

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

func Float64(cmd *cobra.Command, name string) float64 {
	v, _ := cmd.Flags().GetFloat64(name)
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
