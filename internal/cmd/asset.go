package cmd

import (
	"github.com/alpacahq/cli/internal/api"
	"github.com/alpacahq/cli/internal/cmdutil"
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
		params := &api.GetV2AssetsParams{
			Status:     cmdutil.Str(cmd, "status"),
			AssetClass: cmdutil.Str(cmd, "class"),
			Exchange:   cmdutil.Str(cmd, "exchange"),
		}

		assets, err := tradingClient.GetV2Assets(params)
		if err != nil {
			return err
		}
		return output.Render(getOutput(), assetListColumns(), assets)
	},
}

var assetGetCmd = &cobra.Command{
	Use:   "get <symbol>",
	Short: "Get asset details",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		asset, err := tradingClient.GetV2AssetsSymbolOrAssetID(args[0])
		if err != nil {
			return err
		}
		return output.PrintSingle(getOutput(), assetDetailColumns(), asset)
	},
}

func init() {
	assetListCmd.Flags().String("status", "", "Filter: active or inactive")
	assetListCmd.Flags().String("class", "", "Asset class: us_equity, crypto")
	assetListCmd.Flags().String("exchange", "", "Exchange: NYSE, NASDAQ, etc.")

	assetCmd.AddCommand(assetListCmd)
	assetCmd.AddCommand(assetGetCmd)
}
