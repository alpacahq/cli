package cmd

import (
	"bufio"
	"context"
	"fmt"
	"net/http"
	"os"
	"slices"
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
	Example: `  alpaca profile login                    # OAuth via browser (paper, default)
  alpaca profile login --scope trading    # OAuth with specific scopes
  alpaca profile login --api-key          # API key/secret login
  alpaca profile login --api-key --live   # API keys for live trading
  alpaca profile login --api-key --key PKXXXXXXXX --secret XXXXXXXX
  alpaca profile login --api-key --live --name prod`,
	RunE: func(cmd *cobra.Command, args []string) error {
		useAPIKey := cmdutil.Bool(cmd, "api-key")
		if useAPIKey {
			return loginWithAPIKey(cmd)
		}
		return loginWithOAuth(cmd)
	},
}

func loginWithOAuth(cmd *cobra.Command) error {
	if cmdutil.Bool(cmd, "live") {
		return fmt.Errorf("OAuth login is currently supported for paper trading only\nFor live trading, use API keys: alpaca profile login --api-key --live")
	}

	name := cmdutil.Str(cmd, "name")
	noValidate := cmdutil.Bool(cmd, "no-validate")
	scopeFlag := cmdutil.Str(cmd, "scope")

	if name == "" {
		name = defaultProfileName
	}

	baseURL := config.ResolveBaseURL(config.EnvPaper)
	env := defaultProfileName

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
		if err := validateCredentials(baseURL, map[string]string{
			"Authorization": "Bearer " + token.AccessToken,
		}); err != nil {
			return err
		}
	}

	p := &config.Profile{
		AccessToken: token.AccessToken,
		Scopes:      token.Scope,
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
	warnEnvShadowsProfile(name, "  ")
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
	noValidate := cmdutil.Bool(cmd, "no-validate")

	if name == "" {
		name = defaultProfileName
	}

	liveTrade := cmdutil.Bool(cmd, "live")
	envName := config.EnvPaper
	if liveTrade {
		envName = config.EnvLive
	}
	baseURL := config.ResolveBaseURL(envName)

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
		if err := validateCredentials(baseURL, map[string]string{
			"APCA-API-KEY-ID":     key,
			"APCA-API-SECRET-KEY": secret,
		}); err != nil {
			return err
		}
	}

	p := &config.Profile{
		APIKey:    key,
		SecretKey: secret,
	}
	if liveTrade {
		p.LiveTrade = &liveTrade
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
	warnEnvShadowsProfile(name, "  ")
	return nil
}

// envShadowsProfile reports whether env API keys will shadow the named
// profile. Env API keys beat any stored profile, so if env is set AND the
// profile has credentials, the profile is effectively dormant.
func envShadowsProfile(profileName string) bool {
	if os.Getenv("ALPACA_API_KEY") == "" || os.Getenv("ALPACA_SECRET_KEY") == "" {
		return false
	}
	p := config.LoadProfileByName(profileName)
	return p.AccessToken != "" || (p.APIKey != "" && p.SecretKey != "")
}

// warnEnvShadowsProfile prints the shadowing warning when envShadowsProfile
// returns true. Users who don't see it would wonder why "their profile"
// doesn't match their actual account. indent is the leading whitespace for
// the first line; the continuation line is indented two spaces further to
// align under the message. Pass "  " when rendering inside a nested block
// (profile login, doctor) and "" when rendering flush (bare alpaca help).
func warnEnvShadowsProfile(profileName, indent string) {
	if !envShadowsProfile(profileName) {
		return
	}
	color.Yellow(indent+"! ALPACA_API_KEY is set in your environment; it will override profile %q on every command.", profileName)
	fmt.Fprintln(os.Stderr, indent+"  Unset it (`unset ALPACA_API_KEY ALPACA_SECRET_KEY`) to use the profile.")
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
	Example: `  alpaca profile switch live
  alpaca profile switch paper`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		profiles, err := config.ListProfiles()
		if err != nil {
			return err
		}

		if !slices.Contains(profiles, name) {
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

func init() {
	profileLoginCmd.Flags().Bool("api-key", false, "Use API key/secret authentication instead of OAuth")
	profileLoginCmd.Flags().String("key", "", "API key (requires --api-key)")
	profileLoginCmd.Flags().String("secret", "", "Secret key (requires --api-key)")
	profileLoginCmd.Flags().String("scope", "", "OAuth scopes, comma-separated (default: all)")
	profileLoginCmd.Flags().String("name", "", "Profile name (default: paper)")
	profileLoginCmd.Flags().Bool("paper", false, "Target paper trading (default)")
	profileLoginCmd.Flags().Bool("live", false, "Target live trading")
	profileLoginCmd.MarkFlagsMutuallyExclusive("paper", "live")
	profileLoginCmd.Flags().Bool("no-validate", false, "Skip credential validation")

	profileCmd.AddCommand(profileLoginCmd)
	profileCmd.AddCommand(profileLogoutCmd)
	profileCmd.AddCommand(profileListCmd)
	profileCmd.AddCommand(profileSwitchCmd)
}

func validateCredentials(baseURL string, headers map[string]string) error {
	req, _ := http.NewRequest("GET", baseURL+"/v2/account", nil)
	req.Header.Set("User-Agent", "alpaca-cli/"+version)
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := (&http.Client{Timeout: 10 * time.Second}).Do(req)
	if err != nil {
		return fmt.Errorf("failed to connect to %s: %w\nHint: use --no-validate to skip credential check", baseURL, err)
	}
	_ = resp.Body.Close()

	if resp.StatusCode == 401 || resp.StatusCode == 403 {
		return fmt.Errorf("invalid credentials (validated against %s)\nHint: confirm --paper/--live matches the account, or use --no-validate to skip", baseURL)
	}
	if resp.StatusCode >= 400 {
		return fmt.Errorf("unexpected response: HTTP %d from %s", resp.StatusCode, baseURL)
	}
	return nil
}

func loadOrCreateGlobal() *config.Config {
	resolved, _ := config.Load("", "")
	return &config.Config{
		DefaultProfile: resolved.ProfileName,
		Output:         resolved.Output,
		Color:          resolved.Color,
	}
}
