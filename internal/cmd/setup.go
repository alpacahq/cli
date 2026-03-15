package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/alpacahq/cli/internal/cmdutil"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

const (
	shellBash       = "bash"
	shellZsh        = "zsh"
	shellFish       = "fish"
	shellPowerShell = "powershell"

	osWindows = "windows"
)

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Install shell completions",
	Long: `Detects your shell and installs completions.

Runs automatically after 'alpaca update'. Works with any installation method
(go install, Homebrew, binary download).

Installed to user-level directories (no sudo required):
  Completions: ~/.local/share/bash-completion/, ~/.zsh/completions/, ~/.config/fish/`,
	Example: `  alpaca setup
  alpaca setup --shell zsh`,
	RunE: func(cmd *cobra.Command, args []string) error {
		shell := cmdutil.Str(cmd, "shell")

		if shell == "" {
			shell = detectShell()
		}

		if !isSupportedShell(shell) {
			return fmt.Errorf("unsupported shell %q\nSupported: bash, zsh, fish, powershell\nUse --shell to specify your shell explicitly", shell)
		}

		if err := installCompletions(shell); err != nil {
			fmt.Fprintf(os.Stderr, "Warning: completions install failed: %v\n", err)
		}

		return nil
	},
}

func init() {
	setupCmd.Flags().String("shell", "", "Shell to install completions for (bash, zsh, fish, powershell)")
	_ = setupCmd.RegisterFlagCompletionFunc("shell", cobra.FixedCompletions(
		[]string{shellBash, shellZsh, shellFish, shellPowerShell}, cobra.ShellCompDirectiveNoFileComp))

	rootCmd.AddCommand(setupCmd)
}

func isSupportedShell(shell string) bool {
	switch shell {
	case shellBash, shellZsh, shellFish, shellPowerShell:
		return true
	default:
		return false
	}
}

func detectShell() string {
	sh := os.Getenv("SHELL")
	if sh == "" {
		if runtime.GOOS == osWindows {
			return shellPowerShell
		}
		return shellBash
	}
	base := filepath.Base(sh)
	switch {
	case strings.Contains(base, "zsh"):
		return shellZsh
	case strings.Contains(base, "fish"):
		return shellFish
	default:
		return shellBash
	}
}

func installCompletions(shell string) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("cannot determine home directory: %w", err)
	}

	var path string
	var generate func() (string, error)

	switch shell {
	case shellZsh:
		path = filepath.Join(home, ".zsh", "completions", "_alpaca")
		generate = func() (string, error) {
			var buf strings.Builder
			if err := rootCmd.GenZshCompletion(&buf); err != nil {
				return "", err
			}
			return buf.String(), nil
		}
	case shellFish:
		path = filepath.Join(home, ".config", "fish", "completions", "alpaca.fish")
		generate = func() (string, error) {
			var buf strings.Builder
			if err := rootCmd.GenFishCompletion(&buf, true); err != nil {
				return "", err
			}
			return buf.String(), nil
		}
	case shellPowerShell:
		path = filepath.Join(home, ".config", "powershell", "completions", "alpaca.ps1")
		generate = func() (string, error) {
			var buf strings.Builder
			if err := rootCmd.GenPowerShellCompletionWithDesc(&buf); err != nil {
				return "", err
			}
			return buf.String(), nil
		}
	default:
		path = filepath.Join(home, ".local", "share", "bash-completion", "completions", "alpaca")
		generate = func() (string, error) {
			var buf strings.Builder
			if err := rootCmd.GenBashCompletionV2(&buf, true); err != nil {
				return "", err
			}
			return buf.String(), nil
		}
	}

	content, err := generate()
	if err != nil {
		return fmt.Errorf("generating %s completions: %w", shell, err)
	}

	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return fmt.Errorf("creating completions dir: %w", err)
	}

	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		return fmt.Errorf("writing completions: %w", err)
	}

	color.Green("✓ %s completions installed to %s", shell, path)

	switch shell {
	case shellZsh:
		if !zshFpathConfigured(home) {
			fmt.Fprintln(os.Stderr, "  Add to ~/.zshrc if not already present:")
			fmt.Fprintln(os.Stderr, "    fpath=(~/.zsh/completions $fpath)")
			fmt.Fprintln(os.Stderr, "    autoload -Uz compinit && compinit")
		}
	case shellBash:
		fmt.Fprintln(os.Stderr, "  Completions will load automatically if bash-completion is installed.")
		fmt.Fprintln(os.Stderr, "  Otherwise add to ~/.bashrc: source "+path)
	case shellFish:
		fmt.Fprintln(os.Stderr, "  Fish loads completions automatically from this directory.")
	}

	return nil
}

func zshFpathConfigured(home string) bool {
	data, err := os.ReadFile(filepath.Join(home, ".zshrc"))
	if err != nil {
		return false
	}
	return strings.Contains(string(data), ".zsh/completions")
}
