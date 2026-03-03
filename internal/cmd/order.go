package cmd

import (
	"encoding/json"
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
	Short: api.OperationSummary["postOrder"],
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

		applyBracket(body, cmdutil.Str(cmd, "take-profit"), cmdutil.Str(cmd, "stop-loss"))

		if cmdutil.Bool(cmd, "dry-run") {
			return output.JSON(cmd.OutOrStdout(), body)
		}

		order, err := tradingClient.PostOrder(body)
		if err != nil {
			return err
		}
		return output.PrintSingle(cmd.OutOrStdout(), getOutput(), orderColumns(), order)
	},
}

var orderListCmd = &cobra.Command{
	Use:   "list",
	Short: api.OperationSummary["getAllOrders"],
	Example: `  alpaca order list
  alpaca order list --status closed --limit 20
  alpaca order list --symbols AAPL,MSFT --start 2025-01-01`,
	RunE: func(cmd *cobra.Command, args []string) error {
		nested := cmdutil.Bool(cmd, "nested")
		params := &api.GetAllOrdersParams{
			Status:        cmdutil.Str(cmd, "status"),
			Symbols:       cmdutil.Str(cmd, "symbols"),
			After:         cmdutil.Str(cmd, "start"),
			Until:         cmdutil.Str(cmd, "end"),
			Limit:         cmdutil.Int(cmd, "limit"),
			Direction:     cmdutil.Str(cmd, "sort"),
			Nested:        nested,
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

		if !nested {
			return output.Render(cmd.OutOrStdout(), getOutput(), orderColumns(), orders)
		}

		format := getOutput()
		if format == "json" {
			return output.JSON(cmd.OutOrStdout(), orders)
		}
		return output.Render(cmd.OutOrStdout(), format, orderColumns(), expandOrderLegs(orders))
	},
}

var orderGetCmd = &cobra.Command{
	Use:   "get [order-id]",
	Short: api.OperationSummary["getOrderByOrderID"],
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
		return output.PrintSingle(cmd.OutOrStdout(), getOutput(), orderColumns(), order)
	},
}

var orderCancelCmd = &cobra.Command{
	Use:   "cancel <order-id>",
	Short: api.OperationSummary["deleteOrderByOrderID"],
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		warnLive()
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
	Short: api.OperationSummary["deleteAllOrders"],
	RunE: func(cmd *cobra.Command, args []string) error {
		warnLive()
		cancelled, err := tradingClient.DeleteAllOrders()
		if err != nil {
			return err
		}
		return output.JSON(cmd.OutOrStdout(), cancelled)
	},
}

var orderReplaceCmd = &cobra.Command{
	Use:   "replace <order-id>",
	Short: api.OperationSummary["patchOrderByOrderId"],
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
		return output.PrintSingle(cmd.OutOrStdout(), getOutput(), orderColumns(), order)
	},
}

func init() {
	orderSubmitCmd.Flags().String("qty", "", api.ParamDescription["postOrder.qty"])
	orderSubmitCmd.Flags().String("notional", "", api.ParamDescription["postOrder.notional"])
	orderSubmitCmd.Flags().String("side", "", api.ParamDescription["postOrder.side"])
	_ = orderSubmitCmd.RegisterFlagCompletionFunc("side", cobra.FixedCompletions(api.OrderSideValues, cobra.ShellCompDirectiveNoFileComp))
	orderSubmitCmd.Flags().String("type", "market", api.ParamDescription["postOrder.type"])
	_ = orderSubmitCmd.RegisterFlagCompletionFunc("type", cobra.FixedCompletions(api.OrderTypeValues, cobra.ShellCompDirectiveNoFileComp))
	orderSubmitCmd.Flags().String("tif", "", api.ParamDescription["postOrder.time_in_force"])
	_ = orderSubmitCmd.RegisterFlagCompletionFunc("tif", cobra.FixedCompletions(api.TimeInForceValues, cobra.ShellCompDirectiveNoFileComp))
	orderSubmitCmd.Flags().String("limit-price", "", api.ParamDescription["postOrder.limit_price"])
	orderSubmitCmd.Flags().String("stop-price", "", api.ParamDescription["postOrder.stop_price"])
	orderSubmitCmd.Flags().String("trail-percent", "", api.ParamDescription["postOrder.trail_percent"])
	orderSubmitCmd.Flags().String("trail-price", "", api.ParamDescription["postOrder.trail_price"])
	orderSubmitCmd.Flags().Bool("extended-hours", false, api.ParamDescription["postOrder.extended_hours"])
	orderSubmitCmd.Flags().String("take-profit", "", api.ParamDescription["postOrder.take_profit"])
	orderSubmitCmd.Flags().String("stop-loss", "", api.ParamDescription["postOrder.stop_loss"])
	orderSubmitCmd.Flags().String("client-order-id", "", api.ParamDescription["postOrder.client_order_id"])
	orderSubmitCmd.Flags().String("order-class", "", api.ParamDescription["postOrder.order_class"])
	_ = orderSubmitCmd.RegisterFlagCompletionFunc("order-class", cobra.FixedCompletions(api.OrderClassValues, cobra.ShellCompDirectiveNoFileComp))
	orderSubmitCmd.Flags().String("position-intent", "", api.ParamDescription["postOrder.position_intent"])
	_ = orderSubmitCmd.RegisterFlagCompletionFunc("position-intent", cobra.FixedCompletions(api.PositionIntentValues, cobra.ShellCompDirectiveNoFileComp))
	orderSubmitCmd.Flags().Bool("dry-run", false, "Print the request body without submitting")

	orderListCmd.Flags().String("status", "", api.ParamDescription["getAllOrders.status"])
	_ = orderListCmd.RegisterFlagCompletionFunc("status", cobra.FixedCompletions(api.GetAllOrdersParamsStatusValues, cobra.ShellCompDirectiveNoFileComp))
	orderListCmd.Flags().String("symbols", "", api.ParamDescription["getAllOrders.symbols"])
	cmdutil.AddDateRangeFlags(orderListCmd)
	cmdutil.AddLimitFlag(orderListCmd)
	cmdutil.AddSortFlag(orderListCmd, api.SortValues)
	orderListCmd.Flags().Bool("nested", false, api.ParamDescription["getAllOrders.nested"])
	orderListCmd.Flags().String("side", "", api.ParamDescription["getAllOrders.side"])
	_ = orderListCmd.RegisterFlagCompletionFunc("side", cobra.FixedCompletions(api.OrderSideValues, cobra.ShellCompDirectiveNoFileComp))
	orderListCmd.Flags().String("class", "", api.ParamDescription["getAllOrders.asset_class"])
	_ = orderListCmd.RegisterFlagCompletionFunc("class", cobra.FixedCompletions(api.AssetClassValues, cobra.ShellCompDirectiveNoFileComp))
	orderListCmd.Flags().String("before-order-id", "", api.ParamDescription["getAllOrders.before_order_id"])
	orderListCmd.Flags().String("after-order-id", "", api.ParamDescription["getAllOrders.after_order_id"])

	orderReplaceCmd.Flags().String("qty", "", api.ParamDescription["patchOrderByOrderId.qty"])
	orderReplaceCmd.Flags().String("limit-price", "", api.ParamDescription["patchOrderByOrderId.limit_price"])
	orderReplaceCmd.Flags().String("stop-price", "", api.ParamDescription["patchOrderByOrderId.stop_price"])
	orderReplaceCmd.Flags().String("tif", "", api.ParamDescription["patchOrderByOrderId.time_in_force"])
	_ = orderReplaceCmd.RegisterFlagCompletionFunc("tif", cobra.FixedCompletions(api.TimeInForceValues, cobra.ShellCompDirectiveNoFileComp))
	orderReplaceCmd.Flags().String("trail", "", api.ParamDescription["patchOrderByOrderId.trail"])
	orderReplaceCmd.Flags().String("client-order-id", "", api.ParamDescription["patchOrderByOrderId.client_order_id"])

	orderCmd.AddCommand(orderSubmitCmd)
	orderCmd.AddCommand(orderListCmd)
	orderGetCmd.Flags().String("client-id", "", "Look up order by client order ID")
	orderCmd.AddCommand(orderGetCmd)
	orderCmd.AddCommand(orderCancelCmd)
	orderCmd.AddCommand(orderCancelAllCmd)
	orderCmd.AddCommand(orderReplaceCmd)
}

func applyBracket(body *api.PostOrderRequest, takeProfit, stopLoss string) {
	if takeProfit == "" && stopLoss == "" {
		return
	}
	body.OrderClass = "bracket"
	if takeProfit != "" {
		body.TakeProfit = map[string]any{"limit_price": takeProfit}
	}
	if stopLoss != "" {
		body.StopLoss = map[string]any{"stop_price": stopLoss}
	}
}

func expandOrderLegs(orders []api.Order) []map[string]any {
	var rows []map[string]any
	for _, o := range orders {
		b, _ := json.Marshal(o)
		var row map[string]any
		_ = json.Unmarshal(b, &row)
		delete(row, "legs")
		rows = append(rows, row)

		for i, leg := range o.Legs {
			b, _ = json.Marshal(leg)
			var legRow map[string]any
			_ = json.Unmarshal(b, &legRow)
			prefix := " ├─ "
			if i == len(o.Legs)-1 {
				prefix = " └─ "
			}
			if id, ok := legRow["id"].(string); ok {
				legRow["id"] = prefix + id
			}
			delete(legRow, "legs")
			rows = append(rows, legRow)
		}
	}
	return rows
}
