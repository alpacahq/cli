package cmd

import (
	"fmt"
	"strings"

	"github.com/alpacahq/cli/internal/api"
	"github.com/alpacahq/cli/internal/cmdutil"
	"github.com/alpacahq/cli/internal/output"
	"github.com/spf13/cobra"
)

var watchlistCmd = &cobra.Command{
	Use:   "watchlist",
	Short: "Manage watchlists",
}

var watchlistListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all watchlists",
	RunE: func(cmd *cobra.Command, args []string) error {
		watchlists, err := tradingClient.GetWatchlists()
		if err != nil {
			return err
		}
		return output.Render(getOutput(), watchlistColumns(), watchlists)
	},
}

var watchlistGetCmd = &cobra.Command{
	Use:   "get <id>",
	Short: "Get watchlist details",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		wl, err := tradingClient.GetWatchlistByID(args[0])
		if err != nil {
			return err
		}
		return output.JSON(cmd.OutOrStdout(), wl)
	},
}

var watchlistCreateCmd = &cobra.Command{
	Use:   "create <name>",
	Short: "Create a new watchlist",
	Example: `  alpaca watchlist create "Tech Stocks" --symbols AAPL,MSFT,GOOG`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		body := &api.UpdateWatchlistRequest{
			Name: args[0],
		}
		symbols := cmdutil.Str(cmd, "symbols")
		if symbols != "" {
			body.Symbols = strings.Split(symbols, ",")
		}

		wl, err := tradingClient.PostWatchlist(body)
		if err != nil {
			return err
		}
		return output.JSON(cmd.OutOrStdout(), wl)
	},
}

var watchlistUpdateCmd = &cobra.Command{
	Use:   "update <id>",
	Short: "Update a watchlist",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		body := &api.UpdateWatchlistRequest{}
		if cmdutil.Changed(cmd, "name") {
			body.Name = cmdutil.Str(cmd, "name")
		}
		symbols := cmdutil.Str(cmd, "symbols")
		if symbols != "" {
			body.Symbols = strings.Split(symbols, ",")
		}

		wl, err := tradingClient.UpdateWatchlistByID(args[0], body)
		if err != nil {
			return err
		}
		return output.JSON(cmd.OutOrStdout(), wl)
	},
}

var watchlistDeleteCmd = &cobra.Command{
	Use:   "delete <id>",
	Short: "Delete a watchlist",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		_, err := tradingClient.DeleteWatchlistByID(args[0])
		if err != nil {
			return err
		}
		_, _ = fmt.Fprintln(cmd.OutOrStdout(), "Watchlist deleted.")
		return nil
	},
}

var watchlistAddCmd = &cobra.Command{
	Use:   "add <id> <symbol>",
	Short: "Add a symbol to a watchlist",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		wl, err := tradingClient.AddAssetToWatchlist(args[0], &api.AddAssetToWatchlistRequest{
			Symbol: args[1],
		})
		if err != nil {
			return err
		}
		return output.JSON(cmd.OutOrStdout(), wl)
	},
}

var watchlistRemoveCmd = &cobra.Command{
	Use:   "remove <id> <symbol>",
	Short: "Remove a symbol from a watchlist",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		_, err := tradingClient.RemoveAssetFromWatchlist(args[0], args[1])
		if err != nil {
			return err
		}
		_, _ = fmt.Fprintf(cmd.OutOrStdout(), "Removed %s from watchlist.\n", args[1])
		return nil
	},
}

func init() {
	watchlistCreateCmd.Flags().String("symbols", "", "Comma-separated symbols to add")
	watchlistUpdateCmd.Flags().String("name", "", "New watchlist name")
	watchlistUpdateCmd.Flags().String("symbols", "", "New comma-separated symbols list")

	watchlistCmd.AddCommand(watchlistListCmd)
	watchlistCmd.AddCommand(watchlistGetCmd)
	watchlistCmd.AddCommand(watchlistCreateCmd)
	watchlistCmd.AddCommand(watchlistUpdateCmd)
	watchlistCmd.AddCommand(watchlistDeleteCmd)
	watchlistCmd.AddCommand(watchlistAddCmd)
	watchlistCmd.AddCommand(watchlistRemoveCmd)
}
