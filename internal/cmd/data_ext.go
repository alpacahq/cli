package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

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
	Short: api.RatesOp.Summary(),
	Example: `  alpaca data forex rates --currency-pairs EUR/USD,GBP/USD --start 2025-01-01
  alpaca data forex rates --currency-pairs USD/JPY --timeframe 1Hour`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := cmdutil.RequireAll(cmd, api.RatesOp.RequiredFlags()...); err != nil {
			return err
		}

		resp, err := dataClient.Rates(ratesParamsFromFlags(cmd))
		if err != nil {
			return err
		}

		return renderMapValues(cmd.OutOrStdout(), getOutput(), forexRateColumns(), resp.Rates)
	},
}

var dataForexLatestCmd = &cobra.Command{
	Use:     "latest",
	Short:   api.LatestRatesOp.Summary(),
	Example: `  alpaca data forex latest --currency-pairs EUR/USD,GBP/USD`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := cmdutil.RequireAll(cmd, api.LatestRatesOp.RequiredFlags()...); err != nil {
			return err
		}

		resp, err := dataClient.LatestRates(latestRatesParamsFromFlags(cmd))
		if err != nil {
			return err
		}

		return output.JSON(cmd.OutOrStdout(), resp.Rates)
	},
}

// --- Crypto Orderbook ---

var dataCryptoOrderbookCmd = &cobra.Command{
	Use:     "crypto-orderbook",
	Short:   api.CryptoLatestOrderbooksOp.Summary(),
	Example: `  alpaca data crypto-orderbook --symbols BTC/USD,ETH/USD`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := cmdutil.RequireAll(cmd, api.CryptoLatestOrderbooksOp.RequiredFlags()...); err != nil {
			return err
		}

		resp, err := dataClient.CryptoLatestOrderbooks("us", cryptoLatestOrderbooksParamsFromFlags(cmd))
		if err != nil {
			return err
		}

		return output.JSON(cmd.OutOrStdout(), resp.Orderbooks)
	},
}

// --- Auctions ---

var dataAuctionsCmd = &cobra.Command{
	Use:   "auctions",
	Short: api.StockAuctionsOp.Summary(),
	Example: `  alpaca data auctions --symbols AAPL --start 2025-01-01
  alpaca data auctions --symbols AAPL,MSFT --limit 10`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := cmdutil.RequireAll(cmd, api.StockAuctionsOp.RequiredFlags()...); err != nil {
			return err
		}

		resp, err := dataClient.StockAuctions(stockAuctionsParamsFromFlags(cmd))
		if err != nil {
			return err
		}

		return output.JSON(cmd.OutOrStdout(), resp.Auctions)
	},
}

// --- Corporate Actions (market data) ---

var dataCorporateActionsCmd = &cobra.Command{
	Use:     "corporate-actions",
	Short:   api.CorporateActionsOp.Summary(),
	Example: `  alpaca data corporate-actions --symbols AAPL --types forward_split --start 2025-01-01`,
	RunE: func(cmd *cobra.Command, args []string) error {
		resp, err := dataClient.CorporateActions(corporateActionsParamsFromFlags(cmd))
		if err != nil {
			return err
		}

		return output.JSON(cmd.OutOrStdout(), resp)
	},
}

// --- Fixed Income Data ---

var dataFixedIncomeCmd = &cobra.Command{
	Use:     "fixed-income",
	Short:   api.FixedIncomeLatestPricesOp.Summary(),
	Example: `  alpaca data fixed-income --isins 912797KR1,912797LB5`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := cmdutil.RequireAll(cmd, api.FixedIncomeLatestPricesOp.RequiredFlags()...); err != nil {
			return err
		}

		resp, err := dataClient.FixedIncomeLatestPrices(fixedIncomeLatestPricesParamsFromFlags(cmd))
		if err != nil {
			return err
		}

		return output.JSON(cmd.OutOrStdout(), resp.Prices)
	},
}

// --- Logo ---

var dataLogoCmd = &cobra.Command{
	Use:     "logo <symbol>",
	Short:   api.LogosOp.Summary(),
	Example: `  alpaca data logo AAPL`,
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		resp, err := dataClient.Logos(args[0], logosParamsFromFlags(cmd))
		if err != nil {
			return err
		}
		return output.JSON(cmd.OutOrStdout(), resp)
	},
}

// --- Exchange & Condition Metadata ---

var dataMetaCmd = &cobra.Command{
	Use:   "meta",
	Short: "Stock exchange and condition reference data",
}

var dataMetaExchangesCmd = &cobra.Command{
	Use:     "exchanges",
	Short:   api.StockMetaExchangesOp.Summary(),
	Example: `  alpaca data meta exchanges`,
	RunE: func(cmd *cobra.Command, args []string) error {
		resp, err := dataClient.StockMetaExchanges()
		if err != nil {
			return err
		}
		return output.JSON(cmd.OutOrStdout(), resp)
	},
}

var dataMetaConditionsCmd = &cobra.Command{
	Use:     "conditions <ticktype>",
	Short:   api.StockMetaConditionsOp.Summary(),
	Example: `  alpaca data meta conditions trade`,
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		resp, err := dataClient.StockMetaConditions(args[0], stockMetaConditionsParamsFromFlags(cmd))
		if err != nil {
			return err
		}
		return output.JSON(cmd.OutOrStdout(), resp)
	},
}

func init() {
	cmdutil.RegisterFlags(dataForexRatesCmd, api.RatesOp.Flags(), nil)
	cmdutil.RegisterFlags(dataForexLatestCmd, api.LatestRatesOp.Flags(), nil)
	dataForexCmd.AddCommand(dataForexRatesCmd)
	dataForexCmd.AddCommand(dataForexLatestCmd)

	cmdutil.RegisterFlags(dataCryptoOrderbookCmd, api.CryptoLatestOrderbooksOp.Flags(), nil)

	cmdutil.RegisterFlags(dataAuctionsCmd, api.StockAuctionsOp.Flags(), nil)

	cmdutil.RegisterFlags(dataCorporateActionsCmd, api.CorporateActionsOp.Flags(), nil)

	cmdutil.RegisterFlags(dataFixedIncomeCmd, api.FixedIncomeLatestPricesOp.Flags(), nil)

	cmdutil.RegisterFlags(dataLogoCmd, api.LogosOp.Flags(), nil)

	cmdutil.RegisterFlags(dataMetaConditionsCmd, api.StockMetaConditionsOp.Flags(), nil)
	dataMetaCmd.AddCommand(dataMetaExchangesCmd)
	dataMetaCmd.AddCommand(dataMetaConditionsCmd)

	dataCmd.AddCommand(dataOptionCmd)
	dataCmd.AddCommand(dataForexCmd)
	dataCmd.AddCommand(dataCryptoOrderbookCmd)
	dataCmd.AddCommand(dataAuctionsCmd)
	dataCmd.AddCommand(dataCorporateActionsCmd)
	dataCmd.AddCommand(dataFixedIncomeCmd)
	dataCmd.AddCommand(dataLogoCmd)
	dataCmd.AddCommand(dataMetaCmd)
}

func renderMapValues(w io.Writer, format string, cols []output.Column, data any) error {
	j, _ := json.Marshal(data)
	var m map[string]json.RawMessage
	if json.Unmarshal(j, &m) == nil {
		if len(m) == 1 {
			for _, v := range m {
				return output.Render(w, format, cols, v)
			}
		}
		if len(m) > 1 && format != output.FormatJSON && !quietFlag {
			fmt.Fprintf(os.Stderr, "Hint: multi-symbol response (%d symbols); rendering as JSON. Use --json for structured output.\n", len(m))
		}
	}
	return output.JSON(w, data)
}
