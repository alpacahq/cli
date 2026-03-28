package cmd

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/alpacahq/cli/internal/api"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var updateGolden = flag.Bool("update", false, "update golden files")

// TestGeneratedCodeIsUpToDate re-runs the code generator and verifies the
// output matches what's committed. If someone changes the specs or the
// generator without running `make generate`, this test fails.
func TestGeneratedCodeIsUpToDate(t *testing.T) {
	root := projectRoot()

	type genFile struct {
		dir  string
		name string
	}
	apiDir := filepath.Join(root, "internal", "api")
	cmdDir := filepath.Join(root, "internal", "cmd")

	generatedFiles := []genFile{
		{apiDir, "trading_types.gen.go"},
		{apiDir, "trading_client.gen.go"},
		{apiDir, "marketdata_types.gen.go"},
		{apiDir, "marketdata_client.gen.go"},
		{apiDir, "descriptions.gen.go"},
		{cmdDir, "params.gen.go"},
		{cmdDir, "commands.gen.go"},
	}

	snapshots := make(map[string][]byte, len(generatedFiles))
	for _, f := range generatedFiles {
		path := filepath.Join(f.dir, f.name)
		data, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("reading %s: %v", f.name, err)
		}
		snapshots[path] = data
	}

	cmd := exec.Command("go", "run", "./cmd/generate")
	cmd.Dir = root
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("generator failed: %v\n%s", err, out)
	}

	var drifted []string
	for _, f := range generatedFiles {
		path := filepath.Join(f.dir, f.name)
		after, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("reading %s after generate: %v", f.name, err)
		}
		if !bytes.Equal(snapshots[path], after) {
			drifted = append(drifted, f.name)
			_ = os.WriteFile(path, snapshots[path], 0o644)
		}
	}

	if len(drifted) > 0 {
		t.Errorf("generated code is out of date in: %v\nRun 'make generate' and commit the result.", drifted)
	}
}

func TestAllCommandsHaveExamples(t *testing.T) {
	var check func(*cobra.Command)
	check = func(cmd *cobra.Command) {
		if !cmd.HasSubCommands() && cmd.Example == "" {
			t.Errorf("command %q has no example", cmd.CommandPath())
		}
		for _, sub := range cmd.Commands() {
			check(sub)
		}
	}
	check(rootCmd)
}

func TestCommandTreeGolden(t *testing.T) {
	var buf bytes.Buffer
	var walk func(*cobra.Command)
	walk = func(cmd *cobra.Command) {
		if cmd.Hidden {
			return
		}
		if cmd.RunE != nil || cmd.Run != nil {
			fmt.Fprintln(&buf, cmd.CommandPath())
			cmd.NonInheritedFlags().VisitAll(func(f *pflag.Flag) {
				entry := fmt.Sprintf("  --%s (%s", f.Name, f.Value.Type())
				if f.DefValue != "" && f.DefValue != "false" && f.DefValue != "0" {
					entry += ", default=" + f.DefValue
				}
				entry += ")"
				fmt.Fprintln(&buf, entry)
			})
		}
		for _, sub := range cmd.Commands() {
			walk(sub)
		}
	}
	walk(rootCmd)

	compareGolden(t, "command_tree.golden", buf.Bytes())
}

func TestOpsGolden(t *testing.T) {
	var buf bytes.Buffer
	for _, op := range api.AllOps {
		fmt.Fprintf(&buf, "%s: %q\n", op.Name, op.Summary)
		for _, f := range op.Flags {
			fmt.Fprintf(&buf, "  %s (%s, %s", f.Name, f.Source, f.Type)
			if f.Required {
				buf.WriteString(", required")
			}
			if f.Default != "" {
				fmt.Fprintf(&buf, ", default=%s", f.Default)
			}
			fmt.Fprintf(&buf, ") -> %s\n", f.OASName)
		}
		buf.WriteByte('\n')
	}

	compareGolden(t, "ops.golden", buf.Bytes())
}

func compareGolden(t *testing.T, name string, actual []byte) {
	t.Helper()
	golden := filepath.Join("testdata", name)
	if *updateGolden {
		if err := os.MkdirAll("testdata", 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(golden, actual, 0o644); err != nil {
			t.Fatal(err)
		}
		t.Logf("updated %s", golden)
		return
	}
	expected, err := os.ReadFile(golden)
	if err != nil {
		t.Fatalf("missing golden file %s; run: go test ./internal/cmd -run %s -update", golden, t.Name())
	}
	if !bytes.Equal(expected, actual) {
		t.Errorf("%s drifted from golden file; run: go test ./internal/cmd -run %s -update\n\n%s",
			name, t.Name(), diff(expected, actual))
	}
}

func diff(want, got []byte) string {
	wantLines := bytes.Split(want, []byte("\n"))
	gotLines := bytes.Split(got, []byte("\n"))
	var buf bytes.Buffer
	max := len(wantLines)
	if len(gotLines) > max {
		max = len(gotLines)
	}
	shown := 0
	for i := 0; i < max && shown < 20; i++ {
		var w, g []byte
		if i < len(wantLines) {
			w = wantLines[i]
		}
		if i < len(gotLines) {
			g = gotLines[i]
		}
		if !bytes.Equal(w, g) {
			fmt.Fprintf(&buf, "line %d:\n  want: %s\n  got:  %s\n", i+1, w, g)
			shown++
		}
	}
	if shown >= 20 {
		buf.WriteString("... (more differences)\n")
	}
	return buf.String()
}
