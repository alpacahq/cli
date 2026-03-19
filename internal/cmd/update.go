package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/alpacahq/cli/internal/cmdutil"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

const (
	repoOwner = "alpacahq"
	repoName  = "cli"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Check for updates and show upgrade instructions",
	Long: `Check for a newer version of the Alpaca CLI.

Queries GitHub for the latest release, detects your install method
(Homebrew or go install), and prints the appropriate upgrade command.`,
	Example: `  alpaca update
  alpaca update --check`,
	RunE: func(cmd *cobra.Command, args []string) error {
		checkOnly := cmdutil.Bool(cmd, "check")

		latest, err := getLatestVersion()
		if err != nil {
			return fmt.Errorf("checking for updates: %w", err)
		}

		method := detectInstallMethod()
		saveUpdateState(&updateState{
			LatestVersion: latest,
			CheckedAt:     time.Now(),
			InstallMethod: method,
		})

		current := version
		upToDate := versionsEqual(current, latest)

		if checkOnly {
			m := map[string]any{
				"current":          strings.TrimPrefix(current, "v"),
				"latest":           strings.TrimPrefix(latest, "v"),
				"update_available": !upToDate,
				"install_method":   method,
				"update_command":   upgradeCommand(method),
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
		fmt.Println("To update, run:")
		fmt.Println("  " + upgradeCommand(method))
		return nil
	},
}

func init() {
	updateCmd.Flags().Bool("check", false, "Check for updates without installing")
}

type ghRelease struct {
	TagName string `json:"tag_name"`
}

func getLatestVersion() (string, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/releases/latest", repoOwner, repoName)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("User-Agent", "alpaca-cli/"+version)
	resp, err := (&http.Client{Timeout: 10 * time.Second}).Do(req)
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
