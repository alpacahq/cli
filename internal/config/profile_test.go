package config

import (
	"os"
	"path/filepath"
	"testing"
)

func withTempDir(t *testing.T) {
	t.Helper()
	dir := t.TempDir()
	t.Setenv("ALPACA_CONFIG_DIR", dir)
	t.Setenv("ALPACA_API_KEY", "")
	t.Setenv("ALPACA_SECRET_KEY", "")
	t.Setenv("ALPACA_LIVE_TRADE", "")
	t.Setenv("ALPACA_PROFILE", "")
}

// setenvOrUnset calls t.Setenv(k, v) when v is non-empty; otherwise ensures
// the var is unset for the test. Needed because ALPACA_LIVE_TRADE has
// different semantics when unset vs set to empty string.
func setenvOrUnset(t *testing.T, k, v string) {
	t.Helper()
	if v == "" {
		_ = os.Unsetenv(k)
		return
	}
	t.Setenv(k, v)
}

func TestSaveAndLoadProfile(t *testing.T) {
	withTempDir(t)

	live := true
	p := &Profile{
		APIKey:    "PK123",
		SecretKey: "SK456",
		LiveTrade: &live,
	}

	if err := SaveProfile("test", p); err != nil {
		t.Fatalf("SaveProfile: %v", err)
	}

	loaded := LoadProfileByName("test")
	if loaded.APIKey != "PK123" {
		t.Errorf("APIKey = %q, want PK123", loaded.APIKey)
	}
	if loaded.SecretKey != "SK456" {
		t.Errorf("SecretKey = %q, want SK456", loaded.SecretKey)
	}
	if loaded.LiveTrade == nil || !*loaded.LiveTrade {
		t.Errorf("LiveTrade = %v, want &true", loaded.LiveTrade)
	}
}

func TestSaveProfileCreatesDirectory(t *testing.T) {
	withTempDir(t)

	p := &Profile{APIKey: "key", SecretKey: "secret"}
	if err := SaveProfile("newprofile", p); err != nil {
		t.Fatalf("SaveProfile: %v", err)
	}

	loaded := LoadProfileByName("newprofile")
	if loaded.APIKey != "key" {
		t.Errorf("APIKey = %q, want key", loaded.APIKey)
	}
}

func TestDeleteProfile(t *testing.T) {
	withTempDir(t)

	p := &Profile{APIKey: "key", SecretKey: "secret"}
	_ = SaveProfile("deleteme", p)

	if err := DeleteProfile("deleteme"); err != nil {
		t.Fatalf("DeleteProfile: %v", err)
	}

	loaded := LoadProfileByName("deleteme")
	if loaded.APIKey != "" {
		t.Error("deleted profile should return empty")
	}
}

func TestDeleteProfile_NotFound(t *testing.T) {
	withTempDir(t)

	err := DeleteProfile("nonexistent")
	if err == nil {
		t.Error("expected error deleting nonexistent profile")
	}
	if !os.IsNotExist(err) {
		t.Errorf("expected IsNotExist, got: %v", err)
	}
}

func TestListProfiles(t *testing.T) {
	withTempDir(t)

	for _, name := range []string{"paper", "live", "staging"} {
		_ = SaveProfile(name, &Profile{APIKey: name})
	}

	names, err := ListProfiles()
	if err != nil {
		t.Fatalf("ListProfiles: %v", err)
	}
	if len(names) != 3 {
		t.Fatalf("expected 3 profiles, got %d", len(names))
	}

	want := map[string]bool{"paper": true, "live": true, "staging": true}
	for _, n := range names {
		if !want[n] {
			t.Errorf("unexpected profile: %q", n)
		}
	}
}

func TestListProfiles_EmptyDir(t *testing.T) {
	withTempDir(t)

	names, err := ListProfiles()
	if err != nil {
		t.Fatalf("ListProfiles: %v", err)
	}
	if len(names) != 0 {
		t.Errorf("expected 0 profiles, got %d", len(names))
	}
}

func TestSaveAndLoadGlobalConfig(t *testing.T) {
	withTempDir(t)

	cfg := &Config{
		DefaultProfile: "live",
		Output:         "json",
	}

	if err := SaveGlobalConfig(cfg); err != nil {
		t.Fatalf("SaveGlobalConfig: %v", err)
	}

	loaded := loadGlobalConfig()
	if loaded.DefaultProfile != "live" {
		t.Errorf("DefaultProfile = %q, want live", loaded.DefaultProfile)
	}
	if loaded.Output != "json" {
		t.Errorf("Output = %q, want json", loaded.Output)
	}
}

// TestLoad_EnvAPIKeysBeatProfileOAuth guards the core behavior change:
// env API keys now win over a profile's OAuth token. Before atomic bundles,
// the profile OAuth token silently took precedence, which was the footgun
// we set out to fix.
func TestLoad_EnvAPIKeysBeatProfileOAuth(t *testing.T) {
	withTempDir(t)

	_ = SaveProfile("test", &Profile{
		AccessToken: "oauth-from-profile",
	})
	_ = SaveGlobalConfig(&Config{DefaultProfile: "test"})

	t.Setenv("ALPACA_API_KEY", "env-key")
	t.Setenv("ALPACA_SECRET_KEY", "env-secret")

	r, err := Load("", "")
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if r.Source != SourceEnvAPIKey {
		t.Errorf("Source = %q, want %q", r.Source, SourceEnvAPIKey)
	}
	if r.APIKey != "env-key" || r.SecretKey != "env-secret" {
		t.Errorf("got APIKey=%q Secret=%q", r.APIKey, r.SecretKey)
	}
	if r.AccessToken != "" {
		t.Errorf("AccessToken should be empty when env API keys win, got %q", r.AccessToken)
	}
}

// TestLoad_PartialEnvFallsThroughToProfile verifies that a half-bundle
// (just ALPACA_API_KEY without its secret) is treated as absent, not merged.
// This prevents silent field-level mixing across sources.
func TestLoad_PartialEnvFallsThroughToProfile(t *testing.T) {
	withTempDir(t)

	_ = SaveProfile("test", &Profile{
		APIKey:    "profile-key",
		SecretKey: "profile-secret",
	})
	_ = SaveGlobalConfig(&Config{DefaultProfile: "test"})

	t.Setenv("ALPACA_API_KEY", "env-key")

	r, err := Load("", "")
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if r.Source != SourceProfileAPIKey {
		t.Errorf("Source = %q, want %q", r.Source, SourceProfileAPIKey)
	}
	if r.APIKey != "profile-key" {
		t.Errorf("APIKey = %q, want profile-key (no field mixing)", r.APIKey)
	}
}

// TestLoad_EnvCredsDefaultToPaper verifies the safe default: env-sourced
// credentials without an explicit URL go to paper, regardless of what the
// default profile's live_trade flag says (we ignore profile URL state
// entirely when env creds win).
func TestLoad_EnvCredsDefaultToPaper(t *testing.T) {
	withTempDir(t)

	live := true
	_ = SaveProfile("test", &Profile{
		APIKey:    "profile-key",
		SecretKey: "profile-secret",
		LiveTrade: &live,
	})
	_ = SaveGlobalConfig(&Config{DefaultProfile: "test"})

	t.Setenv("ALPACA_API_KEY", "env-key")
	t.Setenv("ALPACA_SECRET_KEY", "env-secret")

	r, err := Load("", "")
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if r.BaseURL != "https://paper-api.alpaca.markets" {
		t.Errorf("BaseURL = %q, want paper default", r.BaseURL)
	}
}

func TestLoad_ProfileLiveGoesLive(t *testing.T) {
	withTempDir(t)

	live := true
	_ = SaveProfile("test", &Profile{
		APIKey:    "profile-key",
		SecretKey: "profile-secret",
		LiveTrade: &live,
	})
	_ = SaveGlobalConfig(&Config{DefaultProfile: "test"})

	r, err := Load("", "")
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if r.BaseURL != "https://api.alpaca.markets" {
		t.Errorf("BaseURL = %q, want live URL", r.BaseURL)
	}
	if r.Source != SourceProfileAPIKey {
		t.Errorf("Source = %q, want %q", r.Source, SourceProfileAPIKey)
	}
}

func TestLoad_ProfilePaperGoesPaper(t *testing.T) {
	withTempDir(t)

	paper := false
	_ = SaveProfile("test", &Profile{
		APIKey:    "profile-key",
		SecretKey: "profile-secret",
		LiveTrade: &paper,
	})
	_ = SaveGlobalConfig(&Config{DefaultProfile: "test"})

	r, err := Load("", "")
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if r.BaseURL != "https://paper-api.alpaca.markets" {
		t.Errorf("BaseURL = %q, want paper URL", r.BaseURL)
	}
}

// TestLoad_ProfileWithoutLiveTradeDefaultsToPaper covers profiles that omit
// the live_trade field (the common case for paper profiles). Missing is a
// safe default - paper, never live.
func TestLoad_ProfileWithoutLiveTradeDefaultsToPaper(t *testing.T) {
	withTempDir(t)

	_ = SaveProfile("test", &Profile{
		APIKey:    "profile-key",
		SecretKey: "profile-secret",
	})
	_ = SaveGlobalConfig(&Config{DefaultProfile: "test"})

	r, err := Load("", "")
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if r.BaseURL != "https://paper-api.alpaca.markets" {
		t.Errorf("BaseURL = %q, want paper default", r.BaseURL)
	}
}

// TestLoad_LiveTradeEnvOverridesProfile verifies ALPACA_LIVE_TRADE beats
// whatever the profile says - the env var is the top rung of URL resolution.
func TestLoad_LiveTradeEnvOverridesProfile(t *testing.T) {
	withTempDir(t)

	live := true
	_ = SaveProfile("test", &Profile{
		APIKey:    "profile-key",
		SecretKey: "profile-secret",
		LiveTrade: &live,
	})
	_ = SaveGlobalConfig(&Config{DefaultProfile: "test"})

	t.Setenv("ALPACA_LIVE_TRADE", "false")

	r, err := Load("", "")
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if r.BaseURL != "https://paper-api.alpaca.markets" {
		t.Errorf("BaseURL = %q, want paper (env overrides live profile)", r.BaseURL)
	}
}

// TestLoad_LiveTradeEnvValues exercises the boolean env var. Only
// case-insensitive "true" means live; anything else - "False", "no", "0",
// "anything" - means paper. Typos fall through to the safe default.
func TestLoad_LiveTradeEnvValues(t *testing.T) {
	for _, tc := range []struct {
		value string
		want  string
	}{
		{"true", "https://api.alpaca.markets"},
		{"True", "https://api.alpaca.markets"},
		{"TRUE", "https://api.alpaca.markets"},
		{"false", "https://paper-api.alpaca.markets"},
		{"False", "https://paper-api.alpaca.markets"},
		{"0", "https://paper-api.alpaca.markets"},
		{"no", "https://paper-api.alpaca.markets"},
		{"yes", "https://paper-api.alpaca.markets"},
		{"1", "https://paper-api.alpaca.markets"},
	} {
		t.Run(tc.value, func(t *testing.T) {
			withTempDir(t)
			t.Setenv("ALPACA_API_KEY", "env-key")
			t.Setenv("ALPACA_SECRET_KEY", "env-secret")
			t.Setenv("ALPACA_LIVE_TRADE", tc.value)

			r, err := Load("", "")
			if err != nil {
				t.Fatalf("Load: %v", err)
			}
			if r.BaseURL != tc.want {
				t.Errorf("ALPACA_LIVE_TRADE=%q -> BaseURL=%q, want %q", tc.value, r.BaseURL, tc.want)
			}
		})
	}
}

func TestLoad_LiveTradeUnsetKeepsPaperDefault(t *testing.T) {
	withTempDir(t)
	setenvOrUnset(t, "ALPACA_LIVE_TRADE", "")
	t.Setenv("ALPACA_API_KEY", "env-key")
	t.Setenv("ALPACA_SECRET_KEY", "env-secret")

	r, err := Load("", "")
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if r.BaseURL != "https://paper-api.alpaca.markets" {
		t.Errorf("unset ALPACA_LIVE_TRADE -> BaseURL=%q, want paper", r.BaseURL)
	}
}

// TestLoad_AccessTokenEnvIgnored guards against the removed ALPACA_ACCESS_TOKEN
// env var being resurrected. OAuth tokens only come from profile files.
func TestLoad_AccessTokenEnvIgnored(t *testing.T) {
	withTempDir(t)
	t.Setenv("ALPACA_ACCESS_TOKEN", "should-be-ignored")
	_ = SaveGlobalConfig(&Config{DefaultProfile: "paper"})

	r, err := Load("", "")
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if r.AccessToken != "" {
		t.Errorf("ALPACA_ACCESS_TOKEN should not be read, got %q", r.AccessToken)
	}
	if r.Source != SourceNone {
		t.Errorf("Source = %q, want %q (env access token must not populate credentials)", r.Source, SourceNone)
	}
}

func TestLoad_NoCredentialsReturnsSourceNone(t *testing.T) {
	withTempDir(t)

	r, err := Load("", "")
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if r.Source != SourceNone {
		t.Errorf("Source = %q, want %q", r.Source, SourceNone)
	}
	if r.HasCredentials() {
		t.Error("HasCredentials should be false with no creds configured")
	}
	if err := r.Validate(); err == nil {
		t.Error("Validate should error with no credentials")
	}
}

func TestLoad_ProfileOAuthSource(t *testing.T) {
	withTempDir(t)
	_ = SaveProfile("test", &Profile{
		AccessToken: "oauth-token",
		Scopes:      "trading data",
	})
	_ = SaveGlobalConfig(&Config{DefaultProfile: "test"})

	r, err := Load("", "")
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if r.Source != SourceProfileOAuth {
		t.Errorf("Source = %q, want %q", r.Source, SourceProfileOAuth)
	}
	if !r.IsOAuth() {
		t.Error("IsOAuth should be true")
	}
	if r.Scopes != "trading data" {
		t.Errorf("Scopes = %q", r.Scopes)
	}
}

func TestLoad_ProfileFlagOverridesDefault(t *testing.T) {
	withTempDir(t)

	_ = SaveProfile("paper", &Profile{APIKey: "paper-key", SecretKey: "paper-secret"})
	_ = SaveProfile("live", &Profile{APIKey: "live-key", SecretKey: "live-secret"})
	_ = SaveGlobalConfig(&Config{DefaultProfile: "paper"})

	r, err := Load("live", "")
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if r.ProfileName != "live" {
		t.Errorf("ProfileName = %q, want live", r.ProfileName)
	}
	if r.APIKey != "live-key" {
		t.Errorf("APIKey = %q, want live-key", r.APIKey)
	}
}

func TestLoad_OutputFlagOverridesConfig(t *testing.T) {
	withTempDir(t)

	_ = SaveGlobalConfig(&Config{Output: "json"})

	r, err := Load("", "csv")
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if r.Output != "csv" {
		t.Errorf("Output = %q, want csv", r.Output)
	}
}

func TestProfileFilePermissions(t *testing.T) {
	withTempDir(t)

	_ = SaveProfile("permtest", &Profile{APIKey: "key", SecretKey: "secret"})

	info, err := os.Stat(Dir() + "/profiles/permtest.yaml")
	if err != nil {
		t.Fatalf("stat: %v", err)
	}

	perm := info.Mode().Perm()
	if perm != 0o600 {
		t.Errorf("profile file permissions = %o, want 0600", perm)
	}
}

func TestLoadProfile_CorruptedYAML(t *testing.T) {
	withTempDir(t)

	dir := filepath.Join(Dir(), "profiles")
	_ = os.MkdirAll(dir, 0o700)
	_ = os.WriteFile(filepath.Join(dir, "corrupt.yaml"), []byte("{{{{not yaml at all!!!!"), 0o600)

	p := LoadProfileByName("corrupt")
	if p.APIKey != "" {
		t.Errorf("corrupted profile should return empty, got APIKey=%q", p.APIKey)
	}
}

func TestLoad_MissingConfigDir(t *testing.T) {
	t.Setenv("ALPACA_CONFIG_DIR", "/nonexistent/path/that/doesnt/exist")
	t.Setenv("ALPACA_API_KEY", "")
	t.Setenv("ALPACA_SECRET_KEY", "")
	setenvOrUnset(t, "ALPACA_LIVE_TRADE", "")

	r, err := Load("", "")
	if err != nil {
		t.Fatalf("Load should not error on missing config dir: %v", err)
	}
	if r.ProfileName != "paper" {
		t.Errorf("ProfileName = %q, want paper (default)", r.ProfileName)
	}
}

func TestLoadGlobalConfig_CorruptedYAML(t *testing.T) {
	withTempDir(t)

	dir := Dir()
	_ = os.MkdirAll(dir, 0o700)
	_ = os.WriteFile(filepath.Join(dir, "config.yaml"), []byte(":::not valid yaml:::"), 0o600)

	cfg := loadGlobalConfig()
	if cfg.DefaultProfile != "" {
		t.Errorf("corrupted config should return empty, got DefaultProfile=%q", cfg.DefaultProfile)
	}
}
