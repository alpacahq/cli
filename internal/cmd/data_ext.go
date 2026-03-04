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
	Short: api.RatesOp.Summary,
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
			PageToken:     cmdutil.Str(cmd, "page-token"),
		})
		if err != nil {
			return err
		}

		return renderMapValues(cmd.OutOrStdout(), getOutput(), forexRateColumns(), resp.Rates)
	},
}

var dataForexLatestCmd = &cobra.Command{
	Use:     "latest",
	Short:   api.LatestRatesOp.Summary,
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
	Use:     "crypto-orderbook",
	Short:   api.CryptoLatestOrderbooksOp.Summary,
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
	Short: api.StockAuctionsOp.Summary,
	Example: `  alpaca data auctions --symbols AAPL --start 2025-01-01
  alpaca data auctions --symbols AAPL,MSFT --limit 10`,
	RunE: func(cmd *cobra.Command, args []string) error {
		symbols, err := cmdutil.RequireStr(cmd, "symbols")
		if err != nil {
			return err
		}

		resp, err := dataClient.StockAuctions(&api.StockAuctionsParams{
			Symbols:   symbols,
			Start:     cmdutil.Str(cmd, "start"),
			End:       cmdutil.Str(cmd, "end"),
			Limit:     cmdutil.Int(cmd, "limit"),
			Sort:      cmdutil.Str(cmd, "sort"),
			Asof:      cmdutil.Str(cmd, "asof"),
			Feed:      cmdutil.Str(cmd, "feed"),
			Currency:  cmdutil.Str(cmd, "currency"),
			PageToken: cmdutil.Str(cmd, "page-token"),
		})
		if err != nil {
			return err
		}

		return output.JSON(cmd.OutOrStdout(), resp.Auctions)
	},
}

// --- Corporate Actions (market data) ---

var dataCorporateActionsCmd = &cobra.Command{
	Use:     "corporate-actions",
	Short:   api.CorporateActionsOp.Summary,
	Example: `  alpaca data corporate-actions --symbols AAPL --types forward_split --start 2025-01-01`,
	RunE: func(cmd *cobra.Command, args []string) error {
		resp, err := dataClient.CorporateActions(&api.CorporateActionsParams{
			Symbols:   cmdutil.Str(cmd, "symbols"),
			Cusips:    cmdutil.Str(cmd, "cusips"),
			Types:     cmdutil.Str(cmd, "types"),
			Start:     cmdutil.Str(cmd, "start"),
			End:       cmdutil.Str(cmd, "end"),
			Ids:       cmdutil.Str(cmd, "ids"),
			Limit:     cmdutil.Int(cmd, "limit"),
			Sort:      cmdutil.Str(cmd, "sort"),
			PageToken: cmdutil.Str(cmd, "page-token"),
		})
		if err != nil {
			return err
		}

		return output.JSON(cmd.OutOrStdout(), resp)
	},
}

// --- Fixed Income Data ---

var dataFixedIncomeCmd = &cobra.Command{
	Use:     "fixed-income",
	Short:   api.FixedIncomeLatestPricesOp.Summary,
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

// --- Logo ---

var dataLogoCmd = &cobra.Command{
	Use:     "logo <symbol>",
	Short:   api.LogosOp.Summary,
	Example: `  alpaca data logo AAPL`,
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		resp, err := dataClient.Logos(args[0], &api.LogosParams{
			Placeholder: cmdutil.Bool(cmd, "placeholder"),
		})
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
	Short:   api.StockMetaExchangesOp.Summary,
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
	Short:   api.StockMetaConditionsOp.Summary,
	Example: `  alpaca data meta conditions trade`,
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		resp, err := dataClient.StockMetaConditions(args[0], &api.StockMetaConditionsParams{
			Tape: cmdutil.Str(cmd, "tape"),
		})
		if err != nil {
			return err
		}
		return output.JSON(cmd.OutOrStdout(), resp)
	},
}

func init() {
	cmdutil.RegisterFlags(dataForexRatesCmd, api.RatesFlags, &cmdutil.FlagOpts{
		Aliases: map[string]string{"currency_pairs": "pairs"},
	})
	cmdutil.RegisterFlags(dataForexLatestCmd, api.LatestRatesFlags, &cmdutil.FlagOpts{
		Aliases: map[string]string{"currency_pairs": "pairs"},
	})
	dataForexCmd.AddCommand(dataForexRatesCmd)
	dataForexCmd.AddCommand(dataForexLatestCmd)

	cmdutil.RegisterFlags(dataCryptoOrderbookCmd, api.CryptoLatestOrderbooksFlags, nil)

	cmdutil.RegisterFlags(dataAuctionsCmd, api.StockAuctionsFlags, nil)

	cmdutil.RegisterFlags(dataCorporateActionsCmd, api.CorporateActionsFlags, nil)

	cmdutil.RegisterFlags(dataFixedIncomeCmd, api.FixedIncomeLatestPricesFlags, &cmdutil.FlagOpts{
		Aliases: map[string]string{"isins": "symbols"},
	})

	cmdutil.RegisterFlags(dataLogoCmd, api.LogosFlags, nil)

	cmdutil.RegisterFlags(dataMetaConditionsCmd, api.StockMetaConditionsFlags, nil)
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
	if json.Unmarshal(j, &m) == nil && len(m) == 1 {
		for _, v := range m {
			return output.Render(w, format, cols, v)
		}
	}
	return output.JSON(w, data)
}
