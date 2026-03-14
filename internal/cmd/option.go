package cmd

import (
	"github.com/alpacahq/cli/internal/api"
	"github.com/alpacahq/cli/internal/cmdutil"
	"github.com/spf13/cobra"
)

var optionCmd = &cobra.Command{
	Use:   "option",
	Short: "Options trading",
}

var optionContractsCmd = fetchCmd("contracts", api.GetOptionsContractsOp, func(cmd *cobra.Command, args []string) (any, error) {
	return tradingClient.GetOptionsContracts(getOptionsContractsParamsFromFlags(cmd))
}, func(c *cobra.Command) {
	c.Long = "List option contracts for an underlying symbol. For market data (greeks, pricing), use `data option chain`."
	c.Example = `  alpaca option contracts --underlying-symbols AAPL
  alpaca option contracts --underlying-symbols AAPL --expiration-date 2025-06-20 --type call
  alpaca option contracts --underlying-symbols SPY --strike-price-gte 400 --strike-price-lte 450`
})

var optionGetCmd = fetchCmd("get", api.GetOptionContractSymbolOrIDOp, func(cmd *cobra.Command, args []string) (any, error) {
	return tradingClient.GetOptionContractSymbolOrID(cmdutil.Str(cmd, "symbol-or-id"))
})

var optionExerciseCmd = fetchCmd("exercise", api.OptionExerciseOp, func(cmd *cobra.Command, args []string) (any, error) {
	return voidResponse(tradingClient.OptionExercise(cmdutil.Str(cmd, "symbol-or-contract-id")))
})

var optionDoNotExerciseCmd = fetchCmd("do-not-exercise", api.OptionDoNotExerciseOp, func(cmd *cobra.Command, args []string) (any, error) {
	return voidResponse(tradingClient.OptionDoNotExercise(cmdutil.Str(cmd, "symbol-or-contract-id")))
})

func init() {
	optionCmd.AddCommand(optionContractsCmd)
	optionCmd.AddCommand(optionGetCmd)
	optionCmd.AddCommand(optionExerciseCmd)
	optionCmd.AddCommand(optionDoNotExerciseCmd)
}
