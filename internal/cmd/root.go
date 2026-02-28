package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/alpacahq/cli/internal/api"
	"github.com/alpacahq/cli/internal/client"
	"github.com/alpacahq/cli/internal/config"
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
			fmt.Fprintf(os.Stderr, "Error: %s\n", apiErr)
			os.Exit(apiErr.ExitCode())
		}
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(exitAPIError)
	}
	return nil
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
