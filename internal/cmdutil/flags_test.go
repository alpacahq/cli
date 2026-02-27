package cmdutil

import (
	"testing"

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
