package cmd

import (
	"github.com/alpacahq/cli/internal/api"
	"github.com/alpacahq/cli/internal/cmdutil"
	"github.com/alpacahq/cli/internal/output"
	"github.com/spf13/cobra"
)

var clockCmd = &cobra.Command{
	Use:   "clock",
	Short: api.LegacyClockOp.Summary(),
	Example: `  alpaca clock
  alpaca clock --markets XNYS,XNAS`,
	RunE: func(cmd *cobra.Command, args []string) error {
		params := clockParamsFromFlags(cmd)

		if params.Markets != "" {
			resp, err := tradingClient.Clock(params)
			if err != nil {
				return err
			}
			return output.Render(cmd.OutOrStdout(), getOutput(), resp)
		}

		clock, err := tradingClient.LegacyClock()
		if err != nil {
			return err
		}
		return output.Render(cmd.OutOrStdout(), getOutput(), clock)
	},
}

var calendarCmd = &cobra.Command{
	Use:   "calendar",
	Short: api.LegacyCalendarOp.Summary(),
	Example: `  alpaca calendar
  alpaca calendar --start 2025-01-01 --end 2025-12-31
  alpaca calendar --market XNYS --start 2025-01-01`,
	RunE: func(cmd *cobra.Command, args []string) error {
		market := cmdutil.Str(cmd, "market")

		if market != "" {
			resp, err := tradingClient.Calendar(market, calendarParamsFromFlags(cmd))
			if err != nil {
				return err
			}
			return output.Render(cmd.OutOrStdout(), getOutput(), resp)
		}

		data, err := tradingClient.LegacyCalendar(legacyCalendarParamsFromFlags(cmd))
		if err != nil {
			return err
		}
		return output.Render(cmd.OutOrStdout(), getOutput(), data)
	},
}

func init() {
	cmdutil.RegisterFlags(clockCmd, api.ClockOp.Flags(), nil)

	cmdutil.RegisterFlags(calendarCmd, api.LegacyCalendarOp.Flags(), nil)
	calendarCmd.Flags().String("market", "", "Market MIC for v3 calendar (e.g. XNYS)")
}
