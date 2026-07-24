// Package useragent builds the CLI's User-Agent header, following the same
// convention used across Alpaca's SDKs (e.g. the Go SDK's "APCA-GO/<version>
// GoRuntime/<go-version>"): APCA-<PLATFORM>/<sdk-version> <OS>/<arch>.
package useragent

import (
	"fmt"
	"runtime"
)

// Build returns the User-Agent string for the given CLI version, e.g.
// "APCA-CLI/0.0.13 darwin/arm64".
func Build(version string) string {
	return fmt.Sprintf("APCA-CLI/%s %s/%s", version, runtime.GOOS, runtime.GOARCH)
}
