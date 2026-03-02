package cmdutil

import (
	"fmt"

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
