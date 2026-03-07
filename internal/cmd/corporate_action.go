package cmd

import (
	"github.com/alpacahq/cli/internal/api"
	"github.com/spf13/cobra"
)

var corporateActionCmd = &cobra.Command{
	Use:     "corporate-action",
	Aliases: []string{"ca"},
	Short:   "Corporate actions announcements",
}

var corporateActionListCmd = fetchCmd("list",
	api.GetV2CorporateActionsAnnouncementsOp,
	func(cmd *cobra.Command, _ []string) (any, error) {
		return tradingClient.GetV2CorporateActionsAnnouncements(
			getV2CorporateActionsAnnouncementsParamsFromFlags(cmd))
	},
	func(c *cobra.Command) {
		c.Example = `  alpaca corporate-action list --ca-types reverse_split --since 2025-01-01 --until 2025-12-31
  alpaca corporate-action list --ca-types cash_dividend --symbol AAPL --since 2025-01-01 --until 2025-06-30`
	})

var corporateActionGetCmd = fetchCmd("get <id>",
	api.GetV2CorporateActionsAnnouncementsIDOp,
	func(_ *cobra.Command, args []string) (any, error) {
		return tradingClient.GetV2CorporateActionsAnnouncementsID(args[0])
	},
	func(c *cobra.Command) {
		c.Args = cobra.ExactArgs(1)
	})

func init() {
	corporateActionCmd.AddCommand(corporateActionListCmd)
	corporateActionCmd.AddCommand(corporateActionGetCmd)
}
