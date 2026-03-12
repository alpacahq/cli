package main

import (
	"runtime/debug"

	"github.com/alpacahq/cli/internal/cmd"
)

var version = "dev"

func main() {
	cmd.SetVersion(resolveVersion())
	_ = cmd.Execute()
}

// resolveVersion returns the build-time version if set via ldflags,
// otherwise falls back to the module version embedded by go install.
func resolveVersion() string {
	if version != "dev" {
		return version
	}
	if info, ok := debug.ReadBuildInfo(); ok {
		if v := info.Main.Version; v != "" && v != "(devel)" {
			return v
		}
	}
	return version
}
