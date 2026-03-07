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
	Short: api.GetWatchlistsOp.Summary,
	RunE: func(cmd *cobra.Command, args []string) error {
		watchlists, err := tradingClient.GetWatchlists()
		if err != nil {
			return err
		}
		return output.RenderWithHint(cmd.OutOrStdout(), getOutput(), columnsForOp(api.GetWatchlistsOp), watchlists, "No watchlists. Create one with `alpaca watchlist create`.")
	},
}

// watchlistFetchRunE builds a RunE that fetches a watchlist by key and prints it.
func watchlistFetchRunE(fetch func(key string) (*api.Watchlist, error)) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		wl, err := fetch(args[0])
		if err != nil {
			return err
		}
		return output.PrintSingle(cmd.OutOrStdout(), getOutput(), columnsForOp(api.GetWatchlistByIDOp), wl)
	}
}

// watchlistDeleteRunE builds a RunE that deletes a watchlist by key.
func watchlistDeleteRunE(del func(key string) error) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		if err := del(args[0]); err != nil {
			return err
		}
		_, _ = fmt.Fprintln(cmd.OutOrStdout(), "Watchlist deleted.")
		return nil
	}
}

// watchlistAddRunE builds a RunE that adds a symbol to a watchlist by key.
func watchlistAddRunE(add func(key, symbol string) (*api.Watchlist, error)) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		wl, err := add(args[0], args[1])
		if err != nil {
			return err
		}
		return output.PrintSingle(cmd.OutOrStdout(), getOutput(), columnsForOp(api.AddAssetToWatchlistOp), wl)
	}
}

var watchlistGetCmd = &cobra.Command{
	Use:   "get <id>",
	Short: api.GetWatchlistByIDOp.Summary,
	Args:  cobra.ExactArgs(1),
	RunE: watchlistFetchRunE(func(key string) (*api.Watchlist, error) {
		return tradingClient.GetWatchlistByID(key)
	}),
}

var watchlistCreateCmd = &cobra.Command{
	Use:     "create <name>",
	Short:   api.PostWatchlistOp.Summary,
	Example: `  alpaca watchlist create "Tech Stocks" --symbols AAPL,MSFT,GOOG`,
	Args:    cobra.ExactArgs(1),
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
		return output.PrintSingle(cmd.OutOrStdout(), getOutput(), columnsForOp(api.PostWatchlistOp), wl)
	},
}

var watchlistUpdateCmd = &cobra.Command{
	Use:   "update <id>",
	Short: api.UpdateWatchlistByIDOp.Summary,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		body, _ := updateWatchlistRequestBodyFromFlags(cmd)
		symbols := cmdutil.Str(cmd, "symbols")
		if symbols != "" {
			body.Symbols = strings.Split(symbols, ",")
		}

		wl, err := tradingClient.UpdateWatchlistByID(args[0], body)
		if err != nil {
			return err
		}
		return output.PrintSingle(cmd.OutOrStdout(), getOutput(), columnsForOp(api.UpdateWatchlistByIDOp), wl)
	},
}

var watchlistDeleteCmd = &cobra.Command{
	Use:   "delete <id>",
	Short: api.DeleteWatchlistByIDOp.Summary,
	Args:  cobra.ExactArgs(1),
	RunE: watchlistDeleteRunE(func(key string) error {
		_, err := tradingClient.DeleteWatchlistByID(key)
		return err
	}),
}

var watchlistAddCmd = &cobra.Command{
	Use:   "add <id> <symbol>",
	Short: api.AddAssetToWatchlistOp.Summary,
	Args:  cobra.ExactArgs(2),
	RunE: watchlistAddRunE(func(key, symbol string) (*api.Watchlist, error) {
		return tradingClient.AddAssetToWatchlist(key, &api.AddAssetToWatchlistRequest{Symbol: symbol})
	}),
}

var watchlistRemoveCmd = &cobra.Command{
	Use:   "remove <id> <symbol>",
	Short: api.RemoveAssetFromWatchlistOp.Summary,
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

var watchlistGetByNameCmd = &cobra.Command{
	Use:     "get-by-name <name>",
	Short:   api.GetWatchlistByNameOp.Summary,
	Example: `  alpaca watchlist get-by-name "Tech Stocks"`,
	Args:    cobra.ExactArgs(1),
	RunE: watchlistFetchRunE(func(key string) (*api.Watchlist, error) {
		return tradingClient.GetWatchlistByName(&api.GetWatchlistByNameParams{Name: key})
	}),
}

var watchlistDeleteByNameCmd = &cobra.Command{
	Use:     "delete-by-name <name>",
	Short:   api.DeleteWatchlistByNameOp.Summary,
	Example: `  alpaca watchlist delete-by-name "Tech Stocks"`,
	Args:    cobra.ExactArgs(1),
	RunE: watchlistDeleteRunE(func(key string) error {
		_, err := tradingClient.DeleteWatchlistByName(&api.DeleteWatchlistByNameParams{Name: key})
		return err
	}),
}

var watchlistUpdateByNameCmd = &cobra.Command{
	Use:     "update-by-name <name>",
	Short:   api.UpdateWatchlistByNameOp.Summary,
	Example: `  alpaca watchlist update-by-name "Tech Stocks" --symbols AAPL,MSFT,GOOG`,
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		body, _ := updateWatchlistRequestBodyFromFlags(cmd)
		if cmdutil.Changed(cmd, "new-name") {
			body.Name = cmdutil.Str(cmd, "new-name")
		}
		symbols := cmdutil.Str(cmd, "symbols")
		if symbols != "" {
			body.Symbols = strings.Split(symbols, ",")
		}

		wl, err := tradingClient.UpdateWatchlistByName(
			&api.UpdateWatchlistByNameParams{Name: args[0]}, body,
		)
		if err != nil {
			return err
		}
		return output.PrintSingle(cmd.OutOrStdout(), getOutput(), columnsForOp(api.UpdateWatchlistByNameOp), wl)
	},
}

var watchlistAddByNameCmd = &cobra.Command{
	Use:     "add-by-name <name> <symbol>",
	Short:   api.AddAssetToWatchlistByNameOp.Summary,
	Example: `  alpaca watchlist add-by-name "Tech Stocks" NVDA`,
	Args:    cobra.ExactArgs(2),
	RunE: watchlistAddRunE(func(key, symbol string) (*api.Watchlist, error) {
		return tradingClient.AddAssetToWatchlistByName(
			&api.AddAssetToWatchlistByNameParams{Name: key},
			&api.AddAssetToWatchlistByNameRequest{Symbol: symbol},
		)
	}),
}

func init() {
	cmdutil.RegisterFlags(watchlistCreateCmd, api.PostWatchlistFlags, &cmdutil.FlagOpts{
		Exclude: map[string]bool{"name": true},
	})
	cmdutil.RegisterFlags(watchlistUpdateCmd, api.UpdateWatchlistByIDFlags, nil)

	cmdutil.RegisterFlags(watchlistUpdateByNameCmd, api.UpdateWatchlistByNameFlags, &cmdutil.FlagOpts{
		Exclude: map[string]bool{"name": true},
	})
	watchlistUpdateByNameCmd.Flags().String("new-name", "", "New name for the watchlist")

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
