package cmd

import (
	"strings"

	"github.com/alpacahq/cli/internal/api"
	"github.com/alpacahq/cli/internal/cmdutil"
	"github.com/spf13/cobra"
)

var watchlistCreateCmd = fetchCmd("create", api.PostWatchlistOp, func(cmd *cobra.Command, args []string) (any, error) {
	body, _ := updateWatchlistRequestBodyFromFlags(cmd)
	if symbols := cmdutil.Str(cmd, "symbols"); symbols != "" {
		body.Symbols = strings.Split(symbols, ",")
	}
	return tradingClient.PostWatchlist(body)
}, func(c *cobra.Command) {
	c.Example = `  alpaca watchlist create --name "Tech Stocks" --symbols AAPL,MSFT,GOOG`
})

var watchlistUpdateCmd = fetchCmd("update", api.UpdateWatchlistByIDOp, func(cmd *cobra.Command, args []string) (any, error) {
	body, _ := updateWatchlistRequestBodyFromFlags(cmd)
	if symbols := cmdutil.Str(cmd, "symbols"); symbols != "" {
		body.Symbols = strings.Split(symbols, ",")
	}
	return tradingClient.UpdateWatchlistByID(cmdutil.Str(cmd, "watchlist-id"), body)
}, func(c *cobra.Command) {
	c.Example = `  alpaca watchlist update --watchlist-id <id> --name "Updated" --symbols AAPL,MSFT`
})

var watchlistAddCmd = fetchCmd("add", api.AddAssetToWatchlistOp, func(cmd *cobra.Command, args []string) (any, error) {
	return tradingClient.AddAssetToWatchlist(cmdutil.Str(cmd, "watchlist-id"), &api.AddAssetToWatchlistRequest{Symbol: cmdutil.Str(cmd, "symbol")})
}, func(c *cobra.Command) {
	c.Example = `  alpaca watchlist add --watchlist-id <id> --symbol AAPL`
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
	watchlistCmd.AddCommand(watchlistCreateCmd)
	watchlistCmd.AddCommand(watchlistUpdateCmd)
	watchlistCmd.AddCommand(watchlistAddCmd)
	watchlistCmd.AddCommand(watchlistUpdateByNameCmd)
	watchlistCmd.AddCommand(watchlistAddByNameCmd)
}
