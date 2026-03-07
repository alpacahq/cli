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
	Short: api.GetV2AssetsOp.Summary(),
	Example: `  alpaca asset list
  alpaca asset list --asset-class us_equity --status active
  alpaca asset list --exchange NYSE`,
	RunE: func(cmd *cobra.Command, args []string) error {
		params := getV2AssetsParamsFromFlags(cmd)
		assets, err := tradingClient.GetV2Assets(params)
		if err != nil {
			return err
		}
		return output.Render(cmd.OutOrStdout(), getOutput(), nil, assets)
	},
}

var assetGetCmd = &cobra.Command{
	Use:   "get <symbol>",
	Short: api.GetV2AssetsSymbolOrAssetIDOp.Summary(),
	Example: `  alpaca asset get AAPL
  alpaca asset get BTC/USD
  alpaca asset get AAPL --json`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		asset, err := tradingClient.GetV2AssetsSymbolOrAssetID(args[0])
		if err != nil {
			return err
		}
		return output.PrintSingle(cmd.OutOrStdout(), getOutput(), nil, asset)
	},
}

var treasuryListCmd = &cobra.Command{
	Use:   "treasury",
	Short: api.UsTreasuriesOp.Summary(),
	Example: `  alpaca asset treasury
  alpaca asset treasury --bond-status active`,
	RunE: func(cmd *cobra.Command, args []string) error {
		resp, err := tradingClient.UsTreasuries(usTreasuriesParamsFromFlags(cmd))
		if err != nil {
			return err
		}
		return output.Render(cmd.OutOrStdout(), getOutput(), nil, resp.UsTreasuries)
	},
}

var bondListCmd = &cobra.Command{
	Use:   "bond",
	Short: api.UsCorporatesOp.Summary(),
	Example: `  alpaca asset bond
  alpaca asset bond --bond-status active`,
	RunE: func(cmd *cobra.Command, args []string) error {
		resp, err := tradingClient.UsCorporates(usCorporatesParamsFromFlags(cmd))
		if err != nil {
			return err
		}
		return output.Render(cmd.OutOrStdout(), getOutput(), nil, resp.UsCorporates)
	},
}

func init() {
	cmdutil.RegisterFlags(assetListCmd, api.GetV2AssetsOp.Flags(), nil)

	cmdutil.RegisterFlags(treasuryListCmd, api.UsTreasuriesOp.Flags(), nil)

	cmdutil.RegisterFlags(bondListCmd, api.UsCorporatesOp.Flags(), nil)

	assetCmd.AddCommand(assetListCmd)
	assetCmd.AddCommand(assetGetCmd)
	assetCmd.AddCommand(treasuryListCmd)
	assetCmd.AddCommand(bondListCmd)
}
