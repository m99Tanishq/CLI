package cmd

import (
	"fmt"
	"strings"

	"github.com/m99Tanishq/CLI/internal/config"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage GLM CLI configuration",
	Long:  `Manage configuration settings for the GLM CLI tool.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Load current configuration
		cfg, err := config.Load()
		if err != nil {
			fmt.Printf("Error loading config: %v\n", err)
			return
		}

		// Check flags
		setFlag, _ := cmd.Flags().GetString("set")
		getFlag, _ := cmd.Flags().GetString("get")
		listFlag, _ := cmd.Flags().GetBool("list")

		if setFlag != "" {
			// Set configuration value
			parts := strings.SplitN(setFlag, "=", 2)
			if len(parts) != 2 {
				fmt.Println("Error: Use format key=value")
				return
			}
			key, value := parts[0], parts[1]

			switch key {
			case "api_key":
				cfg.APIKey = value
			case "model":
				cfg.Model = value
			case "base_url":
				cfg.BaseURL = value
			case "max_history":
				// TODO: Add max_history setting
				fmt.Printf("Setting %s = %s\n", key, value)
			default:
				fmt.Printf("Unknown configuration key: %s\n", key)
				return
			}

			if err := config.Save(cfg); err != nil {
				fmt.Printf("Error saving config: %v\n", err)
				return
			}
			fmt.Printf("Configuration updated: %s = %s\n", key, value)

		} else if getFlag != "" {
			// Get configuration value
			switch getFlag {
			case "api_key":
				if cfg.APIKey == "" {
					fmt.Println("API key not set")
				} else {
					fmt.Printf("API key: %s\n", maskAPIKey(cfg.APIKey))
				}
			case "model":
				fmt.Printf("Model: %s\n", cfg.Model)
			case "base_url":
				fmt.Printf("Base URL: %s\n", cfg.BaseURL)
			case "max_history":
				fmt.Printf("Max history: %d\n", cfg.MaxHistory)
			default:
				fmt.Printf("Unknown configuration key: %s\n", getFlag)
			}

		} else if listFlag {
			// List all configuration values
			fmt.Println("Current configuration:")
			fmt.Printf("  API Key: %s\n", maskAPIKey(cfg.APIKey))
			fmt.Printf("  Model: %s\n", cfg.Model)
			fmt.Printf("  Base URL: %s\n", cfg.BaseURL)
			fmt.Printf("  Max History: %d\n", cfg.MaxHistory)

		} else {
			// Show help
			fmt.Println("Configuration management...")
			fmt.Println("Use --set key=value to set a configuration value")
			fmt.Println("Use --get key to get a configuration value")
			fmt.Println("Use --list to show all configuration values")
		}
	},
}

func init() {
	configCmd.Flags().StringP("set", "s", "", "Set a configuration value")
	configCmd.Flags().StringP("get", "g", "", "Get a configuration value")
	configCmd.Flags().BoolP("list", "l", false, "List all configuration values")
}

// maskAPIKey masks the API key for display
func maskAPIKey(apiKey string) string {
	if apiKey == "" {
		return "not set"
	}
	if len(apiKey) <= 8 {
		return "***"
	}
	return apiKey[:4] + "..." + apiKey[len(apiKey)-4:]
}
