package cmd

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/alpacahq/cli/internal/output"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var clockCmd = &cobra.Command{
	Use:   "clock",
	Short: "Show market clock",
	RunE: func(cmd *cobra.Command, args []string) error {
		data, err := apiClient.Get("/v2/clock", nil)
		if err != nil {
			return err
		}

		if getOutput() == "json" {
			return output.JSON(cmd.OutOrStdout(), data)
		}

		var clock map[string]any
		json.Unmarshal(data, &clock)

		isOpen, _ := clock["is_open"].(bool)
		if isOpen {
			color.Green("Market is OPEN")
		} else {
			color.Yellow("Market is CLOSED")
		}
		fmt.Printf("  Next open:  %v\n", clock["next_open"])
		fmt.Printf("  Next close: %v\n", clock["next_close"])

		return nil
	},
}

var calendarCmd = &cobra.Command{
	Use:   "calendar",
	Short: "Show trading calendar",
	Example: `  alpaca calendar
  alpaca calendar --start 2025-01-01 --end 2025-12-31`,
	RunE: func(cmd *cobra.Command, args []string) error {
		params := url.Values{}
		start, _ := cmd.Flags().GetString("start")
		if start != "" {
			params.Set("start", start)
		}
		end, _ := cmd.Flags().GetString("end")
		if end != "" {
			params.Set("end", end)
		}

		data, err := apiClient.Get("/v2/calendar", params)
		if err != nil {
			return err
		}

		columns := []output.Column{
			{Header: "DATE", Field: "date"},
			{Header: "OPEN", Field: "open"},
			{Header: "CLOSE", Field: "close"},
			{Header: "SESSION OPEN", Field: "session_open"},
			{Header: "SESSION CLOSE", Field: "session_close"},
		}

		return output.Render(getOutput(), columns, data)
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
		params := url.Values{}
		period, _ := cmd.Flags().GetString("period")
		if period != "" {
			params.Set("period", period)
		}
		tf, _ := cmd.Flags().GetString("timeframe")
		if tf != "" {
			params.Set("timeframe", tf)
		}

		data, err := apiClient.Get("/v2/account/portfolio/history", params)
		if err != nil {
			return err
		}

		return output.JSON(cmd.OutOrStdout(), data)
	},
}

var newsCmd = &cobra.Command{
	Use:   "news",
	Short: "Get market news",
	Example: `  alpaca news
  alpaca news --symbols AAPL,MSFT --limit 10`,
	RunE: func(cmd *cobra.Command, args []string) error {
		params := url.Values{}
		symbols, _ := cmd.Flags().GetString("symbols")
		if symbols != "" {
			params.Set("symbols", symbols)
		}
		start, _ := cmd.Flags().GetString("start")
		if start != "" {
			params.Set("start", start)
		}
		end, _ := cmd.Flags().GetString("end")
		if end != "" {
			params.Set("end", end)
		}
		limit, _ := cmd.Flags().GetString("limit")
		if limit != "" {
			params.Set("limit", limit)
		} else {
			params.Set("limit", "10")
		}

		data, err := apiClient.GetData("/v1beta1/news", params)
		if err != nil {
			return err
		}

		// News response: {"news": [...]}
		var resp map[string]json.RawMessage
		if json.Unmarshal(data, &resp) == nil {
			if news, ok := resp["news"]; ok {
				data = news
			}
		}

		columns := []output.Column{
			{Header: "DATE", Field: "created_at"},
			{Header: "HEADLINE", Field: "headline"},
			{Header: "SOURCE", Field: "source"},
			{Header: "SYMBOLS", Field: "symbols"},
			{Header: "URL", Field: "url"},
		}

		return output.Render(getOutput(), columns, data)
	},
}

func init() {
	calendarCmd.Flags().String("start", "", "Start date (YYYY-MM-DD)")
	calendarCmd.Flags().String("end", "", "End date (YYYY-MM-DD)")

	portfolioHistoryCmd.Flags().String("period", "", "Period: 1D, 1W, 1M, 3M, 1A, all")
	portfolioHistoryCmd.Flags().String("timeframe", "", "Timeframe: 1Min, 5Min, 15Min, 1H, 1D")
	portfolioCmd.AddCommand(portfolioHistoryCmd)

	newsCmd.Flags().String("symbols", "", "Filter by symbols (comma-separated)")
	newsCmd.Flags().String("start", "", "Start date")
	newsCmd.Flags().String("end", "", "End date")
	newsCmd.Flags().String("limit", "", "Max articles (default: 10)")
}
