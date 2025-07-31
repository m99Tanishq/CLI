package cmd

import (
	"fmt"
	"strings"

	"github.com/m99Tanishq/CLI/internal/config"
	"github.com/m99Tanishq/CLI/pkg/utils"
	"github.com/spf13/cobra"
)

// maskAPIKey masks the API key for display
func maskAPIKey(apiKey string) string {
	if apiKey == "" {
		return "Not set"
	}
	if len(apiKey) <= 8 {
		return strings.Repeat("*", len(apiKey))
	}
	return apiKey[:4] + strings.Repeat("*", len(apiKey)-8) + apiKey[len(apiKey)-4:]
}

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage configuration",
	Long:  `View and modify CLI configuration settings.`,
	Run: func(cmd *cobra.Command, args []string) {
		ui := utils.NewModernUI()

		cfg, err := config.Load()
		if err != nil {
			ui.PrintError("Failed to load configuration")
			return
		}

		ui.PrintHeader("Configuration")

		// Print current configuration
		ui.PrintSection("Current Settings")

		ui.PrintTable([]string{"Setting", "Value"}, [][]string{
			{"Model", cfg.Model},
			{"API Key", maskAPIKey(cfg.APIKey)},
			{"Base URL", cfg.BaseURL},
			{"Max History", fmt.Sprintf("%d", cfg.MaxHistory)},
			{"Max Tokens", fmt.Sprintf("%d", cfg.MaxTokens)},
			{"Temperature", fmt.Sprintf("%.2f", cfg.Temperature)},
		})

		ui.PrintInfo("Use 'CLI config set <key> <value>' to modify settings")
		ui.PrintInfo("Available keys: model, api_key, base_url, max_history, max_tokens, temperature")
	},
}

var setCmd = &cobra.Command{
	Use:   "set [key] [value]",
	Short: "Set a configuration value",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		ui := utils.NewModernUI()

		key := args[0]
		value := args[1]

		err := config.Set(key, value)
		if err != nil {
			ui.PrintError(fmt.Sprintf("Failed to set %s: %v", key, err))
			return
		}

		ui.PrintSuccess(fmt.Sprintf("Successfully set %s to %s", key, value))

		// Show updated configuration
		cfg, err := config.Load()
		if err != nil {
			ui.PrintError("Failed to load updated configuration")
			return
		}

		ui.PrintSection("Updated Configuration")
		ui.PrintTable([]string{"Setting", "Value"}, [][]string{
			{"Model", cfg.Model},
			{"API Key", maskAPIKey(cfg.APIKey)},
			{"Base URL", cfg.BaseURL},
			{"Max Tokens", fmt.Sprintf("%d", cfg.MaxTokens)},
			{"Temperature", fmt.Sprintf("%.2f", cfg.Temperature)},
		})
	},
}

var listConfigCmd = &cobra.Command{
	Use:   "list",
	Short: "List current configuration",
	Run: func(cmd *cobra.Command, args []string) {
		ui := utils.NewModernUI()

		cfg, err := config.Load()
		if err != nil {
			ui.PrintError("Failed to load configuration")
			return
		}

		ui.PrintHeader("Current Configuration")

		ui.PrintTable([]string{"Setting", "Value"}, [][]string{
			{"Model", cfg.Model},
			{"API Key", maskAPIKey(cfg.APIKey)},
			{"Base URL", cfg.BaseURL},
			{"Max History", fmt.Sprintf("%d", cfg.MaxHistory)},
			{"Max Tokens", fmt.Sprintf("%d", cfg.MaxTokens)},
			{"Temperature", fmt.Sprintf("%.2f", cfg.Temperature)},
		})
	},
}

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset configuration to defaults",
	Run: func(cmd *cobra.Command, args []string) {
		ui := utils.NewModernUI()

		// Create default configuration
		defaultConfig := config.DefaultConfig()
		err := config.Save(defaultConfig)
		if err != nil {
			ui.PrintError("Failed to reset configuration")
			return
		}

		ui.PrintSuccess("Configuration reset to defaults")

		ui.PrintSection("Default Settings")
		ui.PrintTable([]string{"Setting", "Value"}, [][]string{
			{"Model", defaultConfig.Model},
			{"API Key", maskAPIKey(defaultConfig.APIKey)},
			{"Base URL", defaultConfig.BaseURL},
			{"Max History", fmt.Sprintf("%d", defaultConfig.MaxHistory)},
			{"Max Tokens", fmt.Sprintf("%d", defaultConfig.MaxTokens)},
			{"Temperature", fmt.Sprintf("%.2f", defaultConfig.Temperature)},
		})
	},
}

func init() {
	configCmd.AddCommand(setCmd)
	configCmd.AddCommand(listConfigCmd)
	configCmd.AddCommand(resetCmd)
}
