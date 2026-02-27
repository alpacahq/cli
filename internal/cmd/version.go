package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of alpaca CLI",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("alpaca version %s\n", version)
	},
}
