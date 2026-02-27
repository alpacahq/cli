package cmd

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/alpacahq/cli/internal/config"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Manage authentication",
}

var authLoginCmd = &cobra.Command{
	Use:   "login",
	Short: "Authenticate with Alpaca",
	Example: `  alpaca auth login
  alpaca auth login --key PKXXXXXXXX --secret XXXXXXXX
  alpaca auth login --profile myaccount --live
  alpaca auth login --profile dev --base-url https://custom-api.example.com
  alpaca auth login --profile dev --base-url https://custom-api.example.com --no-validate`,
	RunE: func(cmd *cobra.Command, args []string) error {
		key, _ := cmd.Flags().GetString("key")
		secret, _ := cmd.Flags().GetString("secret")
		profile, _ := cmd.Flags().GetString("profile")
		noValidate, _ := cmd.Flags().GetBool("no-validate")

		if profile == "" {
			profile = "paper"
		}

		baseURL, err := resolveBaseURLFlags(cmd)
		if err != nil {
			return err
		}

		if key == "" || secret == "" {
			reader := bufio.NewReader(os.Stdin)
			if key == "" {
				fmt.Print("API Key: ")
				key, _ = reader.ReadString('\n')
				key = strings.TrimSpace(key)
			}
			if secret == "" {
				fmt.Print("Secret Key: ")
				if term.IsTerminal(int(os.Stdin.Fd())) {
					raw, _ := term.ReadPassword(int(os.Stdin.Fd()))
					secret = string(raw)
					fmt.Println()
				} else {
					secret, _ = reader.ReadString('\n')
					secret = strings.TrimSpace(secret)
				}
			}
		}

		if key == "" || secret == "" {
			return fmt.Errorf("both API key and secret key are required")
		}

		if !noValidate {
			req, _ := http.NewRequest("GET", baseURL+"/v2/account", nil)
			req.Header.Set("APCA-API-KEY-ID", key)
			req.Header.Set("APCA-API-SECRET-KEY", secret)
			resp, err := (&http.Client{Timeout: 10 * time.Second}).Do(req)
			if err != nil {
				return fmt.Errorf("failed to connect to %s: %w\nHint: use --no-validate to skip credential check", baseURL, err)
			}
			resp.Body.Close()

			if resp.StatusCode == 401 || resp.StatusCode == 403 {
				return fmt.Errorf("invalid credentials (validated against %s)\nHint: use --base-url to specify the correct API endpoint, or --no-validate to skip", baseURL)
			}
			if resp.StatusCode >= 400 {
				return fmt.Errorf("unexpected response: HTTP %d from %s", resp.StatusCode, baseURL)
			}
		}

		p := &config.Profile{
			APIKey:    key,
			SecretKey: secret,
			BaseURL:   baseURL,
		}
		if err := config.SaveProfile(profile, p); err != nil {
			return fmt.Errorf("saving profile: %w", err)
		}

		globalCfg := loadOrCreateGlobal()
		globalCfg.DefaultProfile = profile
		config.SaveGlobalConfig(globalCfg)

		color.Green("✓ Logged in to %s profile (%s)", profile, baseURL)
		return nil
	},
}

var authLogoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Remove stored credentials",
	RunE: func(cmd *cobra.Command, args []string) error {
		profile, _ := cmd.Flags().GetString("profile")
		if profile == "" {
			profile = "paper"
		}
		if err := config.DeleteProfile(profile); err != nil {
			if os.IsNotExist(err) {
				return fmt.Errorf("profile %q not found", profile)
			}
			return err
		}
		fmt.Printf("Logged out of %s profile.\n", profile)
		return nil
	},
}

var authStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show current authentication status",
	RunE: func(cmd *cobra.Command, args []string) error {
		resolved, err := config.Load(profileFlag, "")
		if err != nil {
			return err
		}

		fmt.Printf("Profile:   %s\n", resolved.ProfileName)
		fmt.Printf("Base URL:  %s\n", resolved.BaseURL)

		if resolved.HasCredentials() {
			masked := resolved.APIKey
			if len(masked) > 6 {
				masked = masked[:6] + strings.Repeat("*", len(masked)-6)
			}
			fmt.Printf("API Key:   %s\n", masked)
			color.Green("✓ Authenticated")
		} else {
			color.Yellow("✗ Not authenticated")
			fmt.Println("Hint: run `alpaca auth login` to set up your credentials")
		}
		return nil
	},
}

var authSwitchCmd = &cobra.Command{
	Use:   "switch <profile>",
	Short: "Switch the default profile",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		profiles, err := config.ListProfiles()
		if err != nil {
			return err
		}

		found := false
		for _, p := range profiles {
			if p == name {
				found = true
				break
			}
		}
		if !found {
			available := "(none)"
			if len(profiles) > 0 {
				available = strings.Join(profiles, ", ")
			}
			return fmt.Errorf("profile %q not found\nAvailable profiles: %s\nHint: run `alpaca auth login --profile %s` to create it", name, available, name)
		}

		globalCfg := loadOrCreateGlobal()
		globalCfg.DefaultProfile = name
		if err := config.SaveGlobalConfig(globalCfg); err != nil {
			return err
		}

		fmt.Printf("Switched to %s profile.\n", name)
		return nil
	},
}

func init() {
	authLoginCmd.Flags().String("key", "", "API key")
	authLoginCmd.Flags().String("secret", "", "Secret key")
	authLoginCmd.Flags().String("profile", "", "Profile name (default: paper)")
	authLoginCmd.Flags().Bool("paper", false, "Use paper trading URL (default)")
	authLoginCmd.Flags().Bool("live", false, "Use live trading URL")
	authLoginCmd.Flags().String("base-url", "", "Custom API base URL")
	authLoginCmd.MarkFlagsMutuallyExclusive("paper", "live", "base-url")
	authLoginCmd.Flags().Bool("no-validate", false, "Skip credential validation")

	authLogoutCmd.Flags().String("profile", "", "Profile name (default: paper)")

	authCmd.AddCommand(authLoginCmd)
	authCmd.AddCommand(authLogoutCmd)
	authCmd.AddCommand(authStatusCmd)
	authCmd.AddCommand(authSwitchCmd)
}

func resolveBaseURLFlags(cmd *cobra.Command) (string, error) {
	live, _ := cmd.Flags().GetBool("live")
	baseURL, _ := cmd.Flags().GetString("base-url")

	if live {
		return "https://api.alpaca.markets", nil
	}
	if baseURL != "" {
		return strings.TrimRight(baseURL, "/"), nil
	}
	return "https://paper-api.alpaca.markets", nil
}

func loadOrCreateGlobal() *config.Config {
	resolved, _ := config.Load("", "")
	return &config.Config{
		DefaultProfile: resolved.ProfileName,
		Output:         resolved.Output,
		Color:          resolved.Color,
	}
}
