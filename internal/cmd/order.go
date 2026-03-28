package cmd

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/alpacahq/cli/internal/api"
	"github.com/alpacahq/cli/internal/cmdutil"
	"github.com/spf13/cobra"
)

func configureOrderSubmit(cmd *cobra.Command) {
	cmd.Flags().Bool("dry-run", false, "Print the request body without submitting")
}

func postOrderHook(cmd *cobra.Command, body *api.PostOrderRequest) (any, error) {
	if body.TimeInForce == "" {
		body.TimeInForce = defaultTimeInForce(body.Symbol)
	}
	if err := applyBracket(body, cmdutil.Str(cmd, "take-profit"), cmdutil.Str(cmd, "stop-loss")); err != nil {
		return nil, err
	}
	if cmdutil.Bool(cmd, "dry-run") {
		return body, nil
	}
	return nil, nil
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
