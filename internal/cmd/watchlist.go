package cmd

import (
	"fmt"
	"strings"

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
		data, err := apiClient.Get("/v2/watchlists", nil)
		if err != nil {
			return err
		}

		columns := []output.Column{
			{Header: "ID", Field: "id"},
			{Header: "NAME", Field: "name"},
			{Header: "CREATED", Field: "created_at"},
			{Header: "UPDATED", Field: "updated_at"},
		}

		return output.Render(getOutput(), columns, data)
	},
}

var watchlistGetCmd = &cobra.Command{
	Use:   "get <id>",
	Short: "Get watchlist details",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		data, err := apiClient.Get("/v2/watchlists/"+args[0], nil)
		if err != nil {
			return err
		}
		return output.JSON(cmd.OutOrStdout(), data)
	},
}

var watchlistCreateCmd = &cobra.Command{
	Use:   "create <name>",
	Short: "Create a new watchlist",
	Example: `  alpaca watchlist create "Tech Stocks" --symbols AAPL,MSFT,GOOG`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		body := map[string]any{
			"name": args[0],
		}
		symbols, _ := cmd.Flags().GetString("symbols")
		if symbols != "" {
			body["symbols"] = strings.Split(symbols, ",")
		}

		data, err := apiClient.Post("/v2/watchlists", body)
		if err != nil {
			return err
		}
		return output.JSON(cmd.OutOrStdout(), data)
	},
}

var watchlistUpdateCmd = &cobra.Command{
	Use:   "update <id>",
	Short: "Update a watchlist",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		body := map[string]any{}
		name, _ := cmd.Flags().GetString("name")
		if name != "" {
			body["name"] = name
		}
		symbols, _ := cmd.Flags().GetString("symbols")
		if symbols != "" {
			body["symbols"] = strings.Split(symbols, ",")
		}

		data, err := apiClient.Put("/v2/watchlists/"+args[0], body)
		if err != nil {
			return err
		}
		return output.JSON(cmd.OutOrStdout(), data)
	},
}

var watchlistDeleteCmd = &cobra.Command{
	Use:   "delete <id>",
	Short: "Delete a watchlist",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		_, err := apiClient.Delete("/v2/watchlists/"+args[0], nil)
		if err != nil {
			return err
		}
		fmt.Println("Watchlist deleted.")
		return nil
	},
}

var watchlistAddCmd = &cobra.Command{
	Use:   "add <id> <symbol>",
	Short: "Add a symbol to a watchlist",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		body := map[string]any{"symbol": args[1]}
		data, err := apiClient.Post("/v2/watchlists/"+args[0], body)
		if err != nil {
			return err
		}
		return output.JSON(cmd.OutOrStdout(), data)
	},
}

var watchlistRemoveCmd = &cobra.Command{
	Use:   "remove <id> <symbol>",
	Short: "Remove a symbol from a watchlist",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		_, err := apiClient.Delete("/v2/watchlists/"+args[0]+"/"+args[1], nil)
		if err != nil {
			return err
		}
		fmt.Printf("Removed %s from watchlist.\n", args[1])
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
