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

var accountGetCmd = fetchCmd("get", api.GetAccountOp, func(cmd *cobra.Command, args []string) (any, error) {
	return tradingClient.GetAccount()
}, func(c *cobra.Command) {
	c.Example = `  alpaca account get
  alpaca account get --json`
})

var accountConfigGetCmd = fetchCmd("get", api.GetAccountConfigOp, func(cmd *cobra.Command, args []string) (any, error) {
	return tradingClient.GetAccountConfig()
}, func(c *cobra.Command) {
	c.Example = `  alpaca account config get
  alpaca account config get --json`
})

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
	accountCmd.AddCommand(activityCmd)
	accountCmd.AddCommand(portfolioCmd)
}
