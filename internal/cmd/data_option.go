package cmd

import (
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
	Short: api.OptionBarsOp.Summary,
	Example: `  alpaca data option bars --symbols AAPL250620C00200000 --start 2025-01-01
  alpaca data option bars --symbols AAPL250620C00200000,AAPL250620P00200000 --timeframe 1Day`,
	RunE: func(cmd *cobra.Command, args []string) error {
		_, err := cmdutil.RequireStr(cmd, "symbols")
		if err != nil {
			return err
		}

		resp, err := dataClient.OptionBars(optionBarsParamsFromFlags(cmd))
		if err != nil {
			return err
		}

		return renderMapValues(cmd.OutOrStdout(), getOutput(), barColumns(), resp.Bars)
	},
}

var dataOptionTradesCmd = &cobra.Command{
	Use:     "trades",
	Short:   api.OptionTradesOp.Summary,
	Example: `  alpaca data option trades --symbols AAPL250620C00200000 --start 2025-01-01`,
	RunE: func(cmd *cobra.Command, args []string) error {
		_, err := cmdutil.RequireStr(cmd, "symbols")
		if err != nil {
			return err
		}

		resp, err := dataClient.OptionTrades(optionTradesParamsFromFlags(cmd))
		if err != nil {
			return err
		}

		return renderMapValues(cmd.OutOrStdout(), getOutput(), tradeColumns(), resp.Trades)
	},
}

var dataOptionSnapshotCmd = &cobra.Command{
	Use:   "snapshot",
	Short: api.OptionSnapshotsOp.Summary,
	Example: `  alpaca data option snapshot --symbols AAPL250620C00200000
  alpaca data option snapshot --symbols AAPL250620C00200000,AAPL250620P00200000`,
	RunE: func(cmd *cobra.Command, args []string) error {
		_, err := cmdutil.RequireStr(cmd, "symbols")
		if err != nil {
			return err
		}

		resp, err := dataClient.OptionSnapshots(optionSnapshotsParamsFromFlags(cmd))
		if err != nil {
			return err
		}

		return output.JSON(cmd.OutOrStdout(), resp.Snapshots)
	},
}

var dataOptionChainCmd = &cobra.Command{
	Use:   "chain <underlying>",
	Short: api.OptionChainOp.Summary,
	Example: `  alpaca data option chain AAPL
  alpaca data option chain SPY --expiration-date 2025-06-20 --type call`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		resp, err := dataClient.OptionChain(args[0], optionChainParamsFromFlags(cmd))
		if err != nil {
			return err
		}

		return output.JSON(cmd.OutOrStdout(), resp.Snapshots)
	},
}

var dataOptionLatestQuotesCmd = &cobra.Command{
	Use:     "latest-quotes",
	Short:   api.OptionLatestQuotesOp.Summary,
	Example: `  alpaca data option latest-quotes --symbols AAPL250620C00200000`,
	RunE: func(cmd *cobra.Command, args []string) error {
		_, err := cmdutil.RequireStr(cmd, "symbols")
		if err != nil {
			return err
		}

		resp, err := dataClient.OptionLatestQuotes(optionLatestQuotesParamsFromFlags(cmd))
		if err != nil {
			return err
		}

		return output.JSON(cmd.OutOrStdout(), resp.Quotes)
	},
}

var dataOptionLatestTradesCmd = &cobra.Command{
	Use:     "latest-trades",
	Short:   api.OptionLatestTradesOp.Summary,
	Example: `  alpaca data option latest-trades --symbols AAPL250620C00200000`,
	RunE: func(cmd *cobra.Command, args []string) error {
		_, err := cmdutil.RequireStr(cmd, "symbols")
		if err != nil {
			return err
		}

		resp, err := dataClient.OptionLatestTrades(optionLatestTradesParamsFromFlags(cmd))
		if err != nil {
			return err
		}

		return output.JSON(cmd.OutOrStdout(), resp.Trades)
	},
}

var dataOptionExchangesCmd = &cobra.Command{
	Use:     "exchanges",
	Short:   api.OptionMetaExchangesOp.Summary,
	Example: `  alpaca data option exchanges`,
	RunE: func(cmd *cobra.Command, args []string) error {
		resp, err := dataClient.OptionMetaExchanges()
		if err != nil {
			return err
		}
		return output.JSON(cmd.OutOrStdout(), resp)
	},
}

var dataOptionConditionsCmd = &cobra.Command{
	Use:     "conditions <ticktype>",
	Short:   api.OptionMetaConditionsOp.Summary,
	Example: `  alpaca data option conditions trade`,
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		resp, err := dataClient.OptionMetaConditions(args[0])
		if err != nil {
			return err
		}
		return output.JSON(cmd.OutOrStdout(), resp)
	},
}

func init() {
	cmdutil.RegisterFlags(dataOptionBarsCmd, api.OptionBarsFlags, &cmdutil.FlagOpts{
		Defaults: map[string]string{"timeframe": "1Day"},
	})

	cmdutil.RegisterFlags(dataOptionTradesCmd, api.OptionTradesFlags, nil)

	cmdutil.RegisterFlags(dataOptionSnapshotCmd, api.OptionSnapshotsFlags, nil)

	cmdutil.RegisterFlags(dataOptionChainCmd, api.OptionChainFlags, nil)

	cmdutil.RegisterFlags(dataOptionLatestQuotesCmd, api.OptionLatestQuotesFlags, nil)
	cmdutil.RegisterFlags(dataOptionLatestTradesCmd, api.OptionLatestTradesFlags, nil)

	dataOptionCmd.AddCommand(dataOptionBarsCmd)
	dataOptionCmd.AddCommand(dataOptionTradesCmd)
	dataOptionCmd.AddCommand(dataOptionSnapshotCmd)
	dataOptionCmd.AddCommand(dataOptionChainCmd)
	dataOptionCmd.AddCommand(dataOptionLatestQuotesCmd)
	dataOptionCmd.AddCommand(dataOptionLatestTradesCmd)
	dataOptionCmd.AddCommand(dataOptionExchangesCmd)
	dataOptionCmd.AddCommand(dataOptionConditionsCmd)
}
