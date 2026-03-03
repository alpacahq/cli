package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

type Config struct {
	DefaultProfile string `yaml:"default_profile"`
	Output         string `yaml:"output"`
	Color          string `yaml:"color"`
}

type Profile struct {
	APIKey           string `yaml:"api_key"`
	SecretKey        string `yaml:"secret_key"`
	BaseURL          string `yaml:"base_url"`
	DataURL          string `yaml:"data_url"`
	SuppressWarnings bool   `yaml:"suppress_warnings,omitempty"`

	// Deprecated: kept for backwards compat with existing profile files.
	// New profiles store base_url directly.
	Environment string `yaml:"environment,omitempty"`
}

type Resolved struct {
	APIKey           string
	SecretKey        string
	BaseURL          string
	DataURL          string
	Output           string
	Color            string
	ProfileName      string
	SuppressWarnings bool
}

func Dir() string {
	if d := os.Getenv("ALPACA_CONFIG_DIR"); d != "" {
		return d
	}
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".config", "alpaca")
}

func Load(profileFlag, outputFlag string) (*Resolved, error) {
	cfg := loadGlobalConfig()
	profileName := resolve(profileFlag, os.Getenv("ALPACA_PROFILE"), cfg.DefaultProfile, "paper")
	profile := loadProfile(profileName)

	r := &Resolved{
		ProfileName:      profileName,
		APIKey:           resolve(os.Getenv("ALPACA_API_KEY"), profile.APIKey),
		SecretKey:        resolve(os.Getenv("ALPACA_SECRET_KEY"), profile.SecretKey),
		BaseURL:          resolve(os.Getenv("ALPACA_BASE_URL"), profile.BaseURL),
		DataURL:          resolve(os.Getenv("ALPACA_DATA_URL"), profile.DataURL),
		Output:           resolve(outputFlag, os.Getenv("ALPACA_OUTPUT"), cfg.Output, "table"),
		Color:            resolve(cfg.Color, "auto"),
		SuppressWarnings: profile.SuppressWarnings,
	}

	// Backwards compat: old profiles may have environment instead of base_url
	if r.BaseURL == "" {
		env := resolve(os.Getenv("ALPACA_ENVIRONMENT"), profile.Environment)
		r.BaseURL = ResolveBaseURL(env)
	}
	if r.DataURL == "" {
		r.DataURL = "https://data.alpaca.markets"
	}

	return r, nil
}

func (r *Resolved) HasCredentials() bool {
	return r.APIKey != "" && r.SecretKey != ""
}

func (r *Resolved) Validate() error {
	if !r.HasCredentials() {
		return fmt.Errorf("authentication required\nHint: run `alpaca profile login` to set up your credentials")
	}
	return nil
}

// ResolveBaseURL takes a value that is either a well-known alias
// ("paper", "live") or a full URL, and returns a URL.
// Empty string defaults to the paper trading URL.
func ResolveBaseURL(value string) string {
	switch strings.ToLower(strings.TrimSpace(value)) {
	case "live":
		return "https://api.alpaca.markets"
	case "", "paper":
		return "https://paper-api.alpaca.markets"
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
