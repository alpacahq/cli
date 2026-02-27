package cmd

import (
	"encoding/json"

	"github.com/alpacahq/cli/internal/api"
	"github.com/alpacahq/cli/internal/cmdutil"
	"github.com/alpacahq/cli/internal/output"
	"github.com/spf13/cobra"
)

var screenerCmd = &cobra.Command{
	Use:   "screener",
	Short: "Stock screener and market movers",
}

var screenerMostActivesCmd = &cobra.Command{
	Use:   "most-actives",
	Short: "Most active stocks by volume or trade count",
	Example: `  alpaca screener most-actives
  alpaca screener most-actives --by trades --top 10`,
	RunE: func(cmd *cobra.Command, args []string) error {
		params := &api.MostActivesParams{
			By:  cmdutil.Str(cmd, "by"),
			Top: cmdutil.Int(cmd, "top"),
		}

		resp, err := dataClient.MostActives(params)
		if err != nil {
			return err
		}

		return output.Render(getOutput(), screenerMostActivesColumns(), resp.MostActives)
	},
}

var screenerMoversCmd = &cobra.Command{
	Use:   "movers",
	Short: "Top market movers (gainers and losers)",
	Example: `  alpaca screener movers
  alpaca screener movers --market crypto --top 5`,
	RunE: func(cmd *cobra.Command, args []string) error {
		market := cmdutil.Str(cmd, "market")
		if market == "" {
			market = "stocks"
		}

		params := &api.MoversParams{
			Top: cmdutil.Int(cmd, "top"),
		}

		resp, err := dataClient.Movers(market, params)
		if err != nil {
			return err
		}

		if getOutput() == "json" || getOutput() == "csv" {
			return output.Render(getOutput(), screenerMoverColumns(), resp)
		}

		cmd.Println("GAINERS")
		gainersJSON, _ := json.Marshal(resp.Gainers)
		if err := output.Render("table", screenerMoverColumns(), json.RawMessage(gainersJSON)); err != nil {
			return err
		}

		cmd.Println("\nLOSERS")
		losersJSON, _ := json.Marshal(resp.Losers)
		return output.Render("table", screenerMoverColumns(), json.RawMessage(losersJSON))
	},
}

func init() {
	screenerMostActivesCmd.Flags().String("by", "", "Sort by: volume or trades (default: volume)")
	screenerMostActivesCmd.Flags().Int("top", 0, "Number of results (default: 10)")

	screenerMoversCmd.Flags().String("market", "", "Market: stocks or crypto (default: stocks)")
	screenerMoversCmd.Flags().Int("top", 0, "Number of results per direction (default: 10)")

	screenerCmd.AddCommand(screenerMostActivesCmd)
	screenerCmd.AddCommand(screenerMoversCmd)
}
