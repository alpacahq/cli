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
}

func TestSaveAndLoadProfile(t *testing.T) {
	withTempDir(t)

	p := &Profile{
		APIKey:    "PK123",
		SecretKey: "SK456",
		BaseURL:   "https://paper-api.alpaca.markets",
		DataURL:   "https://data.alpaca.markets",
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
	if loaded.BaseURL != "https://paper-api.alpaca.markets" {
		t.Errorf("BaseURL = %q", loaded.BaseURL)
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

func TestLoad_EnvVarsOverrideProfile(t *testing.T) {
	withTempDir(t)

	_ = SaveProfile("test", &Profile{
		APIKey:    "profile-key",
		SecretKey: "profile-secret",
		BaseURL:   "https://paper-api.alpaca.markets",
	})
	_ = SaveGlobalConfig(&Config{DefaultProfile: "test"})

	t.Setenv("ALPACA_API_KEY", "env-key")
	t.Setenv("ALPACA_SECRET_KEY", "env-secret")

	r, err := Load("", "")
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if r.APIKey != "env-key" {
		t.Errorf("APIKey = %q, want env-key (env should override profile)", r.APIKey)
	}
	if r.SecretKey != "env-secret" {
		t.Errorf("SecretKey = %q, want env-secret", r.SecretKey)
	}
}

func TestLoad_ProfileFlagOverridesDefault(t *testing.T) {
	withTempDir(t)

	_ = SaveProfile("paper", &Profile{APIKey: "paper-key"})
	_ = SaveProfile("live", &Profile{APIKey: "live-key"})
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

	_ = SaveGlobalConfig(&Config{Output: "table"})

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

func TestLoad_EnvOverridesEverything(t *testing.T) {
	withTempDir(t)

	_ = SaveProfile("envtest", &Profile{
		APIKey:    "profile-key",
		SecretKey: "profile-secret",
		BaseURL:   "https://profile-url.example.com",
		DataURL:   "https://profile-data.example.com",
	})
	_ = SaveGlobalConfig(&Config{DefaultProfile: "envtest"})

	t.Setenv("ALPACA_API_KEY", "env-key")
	t.Setenv("ALPACA_SECRET_KEY", "env-secret")
	t.Setenv("ALPACA_BASE_URL", "https://env-url.example.com")
	t.Setenv("ALPACA_DATA_URL", "https://env-data.example.com")

	r, err := Load("", "")
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if r.APIKey != "env-key" {
		t.Errorf("APIKey = %q, want env-key", r.APIKey)
	}
	if r.SecretKey != "env-secret" {
		t.Errorf("SecretKey = %q, want env-secret", r.SecretKey)
	}
	if r.BaseURL != "https://env-url.example.com" {
		t.Errorf("BaseURL = %q, want env URL", r.BaseURL)
	}
	if r.DataURL != "https://env-data.example.com" {
		t.Errorf("DataURL = %q, want env URL", r.DataURL)
	}
}

func TestLoad_MissingConfigDir(t *testing.T) {
	t.Setenv("ALPACA_CONFIG_DIR", "/nonexistent/path/that/doesnt/exist")

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
