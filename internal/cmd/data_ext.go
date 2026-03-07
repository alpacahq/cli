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

var dataForexLatestCmd = jsonCmd("latest", api.LatestRatesOp, func(cmd *cobra.Command, args []string) (any, error) {
	resp, err := dataClient.LatestRates(latestRatesParamsFromFlags(cmd))
	if err != nil {
		return nil, err
	}
	return resp.Rates, nil
}, func(c *cobra.Command) {
	c.Example = `  alpaca data forex latest --currency-pairs EUR/USD,GBP/USD`
})

// --- Crypto Orderbook ---

var dataCryptoOrderbookCmd = jsonCmd("crypto-orderbook", api.CryptoLatestOrderbooksOp, func(cmd *cobra.Command, args []string) (any, error) {
	resp, err := dataClient.CryptoLatestOrderbooks("us", cryptoLatestOrderbooksParamsFromFlags(cmd))
	if err != nil {
		return nil, err
	}
	return resp.Orderbooks, nil
}, func(c *cobra.Command) {
	c.Example = `  alpaca data crypto-orderbook --symbols BTC/USD,ETH/USD`
})

// --- Auctions ---

var dataAuctionsCmd = jsonCmd("auctions", api.StockAuctionsOp, func(cmd *cobra.Command, args []string) (any, error) {
	resp, err := dataClient.StockAuctions(stockAuctionsParamsFromFlags(cmd))
	if err != nil {
		return nil, err
	}
	return resp.Auctions, nil
}, func(c *cobra.Command) {
	c.Example = `  alpaca data auctions --symbols AAPL --start 2025-01-01
  alpaca data auctions --symbols AAPL,MSFT --limit 10`
})

// --- Corporate Actions (market data) ---

var dataCorporateActionsCmd = jsonCmd("corporate-actions", api.CorporateActionsOp, func(cmd *cobra.Command, args []string) (any, error) {
	return dataClient.CorporateActions(corporateActionsParamsFromFlags(cmd))
}, func(c *cobra.Command) {
	c.Example = `  alpaca data corporate-actions --symbols AAPL --types forward_split --start 2025-01-01`
})

// --- Fixed Income Data ---

var dataFixedIncomeCmd = jsonCmd("fixed-income", api.FixedIncomeLatestPricesOp, func(cmd *cobra.Command, args []string) (any, error) {
	resp, err := dataClient.FixedIncomeLatestPrices(fixedIncomeLatestPricesParamsFromFlags(cmd))
	if err != nil {
		return nil, err
	}
	return resp.Prices, nil
}, func(c *cobra.Command) {
	c.Example = `  alpaca data fixed-income --isins 912797KR1,912797LB5`
})

// --- Logo ---

var dataLogoCmd = jsonCmd("logo <symbol>", api.LogosOp, func(cmd *cobra.Command, args []string) (any, error) {
	return dataClient.Logos(args[0], logosParamsFromFlags(cmd))
}, func(c *cobra.Command) {
	c.Args = cobra.ExactArgs(1)
	c.Example = `  alpaca data logo AAPL`
})

// --- Exchange & Condition Metadata ---

var dataMetaCmd = &cobra.Command{
	Use:   "meta",
	Short: "Stock exchange and condition reference data",
}

var dataMetaExchangesCmd = jsonCmd("exchanges", api.StockMetaExchangesOp, func(cmd *cobra.Command, args []string) (any, error) {
	return dataClient.StockMetaExchanges()
}, func(c *cobra.Command) {
	c.Example = `  alpaca data meta exchanges`
})

var dataMetaConditionsCmd = jsonCmd("conditions <ticktype>", api.StockMetaConditionsOp, func(cmd *cobra.Command, args []string) (any, error) {
	return dataClient.StockMetaConditions(args[0], stockMetaConditionsParamsFromFlags(cmd))
}, func(c *cobra.Command) {
	c.Args = cobra.ExactArgs(1)
	c.Example = `  alpaca data meta conditions trade`
})

func init() {
	cmdutil.RegisterFlags(dataForexRatesCmd, api.RatesOp.Flags(), nil)

	dataForexCmd.AddCommand(dataForexRatesCmd)
	dataForexCmd.AddCommand(dataForexLatestCmd)

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
