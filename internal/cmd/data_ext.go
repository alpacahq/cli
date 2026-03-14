package cmd

import (
	"github.com/alpacahq/cli/internal/api"
	"github.com/alpacahq/cli/internal/cmdutil"
	"github.com/spf13/cobra"
)

// --- Forex ---

var dataForexCmd = &cobra.Command{
	Use:   "forex",
	Short: "Foreign exchange rate data",
}

var dataForexRatesCmd = fetchCmd("rates", api.RatesOp, func(cmd *cobra.Command, args []string) (any, error) {
	return dataClient.Rates(ratesParamsFromFlags(cmd))
}, jsonOnly, func(c *cobra.Command) {
	c.Example = `  alpaca data forex rates --currency-pairs EUR/USD,GBP/USD --start 2025-01-01
  alpaca data forex rates --currency-pairs USD/JPY --timeframe 1Hour`
})

var dataForexLatestCmd = fetchCmd("latest", api.LatestRatesOp, func(cmd *cobra.Command, args []string) (any, error) {
	return dataClient.LatestRates(latestRatesParamsFromFlags(cmd))
}, jsonOnly, func(c *cobra.Command) {
	c.Example = `  alpaca data forex latest --currency-pairs EUR/USD,GBP/USD`
})

// --- Crypto Orderbook ---

var dataCryptoOrderbookCmd = fetchCmd("crypto-orderbook", api.CryptoLatestOrderbooksOp, func(cmd *cobra.Command, args []string) (any, error) {
	return dataClient.CryptoLatestOrderbooks("us", cryptoLatestOrderbooksParamsFromFlags(cmd))
}, jsonOnly, func(c *cobra.Command) {
	c.Example = `  alpaca data crypto-orderbook --symbols BTC/USD,ETH/USD`
})

// --- Auctions ---

var dataAuctionsCmd = fetchCmd("auctions", api.StockAuctionsOp, func(cmd *cobra.Command, args []string) (any, error) {
	return dataClient.StockAuctions(stockAuctionsParamsFromFlags(cmd))
}, jsonOnly, func(c *cobra.Command) {
	c.Example = `  alpaca data auctions --symbols AAPL --start 2025-01-01
  alpaca data auctions --symbols AAPL,MSFT --limit 10`
})

// --- Corporate Actions (market data) ---

var dataCorporateActionsCmd = fetchCmd("corporate-actions", api.CorporateActionsOp, func(cmd *cobra.Command, args []string) (any, error) {
	return dataClient.CorporateActions(corporateActionsParamsFromFlags(cmd))
}, jsonOnly, func(c *cobra.Command) {
	c.Example = `  alpaca data corporate-actions --symbols AAPL --types forward_split --start 2025-01-01`
})

// --- Fixed Income Data ---

var dataFixedIncomeCmd = fetchCmd("fixed-income", api.FixedIncomeLatestPricesOp, func(cmd *cobra.Command, args []string) (any, error) {
	return dataClient.FixedIncomeLatestPrices(fixedIncomeLatestPricesParamsFromFlags(cmd))
}, jsonOnly, func(c *cobra.Command) {
	c.Example = `  alpaca data fixed-income --isins 912797KR1,912797LB5`
})

// --- Logo ---

var dataLogoCmd = fetchCmd("logo <symbol>", api.LogosOp, func(cmd *cobra.Command, args []string) (any, error) {
	return dataClient.Logos(args[0], logosParamsFromFlags(cmd))
}, jsonOnly, func(c *cobra.Command) {
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
}, jsonOnly, func(c *cobra.Command) {
	c.Example = `  alpaca data meta exchanges`
})

var dataMetaConditionsCmd = fetchCmd("conditions <ticktype>", api.StockMetaConditionsOp, func(cmd *cobra.Command, args []string) (any, error) {
	return dataClient.StockMetaConditions(args[0], stockMetaConditionsParamsFromFlags(cmd))
}, jsonOnly, func(c *cobra.Command) {
	c.Args = cobra.ExactArgs(1)
	c.Example = `  alpaca data meta conditions trade`
})

// --- Screener ---

var screenerCmd = &cobra.Command{
	Use:   "screener",
	Short: "Stock and crypto screener and market movers",
}

var screenerMostActivesCmd = fetchCmd("most-actives", api.MostActivesOp, func(cmd *cobra.Command, args []string) (any, error) {
	return dataClient.MostActives(mostActivesParamsFromFlags(cmd))
}, func(c *cobra.Command) {
	c.Example = `  alpaca data screener most-actives
  alpaca data screener most-actives --by trades --top 10`
})

var screenerMoversCmd = fetchCmd("movers", api.MoversOp, func(cmd *cobra.Command, args []string) (any, error) {
	market := cmdutil.Str(cmd, "market")
	if market == "" {
		market = "stocks"
	}
	return dataClient.Movers(market, moversParamsFromFlags(cmd))
}, func(c *cobra.Command) {
	c.Example = `  alpaca data screener movers
  alpaca data screener movers --market crypto --top 5`
})

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

	return dataClient.News(params)
}, func(c *cobra.Command) {
	addPaginationFlags(c)
	c.Example = `  alpaca data news
  alpaca data news --symbols AAPL,MSFT --limit 10
  alpaca data news --symbols AAPL --all --max 100`
})

func init() {
	dataForexCmd.AddCommand(dataForexRatesCmd)
	dataForexCmd.AddCommand(dataForexLatestCmd)

	dataMetaCmd.AddCommand(dataMetaExchangesCmd)
	dataMetaCmd.AddCommand(dataMetaConditionsCmd)

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
