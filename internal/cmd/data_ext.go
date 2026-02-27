package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/alpacahq/cli/internal/api"
	"github.com/alpacahq/cli/internal/cmdutil"
	"github.com/alpacahq/cli/internal/output"
	"github.com/spf13/cobra"
)

// --- Forex ---

var dataForexCmd = &cobra.Command{
	Use:   "forex",
	Short: "Foreign exchange rate data",
}

var dataForexRatesCmd = &cobra.Command{
	Use:   "rates",
	Short: "Get historical forex rates",
	Example: `  alpaca data forex rates --pairs EUR/USD,GBP/USD --start 2025-01-01
  alpaca data forex rates --pairs USD/JPY --timeframe 1Hour`,
	RunE: func(cmd *cobra.Command, args []string) error {
		pairs := cmdutil.Str(cmd, "pairs")
		if pairs == "" {
			return fmt.Errorf("--pairs is required (e.g. EUR/USD,GBP/USD)")
		}

		resp, err := dataClient.Rates(&api.RatesParams{
			CurrencyPairs: pairs,
			Timeframe:     cmdutil.Str(cmd, "timeframe"),
			Start:         cmdutil.Str(cmd, "start"),
			End:           cmdutil.Str(cmd, "end"),
			Limit:         cmdutil.Int(cmd, "limit"),
			Sort:          cmdutil.Str(cmd, "sort"),
		})
		if err != nil {
			return err
		}

		return renderMapValues(getOutput(), forexRateColumns(), resp.Rates)
	},
}

var dataForexLatestCmd = &cobra.Command{
	Use:   "latest",
	Short: "Get latest forex rates",
	Example: `  alpaca data forex latest --pairs EUR/USD,GBP/USD`,
	RunE: func(cmd *cobra.Command, args []string) error {
		pairs := cmdutil.Str(cmd, "pairs")
		if pairs == "" {
			return fmt.Errorf("--pairs is required (e.g. EUR/USD,GBP/USD)")
		}

		resp, err := dataClient.LatestRates(&api.LatestRatesParams{
			CurrencyPairs: pairs,
		})
		if err != nil {
			return err
		}

		return output.JSON(cmd.OutOrStdout(), resp.Rates)
	},
}

// --- Crypto Orderbook ---

var dataCryptoOrderbookCmd = &cobra.Command{
	Use:   "crypto-orderbook",
	Short: "Get latest crypto orderbooks",
	Example: `  alpaca data crypto-orderbook --symbols BTC/USD,ETH/USD`,
	RunE: func(cmd *cobra.Command, args []string) error {
		symbols := cmdutil.Str(cmd, "symbols")
		if symbols == "" {
			return fmt.Errorf("--symbols is required (e.g. BTC/USD,ETH/USD)")
		}

		resp, err := dataClient.CryptoLatestOrderbooks("us", &api.CryptoLatestOrderbooksParams{
			Symbols: symbols,
		})
		if err != nil {
			return err
		}

		return output.JSON(cmd.OutOrStdout(), resp.Orderbooks)
	},
}

// --- Auctions ---

var dataAuctionsCmd = &cobra.Command{
	Use:   "auctions",
	Short: "Get stock auction data",
	Example: `  alpaca data auctions --symbols AAPL --start 2025-01-01
  alpaca data auctions --symbols AAPL,MSFT --limit 10`,
	RunE: func(cmd *cobra.Command, args []string) error {
		symbols := cmdutil.Str(cmd, "symbols")
		if symbols == "" {
			return fmt.Errorf("--symbols is required")
		}

		resp, err := dataClient.StockAuctions(&api.StockAuctionsParams{
			Symbols: symbols,
			Start:   cmdutil.Str(cmd, "start"),
			End:     cmdutil.Str(cmd, "end"),
			Limit:   cmdutil.Int(cmd, "limit"),
			Sort:    cmdutil.Str(cmd, "sort"),
			Asof:    cmdutil.Str(cmd, "asof"),
		})
		if err != nil {
			return err
		}

		return output.JSON(cmd.OutOrStdout(), resp.Auctions)
	},
}

// --- Corporate Actions (market data) ---

var dataCorporateActionsCmd = &cobra.Command{
	Use:   "corporate-actions",
	Short: "Get corporate actions data (market data API)",
	Example: `  alpaca data corporate-actions --symbols AAPL --types forward_split --start 2025-01-01`,
	RunE: func(cmd *cobra.Command, args []string) error {
		resp, err := dataClient.CorporateActions(&api.CorporateActionsParams{
			Symbols: cmdutil.Str(cmd, "symbols"),
			Types:   cmdutil.Str(cmd, "types"),
			Start:   cmdutil.Str(cmd, "start"),
			End:     cmdutil.Str(cmd, "end"),
			Limit:   cmdutil.Int(cmd, "limit"),
			Sort:    cmdutil.Str(cmd, "sort"),
		})
		if err != nil {
			return err
		}

		return output.JSON(cmd.OutOrStdout(), resp)
	},
}

// --- Fixed Income Data ---

var dataFixedIncomeCmd = &cobra.Command{
	Use:   "fixed-income",
	Short: "Get fixed income latest prices",
	Example: `  alpaca data fixed-income --symbols 912797KR1,912797LB5`,
	RunE: func(cmd *cobra.Command, args []string) error {
		symbols := cmdutil.Str(cmd, "symbols")
		if symbols == "" {
			return fmt.Errorf("--symbols is required (CUSIP identifiers)")
		}

		resp, err := dataClient.FixedIncomeLatestPrices(&api.FixedIncomeLatestPricesParams{
			Isins: symbols,
		})
		if err != nil {
			return err
		}

		return output.JSON(cmd.OutOrStdout(), resp.Prices)
	},
}

func init() {
	dataForexRatesCmd.Flags().String("pairs", "", "Currency pairs (comma-separated, e.g. EUR/USD,GBP/USD)")
	dataForexRatesCmd.Flags().String("timeframe", "", "Timeframe: 1Min, 5Min, 1Hour, 1Day")
	dataForexRatesCmd.Flags().String("start", "", "Start date")
	dataForexRatesCmd.Flags().String("end", "", "End date")
	dataForexRatesCmd.Flags().Int("limit", 0, "Max results")
	dataForexRatesCmd.Flags().String("sort", "", "Sort: asc or desc")
	dataForexLatestCmd.Flags().String("pairs", "", "Currency pairs (comma-separated)")
	dataForexCmd.AddCommand(dataForexRatesCmd)
	dataForexCmd.AddCommand(dataForexLatestCmd)

	dataCryptoOrderbookCmd.Flags().String("symbols", "", "Crypto symbols (comma-separated, e.g. BTC/USD)")

	dataAuctionsCmd.Flags().String("symbols", "", "Stock symbols (comma-separated)")
	dataAuctionsCmd.Flags().String("start", "", "Start date")
	dataAuctionsCmd.Flags().String("end", "", "End date")
	dataAuctionsCmd.Flags().Int("limit", 0, "Max results")
	dataAuctionsCmd.Flags().String("sort", "", "Sort: asc or desc")
	dataAuctionsCmd.Flags().String("asof", "", "As-of date for data")

	dataCorporateActionsCmd.Flags().String("symbols", "", "Filter by symbols")
	dataCorporateActionsCmd.Flags().String("types", "", "CA types: forward_split, reverse_split, cash_dividend, etc.")
	dataCorporateActionsCmd.Flags().String("start", "", "Start date")
	dataCorporateActionsCmd.Flags().String("end", "", "End date")
	dataCorporateActionsCmd.Flags().Int("limit", 0, "Max results")
	dataCorporateActionsCmd.Flags().String("sort", "", "Sort: asc or desc")
	dataFixedIncomeCmd.Flags().String("symbols", "", "ISIN identifiers (comma-separated)")

	dataCmd.AddCommand(dataOptionCmd)
	dataCmd.AddCommand(dataForexCmd)
	dataCmd.AddCommand(dataCryptoOrderbookCmd)
	dataCmd.AddCommand(dataAuctionsCmd)
	dataCmd.AddCommand(dataCorporateActionsCmd)
	dataCmd.AddCommand(dataFixedIncomeCmd)
}

func renderMapValues(format string, cols []output.Column, data any) error {
	j, _ := json.Marshal(data)
	var m map[string]json.RawMessage
	if json.Unmarshal(j, &m) == nil && len(m) == 1 {
		for _, v := range m {
			return output.Render(format, cols, v)
		}
	}
	return output.JSON(nil, data)
}
