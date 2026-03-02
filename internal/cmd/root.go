package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/alpacahq/cli/internal/api"
	"github.com/alpacahq/cli/internal/client"
	"github.com/alpacahq/cli/internal/config"
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
	jsonFlag      bool
	csvFlag       bool
	quietFlag     bool
	verboseFlag   bool
	confirmFlag   bool
	profileFlag   string
)

func SetVersion(v string) {
	version = v
	client.Version = v
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
	enc := json.NewEncoder(os.Stderr)
	_ = enc.Encode(map[string]any{
		"error":   apiErr.Message,
		"code":    apiErr.Code,
		"status":  apiErr.StatusCode,
		"hint":    apiErr.Hint(),
	})
}

var rootCmd = &cobra.Command{
	Use:   "alpaca",
	Short: "CLI for Alpaca Trading API",
	Long:  "Trade stocks & crypto, access market data, and manage your Alpaca account from the command line.",
	SilenceUsage:  true,
	SilenceErrors: true,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if cmd.Name() == "help" || cmd.Name() == "completion" || cmd.Name() == "version" {
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

		var err error
		cfg, err = config.Load(profileFlag, resolveOutputFlag())
		if err != nil {
			return err
		}

		if needsAuth(cmd) {
			if err := cfg.Validate(); err != nil {
				return err
			}
			apiClient = client.New(cfg)
			apiClient.Verbose = verboseFlag
			apiClient.Quiet = quietFlag
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
	rootCmd.PersistentFlags().BoolVarP(&quietFlag, "quiet", "q", false, "Suppress non-data output (warnings, hints, color)")
	rootCmd.PersistentFlags().BoolVar(&confirmFlag, "confirm", false, "Authorize destructive operations (cancel-all, close-all on live)")

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

	// Shortcuts
	rootCmd.AddCommand(buyCmd)
	rootCmd.AddCommand(sellCmd)
	rootCmd.AddCommand(priceCmd)
	rootCmd.AddCommand(positionsShortcut)
	rootCmd.AddCommand(ordersShortcut)
}

func needsAuth(cmd *cobra.Command) bool {
	switch cmd.Name() {
	case "version", "help", "completion", "update":
		return false
	}
	return true
}

func resolveOutputFlag() string {
	if jsonFlag {
		return "json"
	}
	if csvFlag {
		return "csv"
	}
	return ""
}

func getOutput() string {
	if jsonFlag {
		return "json"
	}
	if csvFlag {
		return "csv"
	}
	if cfg != nil {
		return cfg.Output
	}
	return "table"
}

func isLive() bool {
	if cfg == nil {
		return false
	}
	return strings.Contains(cfg.BaseURL, "api.alpaca.markets") && !strings.Contains(cfg.BaseURL, "paper")
}

func warnLive() {
	if !isLive() || quietFlag {
		return
	}
	if cfg != nil && cfg.SuppressWarnings {
		return
	}
	_, _ = color.New(color.FgYellow).Fprintln(os.Stderr, "⚠ Live trading account. This order will use real money.")
}

func requireConfirmation(prompt string) error {
	if confirmFlag {
		return nil
	}
	return fmt.Errorf("%s\nHint: pass --confirm to proceed", prompt)
}
