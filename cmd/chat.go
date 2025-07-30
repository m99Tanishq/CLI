package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/m99Tanishq/CLI/internal/api"
	"github.com/m99Tanishq/CLI/internal/config"
	"github.com/m99Tanishq/CLI/pkg/utils"
	"github.com/spf13/cobra"
)

var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: "Start a chat session with GLM",
	Long:  `Start an interactive chat session with the GLM model.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Load configuration
		cfg, err := config.Load()
		if err != nil {
			fmt.Printf("Error loading config: %v\n", err)
			return
		}

		// Check if API key is set
		if cfg.APIKey == "" {
			fmt.Println("Error: API key not configured. Please set it using:")
			fmt.Println("  CLI config --set api_key YOUR_API_KEY")
			return
		}

		// Get model from flag or config
		model, _ := cmd.Flags().GetString("model")
		if model == "" {
			model = cfg.Model
		}

		// Validate model
		if !utils.IsValidModel(model) {
			fmt.Printf("Warning: Unknown model '%s'. Using default model.\n", model)
			model = "zai-org/GLM-4.5:novita"
		}

		// Create API client
		client := api.NewClient(cfg.APIKey)
		client.BaseURL = cfg.BaseURL

		fmt.Printf("Starting chat session with model: %s\n", model)
		fmt.Println("Type 'quit' to exit")
		fmt.Println("----------------------------------------")

		var messages []api.Message
		scanner := bufio.NewScanner(os.Stdin)

		for {
			fmt.Print("You: ")
			if !scanner.Scan() {
				break
			}

			input := strings.TrimSpace(scanner.Text())
			if input == "quit" || input == "exit" {
				fmt.Println("Goodbye!")
				break
			}

			if input == "" {
				continue
			}

			// Add user message to conversation
			userMessage := api.Message{
				Role:    "user",
				Content: input,
			}
			messages = append(messages, userMessage)

			// Send request to GLM API
			fmt.Print("GLM: ")
			start := time.Now()

			req := api.ChatRequest{
				Model:    model,
				Messages: messages,
			}

			resp, err := client.SendChat(req)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				continue
			}

			// Display response
			if len(resp.Choices) > 0 {
				assistantMessage := resp.Choices[0].Message
				fmt.Println(assistantMessage.Content)

				// Add assistant message to conversation
				messages = append(messages, api.Message{
					Role:    "assistant",
					Content: assistantMessage.Content,
				})
			}

			duration := time.Since(start)
			fmt.Printf("(Response time: %s)\n", utils.FormatDuration(duration))
			fmt.Println()
		}
	},
}

func init() {
	chatCmd.Flags().StringP("model", "m", "zai-org/GLM-4.5:novita", "Model to use for chat")
	chatCmd.Flags().BoolP("stream", "s", false, "Enable streaming responses")
}
