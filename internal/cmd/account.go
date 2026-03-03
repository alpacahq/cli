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
	Short: api.GetAccountOp.Summary,
	RunE: func(cmd *cobra.Command, args []string) error {
		account, err := tradingClient.GetAccount()
		if err != nil {
			return err
		}
		return output.PrintSingle(cmd.OutOrStdout(), getOutput(), accountColumns(), account)
	},
}

var accountConfigGetCmd = &cobra.Command{
	Use:   "get",
	Short: api.GetAccountConfigOp.Summary,
	RunE: func(cmd *cobra.Command, args []string) error {
		config, err := tradingClient.GetAccountConfig()
		if err != nil {
			return err
		}
		return output.PrintSingle(cmd.OutOrStdout(), getOutput(), accountConfigColumns(), config)
	},
}

var accountConfigSetCmd = &cobra.Command{
	Use:   "set",
	Short: api.PatchAccountConfigOp.Summary,
	Example: `  alpaca account config set --no-shorting true
  alpaca account config set --dtbp-check entry`,
	RunE: func(cmd *cobra.Command, args []string) error {
		body := &api.AccountConfigurations{}
		anyChanged := false

		if cmdutil.Changed(cmd, "dtbp-check") {
			body.DTBPCheck = cmdutil.Str(cmd, "dtbp-check")
			anyChanged = true
		}
		if cmdutil.Changed(cmd, "no-shorting") {
			body.NoShorting = cmdutil.Bool(cmd, "no-shorting")
			anyChanged = true
		}
		if cmdutil.Changed(cmd, "pdt-check") {
			body.PDTCheck = cmdutil.Str(cmd, "pdt-check")
			anyChanged = true
		}
		if cmdutil.Changed(cmd, "fractional-trading") {
			body.FractionalTrading = cmdutil.Bool(cmd, "fractional-trading")
			anyChanged = true
		}
		if cmdutil.Changed(cmd, "suspend-trade") {
			body.SuspendTrade = cmdutil.Bool(cmd, "suspend-trade")
			anyChanged = true
		}
		if cmdutil.Changed(cmd, "trade-confirm-email") {
			body.TradeConfirmEmail = cmdutil.Str(cmd, "trade-confirm-email")
			anyChanged = true
		}
		if cmdutil.Changed(cmd, "max-margin-multiplier") {
			body.MaxMarginMultiplier = cmdutil.Str(cmd, "max-margin-multiplier")
			anyChanged = true
		}
		if cmdutil.Changed(cmd, "max-options-trading-level") {
			body.MaxOptionsTradingLevel = cmdutil.Int(cmd, "max-options-trading-level")
			anyChanged = true
		}
		if cmdutil.Changed(cmd, "disable-overnight-trading") {
			body.DisableOvernightTrading = cmdutil.Bool(cmd, "disable-overnight-trading")
			anyChanged = true
		}
		if cmdutil.Changed(cmd, "ptp-no-exception-entry") {
			body.PtpNoExceptionEntry = cmdutil.Bool(cmd, "ptp-no-exception-entry")
			anyChanged = true
		}

		if !anyChanged {
			return cmd.Help()
		}

		config, err := tradingClient.PatchAccountConfig(body)
		if err != nil {
			return err
		}
		return output.PrintSingle(cmd.OutOrStdout(), getOutput(), accountConfigColumns(), config)
	},
}

var accountConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage account configuration",
}

func init() {
	cmdutil.RegisterFlags(accountConfigSetCmd, api.PatchAccountConfigFlags, nil)

	accountConfigCmd.AddCommand(accountConfigGetCmd)
	accountConfigCmd.AddCommand(accountConfigSetCmd)

	accountCmd.AddCommand(accountGetCmd)
	accountCmd.AddCommand(accountConfigCmd)
}
