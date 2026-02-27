package cmd

import (
	"fmt"

	"github.com/alpacahq/cli/internal/config"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage CLI configuration",
}

var configGetCmd = &cobra.Command{
	Use:   "get <key>",
	Short: "Get a config value",
	Example: `  alpaca config get output
  alpaca config get default_profile`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		resolved, err := config.Load("", "")
		if err != nil {
			return err
		}

		switch args[0] {
		case "output":
			fmt.Println(resolved.Output)
		case "color":
			fmt.Println(resolved.Color)
		case "default_profile":
			fmt.Println(resolved.ProfileName)
		case "base_url":
			fmt.Println(resolved.BaseURL)
		case "data_url":
			fmt.Println(resolved.DataURL)
		case "environment":
			fmt.Println(resolved.Environment)
		default:
			return fmt.Errorf("unknown config key: %s\nAvailable keys: output, color, default_profile, base_url, data_url, environment", args[0])
		}
		return nil
	},
}

var configSetCmd = &cobra.Command{
	Use:   "set <key> <value>",
	Short: "Set a config value",
	Example: `  alpaca config set output json
  alpaca config set color never
  alpaca config set default_profile live`,
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		key, value := args[0], args[1]

		globalCfg := loadOrCreateGlobal()

		switch key {
		case "output":
			if value != "table" && value != "json" && value != "csv" {
				return fmt.Errorf("output must be one of: table, json, csv")
			}
			globalCfg.Output = value
		case "color":
			if value != "auto" && value != "always" && value != "never" {
				return fmt.Errorf("color must be one of: auto, always, never")
			}
			globalCfg.Color = value
		case "default_profile":
			globalCfg.DefaultProfile = value
		default:
			return fmt.Errorf("unknown config key: %s\nSettable keys: output, color, default_profile", key)
		}

		if err := config.SaveGlobalConfig(globalCfg); err != nil {
			return err
		}

		fmt.Printf("Set %s = %s\n", key, value)
		return nil
	},
}

func init() {
	configCmd.AddCommand(configGetCmd)
	configCmd.AddCommand(configSetCmd)
}
