package cmd

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/spf13/cobra"
)

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
		{apiDir, "trading_types.go"},
		{apiDir, "trading_client.go"},
		{apiDir, "marketdata_types.go"},
		{apiDir, "marketdata_client.go"},
		{apiDir, "descriptions.go"},
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
