package cmd

import (
	"fmt"
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

var watchlistCreateCmd = fetchCmd("create <name>", api.PostWatchlistOp, func(cmd *cobra.Command, args []string) (any, error) {
	body := &api.UpdateWatchlistRequest{
		Name: args[0],
	}
	symbols := cmdutil.Str(cmd, "symbols")
	if symbols != "" {
		body.Symbols = strings.Split(symbols, ",")
	}
	return tradingClient.PostWatchlist(body)
}, func(c *cobra.Command) {
	cmdutil.RegisterFlags(c, api.PostWatchlistOp.Flags(), &cmdutil.FlagOpts{
		Exclude: map[string]bool{"name": true},
	})
	c.Args = cobra.ExactArgs(1)
	c.Example = `  alpaca watchlist create "Tech Stocks" --symbols AAPL,MSFT,GOOG`
})

var watchlistUpdateCmd = fetchCmd("update <id>", api.UpdateWatchlistByIDOp, func(cmd *cobra.Command, args []string) (any, error) {
	body, _ := updateWatchlistRequestBodyFromFlags(cmd)
	symbols := cmdutil.Str(cmd, "symbols")
	if symbols != "" {
		body.Symbols = strings.Split(symbols, ",")
	}
	return tradingClient.UpdateWatchlistByID(args[0], body)
}, func(c *cobra.Command) {
	c.Args = cobra.ExactArgs(1)
})

var watchlistDeleteCmd = actionCmd("delete <id>", api.DeleteWatchlistByIDOp, "Watchlist deleted.", func(cmd *cobra.Command, args []string) error {
	_, err := tradingClient.DeleteWatchlistByID(args[0])
	return err
}, func(c *cobra.Command) {
	c.Args = cobra.ExactArgs(1)
})

var watchlistAddCmd = fetchCmd("add <id> <symbol>", api.AddAssetToWatchlistOp, func(cmd *cobra.Command, args []string) (any, error) {
	return tradingClient.AddAssetToWatchlist(args[0], &api.AddAssetToWatchlistRequest{Symbol: args[1]})
}, func(c *cobra.Command) {
	c.Args = cobra.ExactArgs(2)
})

var watchlistRemoveCmd = actionCmd("remove <id> <symbol>", api.RemoveAssetFromWatchlistOp, "", func(cmd *cobra.Command, args []string) error {
	_, err := tradingClient.RemoveAssetFromWatchlist(args[0], args[1])
	if err != nil {
		return err
	}
	_, _ = fmt.Fprintf(cmd.OutOrStdout(), "Removed %s from watchlist.\n", args[1])
	return nil
}, func(c *cobra.Command) {
	c.Args = cobra.ExactArgs(2)
	c.Example = `  alpaca watchlist remove <id> AAPL`
})

var watchlistGetByNameCmd = fetchCmd("get-by-name <name>", api.GetWatchlistByNameOp, func(cmd *cobra.Command, args []string) (any, error) {
	return tradingClient.GetWatchlistByName(&api.GetWatchlistByNameParams{Name: args[0]})
}, func(c *cobra.Command) {
	cmdutil.RegisterFlags(c, api.GetWatchlistByNameOp.Flags(), &cmdutil.FlagOpts{
		Exclude: map[string]bool{"name": true},
	})
	c.Args = cobra.ExactArgs(1)
	c.Example = `  alpaca watchlist get-by-name "Tech Stocks"`
})

var watchlistDeleteByNameCmd = actionCmd("delete-by-name <name>", api.DeleteWatchlistByNameOp, "Watchlist deleted.", func(cmd *cobra.Command, args []string) error {
	_, err := tradingClient.DeleteWatchlistByName(&api.DeleteWatchlistByNameParams{Name: args[0]})
	return err
}, func(c *cobra.Command) {
	cmdutil.RegisterFlags(c, api.DeleteWatchlistByNameOp.Flags(), &cmdutil.FlagOpts{
		Exclude: map[string]bool{"name": true},
	})
	c.Args = cobra.ExactArgs(1)
	c.Example = `  alpaca watchlist delete-by-name "Tech Stocks"`
})

var watchlistUpdateByNameCmd = fetchCmd("update-by-name <name>", api.UpdateWatchlistByNameOp, func(cmd *cobra.Command, args []string) (any, error) {
	body, _ := updateWatchlistRequestBodyFromFlags(cmd)
	if cmdutil.Changed(cmd, "new-name") {
		body.Name = cmdutil.Str(cmd, "new-name")
	}
	symbols := cmdutil.Str(cmd, "symbols")
	if symbols != "" {
		body.Symbols = strings.Split(symbols, ",")
	}
	return tradingClient.UpdateWatchlistByName(
		&api.UpdateWatchlistByNameParams{Name: args[0]}, body,
	)
}, func(c *cobra.Command) {
	cmdutil.RegisterFlags(c, api.UpdateWatchlistByNameOp.Flags(), &cmdutil.FlagOpts{
		Exclude: map[string]bool{"name": true},
	})
	c.Flags().String("new-name", "", "New name for the watchlist")
	c.Args = cobra.ExactArgs(1)
	c.Example = `  alpaca watchlist update-by-name "Tech Stocks" --symbols AAPL,MSFT,GOOG`
})

var watchlistAddByNameCmd = fetchCmd("add-by-name <name> <symbol>", api.AddAssetToWatchlistByNameOp, func(cmd *cobra.Command, args []string) (any, error) {
	return tradingClient.AddAssetToWatchlistByName(
		&api.AddAssetToWatchlistByNameParams{Name: args[0]},
		&api.AddAssetToWatchlistByNameRequest{Symbol: args[1]},
	)
}, func(c *cobra.Command) {
	cmdutil.RegisterFlags(c, api.AddAssetToWatchlistByNameOp.Flags(), &cmdutil.FlagOpts{
		Exclude: map[string]bool{"name": true, "symbol": true},
	})
	c.Args = cobra.ExactArgs(2)
	c.Example = `  alpaca watchlist add-by-name "Tech Stocks" NVDA`
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
