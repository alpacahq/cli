package cmd

import (
	"fmt"

	"github.com/alpacahq/cli/internal/api"
	"github.com/alpacahq/cli/internal/cmdutil"
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
		positions, err := tradingClient.GetAllOpenPositions()
		if err != nil {
			return err
		}
		return output.Render(cmd.OutOrStdout(), getOutput(), positionColumns(), positions)
	},
}

var positionGetCmd = &cobra.Command{
	Use:   "get <symbol>",
	Short: "Get position for a symbol",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		pos, err := tradingClient.GetOpenPosition(args[0])
		if err != nil {
			return err
		}
		return output.PrintSingle(cmd.OutOrStdout(), getOutput(), positionColumns(), pos)
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
		warnLive()
		params := &api.DeleteOpenPositionParams{
			Qty:        cmdutil.Float64(cmd, "qty"),
			Percentage: cmdutil.Float64(cmd, "pct"),
		}

		order, err := tradingClient.DeleteOpenPosition(args[0], params)
		if err != nil {
			return err
		}
		return output.PrintSingle(cmd.OutOrStdout(), getOutput(), orderColumns(), order)
	},
}

var positionCloseAllCmd = &cobra.Command{
	Use:   "close-all",
	Short: "Close all open positions",
	RunE: func(cmd *cobra.Command, args []string) error {
		warnLive()
		cancelled, err := tradingClient.DeleteAllOpenPositions(&api.DeleteAllOpenPositionsParams{
			CancelOrders: cmdutil.Bool(cmd, "cancel-orders"),
		})
		if err != nil {
			return err
		}
		if getOutput() == "json" || getOutput() == "csv" {
			return output.JSON(cmd.OutOrStdout(), cancelled)
		}
		_, _ = fmt.Fprintln(cmd.OutOrStdout(), "All positions closed.")
		return nil
	},
}

func init() {
	positionCloseCmd.Flags().Float64("qty", 0, "Number of shares to close")
	positionCloseCmd.Flags().Float64("pct", 0, "Percentage of position to close (0-100)")

	positionCloseAllCmd.Flags().Bool("cancel-orders", false, "Also cancel all open orders")

	positionCmd.AddCommand(positionListCmd)
	positionCmd.AddCommand(positionGetCmd)
	positionCmd.AddCommand(positionCloseCmd)
	positionCmd.AddCommand(positionCloseAllCmd)
}
