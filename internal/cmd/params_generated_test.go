package cmd

import (
	"testing"

	"github.com/alpacahq/cli/internal/api"
	"github.com/alpacahq/cli/internal/cmdutil"
	"github.com/spf13/cobra"
)

func TestFromFlagsReadsRegisteredFlags(t *testing.T) {
	cmd := &cobra.Command{Use: "test"}
	cmdutil.RegisterFlags(cmd, api.GetAllOrdersOp.Flags(), nil)
	_ = cmd.Flags().Set("status", "closed")
	_ = cmd.Flags().Set("limit", "10")
	_ = cmd.Flags().Set("symbols", "AAPL,MSFT")

	params := getAllOrdersParamsFromFlags(cmd)
	if params.Status != "closed" {
		t.Errorf("Status = %q, want %q", params.Status, "closed")
	}
	if params.Limit != 10 {
		t.Errorf("Limit = %d, want %d", params.Limit, 10)
	}
	if params.Symbols != "AAPL,MSFT" {
		t.Errorf("Symbols = %q, want %q", params.Symbols, "AAPL,MSFT")
	}
}

func TestFromFlagsSkipsExcludedFlags(t *testing.T) {
	cmd := &cobra.Command{Use: "test"}
	cmdutil.RegisterFlags(cmd, api.GetOptionsContractsOp.Flags(), &cmdutil.FlagOpts{
		Exclude: map[string]bool{"underlying_symbols": true},
	})

	params := getOptionsContractsParamsFromFlags(cmd)
	if params.UnderlyingSymbols != "" {
		t.Errorf("UnderlyingSymbols should be empty when excluded, got %q", params.UnderlyingSymbols)
	}
}

func TestFromFlagsRespectsDefaults(t *testing.T) {
	cmd := &cobra.Command{Use: "test"}
	cmdutil.RegisterFlags(cmd, api.OptionBarsOp.Flags(), &cmdutil.FlagOpts{
		Defaults: map[string]string{"timeframe": "1Day"},
	})

	params := optionBarsParamsFromFlags(cmd)
	if params.Timeframe != "1Day" {
		t.Errorf("Timeframe = %q, want %q", params.Timeframe, "1Day")
	}
}

func TestFromFlagsIntZeroValue(t *testing.T) {
	cmd := &cobra.Command{Use: "test"}
	cmdutil.RegisterFlags(cmd, api.GetAllOrdersOp.Flags(), nil)
	_ = cmd.Flags().Set("limit", "0")

	params := getAllOrdersParamsFromFlags(cmd)
	if params.Limit != 0 {
		t.Errorf("Limit = %d, want 0", params.Limit)
	}
}

func TestFromFlagsBoolExplicitFalse(t *testing.T) {
	cmd := &cobra.Command{Use: "test"}
	cmdutil.RegisterFlags(cmd, api.GetAllOrdersOp.Flags(), nil)
	_ = cmd.Flags().Set("nested", "false")

	params := getAllOrdersParamsFromFlags(cmd)
	if params.Nested != false {
		t.Errorf("Nested = %v, want false", params.Nested)
	}
	if cmd.Flags().Lookup("nested") == nil {
		t.Error("nested flag should be registered")
	}
}

func TestFromFlagsFieldCoverage(t *testing.T) {
	tests := []struct {
		name  string
		flags []api.FlagDef
		check func(*cobra.Command) bool
	}{
		{
			name:  "GetAllOrders",
			flags: api.GetAllOrdersOp.Flags(),
			check: func(cmd *cobra.Command) bool {
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
				p := getAllOrdersParamsFromFlags(cmd)
				return p.Status == "open" && p.Symbols == "AAPL" &&
					p.After == "2025-01-01" && p.Until == "2025-12-31" &&
					p.Limit == 50 && p.Direction == "asc" &&
					p.Nested == true && p.Side == "buy" &&
					p.AssetClass == "us_equity" &&
					p.BeforeOrderID == "id1" && p.AfterOrderID == "id2"
			},
		},
		{
			name:  "News",
			flags: api.NewsOp.Flags(),
			check: func(cmd *cobra.Command) bool {
				_ = cmd.Flags().Set("symbols", "AAPL")
				_ = cmd.Flags().Set("limit", "5")
				_ = cmd.Flags().Set("include-content", "true")
				_ = cmd.Flags().Set("exclude-contentless", "true")
				_ = cmd.Flags().Set("sort", "desc")
				_ = cmd.Flags().Set("start", "2025-01-01")
				_ = cmd.Flags().Set("end", "2025-12-31")
				_ = cmd.Flags().Set("page-token", "tok")
				p := newsParamsFromFlags(cmd)
				return p.Symbols == "AAPL" && p.Limit == 5 &&
					p.IncludeContent == true && p.ExcludeContentless == true &&
					p.Sort == "desc" && p.Start == "2025-01-01" &&
					p.End == "2025-12-31" && p.PageToken == "tok"
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := &cobra.Command{Use: "test"}
			cmdutil.RegisterFlags(cmd, tt.flags, nil)
			if !tt.check(cmd) {
				t.Error("not all fields were populated correctly by FromFlags")
			}
		})
	}
}
