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
  alpaca option chain AAPL --expiry 2025-06-20 --type call
  alpaca option chain SPY --strike-gte 400 --strike-lte 450`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		params := &api.GetOptionsContractsParams{
			UnderlyingSymbols: args[0],
			ShowDeliverables:  cmdutil.Bool(cmd, "show-deliverables"),
			ExpirationDate:    cmdutil.Str(cmd, "expiry"),
			Type:              cmdutil.Str(cmd, "type"),
			Style:             cmdutil.Str(cmd, "style"),
			Status:            cmdutil.Str(cmd, "status"),
			StrikePriceGte:    cmdutil.Float64(cmd, "strike-gte"),
			StrikePriceLte:    cmdutil.Float64(cmd, "strike-lte"),
			ExpirationDateGte: cmdutil.Str(cmd, "expiry-gte"),
			ExpirationDateLte: cmdutil.Str(cmd, "expiry-lte"),
			RootSymbol:        cmdutil.Str(cmd, "root-symbol"),
			Limit:             cmdutil.Int(cmd, "limit"),
			PageToken:         cmdutil.Str(cmd, "page-token"),
		}

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
		warnLive()
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
		warnLive()
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
		Exclude: map[string]bool{"underlying_symbols": true, "ppind": true},
		Aliases: map[string]string{
			"expiration_date":     "expiry",
			"expiration_date_gte": "expiry-gte",
			"expiration_date_lte": "expiry-lte",
			"strike_price_gte":    "strike-gte",
			"strike_price_lte":    "strike-lte",
		},
	})

	optionCmd.AddCommand(optionChainCmd)
	optionCmd.AddCommand(optionGetCmd)
	optionCmd.AddCommand(optionExerciseCmd)
	optionCmd.AddCommand(optionDoNotExerciseCmd)
}
