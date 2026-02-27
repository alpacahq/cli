package cmd

import (
	"fmt"
	"net/url"

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
	Example: `  alpaca corporate-action list --types reverse_split --since 2025-01-01 --until 2025-12-31
  alpaca corporate-action list --types cash_dividend --symbol AAPL --since 2025-01-01 --until 2025-06-30`,
	RunE: func(cmd *cobra.Command, args []string) error {
		params := url.Values{}

		types, _ := cmd.Flags().GetString("types")
		if types == "" {
			return fmt.Errorf("--types is required (e.g. reverse_split, forward_split, cash_dividend, stock_dividend, spin_off, cash_merger, stock_merger)")
		}
		params.Set("ca_types", types)

		since, _ := cmd.Flags().GetString("since")
		if since == "" {
			return fmt.Errorf("--since is required (YYYY-MM-DD)")
		}
		params.Set("since", since)

		until, _ := cmd.Flags().GetString("until")
		if until == "" {
			return fmt.Errorf("--until is required (YYYY-MM-DD)")
		}
		params.Set("until", until)

		symbol, _ := cmd.Flags().GetString("symbol")
		if symbol != "" {
			params.Set("symbol", symbol)
		}
		dateType, _ := cmd.Flags().GetString("date-type")
		if dateType != "" {
			params.Set("date_type", dateType)
		}

		data, err := apiClient.Get("/v2/corporate_actions/announcements", params)
		if err != nil {
			return err
		}

		columns := []output.Column{
			{Header: "ID", Field: "id"},
			{Header: "TYPE", Field: "ca_type"},
			{Header: "SUB TYPE", Field: "ca_sub_type"},
			{Header: "SYMBOL", Field: "symbol"},
			{Header: "EX DATE", Field: "ex_date"},
			{Header: "RECORD DATE", Field: "record_date"},
			{Header: "PAYABLE DATE", Field: "payable_date"},
		}

		return output.Render(getOutput(), columns, data)
	},
}

var corporateActionGetCmd = &cobra.Command{
	Use:   "get <id>",
	Short: "Get a specific corporate action announcement",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		data, err := apiClient.Get("/v2/corporate_actions/announcements/"+args[0], nil)
		if err != nil {
			return err
		}
		return output.JSON(cmd.OutOrStdout(), data)
	},
}

func init() {
	corporateActionListCmd.Flags().String("types", "", "CA types (comma-separated): reverse_split, forward_split, cash_dividend, stock_dividend, spin_off, cash_merger, stock_merger")
	corporateActionListCmd.Flags().String("since", "", "Start date (YYYY-MM-DD, required)")
	corporateActionListCmd.Flags().String("until", "", "End date (YYYY-MM-DD, required)")
	corporateActionListCmd.Flags().String("symbol", "", "Filter by symbol")
	corporateActionListCmd.Flags().String("date-type", "", "Date type: TRADING or SETTLEMENT")

	corporateActionCmd.AddCommand(corporateActionListCmd)
	corporateActionCmd.AddCommand(corporateActionGetCmd)
}
