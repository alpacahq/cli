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
	Use:   "api",
	Short: "Raw API access",
	Long:  "Make raw API calls to any Alpaca endpoint. Paths are relative to the base URL.",
}

var apiGetCmd = &cobra.Command{
	Use:   "get <path>",
	Short: "GET request",
	Example: `  alpaca api get /v2/account
  alpaca api get /v2/orders
  alpaca api get /v2/assets/AAPL`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return rawAPI(cmd, "GET", args[0])
	},
}

var apiPostCmd = &cobra.Command{
	Use:   "post <path>",
	Short: "POST request",
	Example: `  alpaca api post /v2/orders --data '{"symbol":"AAPL","qty":"1","side":"buy","type":"market","time_in_force":"day"}'
  echo '{"symbol":"AAPL","qty":"1","side":"buy","type":"market","time_in_force":"day"}' | alpaca api post /v2/orders`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return rawAPI(cmd, "POST", args[0])
	},
}

var apiPatchCmd = &cobra.Command{
	Use:     "patch <path>",
	Short:   "PATCH request",
	Example: `  alpaca api patch /v2/account/configurations --data '{"no_shorting":true}'`,
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return rawAPI(cmd, "PATCH", args[0])
	},
}

var apiDeleteCmd = &cobra.Command{
	Use:     "delete <path>",
	Short:   "DELETE request",
	Example: `  alpaca api delete /v2/orders/61e69015-8549-4baf-b96f-9c4f3e8d0c35`,
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return rawAPI(cmd, "DELETE", args[0])
	},
}

func rawAPI(cmd *cobra.Command, method, path string) error {
	dataFlag := cmdutil.Str(cmd, "data")
	useData := cmdutil.Bool(cmd, "use-data-api")

	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	var fullURL string
	if useData {
		fullURL = cfg.DataURL + path
	} else {
		fullURL = cfg.BaseURL + path
	}

	var body any
	switch {
	case dataFlag != "":
		var m map[string]any
		if err := json.Unmarshal([]byte(dataFlag), &m); err != nil {
			return fmt.Errorf("invalid JSON in --data: %w", err)
		}
		body = m
	case stdinHasData():
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

	return renderData(cmd.OutOrStdout(), data, false)
}

func stdinHasData() bool {
	stat, err := os.Stdin.Stat()
	if err != nil {
		return false
	}
	return (stat.Mode() & os.ModeCharDevice) == 0
}

func init() {
	for _, cmd := range []*cobra.Command{apiGetCmd, apiPostCmd, apiPatchCmd, apiDeleteCmd} {
		cmd.Flags().String("data", "", "JSON request body")
		cmd.Flags().Bool("use-data-api", false, "Use data API base URL instead of trading API")
	}

	apiCmd.AddCommand(apiGetCmd)
	apiCmd.AddCommand(apiPostCmd)
	apiCmd.AddCommand(apiPatchCmd)
	apiCmd.AddCommand(apiDeleteCmd)
}
