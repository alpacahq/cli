package cmd

import (
	"github.com/alpacahq/cli/internal/api"
	"github.com/spf13/cobra"
)

var positionCmd = &cobra.Command{
	Use:   "position",
	Short: "Manage positions",
}

var positionListCmd = fetchCmd("list", api.GetAllOpenPositionsOp, func(cmd *cobra.Command, args []string) (any, error) {
	return tradingClient.GetAllOpenPositions()
}, func(c *cobra.Command) {
	c.Example = `  alpaca position list
  alpaca position list --json
  alpaca position list --csv`
	cmdColumns[c] = positionColumns()
})

var positionGetCmd = fetchCmd("get <symbol>", api.GetOpenPositionOp, func(cmd *cobra.Command, args []string) (any, error) {
	return tradingClient.GetOpenPosition(args[0])
}, func(c *cobra.Command) {
	c.Args = cobra.ExactArgs(1)
	c.Example = `  alpaca position get AAPL
  alpaca position get BTC/USD --json`
	cmdColumns[c] = positionColumns()
})

var positionCloseCmd = fetchCmd("close <symbol>", api.DeleteOpenPositionOp, func(cmd *cobra.Command, args []string) (any, error) {
	return tradingClient.DeleteOpenPosition(args[0], deleteOpenPositionParamsFromFlags(cmd))
}, func(c *cobra.Command) {
	c.Args = cobra.ExactArgs(1)
	c.Example = `  alpaca position close AAPL
  alpaca position close AAPL --qty 5
  alpaca position close AAPL --percentage 50`
	cmdColumns[c] = orderColumns()
})

var positionCloseAllCmd = actionCmd("close-all", api.DeleteAllOpenPositionsOp, "All positions closed.", func(cmd *cobra.Command, args []string) error {
	_, err := tradingClient.DeleteAllOpenPositions(deleteAllOpenPositionsParamsFromFlags(cmd))
	return err
})

func init() {
	positionCmd.AddCommand(positionListCmd)
	positionCmd.AddCommand(positionGetCmd)
	positionCmd.AddCommand(positionCloseCmd)
	positionCmd.AddCommand(positionCloseAllCmd)
}
