package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/alpacahq/cli/internal/api"
	"github.com/alpacahq/cli/internal/cmdutil"
	"github.com/alpacahq/cli/internal/output"
	"github.com/spf13/cobra"
)

var orderSubmitCmd = &cobra.Command{
	Use:   "submit",
	Short: api.PostOrderOp.Summary(),
	Example: `  alpaca order submit --symbol AAPL --qty 10 --side buy --type market
  alpaca order submit --symbol AAPL --qty 5 --side buy --type limit --limit-price 185.00
  alpaca order submit --symbol AAPL --qty 10 --side sell --type stop --stop-price 175.00
  alpaca order submit --symbol AAPL --notional 1000 --side buy --type market`,
	RunE: func(cmd *cobra.Command, args []string) error {
		body := &api.PostOrderRequest{
			Symbol:         cmdutil.Str(cmd, "symbol"),
			Qty:            cmdutil.Str(cmd, "qty"),
			Notional:       cmdutil.Str(cmd, "notional"),
			Side:           api.OrderSide(cmdutil.Str(cmd, "side")),
			Type:           api.OrderType(cmdutil.Str(cmd, "type")),
			TimeInForce:    api.TimeInForce(cmdutil.Str(cmd, "time-in-force")),
			LimitPrice:     cmdutil.Str(cmd, "limit-price"),
			StopPrice:      cmdutil.Str(cmd, "stop-price"),
			TrailPercent:   cmdutil.Str(cmd, "trail-percent"),
			TrailPrice:     cmdutil.Str(cmd, "trail-price"),
			ExtendedHours:  cmdutil.Bool(cmd, "extended-hours"),
			ClientOrderID:  cmdutil.Str(cmd, "client-order-id"),
			OrderClass:     api.OrderClass(cmdutil.Str(cmd, "order-class")),
			PositionIntent: api.PositionIntent(cmdutil.Str(cmd, "position-intent")),
		}

		if cmdutil.Changed(cmd, "advanced-instructions") {
			if err := json.Unmarshal([]byte(cmdutil.Str(cmd, "advanced-instructions")), &body.AdvancedInstructions); err != nil {
				return fmt.Errorf("--advanced-instructions: %w", err)
			}
		}
		if cmdutil.Changed(cmd, "legs") {
			if err := json.Unmarshal([]byte(cmdutil.Str(cmd, "legs")), &body.Legs); err != nil {
				return fmt.Errorf("--legs: %w", err)
			}
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
		return output.Render(cmd.OutOrStdout(), getOutput(), order)
	},
}

var orderReplaceCmd = fetchCmd("replace", api.PatchOrderByOrderIDOp, func(cmd *cobra.Command, args []string) (any, error) {
	body, _ := patchOrderRequestBodyFromFlags(cmd)
	if cmdutil.Changed(cmd, "advanced-instructions") {
		if err := json.Unmarshal([]byte(cmdutil.Str(cmd, "advanced-instructions")), &body.AdvancedInstructions); err != nil {
			return nil, fmt.Errorf("--advanced-instructions: %w", err)
		}
	}
	return tradingClient.PatchOrderByOrderID(cmdutil.Str(cmd, "order-id"), body)
}, func(c *cobra.Command) {
	c.Example = `  alpaca order replace --order-id <id> --qty 20 --limit-price 190.00`
})

func init() {
	cmdutil.RegisterFlags(orderSubmitCmd, api.PostOrderOp.Flags(), &cmdutil.FlagOpts{
		Defaults: map[string]string{"type": "market"},
	})
	orderSubmitCmd.Flags().Bool("dry-run", false, "Print the request body without submitting")

	orderCmd.AddCommand(orderSubmitCmd)
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
