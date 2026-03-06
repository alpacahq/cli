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
	Example: `  alpaca account get
  alpaca account get --json`,
	RunE: func(cmd *cobra.Command, args []string) error {
		account, err := tradingClient.GetAccount()
		if err != nil {
			return err
		}
		return output.PrintSingle(cmd.OutOrStdout(), getOutput(), nil, account)
	},
}

var accountConfigGetCmd = &cobra.Command{
	Use:   "get",
	Short: api.GetAccountConfigOp.Summary,
	Example: `  alpaca account config get
  alpaca account config get --json`,
	RunE: func(cmd *cobra.Command, args []string) error {
		config, err := tradingClient.GetAccountConfig()
		if err != nil {
			return err
		}
		return output.PrintSingle(cmd.OutOrStdout(), getOutput(), nil, config)
	},
}

var accountConfigSetCmd = &cobra.Command{
	Use:   "set",
	Short: api.PatchAccountConfigOp.Summary,
	Example: `  alpaca account config set --no-shorting true
  alpaca account config set --dtbp-check entry`,
	RunE: func(cmd *cobra.Command, args []string) error {
		body := &api.AccountConfigurations{}
		p := cmdutil.NewPatchHelper(cmd)

		p.Str("dtbp-check", &body.DTBPCheck)
		p.Bool("no-shorting", &body.NoShorting)
		p.Str("pdt-check", &body.PDTCheck)
		p.Bool("fractional-trading", &body.FractionalTrading)
		p.Bool("suspend-trade", &body.SuspendTrade)
		p.Str("trade-confirm-email", &body.TradeConfirmEmail)
		p.Str("max-margin-multiplier", &body.MaxMarginMultiplier)
		p.Int("max-options-trading-level", &body.MaxOptionsTradingLevel)
		p.Bool("disable-overnight-trading", &body.DisableOvernightTrading)
		p.Bool("ptp-no-exception-entry", &body.PtpNoExceptionEntry)

		if !p.AnyChanged() {
			return cmd.Help()
		}

		config, err := tradingClient.PatchAccountConfig(body)
		if err != nil {
			return err
		}
		return output.PrintSingle(cmd.OutOrStdout(), getOutput(), nil, config)
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
