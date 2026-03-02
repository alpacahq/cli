package main

import "github.com/alpacahq/cli/internal/cmd"

var version = "dev"

func main() {
	cmd.SetVersion(version)
	_ = cmd.Execute()
}
