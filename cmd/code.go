package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/m99Tanishq/CLI/internal/api"
	"github.com/m99Tanishq/CLI/internal/config"
	"github.com/spf13/cobra"
)

var codeCmd = &cobra.Command{
	Use:   "code",
	Short: "Analyze and fix code with AI",
	Long:  `Use AI to analyze code files, find issues, and suggest fixes.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Code analysis commands:")
		fmt.Println("  analyze <file>  - Analyze code and find issues")
		fmt.Println("  fix <file>      - Fix code issues with AI")
		fmt.Println("  review <file>   - Code review with AI")
		fmt.Println("  optimize <file> - Optimize code performance")
		fmt.Println("  explain <file>  - Explain code functionality")
	},
}

var analyzeCmd = &cobra.Command{
	Use:   "analyze [file]",
	Short: "Analyze code and find issues",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filePath := args[0]

		// Read file content
		content, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Printf("Error reading file: %v\n", err)
			return
		}

		// Load configuration
		cfg, err := config.Load()
		if err != nil {
			fmt.Printf("Error loading config: %v\n", err)
			return
		}

		if cfg.APIKey == "" {
			fmt.Println("Error: API key not configured")
			return
		}

		// Create API client
		client := api.NewClient(cfg.APIKey)
		client.BaseURL = cfg.BaseURL

		// Prepare analysis prompt
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

		fmt.Printf("Analyzing %s...\n", filePath)
		resp, err := client.SendChat(req)
		if err != nil {
			fmt.Printf("Error analyzing code: %v\n", err)
			return
		}

		if len(resp.Choices) > 0 {
			fmt.Println("\n=== Code Analysis ===")
			fmt.Println(resp.Choices[0].Message.Content)
		}
	},
}

var fixCmd = &cobra.Command{
	Use:   "fix [file]",
	Short: "Fix code issues with AI",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filePath := args[0]

		// Read file content
		content, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Printf("Error reading file: %v\n", err)
			return
		}

		// Load configuration
		cfg, err := config.Load()
		if err != nil {
			fmt.Printf("Error loading config: %v\n", err)
			return
		}

		if cfg.APIKey == "" {
			fmt.Println("Error: API key not configured")
			return
		}

		// Create API client
		client := api.NewClient(cfg.APIKey)
		client.BaseURL = cfg.BaseURL

		// Prepare fix prompt
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

		fmt.Printf("Fixing issues in %s...\n", filePath)
		resp, err := client.SendChat(req)
		if err != nil {
			fmt.Printf("Error fixing code: %v\n", err)
			return
		}

		if len(resp.Choices) > 0 {
			fmt.Println("\n=== Fixed Code ===")
			fmt.Println(resp.Choices[0].Message.Content)

			// Ask if user wants to apply the fix
			fmt.Print("\nDo you want to apply these changes? (y/n): ")
			var response string
			if _, err := fmt.Scanln(&response); err != nil {
				fmt.Printf("Error reading input: %v\n", err)
				return
			}

			if strings.ToLower(response) == "y" || strings.ToLower(response) == "yes" {
				// Extract code from response (simple approach)
				responseContent := resp.Choices[0].Message.Content
				if strings.Contains(responseContent, "```") {
					// Extract code between code blocks
					start := strings.Index(responseContent, "```")
					if start != -1 {
						start = strings.Index(responseContent[start:], "\n") + start + 1
						end := strings.LastIndex(responseContent, "```")
						if end > start {
							code := responseContent[start:end]
							err := os.WriteFile(filePath, []byte(code), 0600)
							if err != nil {
								fmt.Printf("Error writing fixed code: %v\n", err)
								return
							}
							fmt.Printf("Successfully applied fixes to %s\n", filePath)
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
		filePath := args[0]

		// Read file content
		content, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Printf("Error reading file: %v\n", err)
			return
		}

		// Load configuration
		cfg, err := config.Load()
		if err != nil {
			fmt.Printf("Error loading config: %v\n", err)
			return
		}

		if cfg.APIKey == "" {
			fmt.Println("Error: API key not configured")
			return
		}

		// Create API client
		client := api.NewClient(cfg.APIKey)
		client.BaseURL = cfg.BaseURL

		// Prepare review prompt
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

		fmt.Printf("Reviewing %s...\n", filePath)
		resp, err := client.SendChat(req)
		if err != nil {
			fmt.Printf("Error reviewing code: %v\n", err)
			return
		}

		if len(resp.Choices) > 0 {
			fmt.Println("\n=== Code Review ===")
			fmt.Println(resp.Choices[0].Message.Content)
		}
	},
}

func init() {
	codeCmd.AddCommand(analyzeCmd)
	codeCmd.AddCommand(fixCmd)
	codeCmd.AddCommand(reviewCmd)
}
