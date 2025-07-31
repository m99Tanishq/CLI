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
	Short: "Start a chat session with AI",
	Long:  `Start an interactive chat session with the AI model.`,
	Run: func(cmd *cobra.Command, args []string) {
		ui := utils.NewModernUI()

		cfg, err := config.Load()
		if err != nil {
			ui.PrintError("Failed to load configuration")
			ui.PrintInfo("Please run: CLI config set api_key YOUR_HUGGING_FACE_API_KEY")
			return
		}

		if cfg.APIKey == "" {
			ui.PrintError("API key not configured")
			ui.PrintInfo("Please run: CLI config set api_key YOUR_HUGGING_FACE_API_KEY")
			return
		}

		model, _ := cmd.Flags().GetString("model")
		if model == "" {
			model = cfg.Model
		}

		streaming, _ := cmd.Flags().GetBool("stream")
		showProgress, _ := cmd.Flags().GetBool("progress")
		showTimer, _ := cmd.Flags().GetBool("timer")

		client := api.NewClient(cfg.APIKey, cfg.BaseURL)

		ui.PrintHeader("AI Chat Session")
		ui.PrintModelInfo(model, "Hugging Face")
		ui.PrintInfo(fmt.Sprintf("Model: %s", model))

		if streaming {
			ui.PrintSuccess("Streaming mode enabled - responses will appear in real-time")
		}

		ui.PrintInfo("Type 'quit' to exit")
		ui.PrintWarning("Press Ctrl+C to cancel ongoing responses")
		ui.PrintDivider()

		var messages []api.Message
		reader := bufio.NewReader(os.Stdin)

		for {
			fmt.Println()
			ui.PrintPrompt("You: ")
			os.Stdout.Sync()

			input, err := reader.ReadString('\n')
			if err != nil {
				break
			}
			input = strings.TrimSpace(input)
			if input == "quit" || input == "exit" {
				ui.PrintSuccess("Goodbye! 👋")
				break
			}
			if input == "" {
				continue
			}

			userMessage := api.Message{
				Role:    "user",
				Content: input,
			}
			messages = append(messages, userMessage)

			start := time.Now()
			req := api.ChatRequest{
				Model:       model,
				Messages:    messages,
				MaxTokens:   cfg.MaxTokens,
				Temperature: cfg.Temperature,
			}

			var assistantMessage api.Message
			if streaming {
				assistantMessage, err = handleStreamingChat(client, req, showProgress, showTimer, ui)
			} else {
				assistantMessage, err = handleRegularChat(client, req, start, ui)
			}

			if err != nil {
				ui.PrintError(fmt.Sprintf("Error: %v", err))
				continue
			}
			messages = append(messages, assistantMessage)
		}
	},
}

func handleStreamingChat(client *api.Client, req api.ChatRequest, showProgress, showTimer bool, ui *utils.ModernUI) (api.Message, error) {
	handler := utils.NewStreamingHandler(showProgress, showTimer)
	defer handler.Cleanup()

	chunkChan, errChan := client.SendChatStreamWithChannel(req)
	response, err := handler.HandleStream(chunkChan, errChan)
	if err != nil {
		return api.Message{}, err
	}

	return api.Message{
		Role:    "assistant",
		Content: response,
	}, nil
}

func handleRegularChat(client *api.Client, req api.ChatRequest, start time.Time, ui *utils.ModernUI) (api.Message, error) {
	ui.PrintLoading("Thinking...")

	resp, err := client.SendChat(req)
	if err != nil {
		return api.Message{}, err
	}

	if len(resp.Choices) > 0 {
		assistantMessage := resp.Choices[0].Message
		ui.PrintChatMessage("assistant", assistantMessage.Content)

		duration := time.Since(start)
		ui.PrintStreamingComplete(duration)
		fmt.Println()

		return assistantMessage, nil
	}

	return api.Message{}, fmt.Errorf("no response received")
}

func init() {
	chatCmd.Flags().StringP("model", "m", "", "Model to use for chat (defaults to config)")
	chatCmd.Flags().BoolP("stream", "s", false, "Enable streaming responses")
	chatCmd.Flags().BoolP("progress", "p", true, "Show progress indicators")
	chatCmd.Flags().BoolP("timer", "t", true, "Show response time")
}
