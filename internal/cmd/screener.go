package cmd

import (
	"encoding/json"
	"net/url"

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
		params := url.Values{}
		by, _ := cmd.Flags().GetString("by")
		if by != "" {
			params.Set("by", by)
		}
		top, _ := cmd.Flags().GetString("top")
		if top != "" {
			params.Set("top", top)
		}

		data, err := apiClient.GetData("/v1beta1/screener/stocks/most-actives", params)
		if err != nil {
			return err
		}

		var resp map[string]json.RawMessage
		if json.Unmarshal(data, &resp) == nil {
			if actives, ok := resp["most_actives"]; ok {
				data = actives
			}
		}

		columns := []output.Column{
			{Header: "SYMBOL", Field: "symbol"},
			{Header: "VOLUME", Field: "volume"},
			{Header: "TRADE COUNT", Field: "trade_count"},
		}

		return output.Render(getOutput(), columns, data)
	},
}

var screenerMoversCmd = &cobra.Command{
	Use:   "movers",
	Short: "Top market movers (gainers and losers)",
	Example: `  alpaca screener movers
  alpaca screener movers --market crypto --top 5`,
	RunE: func(cmd *cobra.Command, args []string) error {
		market, _ := cmd.Flags().GetString("market")
		if market == "" {
			market = "stocks"
		}

		params := url.Values{}
		top, _ := cmd.Flags().GetString("top")
		if top != "" {
			params.Set("top", top)
		}

		data, err := apiClient.GetData("/v1beta1/screener/"+market+"/movers", params)
		if err != nil {
			return err
		}

		if getOutput() == "json" || getOutput() == "csv" {
			return output.JSON(cmd.OutOrStdout(), data)
		}

		var resp struct {
			Gainers []json.RawMessage `json:"gainers"`
			Losers  []json.RawMessage `json:"losers"`
		}
		if err := json.Unmarshal(data, &resp); err != nil {
			return output.JSON(cmd.OutOrStdout(), data)
		}

		moverColumns := []output.Column{
			{Header: "SYMBOL", Field: "symbol"},
			{Header: "CHANGE %", Field: "percent_change"},
			{Header: "CHANGE", Field: "change"},
			{Header: "PRICE", Field: "price"},
		}

		cmd.Println("GAINERS")
		gainersJSON, _ := json.Marshal(resp.Gainers)
		output.Render("table", moverColumns, gainersJSON)

		cmd.Println("\nLOSERS")
		losersJSON, _ := json.Marshal(resp.Losers)
		return output.Render("table", moverColumns, losersJSON)
	},
}

func init() {
	screenerMostActivesCmd.Flags().String("by", "", "Sort by: volume or trades (default: volume)")
	screenerMostActivesCmd.Flags().String("top", "", "Number of results (default: 10)")

	screenerMoversCmd.Flags().String("market", "", "Market: stocks or crypto (default: stocks)")
	screenerMoversCmd.Flags().String("top", "", "Number of results per direction (default: 10)")

	screenerCmd.AddCommand(screenerMostActivesCmd)
	screenerCmd.AddCommand(screenerMoversCmd)
}
