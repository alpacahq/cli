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
	Short: api.GetAccountActivitiesOp.Summary,
	Example: `  alpaca activity list
  alpaca activity list --types FILL --limit 20
  alpaca activity list --types DIV --start 2025-01-01
  alpaca activity list --types FILL,TRANS --sort desc`,
	RunE: func(cmd *cobra.Command, args []string) error {
		actType := cmdutil.Str(cmd, "types")

		var data json.RawMessage
		var err error

		if actType != "" {
			data, err = tradingClient.GetAccountActivitiesByActivityType(actType, &api.GetAccountActivitiesByActivityTypeParams{
				After:     cmdutil.Str(cmd, "start"),
				Until:     cmdutil.Str(cmd, "end"),
				Date:      cmdutil.Str(cmd, "date"),
				Direction: cmdutil.Str(cmd, "sort"),
				PageSize:  cmdutil.Int(cmd, "limit"),
				PageToken: cmdutil.Str(cmd, "page-token"),
			})
		} else {
			data, err = tradingClient.GetAccountActivities(&api.GetAccountActivitiesParams{
				After:     cmdutil.Str(cmd, "start"),
				Until:     cmdutil.Str(cmd, "end"),
				Date:      cmdutil.Str(cmd, "date"),
				Direction: cmdutil.Str(cmd, "sort"),
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
			return output.Render(cmd.OutOrStdout(), out, tradeActivityColumns(), data)
		}

		_, isTrade := items[0]["cum_qty"]
		if isTrade {
			return output.Render(cmd.OutOrStdout(), out, tradeActivityColumns(), data)
		}
		return output.Render(cmd.OutOrStdout(), out, nonTradeActivityColumns(), data)
	},
}

func init() {
	cmdutil.RegisterFlags(activityListCmd, api.GetAccountActivitiesFlags, &cmdutil.FlagOpts{
		Aliases: map[string]string{
			"activity_types": "types",
			"after":          "start",
			"until":          "end",
			"direction":      "sort",
			"page_size":      "limit",
		},
	})

	activityCmd.AddCommand(activityListCmd)
}
