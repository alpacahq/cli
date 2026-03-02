package cmd

import (
	"fmt"

	"github.com/alpacahq/cli/internal/api"
	"github.com/alpacahq/cli/internal/cmdutil"
	"github.com/alpacahq/cli/internal/output"
	"github.com/spf13/cobra"
)

var orderCmd = &cobra.Command{
	Use:   "order",
	Short: "Manage orders",
}

var orderSubmitCmd = &cobra.Command{
	Use:   "submit <symbol>",
	Short: "Submit a new order",
	Example: `  alpaca order submit AAPL --qty 10 --side buy --type market
  alpaca order submit AAPL --qty 5 --side buy --type limit --limit-price 185.00
  alpaca order submit AAPL --qty 10 --side sell --type stop --stop-price 175.00
  alpaca order submit AAPL --notional 1000 --side buy --type market`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		warnLive()
		body := &api.PostOrderRequest{
			Symbol:        args[0],
			Qty:           cmdutil.Str(cmd, "qty"),
			Notional:      cmdutil.Str(cmd, "notional"),
			Side:          api.OrderSide(cmdutil.Str(cmd, "side")),
			Type:          api.OrderType(cmdutil.Str(cmd, "type")),
			TimeInForce:   api.TimeInForce(cmdutil.Str(cmd, "tif")),
			LimitPrice:    cmdutil.Str(cmd, "limit-price"),
			StopPrice:     cmdutil.Str(cmd, "stop-price"),
			TrailPercent:  cmdutil.Str(cmd, "trail-percent"),
			TrailPrice:    cmdutil.Str(cmd, "trail-price"),
			ExtendedHours: cmdutil.Bool(cmd, "extended-hours"),
			ClientOrderID: cmdutil.Str(cmd, "client-order-id"),
			OrderClass:    api.OrderClass(cmdutil.Str(cmd, "order-class")),
			PositionIntent: api.PositionIntent(cmdutil.Str(cmd, "position-intent")),
		}

		if body.TimeInForce == "" {
			body.TimeInForce = "day"
		}

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
	},
}

var orderListCmd = &cobra.Command{
	Use:   "list",
	Short: "List orders",
	Example: `  alpaca order list
  alpaca order list --status closed --limit 20
  alpaca order list --symbols AAPL,MSFT --start 2025-01-01`,
	RunE: func(cmd *cobra.Command, args []string) error {
		params := &api.GetAllOrdersParams{
			Status:        cmdutil.Str(cmd, "status"),
			Symbols:       cmdutil.Str(cmd, "symbols"),
			After:         cmdutil.Str(cmd, "start"),
			Until:         cmdutil.Str(cmd, "end"),
			Limit:         cmdutil.Int(cmd, "limit"),
			Direction:     cmdutil.Str(cmd, "sort"),
			Nested:        cmdutil.Bool(cmd, "nested"),
			Side:          cmdutil.Str(cmd, "side"),
			AssetClass:    cmdutil.Str(cmd, "class"),
			BeforeOrderID: cmdutil.Str(cmd, "before-order-id"),
			AfterOrderID:  cmdutil.Str(cmd, "after-order-id"),
		}
		if params.Status == "" {
			params.Status = "open"
		}

		orders, err := tradingClient.GetAllOrders(params)
		if err != nil {
			return err
		}
		return output.Render(getOutput(), orderColumns(), orders)
	},
}

var orderGetCmd = &cobra.Command{
	Use:   "get [order-id]",
	Short: "Get order details by ID or client order ID",
	Example: `  alpaca order get 61e69015-8549-4baf-b96f-9c4f3e8d0c35
  alpaca order get --client-id my-order-123`,
	Args: cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		clientID := cmdutil.Str(cmd, "client-id")

		if len(args) == 0 && clientID == "" {
			return fmt.Errorf("either <order-id> argument or --client-id flag is required")
		}

		var order *api.Order
		var err error

		if clientID != "" {
			order, err = tradingClient.GetOrderByClientOrderID(&api.GetOrderByClientOrderIDParams{
				ClientOrderID: clientID,
			})
		} else {
			order, err = tradingClient.GetOrderByOrderID(args[0], nil)
		}
		if err != nil {
			return err
		}
		return output.PrintSingle(getOutput(), orderColumns(), order)
	},
}

var orderCancelCmd = &cobra.Command{
	Use:   "cancel <order-id>",
	Short: "Cancel an order",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		_, err := tradingClient.DeleteOrderByOrderID(args[0])
		if err != nil {
			return err
		}
		_, _ = fmt.Fprintf(cmd.OutOrStdout(), "Order %s cancelled.\n", args[0])
		return nil
	},
}

var orderCancelAllCmd = &cobra.Command{
	Use:   "cancel-all",
	Short: "Cancel all open orders",
	RunE: func(cmd *cobra.Command, args []string) error {
		if isLive() {
			if err := requireConfirmation("Cancel ALL open orders on your live account?"); err != nil {
				return err
			}
		}
		cancelled, err := tradingClient.DeleteAllOrders()
		if err != nil {
			return err
		}
		return output.JSON(cmd.OutOrStdout(), cancelled)
	},
}

var orderReplaceCmd = &cobra.Command{
	Use:   "replace <order-id>",
	Short: "Replace (modify) an existing order",
	Example: `  alpaca order replace <order-id> --qty 20 --limit-price 190.00`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		body := &api.PatchOrderRequest{}

		if cmdutil.Changed(cmd, "qty") {
			body.Qty = cmdutil.Str(cmd, "qty")
		}
		if cmdutil.Changed(cmd, "limit-price") {
			body.LimitPrice = cmdutil.Str(cmd, "limit-price")
		}
		if cmdutil.Changed(cmd, "stop-price") {
			body.StopPrice = cmdutil.Str(cmd, "stop-price")
		}
		if cmdutil.Changed(cmd, "tif") {
			body.TimeInForce = api.TimeInForce(cmdutil.Str(cmd, "tif"))
		}
		if cmdutil.Changed(cmd, "trail") {
			body.Trail = cmdutil.Str(cmd, "trail")
		}
		if cmdutil.Changed(cmd, "client-order-id") {
			body.ClientOrderID = cmdutil.Str(cmd, "client-order-id")
		}

		order, err := tradingClient.PatchOrderByOrderID(args[0], body)
		if err != nil {
			return err
		}
		return output.PrintSingle(getOutput(), orderColumns(), order)
	},
}

func init() {
	orderSubmitCmd.Flags().String("qty", "", "Number of shares")
	orderSubmitCmd.Flags().String("notional", "", "Dollar amount (fractional)")
	orderSubmitCmd.Flags().String("side", "", "buy or sell")
	_ = orderSubmitCmd.RegisterFlagCompletionFunc("side", cobra.FixedCompletions([]string{"buy", "sell"}, cobra.ShellCompDirectiveNoFileComp))
	orderSubmitCmd.Flags().String("type", "market", "Order type: market, limit, stop, stop_limit, trailing_stop")
	_ = orderSubmitCmd.RegisterFlagCompletionFunc("type", cobra.FixedCompletions([]string{"market", "limit", "stop", "stop_limit", "trailing_stop"}, cobra.ShellCompDirectiveNoFileComp))
	orderSubmitCmd.Flags().String("tif", "", "Time in force: day, gtc, ioc, fok, opg, cls (default: day)")
	_ = orderSubmitCmd.RegisterFlagCompletionFunc("tif", cobra.FixedCompletions([]string{"day", "gtc", "ioc", "fok", "opg", "cls"}, cobra.ShellCompDirectiveNoFileComp))
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
	_ = orderListCmd.RegisterFlagCompletionFunc("status", cobra.FixedCompletions([]string{"open", "closed", "all"}, cobra.ShellCompDirectiveNoFileComp))
	orderListCmd.Flags().String("symbols", "", "Filter by symbols (comma-separated)")
	orderListCmd.Flags().String("start", "", "Filter: orders after this date")
	orderListCmd.Flags().String("end", "", "Filter: orders until this date")
	orderListCmd.Flags().Int("limit", 0, "Max number of orders to return")
	orderListCmd.Flags().String("sort", "", "Sort direction: asc or desc")
	_ = orderListCmd.RegisterFlagCompletionFunc("sort", cobra.FixedCompletions([]string{"asc", "desc"}, cobra.ShellCompDirectiveNoFileComp))
	orderListCmd.Flags().Bool("nested", false, "Include nested multi-leg order legs")
	orderListCmd.Flags().String("side", "", "Filter by side: buy or sell")
	_ = orderListCmd.RegisterFlagCompletionFunc("side", cobra.FixedCompletions([]string{"buy", "sell"}, cobra.ShellCompDirectiveNoFileComp))
	orderListCmd.Flags().String("class", "", "Filter by asset class: us_equity, us_option, crypto")
	_ = orderListCmd.RegisterFlagCompletionFunc("class", cobra.FixedCompletions([]string{"us_equity", "us_option", "crypto"}, cobra.ShellCompDirectiveNoFileComp))
	orderListCmd.Flags().String("before-order-id", "", "Cursor: orders before this order ID")
	orderListCmd.Flags().String("after-order-id", "", "Cursor: orders after this order ID")

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
