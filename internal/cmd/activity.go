package cmd

import (
	"encoding/json"
	"net/url"

	"github.com/alpacahq/cli/internal/output"
	"github.com/spf13/cobra"
)

var activityCmd = &cobra.Command{
	Use:   "activity",
	Short: "Account activities (fills, dividends, transfers, etc.)",
}

var activityListCmd = &cobra.Command{
	Use:   "list",
	Short: "List account activities",
	Example: `  alpaca activity list
  alpaca activity list --type FILL --limit 20
  alpaca activity list --type DIV --after 2025-01-01
  alpaca activity list --type FILL,TRANS --direction desc`,
	RunE: func(cmd *cobra.Command, args []string) error {
		params := url.Values{}

		actType, _ := cmd.Flags().GetString("type")
		after, _ := cmd.Flags().GetString("after")
		until, _ := cmd.Flags().GetString("until")
		date, _ := cmd.Flags().GetString("date")
		direction, _ := cmd.Flags().GetString("direction")
		limit, _ := cmd.Flags().GetString("limit")
		category, _ := cmd.Flags().GetString("category")
		pageToken, _ := cmd.Flags().GetString("page-token")

		if after != "" {
			params.Set("after", after)
		}
		if until != "" {
			params.Set("until", until)
		}
		if date != "" {
			params.Set("date", date)
		}
		if direction != "" {
			params.Set("direction", direction)
		}
		if limit != "" {
			params.Set("page_size", limit)
		}
		if category != "" {
			params.Set("category", category)
		}
		if pageToken != "" {
			params.Set("page_token", pageToken)
		}

		var path string
		if actType != "" {
			path = "/v2/account/activities/" + actType
		} else {
			path = "/v2/account/activities"
		}

		data, err := apiClient.Get(path, params)
		if err != nil {
			return err
		}

		out := getOutput()

		var items []map[string]any
		if err := json.Unmarshal(data, &items); err != nil {
			return output.JSON(cmd.OutOrStdout(), data)
		}

		if len(items) == 0 {
			return output.Render(out, tradeActivityColumns(), data)
		}

		_, isTrade := items[0]["cum_qty"]
		if isTrade {
			return output.Render(out, tradeActivityColumns(), data)
		}
		return output.Render(out, nonTradeActivityColumns(), data)
	},
}

func tradeActivityColumns() []output.Column {
	return []output.Column{
		{Header: "ID", Field: "id"},
		{Header: "TYPE", Field: "activity_type"},
		{Header: "SYMBOL", Field: "symbol"},
		{Header: "SIDE", Field: "side"},
		{Header: "QTY", Field: "qty"},
		{Header: "PRICE", Field: "price"},
		{Header: "CUM QTY", Field: "cum_qty"},
		{Header: "ORDER ID", Field: "order_id"},
		{Header: "TIME", Field: "transaction_time"},
	}
}

func nonTradeActivityColumns() []output.Column {
	return []output.Column{
		{Header: "ID", Field: "id"},
		{Header: "TYPE", Field: "activity_type"},
		{Header: "SYMBOL", Field: "symbol"},
		{Header: "NET AMOUNT", Field: "net_amount"},
		{Header: "QTY", Field: "qty"},
		{Header: "PER SHARE", Field: "per_share_amount"},
		{Header: "STATUS", Field: "status"},
		{Header: "DATE", Field: "date"},
	}
}

func init() {
	activityListCmd.Flags().String("type", "", "Activity type: FILL, DIV, TRANS, etc. (comma-separated for multiple)")
	activityListCmd.Flags().String("after", "", "Only activities after this date/time")
	activityListCmd.Flags().String("until", "", "Only activities before this date/time")
	activityListCmd.Flags().String("date", "", "Exact date filter")
	activityListCmd.Flags().String("direction", "", "Sort: asc or desc")
	activityListCmd.Flags().String("limit", "", "Max number of results")
	activityListCmd.Flags().String("category", "", "Category: trade_activity or non_trade_activity")
	activityListCmd.Flags().String("page-token", "", "Pagination token")

	activityCmd.AddCommand(activityListCmd)
}
