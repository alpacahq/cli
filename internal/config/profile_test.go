package config

import (
	"os"
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

func TestProfileSuppressWarnings(t *testing.T) {
	withTempDir(t)

	_ = SaveProfile("quiet", &Profile{
		APIKey:           "key",
		SecretKey:        "secret",
		SuppressWarnings: true,
	})
	_ = SaveGlobalConfig(&Config{DefaultProfile: "quiet"})

	r, err := Load("", "")
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if !r.SuppressWarnings {
		t.Error("expected SuppressWarnings = true")
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
	if perm != 0600 {
		t.Errorf("profile file permissions = %o, want 0600", perm)
	}
}
