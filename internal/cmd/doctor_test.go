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
			contains: []string{"env", "ALPACA_API_KEY"},
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

// TestDetectShadowedEnvVars ensures doctor flags the case where env API
// keys are active but a profile also has credentials - the user's profile
// login is effectively invisible in that state.
func TestDetectShadowedEnvVars(t *testing.T) {
	dir := t.TempDir()
	t.Setenv("ALPACA_CONFIG_DIR", dir)

	_ = config.SaveProfile("paper", &config.Profile{
		APIKey:    "profile-key",
		SecretKey: "profile-secret",
	})

	r := &config.Resolved{
		Source:      config.SourceEnvAPIKey,
		ProfileName: "paper",
	}
	got := detectShadowedEnvVars(r)
	if !strings.Contains(got, "paper") || !strings.Contains(got, "ALPACA_API_KEY") {
		t.Errorf("expected shadow warning mentioning profile name + env var, got: %q", got)
	}

	rProfile := &config.Resolved{
		Source:      config.SourceProfileAPIKey,
		ProfileName: "paper",
	}
	if got := detectShadowedEnvVars(rProfile); got != "" {
		t.Errorf("no shadow expected when source is profile, got: %q", got)
	}
}

func TestDetectShadowedEnvVars_NoProfile(t *testing.T) {
	dir := t.TempDir()
	t.Setenv("ALPACA_CONFIG_DIR", dir)

	r := &config.Resolved{
		Source:      config.SourceEnvAPIKey,
		ProfileName: "paper",
	}
	if got := detectShadowedEnvVars(r); got != "" {
		t.Errorf("no shadow expected when profile file is absent, got: %q", got)
	}
}
