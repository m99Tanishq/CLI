package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/m99Tanishq/CLI/internal/api"
	"github.com/m99Tanishq/CLI/internal/config"
	"github.com/m99Tanishq/CLI/pkg/utils"
	"github.com/spf13/cobra"
)

var codeCmd = &cobra.Command{
	Use:   "code",
	Short: "Code analysis and generation",
	Long:  `Analyze, fix, review, and generate code using AI.`,
}

var analyzeCmd = &cobra.Command{
	Use:   "analyze [file]",
	Short: "Analyze code with AI",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ui := utils.NewModernUI()
		filePath := args[0]

		content, err := os.ReadFile(filePath)
		if err != nil {
			ui.PrintError(fmt.Sprintf("Error reading file: %v", err))
			return
		}

		cfg, err := config.Load()
		if err != nil {
			ui.PrintError("Failed to load configuration")
			return
		}

		if cfg.APIKey == "" {
			ui.PrintError("API key not configured")
			ui.PrintInfo("Please run: CLI config set api_key YOUR_HUGGING_FACE_API_KEY")
			return
		}

		client := api.NewClient(cfg.APIKey, cfg.BaseURL)

		ui.PrintHeader("Code Analysis")
		ui.PrintInfo(fmt.Sprintf("Analyzing: %s", filePath))
		ui.PrintLoading("Processing code analysis...")

		prompt := fmt.Sprintf(`Please analyze this code file and identify any issues, bugs, or areas for improvement:

File: %s
Content:
%s

Please provide:
1. Potential bugs or errors
2. Code quality issues
3. Security concerns
4. Performance improvements
5. Best practices violations

Format your response clearly with sections.`, filePath, string(content))

		req := api.ChatRequest{
			Model: cfg.Model,
			Messages: []api.Message{
				{
					Role:    "user",
					Content: prompt,
				},
			},
		}

		resp, err := client.SendChat(req)
		if err != nil {
			ui.PrintError(fmt.Sprintf("Error analyzing code: %v", err))
			return
		}

		if len(resp.Choices) > 0 {
			ui.PrintSection("Analysis Results")
			ui.PrintCard("Code Analysis Report", resp.Choices[0].Message.Content)
		}
	},
}

var fixCmd = &cobra.Command{
	Use:   "fix [file]",
	Short: "Fix code issues with AI",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ui := utils.NewModernUI()
		filePath := args[0]

		content, err := os.ReadFile(filePath)
		if err != nil {
			ui.PrintError(fmt.Sprintf("Error reading file: %v", err))
			return
		}

		cfg, err := config.Load()
		if err != nil {
			ui.PrintError("Failed to load configuration")
			return
		}

		if cfg.APIKey == "" {
			ui.PrintError("API key not configured")
			ui.PrintInfo("Please run: CLI config set api_key YOUR_HUGGING_FACE_API_KEY")
			return
		}

		client := api.NewClient(cfg.APIKey, cfg.BaseURL)

		ui.PrintHeader("Code Fix")
		ui.PrintInfo(fmt.Sprintf("Fixing issues in: %s", filePath))
		ui.PrintLoading("Analyzing and fixing code...")

		prompt := fmt.Sprintf(`Please analyze and fix any issues in this code file. Return the corrected code:

File: %s
Content:
%s

Please:
1. Identify any bugs, errors, or issues
2. Provide the corrected code
3. Explain what was fixed
4. Ensure the code follows best practices

Return the corrected code in a code block.`, filePath, string(content))

		req := api.ChatRequest{
			Model: cfg.Model,
			Messages: []api.Message{
				{
					Role:    "user",
					Content: prompt,
				},
			},
		}

		resp, err := client.SendChat(req)
		if err != nil {
			ui.PrintError(fmt.Sprintf("Error fixing code: %v", err))
			return
		}

		if len(resp.Choices) > 0 {
			ui.PrintSection("Fixed Code")
			ui.PrintCard("Code Fix Results", resp.Choices[0].Message.Content)

			ui.PrintPrompt("Do you want to apply these changes? (y/n): ")
			var response string
			if _, err := fmt.Scanln(&response); err != nil {
				ui.PrintError("Error reading input")
				return
			}

			if strings.ToLower(response) == "y" || strings.ToLower(response) == "yes" {
				responseContent := resp.Choices[0].Message.Content
				if strings.Contains(responseContent, "```") {
					start := strings.Index(responseContent, "```")
					if start != -1 {
						start = strings.Index(responseContent[start:], "\n") + start + 1
						end := strings.LastIndex(responseContent, "```")
						if end > start {
							code := responseContent[start:end]
							err := os.WriteFile(filePath, []byte(code), 0600)
							if err != nil {
								ui.PrintError(fmt.Sprintf("Error writing fixed code: %v", err))
								return
							}
							ui.PrintSuccess(fmt.Sprintf("Successfully applied fixes to %s", filePath))
						}
					}
				}
			}
		}
	},
}

var reviewCmd = &cobra.Command{
	Use:   "review [file]",
	Short: "Code review with AI",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ui := utils.NewModernUI()
		filePath := args[0]

		content, err := os.ReadFile(filePath)
		if err != nil {
			ui.PrintError(fmt.Sprintf("Error reading file: %v", err))
			return
		}

		cfg, err := config.Load()
		if err != nil {
			ui.PrintError("Failed to load configuration")
			return
		}

		if cfg.APIKey == "" {
			ui.PrintError("API key not configured")
			ui.PrintInfo("Please run: CLI config set api_key YOUR_HUGGING_FACE_API_KEY")
			return
		}

		client := api.NewClient(cfg.APIKey, cfg.BaseURL)

		ui.PrintHeader("Code Review")
		ui.PrintInfo(fmt.Sprintf("Reviewing: %s", filePath))
		ui.PrintLoading("Performing comprehensive code review...")

		prompt := fmt.Sprintf(`Please perform a comprehensive code review of this file:

File: %s
Content:
%s

Please provide a detailed review covering:
1. Code quality and readability
2. Architecture and design patterns
3. Performance considerations
4. Security implications
5. Maintainability
6. Suggestions for improvement
7. Overall rating (1-10)

Format your response as a professional code review.`, filePath, string(content))

		req := api.ChatRequest{
			Model: cfg.Model,
			Messages: []api.Message{
				{
					Role:    "user",
					Content: prompt,
				},
			},
		}

		resp, err := client.SendChat(req)
		if err != nil {
			ui.PrintError(fmt.Sprintf("Error reviewing code: %v", err))
			return
		}

		if len(resp.Choices) > 0 {
			ui.PrintSection("Code Review Results")
			ui.PrintCard("Review Report", resp.Choices[0].Message.Content)
		}
	},
}

func init() {
	codeCmd.AddCommand(analyzeCmd)
	codeCmd.AddCommand(fixCmd)
	codeCmd.AddCommand(reviewCmd)
}
