package cmd

import (
	"encoding/json"
	"net/url"

	"github.com/alpacahq/cli/internal/api"
	"github.com/alpacahq/cli/internal/cmdutil"
	"github.com/spf13/cobra"
)

var dataCmd = &cobra.Command{
	Use:   "data",
	Short: "Access market data",
}

// fetchPaginated returns raw API response for a single page, or accumulates
// all pages when --all is set. The key param identifies the array field inside
// the response envelope (e.g. "bars", "trades", "quotes") for pagination extraction.
func fetchPaginated(
	cmd *cobra.Command,
	symbol string,
	params url.Values,
	fetch func(string, url.Values) (json.RawMessage, error),
	key string,
) (any, error) {
	if cmdutil.Bool(cmd, "all") {
		return fetchAllDataPages(
			func(pt string) (json.RawMessage, error) {
				if pt != "" {
					params.Set("page_token", pt)
				}
				return fetch(symbol, params)
			},
			func(raw json.RawMessage) json.RawMessage { return extractArray(raw, symbol, key) },
			cmdutil.Int(cmd, "max"),
		)
	}
	return fetch(symbol, params)
}

var dataBarsCmd = fetchCmd("bars <symbol>", api.StockBarSingleOp, func(cmd *cobra.Command, args []string) (any, error) {
	p := stockBarSingleParamsFromFlags(cmd)
	if p.Timeframe == "" {
		p.Timeframe = "1Day"
	}
	return fetchPaginated(cmd, args[0], p.Values(), dataClient.Bars, "bars")
}, flagOpts(&cmdutil.FlagOpts{Defaults: map[string]string{"timeframe": "1Day"}}),
	func(c *cobra.Command) {
		c.Args = cobra.ExactArgs(1)
		c.Example = `  alpaca data bars AAPL --start 2025-01-01 --timeframe 1Day
  alpaca data bars BTC/USD --start 2025-01-01 --timeframe 1Hour
  alpaca data bars AAPL --start 2025-01-01 --end 2025-06-01 --limit 100
  alpaca data bars AAPL --start 2025-01-01 --all`
		addPaginationFlags(c)
	})

var dataQuotesCmd = fetchCmd("quotes <symbol>", api.StockQuoteSingleOp, func(cmd *cobra.Command, args []string) (any, error) {
	return fetchPaginated(cmd, args[0], stockQuoteSingleParamsFromFlags(cmd).Values(), dataClient.Quotes, "quotes")
}, func(c *cobra.Command) {
	c.Args = cobra.ExactArgs(1)
	c.Example = `  alpaca data quotes AAPL --start 2025-01-01
  alpaca data quotes AAPL --start 2025-01-01 --end 2025-01-31 --limit 50
  alpaca data quotes AAPL --start 2025-01-01 --all --max 5000`
	addPaginationFlags(c)
})

var dataTradesCmd = fetchCmd("trades <symbol>", api.StockTradeSingleOp, func(cmd *cobra.Command, args []string) (any, error) {
	return fetchPaginated(cmd, args[0], stockTradeSingleParamsFromFlags(cmd).Values(), dataClient.Trades, "trades")
}, func(c *cobra.Command) {
	c.Args = cobra.ExactArgs(1)
	c.Example = `  alpaca data trades AAPL --start 2025-01-01
  alpaca data trades AAPL --start 2025-01-01 --limit 100
  alpaca data trades AAPL --start 2025-01-01 --all`
	addPaginationFlags(c)
})

var dataSnapshotCmd = fetchCmd("snapshot <symbol>", api.StockSnapshotSingleOp, func(cmd *cobra.Command, args []string) (any, error) {
	return dataClient.Snapshot(args[0], stockSnapshotSingleParamsFromFlags(cmd).Values())
}, jsonOnly, func(c *cobra.Command) {
	c.Args = cobra.ExactArgs(1)
	c.Long = "Returns the latest snapshot for a symbol. Output is always JSON due to complex nested structure."
	c.Example = `  alpaca data snapshot AAPL
  alpaca data snapshot BTC/USD --feed sip`
})

var dataLatestCmd = &cobra.Command{
	Use:   "latest",
	Short: "Get latest market data",
}

var dataLatestTradeCmd = fetchCmd("trade <symbol>", api.StockLatestTradeSingleOp, func(cmd *cobra.Command, args []string) (any, error) {
	return dataClient.LatestTrade(args[0], stockLatestTradeSingleParamsFromFlags(cmd).Values())
}, func(c *cobra.Command) {
	c.Args = cobra.ExactArgs(1)
	c.Example = `  alpaca data latest trade AAPL
  alpaca data latest trade AAPL --feed sip`
})

var dataLatestQuoteCmd = fetchCmd("quote <symbol>", api.StockLatestQuoteSingleOp, func(cmd *cobra.Command, args []string) (any, error) {
	return dataClient.LatestQuote(args[0], stockLatestQuoteSingleParamsFromFlags(cmd).Values())
}, func(c *cobra.Command) {
	c.Args = cobra.ExactArgs(1)
	c.Example = `  alpaca data latest quote AAPL`
})

var dataLatestBarCmd = fetchCmd("bar <symbol>", api.StockLatestBarSingleOp, func(cmd *cobra.Command, args []string) (any, error) {
	return dataClient.LatestBar(args[0], stockLatestBarSingleParamsFromFlags(cmd).Values())
}, func(c *cobra.Command) {
	c.Args = cobra.ExactArgs(1)
	c.Example = `  alpaca data latest bar AAPL
  alpaca data latest bar BTC/USD`
})

func addPaginationFlags(cmd *cobra.Command) {
	cmd.Flags().Bool("all", false, "Fetch all pages")
	cmd.Flags().Int("max", 10000, "Maximum items when using --all")
}

func init() {
	_ = dataBarsCmd.RegisterFlagCompletionFunc("timeframe", cobra.FixedCompletions([]string{"1Min", "5Min", "15Min", "1Hour", "1Day", "1Week", "1Month"}, cobra.ShellCompDirectiveNoFileComp))
	_ = dataBarsCmd.RegisterFlagCompletionFunc("adjustment", cobra.FixedCompletions([]string{"raw", "split", "dividend", "all"}, cobra.ShellCompDirectiveNoFileComp))

	dataLatestCmd.AddCommand(dataLatestTradeCmd)
	dataLatestCmd.AddCommand(dataLatestQuoteCmd)
	dataLatestCmd.AddCommand(dataLatestBarCmd)

	dataCmd.AddCommand(dataBarsCmd)
	dataCmd.AddCommand(dataQuotesCmd)
	dataCmd.AddCommand(dataTradesCmd)
	dataCmd.AddCommand(dataSnapshotCmd)
	dataCmd.AddCommand(dataLatestCmd)
}

func extractArray(data json.RawMessage, symbol, key string) json.RawMessage {
	var m map[string]json.RawMessage
	if err := json.Unmarshal(data, &m); err != nil {
		verboseLog("extractArray: unmarshal top-level: %v", err)
		return data
	}
	if arr, ok := m[key]; ok {
		var multi map[string]json.RawMessage
		if err := json.Unmarshal(arr, &multi); err == nil {
			if symArr, ok := multi[symbol]; ok {
				return symArr
			}
		} else {
			verboseLog("extractArray: unmarshal %s: %v", key, err)
		}
		return arr
	}
	return data
}
