package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/alpacahq/cli/internal/api"
	"github.com/alpacahq/cli/internal/client"
	"github.com/alpacahq/cli/internal/config"
	"github.com/alpacahq/cli/internal/oauth"
	"github.com/alpacahq/cli/internal/output"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

const exitAPIError = 1

var (
	version       = "dev"
	cfg           *config.Resolved
	apiClient     *client.Client
	tradingClient *api.TradingClient
	dataClient    *api.MarketDataClient
	csvFlag       bool
	jqFlag        string
	quietFlag     bool
	verboseFlag   bool
	debugFlag     bool
	traceFlag     bool
	schemaFlag    bool
	versionFlag   bool
	profileFlag   string
	timeoutFlag   int
)

func SetVersion(v string) {
	version = v
	client.Version = v
	oauth.UserAgent = "alpaca-cli/" + v
}

func Root() *cobra.Command {
	return rootCmd
}

func Execute() error {
	err := rootCmd.Execute()
	if err != nil {
		var apiErr *client.APIError
		if errors.As(err, &apiErr) {
			printJSONError(apiErr)
			os.Exit(apiErr.ExitCode())
		}

		printJSONError(&client.APIError{Message: err.Error()})
		os.Exit(exitAPIError)
	}
	return nil
}

func printJSONError(apiErr *client.APIError) {
	m := map[string]any{
		"error":  apiErr.Message,
		"code":   apiErr.Code,
		"status": apiErr.StatusCode,
		"hint":   apiErr.Hint(),
	}
	if apiErr.Method != "" {
		m["method"] = apiErr.Method
	}
	if apiErr.Path != "" {
		m["path"] = apiErr.Path
	}
	if apiErr.RequestID != "" {
		m["request_id"] = apiErr.RequestID
	}
	enc := json.NewEncoder(os.Stderr)
	enc.SetIndent("", "  ")
	_ = enc.Encode(m)
}

var versionCmd = &cobra.Command{
	Use:     "version",
	Short:   "Print the version of alpaca CLI",
	Long:    `Print the version of alpaca CLI.`,
	Example: `  alpaca version`,
	RunE: func(cmd *cobra.Command, args []string) error {
		_, err := fmt.Fprintln(cmd.OutOrStdout(), version)
		return err
	},
}

var rootCmd = &cobra.Command{
	Use:   "alpaca",
	Short: "CLI for Alpaca Trading API",
	Long: `Trade stocks & crypto, access market data, and manage your Alpaca account from the command line.

To update:  alpaca update`,
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		if versionFlag {
			_, err := fmt.Fprintln(cmd.OutOrStdout(), version)
			return err
		}
		if ha, _ := cmd.Flags().GetBool("help-all"); ha {
			printCommandTree(cmd.OutOrStdout(), cmd, 0)
			return nil
		}
		shouldWarn := cfg != nil && envShadowsProfile(cfg.ProfileName)
		if shouldWarn {
			warnEnvShadowsProfile(cfg.ProfileName, "")
			fmt.Fprintln(cmd.OutOrStdout())
		}
		if err := cmd.Help(); err != nil {
			return err
		}
		if shouldWarn {
			fmt.Fprintln(cmd.OutOrStdout())
			warnEnvShadowsProfile(cfg.ProfileName, "")
		}
		printUpdateNoticeIfAvailable(cmd.OutOrStdout(), 2*time.Second)
		return nil
	},
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if cmd.Name() == "help" {
			return nil
		}
		if ha, _ := cmd.Flags().GetBool("help-all"); ha {
			return nil
		}
		if schemaFlag {
			err := printCommandSchema(cmd)
			cmd.RunE = func(*cobra.Command, []string) error { return nil }
			return err
		}
		if cmd.Parent() != nil && cmd.Parent().Name() == "profile" {
			return nil
		}

		if os.Getenv("ALPACA_QUIET") != "" {
			quietFlag = true
		}
		if quietFlag {
			color.NoColor = true
		}

		if os.Getenv("ALPACA_VERBOSE") != "" {
			verboseFlag = true
		}
		if os.Getenv("ALPACA_DEBUG") != "" {
			debugFlag = true
		}
		if os.Getenv("ALPACA_TRACE") != "" {
			traceFlag = true
		}

		var err error
		outputOverride := ""
		if csvFlag {
			outputOverride = string(output.FormatCSV)
		}
		cfg, err = config.Load(profileFlag, outputOverride)
		if err != nil {
			return err
		}

		if needsAuth(cmd) {
			if err := cfg.Validate(); err != nil {
				return err
			}
			// Tests can pre-inject apiClient (and trading/dataClient) to
			// redirect traffic at a mock server; only build a fresh client
			// when nothing's been injected. CLI flags are applied either
			// way so --debug/--verbose/--quiet/--trace/--timeout work the
			// same under tests as in production.
			if apiClient == nil {
				apiClient = client.New(cfg)
				tradingClient = api.NewTradingClient(apiClient)
				dataClient = api.NewMarketDataClient(apiClient)
			}
			apiClient.Verbose = verboseFlag
			apiClient.Debug = debugFlag
			apiClient.Quiet = quietFlag
			apiClient.Trace = traceFlag
			if timeoutFlag != 30 {
				apiClient.SetTimeout(time.Duration(timeoutFlag) * time.Second)
			}
		}

		return nil
	},
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&csvFlag, "csv", false, "Output as CSV")
	rootCmd.PersistentFlags().StringVar(&jqFlag, "jq", "", "Filter output with a jq expression")
	rootCmd.PersistentFlags().StringVarP(&profileFlag, "profile", "p", "", "Config profile to use")
	rootCmd.PersistentFlags().BoolVarP(&verboseFlag, "verbose", "v", false, "Show HTTP request details on stderr")
	rootCmd.PersistentFlags().BoolVar(&debugFlag, "debug", false, "Show HTTP request/response headers and bodies on stderr")
	rootCmd.PersistentFlags().BoolVar(&traceFlag, "trace", false, "Show HTTP timing breakdown on stderr (DNS, TLS, TTFB)")
	rootCmd.PersistentFlags().BoolVarP(&quietFlag, "quiet", "q", false, "Suppress non-data output (warnings, hints, color)")
	rootCmd.PersistentFlags().IntVar(&timeoutFlag, "timeout", 30, "HTTP request timeout in seconds")
	rootCmd.PersistentFlags().BoolVar(&schemaFlag, "schema", false, "Show response schema for this command and exit")

	rootCmd.Flags().Bool("help-all", false, "Print full reference for every command")
	rootCmd.Flags().BoolVar(&versionFlag, "version", false, "Print the version of alpaca CLI")

	tradingGroup := &cobra.Group{ID: "trading", Title: "Trading"}
	accountGroup := &cobra.Group{ID: "account", Title: "Account & Assets"}
	utilGroup := &cobra.Group{ID: "util", Title: "Utilities"}

	rootCmd.AddGroup(tradingGroup, accountGroup, utilGroup)

	addGroup(rootCmd, tradingGroup.ID, orderCmd, positionCmd, optionCmd, locateCmd, clockCmd, calendarCmd, dataCmd, cryptoPerpCmd)
	addGroup(rootCmd, accountGroup.ID, accountCmd, assetCmd, corporateActionCmd, watchlistCmd, walletCmd)
	addGroup(rootCmd, utilGroup.ID, profileCmd, apiCmd, updateCmd, versionCmd, doctorCmd)
}

func addGroup(parent *cobra.Command, groupID string, cmds ...*cobra.Command) {
	for _, c := range cmds {
		c.GroupID = groupID
		parent.AddCommand(c)
	}
}

func needsAuth(cmd *cobra.Command) bool {
	for c := cmd; c != nil; c = c.Parent() {
		if c.Parent() != nil && c.Parent().Parent() == nil {
			switch c.Name() {
			case "version", "help", "completion", "update", "doctor":
				return false
			}
		}
	}
	return true
}

func getOutput() output.Format {
	if csvFlag {
		return output.FormatCSV
	}
	if cfg != nil {
		return output.Format(cfg.Output)
	}
	return output.FormatJSON
}

// renderData is the unified output pipeline: jq filter → format render.
func renderData(cmd *cobra.Command, data any) error {
	if jqFlag != "" {
		var err error
		data, err = output.ApplyJQ(data, jqFlag)
		if err != nil {
			return err
		}
	}
	if getOutput() == output.FormatCSV {
		return output.CSVWithHeaders(cmd.OutOrStdout(), data, csvHeadersForCommand(cmd))
	}
	return output.Render(cmd.OutOrStdout(), getOutput(), data)
}

func csvHeadersForCommand(cmd *cobra.Command) []string {
	if jqFlag != "" || cmd == nil {
		return nil
	}
	opName := cmd.Annotations["op"]
	if opName == "" {
		return nil
	}
	fields, ok := api.ResponseSchema(opName)
	if !ok || len(fields) == 0 {
		return nil
	}
	headers := make([]string, 0, len(fields))
	for _, f := range fields {
		headers = append(headers, f.Name)
	}
	return headers
}

var (
	schemaComment = color.New(color.Faint)
	schemaType    = color.New(color.FgCyan)
	schemaEnum    = color.New(color.FgGreen)
)

func printCommandSchema(cmd *cobra.Command) error {
	opName := cmd.Annotations["op"]
	if opName == "" {
		return fmt.Errorf("no response schema available for %q", cmd.CommandPath())
	}
	fields, ok := api.ResponseSchema(opName)
	if !ok {
		return fmt.Errorf("no response schema available for %q", cmd.CommandPath())
	}

	w := cmd.OutOrStdout()

	if op, ok := api.OpByName(opName); ok {
		if op.ReturnsArray {
			schemaComment.Fprintf(w, "// %s — returns an array of:\n", op.Summary)
		} else if op.Summary != "" {
			schemaComment.Fprintf(w, "// %s\n", op.Summary)
		}
	}
	fmt.Fprintln(w, "{")

	for _, f := range fields {
		ts := tsTypeColorized(f)
		desc := firstLine(f.Description)
		if desc != "" {
			fmt.Fprintf(w, "  %s: %s; %s\n", f.Name, ts, schemaComment.Sprintf("// %s", desc))
		} else {
			fmt.Fprintf(w, "  %s: %s;\n", f.Name, ts)
		}
	}

	fmt.Fprintln(w, "}")
	return nil
}

var oasToTS = map[string]string{
	"string":    "string",
	"boolean":   "boolean",
	"integer":   "number",
	"number":    "number",
	"enum":      "string",
	"object":    "object",
	"any":       "unknown",
	"[]string":  "string[]",
	"[]integer": "number[]",
	"[]number":  "number[]",
	"[]boolean": "boolean[]",
	"[]object":  "object[]",
	"[]enum":    "string[]",
}

func tsTypeColorized(f api.ResponseField) string {
	if len(f.EnumValues) > 0 {
		isArray := strings.HasPrefix(f.Type, "[]")
		parts := make([]string, len(f.EnumValues))
		for i, v := range f.EnumValues {
			parts[i] = schemaEnum.Sprintf("%q", v)
		}
		union := strings.Join(parts, " | ")
		if isArray {
			return "(" + union + ")[]"
		}
		return union
	}

	raw := tsTypePlain(f.Type)
	return schemaType.Sprint(raw)
}

func tsTypePlain(t string) string {
	if ts, ok := oasToTS[t]; ok {
		return ts
	}
	if strings.HasPrefix(t, "map[string]") {
		return "Record<string, " + t[len("map[string]"):] + ">"
	}
	if strings.HasPrefix(t, "[]") {
		return t[2:] + "[]"
	}
	return t
}

func firstLine(s string) string {
	s = strings.TrimSpace(s)
	if i := strings.IndexAny(s, "\n\r"); i >= 0 {
		s = s[:i]
	}
	return strings.TrimSpace(s)
}
