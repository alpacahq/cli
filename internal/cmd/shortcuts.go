package cmd

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/alpacahq/cli/internal/api"
	"github.com/alpacahq/cli/internal/cmdutil"
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

	body := &api.PostOrderRequest{
		Symbol: symbol,
		Side:   api.OrderSide(side),
		Type:   "market",
	}

	if len(args) > 1 {
		body.Qty = args[1]
	}

	notional := cmdutil.Str(cmd, "notional")
	if notional != "" {
		body.Notional = notional
		body.Qty = ""
	}

	if body.Qty == "" && notional == "" {
		return fmt.Errorf("either <qty> argument or --notional flag is required")
	}

	limitPrice := cmdutil.Str(cmd, "limit")
	if limitPrice != "" {
		body.Type = "limit"
		body.LimitPrice = limitPrice
	}

	tif := cmdutil.Str(cmd, "tif")
	if tif != "" {
		body.TimeInForce = api.TimeInForce(tif)
	} else if isCrypto(symbol) {
		body.TimeInForce = "gtc"
	} else {
		body.TimeInForce = "day"
	}

	body.ExtendedHours = cmdutil.Bool(cmd, "extended-hours")

	takeProfit := cmdutil.Str(cmd, "take-profit")
	stopLoss := cmdutil.Str(cmd, "stop-loss")
	if takeProfit != "" || stopLoss != "" {
		body.OrderClass = "bracket"
		if takeProfit != "" {
			body.TakeProfit = map[string]any{"limit_price": takeProfit}
		}
		if stopLoss != "" {
			body.StopLoss = map[string]any{"stop_price": stopLoss}
		}
	}

	order, err := tradingClient.PostOrder(body)
	if err != nil {
		return err
	}
	return output.PrintSingle(getOutput(), orderColumns(), order)
}

var priceCmd = &cobra.Command{
	Use:   "price <symbol>",
	Short: "Get latest price for a symbol",
	Example: `  alpaca price AAPL
  alpaca price BTC/USD`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		symbol := args[0]

		var path string
		var params url.Values
		if isCrypto(symbol) {
			path = "/v1beta3/crypto/us/snapshots"
			params = url.Values{"symbols": {symbol}}
		} else {
			path = "/v2/stocks/" + symbol + "/snapshot"
		}

		data, err := apiClient.GetData(path, params)
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

		snapshot := m
		if isCrypto(symbol) {
			if snapshots, ok := m["snapshots"].(map[string]any); ok {
				if s, ok := snapshots[symbol].(map[string]any); ok {
					snapshot = s
				}
			}
		}

		latestTrade, _ := snapshot["latestTrade"].(map[string]any)
		latestQuote, _ := snapshot["latestQuote"].(map[string]any)

		fmt.Printf("%s\n", symbol)
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
		positions, err := tradingClient.GetAllOpenPositions()
		if err != nil {
			return err
		}
		return output.Render(getOutput(), positionColumns(), positions)
	},
}

var ordersShortcut = &cobra.Command{
	Use:   "orders",
	Short: "List open orders (shortcut)",
	RunE: func(cmd *cobra.Command, args []string) error {
		orders, err := tradingClient.GetAllOrders(&api.GetAllOrdersParams{
			Status: "open",
		})
		if err != nil {
			return err
		}
		return output.Render(getOutput(), orderColumns(), orders)
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
