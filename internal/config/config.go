package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

const (
	EnvLive  = "live"
	EnvPaper = "paper"

	paperTradingURL = "https://paper-api.alpaca.markets"
	liveTradingURL  = "https://api.alpaca.markets"
	marketDataURL   = "https://data.alpaca.markets"
)

// Source identifies where the resolved credentials came from.
// Credentials resolve as an atomic bundle - never mixed across sources -
// so the winning source fully determines which auth headers are sent.
type Source string

const (
	SourceNone          Source = ""
	SourceEnvAPIKey     Source = "env-apikey"
	SourceProfileOAuth  Source = "profile-oauth"
	SourceProfileAPIKey Source = "profile-apikey"
)

type Config struct {
	DefaultProfile string `yaml:"default_profile"`
	Output         string `yaml:"output"`
	Color          string `yaml:"color"`
}

type Profile struct {
	APIKey      string `yaml:"api_key"`
	SecretKey   string `yaml:"secret_key"`
	AccessToken string `yaml:"access_token,omitempty"`
	Scopes      string `yaml:"scopes,omitempty"`
	// PaperTrade records whether this profile targets paper or live trading.
	// Pointer so we can distinguish "not specified" (nil -> safe paper default)
	// from "explicitly live" (false).
	PaperTrade *bool `yaml:"paper_trade,omitempty"`
}

type Resolved struct {
	APIKey      string
	SecretKey   string
	AccessToken string
	Scopes      string
	BaseURL     string
	DataURL     string
	Output      string
	Color       string
	ProfileName string
	Source      Source
}

func Dir() string {
	if d := os.Getenv("ALPACA_CONFIG_DIR"); d != "" {
		return d
	}
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".config", "alpaca")
}

// Load resolves credentials and URLs from env vars and the named profile.
//
// Credentials resolve as an atomic bundle - the first complete source wins,
// and field-level mixing across sources is not allowed. Order:
//  1. env ALPACA_API_KEY + ALPACA_SECRET_KEY (both required)
//  2. profile access_token
//  3. profile api_key + secret_key (both required)
//
// Paper-vs-live resolves independently:
//  1. ALPACA_PAPER_TRADE boolean (matches the MCP server's env var)
//  2. profile.paper_trade (only when credentials came from the profile)
//  3. paper default
//
// The paper default is deliberate: scripts and agents that forget to opt
// into live should hit paper, not live.
func Load(profileFlag, outputFlag string) (*Resolved, error) {
	cfg := loadGlobalConfig()
	profileName := resolve(profileFlag, os.Getenv("ALPACA_PROFILE"), cfg.DefaultProfile, EnvPaper)
	profile := loadProfile(profileName)

	r := &Resolved{
		ProfileName: profileName,
		DataURL:     marketDataURL,
		Output:      resolve(outputFlag, os.Getenv("ALPACA_OUTPUT"), cfg.Output, "json"),
		Color:       resolve(cfg.Color, "auto"),
	}

	resolveCredentials(r, &profile)
	resolveBaseURL(r, &profile)

	return r, nil
}

// resolveCredentials picks an atomic credential bundle - env or profile,
// never a mix. Sets Source to record which bundle won.
func resolveCredentials(r *Resolved, profile *Profile) {
	envKey := os.Getenv("ALPACA_API_KEY")
	envSecret := os.Getenv("ALPACA_SECRET_KEY")

	switch {
	case envKey != "" && envSecret != "":
		r.APIKey = envKey
		r.SecretKey = envSecret
		r.Source = SourceEnvAPIKey
	case profile.AccessToken != "":
		r.AccessToken = profile.AccessToken
		r.Scopes = profile.Scopes
		r.Source = SourceProfileOAuth
	case profile.APIKey != "" && profile.SecretKey != "":
		r.APIKey = profile.APIKey
		r.SecretKey = profile.SecretKey
		r.Source = SourceProfileAPIKey
	default:
		r.Source = SourceNone
	}
}

// resolveBaseURL determines the trading API URL. Defaults to paper unless
// something explicitly asks for live. Profile credentials honor the profile's
// paper_trade field.
func resolveBaseURL(r *Resolved, profile *Profile) {
	if pt := os.Getenv("ALPACA_PAPER_TRADE"); pt != "" {
		if isPaper(pt) {
			r.BaseURL = ResolveBaseURL(EnvPaper)
		} else {
			r.BaseURL = ResolveBaseURL(EnvLive)
		}
		return
	}
	profileIsLive := (r.Source == SourceProfileOAuth || r.Source == SourceProfileAPIKey) &&
		profile.PaperTrade != nil && !*profile.PaperTrade
	if profileIsLive {
		r.BaseURL = ResolveBaseURL(EnvLive)
		return
	}
	r.BaseURL = ResolveBaseURL(EnvPaper)
}

// isPaper interprets ALPACA_PAPER_TRADE. Matches MCP server semantics:
// case-insensitive "true" = paper, anything else = live.
func isPaper(v string) bool {
	return strings.EqualFold(strings.TrimSpace(v), "true")
}

func (r *Resolved) HasCredentials() bool {
	return r.Source != SourceNone
}

func (r *Resolved) IsOAuth() bool {
	return r.Source == SourceProfileOAuth
}

func (r *Resolved) Validate() error {
	if !r.HasCredentials() {
		return fmt.Errorf("authentication required\nHint: run `alpaca profile login` to authenticate")
	}
	return nil
}

// ResolveBaseURL takes a value that is either a well-known alias
// ("paper", "live") or a full URL, and returns a URL.
// Empty string defaults to the paper trading URL.
func ResolveBaseURL(value string) string {
	switch strings.ToLower(strings.TrimSpace(value)) {
	case EnvLive:
		return liveTradingURL
	case "", EnvPaper:
		return paperTradingURL
	default:
		return strings.TrimRight(value, "/")
	}
}

func loadGlobalConfig() Config {
	var cfg Config
	path := filepath.Join(Dir(), "config.yaml")
	data, err := os.ReadFile(path)
	if err != nil {
		return cfg
	}
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		fmt.Fprintf(os.Stderr, "Warning: failed to parse %s: %v\n", path, err)
	}
	return cfg
}

func LoadProfileByName(name string) *Profile {
	p := loadProfile(name)
	return &p
}

func loadProfile(name string) Profile {
	var p Profile
	path := filepath.Join(Dir(), "profiles", name+".yaml")
	data, err := os.ReadFile(path)
	if err != nil {
		return p
	}
	if err := yaml.Unmarshal(data, &p); err != nil {
		fmt.Fprintf(os.Stderr, "Warning: failed to parse %s: %v\n", path, err)
	}
	return p
}

func SaveGlobalConfig(cfg *Config) error {
	dir := Dir()
	if err := os.MkdirAll(dir, 0o700); err != nil {
		return err
	}
	data, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(dir, "config.yaml"), data, 0o600)
}

func SaveProfile(name string, p *Profile) error {
	dir := filepath.Join(Dir(), "profiles")
	if err := os.MkdirAll(dir, 0o700); err != nil {
		return err
	}
	data, err := yaml.Marshal(p)
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(dir, name+".yaml"), data, 0o600)
}

func DeleteProfile(name string) error {
	return os.Remove(filepath.Join(Dir(), "profiles", name+".yaml"))
}

func ListProfiles() ([]string, error) {
	entries, err := os.ReadDir(filepath.Join(Dir(), "profiles"))
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}
	var names []string
	for _, e := range entries {
		if !e.IsDir() && filepath.Ext(e.Name()) == ".yaml" {
			names = append(names, e.Name()[:len(e.Name())-5])
		}
	}
	return names, nil
}

// resolve returns the first non-empty value.
func resolve(values ...string) string {
	for _, v := range values {
		if v != "" {
			return v
		}
	}
	return ""
}
