package cmd

import (
	"fmt"

	"github.com/alpacahq/cli/internal/api"
	"github.com/alpacahq/cli/internal/cmdutil"
	"github.com/alpacahq/cli/internal/output"
	"github.com/spf13/cobra"
)

var walletCmd = &cobra.Command{
	Use:   "wallet",
	Short: "Crypto funding wallets and transfers",
}

var walletListCmd = &cobra.Command{
	Use:   "list",
	Short: api.ListCryptoFundingWalletsOp.Summary,
	RunE: func(cmd *cobra.Command, args []string) error {
		wallet, err := tradingClient.ListCryptoFundingWallets(&api.ListCryptoFundingWalletsParams{
			Asset: cmdutil.Str(cmd, "asset"),
		})
		if err != nil {
			return err
		}
		return output.Render(cmd.OutOrStdout(), getOutput(), walletColumns(), wallet)
	},
}

// --- wallet transfer ---

var walletTransferCmd = &cobra.Command{
	Use:   "transfer",
	Short: "Manage crypto transfers",
}

var walletTransferListCmd = &cobra.Command{
	Use:   "list",
	Short: api.ListCryptoFundingTransfersOp.Summary,
	RunE: func(cmd *cobra.Command, args []string) error {
		transfers, err := tradingClient.ListCryptoFundingTransfers()
		if err != nil {
			return err
		}
		return output.Render(cmd.OutOrStdout(), getOutput(), transferColumns(), transfers)
	},
}

var walletTransferGetCmd = &cobra.Command{
	Use:   "get <id>",
	Short: api.GetCryptoFundingTransferOp.Summary,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		transfer, err := tradingClient.GetCryptoFundingTransfer(args[0])
		if err != nil {
			return err
		}
		return output.PrintSingle(cmd.OutOrStdout(), getOutput(), transferColumns(), transfer)
	},
}

var walletTransferCreateCmd = &cobra.Command{
	Use:   "create",
	Short: api.CreateCryptoTransferForAccountOp.Summary,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := cmdutil.RequireAll(cmd, "amount", "address", "asset"); err != nil {
			return err
		}
		body := &api.CreateCryptoTransferRequest{
			Amount:  cmdutil.Str(cmd, "amount"),
			Address: cmdutil.Str(cmd, "address"),
			Asset:   cmdutil.Str(cmd, "asset"),
		}

		transfer, err := tradingClient.CreateCryptoTransferForAccount(body)
		if err != nil {
			return err
		}
		return output.PrintSingle(cmd.OutOrStdout(), getOutput(), transferColumns(), transfer)
	},
}

// --- transfer estimate ---

var walletTransferEstimateCmd = &cobra.Command{
	Use:   "estimate",
	Short: api.GetCryptoTransferEstimateOp.Summary,
	Example: `  alpaca wallet transfer estimate --asset BTC --amount 0.5 \
    --from-address 0xabc... --to-address 0xdef...`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := cmdutil.RequireAll(cmd, "asset", "amount"); err != nil {
			return err
		}
		resp, err := tradingClient.GetCryptoTransferEstimate(&api.GetCryptoTransferEstimateParams{
			Asset:       cmdutil.Str(cmd, "asset"),
			Amount:      cmdutil.Str(cmd, "amount"),
			FromAddress: cmdutil.Str(cmd, "from-address"),
			ToAddress:   cmdutil.Str(cmd, "to-address"),
		})
		if err != nil {
			return err
		}
		return output.JSON(cmd.OutOrStdout(), resp)
	},
}

// --- wallet whitelist ---

var walletWhitelistCmd = &cobra.Command{
	Use:   "whitelist",
	Short: "Manage whitelisted crypto addresses",
}

var walletWhitelistListCmd = &cobra.Command{
	Use:   "list",
	Short: api.ListWhitelistedAddressOp.Summary,
	RunE: func(cmd *cobra.Command, args []string) error {
		addrs, err := tradingClient.ListWhitelistedAddress()
		if err != nil {
			return err
		}
		return output.Render(cmd.OutOrStdout(), getOutput(), whitelistColumns(), addrs)
	},
}

var walletWhitelistAddCmd = &cobra.Command{
	Use:   "add",
	Short: api.CreateWhitelistedAddressOp.Summary,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := cmdutil.RequireAll(cmd, "address", "asset"); err != nil {
			return err
		}
		body := &api.CreateWhitelistedAddressRequest{
			Address: cmdutil.Str(cmd, "address"),
			Asset:   cmdutil.Str(cmd, "asset"),
		}

		addr, err := tradingClient.CreateWhitelistedAddress(body)
		if err != nil {
			return err
		}
		return output.PrintSingle(cmd.OutOrStdout(), getOutput(), whitelistColumns(), addr)
	},
}

var walletWhitelistDeleteCmd = &cobra.Command{
	Use:   "delete <id>",
	Short: api.DeleteWhitelistedAddressOp.Summary,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		_, err := tradingClient.DeleteWhitelistedAddress(args[0])
		if err != nil {
			return err
		}
		_, _ = fmt.Fprintln(cmd.OutOrStdout(), "Whitelisted address removed.")
		return nil
	},
}

func init() {
	cmdutil.RegisterFlags(walletListCmd, api.ListCryptoFundingWalletsFlags, &cmdutil.FlagOpts{
		Exclude: map[string]bool{"network": true},
	})
	cmdutil.RegisterFlags(walletTransferCreateCmd, api.CreateCryptoTransferForAccountFlags, nil)
	cmdutil.RegisterFlags(walletTransferEstimateCmd, api.GetCryptoTransferEstimateFlags, nil)
	cmdutil.RegisterFlags(walletWhitelistAddCmd, api.CreateWhitelistedAddressFlags, nil)

	walletTransferCmd.AddCommand(walletTransferListCmd)
	walletTransferCmd.AddCommand(walletTransferGetCmd)
	walletTransferCmd.AddCommand(walletTransferCreateCmd)
	walletTransferCmd.AddCommand(walletTransferEstimateCmd)

	walletWhitelistCmd.AddCommand(walletWhitelistListCmd)
	walletWhitelistCmd.AddCommand(walletWhitelistAddCmd)
	walletWhitelistCmd.AddCommand(walletWhitelistDeleteCmd)

	walletCmd.AddCommand(walletListCmd)
	walletCmd.AddCommand(walletTransferCmd)
	walletCmd.AddCommand(walletWhitelistCmd)
}
