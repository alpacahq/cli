package cmd

import (
	"strings"

	"github.com/alpacahq/cli/internal/api"
	"github.com/spf13/cobra"
)

// normalizeSymbol converts order-style crypto symbols (BTC/USD) to
// position-style symbols (BTCUSD) since the positions API doesn't
// accept slashes in the path.
func normalizeSymbol(s string) string {
	return strings.ReplaceAll(s, "/", "")
}

var positionCmd = &cobra.Command{
	Use:   "position",
	Short: "Manage positions",
}

var positionListCmd = fetchCmd("list", api.GetAllOpenPositionsOp, func(cmd *cobra.Command, args []string) (any, error) {
	return tradingClient.GetAllOpenPositions()
}, func(c *cobra.Command) {
	c.Example = `  alpaca position list
  alpaca position list --csv`
})

var positionGetCmd = fetchCmd("get <symbol>", api.GetOpenPositionOp, func(cmd *cobra.Command, args []string) (any, error) {
	return tradingClient.GetOpenPosition(normalizeSymbol(args[0]))
}, func(c *cobra.Command) {
	c.Args = cobra.ExactArgs(1)
	c.Example = `  alpaca position get AAPL
  alpaca position get BTC/USD`
})

var positionCloseCmd = fetchCmd("close <symbol>", api.DeleteOpenPositionOp, func(cmd *cobra.Command, args []string) (any, error) {
	return tradingClient.DeleteOpenPosition(normalizeSymbol(args[0]), deleteOpenPositionParamsFromFlags(cmd))
}, func(c *cobra.Command) {
	c.Args = cobra.ExactArgs(1)
	c.Example = `  alpaca position close AAPL
  alpaca position close AAPL --qty 5
  alpaca position close AAPL --percentage 50`
})

var positionCloseAllCmd = fetchCmd("close-all", api.DeleteAllOpenPositionsOp, func(cmd *cobra.Command, args []string) (any, error) {
	return tradingClient.DeleteAllOpenPositions(deleteAllOpenPositionsParamsFromFlags(cmd))
})

func init() {
	positionCmd.AddCommand(positionListCmd)
	positionCmd.AddCommand(positionGetCmd)
	positionCmd.AddCommand(positionCloseCmd)
	positionCmd.AddCommand(positionCloseAllCmd)
}
