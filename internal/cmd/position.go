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
	Short: api.GetAllOpenPositionsOp.Summary,
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
	Short: api.GetOpenPositionOp.Summary,
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
	Short: api.DeleteOpenPositionOp.Summary,
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
	Short: api.DeleteAllOpenPositionsOp.Summary,
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
	cmdutil.RegisterFlags(positionCloseCmd, api.DeleteOpenPositionFlags, &cmdutil.FlagOpts{
		Exclude: map[string]bool{"symbol_or_asset_id": true},
		Aliases: map[string]string{"percentage": "pct"},
	})
	cmdutil.RegisterFlags(positionCloseAllCmd, api.DeleteAllOpenPositionsFlags, nil)

	positionCmd.AddCommand(positionListCmd)
	positionCmd.AddCommand(positionGetCmd)
	positionCmd.AddCommand(positionCloseCmd)
	positionCmd.AddCommand(positionCloseAllCmd)
}
