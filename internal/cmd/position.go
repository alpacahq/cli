package cmd

import (
	"fmt"
	"net/url"

	"github.com/alpacahq/cli/internal/output"
	"github.com/spf13/cobra"
)

var positionCmd = &cobra.Command{
	Use:   "position",
	Short: "Manage positions",
}

var positionListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all open positions",
	RunE: func(cmd *cobra.Command, args []string) error {
		data, err := apiClient.Get("/v2/positions", nil)
		if err != nil {
			return err
		}

		return output.Render(getOutput(), positionColumns(), data)
	},
}

var positionGetCmd = &cobra.Command{
	Use:   "get <symbol>",
	Short: "Get position for a symbol",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		data, err := apiClient.Get("/v2/positions/"+args[0], nil)
		if err != nil {
			return err
		}
		return output.PrintSingle(getOutput(), positionColumns(), data)
	},
}

var positionCloseCmd = &cobra.Command{
	Use:   "close <symbol>",
	Short: "Close a position",
	Example: `  alpaca position close AAPL
  alpaca position close AAPL --qty 5
  alpaca position close AAPL --pct 50`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		params := url.Values{}
		qty, _ := cmd.Flags().GetString("qty")
		pct, _ := cmd.Flags().GetString("pct")

		if qty != "" {
			params.Set("qty", qty)
		}
		if pct != "" {
			params.Set("percentage", pct)
		}

		data, err := apiClient.Delete("/v2/positions/"+args[0], params)
		if err != nil {
			return err
		}

		return output.PrintSingle(getOutput(), orderColumns(), data)
	},
}

var positionCloseAllCmd = &cobra.Command{
	Use:   "close-all",
	Short: "Close all open positions",
	RunE: func(cmd *cobra.Command, args []string) error {
		cancelOrders, _ := cmd.Flags().GetBool("cancel-orders")
		params := url.Values{}
		if cancelOrders {
			params.Set("cancel_orders", "true")
		}

		data, err := apiClient.Delete("/v2/positions", params)
		if err != nil {
			return err
		}

		fmt.Println("All positions closed.")
		return output.JSON(cmd.OutOrStdout(), data)
	},
}

func init() {
	positionCloseCmd.Flags().String("qty", "", "Number of shares to close")
	positionCloseCmd.Flags().String("pct", "", "Percentage of position to close (0-100)")

	positionCloseAllCmd.Flags().Bool("cancel-orders", false, "Also cancel all open orders")

	positionCmd.AddCommand(positionListCmd)
	positionCmd.AddCommand(positionGetCmd)
	positionCmd.AddCommand(positionCloseCmd)
	positionCmd.AddCommand(positionCloseAllCmd)
}

func positionColumns() []output.Column {
	return []output.Column{
		{Header: "SYMBOL", Field: "symbol"},
		{Header: "QTY", Field: "qty"},
		{Header: "SIDE", Field: "side"},
		{Header: "AVG ENTRY", Field: "avg_entry_price"},
		{Header: "CURRENT", Field: "current_price"},
		{Header: "MKT VALUE", Field: "market_value"},
		{Header: "P/L", Field: "unrealized_pl", Format: func(v any) string {
			return output.DollarPL(fmt.Sprintf("%v", v))
		}},
		{Header: "P/L %", Field: "unrealized_plpc", Format: func(v any) string {
			return output.PercentPL(fmt.Sprintf("%v", v))
		}},
	}
}
