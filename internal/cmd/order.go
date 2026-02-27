package cmd

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/alpacahq/cli/internal/output"
	"github.com/spf13/cobra"
)

var orderCmd = &cobra.Command{
	Use:   "order",
	Short: "Manage orders",
}

var orderSubmitCmd = &cobra.Command{
	Use:   "submit",
	Short: "Submit a new order",
	Example: `  alpaca order submit --symbol AAPL --qty 10 --side buy --type market
  alpaca order submit --symbol AAPL --qty 5 --side buy --type limit --limit-price 185.00
  alpaca order submit --symbol AAPL --qty 10 --side sell --type stop --stop-price 175.00
  alpaca order submit --symbol AAPL --notional 1000 --side buy --type market`,
	RunE: func(cmd *cobra.Command, args []string) error {
		symbol, _ := cmd.Flags().GetString("symbol")
		qty, _ := cmd.Flags().GetString("qty")
		notional, _ := cmd.Flags().GetString("notional")
		side, _ := cmd.Flags().GetString("side")
		orderType, _ := cmd.Flags().GetString("type")
		tif, _ := cmd.Flags().GetString("tif")
		limitPrice, _ := cmd.Flags().GetString("limit-price")
		stopPrice, _ := cmd.Flags().GetString("stop-price")
		trailPercent, _ := cmd.Flags().GetString("trail-percent")
		trailPrice, _ := cmd.Flags().GetString("trail-price")
		extendedHours, _ := cmd.Flags().GetBool("extended-hours")
		takeProfit, _ := cmd.Flags().GetString("take-profit")
		stopLoss, _ := cmd.Flags().GetString("stop-loss")
		clientOrderID, _ := cmd.Flags().GetString("client-order-id")

		if symbol == "" {
			return fmt.Errorf("--symbol is required")
		}
		if side == "" {
			return fmt.Errorf("--side is required (buy or sell)")
		}
		if qty == "" && notional == "" {
			return fmt.Errorf("either --qty or --notional is required")
		}

		body := map[string]any{
			"symbol": symbol,
			"side":   side,
			"type":   orderType,
		}

		if qty != "" {
			body["qty"] = qty
		}
		if notional != "" {
			body["notional"] = notional
		}
		if tif != "" {
			body["time_in_force"] = tif
		} else {
			body["time_in_force"] = "day"
		}
		if limitPrice != "" {
			body["limit_price"] = limitPrice
		}
		if stopPrice != "" {
			body["stop_price"] = stopPrice
		}
		if trailPercent != "" {
			body["trail_percent"] = trailPercent
		}
		if trailPrice != "" {
			body["trail_price"] = trailPrice
		}
		if extendedHours {
			body["extended_hours"] = true
		}
		if clientOrderID != "" {
			body["client_order_id"] = clientOrderID
		}

		orderClass, _ := cmd.Flags().GetString("order-class")
		if orderClass != "" {
			body["order_class"] = orderClass
		}
		positionIntent, _ := cmd.Flags().GetString("position-intent")
		if positionIntent != "" {
			body["position_intent"] = positionIntent
		}

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
	},
}

var orderListCmd = &cobra.Command{
	Use:   "list",
	Short: "List orders",
	Example: `  alpaca order list
  alpaca order list --status closed --limit 20
  alpaca order list --symbols AAPL,MSFT --after 2025-01-01`,
	RunE: func(cmd *cobra.Command, args []string) error {
		params := url.Values{}

		status, _ := cmd.Flags().GetString("status")
		if status != "" {
			params.Set("status", status)
		} else {
			params.Set("status", "open")
		}

		symbols, _ := cmd.Flags().GetString("symbols")
		if symbols != "" {
			params.Set("symbols", symbols)
		}
		after, _ := cmd.Flags().GetString("after")
		if after != "" {
			params.Set("after", after)
		}
		until, _ := cmd.Flags().GetString("until")
		if until != "" {
			params.Set("until", until)
		}
		limit, _ := cmd.Flags().GetString("limit")
		if limit != "" {
			params.Set("limit", limit)
		}
		direction, _ := cmd.Flags().GetString("direction")
		if direction != "" {
			params.Set("direction", direction)
		}

		nested, _ := cmd.Flags().GetBool("nested")
		if nested {
			params.Set("nested", "true")
		}
		side, _ := cmd.Flags().GetString("side")
		if side != "" {
			params.Set("side", side)
		}
		assetClass, _ := cmd.Flags().GetString("asset-class")
		if assetClass != "" {
			params.Set("asset_class", assetClass)
		}

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
			{Header: "STOP", Field: "stop_price"},
			{Header: "FILLED QTY", Field: "filled_qty"},
			{Header: "FILLED AVG", Field: "filled_avg_price"},
			{Header: "SUBMITTED", Field: "submitted_at"},
		}

		return output.Render(getOutput(), columns, data)
	},
}

var orderGetCmd = &cobra.Command{
	Use:   "get [order-id]",
	Short: "Get order details by ID or client order ID",
	Example: `  alpaca order get 61e69015-8549-4baf-b96f-9c4f3e8d0c35
  alpaca order get --client-id my-order-123`,
	Args: cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		clientID, _ := cmd.Flags().GetString("client-id")

		if len(args) == 0 && clientID == "" {
			return fmt.Errorf("either <order-id> argument or --client-id flag is required")
		}

		var data json.RawMessage
		var err error

		if clientID != "" {
			params := url.Values{}
			params.Set("client_order_id", clientID)
			data, err = apiClient.Get("/v2/orders:by_client_order_id", params)
		} else {
			data, err = apiClient.Get("/v2/orders/"+args[0], nil)
		}
		if err != nil {
			return err
		}
		return output.PrintSingle(getOutput(), orderColumns(), data)
	},
}

var orderCancelCmd = &cobra.Command{
	Use:   "cancel <order-id>",
	Short: "Cancel an order",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		_, err := apiClient.Delete("/v2/orders/"+args[0], nil)
		if err != nil {
			return err
		}
		fmt.Printf("Order %s cancelled.\n", args[0])
		return nil
	},
}

var orderCancelAllCmd = &cobra.Command{
	Use:   "cancel-all",
	Short: "Cancel all open orders",
	RunE: func(cmd *cobra.Command, args []string) error {
		data, err := apiClient.Delete("/v2/orders", nil)
		if err != nil {
			return err
		}
		return output.JSON(cmd.OutOrStdout(), data)
	},
}

var orderReplaceCmd = &cobra.Command{
	Use:   "replace <order-id>",
	Short: "Replace (modify) an existing order",
	Example: `  alpaca order replace <order-id> --qty 20 --limit-price 190.00`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		body := map[string]any{}

		if cmd.Flags().Changed("qty") {
			v, _ := cmd.Flags().GetString("qty")
			body["qty"] = v
		}
		if cmd.Flags().Changed("limit-price") {
			v, _ := cmd.Flags().GetString("limit-price")
			body["limit_price"] = v
		}
		if cmd.Flags().Changed("stop-price") {
			v, _ := cmd.Flags().GetString("stop-price")
			body["stop_price"] = v
		}
		if cmd.Flags().Changed("tif") {
			v, _ := cmd.Flags().GetString("tif")
			body["time_in_force"] = v
		}
		if cmd.Flags().Changed("trail") {
			v, _ := cmd.Flags().GetString("trail")
			body["trail"] = v
		}
		if cmd.Flags().Changed("client-order-id") {
			v, _ := cmd.Flags().GetString("client-order-id")
			body["client_order_id"] = v
		}

		if len(body) == 0 {
			return fmt.Errorf("at least one field to update is required")
		}

		data, err := apiClient.Patch("/v2/orders/"+args[0], body)
		if err != nil {
			return err
		}

		return output.PrintSingle(getOutput(), orderColumns(), data)
	},
}

func init() {
	orderSubmitCmd.Flags().String("symbol", "", "Ticker symbol")
	orderSubmitCmd.Flags().String("qty", "", "Number of shares")
	orderSubmitCmd.Flags().String("notional", "", "Dollar amount (fractional)")
	orderSubmitCmd.Flags().String("side", "", "buy or sell")
	orderSubmitCmd.Flags().String("type", "market", "Order type: market, limit, stop, stop_limit, trailing_stop")
	orderSubmitCmd.Flags().String("tif", "", "Time in force: day, gtc, ioc, fok, opg, cls (default: day)")
	orderSubmitCmd.Flags().String("limit-price", "", "Limit price")
	orderSubmitCmd.Flags().String("stop-price", "", "Stop price")
	orderSubmitCmd.Flags().String("trail-percent", "", "Trailing stop percent")
	orderSubmitCmd.Flags().String("trail-price", "", "Trailing stop price offset")
	orderSubmitCmd.Flags().Bool("extended-hours", false, "Allow extended hours trading")
	orderSubmitCmd.Flags().String("take-profit", "", "Take profit limit price (bracket order)")
	orderSubmitCmd.Flags().String("stop-loss", "", "Stop loss price (bracket order)")
	orderSubmitCmd.Flags().String("client-order-id", "", "Client-specified order ID")
	orderSubmitCmd.Flags().String("order-class", "", "Order class: simple, bracket, oco, oto, mleg")
	orderSubmitCmd.Flags().String("position-intent", "", "Position intent: buy_to_open, buy_to_close, sell_to_open, sell_to_close")

	orderListCmd.Flags().String("status", "", "Filter: open, closed, all (default: open)")
	orderListCmd.Flags().String("symbols", "", "Filter by symbols (comma-separated)")
	orderListCmd.Flags().String("after", "", "Filter: orders after this date")
	orderListCmd.Flags().String("until", "", "Filter: orders until this date")
	orderListCmd.Flags().String("limit", "", "Max number of orders to return")
	orderListCmd.Flags().String("direction", "", "Sort direction: asc or desc")
	orderListCmd.Flags().Bool("nested", false, "Include nested multi-leg order legs")
	orderListCmd.Flags().String("side", "", "Filter by side: buy or sell")
	orderListCmd.Flags().String("asset-class", "", "Filter by asset class: us_equity, us_option, crypto")

	orderReplaceCmd.Flags().String("qty", "", "New quantity")
	orderReplaceCmd.Flags().String("limit-price", "", "New limit price")
	orderReplaceCmd.Flags().String("stop-price", "", "New stop price")
	orderReplaceCmd.Flags().String("tif", "", "New time in force")
	orderReplaceCmd.Flags().String("trail", "", "New trail value")
	orderReplaceCmd.Flags().String("client-order-id", "", "New client order ID")

	orderCmd.AddCommand(orderSubmitCmd)
	orderCmd.AddCommand(orderListCmd)
	orderGetCmd.Flags().String("client-id", "", "Look up order by client order ID")

	orderCmd.AddCommand(orderGetCmd)
	orderCmd.AddCommand(orderCancelCmd)
	orderCmd.AddCommand(orderCancelAllCmd)
	orderCmd.AddCommand(orderReplaceCmd)
}

func orderColumns() []output.Column {
	return []output.Column{
		{Header: "ID", Field: "id"},
		{Header: "SYMBOL", Field: "symbol"},
		{Header: "SIDE", Field: "side"},
		{Header: "QTY", Field: "qty"},
		{Header: "TYPE", Field: "type"},
		{Header: "STATUS", Field: "status"},
		{Header: "LIMIT PRICE", Field: "limit_price"},
		{Header: "STOP PRICE", Field: "stop_price"},
		{Header: "FILLED QTY", Field: "filled_qty"},
		{Header: "FILLED AVG", Field: "filled_avg_price"},
		{Header: "TIME IN FORCE", Field: "time_in_force"},
		{Header: "SUBMITTED", Field: "submitted_at"},
	}
}
