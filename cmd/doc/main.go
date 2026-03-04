package main

import (
	"fmt"
	"os"

	"github.com/alpacahq/cli/internal/cmd"
	"github.com/spf13/cobra/doc"
)

func main() {
	dir := "man"
	if len(os.Args) > 1 {
		dir = os.Args[1]
	}

	if err := os.MkdirAll(dir, 0o755); err != nil {
		fmt.Fprintf(os.Stderr, "creating output dir: %v\n", err)
		os.Exit(1)
	}

	header := &doc.GenManHeader{
		Title:   "ALPACA",
		Section: "1",
		Source:  "Alpaca CLI",
		Manual:  "Alpaca CLI Manual",
	}

	root := cmd.Root()
	root.DisableAutoGenTag = true

	if err := doc.GenManTree(root, header, dir); err != nil {
		fmt.Fprintf(os.Stderr, "generating man pages: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Man pages generated in %s/\n", dir)
}
