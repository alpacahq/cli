package cmd

import (
	"github.com/alpacahq/cli/internal/api"
	"github.com/alpacahq/cli/internal/cmdutil"
	"github.com/alpacahq/cli/internal/output"
	"github.com/spf13/cobra"
)

var corporateActionCmd = &cobra.Command{
	Use:     "corporate-action",
	Aliases: []string{"ca"},
	Short:   "Corporate actions announcements",
}

var corporateActionListCmd = &cobra.Command{
	Use:   "list",
	Short: api.OperationSummary["get-v2-corporate_actions-announcements"],
	Example: `  alpaca corporate-action list --types reverse_split --start 2025-01-01 --end 2025-12-31
  alpaca corporate-action list --types cash_dividend --symbols AAPL --start 2025-01-01 --end 2025-06-30`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := cmdutil.RequireAll(cmd, "types", "start", "end"); err != nil {
			return err
		}

		params := &api.GetV2CorporateActionsAnnouncementsParams{
			CaTypes:  cmdutil.Str(cmd, "types"),
			Since:    cmdutil.Str(cmd, "start"),
			Until:    cmdutil.Str(cmd, "end"),
			Symbol:   cmdutil.Str(cmd, "symbols"),
			DateType: cmdutil.Str(cmd, "date-type"),
		}

		data, err := tradingClient.GetV2CorporateActionsAnnouncements(params)
		if err != nil {
			return err
		}

		return output.Render(cmd.OutOrStdout(), getOutput(), corporateActionColumns(), data)
	},
}

var corporateActionGetCmd = &cobra.Command{
	Use:   "get <id>",
	Short: api.OperationSummary["get-v2-corporate_actions-announcements-id"],
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		data, err := tradingClient.GetV2CorporateActionsAnnouncementsID(args[0])
		if err != nil {
			return err
		}
		return output.PrintSingle(cmd.OutOrStdout(), getOutput(), corporateActionColumns(), data)
	},
}

func init() {
	corporateActionListCmd.Flags().String("types", "", "CA types (comma-separated): reverse_split, forward_split, cash_dividend, stock_dividend, spin_off, cash_merger, stock_merger")
	corporateActionListCmd.Flags().String("start", "", "Start date (YYYY-MM-DD, required)")
	corporateActionListCmd.Flags().String("end", "", "End date (YYYY-MM-DD, required)")
	corporateActionListCmd.Flags().String("symbols", "", "Filter by symbols (comma-separated)")
	corporateActionListCmd.Flags().String("date-type", "", "Date type: TRADING or SETTLEMENT")
	_ = corporateActionListCmd.RegisterFlagCompletionFunc("date-type", cobra.FixedCompletions(api.LegacyCalendarParamsDateTypeValues, cobra.ShellCompDirectiveNoFileComp))

	corporateActionCmd.AddCommand(corporateActionListCmd)
	corporateActionCmd.AddCommand(corporateActionGetCmd)
}
