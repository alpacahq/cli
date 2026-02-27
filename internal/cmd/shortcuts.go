package cmd

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/alpacahq/cli/internal/output"
	"github.com/spf13/cobra"
)

var buyCmd = &cobra.Command{
	Use:   "buy <symbol> <qty>",
	Short: "Buy shares (market order shortcut)",
	Example: `  alpaca buy AAPL 10
  alpaca buy AAPL 10 --limit 185.00
  alpaca buy AAPL --notional 1000
  alpaca buy AAPL 10 --tif gtc`,
	Args: cobra.RangeArgs(1, 2),
	RunE: func(cmd *cobra.Command, args []string) error {
		return submitShortcut(cmd, args, "buy")
	},
}

var sellCmd = &cobra.Command{
	Use:   "sell <symbol> <qty>",
	Short: "Sell shares (market order shortcut)",
	Example: `  alpaca sell AAPL 10
  alpaca sell AAPL 10 --limit 190.00`,
	Args: cobra.RangeArgs(1, 2),
	RunE: func(cmd *cobra.Command, args []string) error {
		return submitShortcut(cmd, args, "sell")
	},
}

func submitShortcut(cmd *cobra.Command, args []string, side string) error {
	symbol := args[0]

	body := map[string]any{
		"symbol": symbol,
		"side":   side,
		"type":   "market",
	}

	if len(args) > 1 {
		body["qty"] = args[1]
	}

	notional, _ := cmd.Flags().GetString("notional")
	if notional != "" {
		body["notional"] = notional
		delete(body, "qty")
	}

	if body["qty"] == nil && notional == "" {
		return fmt.Errorf("either <qty> argument or --notional flag is required")
	}

	limitPrice, _ := cmd.Flags().GetString("limit")
	if limitPrice != "" {
		body["type"] = "limit"
		body["limit_price"] = limitPrice
	}

	tif, _ := cmd.Flags().GetString("tif")
	if tif != "" {
		body["time_in_force"] = tif
	} else {
		body["time_in_force"] = "day"
	}

	extendedHours, _ := cmd.Flags().GetBool("extended-hours")
	if extendedHours {
		body["extended_hours"] = true
	}

	takeProfit, _ := cmd.Flags().GetString("take-profit")
	stopLoss, _ := cmd.Flags().GetString("stop-loss")
	if takeProfit != "" || stopLoss != "" {
		body["order_class"] = "bracket"
		if takeProfit != "" {
			body["take_profit"] = map[string]any{"limit_price": takeProfit}
		}
		if stopLoss != "" {
			body["stop_loss"] = map[string]any{"stop_price": stopLoss}
		}
	}

	data, err := apiClient.Post("/v2/orders", body)
	if err != nil {
		return err
	}

	return output.PrintSingle(getOutput(), orderColumns(), data)
}

var priceCmd = &cobra.Command{
	Use:   "price <symbol>",
	Short: "Get latest price for a symbol",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		data, err := apiClient.GetData("/v2/stocks/"+args[0]+"/snapshot", nil)
		if err != nil {
			return err
		}

		if getOutput() == "json" {
			return output.JSON(cmd.OutOrStdout(), data)
		}

		m := make(map[string]any)
		if err := json.Unmarshal(data, &m); err != nil {
			return err
		}

		latestTrade, _ := m["latestTrade"].(map[string]any)
		latestQuote, _ := m["latestQuote"].(map[string]any)

		fmt.Printf("%s\n", args[0])
		if latestTrade != nil {
			fmt.Printf("  Price:  $%v\n", latestTrade["p"])
		}
		if latestQuote != nil {
			fmt.Printf("  Bid:    $%v\n", latestQuote["bp"])
			fmt.Printf("  Ask:    $%v\n", latestQuote["ap"])
		}

		return nil
	},
}

var positionsShortcut = &cobra.Command{
	Use:   "positions",
	Short: "List all open positions (shortcut)",
	RunE: func(cmd *cobra.Command, args []string) error {
		data, err := apiClient.Get("/v2/positions", nil)
		if err != nil {
			return err
		}
		return output.Render(getOutput(), positionColumns(), data)
	},
}

var ordersShortcut = &cobra.Command{
	Use:   "orders",
	Short: "List open orders (shortcut)",
	RunE: func(cmd *cobra.Command, args []string) error {
		params := url.Values{}
		params.Set("status", "open")
		data, err := apiClient.Get("/v2/orders", params)
		if err != nil {
			return err
		}
		columns := []output.Column{
			{Header: "ID", Field: "id"},
			{Header: "SYMBOL", Field: "symbol"},
			{Header: "SIDE", Field: "side"},
			{Header: "QTY", Field: "qty"},
			{Header: "TYPE", Field: "type"},
			{Header: "STATUS", Field: "status"},
			{Header: "LIMIT", Field: "limit_price"},
			{Header: "FILLED QTY", Field: "filled_qty"},
			{Header: "SUBMITTED", Field: "submitted_at"},
		}
		return output.Render(getOutput(), columns, data)
	},
}

func init() {
	for _, cmd := range []*cobra.Command{buyCmd, sellCmd} {
		cmd.Flags().String("limit", "", "Limit price (converts to limit order)")
		cmd.Flags().String("notional", "", "Dollar amount instead of qty (fractional)")
		cmd.Flags().String("tif", "", "Time in force: day, gtc, ioc, fok")
		cmd.Flags().Bool("extended-hours", false, "Allow extended hours trading")
		cmd.Flags().String("take-profit", "", "Take profit price (bracket)")
		cmd.Flags().String("stop-loss", "", "Stop loss price (bracket)")
	}
}
