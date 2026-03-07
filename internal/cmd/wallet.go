package cmd

import (
	"github.com/alpacahq/cli/internal/api"
	"github.com/alpacahq/cli/internal/cmdutil"
	"github.com/spf13/cobra"
)

var walletCmd = &cobra.Command{
	Use:   "wallet",
	Short: "Crypto funding wallets and transfers",
}

var walletListCmd = fetchCmd("list", api.ListCryptoFundingWalletsOp, func(cmd *cobra.Command, args []string) (any, error) {
	return tradingClient.ListCryptoFundingWallets(listCryptoFundingWalletsParamsFromFlags(cmd))
})

// --- wallet transfer ---

var walletTransferCmd = &cobra.Command{
	Use:   "transfer",
	Short: "Manage crypto transfers",
}

var walletTransferListCmd = fetchCmd("list", api.ListCryptoFundingTransfersOp, func(cmd *cobra.Command, args []string) (any, error) {
	return tradingClient.ListCryptoFundingTransfers()
})

var walletTransferGetCmd = fetchCmd("get <id>", api.GetCryptoFundingTransferOp, func(cmd *cobra.Command, args []string) (any, error) {
	return tradingClient.GetCryptoFundingTransfer(args[0])
}, func(c *cobra.Command) {
	c.Args = cobra.ExactArgs(1)
})

var walletTransferCreateCmd = fetchCmd("create", api.CreateCryptoTransferForAccountOp, func(cmd *cobra.Command, args []string) (any, error) {
	if err := cmdutil.RequireAll(cmd, "amount", "address", "asset"); err != nil {
		return nil, err
	}
	body := &api.CreateCryptoTransferRequest{
		Amount:  cmdutil.Str(cmd, "amount"),
		Address: cmdutil.Str(cmd, "address"),
		Asset:   cmdutil.Str(cmd, "asset"),
	}
	return tradingClient.CreateCryptoTransferForAccount(body)
})

// --- transfer estimate ---

var walletTransferEstimateCmd = fetchCmd("estimate", api.GetCryptoTransferEstimateOp, func(cmd *cobra.Command, args []string) (any, error) {
	if err := cmdutil.RequireAll(cmd, "asset", "amount"); err != nil {
		return nil, err
	}
	return tradingClient.GetCryptoTransferEstimate(getCryptoTransferEstimateParamsFromFlags(cmd))
}, func(c *cobra.Command) {
	c.Example = `  alpaca wallet transfer estimate --asset BTC --amount 0.5 \
    --from-address 0xabc... --to-address 0xdef...`
})

// --- wallet whitelist ---

var walletWhitelistCmd = &cobra.Command{
	Use:   "whitelist",
	Short: "Manage whitelisted crypto addresses",
}

var walletWhitelistListCmd = fetchCmd("list", api.ListWhitelistedAddressOp, func(cmd *cobra.Command, args []string) (any, error) {
	return tradingClient.ListWhitelistedAddress()
})

var walletWhitelistAddCmd = fetchCmd("add", api.CreateWhitelistedAddressOp, func(cmd *cobra.Command, args []string) (any, error) {
	if err := cmdutil.RequireAll(cmd, "address", "asset"); err != nil {
		return nil, err
	}
	body := &api.CreateWhitelistedAddressRequest{
		Address: cmdutil.Str(cmd, "address"),
		Asset:   cmdutil.Str(cmd, "asset"),
	}
	return tradingClient.CreateWhitelistedAddress(body)
})

var walletWhitelistDeleteCmd = actionCmd("delete <id>", api.DeleteWhitelistedAddressOp, "Whitelisted address removed.", func(cmd *cobra.Command, args []string) error {
	_, err := tradingClient.DeleteWhitelistedAddress(args[0])
	return err
}, func(c *cobra.Command) {
	c.Args = cobra.ExactArgs(1)
})

func init() {
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
