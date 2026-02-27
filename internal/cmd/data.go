package cmd

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/alpacahq/cli/internal/output"
	"github.com/spf13/cobra"
)

var dataCmd = &cobra.Command{
	Use:   "data",
	Short: "Access market data",
}

var dataBarsCmd = &cobra.Command{
	Use:   "bars",
	Short: "Get historical price bars",
	Example: `  alpaca data bars --symbol AAPL --start 2025-01-01 --timeframe 1Day
  alpaca data bars --symbol BTC/USD --start 2025-01-01 --timeframe 1Hour
  alpaca data bars --symbol AAPL --start 2025-01-01 --end 2025-06-01 --limit 100`,
	RunE: func(cmd *cobra.Command, args []string) error {
		symbol, _ := cmd.Flags().GetString("symbol")
		if symbol == "" {
			return fmt.Errorf("--symbol is required")
		}

		params := dataParams(cmd)

		path, p := stockOrCryptoPath(symbol, "/bars", params)
		data, err := apiClient.GetData(path, p)
		if err != nil {
			return err
		}

		bars := extractBars(data, symbol)
		columns := []output.Column{
			{Header: "TIMESTAMP", Field: "t"},
			{Header: "OPEN", Field: "o"},
			{Header: "HIGH", Field: "h"},
			{Header: "LOW", Field: "l"},
			{Header: "CLOSE", Field: "c"},
			{Header: "VOLUME", Field: "v"},
			{Header: "VWAP", Field: "vw"},
		}

		return output.Render(getOutput(), columns, bars)
	},
}

var dataQuotesCmd = &cobra.Command{
	Use:   "quotes",
	Short: "Get historical quotes",
	RunE: func(cmd *cobra.Command, args []string) error {
		symbol, _ := cmd.Flags().GetString("symbol")
		if symbol == "" {
			return fmt.Errorf("--symbol is required")
		}

		params := dataParams(cmd)
		path, p := stockOrCryptoPath(symbol, "/quotes", params)
		data, err := apiClient.GetData(path, p)
		if err != nil {
			return err
		}

		quotes := extractArray(data, symbol, "quotes")
		columns := []output.Column{
			{Header: "TIMESTAMP", Field: "t"},
			{Header: "BID", Field: "bp"},
			{Header: "BID SIZE", Field: "bs"},
			{Header: "ASK", Field: "ap"},
			{Header: "ASK SIZE", Field: "as"},
		}

		return output.Render(getOutput(), columns, quotes)
	},
}

var dataTradesCmd = &cobra.Command{
	Use:   "trades",
	Short: "Get historical trades",
	RunE: func(cmd *cobra.Command, args []string) error {
		symbol, _ := cmd.Flags().GetString("symbol")
		if symbol == "" {
			return fmt.Errorf("--symbol is required")
		}

		params := dataParams(cmd)
		path, p := stockOrCryptoPath(symbol, "/trades", params)
		data, err := apiClient.GetData(path, p)
		if err != nil {
			return err
		}

		trades := extractArray(data, symbol, "trades")
		columns := []output.Column{
			{Header: "TIMESTAMP", Field: "t"},
			{Header: "PRICE", Field: "p"},
			{Header: "SIZE", Field: "s"},
			{Header: "EXCHANGE", Field: "x"},
		}

		return output.Render(getOutput(), columns, trades)
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
	for _, cmd := range []*cobra.Command{dataBarsCmd, dataQuotesCmd, dataTradesCmd} {
		cmd.Flags().String("symbol", "", "Ticker symbol (e.g. AAPL, BTC/USD)")
		cmd.Flags().String("start", "", "Start date (YYYY-MM-DD or RFC3339)")
		cmd.Flags().String("end", "", "End date (YYYY-MM-DD or RFC3339)")
		cmd.Flags().String("limit", "", "Max number of results")
		cmd.Flags().String("feed", "", "Data feed: iex, sip, otc, delayed_sip")
		cmd.Flags().String("currency", "", "Currency for prices (e.g. USD, EUR)")
		cmd.Flags().String("sort", "", "Sort order: asc or desc")
	}
	dataBarsCmd.Flags().String("timeframe", "1Day", "Bar timeframe: 1Min, 5Min, 15Min, 1Hour, 1Day, 1Week, 1Month")
	dataBarsCmd.Flags().String("adjustment", "", "Price adjustment: raw, split, dividend, all")

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
	start, _ := cmd.Flags().GetString("start")
	if start != "" {
		params.Set("start", start)
	}
	end, _ := cmd.Flags().GetString("end")
	if end != "" {
		params.Set("end", end)
	}
	limit, _ := cmd.Flags().GetString("limit")
	if limit != "" {
		params.Set("limit", limit)
	}
	tf, _ := cmd.Flags().GetString("timeframe")
	if tf != "" {
		params.Set("timeframe", tf)
	}
	feed, _ := cmd.Flags().GetString("feed")
	if feed != "" {
		params.Set("feed", feed)
	}
	currency, _ := cmd.Flags().GetString("currency")
	if currency != "" {
		params.Set("currency", currency)
	}
	sort, _ := cmd.Flags().GetString("sort")
	if sort != "" {
		params.Set("sort", sort)
	}
	adjustment, _ := cmd.Flags().GetString("adjustment")
	if adjustment != "" {
		params.Set("adjustment", adjustment)
	}
	return params
}

func extractBars(data json.RawMessage, symbol string) json.RawMessage {
	var m map[string]json.RawMessage
	if json.Unmarshal(data, &m) != nil {
		return data
	}

	// Stock single-symbol: {"bars": [...]}
	if bars, ok := m["bars"]; ok {
		// Multi-symbol: {"bars": {"AAPL": [...]}}
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

