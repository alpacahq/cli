package cmd

import (
	"bufio"
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/alpacahq/cli/internal/cmdutil"
	"github.com/alpacahq/cli/internal/config"
	"github.com/alpacahq/cli/internal/oauth"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/yarlson/tap"
	"golang.org/x/term"
)

const defaultProfileName = config.EnvPaper

var profileCmd = &cobra.Command{
	Use:   "profile",
	Short: "Manage connection profiles",
}

var profileLoginCmd = &cobra.Command{
	Use:   "login",
	Short: "Authenticate via browser OAuth or API keys",
	Example: `  alpaca profile login                    # OAuth via browser (default)
  alpaca profile login --live             # OAuth for live account
  alpaca profile login --scope trading    # OAuth with specific scopes
  alpaca profile login --api-key          # API key/secret login
  alpaca profile login --api-key --key PKXXXXXXXX --secret XXXXXXXX
  alpaca profile login --name dev --base-url https://custom-api.example.com`,
	RunE: func(cmd *cobra.Command, args []string) error {
		useAPIKey := cmdutil.Bool(cmd, "api-key")
		if useAPIKey {
			return loginWithAPIKey(cmd)
		}
		return loginWithOAuth(cmd)
	},
}

func loginWithOAuth(cmd *cobra.Command) error {
	name := cmdutil.Str(cmd, "name")
	dataURL := cmdutil.Str(cmd, "data-url")
	noValidate := cmdutil.Bool(cmd, "no-validate")
	scopeFlag := cmdutil.Str(cmd, "scope")

	if name == "" {
		name = defaultProfileName
	}

	baseURL, err := resolveBaseURLFlags(cmd)
	if err != nil {
		return err
	}

	env := defaultProfileName
	if cmdutil.Bool(cmd, "live") {
		env = config.EnvLive
		if name == defaultProfileName {
			name = config.EnvLive
		}
	}

	scope := oauth.DefaultScopes
	if scopeFlag != "" {
		scope = strings.ReplaceAll(scopeFlag, ",", " ")
	} else if term.IsTerminal(int(os.Stdin.Fd())) {
		selected, err := promptScopes()
		if err != nil {
			return err
		}
		if len(selected) > 0 {
			scope = strings.Join(selected, " ")
		}
	}

	token, err := oauth.Login(env, scope)
	if err != nil {
		return err
	}

	if !noValidate {
		req, _ := http.NewRequest("GET", baseURL+"/v2/account", nil)
		req.Header.Set("Authorization", "Bearer "+token.AccessToken)
		if err := validateCredentials(req, baseURL); err != nil {
			return err
		}
	}

	p := &config.Profile{
		AccessToken: token.AccessToken,
		Scopes:      token.Scope,
		BaseURL:     baseURL,
		DataURL:     dataURL,
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
	if token.Scope != "" {
		fmt.Fprintf(os.Stderr, "  Scopes: %s\n", token.Scope)
	}
	fmt.Fprintf(os.Stderr, "  Credentials stored in %s/profiles/\n", config.Dir())
	return nil
}

type scopeOption struct {
	Value       string
	Description string
}

var availableScopes = []scopeOption{
	{"account:write", "Manage account settings and watchlists"},
	{"trading", "Place, cancel, and modify orders"},
	{"data", "Access market data"},
}

func promptScopes() ([]string, error) {
	allValues := make([]string, len(availableScopes))
	options := make([]tap.SelectOption[string], len(availableScopes))
	for i, s := range availableScopes {
		allValues[i] = s.Value
		options[i] = tap.SelectOption[string]{
			Value: s.Value,
			Label: s.Value,
			Hint:  s.Description,
		}
	}

	selected := tap.MultiSelect(context.Background(), tap.MultiSelectOptions[string]{
		Message:       "Select scopes to authorize (space to toggle, enter to confirm)",
		Options:       options,
		InitialValues: allValues,
	})
	if selected == nil {
		return nil, fmt.Errorf("login canceled")
	}
	return selected, nil
}

func loginWithAPIKey(cmd *cobra.Command) error {
	key := cmdutil.Str(cmd, "key")
	secret := cmdutil.Str(cmd, "secret")
	name := cmdutil.Str(cmd, "name")
	dataURL := cmdutil.Str(cmd, "data-url")
	noValidate := cmdutil.Bool(cmd, "no-validate")

	if name == "" {
		name = defaultProfileName
	}

	baseURL, err := resolveBaseURLFlags(cmd)
	if err != nil {
		return err
	}

	if cmdutil.Changed(cmd, "secret") {
		fmt.Fprintln(os.Stderr, "Warning: passing secrets via flags may expose them in shell history.")
		fmt.Fprintln(os.Stderr, "  Use `alpaca profile login --api-key` interactively or ALPACA_SECRET_KEY env var.")
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
		if err := validateCredentials(req, baseURL); err != nil {
			return err
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
}

var profileLogoutCmd = &cobra.Command{
	Use:   "logout [name]",
	Short: "Remove a profile",
	Example: `  alpaca profile logout
  alpaca profile logout live`,
	Args: cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := defaultProfileName
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
	Use:     "status",
	Short:   "Show the active profile",
	Example: `  alpaca profile status`,
	RunE: func(cmd *cobra.Command, args []string) error {
		resolved, err := config.Load(profileFlag, "")
		if err != nil {
			return err
		}

		fmt.Printf("Profile:   %s\n", resolved.ProfileName)
		fmt.Printf("Base URL:  %s\n", resolved.BaseURL)
		fmt.Printf("Data URL:  %s\n", resolved.DataURL)

		if resolved.HasCredentials() {
			if resolved.IsOAuth() {
				fmt.Printf("Auth:      OAuth (bearer token: %s)\n", maskCredential(resolved.AccessToken, 8))
				if resolved.Scopes != "" {
					fmt.Printf("Scopes:    %s\n", strings.ReplaceAll(resolved.Scopes, " ", ", "))
				}
			} else {
				fmt.Printf("Auth:      API key (%s)\n", maskCredential(resolved.APIKey, 6))
			}
			color.Green("✓ Authenticated")
		} else {
			color.Yellow("✗ Not authenticated")
			fmt.Println("Hint: run `alpaca profile login` to authenticate")
		}
		return nil
	},
}

var profileListCmd = &cobra.Command{
	Use:     "list",
	Short:   "List all profiles",
	Example: `  alpaca profile list`,
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
	profileLoginCmd.Flags().Bool("api-key", false, "Use API key/secret authentication instead of OAuth")
	profileLoginCmd.Flags().String("key", "", "API key (requires --api-key)")
	profileLoginCmd.Flags().String("secret", "", "Secret key (requires --api-key)")
	profileLoginCmd.Flags().String("scope", "", "OAuth scopes, comma-separated (default: all)")
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

func validateCredentials(req *http.Request, baseURL string) error {
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
	return nil
}

func resolveBaseURLFlags(cmd *cobra.Command) (string, error) {
	if cmdutil.Bool(cmd, "live") {
		return config.ResolveBaseURL(config.EnvLive), nil
	}
	if u := cmdutil.Str(cmd, "base-url"); u != "" {
		return config.ResolveBaseURL(u), nil
	}
	return config.ResolveBaseURL(defaultProfileName), nil
}

// maskCredential returns a string showing the first prefixLen characters
// followed by asterisks. The prefix is copied to a new string to break
// static-analysis taint tracking from the original credential.
func maskCredential(secret string, prefixLen int) string {
	n := len(secret)
	if n == 0 {
		return ""
	}
	if n <= prefixLen {
		return strings.Repeat("*", n)
	}
	prefix := make([]byte, prefixLen)
	copy(prefix, secret[:prefixLen])
	return string(prefix) + strings.Repeat("*", n-prefixLen)
}

func loadOrCreateGlobal() *config.Config {
	resolved, _ := config.Load("", "")
	return &config.Config{
		DefaultProfile: resolved.ProfileName,
		Output:         resolved.Output,
		Color:          resolved.Color,
	}
}
