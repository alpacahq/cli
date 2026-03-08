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

var dataForexLatestCmd = fetchCmd("latest", api.LatestRatesOp, func(cmd *cobra.Command, args []string) (any, error) {
	resp, err := dataClient.LatestRates(latestRatesParamsFromFlags(cmd))
	if err != nil {
		return nil, err
	}
	return resp.Rates, nil
}, withJSON, func(c *cobra.Command) {
	c.Example = `  alpaca data forex latest --currency-pairs EUR/USD,GBP/USD`
})

// --- Crypto Orderbook ---

var dataCryptoOrderbookCmd = fetchCmd("crypto-orderbook", api.CryptoLatestOrderbooksOp, func(cmd *cobra.Command, args []string) (any, error) {
	resp, err := dataClient.CryptoLatestOrderbooks("us", cryptoLatestOrderbooksParamsFromFlags(cmd))
	if err != nil {
		return nil, err
	}
	return resp.Orderbooks, nil
}, withJSON, func(c *cobra.Command) {
	c.Example = `  alpaca data crypto-orderbook --symbols BTC/USD,ETH/USD`
})

// --- Auctions ---

var dataAuctionsCmd = fetchCmd("auctions", api.StockAuctionsOp, func(cmd *cobra.Command, args []string) (any, error) {
	resp, err := dataClient.StockAuctions(stockAuctionsParamsFromFlags(cmd))
	if err != nil {
		return nil, err
	}
	return resp.Auctions, nil
}, withJSON, func(c *cobra.Command) {
	c.Example = `  alpaca data auctions --symbols AAPL --start 2025-01-01
  alpaca data auctions --symbols AAPL,MSFT --limit 10`
})

// --- Corporate Actions (market data) ---

var dataCorporateActionsCmd = fetchCmd("corporate-actions", api.CorporateActionsOp, func(cmd *cobra.Command, args []string) (any, error) {
	return dataClient.CorporateActions(corporateActionsParamsFromFlags(cmd))
}, withJSON, func(c *cobra.Command) {
	c.Example = `  alpaca data corporate-actions --symbols AAPL --types forward_split --start 2025-01-01`
})

// --- Fixed Income Data ---

var dataFixedIncomeCmd = fetchCmd("fixed-income", api.FixedIncomeLatestPricesOp, func(cmd *cobra.Command, args []string) (any, error) {
	resp, err := dataClient.FixedIncomeLatestPrices(fixedIncomeLatestPricesParamsFromFlags(cmd))
	if err != nil {
		return nil, err
	}
	return resp.Prices, nil
}, withJSON, func(c *cobra.Command) {
	c.Example = `  alpaca data fixed-income --isins 912797KR1,912797LB5`
})

// --- Logo ---

var dataLogoCmd = fetchCmd("logo <symbol>", api.LogosOp, func(cmd *cobra.Command, args []string) (any, error) {
	return dataClient.Logos(args[0], logosParamsFromFlags(cmd))
}, withJSON, func(c *cobra.Command) {
	c.Args = cobra.ExactArgs(1)
	c.Example = `  alpaca data logo AAPL`
})

// --- Exchange & Condition Metadata ---

var dataMetaCmd = &cobra.Command{
	Use:   "meta",
	Short: "Stock exchange and condition reference data",
}

var dataMetaExchangesCmd = fetchCmd("exchanges", api.StockMetaExchangesOp, func(cmd *cobra.Command, args []string) (any, error) {
	return dataClient.StockMetaExchanges()
}, withJSON, func(c *cobra.Command) {
	c.Example = `  alpaca data meta exchanges`
})

var dataMetaConditionsCmd = fetchCmd("conditions <ticktype>", api.StockMetaConditionsOp, func(cmd *cobra.Command, args []string) (any, error) {
	return dataClient.StockMetaConditions(args[0], stockMetaConditionsParamsFromFlags(cmd))
}, withJSON, func(c *cobra.Command) {
	c.Args = cobra.ExactArgs(1)
	c.Example = `  alpaca data meta conditions trade`
})

// --- Screener ---

var screenerCmd = &cobra.Command{
	Use:   "screener",
	Short: "Stock and crypto screener and market movers",
}

var screenerMostActivesCmd = fetchCmd("most-actives", api.MostActivesOp, func(cmd *cobra.Command, args []string) (any, error) {
	resp, err := dataClient.MostActives(mostActivesParamsFromFlags(cmd))
	if err != nil {
		return nil, err
	}
	return resp.MostActives, nil
}, func(c *cobra.Command) {
	c.Example = `  alpaca data screener most-actives
  alpaca data screener most-actives --by trades --top 10`
})

var screenerMoversCmd = &cobra.Command{
	Use:   "movers",
	Short: api.MoversOp.Summary(),
	Example: `  alpaca data screener movers
  alpaca data screener movers --market crypto --top 5`,
	RunE: func(cmd *cobra.Command, args []string) error {
		market := cmdutil.Str(cmd, "market")
		if market == "" {
			market = "stocks"
		}

		resp, err := dataClient.Movers(market, moversParamsFromFlags(cmd))
		if err != nil {
			return err
		}

		w := cmd.OutOrStdout()
		format := getOutput()
		if format == output.FormatJSON || format == output.FormatCSV {
			return output.Render(w, format, nil, resp)
		}

		cmd.Println("GAINERS")
		gainersJSON, _ := json.Marshal(resp.Gainers)
		if err := output.Render(w, output.FormatTable, nil, json.RawMessage(gainersJSON)); err != nil {
			return err
		}

		cmd.Println("\nLOSERS")
		losersJSON, _ := json.Marshal(resp.Losers)
		return output.Render(w, output.FormatTable, nil, json.RawMessage(losersJSON))
	},
}

// --- News ---

var newsCmd = fetchCmd("news", api.NewsOp, func(cmd *cobra.Command, args []string) (any, error) {
	params := newsParamsFromFlags(cmd)
	if params.Limit == 0 {
		params.Limit = 10
	}

	if cmdutil.Bool(cmd, "all") {
		max := cmdutil.Int(cmd, "max")
		var allNews []api.News
		for page := 0; page < maxPages; page++ {
			resp, err := dataClient.News(params)
			if err != nil {
				return nil, err
			}
			allNews = append(allNews, resp.News...)
			if max > 0 && len(allNews) >= max {
				allNews = allNews[:max]
				break
			}
			if resp.NextPageToken == "" {
				break
			}
			params.PageToken = resp.NextPageToken
		}
		return allNews, nil
	}

	resp, err := dataClient.News(params)
	if err != nil {
		return nil, err
	}
	return resp.News, nil
}, func(c *cobra.Command) {
	addPaginationFlags(c)
	c.Example = `  alpaca data news
  alpaca data news --symbols AAPL,MSFT --limit 10
  alpaca data news --symbols AAPL --all --max 100`
})

func init() {
	cmdutil.RegisterFlags(dataForexRatesCmd, api.RatesOp.Flags(), nil)

	dataForexCmd.AddCommand(dataForexRatesCmd)
	dataForexCmd.AddCommand(dataForexLatestCmd)

	dataMetaCmd.AddCommand(dataMetaExchangesCmd)
	dataMetaCmd.AddCommand(dataMetaConditionsCmd)

	cmdutil.RegisterFlags(screenerMoversCmd, api.MoversOp.Flags(), nil)
	screenerMoversCmd.Flags().String("market", "", "Market: stocks or crypto (default: stocks)")
	_ = screenerMoversCmd.RegisterFlagCompletionFunc("market", cobra.FixedCompletions(api.MarketTypeValues, cobra.ShellCompDirectiveNoFileComp))
	screenerCmd.AddCommand(screenerMostActivesCmd)
	screenerCmd.AddCommand(screenerMoversCmd)

	dataCmd.AddCommand(dataOptionCmd)
	dataCmd.AddCommand(dataForexCmd)
	dataCmd.AddCommand(dataCryptoOrderbookCmd)
	dataCmd.AddCommand(dataAuctionsCmd)
	dataCmd.AddCommand(dataCorporateActionsCmd)
	dataCmd.AddCommand(dataFixedIncomeCmd)
	dataCmd.AddCommand(dataLogoCmd)
	dataCmd.AddCommand(dataMetaCmd)
	dataCmd.AddCommand(screenerCmd)
	dataCmd.AddCommand(newsCmd)
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
