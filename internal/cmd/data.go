package cmd

import (
	"encoding/json"
	"net/url"

	"github.com/alpacahq/cli/internal/api"
	"github.com/alpacahq/cli/internal/cmdutil"
	"github.com/spf13/cobra"
)

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

var dataBarsCmd = fetchCmd("bars", api.StockBarSingleOp, func(cmd *cobra.Command, args []string) (any, error) {
	p := stockBarSingleParamsFromFlags(cmd)
	return fetchPaginated(cmd, cmdutil.Str(cmd, "symbol"), p.Values(), dataClient.Bars, "bars")
}, func(c *cobra.Command) {
	c.Example = `  alpaca data bars --symbol AAPL --start 2025-01-01 --timeframe 1Day
  alpaca data bars --symbol BTC/USD --start 2025-01-01 --timeframe 1Hour
  alpaca data bars --symbol AAPL --start 2025-01-01 --end 2025-06-01 --limit 100
  alpaca data bars --symbol AAPL --start 2025-01-01 --all`
	addPaginationFlags(c)
})

var dataQuotesCmd = fetchCmd("quotes", api.StockQuoteSingleOp, func(cmd *cobra.Command, args []string) (any, error) {
	return fetchPaginated(cmd, cmdutil.Str(cmd, "symbol"), stockQuoteSingleParamsFromFlags(cmd).Values(), dataClient.Quotes, "quotes")
}, func(c *cobra.Command) {
	c.Example = `  alpaca data quotes --symbol AAPL --start 2025-01-01
  alpaca data quotes --symbol AAPL --start 2025-01-01 --end 2025-01-31 --limit 50
  alpaca data quotes --symbol AAPL --start 2025-01-01 --all --max 5000`
	addPaginationFlags(c)
})

var dataTradesCmd = fetchCmd("trades", api.StockTradeSingleOp, func(cmd *cobra.Command, args []string) (any, error) {
	return fetchPaginated(cmd, cmdutil.Str(cmd, "symbol"), stockTradeSingleParamsFromFlags(cmd).Values(), dataClient.Trades, "trades")
}, func(c *cobra.Command) {
	c.Example = `  alpaca data trades --symbol AAPL --start 2025-01-01
  alpaca data trades --symbol AAPL --start 2025-01-01 --limit 100
  alpaca data trades --symbol AAPL --start 2025-01-01 --all`
	addPaginationFlags(c)
})

var dataSnapshotCmd = fetchCmd("snapshot", api.StockSnapshotSingleOp, func(cmd *cobra.Command, args []string) (any, error) {
	return dataClient.Snapshot(cmdutil.Str(cmd, "symbol"), stockSnapshotSingleParamsFromFlags(cmd).Values())
}, jsonOnly, func(c *cobra.Command) {
	c.Long = "Returns the latest snapshot for a symbol. Output is always JSON due to complex nested structure."
	c.Example = `  alpaca data snapshot --symbol AAPL
  alpaca data snapshot --symbol BTC/USD --feed sip`
})

var dataLatestTradeCmd = fetchCmd("trade", api.StockLatestTradeSingleOp, func(cmd *cobra.Command, args []string) (any, error) {
	return dataClient.LatestTrade(cmdutil.Str(cmd, "symbol"), stockLatestTradeSingleParamsFromFlags(cmd).Values())
}, func(c *cobra.Command) {
	c.Example = `  alpaca data latest trade --symbol AAPL
  alpaca data latest trade --symbol AAPL --feed sip`
})

var dataLatestQuoteCmd = fetchCmd("quote", api.StockLatestQuoteSingleOp, func(cmd *cobra.Command, args []string) (any, error) {
	return dataClient.LatestQuote(cmdutil.Str(cmd, "symbol"), stockLatestQuoteSingleParamsFromFlags(cmd).Values())
}, func(c *cobra.Command) {
	c.Example = `  alpaca data latest quote --symbol AAPL`
})

var dataLatestBarCmd = fetchCmd("bar", api.StockLatestBarSingleOp, func(cmd *cobra.Command, args []string) (any, error) {
	return dataClient.LatestBar(cmdutil.Str(cmd, "symbol"), stockLatestBarSingleParamsFromFlags(cmd).Values())
}, func(c *cobra.Command) {
	c.Example = `  alpaca data latest bar --symbol AAPL
  alpaca data latest bar --symbol BTC/USD`
})

func addPaginationFlags(cmd *cobra.Command) {
	cmd.Flags().Bool("all", false, "Fetch all pages")
	cmd.Flags().Int("max", 10000, "Maximum items when using --all")
}

func init() {
	dataLatestCmd.AddCommand(dataLatestTradeCmd)
	dataLatestCmd.AddCommand(dataLatestQuoteCmd)
	dataLatestCmd.AddCommand(dataLatestBarCmd)

	dataCmd.AddCommand(dataBarsCmd)
	dataCmd.AddCommand(dataQuotesCmd)
	dataCmd.AddCommand(dataTradesCmd)
	dataCmd.AddCommand(dataSnapshotCmd)
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
