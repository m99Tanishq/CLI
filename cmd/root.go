package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "CLI",
	Short: "A CLI tool for interacting with Rzork models",
	Long: `Rzork CLI is a command-line interface for interacting with Rzork (General Language Model) APIs.
It provides features for chat, configuration management, and conversation history.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// Add subcommands
	rootCmd.AddCommand(chatCmd)
	rootCmd.AddCommand(configCmd)
	rootCmd.AddCommand(historyCmd)
	rootCmd.AddCommand(filesCmd)
	rootCmd.AddCommand(codeCmd)
	rootCmd.AddCommand(memoryCmd)
	rootCmd.AddCommand(versionCmd)
}
