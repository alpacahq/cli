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
	Example: `  alpaca position list
  alpaca position list --json
  alpaca position list --csv`,
	RunE: func(cmd *cobra.Command, args []string) error {
		positions, err := tradingClient.GetAllOpenPositions()
		if err != nil {
			return err
		}
		return output.RenderWithHint(cmd.OutOrStdout(), getOutput(), positionColumns(), positions, "No open positions.")
	},
}

var positionGetCmd = &cobra.Command{
	Use:   "get <symbol>",
	Short: api.GetOpenPositionOp.Summary,
	Example: `  alpaca position get AAPL
  alpaca position get BTC/USD --json`,
	Args: cobra.ExactArgs(1),
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
  alpaca position close AAPL --percentage 50`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		params := deleteOpenPositionParamsFromFlags(cmd)
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
		canceled, err := tradingClient.DeleteAllOpenPositions(deleteAllOpenPositionsParamsFromFlags(cmd))
		if err != nil {
			return err
		}
		format := getOutput()
		if format == output.FormatJSON || format == output.FormatCSV {
			return output.JSON(cmd.OutOrStdout(), canceled)
		}
		_, _ = fmt.Fprintln(cmd.OutOrStdout(), "All positions closed.")
		return nil
	},
}

func init() {
	cmdutil.RegisterFlags(positionCloseCmd, api.DeleteOpenPositionFlags, nil)
	cmdutil.RegisterFlags(positionCloseAllCmd, api.DeleteAllOpenPositionsFlags, nil)

	positionCmd.AddCommand(positionListCmd)
	positionCmd.AddCommand(positionGetCmd)
	positionCmd.AddCommand(positionCloseCmd)
	positionCmd.AddCommand(positionCloseAllCmd)
}
