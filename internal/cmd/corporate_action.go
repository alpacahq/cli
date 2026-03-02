package cmd

import (
	"fmt"

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
	Short: "List corporate action announcements",
	Example: `  alpaca corporate-action list --types reverse_split --start 2025-01-01 --end 2025-12-31
  alpaca corporate-action list --types cash_dividend --symbols AAPL --start 2025-01-01 --end 2025-06-30`,
	RunE: func(cmd *cobra.Command, args []string) error {
		types := cmdutil.Str(cmd, "types")
		if types == "" {
			return fmt.Errorf("--types is required (e.g. reverse_split, forward_split, cash_dividend, stock_dividend, spin_off, cash_merger, stock_merger)")
		}
		start := cmdutil.Str(cmd, "start")
		if start == "" {
			return fmt.Errorf("--start is required (YYYY-MM-DD)")
		}
		end := cmdutil.Str(cmd, "end")
		if end == "" {
			return fmt.Errorf("--end is required (YYYY-MM-DD)")
		}

		params := &api.GetV2CorporateActionsAnnouncementsParams{
			CaTypes:  types,
			Since:    start,
			Until:    end,
			Symbol:   cmdutil.Str(cmd, "symbols"),
			DateType: cmdutil.Str(cmd, "date-type"),
		}

		data, err := tradingClient.GetV2CorporateActionsAnnouncements(params)
		if err != nil {
			return err
		}

		return output.Render(getOutput(), corporateActionColumns(), data)
	},
}

var corporateActionGetCmd = &cobra.Command{
	Use:   "get <id>",
	Short: "Get a specific corporate action announcement",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		data, err := tradingClient.GetV2CorporateActionsAnnouncementsID(args[0])
		if err != nil {
			return err
		}
		return output.PrintSingle(getOutput(), corporateActionColumns(), data)
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
