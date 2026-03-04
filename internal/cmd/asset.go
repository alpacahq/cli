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
	Short: api.GetV2AssetsOp.Summary,
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
		return output.Render(cmd.OutOrStdout(), getOutput(), assetListColumns(), assets)
	},
}

var assetGetCmd = &cobra.Command{
	Use:   "get <symbol>",
	Short: api.GetV2AssetsSymbolOrAssetIDOp.Summary,
	Example: `  alpaca asset get AAPL
  alpaca asset get BTC/USD
  alpaca asset get AAPL --json`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		asset, err := tradingClient.GetV2AssetsSymbolOrAssetID(args[0])
		if err != nil {
			return err
		}
		return output.PrintSingle(cmd.OutOrStdout(), getOutput(), assetDetailColumns(), asset)
	},
}

var treasuryListCmd = &cobra.Command{
	Use:   "treasury",
	Short: api.UsTreasuriesOp.Summary,
	Example: `  alpaca asset treasury
  alpaca asset treasury --status active`,
	RunE: func(cmd *cobra.Command, args []string) error {
		resp, err := tradingClient.UsTreasuries(&api.UsTreasuriesParams{
			BondStatus: cmdutil.Str(cmd, "status"),
			Cusips:     cmdutil.Str(cmd, "cusips"),
		})
		if err != nil {
			return err
		}
		return output.Render(cmd.OutOrStdout(), getOutput(), treasuryColumns(), resp.UsTreasuries)
	},
}

var bondListCmd = &cobra.Command{
	Use:   "bond",
	Short: api.UsCorporatesOp.Summary,
	Example: `  alpaca asset bond
  alpaca asset bond --status active`,
	RunE: func(cmd *cobra.Command, args []string) error {
		resp, err := tradingClient.UsCorporates(&api.UsCorporatesParams{
			BondStatus: cmdutil.Str(cmd, "status"),
			Cusips:     cmdutil.Str(cmd, "cusips"),
		})
		if err != nil {
			return err
		}
		return output.Render(cmd.OutOrStdout(), getOutput(), bondColumns(), resp.UsCorporates)
	},
}

func init() {
	cmdutil.RegisterFlags(assetListCmd, api.GetV2AssetsFlags, &cmdutil.FlagOpts{
		Exclude: map[string]bool{"attributes": true},
		Aliases: map[string]string{"asset_class": "class"},
	})

	cmdutil.RegisterFlags(treasuryListCmd, api.UsTreasuriesFlags, &cmdutil.FlagOpts{
		Exclude: map[string]bool{"isins": true, "subtype": true},
		Aliases: map[string]string{"bond_status": "status"},
	})

	cmdutil.RegisterFlags(bondListCmd, api.UsCorporatesFlags, &cmdutil.FlagOpts{
		Exclude: map[string]bool{"isins": true, "tickers": true},
		Aliases: map[string]string{"bond_status": "status"},
	})

	assetCmd.AddCommand(assetListCmd)
	assetCmd.AddCommand(assetGetCmd)
	assetCmd.AddCommand(treasuryListCmd)
	assetCmd.AddCommand(bondListCmd)
}
