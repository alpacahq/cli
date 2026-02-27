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
	Short: "Show market clock",
	RunE: func(cmd *cobra.Command, args []string) error {
		clock, err := tradingClient.LegacyClock()
		if err != nil {
			return err
		}

		if getOutput() == "json" {
			return output.JSON(cmd.OutOrStdout(), clock)
		}

		if clock.IsOpen {
			color.Green("Market is OPEN")
		} else {
			color.Yellow("Market is CLOSED")
		}
		fmt.Printf("  Next open:  %v\n", clock.NextOpen)
		fmt.Printf("  Next close: %v\n", clock.NextClose)
		return nil
	},
}

var calendarCmd = &cobra.Command{
	Use:   "calendar",
	Short: "Show trading calendar",
	Example: `  alpaca calendar
  alpaca calendar --start 2025-01-01 --end 2025-12-31`,
	RunE: func(cmd *cobra.Command, args []string) error {
		params := &api.LegacyCalendarParams{
			Start: cmdutil.Str(cmd, "start"),
			End:   cmdutil.Str(cmd, "end"),
		}

		data, err := tradingClient.LegacyCalendar(params)
		if err != nil {
			return err
		}
		return output.Render(getOutput(), calendarColumns(), data)
	},
}

var portfolioCmd = &cobra.Command{
	Use:   "portfolio",
	Short: "Portfolio analytics",
}

var portfolioHistoryCmd = &cobra.Command{
	Use:   "history",
	Short: "Get portfolio value history",
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
	Short: "Get market news",
	Example: `  alpaca news
  alpaca news --symbols AAPL,MSFT --limit 10`,
	RunE: func(cmd *cobra.Command, args []string) error {
		params := &api.NewsParams{
			Symbols:        cmdutil.Str(cmd, "symbols"),
			Start:          cmdutil.Str(cmd, "start"),
			End:            cmdutil.Str(cmd, "end"),
			Sort:           cmdutil.Str(cmd, "sort"),
			IncludeContent: cmdutil.Bool(cmd, "include-content"),
			Limit:          cmdutil.Int(cmd, "limit"),
		}
		if params.Limit == 0 {
			params.Limit = 10
		}

		resp, err := dataClient.News(params)
		if err != nil {
			return err
		}

		// Unwrap the news array from the response
		newsData, _ := json.Marshal(resp.News)
		return output.Render(getOutput(), newsColumns(), json.RawMessage(newsData))
	},
}

func init() {
	calendarCmd.Flags().String("start", "", "Start date (YYYY-MM-DD)")
	calendarCmd.Flags().String("end", "", "End date (YYYY-MM-DD)")

	portfolioHistoryCmd.Flags().String("period", "", "Period: 1D, 1W, 1M, 3M, 1A, all")
	portfolioHistoryCmd.Flags().String("timeframe", "", "Timeframe: 1Min, 5Min, 15Min, 1H, 1D")
	portfolioHistoryCmd.Flags().String("start", "", "Start date (RFC3339)")
	portfolioHistoryCmd.Flags().String("end", "", "End date (RFC3339)")
	portfolioHistoryCmd.Flags().String("intraday-reporting", "", "Intraday reporting: market_hours, extended_hours, continuous")
	portfolioHistoryCmd.Flags().String("pnl-reset", "", "P&L reset mode: no_reset, per_day")
	portfolioCmd.AddCommand(portfolioHistoryCmd)

	newsCmd.Flags().String("symbols", "", "Filter by symbols (comma-separated)")
	newsCmd.Flags().String("start", "", "Start date")
	newsCmd.Flags().String("end", "", "End date")
	newsCmd.Flags().Int("limit", 0, "Max articles (default: 10)")
	newsCmd.Flags().String("sort", "", "Sort order: asc or desc")
	newsCmd.Flags().Bool("include-content", false, "Include full article content")
}
