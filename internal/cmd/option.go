package cmd

import (
	"fmt"

	"github.com/alpacahq/cli/internal/api"
	"github.com/alpacahq/cli/internal/cmdutil"
	"github.com/alpacahq/cli/internal/output"
	"github.com/spf13/cobra"
)

var optionCmd = &cobra.Command{
	Use:   "option",
	Short: "Options trading",
}

var optionChainCmd = &cobra.Command{
	Use:   "chain <underlying>",
	Short: api.GetOptionsContractsOp.Summary,
	Long:  "List option contracts for an underlying symbol. For market data (greeks, pricing), use `data option chain`.",
	Example: `  alpaca option chain AAPL
  alpaca option chain AAPL --expiration-date 2025-06-20 --type call
  alpaca option chain SPY --strike-price-gte 400 --strike-price-lte 450`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		params := getOptionsContractsParamsFromFlags(cmd)
		params.UnderlyingSymbols = args[0]
		data, err := tradingClient.GetOptionsContracts(params)
		if err != nil {
			return err
		}
		return output.Render(cmd.OutOrStdout(), getOutput(), optionChainColumns(), data)
	},
}

var optionGetCmd = &cobra.Command{
	Use:   "get <symbol-or-id>",
	Short: api.GetOptionContractSymbolOrIDOp.Summary,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		contract, err := tradingClient.GetOptionContractSymbolOrID(args[0])
		if err != nil {
			return err
		}
		return output.PrintSingle(cmd.OutOrStdout(), getOutput(), optionChainColumns(), contract)
	},
}

var optionExerciseCmd = &cobra.Command{
	Use:   "exercise <symbol-or-id>",
	Short: api.OptionExerciseOp.Summary,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		_, err := tradingClient.OptionExercise(args[0])
		if err != nil {
			return err
		}
		_, _ = fmt.Fprintf(cmd.OutOrStdout(), "Option %s exercise requested.\n", args[0])
		return nil
	},
}

var optionDoNotExerciseCmd = &cobra.Command{
	Use:   "do-not-exercise <symbol-or-id>",
	Short: api.OptionDoNotExerciseOp.Summary,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		_, err := tradingClient.OptionDoNotExercise(args[0])
		if err != nil {
			return err
		}
		_, _ = fmt.Fprintf(cmd.OutOrStdout(), "Option %s marked as do-not-exercise.\n", args[0])
		return nil
	},
}

func init() {
	cmdutil.RegisterFlags(optionChainCmd, api.GetOptionsContractsFlags, &cmdutil.FlagOpts{
		Exclude: map[string]bool{"underlying_symbols": true},
	})

	optionCmd.AddCommand(optionChainCmd)
	optionCmd.AddCommand(optionGetCmd)
	optionCmd.AddCommand(optionExerciseCmd)
	optionCmd.AddCommand(optionDoNotExerciseCmd)
}
