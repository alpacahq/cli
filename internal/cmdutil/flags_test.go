package cmdutil

import (
	"testing"

	"github.com/alpacahq/cli/internal/api"
	"github.com/spf13/cobra"
)

func newTestCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "test"}
	cmd.Flags().String("name", "default", "")
	cmd.Flags().Bool("verbose", false, "")
	cmd.Flags().Int("count", 0, "")
	cmd.Flags().Float64("price", 0, "")
	return cmd
}

func TestStr(t *testing.T) {
	cmd := newTestCmd()
	if got := Str(cmd, "name"); got != "default" {
		t.Errorf("Str(name) = %q, want %q", got, "default")
	}
}

func TestBool(t *testing.T) {
	cmd := newTestCmd()
	if got := Bool(cmd, "verbose"); got != false {
		t.Errorf("Bool(verbose) = %v, want false", got)
	}
}

func TestInt(t *testing.T) {
	cmd := newTestCmd()
	if got := Int(cmd, "count"); got != 0 {
		t.Errorf("Int(count) = %d, want 0", got)
	}
}

func TestFloat64(t *testing.T) {
	cmd := newTestCmd()
	if got := Float64(cmd, "price"); got != 0 {
		t.Errorf("Float64(price) = %f, want 0", got)
	}
}

func TestChanged(t *testing.T) {
	cmd := newTestCmd()
	if got := Changed(cmd, "name"); got != false {
		t.Errorf("Changed(name) = %v, want false (not set)", got)
	}
}

func TestRegisterFlags_TypeDispatch(t *testing.T) {
	cmd := &cobra.Command{Use: "test"}
	defs := []api.FlagDef{
		{Name: "name", OASName: "name", Type: "string", Default: "hello", Description: "A name"},
		{Name: "count", OASName: "count", Type: "int", Default: "42", Description: "A count"},
		{Name: "active", OASName: "active", Type: "bool", Default: "true", Description: "Active flag"},
		{Name: "price", OASName: "price", Type: "float64", Default: "9.99", Description: "A price"},
	}
	RegisterFlags(cmd, defs, nil)

	tests := []struct {
		flag     string
		wantType string
		wantDef  string
	}{
		{"name", "string", "hello"},
		{"count", "int", "42"},
		{"active", "bool", "true"},
		{"price", "float64", "9.99"},
	}
	for _, tt := range tests {
		f := cmd.Flags().Lookup(tt.flag)
		if f == nil {
			t.Errorf("flag %q not registered", tt.flag)
			continue
		}
		if f.Value.Type() != tt.wantType {
			t.Errorf("flag %q type = %s, want %s", tt.flag, f.Value.Type(), tt.wantType)
		}
		if f.DefValue != tt.wantDef {
			t.Errorf("flag %q default = %s, want %s", tt.flag, f.DefValue, tt.wantDef)
		}
	}
}

func TestRegisterFlags_Exclude(t *testing.T) {
	cmd := &cobra.Command{Use: "test"}
	defs := []api.FlagDef{
		{Name: "keep", OASName: "keep", Type: "string", Description: "Keep me"},
		{Name: "drop", OASName: "drop", Type: "string", Description: "Drop me"},
		{Name: "also-keep", OASName: "also_keep", Type: "string", Description: "Also keep"},
	}
	RegisterFlags(cmd, defs, &FlagOpts{
		Exclude: map[string]bool{"drop": true},
	})

	if cmd.Flags().Lookup("keep") == nil {
		t.Error("expected 'keep' to be registered")
	}
	if cmd.Flags().Lookup("drop") != nil {
		t.Error("expected 'drop' to be excluded")
	}
	if cmd.Flags().Lookup("also-keep") == nil {
		t.Error("expected 'also-keep' to be registered")
	}
}

func TestRegisterFlags_Aliases(t *testing.T) {
	cmd := &cobra.Command{Use: "test"}
	defs := []api.FlagDef{
		{Name: "activity-type", OASName: "activity_type", Type: "string", Description: "Activity type"},
	}
	RegisterFlags(cmd, defs, &FlagOpts{
		Aliases: map[string]string{"activity_type": "type"},
	})

	if cmd.Flags().Lookup("activity-type") != nil {
		t.Error("original name should not be registered when alias exists")
	}
	if cmd.Flags().Lookup("type") == nil {
		t.Error("alias 'type' should be registered")
	}
}

func TestRegisterFlags_DefaultOverride(t *testing.T) {
	cmd := &cobra.Command{Use: "test"}
	defs := []api.FlagDef{
		{Name: "timeframe", OASName: "timeframe", Type: "string", Default: "1Min", Description: "Timeframe"},
	}
	RegisterFlags(cmd, defs, &FlagOpts{
		Defaults: map[string]string{"timeframe": "1Day"},
	})

	f := cmd.Flags().Lookup("timeframe")
	if f == nil {
		t.Fatal("timeframe not registered")
	}
	if f.DefValue != "1Day" {
		t.Errorf("default = %s, want 1Day (overridden from 1Min)", f.DefValue)
	}
}

func TestRegisterFlags_Completions(t *testing.T) {
	cmd := &cobra.Command{Use: "test", Run: func(cmd *cobra.Command, args []string) {}}
	defs := []api.FlagDef{
		{
			Name:        "status",
			OASName:     "status",
			Type:        "string",
			Description: "Status filter",
			Completions: []string{"active", "inactive"},
		},
		{
			Name:        "name",
			OASName:     "name",
			Type:        "string",
			Description: "A name (no completions)",
		},
	}
	RegisterFlags(cmd, defs, nil)

	if cmd.Flags().Lookup("status") == nil {
		t.Error("status flag not registered")
	}
	if cmd.Flags().Lookup("name") == nil {
		t.Error("name flag not registered")
	}
	if cmd.Flags().Lookup("status").Usage != "Status filter" {
		t.Errorf("status description = %q", cmd.Flags().Lookup("status").Usage)
	}
}

func TestRegisterFlags_NilOpts(t *testing.T) {
	cmd := &cobra.Command{Use: "test"}
	defs := []api.FlagDef{
		{Name: "symbol", OASName: "symbol", Type: "string", Description: "Symbol"},
	}
	RegisterFlags(cmd, defs, nil)

	if cmd.Flags().Lookup("symbol") == nil {
		t.Error("flag not registered with nil opts")
	}
}

func TestRegisterFlags_EmptyDefs(t *testing.T) {
	cmd := &cobra.Command{Use: "test"}
	RegisterFlags(cmd, nil, nil)
	RegisterFlags(cmd, []api.FlagDef{}, nil)

	if cmd.Flags().HasFlags() {
		t.Error("expected no flags with empty defs")
	}
}
