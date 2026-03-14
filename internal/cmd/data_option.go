package cmd

import (
	"github.com/alpacahq/cli/internal/api"
	"github.com/alpacahq/cli/internal/cmdutil"
	"github.com/spf13/cobra"
)

var dataOptionCmd = &cobra.Command{
	Use:   "option",
	Short: "Options market data",
}

var dataOptionBarsCmd = fetchCmd("bars", api.OptionBarsOp, func(cmd *cobra.Command, args []string) (any, error) {
	return dataClient.OptionBars(optionBarsParamsFromFlags(cmd))
}, jsonOnly, func(c *cobra.Command) {
	c.Example = `  alpaca data option bars --symbols AAPL250620C00200000 --start 2025-01-01
  alpaca data option bars --symbols AAPL250620C00200000,AAPL250620P00200000 --timeframe 1Day`
})

var dataOptionTradesCmd = fetchCmd("trades", api.OptionTradesOp, func(cmd *cobra.Command, args []string) (any, error) {
	return dataClient.OptionTrades(optionTradesParamsFromFlags(cmd))
}, jsonOnly, func(c *cobra.Command) {
	c.Example = `  alpaca data option trades --symbols AAPL250620C00200000 --start 2025-01-01`
})

var dataOptionSnapshotCmd = fetchCmd("snapshot", api.OptionSnapshotsOp, func(cmd *cobra.Command, args []string) (any, error) {
	return dataClient.OptionSnapshots(optionSnapshotsParamsFromFlags(cmd))
}, jsonOnly, func(c *cobra.Command) {
	c.Example = `  alpaca data option snapshot --symbols AAPL250620C00200000
  alpaca data option snapshot --symbols AAPL250620C00200000,AAPL250620P00200000`
})

var dataOptionChainCmd = fetchCmd("chain", api.OptionChainOp, func(cmd *cobra.Command, args []string) (any, error) {
	return dataClient.OptionChain(cmdutil.Str(cmd, "underlying-symbol"), optionChainParamsFromFlags(cmd))
}, jsonOnly, func(c *cobra.Command) {
	c.Example = `  alpaca data option chain --underlying-symbol AAPL
  alpaca data option chain --underlying-symbol SPY --expiration-date 2025-06-20 --type call`
})

var dataOptionLatestQuotesCmd = fetchCmd("latest-quotes", api.OptionLatestQuotesOp, func(cmd *cobra.Command, args []string) (any, error) {
	return dataClient.OptionLatestQuotes(optionLatestQuotesParamsFromFlags(cmd))
}, jsonOnly, func(c *cobra.Command) {
	c.Example = `  alpaca data option latest-quotes --symbols AAPL250620C00200000`
})

var dataOptionLatestTradesCmd = fetchCmd("latest-trades", api.OptionLatestTradesOp, func(cmd *cobra.Command, args []string) (any, error) {
	return dataClient.OptionLatestTrades(optionLatestTradesParamsFromFlags(cmd))
}, jsonOnly, func(c *cobra.Command) {
	c.Example = `  alpaca data option latest-trades --symbols AAPL250620C00200000`
})

var dataOptionExchangesCmd = fetchCmd("exchanges", api.OptionMetaExchangesOp, func(cmd *cobra.Command, args []string) (any, error) {
	return dataClient.OptionMetaExchanges()
}, jsonOnly, func(c *cobra.Command) {
	c.Example = `  alpaca data option exchanges`
})

var dataOptionConditionsCmd = fetchCmd("conditions", api.OptionMetaConditionsOp, func(cmd *cobra.Command, args []string) (any, error) {
	return dataClient.OptionMetaConditions(cmdutil.Str(cmd, "ticktype"))
}, jsonOnly, func(c *cobra.Command) {
	c.Example = `  alpaca data option conditions --ticktype trade`
})

func init() {
	dataOptionCmd.AddCommand(dataOptionBarsCmd)
	dataOptionCmd.AddCommand(dataOptionTradesCmd)
	dataOptionCmd.AddCommand(dataOptionSnapshotCmd)
	dataOptionCmd.AddCommand(dataOptionChainCmd)
	dataOptionCmd.AddCommand(dataOptionLatestQuotesCmd)
	dataOptionCmd.AddCommand(dataOptionLatestTradesCmd)
	dataOptionCmd.AddCommand(dataOptionExchangesCmd)
	dataOptionCmd.AddCommand(dataOptionConditionsCmd)
}
