package cmd

import (
	"strings"

	"github.com/alpacahq/cli/internal/api"
	"github.com/alpacahq/cli/internal/cmdutil"
	"github.com/spf13/cobra"
)

var watchlistCmd = &cobra.Command{
	Use:   "watchlist",
	Short: "Manage watchlists",
}

var watchlistListCmd = fetchCmd("list", api.GetWatchlistsOp, func(cmd *cobra.Command, args []string) (any, error) {
	return tradingClient.GetWatchlists()
}, func(c *cobra.Command) {
	c.Example = `  alpaca watchlist list`
})

var watchlistGetCmd = fetchCmd("get <id>", api.GetWatchlistByIDOp, func(cmd *cobra.Command, args []string) (any, error) {
	return tradingClient.GetWatchlistByID(args[0])
}, func(c *cobra.Command) {
	c.Args = cobra.ExactArgs(1)
})

var watchlistCreateCmd = fetchCmd("create", api.PostWatchlistOp, func(cmd *cobra.Command, args []string) (any, error) {
	body, _ := updateWatchlistRequestBodyFromFlags(cmd)
	if symbols := cmdutil.Str(cmd, "symbols"); symbols != "" {
		body.Symbols = strings.Split(symbols, ",")
	}
	return tradingClient.PostWatchlist(body)
}, func(c *cobra.Command) {
	c.Example = `  alpaca watchlist create --name "Tech Stocks" --symbols AAPL,MSFT,GOOG`
})

var watchlistUpdateCmd = fetchCmd("update <id>", api.UpdateWatchlistByIDOp, func(cmd *cobra.Command, args []string) (any, error) {
	body, _ := updateWatchlistRequestBodyFromFlags(cmd)
	if symbols := cmdutil.Str(cmd, "symbols"); symbols != "" {
		body.Symbols = strings.Split(symbols, ",")
	}
	return tradingClient.UpdateWatchlistByID(args[0], body)
}, func(c *cobra.Command) {
	c.Args = cobra.ExactArgs(1)
})

var watchlistDeleteCmd = fetchCmd("delete <id>", api.DeleteWatchlistByIDOp, func(cmd *cobra.Command, args []string) (any, error) {
	return voidResponse(tradingClient.DeleteWatchlistByID(args[0]))
}, func(c *cobra.Command) {
	c.Args = cobra.ExactArgs(1)
})

var watchlistAddCmd = fetchCmd("add <id>", api.AddAssetToWatchlistOp, func(cmd *cobra.Command, args []string) (any, error) {
	return tradingClient.AddAssetToWatchlist(args[0], &api.AddAssetToWatchlistRequest{Symbol: cmdutil.Str(cmd, "symbol")})
}, func(c *cobra.Command) {
	c.Args = cobra.ExactArgs(1)
	c.Example = `  alpaca watchlist add <id> --symbol AAPL`
})

var watchlistRemoveCmd = fetchCmd("remove <id> <symbol>", api.RemoveAssetFromWatchlistOp, func(cmd *cobra.Command, args []string) (any, error) {
	return tradingClient.RemoveAssetFromWatchlist(args[0], args[1])
}, func(c *cobra.Command) {
	c.Args = cobra.ExactArgs(2)
	c.Example = `  alpaca watchlist remove <id> AAPL`
})

var watchlistGetByNameCmd = fetchCmd("get-by-name", api.GetWatchlistByNameOp, func(cmd *cobra.Command, args []string) (any, error) {
	return tradingClient.GetWatchlistByName(getWatchlistByNameParamsFromFlags(cmd))
}, func(c *cobra.Command) {
	c.Example = `  alpaca watchlist get-by-name --name "Tech Stocks"`
})

var watchlistDeleteByNameCmd = fetchCmd("delete-by-name", api.DeleteWatchlistByNameOp, func(cmd *cobra.Command, args []string) (any, error) {
	return voidResponse(tradingClient.DeleteWatchlistByName(deleteWatchlistByNameParamsFromFlags(cmd)))
}, func(c *cobra.Command) {
	c.Example = `  alpaca watchlist delete-by-name --name "Tech Stocks"`
})

var watchlistUpdateByNameCmd = fetchCmd("update-by-name", api.UpdateWatchlistByNameOp, func(cmd *cobra.Command, args []string) (any, error) {
	body, _ := updateWatchlistRequestBodyFromFlags(cmd)
	if cmdutil.Changed(cmd, "new-name") {
		body.Name = cmdutil.Str(cmd, "new-name")
	}
	if symbols := cmdutil.Str(cmd, "symbols"); symbols != "" {
		body.Symbols = strings.Split(symbols, ",")
	}
	return tradingClient.UpdateWatchlistByName(
		updateWatchlistByNameParamsFromFlags(cmd), body,
	)
}, func(c *cobra.Command) {
	c.Flags().String("new-name", "", "New name for the watchlist")
	c.Example = `  alpaca watchlist update-by-name --name "Tech Stocks" --symbols AAPL,MSFT,GOOG`
})

var watchlistAddByNameCmd = fetchCmd("add-by-name", api.AddAssetToWatchlistByNameOp, func(cmd *cobra.Command, args []string) (any, error) {
	return tradingClient.AddAssetToWatchlistByName(
		addAssetToWatchlistByNameParamsFromFlags(cmd),
		&api.AddAssetToWatchlistByNameRequest{Symbol: cmdutil.Str(cmd, "symbol")},
	)
}, func(c *cobra.Command) {
	c.Example = `  alpaca watchlist add-by-name --name "Tech Stocks" --symbol NVDA`
})

func init() {
	watchlistCmd.AddCommand(watchlistListCmd)
	watchlistCmd.AddCommand(watchlistGetCmd)
	watchlistCmd.AddCommand(watchlistCreateCmd)
	watchlistCmd.AddCommand(watchlistUpdateCmd)
	watchlistCmd.AddCommand(watchlistDeleteCmd)
	watchlistCmd.AddCommand(watchlistAddCmd)
	watchlistCmd.AddCommand(watchlistRemoveCmd)
	watchlistCmd.AddCommand(watchlistGetByNameCmd)
	watchlistCmd.AddCommand(watchlistDeleteByNameCmd)
	watchlistCmd.AddCommand(watchlistUpdateByNameCmd)
	watchlistCmd.AddCommand(watchlistAddByNameCmd)
}
