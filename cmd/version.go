package cmd

import (
	"github.com/m99Tanishq/CLI/pkg/utils"
	"github.com/spf13/cobra"
)

var Version = "dev"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version information",
	Long:  `Display the current version of Rzork CLI and its capabilities.`,
	Run: func(cmd *cobra.Command, args []string) {
		ui := utils.NewModernUI()

		// Print banner
		ui.PrintBanner("CLI - AI-Powered Development Assistant")

		// Print version info
		ui.PrintModelInfo("CLI", Version)

		// Print capabilities
		capabilities := []string{
			"AI Chat with advanced language models",
			"Real-time streaming responses",
			"Code analysis and intelligent review",
			"File manipulation and management",
			"Memory system for codebase indexing",
			"Professional colored output",
			"Progress indicators and animations",
			"Cancellation support (Ctrl+C)",
			"Cross-platform compatibility",
			"Modern command-line interface",
		}
		ui.PrintCapabilities(capabilities)

		// Print usage examples
		examples := []string{
			"CLI chat --stream",
			"CLI code analyze main.go",
			"CLI memory index .",
			"CLI files list src/",
		}
		ui.PrintExamples(examples)

		// Print status
		ui.PrintStatus("🚀", "Ready for development")
		ui.PrintStatus("⚡", "Streaming enabled")
		ui.PrintStatus("🎨", "Modern UI active")

		// Print footer
		ui.PrintFooter()
	},
}

func init() {
	// versionCmd is added in root.go
}
