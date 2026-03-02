package cmd

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/alpacahq/cli/internal/cmdutil"

	"github.com/alpacahq/cli/internal/config"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

var profileCmd = &cobra.Command{
	Use:   "profile",
	Short: "Manage connection profiles",
}

var profileLoginCmd = &cobra.Command{
	Use:   "login",
	Short: "Create a profile and authenticate",
	Example: `  alpaca profile login
  alpaca profile login --key PKXXXXXXXX --secret XXXXXXXX
  alpaca profile login --name myaccount --live
  alpaca profile login --name dev --base-url https://custom-api.example.com
  alpaca profile login --name dev --base-url https://custom-api.example.com --data-url https://custom-data.example.com`,
	RunE: func(cmd *cobra.Command, args []string) error {
		key := cmdutil.Str(cmd, "key")
		secret := cmdutil.Str(cmd, "secret")
		name := cmdutil.Str(cmd, "name")
		dataURL := cmdutil.Str(cmd, "data-url")
		noValidate := cmdutil.Bool(cmd, "no-validate")

		if name == "" {
			name = "paper"
		}

		baseURL, err := resolveBaseURLFlags(cmd)
		if err != nil {
			return err
		}

		if cmdutil.Changed(cmd, "secret") {
			fmt.Fprintln(os.Stderr, "Warning: passing secrets via flags may expose them in shell history.")
			fmt.Fprintln(os.Stderr, "  Use `alpaca profile login` interactively or ALPACA_SECRET_KEY env var.")
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
			_ = resp.Body.Close()

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
			DataURL:   dataURL,
		}
		if err := config.SaveProfile(name, p); err != nil {
			return fmt.Errorf("saving profile: %w", err)
		}

		globalCfg := loadOrCreateGlobal()
		globalCfg.DefaultProfile = name
		if err := config.SaveGlobalConfig(globalCfg); err != nil {
			return fmt.Errorf("saving global config: %w", err)
		}

		color.Green("✓ Logged in to %s (%s)", name, baseURL)
		fmt.Fprintf(os.Stderr, "  Credentials stored in %s/profiles/\n", config.Dir())
		fmt.Fprintln(os.Stderr, "  For CI/automation, use ALPACA_API_KEY and ALPACA_SECRET_KEY env vars instead.")
		return nil
	},
}

var profileLogoutCmd = &cobra.Command{
	Use:   "logout [name]",
	Short: "Remove a profile",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := "paper"
		if len(args) > 0 {
			name = args[0]
		}
		if err := config.DeleteProfile(name); err != nil {
			if os.IsNotExist(err) {
				return fmt.Errorf("profile %q not found", name)
			}
			return err
		}
		fmt.Printf("Removed profile %s.\n", name)
		return nil
	},
}

var profileStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show the active profile",
	RunE: func(cmd *cobra.Command, args []string) error {
		resolved, err := config.Load(profileFlag, "")
		if err != nil {
			return err
		}

		fmt.Printf("Profile:   %s\n", resolved.ProfileName)
		fmt.Printf("Base URL:  %s\n", resolved.BaseURL)
		fmt.Printf("Data URL:  %s\n", resolved.DataURL)

		if resolved.HasCredentials() {
			masked := resolved.APIKey
			if len(masked) > 6 {
				masked = masked[:6] + strings.Repeat("*", len(masked)-6)
			}
			fmt.Printf("API Key:   %s\n", masked)
			color.Green("✓ Authenticated")
		} else {
			color.Yellow("✗ Not authenticated")
			fmt.Println("Hint: run `alpaca profile login` to set up credentials")
		}
		return nil
	},
}

var profileListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all profiles",
	RunE: func(cmd *cobra.Command, args []string) error {
		profiles, err := config.ListProfiles()
		if err != nil {
			return err
		}
		if len(profiles) == 0 {
			fmt.Println("No profiles configured.")
			fmt.Println("Hint: run `alpaca profile login` to create one")
			return nil
		}

		resolved, _ := config.Load("", "")
		for _, name := range profiles {
			if name == resolved.ProfileName {
				color.Green("* %s (active)", name)
			} else {
				fmt.Printf("  %s\n", name)
			}
		}
		return nil
	},
}

var profileSwitchCmd = &cobra.Command{
	Use:   "switch <name>",
	Short: "Switch the active profile",
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
			return fmt.Errorf("profile %q not found\nAvailable: %s\nHint: run `alpaca profile login --name %s` to create it", name, available, name)
		}

		globalCfg := loadOrCreateGlobal()
		globalCfg.DefaultProfile = name
		if err := config.SaveGlobalConfig(globalCfg); err != nil {
			return err
		}

		fmt.Printf("Switched to %s.\n", name)
		return nil
	},
}

var profileSetCmd = &cobra.Command{
	Use:   "set <key> <value>",
	Short: "Update a profile setting",
	Long: `Update a setting on the active profile.

Available keys:
  base_url    API base URL for trading
  data_url    API base URL for market data`,
	Example: `  alpaca profile set data_url https://data.example.com
  alpaca profile set base_url https://api.example.com`,
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		key, value := args[0], args[1]

		resolved, err := config.Load(profileFlag, "")
		if err != nil {
			return err
		}

		profile := config.LoadProfileByName(resolved.ProfileName)

		switch key {
		case "base_url":
			profile.BaseURL = value
		case "data_url":
			profile.DataURL = value
		default:
			return fmt.Errorf("unknown key: %s\nAvailable keys: base_url, data_url", key)
		}

		if err := config.SaveProfile(resolved.ProfileName, profile); err != nil {
			return err
		}
		fmt.Printf("Set %s = %s (profile: %s)\n", key, value, resolved.ProfileName)
		return nil
	},
}

func init() {
	profileLoginCmd.Flags().String("key", "", "API key")
	profileLoginCmd.Flags().String("secret", "", "Secret key")
	profileLoginCmd.Flags().String("name", "", "Profile name (default: paper)")
	profileLoginCmd.Flags().Bool("paper", false, "Use paper trading URL (default)")
	profileLoginCmd.Flags().Bool("live", false, "Use live trading URL")
	profileLoginCmd.Flags().String("base-url", "", "Custom API base URL")
	profileLoginCmd.MarkFlagsMutuallyExclusive("paper", "live", "base-url")
	profileLoginCmd.Flags().String("data-url", "", "Custom market data API URL")
	profileLoginCmd.Flags().Bool("no-validate", false, "Skip credential validation")

	profileCmd.AddCommand(profileLoginCmd)
	profileCmd.AddCommand(profileLogoutCmd)
	profileCmd.AddCommand(profileStatusCmd)
	profileCmd.AddCommand(profileListCmd)
	profileCmd.AddCommand(profileSwitchCmd)
	profileCmd.AddCommand(profileSetCmd)
}

func resolveBaseURLFlags(cmd *cobra.Command) (string, error) {
	if cmdutil.Bool(cmd, "live") {
		return config.ResolveBaseURL("live"), nil
	}
	if u := cmdutil.Str(cmd, "base-url"); u != "" {
		return config.ResolveBaseURL(u), nil
	}
	return config.ResolveBaseURL("paper"), nil
}

func loadOrCreateGlobal() *config.Config {
	resolved, _ := config.Load("", "")
	return &config.Config{
		DefaultProfile: resolved.ProfileName,
		Output:         resolved.Output,
		Color:          resolved.Color,
	}
}
