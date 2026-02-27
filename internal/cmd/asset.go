package cmd

import (
	"net/url"

	"github.com/alpacahq/cli/internal/output"
	"github.com/spf13/cobra"
)

var assetCmd = &cobra.Command{
	Use:   "asset",
	Short: "Browse assets",
}

var assetListCmd = &cobra.Command{
	Use:   "list",
	Short: "List available assets",
	Example: `  alpaca asset list
  alpaca asset list --class us_equity --status active
  alpaca asset list --exchange NYSE`,
	RunE: func(cmd *cobra.Command, args []string) error {
		params := url.Values{}
		status, _ := cmd.Flags().GetString("status")
		if status != "" {
			params.Set("status", status)
		}
		class, _ := cmd.Flags().GetString("class")
		if class != "" {
			params.Set("asset_class", class)
		}
		exchange, _ := cmd.Flags().GetString("exchange")
		if exchange != "" {
			params.Set("exchange", exchange)
		}

		data, err := apiClient.Get("/v2/assets", params)
		if err != nil {
			return err
		}

		columns := []output.Column{
			{Header: "SYMBOL", Field: "symbol"},
			{Header: "NAME", Field: "name"},
			{Header: "CLASS", Field: "class"},
			{Header: "EXCHANGE", Field: "exchange"},
			{Header: "STATUS", Field: "status"},
			{Header: "TRADABLE", Field: "tradable"},
			{Header: "SHORTABLE", Field: "shortable"},
			{Header: "FRACTIONABLE", Field: "fractionable"},
		}

		return output.Render(getOutput(), columns, data)
	},
}

var assetGetCmd = &cobra.Command{
	Use:   "get <symbol>",
	Short: "Get asset details",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		data, err := apiClient.Get("/v2/assets/"+args[0], nil)
		if err != nil {
			return err
		}

		columns := []output.Column{
			{Header: "SYMBOL", Field: "symbol"},
			{Header: "NAME", Field: "name"},
			{Header: "CLASS", Field: "class"},
			{Header: "EXCHANGE", Field: "exchange"},
			{Header: "STATUS", Field: "status"},
			{Header: "TRADABLE", Field: "tradable"},
			{Header: "SHORTABLE", Field: "shortable"},
			{Header: "FRACTIONABLE", Field: "fractionable"},
			{Header: "MARGINABLE", Field: "marginable"},
			{Header: "EASY TO BORROW", Field: "easy_to_borrow"},
		}

		return output.PrintSingle(getOutput(), columns, data)
	},
}

func init() {
	assetListCmd.Flags().String("status", "", "Filter: active or inactive")
	assetListCmd.Flags().String("class", "", "Asset class: us_equity, crypto")
	assetListCmd.Flags().String("exchange", "", "Exchange: NYSE, NASDAQ, etc.")

	assetCmd.AddCommand(assetListCmd)
	assetCmd.AddCommand(assetGetCmd)
}
