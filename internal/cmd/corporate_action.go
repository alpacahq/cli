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
	Short: api.GetV2CorporateActionsAnnouncementsOp.Summary,
	Example: `  alpaca corporate-action list --ca-types reverse_split --since 2025-01-01 --until 2025-12-31
  alpaca corporate-action list --ca-types cash_dividend --symbol AAPL --since 2025-01-01 --until 2025-06-30`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := cmdutil.RequireAll(cmd, "ca-types", "since", "until"); err != nil {
			return err
		}

		params := &api.GetV2CorporateActionsAnnouncementsParams{
			CaTypes:  cmdutil.Str(cmd, "ca-types"),
			Since:    cmdutil.Str(cmd, "since"),
			Until:    cmdutil.Str(cmd, "until"),
			Symbol:   cmdutil.Str(cmd, "symbol"),
			Cusip:    cmdutil.Str(cmd, "cusip"),
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
	Short: api.GetV2CorporateActionsAnnouncementsIDOp.Summary,
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
	cmdutil.RegisterFlags(corporateActionListCmd, api.GetV2CorporateActionsAnnouncementsFlags, nil)

	corporateActionCmd.AddCommand(corporateActionListCmd)
	corporateActionCmd.AddCommand(corporateActionGetCmd)
}
