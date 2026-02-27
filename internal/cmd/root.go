package cmd

import (
	"github.com/alpacahq/cli/internal/client"
	"github.com/alpacahq/cli/internal/config"
	"github.com/spf13/cobra"
)

var (
	version    = "dev"
	cfg        *config.Resolved
	apiClient  *client.Client
	outputFlag string
	profileFlag string
)

func SetVersion(v string) {
	version = v
	client.Version = v
}

func Execute() error {
	return rootCmd.Execute()
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
		// Auth commands handle their own config
		if cmd.Parent() != nil && cmd.Parent().Name() == "auth" {
			return nil
		}

		var err error
		cfg, err = config.Load(profileFlag, outputFlag)
		if err != nil {
			return err
		}

		if needsAuth(cmd) {
			if err := cfg.Validate(); err != nil {
				return err
			}
			apiClient = client.New(cfg)
		}

		return nil
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&outputFlag, "output", "o", "", "Output format: table, json, csv")
	rootCmd.PersistentFlags().StringVarP(&profileFlag, "profile", "p", "", "Config profile to use")

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(authCmd)
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
	rootCmd.AddCommand(apiCmd)
	rootCmd.AddCommand(updateCmd)
	rootCmd.AddCommand(configCmd)

	// Shortcuts
	rootCmd.AddCommand(buyCmd)
	rootCmd.AddCommand(sellCmd)
	rootCmd.AddCommand(priceCmd)
	rootCmd.AddCommand(positionsShortcut)
	rootCmd.AddCommand(ordersShortcut)
}

func needsAuth(cmd *cobra.Command) bool {
	switch cmd.Name() {
	case "version", "help", "completion", "update", "login", "logout", "status", "switch":
		return false
	}
	// config get/set doesn't need auth
	if cmd.Parent() != nil && cmd.Parent().Name() == "config" {
		return false
	}
	return true
}

func getOutput() string {
	if cfg != nil {
		return cfg.Output
	}
	if outputFlag != "" {
		return outputFlag
	}
	return "table"
}
