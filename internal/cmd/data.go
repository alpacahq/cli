package cmd

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

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
  alpaca data bars AAPL --start 2025-01-01 --end 2025-06-01 --limit 100`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		symbol := args[0]

		params := dataParams(cmd)
		path, p := stockOrCryptoPath(symbol, "/bars", params)
		data, err := apiClient.GetData(path, p)
		if err != nil {
			return err
		}

		return output.Render(getOutput(), barColumns(), extractBars(data, symbol))
	},
}

var dataQuotesCmd = &cobra.Command{
	Use:   "quotes <symbol>",
	Short: "Get historical quotes",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		symbol := args[0]

		params := dataParams(cmd)
		path, p := stockOrCryptoPath(symbol, "/quotes", params)
		data, err := apiClient.GetData(path, p)
		if err != nil {
			return err
		}

		return output.Render(getOutput(), quoteColumns(), extractArray(data, symbol, "quotes"))
	},
}

var dataTradesCmd = &cobra.Command{
	Use:   "trades <symbol>",
	Short: "Get historical trades",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		symbol := args[0]

		params := dataParams(cmd)
		path, p := stockOrCryptoPath(symbol, "/trades", params)
		data, err := apiClient.GetData(path, p)
		if err != nil {
			return err
		}

		return output.Render(getOutput(), tradeColumns(), extractArray(data, symbol, "trades"))
	},
}

var dataSnapshotCmd = &cobra.Command{
	Use:   "snapshot <symbol>",
	Short: "Get latest snapshot (bar + quote + trade)",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		symbol := args[0]

		var path string
		if isCrypto(symbol) {
			encoded := strings.ReplaceAll(symbol, "/", "%2F")
			path = "/v1beta3/crypto/us/snapshots/" + encoded
		} else {
			path = "/v2/stocks/" + symbol + "/snapshot"
		}

		data, err := apiClient.GetData(path, nil)
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
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		symbol := args[0]
		var path string
		if isCrypto(symbol) {
			path = "/v1beta3/crypto/us/latest/trades"
		} else {
			path = "/v2/stocks/" + symbol + "/trades/latest"
		}

		params := url.Values{}
		if isCrypto(symbol) {
			params.Set("symbols", symbol)
		}

		data, err := apiClient.GetData(path, params)
		if err != nil {
			return err
		}

		if getOutput() == "json" {
			return output.JSON(cmd.OutOrStdout(), data)
		}

		m := make(map[string]any)
		if err := json.Unmarshal(data, &m); err != nil {
			return err
		}

		trade := extractTrade(m, symbol)
		if trade != nil {
			fmt.Printf("%s  $%v  (size: %v)\n", symbol, trade["p"], trade["s"])
		}
		return nil
	},
}

var dataLatestQuoteCmd = &cobra.Command{
	Use:   "quote <symbol>",
	Short: "Get latest quote",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		symbol := args[0]
		var path string
		if isCrypto(symbol) {
			path = "/v1beta3/crypto/us/latest/quotes"
		} else {
			path = "/v2/stocks/" + symbol + "/quotes/latest"
		}

		params := url.Values{}
		if isCrypto(symbol) {
			params.Set("symbols", symbol)
		}

		data, err := apiClient.GetData(path, params)
		if err != nil {
			return err
		}

		if getOutput() == "json" {
			return output.JSON(cmd.OutOrStdout(), data)
		}

		m := make(map[string]any)
		if err := json.Unmarshal(data, &m); err != nil {
			return err
		}

		quote := extractQuote(m, symbol)
		if quote != nil {
			fmt.Printf("%s  Bid: $%v (%v)  Ask: $%v (%v)\n",
				symbol, quote["bp"], quote["bs"], quote["ap"], quote["as"])
		}
		return nil
	},
}

var dataLatestBarCmd = &cobra.Command{
	Use:   "bar <symbol>",
	Short: "Get latest bar",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		symbol := args[0]
		var path string
		if isCrypto(symbol) {
			path = "/v1beta3/crypto/us/latest/bars"
		} else {
			path = "/v2/stocks/" + symbol + "/bars/latest"
		}

		params := url.Values{}
		if isCrypto(symbol) {
			params.Set("symbols", symbol)
		}

		data, err := apiClient.GetData(path, params)
		if err != nil {
			return err
		}

		return output.JSON(cmd.OutOrStdout(), data)
	},
}

func init() {
	for _, c := range []*cobra.Command{dataBarsCmd, dataQuotesCmd, dataTradesCmd} {
		c.Flags().String("start", "", "Start date (YYYY-MM-DD or RFC3339)")
		c.Flags().String("end", "", "End date (YYYY-MM-DD or RFC3339)")
		c.Flags().Int("limit", 0, "Max number of results")
		c.Flags().String("feed", "", "Data feed: iex, sip, otc, delayed_sip")
		_ = c.RegisterFlagCompletionFunc("feed", cobra.FixedCompletions([]string{"iex", "sip", "otc", "delayed_sip"}, cobra.ShellCompDirectiveNoFileComp))
		c.Flags().String("currency", "", "Currency for prices (e.g. USD, EUR)")
		c.Flags().String("sort", "", "Sort order: asc or desc")
		_ = c.RegisterFlagCompletionFunc("sort", cobra.FixedCompletions([]string{"asc", "desc"}, cobra.ShellCompDirectiveNoFileComp))
		c.Flags().String("asof", "", "As-of date for point-in-time data")
	}
	dataBarsCmd.Flags().String("timeframe", "1Day", "Bar timeframe: 1Min, 5Min, 15Min, 1Hour, 1Day, 1Week, 1Month")
	_ = dataBarsCmd.RegisterFlagCompletionFunc("timeframe", cobra.FixedCompletions([]string{"1Min", "5Min", "15Min", "1Hour", "1Day", "1Week", "1Month"}, cobra.ShellCompDirectiveNoFileComp))
	dataBarsCmd.Flags().String("adjustment", "", "Price adjustment: raw, split, dividend, all")
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

func isCrypto(symbol string) bool {
	return strings.Contains(symbol, "/")
}

func stockOrCryptoPath(symbol, endpoint string, params url.Values) (string, url.Values) {
	if isCrypto(symbol) {
		params.Set("symbols", symbol)
		return "/v1beta3/crypto/us" + endpoint, params
	}
	return "/v2/stocks/" + symbol + endpoint, params
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
	return params
}

func extractBars(data json.RawMessage, symbol string) json.RawMessage {
	var m map[string]json.RawMessage
	if json.Unmarshal(data, &m) != nil {
		return data
	}

	if bars, ok := m["bars"]; ok {
		var multi map[string]json.RawMessage
		if json.Unmarshal(bars, &multi) == nil {
			if symBars, ok := multi[symbol]; ok {
				return symBars
			}
		}
		return bars
	}

	return data
}

func extractArray(data json.RawMessage, symbol, key string) json.RawMessage {
	var m map[string]json.RawMessage
	if json.Unmarshal(data, &m) != nil {
		return data
	}
	if arr, ok := m[key]; ok {
		var multi map[string]json.RawMessage
		if json.Unmarshal(arr, &multi) == nil {
			if symArr, ok := multi[symbol]; ok {
				return symArr
			}
		}
		return arr
	}
	return data
}

func extractTrade(m map[string]any, symbol string) map[string]any {
	if trade, ok := m["trade"].(map[string]any); ok {
		return trade
	}
	if trades, ok := m["trades"].(map[string]any); ok {
		if t, ok := trades[symbol].(map[string]any); ok {
			return t
		}
	}
	return m
}

func extractQuote(m map[string]any, symbol string) map[string]any {
	if quote, ok := m["quote"].(map[string]any); ok {
		return quote
	}
	if quotes, ok := m["quotes"].(map[string]any); ok {
		if q, ok := quotes[symbol].(map[string]any); ok {
			return q
		}
	}
	return m
}
