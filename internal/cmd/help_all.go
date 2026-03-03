package cmd

import (
	"fmt"
	"io"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var skipCommands = map[string]bool{
	"help":       true,
	"completion": true,
}

func printCommandTree(w io.Writer, root *cobra.Command, depth int) {
	if depth == 0 {
		printRootSection(w, root)
		fmt.Fprintln(w)
		for _, sub := range root.Commands() {
			if sub.IsAvailableCommand() && !skipCommands[sub.Name()] {
				printCommandTree(w, sub, 1)
			}
		}
		return
	}

	if !root.IsAvailableCommand() || skipCommands[root.Name()] {
		return
	}

	path := root.CommandPath()
	prefix := strings.Repeat("  ", depth-1)

	if root.HasAvailableSubCommands() {
		fmt.Fprintf(w, "%s── %s  —  %s\n", prefix, path, root.Short)
	} else {
		printLeafCommand(w, root, prefix)
	}

	for _, sub := range root.Commands() {
		if sub.IsAvailableCommand() && !skipCommands[sub.Name()] {
			printCommandTree(w, sub, depth+1)
		}
	}
}

func printRootSection(w io.Writer, cmd *cobra.Command) {
	fmt.Fprintf(w, "ALPACA CLI — Complete Reference\n")
	fmt.Fprintf(w, "%s\n\n", strings.Repeat("─", 40))

	if cmd.Long != "" {
		fmt.Fprintf(w, "%s\n\n", cmd.Long)
	}

	fmt.Fprintf(w, "GLOBAL FLAGS\n")
	cmd.NonInheritedFlags().VisitAll(func(f *pflag.Flag) {
		if f.Hidden || skipCommands[f.Name] {
			return
		}
		printFlag(w, f, "  ")
	})
}

func printLeafCommand(w io.Writer, cmd *cobra.Command, prefix string) {
	path := cmd.CommandPath()

	fmt.Fprintf(w, "%s── %s\n", prefix, path)
	fmt.Fprintf(w, "%s   %s\n", prefix, cmd.Short)
	fmt.Fprintf(w, "%s   Usage: %s\n", prefix, cmd.UseLine())

	localFlags := cmd.NonInheritedFlags()
	hasFlags := false
	localFlags.VisitAll(func(f *pflag.Flag) {
		if !f.Hidden {
			hasFlags = true
		}
	})

	if hasFlags {
		fmt.Fprintf(w, "%s   Flags:\n", prefix)
		localFlags.VisitAll(func(f *pflag.Flag) {
			if f.Hidden {
				return
			}
			printFlag(w, f, prefix+"     ")
		})
	}

	if cmd.Example != "" {
		fmt.Fprintf(w, "%s   Examples:\n", prefix)
		for _, line := range strings.Split(strings.TrimSpace(cmd.Example), "\n") {
			fmt.Fprintf(w, "%s     %s\n", prefix, strings.TrimSpace(line))
		}
	}

	fmt.Fprintln(w)
}

func printFlag(w io.Writer, f *pflag.Flag, indent string) {
	name := "--" + f.Name
	if f.Shorthand != "" {
		name = "-" + f.Shorthand + ", " + name
	}

	def := ""
	if f.DefValue != "" && f.DefValue != "false" && f.DefValue != "0" && f.DefValue != "[]" {
		def = " (default: " + f.DefValue + ")"
	}

	fmt.Fprintf(w, "%s%-24s %s%s\n", indent, name, f.Usage, def)
}
