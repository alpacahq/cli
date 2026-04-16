package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/alpacahq/cli/internal/cmdutil"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

const (
	repoOwner = "alpacahq"
	repoName  = "cli"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Check for updates and (optionally) install them",
	Long: `Check for a newer version of the Alpaca CLI.

Queries GitHub for the latest release, detects your install method
(Homebrew or go install), and prompts to run the upgrade. Use --check
to print structured JSON without prompting.`,
	Example: `  alpaca update
  alpaca update --yes
  alpaca update --check`,
	RunE: func(cmd *cobra.Command, args []string) error {
		checkOnly := cmdutil.Bool(cmd, "check")
		assumeYes := cmdutil.Bool(cmd, "yes")

		latest, err := getLatestVersion(10 * time.Second)
		if err != nil {
			return fmt.Errorf("checking for updates: %w", err)
		}

		method := detectInstallMethod()
		current := version
		upToDate := !versionNewer(latest, current)
		upgradeCmd := upgradeCommand(method)

		if checkOnly {
			m := map[string]any{
				"current":          strings.TrimPrefix(current, "v"),
				"latest":           strings.TrimPrefix(latest, "v"),
				"update_available": !upToDate,
				"install_method":   method,
				"update_command":   upgradeCmd,
			}
			enc := json.NewEncoder(os.Stdout)
			enc.SetIndent("", "  ")
			return enc.Encode(m)
		}

		if upToDate {
			color.Green("Already up to date (%s)", current)
			return nil
		}

		fmt.Printf("Current version: %s\n", current)
		fmt.Printf("Latest version:  %s\n", latest)
		fmt.Println()

		if !assumeYes && !term.IsTerminal(int(os.Stdin.Fd())) {
			fmt.Println("To update, run:")
			fmt.Println("  " + upgradeCmd)
			return nil
		}

		if !assumeYes {
			fmt.Printf("Upgrade %s -> %s using `%s`? [y/N] ", current, latest, upgradeCmd)
			reader := bufio.NewReader(os.Stdin)
			ans, _ := reader.ReadString('\n')
			ans = strings.ToLower(strings.TrimSpace(ans))
			if ans != "y" && ans != "yes" {
				fmt.Println("Canceled. To update later, run:")
				fmt.Println("  " + upgradeCmd)
				return nil
			}
		}

		return runUpgrade(upgradeCmd)
	},
}

// runUpgrade shells out to the install-method-specific upgrade command and
// streams its output through to the user's terminal. We use sh -c so users
// see the same command they'd run by hand, without us having to parse its
// arguments.
func runUpgrade(upgradeCmd string) error {
	fmt.Fprintln(os.Stderr, "Running:", upgradeCmd)
	c := exec.Command("sh", "-c", upgradeCmd)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	c.Stdin = os.Stdin
	if err := c.Run(); err != nil {
		return fmt.Errorf("upgrade command failed: %w", err)
	}
	return nil
}

func init() {
	updateCmd.Flags().Bool("check", false, "Print update status as JSON without prompting")
	updateCmd.Flags().Bool("yes", false, "Skip the confirmation prompt and run the upgrade immediately")
}

// printUpdateNoticeIfAvailable does a best-effort GitHub check and prints a
// short notice when a newer release exists. It mirrors the doctor command's
// "Update:" section but stays silent on errors and when already up to date,
// since the bare `alpaca` command is mostly help output and any noise here
// would be surfaced on every invocation.
func printUpdateNoticeIfAvailable(w io.Writer, timeout time.Duration) {
	latest, err := getLatestVersion(timeout)
	if err != nil || !versionNewer(latest, version) {
		return
	}
	method := detectInstallMethod()
	fmt.Fprintln(w)
	fmt.Fprintln(w, "Update:")
	fmt.Fprintf(w, "  - update available: %s -> %s, run `%s` (or `alpaca update`)\n",
		version, latest, upgradeCommand(method))
}

type ghRelease struct {
	TagName string `json:"tag_name"`
}

func getLatestVersion(timeout time.Duration) (string, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/releases/latest", repoOwner, repoName)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("User-Agent", "alpaca-cli/"+version)
	resp, err := (&http.Client{Timeout: timeout}).Do(req)
	if err != nil {
		return "", err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("GitHub API returned %d", resp.StatusCode)
	}

	var release ghRelease
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		return "", err
	}

	return release.TagName, nil
}
