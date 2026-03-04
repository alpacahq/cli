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
		return output.RenderWithHint(cmd.OutOrStdout(), getOutput(), watchlistColumns(), watchlists, "No watchlists. Create one with `alpaca watchlist create`.")
	},
}

var watchlistGetCmd = &cobra.Command{
	Use:   "get <id>",
	Short: api.GetWatchlistByIDOp.Summary,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		wl, err := tradingClient.GetWatchlistByID(args[0])
		if err != nil {
			return err
		}
		return output.PrintSingle(cmd.OutOrStdout(), getOutput(), watchlistColumns(), wl)
	},
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
		return output.PrintSingle(cmd.OutOrStdout(), getOutput(), watchlistColumns(), wl)
	},
}

var watchlistUpdateCmd = &cobra.Command{
	Use:   "update <id>",
	Short: api.UpdateWatchlistByIDOp.Summary,
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
		return output.PrintSingle(cmd.OutOrStdout(), getOutput(), watchlistColumns(), wl)
	},
}

var watchlistDeleteCmd = &cobra.Command{
	Use:   "delete <id>",
	Short: api.DeleteWatchlistByIDOp.Summary,
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
	Short: api.AddAssetToWatchlistOp.Summary,
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		wl, err := tradingClient.AddAssetToWatchlist(args[0], &api.AddAssetToWatchlistRequest{
			Symbol: args[1],
		})
		if err != nil {
			return err
		}
		return output.PrintSingle(cmd.OutOrStdout(), getOutput(), watchlistColumns(), wl)
	},
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
	RunE: func(cmd *cobra.Command, args []string) error {
		wl, err := tradingClient.GetWatchlistByName(&api.GetWatchlistByNameParams{
			Name: args[0],
		})
		if err != nil {
			return err
		}
		return output.PrintSingle(cmd.OutOrStdout(), getOutput(), watchlistColumns(), wl)
	},
}

var watchlistDeleteByNameCmd = &cobra.Command{
	Use:     "delete-by-name <name>",
	Short:   api.DeleteWatchlistByNameOp.Summary,
	Example: `  alpaca watchlist delete-by-name "Tech Stocks"`,
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		_, err := tradingClient.DeleteWatchlistByName(&api.DeleteWatchlistByNameParams{
			Name: args[0],
		})
		if err != nil {
			return err
		}
		_, _ = fmt.Fprintln(cmd.OutOrStdout(), "Watchlist deleted.")
		return nil
	},
}

var watchlistUpdateByNameCmd = &cobra.Command{
	Use:     "update-by-name <name>",
	Short:   api.UpdateWatchlistByNameOp.Summary,
	Example: `  alpaca watchlist update-by-name "Tech Stocks" --symbols AAPL,MSFT,GOOG`,
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		body := &api.UpdateWatchlistRequest{}
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
		return output.PrintSingle(cmd.OutOrStdout(), getOutput(), watchlistColumns(), wl)
	},
}

var watchlistAddByNameCmd = &cobra.Command{
	Use:     "add-by-name <name> <symbol>",
	Short:   api.AddAssetToWatchlistByNameOp.Summary,
	Example: `  alpaca watchlist add-by-name "Tech Stocks" NVDA`,
	Args:    cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		wl, err := tradingClient.AddAssetToWatchlistByName(
			&api.AddAssetToWatchlistByNameParams{Name: args[0]},
			&api.AddAssetToWatchlistByNameRequest{Symbol: args[1]},
		)
		if err != nil {
			return err
		}
		return output.PrintSingle(cmd.OutOrStdout(), getOutput(), watchlistColumns(), wl)
	},
}

func init() {
	cmdutil.RegisterFlags(watchlistCreateCmd, api.PostWatchlistFlags, &cmdutil.FlagOpts{
		Exclude: map[string]bool{"name": true},
	})
	cmdutil.RegisterFlags(watchlistUpdateCmd, api.UpdateWatchlistByIDFlags, &cmdutil.FlagOpts{
		Exclude: map[string]bool{"watchlist_id": true},
	})

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
