package cmd

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/alpacahq/cli/internal/api"
	"github.com/alpacahq/cli/internal/cmdutil"
	"github.com/alpacahq/cli/internal/output"
	"github.com/spf13/cobra"
)

var dataCmd = &cobra.Command{
	Use:   "data",
	Short: "Access market data",
}

var dataBarsCmd = &cobra.Command{
	Use:   "bars <symbol>",
	Short: "Get historical price bars",
	Example: `  alpaca data bars AAPL --start 2025-01-01 --timeframe 1Day
  alpaca data bars BTC/USD --start 2025-01-01 --timeframe 1Hour
  alpaca data bars AAPL --start 2025-01-01 --end 2025-06-01 --limit 100
  alpaca data bars AAPL --start 2025-01-01 --all`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		symbol := args[0]
		params := dataParams(cmd)

		if cmdutil.Bool(cmd, "all") {
			data, err := fetchAllDataPages(
				func(pt string) (json.RawMessage, error) {
					if pt != "" {
						params.Set("page_token", pt)
					}
					return dataClient.Bars(symbol, params)
				},
				func(raw json.RawMessage) json.RawMessage { return extractBars(raw, symbol) },
				cmdutil.Int(cmd, "max"),
			)
			if err != nil {
				return err
			}
			return output.Render(cmd.OutOrStdout(), getOutput(), barColumns(), data)
		}

		data, err := dataClient.Bars(symbol, params)
		if err != nil {
			return err
		}
		return output.Render(cmd.OutOrStdout(), getOutput(), barColumns(), extractBars(data, symbol))
	},
}

var dataQuotesCmd = &cobra.Command{
	Use:   "quotes <symbol>",
	Short: "Get historical quotes",
	Example: `  alpaca data quotes AAPL --start 2025-01-01
  alpaca data quotes AAPL --start 2025-01-01 --end 2025-01-31 --limit 50
  alpaca data quotes AAPL --start 2025-01-01 --all --max 5000`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		symbol := args[0]
		params := dataParams(cmd)

		if cmdutil.Bool(cmd, "all") {
			data, err := fetchAllDataPages(
				func(pt string) (json.RawMessage, error) {
					if pt != "" {
						params.Set("page_token", pt)
					}
					return dataClient.Quotes(symbol, params)
				},
				func(raw json.RawMessage) json.RawMessage { return extractArray(raw, symbol, "quotes") },
				cmdutil.Int(cmd, "max"),
			)
			if err != nil {
				return err
			}
			return output.Render(cmd.OutOrStdout(), getOutput(), quoteColumns(), data)
		}

		data, err := dataClient.Quotes(symbol, dataParams(cmd))
		if err != nil {
			return err
		}
		return output.Render(cmd.OutOrStdout(), getOutput(), quoteColumns(), extractArray(data, symbol, "quotes"))
	},
}

var dataTradesCmd = &cobra.Command{
	Use:   "trades <symbol>",
	Short: "Get historical trades",
	Example: `  alpaca data trades AAPL --start 2025-01-01
  alpaca data trades AAPL --start 2025-01-01 --limit 100
  alpaca data trades AAPL --start 2025-01-01 --all`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		symbol := args[0]
		params := dataParams(cmd)

		if cmdutil.Bool(cmd, "all") {
			data, err := fetchAllDataPages(
				func(pt string) (json.RawMessage, error) {
					if pt != "" {
						params.Set("page_token", pt)
					}
					return dataClient.Trades(symbol, params)
				},
				func(raw json.RawMessage) json.RawMessage { return extractArray(raw, symbol, "trades") },
				cmdutil.Int(cmd, "max"),
			)
			if err != nil {
				return err
			}
			return output.Render(cmd.OutOrStdout(), getOutput(), tradeColumns(), data)
		}

		data, err := dataClient.Trades(symbol, dataParams(cmd))
		if err != nil {
			return err
		}
		return output.Render(cmd.OutOrStdout(), getOutput(), tradeColumns(), extractArray(data, symbol, "trades"))
	},
}

var dataSnapshotCmd = &cobra.Command{
	Use:   "snapshot <symbol>",
	Short: "Get latest snapshot (bar + quote + trade)",
	Long:  "Returns the latest snapshot for a symbol. Output is always JSON due to complex nested structure.",
	Example: `  alpaca data snapshot AAPL
  alpaca data snapshot BTC/USD --feed sip`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		data, err := dataClient.Snapshot(args[0], latestParams(cmd))
		if err != nil {
			return err
		}
		return output.JSON(cmd.OutOrStdout(), data)
	},
}

var dataLatestCmd = &cobra.Command{
	Use:   "latest",
	Short: "Get latest market data",
}

var dataLatestTradeCmd = &cobra.Command{
	Use:   "trade <symbol>",
	Short: "Get latest trade",
	Example: `  alpaca data latest trade AAPL
  alpaca data latest trade AAPL --feed sip`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		symbol := args[0]
		data, err := dataClient.LatestTrade(symbol, latestParams(cmd))
		if err != nil {
			return err
		}

		var m map[string]any
		if err := json.Unmarshal(data, &m); err != nil {
			return err
		}

		trade := extractSingle(m, symbol, "trade", "trades")
		return output.PrintSingle(cmd.OutOrStdout(), getOutput(), tradeColumns(), trade)
	},
}

var dataLatestQuoteCmd = &cobra.Command{
	Use:   "quote <symbol>",
	Short: "Get latest quote",
	Example: `  alpaca data latest quote AAPL
  alpaca data latest quote AAPL --json`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		symbol := args[0]
		data, err := dataClient.LatestQuote(symbol, latestParams(cmd))
		if err != nil {
			return err
		}

		var m map[string]any
		if err := json.Unmarshal(data, &m); err != nil {
			return err
		}

		quote := extractSingle(m, symbol, "quote", "quotes")
		return output.PrintSingle(cmd.OutOrStdout(), getOutput(), quoteColumns(), quote)
	},
}

var dataLatestBarCmd = &cobra.Command{
	Use:   "bar <symbol>",
	Short: "Get latest bar",
	Example: `  alpaca data latest bar AAPL
  alpaca data latest bar BTC/USD`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		symbol := args[0]
		data, err := dataClient.LatestBar(symbol, latestParams(cmd))
		if err != nil {
			return err
		}

		var m map[string]any
		if err := json.Unmarshal(data, &m); err != nil {
			return err
		}

		bar := extractSingle(m, symbol, "bar", "bars")
		return output.PrintSingle(cmd.OutOrStdout(), getOutput(), barColumns(), bar)
	},
}

func addPaginationFlags(cmd *cobra.Command) {
	cmd.Flags().Bool("all", false, "Fetch all pages")
	cmd.Flags().Int("max", 10000, "Maximum items when using --all")
}

func init() {
	feedCompletions := cobra.FixedCompletions([]string{"iex", "sip", "otc", "delayed_sip"}, cobra.ShellCompDirectiveNoFileComp)

	cmdutil.RegisterFlags(dataBarsCmd, api.StockBarSingleFlags, &cmdutil.FlagOpts{
		Defaults: map[string]string{"timeframe": "1Day"},
	})
	addPaginationFlags(dataBarsCmd)
	_ = dataBarsCmd.RegisterFlagCompletionFunc("feed", feedCompletions)
	_ = dataBarsCmd.RegisterFlagCompletionFunc("timeframe", cobra.FixedCompletions([]string{"1Min", "5Min", "15Min", "1Hour", "1Day", "1Week", "1Month"}, cobra.ShellCompDirectiveNoFileComp))
	_ = dataBarsCmd.RegisterFlagCompletionFunc("adjustment", cobra.FixedCompletions([]string{"raw", "split", "dividend", "all"}, cobra.ShellCompDirectiveNoFileComp))

	cmdutil.RegisterFlags(dataQuotesCmd, api.StockQuoteSingleFlags, nil)
	_ = dataQuotesCmd.RegisterFlagCompletionFunc("feed", feedCompletions)
	addPaginationFlags(dataQuotesCmd)

	cmdutil.RegisterFlags(dataTradesCmd, api.StockTradeSingleFlags, nil)
	_ = dataTradesCmd.RegisterFlagCompletionFunc("feed", feedCompletions)
	addPaginationFlags(dataTradesCmd)

	cmdutil.RegisterFlags(dataSnapshotCmd, api.StockSnapshotSingleFlags, nil)
	_ = dataSnapshotCmd.RegisterFlagCompletionFunc("feed", feedCompletions)

	cmdutil.RegisterFlags(dataLatestTradeCmd, api.StockLatestTradeSingleFlags, nil)
	_ = dataLatestTradeCmd.RegisterFlagCompletionFunc("feed", feedCompletions)

	cmdutil.RegisterFlags(dataLatestQuoteCmd, api.StockLatestQuoteSingleFlags, nil)
	_ = dataLatestQuoteCmd.RegisterFlagCompletionFunc("feed", feedCompletions)

	cmdutil.RegisterFlags(dataLatestBarCmd, api.StockLatestBarSingleFlags, nil)
	_ = dataLatestBarCmd.RegisterFlagCompletionFunc("feed", feedCompletions)

	dataLatestCmd.AddCommand(dataLatestTradeCmd)
	dataLatestCmd.AddCommand(dataLatestQuoteCmd)
	dataLatestCmd.AddCommand(dataLatestBarCmd)

	dataCmd.AddCommand(dataBarsCmd)
	dataCmd.AddCommand(dataQuotesCmd)
	dataCmd.AddCommand(dataTradesCmd)
	dataCmd.AddCommand(dataSnapshotCmd)
	dataCmd.AddCommand(dataLatestCmd)
}

func dataParams(cmd *cobra.Command) url.Values {
	params := url.Values{}
	for _, key := range []string{"start", "end", "timeframe", "feed", "currency", "sort", "adjustment", "asof"} {
		if v := cmdutil.Str(cmd, key); v != "" {
			params.Set(key, v)
		}
	}
	if v := cmdutil.Int(cmd, "limit"); v != 0 {
		params.Set("limit", fmt.Sprint(v))
	}
	if v := cmdutil.Str(cmd, "page-token"); v != "" {
		params.Set("page_token", v)
	}
	return params
}

func latestParams(cmd *cobra.Command) url.Values {
	params := url.Values{}
	for _, key := range []string{"feed", "currency"} {
		if v := cmdutil.Str(cmd, key); v != "" {
			params.Set(key, v)
		}
	}
	return params
}

func extractBars(data json.RawMessage, symbol string) json.RawMessage {
	var m map[string]json.RawMessage
	if err := json.Unmarshal(data, &m); err != nil {
		verboseLog("extractBars: unmarshal top-level: %v", err)
		return data
	}

	if bars, ok := m["bars"]; ok {
		var multi map[string]json.RawMessage
		if err := json.Unmarshal(bars, &multi); err == nil {
			if symBars, ok := multi[symbol]; ok {
				return symBars
			}
		} else {
			verboseLog("extractBars: unmarshal bars: %v", err)
		}
		return bars
	}

	return data
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
