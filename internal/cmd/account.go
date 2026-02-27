package cmd

import (
	"github.com/alpacahq/cli/internal/output"
	"github.com/spf13/cobra"
)

var accountCmd = &cobra.Command{
	Use:   "account",
	Short: "Manage your trading account",
}

var accountGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Show account details",
	RunE: func(cmd *cobra.Command, args []string) error {
		data, err := apiClient.Get("/v2/account", nil)
		if err != nil {
			return err
		}

		columns := []output.Column{
			{Header: "ACCOUNT #", Field: "account_number"},
			{Header: "STATUS", Field: "status"},
			{Header: "EQUITY", Field: "equity"},
			{Header: "CASH", Field: "cash"},
			{Header: "BUYING POWER", Field: "buying_power"},
			{Header: "PORTFOLIO VALUE", Field: "portfolio_value"},
			{Header: "CURRENCY", Field: "currency"},
			{Header: "PDT", Field: "pattern_day_trader"},
			{Header: "TRADING BLOCKED", Field: "trading_blocked"},
		}

		return output.PrintSingle(getOutput(), columns, data)
	},
}

var accountConfigGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Show account configuration",
	RunE: func(cmd *cobra.Command, args []string) error {
		data, err := apiClient.Get("/v2/account/configurations", nil)
		if err != nil {
			return err
		}

		columns := []output.Column{
			{Header: "DTBP CHECK", Field: "dtbp_check"},
			{Header: "FRACTIONAL TRADING", Field: "fractional_trading"},
			{Header: "MAX MARGIN MULTIPLIER", Field: "max_margin_multiplier"},
			{Header: "NO SHORTING", Field: "no_shorting"},
			{Header: "PDT CHECK", Field: "pdt_check"},
			{Header: "SUSPEND TRADE", Field: "suspend_trade"},
			{Header: "TRADE CONFIRM EMAIL", Field: "trade_confirm_email"},
		}

		return output.PrintSingle(getOutput(), columns, data)
	},
}

var accountConfigSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Update account configuration",
	Example: `  alpaca account config set --no-shorting true
  alpaca account config set --dtbp-check entry`,
	RunE: func(cmd *cobra.Command, args []string) error {
		body := map[string]any{}

		if cmd.Flags().Changed("dtbp-check") {
			v, _ := cmd.Flags().GetString("dtbp-check")
			body["dtbp_check"] = v
		}
		if cmd.Flags().Changed("no-shorting") {
			v, _ := cmd.Flags().GetBool("no-shorting")
			body["no_shorting"] = v
		}
		if cmd.Flags().Changed("pdt-check") {
			v, _ := cmd.Flags().GetString("pdt-check")
			body["pdt_check"] = v
		}
		if cmd.Flags().Changed("fractional-trading") {
			v, _ := cmd.Flags().GetBool("fractional-trading")
			body["fractional_trading"] = v
		}
		if cmd.Flags().Changed("suspend-trade") {
			v, _ := cmd.Flags().GetBool("suspend-trade")
			body["suspend_trade"] = v
		}
		if cmd.Flags().Changed("trade-confirm-email") {
			v, _ := cmd.Flags().GetString("trade-confirm-email")
			body["trade_confirm_email"] = v
		}
		if cmd.Flags().Changed("max-margin-multiplier") {
			v, _ := cmd.Flags().GetString("max-margin-multiplier")
			body["max_margin_multiplier"] = v
		}

		if len(body) == 0 {
			return cmd.Help()
		}

		data, err := apiClient.Patch("/v2/account/configurations", body)
		if err != nil {
			return err
		}

		return output.JSON(cmd.OutOrStdout(), data)
	},
}

var accountConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage account configuration",
}

func init() {
	accountConfigSetCmd.Flags().String("dtbp-check", "", "Day trading buying power check: entry or exit")
	accountConfigSetCmd.Flags().Bool("no-shorting", false, "Disable short selling")
	accountConfigSetCmd.Flags().String("pdt-check", "", "PDT check: entry or exit")
	accountConfigSetCmd.Flags().Bool("fractional-trading", false, "Enable fractional trading")
	accountConfigSetCmd.Flags().Bool("suspend-trade", false, "Suspend trading")
	accountConfigSetCmd.Flags().String("trade-confirm-email", "", "Trade confirm email: all or none")
	accountConfigSetCmd.Flags().String("max-margin-multiplier", "", "Max margin multiplier: 1 or 4")

	accountConfigCmd.AddCommand(accountConfigGetCmd)
	accountConfigCmd.AddCommand(accountConfigSetCmd)

	accountCmd.AddCommand(accountGetCmd)
	accountCmd.AddCommand(accountConfigCmd)
}
