package cmd

import (
	"fmt"
	"net/url"

	"github.com/alpacahq/cli/internal/output"
	"github.com/spf13/cobra"
)

var optionCmd = &cobra.Command{
	Use:   "option",
	Short: "Options trading",
}

var optionChainCmd = &cobra.Command{
	Use:   "chain <underlying>",
	Short: "Get option chain for an underlying symbol",
	Example: `  alpaca option chain AAPL
  alpaca option chain AAPL --expiry 2025-06-20 --type call
  alpaca option chain SPY --strike-gte 400 --strike-lte 450`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		params := url.Values{}
		params.Set("underlying_symbols", args[0])

		expiry, _ := cmd.Flags().GetString("expiry")
		if expiry != "" {
			params.Set("expiration_date", expiry)
		}
		optType, _ := cmd.Flags().GetString("type")
		if optType != "" {
			params.Set("type", optType)
		}
		strikeGte, _ := cmd.Flags().GetString("strike-gte")
		if strikeGte != "" {
			params.Set("strike_price_gte", strikeGte)
		}
		strikeLte, _ := cmd.Flags().GetString("strike-lte")
		if strikeLte != "" {
			params.Set("strike_price_lte", strikeLte)
		}
		limit, _ := cmd.Flags().GetString("limit")
		if limit != "" {
			params.Set("limit", limit)
		}

		data, err := apiClient.Get("/v2/options/contracts", params)
		if err != nil {
			return err
		}

		columns := []output.Column{
			{Header: "SYMBOL", Field: "symbol"},
			{Header: "TYPE", Field: "type"},
			{Header: "STRIKE", Field: "strike_price"},
			{Header: "EXPIRY", Field: "expiration_date"},
			{Header: "STATUS", Field: "status"},
			{Header: "UNDERLYING", Field: "underlying_symbol"},
		}

		return output.Render(getOutput(), columns, data)
	},
}

var optionGetCmd = &cobra.Command{
	Use:   "get <symbol-or-id>",
	Short: "Get option contract details",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		data, err := apiClient.Get("/v2/options/contracts/"+args[0], nil)
		if err != nil {
			return err
		}
		return output.JSON(cmd.OutOrStdout(), data)
	},
}

var optionExerciseCmd = &cobra.Command{
	Use:   "exercise <symbol-or-id>",
	Short: "Exercise an option position",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		_, err := apiClient.Post("/v2/positions/"+args[0]+"/exercise", nil)
		if err != nil {
			return err
		}
		fmt.Printf("Option %s exercise requested.\n", args[0])
		return nil
	},
}

func init() {
	optionChainCmd.Flags().String("expiry", "", "Expiration date (YYYY-MM-DD)")
	optionChainCmd.Flags().String("type", "", "Option type: call or put")
	optionChainCmd.Flags().String("strike-gte", "", "Minimum strike price")
	optionChainCmd.Flags().String("strike-lte", "", "Maximum strike price")
	optionChainCmd.Flags().String("limit", "", "Max results")

	optionCmd.AddCommand(optionChainCmd)
	optionCmd.AddCommand(optionGetCmd)
	optionCmd.AddCommand(optionExerciseCmd)
}
