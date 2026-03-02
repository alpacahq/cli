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
	Short: "List crypto funding wallets",
	RunE: func(cmd *cobra.Command, args []string) error {
		wallet, err := tradingClient.ListCryptoFundingWallets(&api.ListCryptoFundingWalletsParams{
			Asset: cmdutil.Str(cmd, "asset"),
		})
		if err != nil {
			return err
		}
		return output.Render(getOutput(), walletColumns(), wallet)
	},
}

var walletTransfersCmd = &cobra.Command{
	Use:   "transfers",
	Short: "List crypto funding transfers",
	RunE: func(cmd *cobra.Command, args []string) error {
		transfers, err := tradingClient.ListCryptoFundingTransfers()
		if err != nil {
			return err
		}
		return output.Render(getOutput(), transferColumns(), transfers)
	},
}

var walletTransferGetCmd = &cobra.Command{
	Use:   "transfer-get <id>",
	Short: "Get a specific crypto transfer",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		transfer, err := tradingClient.GetCryptoFundingTransfer(args[0])
		if err != nil {
			return err
		}
		return output.PrintSingle(getOutput(), transferColumns(), transfer)
	},
}

var walletTransferCmd = &cobra.Command{
	Use:   "transfer",
	Short: "Create a crypto transfer",
	RunE: func(cmd *cobra.Command, args []string) error {
		body := &api.CreateCryptoTransferRequest{
			Amount:  cmdutil.Str(cmd, "amount"),
			Address: cmdutil.Str(cmd, "address"),
			Asset:   cmdutil.Str(cmd, "asset"),
		}
		if body.Amount == "" || body.Address == "" || body.Asset == "" {
			return fmt.Errorf("--amount, --address, and --asset are all required")
		}

		transfer, err := tradingClient.CreateCryptoTransferForAccount(body)
		if err != nil {
			return err
		}
		return output.PrintSingle(getOutput(), transferColumns(), transfer)
	},
}

var walletWhitelistListCmd = &cobra.Command{
	Use:   "whitelist",
	Short: "List whitelisted crypto addresses",
	RunE: func(cmd *cobra.Command, args []string) error {
		addrs, err := tradingClient.ListWhitelistedAddress()
		if err != nil {
			return err
		}
		return output.Render(getOutput(), whitelistColumns(), addrs)
	},
}

var walletWhitelistAddCmd = &cobra.Command{
	Use:   "whitelist-add",
	Short: "Add a whitelisted crypto address",
	RunE: func(cmd *cobra.Command, args []string) error {
		body := &api.CreateWhitelistedAddressRequest{
			Address: cmdutil.Str(cmd, "address"),
			Asset:   cmdutil.Str(cmd, "asset"),
		}
		if body.Address == "" || body.Asset == "" {
			return fmt.Errorf("--address and --asset are required")
		}

		addr, err := tradingClient.CreateWhitelistedAddress(body)
		if err != nil {
			return err
		}
		return output.PrintSingle(getOutput(), whitelistColumns(), addr)
	},
}

var walletWhitelistDeleteCmd = &cobra.Command{
	Use:   "whitelist-delete <id>",
	Short: "Remove a whitelisted crypto address",
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

	walletTransferCmd.Flags().String("amount", "", "Transfer amount")
	walletTransferCmd.Flags().String("address", "", "Destination address")
	walletTransferCmd.Flags().String("asset", "", "Crypto asset (e.g. BTC, ETH)")

	walletWhitelistAddCmd.Flags().String("address", "", "Crypto address to whitelist")
	walletWhitelistAddCmd.Flags().String("asset", "", "Crypto asset (e.g. BTC, ETH)")

	walletCmd.AddCommand(walletListCmd)
	walletCmd.AddCommand(walletTransfersCmd)
	walletCmd.AddCommand(walletTransferGetCmd)
	walletCmd.AddCommand(walletTransferCmd)
	walletCmd.AddCommand(walletWhitelistListCmd)
	walletCmd.AddCommand(walletWhitelistAddCmd)
	walletCmd.AddCommand(walletWhitelistDeleteCmd)
}
