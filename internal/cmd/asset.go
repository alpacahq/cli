package cmd

import (
	"github.com/alpacahq/cli/internal/api"
	"github.com/alpacahq/cli/internal/cmdutil"
	"github.com/spf13/cobra"
)

var assetCmd = &cobra.Command{
	Use:   "asset",
	Short: "Browse assets",
}

var assetListCmd = fetchCmd("list", api.GetV2AssetsOp, func(cmd *cobra.Command, args []string) (any, error) {
	return tradingClient.GetV2Assets(getV2AssetsParamsFromFlags(cmd))
}, func(c *cobra.Command) {
	c.Example = `  alpaca asset list
  alpaca asset list --asset-class us_equity --status active
  alpaca asset list --exchange NYSE`
})

var assetGetCmd = fetchCmd("get", api.GetV2AssetsSymbolOrAssetIDOp, func(cmd *cobra.Command, args []string) (any, error) {
	return tradingClient.GetV2AssetsSymbolOrAssetID(cmdutil.Str(cmd, "symbol-or-asset-id"))
}, func(c *cobra.Command) {
	c.Example = `  alpaca asset get --symbol-or-asset-id AAPL
  alpaca asset get --symbol-or-asset-id BTC/USD`
})

var treasuryListCmd = fetchCmd("treasury", api.UsTreasuriesOp, func(cmd *cobra.Command, args []string) (any, error) {
	return tradingClient.UsTreasuries(usTreasuriesParamsFromFlags(cmd))
}, func(c *cobra.Command) {
	c.Example = `  alpaca asset treasury
  alpaca asset treasury --bond-status active`
})

var bondListCmd = fetchCmd("bond", api.UsCorporatesOp, func(cmd *cobra.Command, args []string) (any, error) {
	return tradingClient.UsCorporates(usCorporatesParamsFromFlags(cmd))
}, func(c *cobra.Command) {
	c.Example = `  alpaca asset bond
  alpaca asset bond --bond-status active`
})

func init() {
	assetCmd.AddCommand(assetListCmd)
	assetCmd.AddCommand(assetGetCmd)
	assetCmd.AddCommand(treasuryListCmd)
	assetCmd.AddCommand(bondListCmd)
}
