package cmd

import (
	"strings"

	"github.com/alpacahq/cli/internal/api"
	"github.com/alpacahq/cli/internal/cmdutil"
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

var positionGetCmd = fetchCmd("get", api.GetOpenPositionOp, func(cmd *cobra.Command, args []string) (any, error) {
	return tradingClient.GetOpenPosition(normalizeSymbol(cmdutil.Str(cmd, "symbol-or-asset-id")))
}, func(c *cobra.Command) {
	c.Example = `  alpaca position get --symbol-or-asset-id AAPL
  alpaca position get --symbol-or-asset-id BTC/USD`
})

var positionCloseCmd = fetchCmd("close", api.DeleteOpenPositionOp, func(cmd *cobra.Command, args []string) (any, error) {
	return tradingClient.DeleteOpenPosition(normalizeSymbol(cmdutil.Str(cmd, "symbol-or-asset-id")), deleteOpenPositionParamsFromFlags(cmd))
}, func(c *cobra.Command) {
	c.Example = `  alpaca position close --symbol-or-asset-id AAPL
  alpaca position close --symbol-or-asset-id AAPL --qty 5
  alpaca position close --symbol-or-asset-id AAPL --percentage 50`
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
