package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/alpacahq/cli/internal/api"
	"github.com/alpacahq/cli/internal/cmdutil"
	"github.com/alpacahq/cli/internal/output"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var clockCmd = &cobra.Command{
	Use:   "clock",
	Short: api.LegacyClockOp.Summary,
	Example: `  alpaca clock
  alpaca clock --markets XNYS,XNAS`,
	RunE: func(cmd *cobra.Command, args []string) error {
		params := clockParamsFromFlags(cmd)

		if params.Markets != "" {
			resp, err := tradingClient.Clock(params)
			if err != nil {
				return err
			}
			return output.JSON(cmd.OutOrStdout(), resp)
		}

		clock, err := tradingClient.LegacyClock()
		if err != nil {
			return err
		}

		if getOutput() == outputJSON {
			return output.JSON(cmd.OutOrStdout(), clock)
		}

		w := cmd.OutOrStdout()
		if clock.IsOpen {
			_, _ = color.New(color.FgGreen).Fprintln(w, "Market is OPEN")
		} else {
			_, _ = color.New(color.FgYellow).Fprintln(w, "Market is CLOSED")
		}
		_, _ = fmt.Fprintf(w, "  Next open:  %v\n", clock.NextOpen)
		_, _ = fmt.Fprintf(w, "  Next close: %v\n", clock.NextClose)
		return nil
	},
}

var calendarCmd = &cobra.Command{
	Use:   "calendar",
	Short: api.LegacyCalendarOp.Summary,
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
			return output.JSON(cmd.OutOrStdout(), resp)
		}

		params := legacyCalendarParamsFromFlags(cmd)

		data, err := tradingClient.LegacyCalendar(params)
		if err != nil {
			return err
		}
		return output.Render(cmd.OutOrStdout(), getOutput(), calendarColumns(), data)
	},
}

var portfolioCmd = &cobra.Command{
	Use:   "portfolio",
	Short: "Portfolio equity and P&L history",
}

var portfolioHistoryCmd = &cobra.Command{
	Use:   "history",
	Short: api.GetAccountPortfolioHistoryOp.Summary,
	Long:  "Returns portfolio equity and P&L history. Output is always JSON due to complex time-series structure.",
	Example: `  alpaca portfolio history
  alpaca portfolio history --period 1M --timeframe 1D`,
	RunE: func(cmd *cobra.Command, args []string) error {
		params := getAccountPortfolioHistoryParamsFromFlags(cmd)

		history, err := tradingClient.GetAccountPortfolioHistory(params)
		if err != nil {
			return err
		}
		return output.JSON(cmd.OutOrStdout(), history)
	},
}

var newsCmd = &cobra.Command{
	Use:   "news",
	Short: api.NewsOp.Summary,
	Example: `  alpaca news
  alpaca news --symbols AAPL,MSFT --limit 10`,
	RunE: func(cmd *cobra.Command, args []string) error {
		params := newsParamsFromFlags(cmd)
		if params.Limit == 0 {
			params.Limit = 10
		}

		resp, err := dataClient.News(params)
		if err != nil {
			return err
		}

		newsData, _ := json.Marshal(resp.News)
		return output.Render(cmd.OutOrStdout(), getOutput(), newsColumns(), json.RawMessage(newsData))
	},
}

func init() {
	cmdutil.RegisterFlags(clockCmd, api.ClockFlags, nil)

	cmdutil.RegisterFlags(calendarCmd, api.LegacyCalendarFlags, nil)
	calendarCmd.Flags().String("market", "", "Market MIC for v3 calendar (e.g. XNYS)")

	cmdutil.RegisterFlags(portfolioHistoryCmd, api.GetAccountPortfolioHistoryFlags, &cmdutil.FlagOpts{
		Exclude: map[string]bool{"extended_hours": true},
	})
	portfolioCmd.AddCommand(portfolioHistoryCmd)

	cmdutil.RegisterFlags(newsCmd, api.NewsFlags, nil)
}
