package cmd

import (
	"github.com/alpacahq/cli/internal/api"
	"github.com/alpacahq/cli/internal/cmdutil"
	"github.com/spf13/cobra"
)

var clockCmd = fetchCmd("clock", api.LegacyClockOp, func(cmd *cobra.Command, args []string) (any, error) {
	return tradingClient.LegacyClock()
}, func(c *cobra.Command) {
	c.Example = `  alpaca clock`
})

var clockMarketsCmd = fetchCmd("markets", api.ClockOp, func(cmd *cobra.Command, args []string) (any, error) {
	return tradingClient.Clock(clockParamsFromFlags(cmd))
}, jsonOnly, func(c *cobra.Command) {
	c.Example = `  alpaca clock markets --markets XNYS,XNAS`
})

var calendarCmd = fetchCmd("calendar", api.LegacyCalendarOp, func(cmd *cobra.Command, args []string) (any, error) {
	return tradingClient.LegacyCalendar(legacyCalendarParamsFromFlags(cmd))
}, func(c *cobra.Command) {
	c.Example = `  alpaca calendar
  alpaca calendar --start 2025-01-01 --end 2025-12-31`
})

var calendarMarketCmd = fetchCmd("market", api.CalendarOp, func(cmd *cobra.Command, args []string) (any, error) {
	return tradingClient.Calendar(cmdutil.Str(cmd, "market"), calendarParamsFromFlags(cmd))
}, func(c *cobra.Command) {
	c.Example = `  alpaca calendar market --market XNYS --start 2025-01-01`
})

func init() {
	clockCmd.AddCommand(clockMarketsCmd)
	calendarCmd.AddCommand(calendarMarketCmd)
}
