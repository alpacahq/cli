package cmd

import (
	"fmt"

	"github.com/alpacahq/cli/internal/api"
	"github.com/alpacahq/cli/internal/cmdutil"
	"github.com/spf13/cobra"
)

var removeWatchlistAssetByNameCmd = &cobra.Command{
	Use:     "remove-by-name",
	Short:   "Delete symbol from watchlist by name",
	Example: `  alpaca watchlist remove-by-name --name "Tech Stocks" --symbol AAPL`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := cmdutil.RequireAll(cmd, "name", "symbol"); err != nil {
			return err
		}

		wl, err := tradingClient.GetWatchlistByName(queryFromFlags(cmd, api.GetWatchlistByNameOp))
		if err != nil {
			return err
		}
		if wl == nil || wl.ID == "" {
			return fmt.Errorf("watchlist %q not found", cmdutil.Str(cmd, "name"))
		}

		updated, err := tradingClient.RemoveAssetFromWatchlist(wl.ID, cmdutil.Str(cmd, "symbol"))
		if err != nil {
			return err
		}
		return renderData(cmd, updated)
	},
}

func init() {
	removeWatchlistAssetByNameCmd.Annotations = map[string]string{
		"op": api.RemoveAssetFromWatchlistOp.Name,
	}
	removeWatchlistAssetByNameCmd.Flags().String("name", "", "name of the watchlist")
	removeWatchlistAssetByNameCmd.Flags().String("symbol", "", "symbol name to remove from the watchlist")
	watchlistCmd.AddCommand(removeWatchlistAssetByNameCmd)
}
