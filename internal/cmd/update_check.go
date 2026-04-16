package cmd

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	installHomebrew  = "homebrew"
	installGoInstall = "goinstall"
)

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
		return "brew upgrade alpacahq/tap/cli"
	default:
		return "go install github.com/alpacahq/cli/cmd/alpaca@latest"
	}
}

// versionNewer reports whether latest is strictly greater than current
// using numeric major.minor.patch comparison.
func versionNewer(latest, current string) bool {
	parseVer := func(s string) [3]int {
		s = strings.TrimPrefix(s, "v")
		if idx := strings.IndexByte(s, '-'); idx != -1 {
			s = s[:idx]
		}
		parts := strings.SplitN(s, ".", 3)
		var v [3]int
		for i := 0; i < len(parts) && i < 3; i++ {
			n, _ := strconv.Atoi(parts[i])
			v[i] = n
		}
		return v
	}
	l, c := parseVer(latest), parseVer(current)
	for i := range 3 {
		if l[i] != c[i] {
			return l[i] > c[i]
		}
	}
	return false
}
