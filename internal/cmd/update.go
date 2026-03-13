package cmd

import (
	"archive/tar"
	"archive/zip"
	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/alpacahq/cli/internal/cmdutil"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

const (
	repoOwner = "alpacahq"
	repoName  = "cli"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update alpaca CLI to the latest version",
	Long: `Update alpaca CLI to the latest version.

Checks GitHub for the latest release and installs it. The install method
is auto-detected (Homebrew, go install, or binary download) and the
appropriate upgrade command is shown.

Use --check to see if an update is available without installing.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		suppressUpdateNotice = true
		checkOnly := cmdutil.Bool(cmd, "check")

		latest, downloadURL, checksumURL, err := getLatestRelease()
		if err != nil {
			return fmt.Errorf("checking for updates: %w", err)
		}

		method := detectInstallMethod()
		saveUpdateState(&updateState{
			LatestVersion: latest,
			CheckedAt:     time.Now(),
			InstallMethod: method,
		})

		current := version
		upToDate := versionsEqual(current, latest)

		if checkOnly {
			m := map[string]any{
				"current":          strings.TrimPrefix(current, "v"),
				"latest":           strings.TrimPrefix(latest, "v"),
				"update_available": !upToDate,
				"install_method":   method,
				"update_command":   upgradeCommand(method),
			}
			enc := json.NewEncoder(os.Stdout)
			enc.SetIndent("", "  ")
			return enc.Encode(m)
		}

		if upToDate {
			color.Green("Already up to date (%s)", current)
			return nil
		}

		fmt.Printf("Current version: %s\n", current)
		fmt.Printf("Latest version:  %s\n", latest)

		if method == installHomebrew {
			fmt.Println()
			fmt.Println("This CLI was installed via Homebrew. To update, run:")
			fmt.Println("  brew upgrade alpaca")
			return nil
		}
		if method == installGoInstall {
			fmt.Println()
			fmt.Println("This CLI was installed via go install. To update, run:")
			fmt.Println("  go install github.com/alpacahq/cli/cmd/alpaca@latest")
			return nil
		}

		if downloadURL == "" {
			return fmt.Errorf("no prebuilt binary available for %s/%s\nHint: install via Go instead:\n  go install github.com/alpacahq/cli/cmd/alpaca@%s", runtime.GOOS, runtime.GOARCH, latest)
		}

		fmt.Printf("Downloading %s...\n", latest)

		archive, err := downloadBinary(downloadURL)
		if err != nil {
			return fmt.Errorf("downloading update: %w", err)
		}

		if checksumURL != "" {
			if err := verifyChecksum(archive, checksumURL, downloadURL); err != nil {
				return fmt.Errorf("checksum verification failed: %w", err)
			}
		}

		binaryName := "alpaca"
		if runtime.GOOS == osWindows {
			binaryName = "alpaca.exe"
		}

		binary, err := extractBinaryFromArchive(archive, downloadURL, binaryName)
		if err != nil {
			return fmt.Errorf("extracting binary: %w", err)
		}

		execPath, err := os.Executable()
		if err != nil {
			return fmt.Errorf("locating current binary: %w", err)
		}

		if err := os.WriteFile(execPath, binary, 0o755); err != nil {
			return fmt.Errorf("replacing binary at %s: %w\nHint: try one of:\n  sudo alpaca update\n  go install github.com/alpacahq/cli/cmd/alpaca@latest\n  Download from https://github.com/%s/%s/releases", execPath, err, repoOwner, repoName)
		}

		color.Green("Updated to %s", latest)

		fmt.Println()
		if err := installCompletions(detectShell()); err != nil {
			verboseLog("completions install: %v", err)
		}
		if err := installManPages(); err != nil {
			verboseLog("man pages install: %v", err)
		}

		return nil
	},
}

func init() {
	updateCmd.Flags().Bool("check", false, "Check for updates without installing")
}

type ghRelease struct {
	TagName string    `json:"tag_name"`
	Assets  []ghAsset `json:"assets"`
}

type ghAsset struct {
	Name               string `json:"name"`
	BrowserDownloadURL string `json:"browser_download_url"`
}

func getLatestRelease() (tag, downloadURL, checksumURL string, err error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/releases/latest", repoOwner, repoName)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", "", "", err
	}
	req.Header.Set("User-Agent", "alpaca-cli/"+version)
	resp, err := (&http.Client{Timeout: 10 * time.Second}).Do(req)
	if err != nil {
		return "", "", "", err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return "", "", "", fmt.Errorf("GitHub API returned %d", resp.StatusCode)
	}

	var release ghRelease
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		return "", "", "", err
	}

	archMap := map[string]string{
		"amd64": "x86_64",
		"arm64": "arm64",
	}
	osMap := map[string]string{
		"darwin":  "Darwin",
		"linux":   "Linux",
		"windows": "Windows",
	}

	goosName := osMap[runtime.GOOS]
	goarchName := archMap[runtime.GOARCH]

	for _, a := range release.Assets {
		if a.Name == "checksums.txt" {
			checksumURL = a.BrowserDownloadURL
		}
		if strings.Contains(a.Name, goosName) && strings.Contains(a.Name, goarchName) {
			downloadURL = a.BrowserDownloadURL
		}
	}

	return release.TagName, downloadURL, checksumURL, nil
}

func downloadBinary(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "alpaca-cli/"+version)
	resp, err := (&http.Client{Timeout: 60 * time.Second}).Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("download returned HTTP %d", resp.StatusCode)
	}

	return io.ReadAll(resp.Body)
}

func verifyChecksum(binary []byte, checksumURL, binaryURL string) error {
	checksumData, err := downloadBinary(checksumURL)
	if err != nil {
		return fmt.Errorf("downloading checksums: %w", err)
	}

	hash := sha256.Sum256(binary)
	got := hex.EncodeToString(hash[:])

	binaryName := binaryURL[strings.LastIndex(binaryURL, "/")+1:]
	for _, line := range strings.Split(string(checksumData), "\n") {
		parts := strings.Fields(line)
		if len(parts) == 2 && parts[1] == binaryName {
			if parts[0] != got {
				return fmt.Errorf("expected %s, got %s", parts[0], got)
			}
			return nil
		}
	}

	return fmt.Errorf("no checksum found for %s", binaryName)
}

func extractBinaryFromArchive(data []byte, archiveURL, binaryName string) ([]byte, error) {
	if strings.HasSuffix(archiveURL, ".zip") {
		return extractFromZip(data, binaryName)
	}
	return extractFromTarGz(data, binaryName)
}

func extractFromTarGz(data []byte, binaryName string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("decompressing: %w", err)
	}
	defer func() { _ = gz.Close() }()

	tr := tar.NewReader(gz)
	for {
		header, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("reading tar: %w", err)
		}
		if header.Typeflag == tar.TypeReg && strings.HasSuffix(header.Name, binaryName) {
			return io.ReadAll(tr)
		}
	}
	return nil, fmt.Errorf("binary %q not found in archive", binaryName)
}

func extractFromZip(data []byte, binaryName string) ([]byte, error) {
	r, err := zip.NewReader(bytes.NewReader(data), int64(len(data)))
	if err != nil {
		return nil, fmt.Errorf("reading zip: %w", err)
	}

	for _, f := range r.File {
		if strings.HasSuffix(f.Name, binaryName) {
			rc, err := f.Open()
			if err != nil {
				return nil, err
			}
			defer func() { _ = rc.Close() }()
			return io.ReadAll(rc)
		}
	}
	return nil, fmt.Errorf("binary %q not found in archive", binaryName)
}
