package cmd

import (
	"encoding/json"
	"io"

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
	Short: api.OperationSummary["Rates"],
	Example: `  alpaca data forex rates --pairs EUR/USD,GBP/USD --start 2025-01-01
  alpaca data forex rates --pairs USD/JPY --timeframe 1Hour`,
	RunE: func(cmd *cobra.Command, args []string) error {
		pairs, err := cmdutil.RequireStr(cmd, "pairs")
		if err != nil {
			return err
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

		return renderMapValues(cmd.OutOrStdout(), getOutput(), forexRateColumns(), resp.Rates)
	},
}

var dataForexLatestCmd = &cobra.Command{
	Use:   "latest",
	Short: api.OperationSummary["LatestRates"],
	Example: `  alpaca data forex latest --pairs EUR/USD,GBP/USD`,
	RunE: func(cmd *cobra.Command, args []string) error {
		pairs, err := cmdutil.RequireStr(cmd, "pairs")
		if err != nil {
			return err
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
	Short: api.OperationSummary["CryptoLatestOrderbooks"],
	Example: `  alpaca data crypto-orderbook --symbols BTC/USD,ETH/USD`,
	RunE: func(cmd *cobra.Command, args []string) error {
		symbols, err := cmdutil.RequireStr(cmd, "symbols")
		if err != nil {
			return err
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
	Short: api.OperationSummary["StockAuctions"],
	Example: `  alpaca data auctions --symbols AAPL --start 2025-01-01
  alpaca data auctions --symbols AAPL,MSFT --limit 10`,
	RunE: func(cmd *cobra.Command, args []string) error {
		symbols, err := cmdutil.RequireStr(cmd, "symbols")
		if err != nil {
			return err
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
	Short: api.OperationSummary["CorporateActions"],
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
	Short: api.OperationSummary["FixedIncomeLatestPrices"],
	Example: `  alpaca data fixed-income --symbols 912797KR1,912797LB5`,
	RunE: func(cmd *cobra.Command, args []string) error {
		symbols, err := cmdutil.RequireStr(cmd, "symbols")
		if err != nil {
			return err
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
	_ = dataForexRatesCmd.RegisterFlagCompletionFunc("timeframe", cobra.FixedCompletions([]string{"1Min", "5Min", "1Hour", "1Day"}, cobra.ShellCompDirectiveNoFileComp))
	cmdutil.AddDateRangeFlags(dataForexRatesCmd)
	cmdutil.AddLimitFlag(dataForexRatesCmd)
	cmdutil.AddSortFlag(dataForexRatesCmd, api.SortValues)
	dataForexLatestCmd.Flags().String("pairs", "", "Currency pairs (comma-separated)")
	dataForexCmd.AddCommand(dataForexRatesCmd)
	dataForexCmd.AddCommand(dataForexLatestCmd)

	dataCryptoOrderbookCmd.Flags().String("symbols", "", "Crypto symbols (comma-separated, e.g. BTC/USD)")

	dataAuctionsCmd.Flags().String("symbols", "", "Stock symbols (comma-separated)")
	cmdutil.AddDateRangeFlags(dataAuctionsCmd)
	cmdutil.AddLimitFlag(dataAuctionsCmd)
	cmdutil.AddSortFlag(dataAuctionsCmd, api.SortValues)
	dataAuctionsCmd.Flags().String("asof", "", "As-of date for data")

	dataCorporateActionsCmd.Flags().String("symbols", "", "Filter by symbols")
	dataCorporateActionsCmd.Flags().String("types", "", "CA types: forward_split, reverse_split, cash_dividend, etc.")
	cmdutil.AddDateRangeFlags(dataCorporateActionsCmd)
	cmdutil.AddLimitFlag(dataCorporateActionsCmd)
	cmdutil.AddSortFlag(dataCorporateActionsCmd, api.SortValues)
	dataFixedIncomeCmd.Flags().String("symbols", "", "ISIN identifiers (comma-separated)")

	dataCmd.AddCommand(dataOptionCmd)
	dataCmd.AddCommand(dataForexCmd)
	dataCmd.AddCommand(dataCryptoOrderbookCmd)
	dataCmd.AddCommand(dataAuctionsCmd)
	dataCmd.AddCommand(dataCorporateActionsCmd)
	dataCmd.AddCommand(dataFixedIncomeCmd)
}

func renderMapValues(w io.Writer, format string, cols []output.Column, data any) error {
	j, _ := json.Marshal(data)
	var m map[string]json.RawMessage
	if json.Unmarshal(j, &m) == nil && len(m) == 1 {
		for _, v := range m {
			return output.Render(w, format, cols, v)
		}
	}
	return output.JSON(w, data)
}
