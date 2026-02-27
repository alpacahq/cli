package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/alpacahq/cli/internal/client"
	"github.com/alpacahq/cli/internal/cmd"
)

var version = "dev"

func main() {
	cmd.SetVersion(version)
	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)

		var apiErr *client.APIError
		if errors.As(err, &apiErr) {
			os.Exit(apiErr.ExitCode())
		}
		os.Exit(1)
	}
}
