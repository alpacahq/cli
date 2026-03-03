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
		markets := cmdutil.Str(cmd, "markets")

		if markets != "" {
			resp, err := tradingClient.Clock(&api.ClockParams{
				Markets: markets,
				Time:    cmdutil.Str(cmd, "time"),
			})
			if err != nil {
				return err
			}
			return output.JSON(cmd.OutOrStdout(), resp)
		}

		clock, err := tradingClient.LegacyClock()
		if err != nil {
			return err
		}

		if getOutput() == "json" {
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
			resp, err := tradingClient.Calendar(market, &api.CalendarParams{
				Start: cmdutil.Str(cmd, "start"),
				End:   cmdutil.Str(cmd, "end"),
			})
			if err != nil {
				return err
			}
			return output.JSON(cmd.OutOrStdout(), resp)
		}

		params := &api.LegacyCalendarParams{
			Start:    cmdutil.Str(cmd, "start"),
			End:      cmdutil.Str(cmd, "end"),
			DateType: cmdutil.Str(cmd, "date-type"),
		}

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
	Example: `  alpaca portfolio history
  alpaca portfolio history --period 1M --timeframe 1D`,
	RunE: func(cmd *cobra.Command, args []string) error {
		params := &api.GetAccountPortfolioHistoryParams{
			Period:            cmdutil.Str(cmd, "period"),
			Timeframe:         cmdutil.Str(cmd, "timeframe"),
			Start:             cmdutil.Str(cmd, "start"),
			End:               cmdutil.Str(cmd, "end"),
			IntradayReporting: cmdutil.Str(cmd, "intraday-reporting"),
			PNLReset:          cmdutil.Str(cmd, "pnl-reset"),
			CashflowTypes:     cmdutil.Str(cmd, "cashflow-types"),
		}

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
		params := &api.NewsParams{
			Symbols:            cmdutil.Str(cmd, "symbols"),
			Start:              cmdutil.Str(cmd, "start"),
			End:                cmdutil.Str(cmd, "end"),
			Sort:               cmdutil.Str(cmd, "sort"),
			IncludeContent:     cmdutil.Bool(cmd, "include-content"),
			ExcludeContentless: cmdutil.Bool(cmd, "exclude-contentless"),
			Limit:              cmdutil.Int(cmd, "limit"),
			PageToken:          cmdutil.Str(cmd, "page-token"),
		}
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
