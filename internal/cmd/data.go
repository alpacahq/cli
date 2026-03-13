package cmd

import (
	"encoding/json"

	"github.com/alpacahq/cli/internal/api"
	"github.com/alpacahq/cli/internal/cmdutil"
	"github.com/spf13/cobra"
)

var dataCmd = &cobra.Command{
	Use:   "data",
	Short: "Access market data",
}

func fetchPaginated(
	cmd *cobra.Command,
	symbol string,
	fetch func(string, string) (json.RawMessage, error),
	extract func(json.RawMessage, string) json.RawMessage,
) (any, error) {
	if cmdutil.Bool(cmd, "all") {
		return fetchAllDataPages(
			func(pt string) (json.RawMessage, error) { return fetch(symbol, pt) },
			func(raw json.RawMessage) json.RawMessage { return extract(raw, symbol) },
			cmdutil.Int(cmd, "max"),
		)
	}
	data, err := fetch(symbol, "")
	if err != nil {
		return nil, err
	}
	return extract(data, symbol), nil
}

var dataBarsCmd = fetchCmd("bars <symbol>", api.StockBarSingleOp, func(cmd *cobra.Command, args []string) (any, error) {
	p := stockBarSingleParamsFromFlags(cmd)
	if p.Timeframe == "" {
		p.Timeframe = "1Day"
	}
	params := p.Values()
	return fetchPaginated(cmd, args[0],
		func(sym, pt string) (json.RawMessage, error) {
			if pt != "" {
				params.Set("page_token", pt)
			}
			return dataClient.Bars(sym, params)
		},
		func(data json.RawMessage, sym string) json.RawMessage { return extractArray(data, sym, "bars") },
	)
}, flagOpts(&cmdutil.FlagOpts{Defaults: map[string]string{"timeframe": "1Day"}}),
	func(c *cobra.Command) {
		c.Args = cobra.ExactArgs(1)
		c.Example = `  alpaca data bars AAPL --start 2025-01-01 --timeframe 1Day
  alpaca data bars BTC/USD --start 2025-01-01 --timeframe 1Hour
  alpaca data bars AAPL --start 2025-01-01 --end 2025-06-01 --limit 100
  alpaca data bars AAPL --start 2025-01-01 --all`
		cmdColumns[c] = barColumns()
		addPaginationFlags(c)
	})

var dataQuotesCmd = fetchCmd("quotes <symbol>", api.StockQuoteSingleOp, func(cmd *cobra.Command, args []string) (any, error) {
	params := stockQuoteSingleParamsFromFlags(cmd).Values()
	return fetchPaginated(cmd, args[0],
		func(sym, pt string) (json.RawMessage, error) {
			if pt != "" {
				params.Set("page_token", pt)
			}
			return dataClient.Quotes(sym, params)
		},
		func(data json.RawMessage, sym string) json.RawMessage { return extractArray(data, sym, "quotes") },
	)
}, func(c *cobra.Command) {
	c.Args = cobra.ExactArgs(1)
	c.Example = `  alpaca data quotes AAPL --start 2025-01-01
  alpaca data quotes AAPL --start 2025-01-01 --end 2025-01-31 --limit 50
  alpaca data quotes AAPL --start 2025-01-01 --all --max 5000`
	cmdColumns[c] = quoteColumns()
	addPaginationFlags(c)
})

var dataTradesCmd = fetchCmd("trades <symbol>", api.StockTradeSingleOp, func(cmd *cobra.Command, args []string) (any, error) {
	params := stockTradeSingleParamsFromFlags(cmd).Values()
	return fetchPaginated(cmd, args[0],
		func(sym, pt string) (json.RawMessage, error) {
			if pt != "" {
				params.Set("page_token", pt)
			}
			return dataClient.Trades(sym, params)
		},
		func(data json.RawMessage, sym string) json.RawMessage { return extractArray(data, sym, "trades") },
	)
}, func(c *cobra.Command) {
	c.Args = cobra.ExactArgs(1)
	c.Example = `  alpaca data trades AAPL --start 2025-01-01
  alpaca data trades AAPL --start 2025-01-01 --limit 100
  alpaca data trades AAPL --start 2025-01-01 --all`
	cmdColumns[c] = tradeColumns()
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
	symbol := args[0]
	data, err := dataClient.LatestTrade(symbol, stockLatestTradeSingleParamsFromFlags(cmd).Values())
	if err != nil {
		return nil, err
	}
	var m map[string]any
	if err := json.Unmarshal(data, &m); err != nil {
		return nil, err
	}
	return extractSingle(m, symbol, "trade", "trades"), nil
}, func(c *cobra.Command) {
	c.Args = cobra.ExactArgs(1)
	c.Example = `  alpaca data latest trade AAPL
  alpaca data latest trade AAPL --feed sip`
	cmdColumns[c] = tradeColumns()
})

var dataLatestQuoteCmd = fetchCmd("quote <symbol>", api.StockLatestQuoteSingleOp, func(cmd *cobra.Command, args []string) (any, error) {
	symbol := args[0]
	data, err := dataClient.LatestQuote(symbol, stockLatestQuoteSingleParamsFromFlags(cmd).Values())
	if err != nil {
		return nil, err
	}
	var m map[string]any
	if err := json.Unmarshal(data, &m); err != nil {
		return nil, err
	}
	return extractSingle(m, symbol, "quote", "quotes"), nil
}, func(c *cobra.Command) {
	c.Args = cobra.ExactArgs(1)
	c.Example = `  alpaca data latest quote AAPL
  alpaca data latest quote AAPL --json`
	cmdColumns[c] = quoteColumns()
})

var dataLatestBarCmd = fetchCmd("bar <symbol>", api.StockLatestBarSingleOp, func(cmd *cobra.Command, args []string) (any, error) {
	symbol := args[0]
	data, err := dataClient.LatestBar(symbol, stockLatestBarSingleParamsFromFlags(cmd).Values())
	if err != nil {
		return nil, err
	}
	var m map[string]any
	if err := json.Unmarshal(data, &m); err != nil {
		return nil, err
	}
	return extractSingle(m, symbol, "bar", "bars"), nil
}, func(c *cobra.Command) {
	c.Args = cobra.ExactArgs(1)
	c.Example = `  alpaca data latest bar AAPL
  alpaca data latest bar BTC/USD`
	cmdColumns[c] = barColumns()
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

func extractSingle(m map[string]any, symbol, singular, plural string) map[string]any {
	if v, ok := m[singular].(map[string]any); ok {
		return v
	}
	if multi, ok := m[plural].(map[string]any); ok {
		if v, ok := multi[symbol].(map[string]any); ok {
			return v
		}
	}
	return m
}
