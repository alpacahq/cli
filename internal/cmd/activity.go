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
	Short: api.GetAccountActivitiesOp.Summary(),
	Example: `  alpaca account activity list
  alpaca account activity list --activity-types FILL --page-size 20
  alpaca account activity list --activity-types DIV --after 2025-01-01
  alpaca account activity list --activity-types FILL,TRANS --direction desc`,
	RunE: func(cmd *cobra.Command, args []string) error {
		actType := cmdutil.Str(cmd, "activity-types")

		var data json.RawMessage
		var err error

		if actType != "" {
			data, err = tradingClient.GetAccountActivitiesByActivityType(actType, getAccountActivitiesByActivityTypeParamsFromFlags(cmd))
		} else {
			data, err = tradingClient.GetAccountActivities(getAccountActivitiesParamsFromFlags(cmd))
		}
		if err != nil {
			return err
		}

		out := getOutput()

		var items []map[string]any
		if err := json.Unmarshal(data, &items); err != nil {
			return output.JSON(cmd.OutOrStdout(), data)
		}

		return output.Render(cmd.OutOrStdout(), out, nil, data)
	},
}

func init() {
	cmdutil.RegisterFlags(activityListCmd, api.GetAccountActivitiesOp.Flags(), nil)

	activityCmd.AddCommand(activityListCmd)
}
