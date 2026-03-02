package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/alpacahq/cli/internal/api"
	"github.com/alpacahq/cli/internal/cmdutil"
	"github.com/alpacahq/cli/internal/output"
	"github.com/spf13/cobra"
)

var dataOptionCmd = &cobra.Command{
	Use:   "option",
	Short: "Options market data",
}

var dataOptionBarsCmd = &cobra.Command{
	Use:   "bars",
	Short: "Get historical option bars",
	Example: `  alpaca data option bars --symbols AAPL250620C00200000 --start 2025-01-01
  alpaca data option bars --symbols AAPL250620C00200000,AAPL250620P00200000 --timeframe 1Day`,
	RunE: func(cmd *cobra.Command, args []string) error {
		symbols := cmdutil.Str(cmd, "symbols")
		if symbols == "" {
			return fmt.Errorf("--symbols is required")
		}

		resp, err := dataClient.OptionBars(&api.OptionBarsParams{
			Symbols:   symbols,
			Timeframe: cmdutil.Str(cmd, "timeframe"),
			Start:     cmdutil.Str(cmd, "start"),
			End:       cmdutil.Str(cmd, "end"),
			Limit:     cmdutil.Int(cmd, "limit"),
			Sort:      cmdutil.Str(cmd, "sort"),
		})
		if err != nil {
			return err
		}

		return renderOptionMap(getOutput(), barColumns(), resp.Bars)
	},
}

var dataOptionTradesCmd = &cobra.Command{
	Use:   "trades",
	Short: "Get historical option trades",
	Example: `  alpaca data option trades --symbols AAPL250620C00200000 --start 2025-01-01`,
	RunE: func(cmd *cobra.Command, args []string) error {
		symbols := cmdutil.Str(cmd, "symbols")
		if symbols == "" {
			return fmt.Errorf("--symbols is required")
		}

		resp, err := dataClient.OptionTrades(&api.OptionTradesParams{
			Symbols: symbols,
			Start:   cmdutil.Str(cmd, "start"),
			End:     cmdutil.Str(cmd, "end"),
			Limit:   cmdutil.Int(cmd, "limit"),
			Sort:    cmdutil.Str(cmd, "sort"),
		})
		if err != nil {
			return err
		}

		return renderOptionMap(getOutput(), tradeColumns(), resp.Trades)
	},
}

var dataOptionSnapshotCmd = &cobra.Command{
	Use:   "snapshot",
	Short: "Get option snapshots",
	Example: `  alpaca data option snapshot --symbols AAPL250620C00200000
  alpaca data option snapshot --symbols AAPL250620C00200000,AAPL250620P00200000`,
	RunE: func(cmd *cobra.Command, args []string) error {
		symbols := cmdutil.Str(cmd, "symbols")
		if symbols == "" {
			return fmt.Errorf("--symbols is required")
		}

		resp, err := dataClient.OptionSnapshots(&api.OptionSnapshotsParams{
			Symbols: symbols,
			Feed:    cmdutil.Str(cmd, "feed"),
		})
		if err != nil {
			return err
		}

		return output.JSON(cmd.OutOrStdout(), resp.Snapshots)
	},
}

var dataOptionChainCmd = &cobra.Command{
	Use:   "chain <underlying>",
	Short: "Get option chain snapshots with greeks and pricing",
	Example: `  alpaca data option chain AAPL
  alpaca data option chain SPY --expiry 2025-06-20 --type call`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		resp, err := dataClient.OptionChain(args[0], &api.OptionChainParams{
			Feed:              cmdutil.Str(cmd, "feed"),
			ExpirationDate:    cmdutil.Str(cmd, "expiry"),
			ExpirationDateGte: cmdutil.Str(cmd, "expiry-gte"),
			ExpirationDateLte: cmdutil.Str(cmd, "expiry-lte"),
			StrikePriceGte:    cmdutil.Float64(cmd, "strike-gte"),
			StrikePriceLte:    cmdutil.Float64(cmd, "strike-lte"),
			RootSymbol:        cmdutil.Str(cmd, "root-symbol"),
			Type:              cmdutil.Str(cmd, "type"),
			Limit:             cmdutil.Int(cmd, "limit"),
		})
		if err != nil {
			return err
		}

		return output.JSON(cmd.OutOrStdout(), resp.Snapshots)
	},
}

var dataOptionLatestQuotesCmd = &cobra.Command{
	Use:   "latest-quotes",
	Short: "Get latest option quotes",
	Example: `  alpaca data option latest-quotes --symbols AAPL250620C00200000`,
	RunE: func(cmd *cobra.Command, args []string) error {
		symbols := cmdutil.Str(cmd, "symbols")
		if symbols == "" {
			return fmt.Errorf("--symbols is required")
		}

		resp, err := dataClient.OptionLatestQuotes(&api.OptionLatestQuotesParams{
			Symbols: symbols,
			Feed:    cmdutil.Str(cmd, "feed"),
		})
		if err != nil {
			return err
		}

		return output.JSON(cmd.OutOrStdout(), resp.Quotes)
	},
}

var dataOptionLatestTradesCmd = &cobra.Command{
	Use:   "latest-trades",
	Short: "Get latest option trades",
	Example: `  alpaca data option latest-trades --symbols AAPL250620C00200000`,
	RunE: func(cmd *cobra.Command, args []string) error {
		symbols := cmdutil.Str(cmd, "symbols")
		if symbols == "" {
			return fmt.Errorf("--symbols is required")
		}

		resp, err := dataClient.OptionLatestTrades(&api.OptionLatestTradesParams{
			Symbols: symbols,
			Feed:    cmdutil.Str(cmd, "feed"),
		})
		if err != nil {
			return err
		}

		return output.JSON(cmd.OutOrStdout(), resp.Trades)
	},
}

func init() {
	dataOptionBarsCmd.Flags().String("symbols", "", "Option symbols (comma-separated, required)")
	dataOptionBarsCmd.Flags().String("timeframe", "1Day", "Timeframe: 1Min, 5Min, 15Min, 1Hour, 1Day, 1Week, 1Month")
	dataOptionBarsCmd.Flags().String("start", "", "Start date (YYYY-MM-DD or RFC3339)")
	dataOptionBarsCmd.Flags().String("end", "", "End date")
	dataOptionBarsCmd.Flags().Int("limit", 0, "Max results")
	dataOptionBarsCmd.Flags().String("sort", "", "Sort: asc or desc")

	dataOptionTradesCmd.Flags().String("symbols", "", "Option symbols (comma-separated, required)")
	dataOptionTradesCmd.Flags().String("start", "", "Start date")
	dataOptionTradesCmd.Flags().String("end", "", "End date")
	dataOptionTradesCmd.Flags().Int("limit", 0, "Max results")
	dataOptionTradesCmd.Flags().String("sort", "", "Sort: asc or desc")

	dataOptionSnapshotCmd.Flags().String("symbols", "", "Option symbols (comma-separated, required)")
	dataOptionSnapshotCmd.Flags().String("feed", "", "Feed: indicative or opra")

	dataOptionChainCmd.Flags().String("feed", "", "Feed: indicative or opra")
	dataOptionChainCmd.Flags().String("expiry", "", "Exact expiration date (YYYY-MM-DD)")
	dataOptionChainCmd.Flags().String("expiry-gte", "", "Expiration on or after")
	dataOptionChainCmd.Flags().String("expiry-lte", "", "Expiration on or before")
	dataOptionChainCmd.Flags().Float64("strike-gte", 0, "Min strike price")
	dataOptionChainCmd.Flags().Float64("strike-lte", 0, "Max strike price")
	dataOptionChainCmd.Flags().String("root-symbol", "", "Root symbol")
	dataOptionChainCmd.Flags().String("type", "", "Option type: call or put")
	dataOptionChainCmd.Flags().Int("limit", 0, "Max results")

	dataOptionLatestQuotesCmd.Flags().String("symbols", "", "Option symbols (comma-separated, required)")
	dataOptionLatestQuotesCmd.Flags().String("feed", "", "Feed: indicative or opra")

	dataOptionLatestTradesCmd.Flags().String("symbols", "", "Option symbols (comma-separated, required)")
	dataOptionLatestTradesCmd.Flags().String("feed", "", "Feed: indicative or opra")

	dataOptionCmd.AddCommand(dataOptionBarsCmd)
	dataOptionCmd.AddCommand(dataOptionTradesCmd)
	dataOptionCmd.AddCommand(dataOptionSnapshotCmd)
	dataOptionCmd.AddCommand(dataOptionChainCmd)
	dataOptionCmd.AddCommand(dataOptionLatestQuotesCmd)
	dataOptionCmd.AddCommand(dataOptionLatestTradesCmd)
}

func renderOptionMap(format string, cols []output.Column, data any) error {
	j, _ := json.Marshal(data)
	var m map[string]json.RawMessage
	if json.Unmarshal(j, &m) == nil && len(m) == 1 {
		for _, v := range m {
			return output.Render(format, cols, v)
		}
	}
	return output.JSON(os.Stdout, data)
}
