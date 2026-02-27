package cmd

import (
	"encoding/json"

	"github.com/alpacahq/cli/internal/api"
	"github.com/alpacahq/cli/internal/cmdutil"
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
		actType := cmdutil.Str(cmd, "type")

		var data json.RawMessage
		var err error

		if actType != "" {
			data, err = tradingClient.GetAccountActivitiesByActivityType(actType, &api.GetAccountActivitiesByActivityTypeParams{
				After:     cmdutil.Str(cmd, "after"),
				Until:     cmdutil.Str(cmd, "until"),
				Date:      cmdutil.Str(cmd, "date"),
				Direction: cmdutil.Str(cmd, "direction"),
				PageSize:  cmdutil.Int(cmd, "limit"),
				PageToken: cmdutil.Str(cmd, "page-token"),
			})
		} else {
			data, err = tradingClient.GetAccountActivities(&api.GetAccountActivitiesParams{
				After:     cmdutil.Str(cmd, "after"),
				Until:     cmdutil.Str(cmd, "until"),
				Date:      cmdutil.Str(cmd, "date"),
				Direction: cmdutil.Str(cmd, "direction"),
				PageSize:  cmdutil.Int(cmd, "limit"),
				Category:  cmdutil.Str(cmd, "category"),
				PageToken: cmdutil.Str(cmd, "page-token"),
			})
		}
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

func init() {
	activityListCmd.Flags().String("type", "", "Activity type: FILL, DIV, TRANS, etc. (comma-separated for multiple)")
	activityListCmd.Flags().String("after", "", "Only activities after this date/time")
	activityListCmd.Flags().String("until", "", "Only activities before this date/time")
	activityListCmd.Flags().String("date", "", "Exact date filter")
	activityListCmd.Flags().String("direction", "", "Sort: asc or desc")
	activityListCmd.Flags().Int("limit", 0, "Max number of results")
	activityListCmd.Flags().String("category", "", "Category: trade_activity or non_trade_activity")
	activityListCmd.Flags().String("page-token", "", "Pagination token")

	activityCmd.AddCommand(activityListCmd)
}
