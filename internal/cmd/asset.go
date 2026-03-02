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

var treasuryListCmd = &cobra.Command{
	Use:   "treasury",
	Short: "List US Treasury bonds",
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
		return output.Render(getOutput(), treasuryColumns(), resp.UsTreasuries)
	},
}

var bondListCmd = &cobra.Command{
	Use:   "bond",
	Short: "List US Corporate bonds",
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
		return output.Render(getOutput(), bondColumns(), resp.UsCorporates)
	},
}

func init() {
	assetListCmd.Flags().String("status", "", "Filter: active or inactive")
	_ = assetListCmd.RegisterFlagCompletionFunc("status", cobra.FixedCompletions(api.AssetsStatusValues, cobra.ShellCompDirectiveNoFileComp))
	assetListCmd.Flags().String("class", "", "Asset class: us_equity, crypto")
	_ = assetListCmd.RegisterFlagCompletionFunc("class", cobra.FixedCompletions([]string{"us_equity", "crypto", "us_option", "fixed_income"}, cobra.ShellCompDirectiveNoFileComp))
	assetListCmd.Flags().String("exchange", "", "Exchange: NYSE, NASDAQ, etc.")
	_ = assetListCmd.RegisterFlagCompletionFunc("exchange", cobra.FixedCompletions([]string{"NYSE", "NASDAQ", "AMEX", "ARCA", "BATS", "OTC", "FTXU", "CBSE", "ERSX"}, cobra.ShellCompDirectiveNoFileComp))

	treasuryListCmd.Flags().String("status", "", "Bond status: active or inactive")
	_ = treasuryListCmd.RegisterFlagCompletionFunc("status", cobra.FixedCompletions(api.AssetsStatusValues, cobra.ShellCompDirectiveNoFileComp))
	treasuryListCmd.Flags().String("cusips", "", "Filter by CUSIPs (comma-separated)")
	bondListCmd.Flags().String("status", "", "Bond status: active or inactive")
	_ = bondListCmd.RegisterFlagCompletionFunc("status", cobra.FixedCompletions(api.AssetsStatusValues, cobra.ShellCompDirectiveNoFileComp))
	bondListCmd.Flags().String("cusips", "", "Filter by CUSIPs (comma-separated)")

	assetCmd.AddCommand(assetListCmd)
	assetCmd.AddCommand(assetGetCmd)
	assetCmd.AddCommand(treasuryListCmd)
	assetCmd.AddCommand(bondListCmd)
}
