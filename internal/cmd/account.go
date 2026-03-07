package cmd

import (
	"github.com/alpacahq/cli/internal/api"
	"github.com/alpacahq/cli/internal/cmdutil"
	"github.com/alpacahq/cli/internal/output"
	"github.com/spf13/cobra"
)

var accountCmd = &cobra.Command{
	Use:   "account",
	Short: "Manage your trading account",
}

var accountGetCmd = &cobra.Command{
	Use:   "get",
	Short: api.GetAccountOp.Summary(),
	Example: `  alpaca account get
  alpaca account get --json`,
	RunE: func(cmd *cobra.Command, args []string) error {
		account, err := tradingClient.GetAccount()
		if err != nil {
			return err
		}
		return output.PrintSingle(cmd.OutOrStdout(), getOutput(), columnsForOp(api.GetAccountOp), account)
	},
}

var accountConfigGetCmd = &cobra.Command{
	Use:   "get",
	Short: api.GetAccountConfigOp.Summary(),
	Example: `  alpaca account config get
  alpaca account config get --json`,
	RunE: func(cmd *cobra.Command, args []string) error {
		config, err := tradingClient.GetAccountConfig()
		if err != nil {
			return err
		}
		return output.PrintSingle(cmd.OutOrStdout(), getOutput(), columnsForOp(api.GetAccountConfigOp), config)
	},
}

var accountConfigSetCmd = &cobra.Command{
	Use:   "set",
	Short: api.PatchAccountConfigOp.Summary(),
	Example: `  alpaca account config set --no-shorting true
  alpaca account config set --dtbp-check entry`,
	RunE: func(cmd *cobra.Command, args []string) error {
		body, changed := accountConfigurationsBodyFromFlags(cmd)
		if !changed {
			return cmd.Help()
		}

		config, err := tradingClient.PatchAccountConfig(body)
		if err != nil {
			return err
		}
		return output.PrintSingle(cmd.OutOrStdout(), getOutput(), columnsForOp(api.PatchAccountConfigOp), config)
	},
}

var accountConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage account configuration",
}

func init() {
	cmdutil.RegisterFlags(accountConfigSetCmd, api.PatchAccountConfigOp.Flags(), nil)

	accountConfigCmd.AddCommand(accountConfigGetCmd)
	accountConfigCmd.AddCommand(accountConfigSetCmd)

	accountCmd.AddCommand(accountGetCmd)
	accountCmd.AddCommand(accountConfigCmd)
}
