package cmd

import (
	"fmt"
	"net/url"

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
  alpaca position close AAPL --pct 50`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		v := url.Values{}
		if q := cmdutil.Str(cmd, "qty"); q != "" {
			v.Set("qty", q)
		}
		if p := cmdutil.Str(cmd, "pct"); p != "" {
			v.Set("percentage", p)
		}

		path := fmt.Sprintf("/v2/positions/%s", args[0])
		data, err := apiClient.Delete(path, v)
		if err != nil {
			return err
		}
		return output.PrintSingle(cmd.OutOrStdout(), getOutput(), orderColumns(), data)
	},
}

var positionCloseAllCmd = &cobra.Command{
	Use:   "close-all",
	Short: api.DeleteAllOpenPositionsOp.Summary,
	RunE: func(cmd *cobra.Command, args []string) error {
		canceled, err := tradingClient.DeleteAllOpenPositions(&api.DeleteAllOpenPositionsParams{
			CancelOrders: cmdutil.Bool(cmd, "cancel-orders"),
		})
		if err != nil {
			return err
		}
		if getOutput() == outputJSON || getOutput() == outputCSV {
			return output.JSON(cmd.OutOrStdout(), canceled)
		}
		_, _ = fmt.Fprintln(cmd.OutOrStdout(), "All positions closed.")
		return nil
	},
}

func init() {
	cmdutil.RegisterFlags(positionCloseCmd, api.DeleteOpenPositionFlags, &cmdutil.FlagOpts{
		Exclude: map[string]bool{"symbol_or_asset_id": true, "qty": true, "percentage": true},
	})
	positionCloseCmd.Flags().String("qty", "", "Number of shares to liquidate (up to 9 decimal points). Cannot use with --pct")
	positionCloseCmd.Flags().String("pct", "", "Percentage of position to liquidate. Cannot use with --qty")
	cmdutil.RegisterFlags(positionCloseAllCmd, api.DeleteAllOpenPositionsFlags, nil)

	positionCmd.AddCommand(positionListCmd)
	positionCmd.AddCommand(positionGetCmd)
	positionCmd.AddCommand(positionCloseCmd)
	positionCmd.AddCommand(positionCloseAllCmd)
}
