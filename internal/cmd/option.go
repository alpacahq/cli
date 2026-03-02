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
	Short: "Get option chain for an underlying symbol (contracts)",
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
			StrikePriceGte:    cmdutil.Float64(cmd, "strike-gte"),
			StrikePriceLte:    cmdutil.Float64(cmd, "strike-lte"),
			ExpirationDateGte: cmdutil.Str(cmd, "expiry-gte"),
			ExpirationDateLte: cmdutil.Str(cmd, "expiry-lte"),
			RootSymbol:        cmdutil.Str(cmd, "root-symbol"),
			Limit:             cmdutil.Int(cmd, "limit"),
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
	Short: "Get option contract details",
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
	Short: "Exercise an option position",
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
	Short: "Mark an option position as do-not-exercise",
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
	optionChainCmd.Flags().String("expiry", "", "Exact expiration date (YYYY-MM-DD)")
	optionChainCmd.Flags().String("expiry-gte", "", "Expiration date on or after (YYYY-MM-DD)")
	optionChainCmd.Flags().String("expiry-lte", "", "Expiration date on or before (YYYY-MM-DD)")
	optionChainCmd.Flags().String("type", "", "Option type: call or put")
	_ = optionChainCmd.RegisterFlagCompletionFunc("type", cobra.FixedCompletions(api.OptionContractTypeValues, cobra.ShellCompDirectiveNoFileComp))
	optionChainCmd.Flags().Float64("strike-gte", 0, "Minimum strike price")
	optionChainCmd.Flags().Float64("strike-lte", 0, "Maximum strike price")
	optionChainCmd.Flags().String("root-symbol", "", "Root symbol for options")
	optionChainCmd.Flags().Int("limit", 0, "Max results")
	optionChainCmd.Flags().Bool("show-deliverables", false, "Include deliverables info")

	optionCmd.AddCommand(optionChainCmd)
	optionCmd.AddCommand(optionGetCmd)
	optionCmd.AddCommand(optionExerciseCmd)
	optionCmd.AddCommand(optionDoNotExerciseCmd)
}
