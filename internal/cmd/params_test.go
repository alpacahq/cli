package cmd

import (
	"testing"

	"github.com/alpacahq/cli/internal/api"
	"github.com/alpacahq/cli/internal/cmdutil"
	"github.com/spf13/cobra"
)

func TestQueryFromFlagsReadsRegisteredFlags(t *testing.T) {
	cmd := &cobra.Command{Use: "test"}
	cmdutil.RegisterFlags(cmd, api.GetAllOrdersOp.Flags(), "", nil)
	_ = cmd.Flags().Set("status", "closed")
	_ = cmd.Flags().Set("limit", "10")
	_ = cmd.Flags().Set("symbols", "AAPL,MSFT")

	v := queryFromFlags(cmd, api.GetAllOrdersOp)
	if got := v.Get("status"); got != "closed" {
		t.Errorf("status = %q, want %q", got, "closed")
	}
	if got := v.Get("limit"); got != "10" {
		t.Errorf("limit = %q, want %q", got, "10")
	}
	if got := v.Get("symbols"); got != "AAPL,MSFT" {
		t.Errorf("symbols = %q, want %q", got, "AAPL,MSFT")
	}
}

func TestQueryFromFlagsOmitsUnchanged(t *testing.T) {
	cmd := &cobra.Command{Use: "test"}
	cmdutil.RegisterFlags(cmd, api.OptionBarsOp.Flags(), "", &cmdutil.FlagOpts{
		Defaults: map[string]string{"timeframe": "1Day"},
	})

	v := queryFromFlags(cmd, api.OptionBarsOp)
	if got := v.Get("timeframe"); got != "" {
		t.Errorf("timeframe = %q, want empty (flag not explicitly set)", got)
	}

	_ = cmd.Flags().Set("timeframe", "1Hour")
	v = queryFromFlags(cmd, api.OptionBarsOp)
	if got := v.Get("timeframe"); got != "1Hour" {
		t.Errorf("timeframe = %q, want %q", got, "1Hour")
	}
}

func TestQueryFromFlagsIgnoresNonQueryFlags(t *testing.T) {
	cmd := &cobra.Command{Use: "test"}
	cmdutil.RegisterFlags(cmd, api.GetOrderByOrderIDOp.Flags(), "", nil)
	_ = cmd.Flags().Set("order-id", "abc-123")

	v := queryFromFlags(cmd, api.GetOrderByOrderIDOp)
	if v.Get("order_id") != "" {
		t.Error("path param should not appear in query values")
	}
}

func TestQueryFromFlagsIntValue(t *testing.T) {
	cmd := &cobra.Command{Use: "test"}
	cmdutil.RegisterFlags(cmd, api.GetAllOrdersOp.Flags(), "", nil)
	_ = cmd.Flags().Set("limit", "0")

	v := queryFromFlags(cmd, api.GetAllOrdersOp)
	if got := v.Get("limit"); got != "0" {
		t.Errorf("limit = %q, want %q", got, "0")
	}
}

func TestQueryFromFlagsBoolValue(t *testing.T) {
	cmd := &cobra.Command{Use: "test"}
	cmdutil.RegisterFlags(cmd, api.GetAllOrdersOp.Flags(), "", nil)
	_ = cmd.Flags().Set("nested", "true")

	v := queryFromFlags(cmd, api.GetAllOrdersOp)
	if got := v.Get("nested"); got != "true" {
		t.Errorf("nested = %q, want %q", got, "true")
	}
}

func TestQueryFromFlagsFieldCoverage(t *testing.T) {
	cmd := &cobra.Command{Use: "test"}
	cmdutil.RegisterFlags(cmd, api.GetAllOrdersOp.Flags(), "", nil)
	_ = cmd.Flags().Set("status", "open")
	_ = cmd.Flags().Set("symbols", "AAPL")
	_ = cmd.Flags().Set("after", "2025-01-01")
	_ = cmd.Flags().Set("until", "2025-12-31")
	_ = cmd.Flags().Set("limit", "50")
	_ = cmd.Flags().Set("direction", "asc")
	_ = cmd.Flags().Set("nested", "true")
	_ = cmd.Flags().Set("side", "buy")
	_ = cmd.Flags().Set("asset-class", "us_equity")
	_ = cmd.Flags().Set("before-order-id", "id1")
	_ = cmd.Flags().Set("after-order-id", "id2")

	v := queryFromFlags(cmd, api.GetAllOrdersOp)
	expected := map[string]string{
		"status":          "open",
		"symbols":         "AAPL",
		"after":           "2025-01-01",
		"until":           "2025-12-31",
		"limit":           "50",
		"direction":       "asc",
		"nested":          "true",
		"side":            "buy",
		"asset_class":     "us_equity",
		"before_order_id": "id1",
		"after_order_id":  "id2",
	}
	for k, want := range expected {
		if got := v.Get(k); got != want {
			t.Errorf("%s = %q, want %q", k, got, want)
		}
	}
}
