package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/alpacahq/cli/internal/cmdutil"
	"github.com/spf13/cobra"
)

var apiCmd = &cobra.Command{
	Use:   "api [METHOD] <path>",
	Short: "Raw API access",
	Long: `Make raw API calls to any Alpaca endpoint. Paths are relative to the base URL.

METHOD defaults to GET if omitted.`,
	Example: `  alpaca api /v2/account
  alpaca api GET /v2/orders
  alpaca api POST /v2/orders --body '{"symbol":"AAPL","qty":"1","side":"buy","type":"market","time_in_force":"day"}'
  alpaca api GET /v2/orders --query "status=open&limit=10"
  echo '{"symbol":"AAPL","qty":"1","side":"buy","type":"market","time_in_force":"day"}' | alpaca api POST /v2/orders`,
	Args: cobra.RangeArgs(1, 2),
	RunE: func(cmd *cobra.Command, args []string) error {
		method, path := parseMethodPath(args)
		bodyFlag := cmdutil.Str(cmd, "body")
		queryFlag := cmdutil.Str(cmd, "query")
		useData := cmdutil.Bool(cmd, "use-data-api")

		if !strings.HasPrefix(path, "/") {
			path = "/" + path
		}
		if queryFlag != "" {
			path += "?" + queryFlag
		}

		var fullURL string
		if useData {
			fullURL = cfg.DataURL + path
		} else {
			fullURL = cfg.BaseURL + path
		}

		var body any
		switch {
		case bodyFlag != "":
			var m map[string]any
			if err := json.Unmarshal([]byte(bodyFlag), &m); err != nil {
				return fmt.Errorf("invalid JSON in --body: %w", err)
			}
			body = m
		case methodSupportsBody(method) && stdinHasData():
			raw, err := io.ReadAll(os.Stdin)
			if err != nil {
				return fmt.Errorf("reading stdin: %w", err)
			}
			if len(raw) > 0 {
				var m map[string]any
				if err := json.Unmarshal(raw, &m); err != nil {
					return fmt.Errorf("invalid JSON on stdin: %w", err)
				}
				body = m
			}
		}

		data, err := apiClient.RawRequest(method, fullURL, body)
		if err != nil {
			return err
		}

		return renderData(cmd, data)
	},
}

var validMethods = map[string]bool{
	"GET": true, "POST": true, "PUT": true, "PATCH": true, "DELETE": true,
}

func parseMethodPath(args []string) (method, path string) {
	if len(args) == 1 {
		return "GET", args[0]
	}
	upper := strings.ToUpper(args[0])
	if validMethods[upper] {
		return upper, args[1]
	}
	return "GET", args[0]
}

func stdinHasData() bool {
	stat, err := os.Stdin.Stat()
	if err != nil {
		return false
	}
	return (stat.Mode() & os.ModeCharDevice) == 0
}

func methodSupportsBody(method string) bool {
	switch method {
	case "POST", "PUT", "PATCH":
		return true
	default:
		return false
	}
}

func init() {
	apiCmd.Flags().String("body", "", "JSON request body")
	apiCmd.Flags().String("query", "", "Query string to append (e.g. \"status=open&limit=10\")")
	apiCmd.Flags().Bool("use-data-api", false, "Use data API base URL instead of trading API")
}
