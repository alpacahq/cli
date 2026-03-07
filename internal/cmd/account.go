package cmd

import (
	"encoding/json"

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

// --- Activity ---

var activityCmd = &cobra.Command{
	Use:   "activity",
	Short: "Account activities (fills, dividends, transfers, etc.)",
}

var activityListCmd = &cobra.Command{
	Use:   "list",
	Short: api.GetAccountActivitiesOp.Summary(),
	Example: `  alpaca account activity list
  alpaca account activity list --activity-types FILL --page-size 20
  alpaca account activity list --activity-types DIV --after 2025-01-01
  alpaca account activity list --activity-types FILL,TRANS --direction desc`,
	RunE: func(cmd *cobra.Command, args []string) error {
		actType := cmdutil.Str(cmd, "activity-types")

		var data json.RawMessage
		var err error

		if actType != "" {
			data, err = tradingClient.GetAccountActivitiesByActivityType(actType, getAccountActivitiesByActivityTypeParamsFromFlags(cmd))
		} else {
			data, err = tradingClient.GetAccountActivities(getAccountActivitiesParamsFromFlags(cmd))
		}
		if err != nil {
			return err
		}

		out := getOutput()

		var items []map[string]any
		if err := json.Unmarshal(data, &items); err != nil {
			return output.JSON(cmd.OutOrStdout(), data)
		}

		return output.Render(cmd.OutOrStdout(), out, nil, data)
	},
}

// --- Portfolio ---

var portfolioCmd = &cobra.Command{
	Use:   "portfolio",
	Short: api.GetAccountPortfolioHistoryOp.Summary(),
	Long:  "Returns portfolio equity and P&L history. Output is always JSON due to complex time-series structure.",
	Example: `  alpaca account portfolio
  alpaca account portfolio --period 1M --timeframe 1D`,
	RunE: func(cmd *cobra.Command, args []string) error {
		params := getAccountPortfolioHistoryParamsFromFlags(cmd)

		history, err := tradingClient.GetAccountPortfolioHistory(params)
		if err != nil {
			return err
		}
		return output.JSON(cmd.OutOrStdout(), history)
	},
}

func init() {
	cmdutil.RegisterFlags(accountConfigSetCmd, api.PatchAccountConfigOp.Flags(), nil)

	accountConfigCmd.AddCommand(accountConfigGetCmd)
	accountConfigCmd.AddCommand(accountConfigSetCmd)

	cmdutil.RegisterFlags(activityListCmd, api.GetAccountActivitiesOp.Flags(), nil)
	activityCmd.AddCommand(activityListCmd)

	cmdutil.RegisterFlags(portfolioCmd, api.GetAccountPortfolioHistoryOp.Flags(), &cmdutil.FlagOpts{
		Exclude: map[string]bool{"extended_hours": true},
	})

	accountCmd.AddCommand(accountGetCmd)
	accountCmd.AddCommand(accountConfigCmd)
	accountCmd.AddCommand(activityCmd)
	accountCmd.AddCommand(portfolioCmd)
}
