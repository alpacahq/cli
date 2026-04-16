package cmd

import (
	"strings"
	"testing"

	"github.com/alpacahq/cli/internal/config"
)

func TestCredentialSourceDescription(t *testing.T) {
	cases := []struct {
		name     string
		resolved *config.Resolved
		contains []string
	}{
		{
			name:     "env api key",
			resolved: &config.Resolved{Source: config.SourceEnvAPIKey},
			contains: []string{"env", "ALPACA_API_KEY", "ALPACA_SECRET_KEY"},
		},
		{
			name:     "profile oauth",
			resolved: &config.Resolved{Source: config.SourceProfileOAuth, ProfileName: "paper"},
			contains: []string{"OAuth", `"paper"`},
		},
		{
			name:     "profile api key",
			resolved: &config.Resolved{Source: config.SourceProfileAPIKey, ProfileName: "live"},
			contains: []string{"API key", `"live"`},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := credentialSourceDescription(tc.resolved)
			for _, want := range tc.contains {
				if !strings.Contains(got, want) {
					t.Errorf("%q missing %q", got, want)
				}
			}
		})
	}
}

// TestEnvShadowsProfile ensures we detect when env API keys are active and a
// profile also has credentials - the user's profile login is effectively
// invisible in that state.
func TestEnvShadowsProfile(t *testing.T) {
	dir := t.TempDir()
	t.Setenv("ALPACA_CONFIG_DIR", dir)
	t.Setenv("ALPACA_API_KEY", "env-key")
	t.Setenv("ALPACA_SECRET_KEY", "env-secret")

	_ = config.SaveProfile("paper", &config.Profile{
		APIKey:    "profile-key",
		SecretKey: "profile-secret",
	})

	if !envShadowsProfile("paper") {
		t.Error("expected shadowing when env keys set and profile has creds")
	}
}

func TestEnvShadowsProfile_NoEnv(t *testing.T) {
	dir := t.TempDir()
	t.Setenv("ALPACA_CONFIG_DIR", dir)
	t.Setenv("ALPACA_API_KEY", "")
	t.Setenv("ALPACA_SECRET_KEY", "")

	_ = config.SaveProfile("paper", &config.Profile{
		APIKey:    "profile-key",
		SecretKey: "profile-secret",
	})

	if envShadowsProfile("paper") {
		t.Error("no shadow expected when env vars are unset")
	}
}

func TestEnvShadowsProfile_NoProfile(t *testing.T) {
	dir := t.TempDir()
	t.Setenv("ALPACA_CONFIG_DIR", dir)
	t.Setenv("ALPACA_API_KEY", "env-key")
	t.Setenv("ALPACA_SECRET_KEY", "env-secret")

	if envShadowsProfile("paper") {
		t.Error("no shadow expected when profile file is absent")
	}
}
