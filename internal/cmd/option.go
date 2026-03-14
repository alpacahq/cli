package cmd

import (
	"github.com/alpacahq/cli/internal/api"
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

var optionGetCmd = fetchCmd("get <symbol-or-id>", api.GetOptionContractSymbolOrIDOp, func(_ *cobra.Command, args []string) (any, error) {
	return tradingClient.GetOptionContractSymbolOrID(args[0])
}, func(c *cobra.Command) {
	c.Args = cobra.ExactArgs(1)
})

var optionExerciseCmd = fetchCmd("exercise <symbol-or-id>", api.OptionExerciseOp, func(cmd *cobra.Command, args []string) (any, error) {
	return voidResponse(tradingClient.OptionExercise(args[0]))
}, func(c *cobra.Command) {
	c.Args = cobra.ExactArgs(1)
})

var optionDoNotExerciseCmd = fetchCmd("do-not-exercise <symbol-or-id>", api.OptionDoNotExerciseOp, func(cmd *cobra.Command, args []string) (any, error) {
	return voidResponse(tradingClient.OptionDoNotExercise(args[0]))
}, func(c *cobra.Command) {
	c.Args = cobra.ExactArgs(1)
})

func init() {
	optionCmd.AddCommand(optionContractsCmd)
	optionCmd.AddCommand(optionGetCmd)
	optionCmd.AddCommand(optionExerciseCmd)
	optionCmd.AddCommand(optionDoNotExerciseCmd)
}
