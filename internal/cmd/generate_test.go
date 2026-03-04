package cmd

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

// TestGeneratedCodeIsUpToDate re-runs the code generator and verifies the
// output matches what's committed. If someone changes the specs or the
// generator without running `make generate`, this test fails.
func TestGeneratedCodeIsUpToDate(t *testing.T) {
	root := projectRoot()
	dir := filepath.Join(root, "internal", "api")

	generatedFiles := []string{
		"trading_types.go",
		"trading_client.go",
		"marketdata_types.go",
		"marketdata_client.go",
		"descriptions.go",
	}

	snapshots := make(map[string][]byte, len(generatedFiles))
	for _, f := range generatedFiles {
		data, err := os.ReadFile(filepath.Join(dir, f))
		if err != nil {
			t.Fatalf("reading %s: %v", f, err)
		}
		snapshots[f] = data
	}

	cmd := exec.Command("go", "run", "./cmd/generate")
	cmd.Dir = root
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("generator failed: %v\n%s", err, out)
	}

	var drifted []string
	for _, f := range generatedFiles {
		path := filepath.Join(dir, f)
		after, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("reading %s after generate: %v", f, err)
		}
		if !bytes.Equal(snapshots[f], after) {
			drifted = append(drifted, f)
			_ = os.WriteFile(path, snapshots[f], 0o644)
		}
	}

	if len(drifted) > 0 {
		t.Errorf("generated code is out of date in: %v\nRun 'make generate' and commit the result.", drifted)
	}
}
