package cmd

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"time"

	"github.com/alpacahq/cli/internal/client"
	"github.com/alpacahq/cli/internal/config"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var doctorCmd = &cobra.Command{
	Use:     "doctor",
	Short:   "Check CLI configuration and connectivity",
	Long:    "Run diagnostic checks on your Alpaca CLI setup: config files, credentials, and API connectivity.",
	Example: `  alpaca doctor`,
	RunE: func(cmd *cobra.Command, args []string) error {
		w := cmd.OutOrStdout()
		allOK := true
		usingEnvCredentials := os.Getenv("ALPACA_API_KEY") != "" && os.Getenv("ALPACA_SECRET_KEY") != ""

		fmt.Fprintf(w, "Alpaca CLI %s\n", version)
		fmt.Fprintf(w, "  Go:       %s\n", runtime.Version())
		fmt.Fprintf(w, "  OS/Arch:  %s/%s\n\n", runtime.GOOS, runtime.GOARCH)

		configDir := config.Dir()
		fmt.Fprintf(w, "Config:     %s\n", configDir)
		if _, err := os.Stat(configDir); os.IsNotExist(err) {
			if usingEnvCredentials {
				printCheck(w, true, "config directory does not exist (ok when using env vars)")
			} else {
				allOK = printCheck(w, false, "config directory does not exist")
			}
		} else {
			printCheck(w, true, "config directory exists")
		}

		profiles, _ := config.ListProfiles()
		if len(profiles) == 0 {
			if usingEnvCredentials {
				printCheck(w, true, "no saved profiles configured (using env var credentials)")
			} else {
				allOK = printCheck(w, false, "no profiles configured — run `alpaca profile login`")
			}
		} else {
			printCheck(w, true, fmt.Sprintf("%d profile(s): %s", len(profiles), joinMax(profiles, 5)))
		}

		resolved, err := config.Load("", "")
		if err != nil {
			allOK = printCheck(w, false, "failed to load config: "+err.Error())
			return doctorResult(w, allOK)
		}
		printCheck(w, true, "active profile: "+resolved.ProfileName)

		if !resolved.HasCredentials() {
			allOK = printCheck(w, false, "no credentials — run `alpaca profile login`")
			return doctorResult(w, allOK)
		}
		printCheck(w, true, credentialSourceDescription(resolved))
		if resolved.IsOAuth() && resolved.Scopes != "" {
			fmt.Fprintf(w, "  Scopes:   %s\n", resolved.Scopes)
		}
		warnEnvShadowsProfile(resolved.ProfileName, "  ")

		fmt.Fprintf(w, "\nConnectivity:\n")
		c := client.New(resolved)
		c.SetTimeout(10 * time.Second)

		fmt.Fprintf(w, "  Trading:  %s\n", resolved.BaseURL)
		if _, err := c.Get("/v2/account", nil); err != nil {
			allOK = printCheck(w, false, "trading API: "+err.Error())
		} else {
			printCheck(w, true, "trading API: connected")
		}

		fmt.Fprintf(w, "  Data:     %s\n", resolved.DataURL)
		if _, err := c.GetData("/v2/stocks/AAPL/trades/latest", nil); err != nil {
			allOK = printCheck(w, false, "data API: "+err.Error())
		} else {
			printCheck(w, true, "data API: connected")
		}

		fmt.Fprintf(w, "\nUpdate:\n")
		latest, err := getLatestVersion()
		if err != nil {
			fmt.Fprintf(w, "  - could not check for updates: %v\n", err)
		} else if !versionNewer(latest, version) {
			printCheck(w, true, fmt.Sprintf("up to date (%s)", version))
		} else {
			method := detectInstallMethod()
			fmt.Fprintf(w, "  - update available: %s → %s — run `%s`\n",
				version, latest, upgradeCommand(method))
		}

		return doctorResult(w, allOK)
	},
}

func printCheck(w io.Writer, ok bool, msg string) bool {
	if ok {
		color.New(color.FgGreen).Fprintf(w, "  ✓ ")
		fmt.Fprintln(w, msg)
	} else {
		color.New(color.FgRed).Fprintf(w, "  ✗ ")
		fmt.Fprintln(w, msg)
	}
	return ok
}

func doctorResult(w io.Writer, allOK bool) error {
	fmt.Fprintln(w)
	if allOK {
		color.New(color.FgGreen).Fprintln(w, "All checks passed.")
		return nil
	}
	return fmt.Errorf("some checks failed")
}

func joinMax(items []string, max int) string {
	if len(items) <= max {
		return fmt.Sprintf("%v", items)
	}
	shown := items[:max]
	return fmt.Sprintf("%v (+%d more)", shown, len(items)-max)
}

func credentialSourceDescription(r *config.Resolved) string {
	switch r.Source {
	case config.SourceEnvAPIKey:
		return "API key credentials from env (ALPACA_API_KEY + ALPACA_SECRET_KEY)"
	case config.SourceProfileOAuth:
		return fmt.Sprintf("OAuth token from profile %q", r.ProfileName)
	case config.SourceProfileAPIKey:
		return fmt.Sprintf("API key credentials from profile %q", r.ProfileName)
	default:
		return "credentials configured"
	}
}
