package cmd

import (
	"fmt"

	"github.com/alpacahq/cli/internal/api"
	"github.com/alpacahq/cli/internal/cmdutil"
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

var accountConfigSetCmd = fetchCmd("set", api.PatchAccountConfigOp, func(cmd *cobra.Command, args []string) (any, error) {
	body, changed := accountConfigurationsBodyFromFlags(cmd)
	if !changed {
		return nil, fmt.Errorf("specify at least one flag to change (see '%s --help')", cmd.CommandPath())
	}
	return tradingClient.PatchAccountConfig(body)
}, func(c *cobra.Command) {
	c.Example = `  alpaca account config set --no-shorting true
  alpaca account config set --dtbp-check entry`
})

var accountConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage account configuration",
}

// --- Activity ---

var activityCmd = &cobra.Command{
	Use:   "activity",
	Short: "Account activities (fills, dividends, transfers, etc.)",
}

var activityListCmd = fetchCmd("list", api.GetAccountActivitiesOp, func(cmd *cobra.Command, args []string) (any, error) {
	actType := cmdutil.Str(cmd, "activity-types")
	if actType != "" {
		return tradingClient.GetAccountActivitiesByActivityType(actType, getAccountActivitiesByActivityTypeParamsFromFlags(cmd))
	}
	return tradingClient.GetAccountActivities(getAccountActivitiesParamsFromFlags(cmd))
}, func(c *cobra.Command) {
	c.Example = `  alpaca account activity list
  alpaca account activity list --activity-types FILL --page-size 20
  alpaca account activity list --activity-types DIV --after 2025-01-01
  alpaca account activity list --activity-types FILL,TRANS --direction desc`
})

// --- Portfolio ---

var portfolioCmd = fetchCmd("portfolio", api.GetAccountPortfolioHistoryOp, func(cmd *cobra.Command, args []string) (any, error) {
	return tradingClient.GetAccountPortfolioHistory(getAccountPortfolioHistoryParamsFromFlags(cmd))
}, withJSON, func(c *cobra.Command) {
	cmdutil.RegisterFlags(c, api.GetAccountPortfolioHistoryOp.Flags(), &cmdutil.FlagOpts{
		Exclude: map[string]bool{"extended_hours": true},
	})
	c.Long = "Returns portfolio equity and P&L history. Output is always JSON due to complex time-series structure."
	c.Example = `  alpaca account portfolio
  alpaca account portfolio --period 1M --timeframe 1D`
})

func init() {
	accountConfigCmd.AddCommand(accountConfigGetCmd)
	accountConfigCmd.AddCommand(accountConfigSetCmd)

	activityCmd.AddCommand(activityListCmd)

	accountCmd.AddCommand(accountGetCmd)
	accountCmd.AddCommand(accountConfigCmd)
	accountCmd.AddCommand(activityCmd)
	accountCmd.AddCommand(portfolioCmd)
}
