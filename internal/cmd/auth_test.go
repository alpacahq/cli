package cmd

import (
	"bytes"
	"strings"
	"testing"

	"github.com/alpacahq/cli/internal/config"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// captureColorOutput redirects color.Output (where color.Yellow etc. write)
// to a buffer and disables color codes for the duration of the test. Returns
// the buffer plus a cleanup func.
func captureColorOutput(t *testing.T) (*bytes.Buffer, func()) {
	t.Helper()
	buf := new(bytes.Buffer)
	oldOut := color.Output
	oldNoColor := color.NoColor
	color.Output = buf
	color.NoColor = true
	return buf, func() {
		color.Output = oldOut
		color.NoColor = oldNoColor
	}
}

// runLoginAPIKey invokes loginWithAPIKey with the given flags set on a
// throwaway cobra.Command. We don't use Root().Execute() here because cobra
// auto-registers completion/help subcommands on first Execute(), polluting
// the shared rootCmd tree that the golden tests rely on.
func runLoginAPIKey(t *testing.T, name, key, secret string, live bool) error {
	t.Helper()
	cmd := &cobra.Command{Use: "login"}
	cmd.Flags().String("key", key, "")
	cmd.Flags().String("secret", secret, "")
	cmd.Flags().String("name", name, "")
	cmd.Flags().Bool("no-validate", true, "")
	cmd.Flags().Bool("live", live, "")
	return loginWithAPIKey(cmd)
}

// TestLogin_WarnsWhenEnvShadows verifies that logging in while
// ALPACA_API_KEY/ALPACA_SECRET_KEY are set triggers the shadowing warning -
// users must know their fresh credentials are about to be overridden by env.
func TestLogin_WarnsWhenEnvShadows(t *testing.T) {
	dir := t.TempDir()
	t.Setenv("ALPACA_CONFIG_DIR", dir)
	t.Setenv("ALPACA_API_KEY", "env-key")
	t.Setenv("ALPACA_SECRET_KEY", "env-secret")
	t.Setenv("ALPACA_PAPER_TRADE", "")
	t.Setenv("ALPACA_PROFILE", "")

	buf, cleanup := captureColorOutput(t)
	defer cleanup()

	if err := runLoginAPIKey(t, "paper", "new-key", "new-secret", false); err != nil {
		t.Fatalf("login failed: %v", err)
	}

	out := buf.String()
	if !strings.Contains(out, "ALPACA_API_KEY is set in your environment") {
		t.Errorf("expected shadow warning in output, got: %q", out)
	}
	if !strings.Contains(out, `"paper"`) {
		t.Errorf("expected warning to name the shadowed profile, got: %q", out)
	}

	p := config.LoadProfileByName("paper")
	if p.APIKey != "new-key" || p.SecretKey != "new-secret" {
		t.Errorf("expected login to persist new credentials, got: %+v", p)
	}
}

// TestLogin_NoWarnWhenEnvUnset verifies the warning stays quiet when the
// env is clean - nothing is shadowing the freshly-saved profile.
func TestLogin_NoWarnWhenEnvUnset(t *testing.T) {
	dir := t.TempDir()
	t.Setenv("ALPACA_CONFIG_DIR", dir)
	t.Setenv("ALPACA_API_KEY", "")
	t.Setenv("ALPACA_SECRET_KEY", "")
	t.Setenv("ALPACA_PAPER_TRADE", "")
	t.Setenv("ALPACA_PROFILE", "")

	buf, cleanup := captureColorOutput(t)
	defer cleanup()

	if err := runLoginAPIKey(t, "paper", "new-key", "new-secret", false); err != nil {
		t.Fatalf("login failed: %v", err)
	}

	out := buf.String()
	if strings.Contains(out, "ALPACA_API_KEY is set in your environment") {
		t.Errorf("unexpected shadow warning, got: %q", out)
	}
}

// TestLogin_WarnTargetsNewProfile verifies the warning names the profile
// the user just logged into, not some pre-existing default. Pre-existing
// profiles under other names are irrelevant to the shadowing check.
func TestLogin_WarnTargetsNewProfile(t *testing.T) {
	dir := t.TempDir()
	t.Setenv("ALPACA_CONFIG_DIR", dir)
	t.Setenv("ALPACA_API_KEY", "env-key")
	t.Setenv("ALPACA_SECRET_KEY", "env-secret")
	t.Setenv("ALPACA_PAPER_TRADE", "")
	t.Setenv("ALPACA_PROFILE", "")

	paper := true
	if err := config.SaveProfile("paper", &config.Profile{
		APIKey: "pre-existing-key", SecretKey: "pre-existing-secret", PaperTrade: &paper,
	}); err != nil {
		t.Fatalf("SaveProfile: %v", err)
	}

	buf, cleanup := captureColorOutput(t)
	defer cleanup()

	if err := runLoginAPIKey(t, "prod", "new-key", "new-secret", false); err != nil {
		t.Fatalf("login failed: %v", err)
	}

	out := buf.String()
	if !strings.Contains(out, `"prod"`) {
		t.Errorf("expected warning to name the just-logged-in profile %q, got: %q", "prod", out)
	}
	if strings.Contains(out, `"paper"`) {
		t.Errorf("warning should not name the pre-existing profile, got: %q", out)
	}
}
