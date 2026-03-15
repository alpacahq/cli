package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/alpacahq/cli/internal/config"
	"github.com/fatih/color"
)

const (
	updateCheckTTL   = 24 * time.Hour
	installHomebrew  = "homebrew"
	installGoInstall = "goinstall"
)

type updateState struct {
	LatestVersion string    `json:"latest_version"`
	CheckedAt     time.Time `json:"checked_at"`
	InstallMethod string    `json:"install_method"`
}

// suppressUpdateNotice is set by commands that handle update info themselves
// (update, version) to avoid a redundant stderr notice.
var suppressUpdateNotice bool

func detectInstallMethod() string {
	exe, err := os.Executable()
	if err != nil {
		return installGoInstall
	}
	resolved, err := filepath.EvalSymlinks(exe)
	if err != nil {
		resolved = exe
	}

	if strings.Contains(resolved, "/Cellar/") || strings.Contains(resolved, "/homebrew/") {
		return installHomebrew
	}

	gobin := os.Getenv("GOBIN")
	if gobin != "" && strings.HasPrefix(resolved, gobin) {
		return installGoInstall
	}

	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		home, _ := os.UserHomeDir()
		gopath = filepath.Join(home, "go")
	}
	if strings.HasPrefix(resolved, filepath.Join(gopath, "bin")) {
		return installGoInstall
	}

	return installGoInstall
}

func upgradeCommand(method string) string {
	switch method {
	case installHomebrew:
		return "brew upgrade alpaca"
	default:
		return "go install github.com/alpacahq/cli/cmd/alpaca@latest"
	}
}

func updateStatePath() string {
	return filepath.Join(config.Dir(), "update-state.json")
}

func loadUpdateState() *updateState {
	data, err := os.ReadFile(updateStatePath())
	if err != nil {
		return nil
	}
	var s updateState
	if err := json.Unmarshal(data, &s); err != nil {
		return nil
	}
	return &s
}

func saveUpdateState(s *updateState) {
	dir := config.Dir()
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return
	}
	data, err := json.Marshal(s)
	if err != nil {
		return
	}
	_ = os.WriteFile(updateStatePath(), data, 0o600)
}

func versionsEqual(a, b string) bool {
	return strings.TrimPrefix(a, "v") == strings.TrimPrefix(b, "v")
}

// checkForUpdateAsync spawns a background goroutine to check for updates.
// Returns a channel that receives an *updateState if an update is available,
// or is closed with no value if no update is needed.
func checkForUpdateAsync() <-chan *updateState {
	ch := make(chan *updateState, 1)

	if version == "dev" || quietFlag || os.Getenv("ALPACA_NO_UPDATE_NOTIFY") != "" {
		close(ch)
		return ch
	}

	go func() {
		defer close(ch)

		state := loadUpdateState()
		if state != nil && time.Since(state.CheckedAt) < updateCheckTTL {
			if state.LatestVersion != "" && !versionsEqual(version, state.LatestVersion) {
				ch <- state
			}
			return
		}

		latest, err := getLatestVersion()
		if err != nil {
			return
		}

		newState := &updateState{
			LatestVersion: latest,
			CheckedAt:     time.Now(),
			InstallMethod: detectInstallMethod(),
		}
		saveUpdateState(newState)

		// Don't notify on the very first check (user just installed).
		if state == nil {
			return
		}

		if !versionsEqual(version, latest) {
			ch <- newState
		}
	}()

	return ch
}

func showUpdateNotice(ch <-chan *updateState) {
	if suppressUpdateNotice {
		return
	}

	var state *updateState
	select {
	case s, ok := <-ch:
		if !ok {
			return
		}
		state = s
	case <-time.After(200 * time.Millisecond):
		return
	}

	if state == nil {
		return
	}

	method := state.InstallMethod
	if method == "" {
		method = detectInstallMethod()
	}

	fmt.Fprintln(os.Stderr)
	color.New(color.FgYellow).Fprintf(os.Stderr,
		"A new version of alpaca is available: %s → %s\n", version, state.LatestVersion)
	fmt.Fprintf(os.Stderr, "Run `%s` to upgrade.\n", upgradeCommand(method))
}
