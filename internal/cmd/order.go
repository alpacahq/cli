package cmd

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/alpacahq/cli/internal/api"
	"github.com/alpacahq/cli/internal/cmdutil"
	"github.com/alpacahq/cli/internal/output"
	"github.com/spf13/cobra"
)

var orderSubmitCmd = &cobra.Command{
	Use:   "submit",
	Short: api.PostOrderOp.Summary,
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
			body.TimeInForce = defaultTimeInForce(body.Symbol)
		}

		if err := applyBracket(body, cmdutil.Str(cmd, "take-profit"), cmdutil.Str(cmd, "stop-loss")); err != nil {
			return err
		}

		if cmdutil.Bool(cmd, "dry-run") {
			return output.JSON(cmd.OutOrStdout(), body)
		}

		order, err := tradingClient.PostOrder(body)
		if err != nil {
			return err
		}
		return renderData(cmd, order)
	},
}

func init() {
	cmdutil.RegisterFlags(orderSubmitCmd, api.PostOrderOp.Flags, api.PostOrderOp.Name, &cmdutil.FlagOpts{
		Defaults: map[string]string{"type": "market"},
	})
	orderSubmitCmd.Flags().Bool("dry-run", false, "Print the request body without submitting")

	orderCmd.AddCommand(orderSubmitCmd)
}

func defaultTimeInForce(symbol string) api.TimeInForce {
	if strings.Contains(symbol, "/") {
		return "gtc"
	}
	return "day"
}

func applyBracket(body *api.PostOrderRequest, takeProfit, stopLoss string) error {
	if takeProfit == "" && stopLoss == "" {
		return nil
	}
	body.OrderClass = "bracket"
	if takeProfit != "" {
		val, err := parseOrderObjectFlag("--take-profit", takeProfit, "limit_price")
		if err != nil {
			return err
		}
		body.TakeProfit = val
	}
	if stopLoss != "" {
		val, err := parseOrderObjectFlag("--stop-loss", stopLoss, "stop_price")
		if err != nil {
			return err
		}
		body.StopLoss = val
	}
	return nil
}

func parseOrderObjectFlag(flagName, raw, shorthandKey string) (map[string]any, error) {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return nil, nil
	}
	if strings.HasPrefix(raw, "{") {
		var out map[string]any
		if err := json.Unmarshal([]byte(raw), &out); err != nil {
			return nil, fmt.Errorf("%s: %w", flagName, err)
		}
		return out, nil
	}
	return map[string]any{shorthandKey: raw}, nil
}
