package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/alpacahq/cli/internal/api"
	"github.com/alpacahq/cli/internal/client"
	"github.com/alpacahq/cli/internal/config"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

const (
	exitAPIError = 1

	outputJSON  = "json"
	outputCSV   = "csv"
	outputTable = "table"
)

var (
	version       = "dev"
	cfg           *config.Resolved
	apiClient     *client.Client
	tradingClient *api.TradingClient
	dataClient    *api.MarketDataClient
	jsonFlag      bool
	csvFlag       bool
	quietFlag     bool
	verboseFlag   bool
	debugFlag     bool
	profileFlag   string
	timeoutFlag   int
)

func SetVersion(v string) {
	version = v
	client.Version = v
}

func Root() *cobra.Command {
	return rootCmd
}

func Execute() error {
	err := rootCmd.Execute()
	if err != nil {
		var apiErr *client.APIError
		if errors.As(err, &apiErr) {
			if jsonFlag || quietFlag {
				printJSONError(apiErr)
			} else {
				fmt.Fprintf(os.Stderr, "Error: %s\n", apiErr)
				if hint := apiErr.Hint(); hint != "" {
					fmt.Fprintf(os.Stderr, "Hint: %s\n", hint)
				}
			}
			os.Exit(apiErr.ExitCode())
		}

		if jsonFlag || quietFlag {
			printJSONError(&client.APIError{Message: err.Error()})
		} else {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		}
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
	enc := json.NewEncoder(os.Stderr)
	_ = enc.Encode(m)
}

var rootCmd = &cobra.Command{
	Use:           "alpaca",
	Short:         "CLI for Alpaca Trading API",
	Long:          "Trade stocks & crypto, access market data, and manage your Alpaca account from the command line.",
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		if ha, _ := cmd.Flags().GetBool("help-all"); ha {
			printCommandTree(cmd.OutOrStdout(), cmd, 0)
			return nil
		}
		return cmd.Help()
	},
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if cmd.Name() == "help" || cmd.Name() == "completion" || cmd.Name() == "version" {
			return nil
		}
		if ha, _ := cmd.Flags().GetBool("help-all"); ha {
			return nil
		}
		if cmd.Parent() != nil && cmd.Parent().Name() == "profile" {
			return nil
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

		var err error
		outputOverride := ""
		if jsonFlag {
			outputOverride = outputJSON
		} else if csvFlag {
			outputOverride = outputCSV
		}
		cfg, err = config.Load(profileFlag, outputOverride)
		if err != nil {
			return err
		}

		if needsAuth(cmd) {
			if err := cfg.Validate(); err != nil {
				return err
			}
			apiClient = client.New(cfg)
			apiClient.Verbose = verboseFlag || debugFlag
			apiClient.Debug = debugFlag
			apiClient.Quiet = quietFlag
			if timeoutFlag != 30 {
				apiClient.SetTimeout(time.Duration(timeoutFlag) * time.Second)
			}
			tradingClient = api.NewTradingClient(apiClient)
			dataClient = api.NewMarketDataClient(apiClient)
		}

		return nil
	},
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&jsonFlag, "json", false, "Output as JSON")
	rootCmd.PersistentFlags().BoolVar(&csvFlag, "csv", false, "Output as CSV")
	rootCmd.MarkFlagsMutuallyExclusive("json", "csv")
	rootCmd.PersistentFlags().StringVarP(&profileFlag, "profile", "p", "", "Config profile to use")
	rootCmd.PersistentFlags().BoolVarP(&verboseFlag, "verbose", "v", false, "Show HTTP request details on stderr")
	rootCmd.PersistentFlags().BoolVar(&debugFlag, "debug", false, "Show full HTTP request/response bodies on stderr (implies --verbose)")
	rootCmd.PersistentFlags().BoolVarP(&quietFlag, "quiet", "q", false, "Suppress non-data output (warnings, hints, color)")
	rootCmd.PersistentFlags().IntVar(&timeoutFlag, "timeout", 30, "HTTP request timeout in seconds")

	rootCmd.Flags().Bool("help-all", false, "Print full reference for every command")

	tradingGroup := &cobra.Group{ID: "trading", Title: "Trading"}
	dataGroup := &cobra.Group{ID: "data", Title: "Market Data"}
	accountGroup := &cobra.Group{ID: "account", Title: "Account & Assets"}
	utilGroup := &cobra.Group{ID: "util", Title: "Utilities"}

	rootCmd.AddGroup(tradingGroup, dataGroup, accountGroup, utilGroup)

	orderCmd.GroupID = tradingGroup.ID
	positionCmd.GroupID = tradingGroup.ID
	optionCmd.GroupID = tradingGroup.ID

	dataCmd.GroupID = dataGroup.ID
	screenerCmd.GroupID = dataGroup.ID
	newsCmd.GroupID = dataGroup.ID

	accountCmd.GroupID = accountGroup.ID
	activityCmd.GroupID = accountGroup.ID
	assetCmd.GroupID = accountGroup.ID
	portfolioCmd.GroupID = accountGroup.ID
	corporateActionCmd.GroupID = accountGroup.ID
	watchlistCmd.GroupID = accountGroup.ID
	walletCmd.GroupID = accountGroup.ID

	profileCmd.GroupID = utilGroup.ID
	clockCmd.GroupID = utilGroup.ID
	calendarCmd.GroupID = utilGroup.ID
	apiCmd.GroupID = utilGroup.ID
	setupCmd.GroupID = utilGroup.ID
	updateCmd.GroupID = utilGroup.ID
	versionCmd.GroupID = utilGroup.ID

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(profileCmd)
	rootCmd.AddCommand(accountCmd)
	rootCmd.AddCommand(orderCmd)
	rootCmd.AddCommand(positionCmd)
	rootCmd.AddCommand(assetCmd)
	rootCmd.AddCommand(dataCmd)
	rootCmd.AddCommand(watchlistCmd)
	rootCmd.AddCommand(clockCmd)
	rootCmd.AddCommand(calendarCmd)
	rootCmd.AddCommand(portfolioCmd)
	rootCmd.AddCommand(newsCmd)
	rootCmd.AddCommand(optionCmd)
	rootCmd.AddCommand(activityCmd)
	rootCmd.AddCommand(screenerCmd)
	rootCmd.AddCommand(corporateActionCmd)
	rootCmd.AddCommand(walletCmd)
	rootCmd.AddCommand(apiCmd)
	rootCmd.AddCommand(updateCmd)
}

func needsAuth(cmd *cobra.Command) bool {
	switch cmd.Name() {
	case "version", "help", "completion", "update", "setup":
		return false
	}
	return true
}

func getOutput() string {
	if jsonFlag {
		return outputJSON
	}
	if csvFlag {
		return outputCSV
	}
	if cfg != nil {
		return cfg.Output
	}
	return outputTable
}

func verboseLog(format string, args ...any) {
	if verboseFlag {
		fmt.Fprintf(os.Stderr, format+"\n", args...)
	}
}
