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
	Short: api.OperationSummary["listCryptoFundingWallets"],
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
	Short: api.OperationSummary["listCryptoFundingTransfers"],
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
	Short: api.OperationSummary["getCryptoFundingTransfer"],
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
	Short: api.OperationSummary["createCryptoTransferForAccount"],
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := cmdutil.RequireAll(cmd, "amount", "address", "asset"); err != nil {
			return err
		}
		warnLive()
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

// --- wallet whitelist ---

var walletWhitelistCmd = &cobra.Command{
	Use:   "whitelist",
	Short: "Manage whitelisted crypto addresses",
}

var walletWhitelistListCmd = &cobra.Command{
	Use:   "list",
	Short: api.OperationSummary["listWhitelistedAddress"],
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
	Short: api.OperationSummary["createWhitelistedAddress"],
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
	Short: api.OperationSummary["deleteWhitelistedAddress"],
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
	walletListCmd.Flags().String("asset", "", "Filter by crypto asset (e.g. BTC, ETH)")

	walletTransferCreateCmd.Flags().String("amount", "", "Amount to withdraw")
	walletTransferCreateCmd.Flags().String("address", "", "Destination address (must be whitelisted)")
	walletTransferCreateCmd.Flags().String("asset", "", "Crypto asset (e.g. BTC, ETH)")

	walletWhitelistAddCmd.Flags().String("address", "", "Crypto address to whitelist")
	walletWhitelistAddCmd.Flags().String("asset", "", "Crypto asset (e.g. BTC, ETH)")

	walletTransferCmd.AddCommand(walletTransferListCmd)
	walletTransferCmd.AddCommand(walletTransferGetCmd)
	walletTransferCmd.AddCommand(walletTransferCreateCmd)

	walletWhitelistCmd.AddCommand(walletWhitelistListCmd)
	walletWhitelistCmd.AddCommand(walletWhitelistAddCmd)
	walletWhitelistCmd.AddCommand(walletWhitelistDeleteCmd)

	walletCmd.AddCommand(walletListCmd)
	walletCmd.AddCommand(walletTransferCmd)
	walletCmd.AddCommand(walletWhitelistCmd)
}
