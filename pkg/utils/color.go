package utils

import (
	"strings"

	"github.com/fatih/color"
)

// Color themes for consistent styling
var (
	// Primary colors
	Primary   = color.New(color.FgCyan, color.Bold)
	Secondary = color.New(color.FgBlue)
	Success   = color.New(color.FgGreen, color.Bold)
	Warning   = color.New(color.FgYellow, color.Bold)
	Error     = color.New(color.FgRed, color.Bold)
	Info      = color.New(color.FgMagenta)

	// UI elements
	Prompt    = color.New(color.FgCyan, color.Bold)
	Response  = color.New(color.FgWhite)
	Code      = color.New(color.FgYellow)
	Progress  = color.New(color.FgBlue)
	Timer     = color.New(color.FgHiBlack)
	SpinnerUI = color.New(color.FgCyan)

	// Status indicators
	Thinking  = color.New(color.FgCyan)
	Working   = color.New(color.FgBlue)
	Done      = color.New(color.FgGreen)
	Cancelled = color.New(color.FgRed)
)

// PrintSuccess prints a success message
func PrintSuccess(message string) {
	Success.Println("✅ " + message)
}

// PrintError prints an error message
func PrintError(message string) {
	Error.Println("❌ " + message)
}

// PrintWarning prints a warning message
func PrintWarning(message string) {
	Warning.Println("⚠️  " + message)
}

// PrintInfo prints an info message
func PrintInfo(message string) {
	Info.Println("ℹ️  " + message)
}

// PrintPrompt prints a prompt
func PrintPrompt(message string) {
	Prompt.Print(message)
}

// PrintResponse prints a response
func PrintResponse(message string) {
	Response.Print(message)
}

// PrintCode prints code with syntax highlighting
func PrintCode(code string) {
	Code.Print(code)
}

// PrintProgress prints progress information
func PrintProgress(message string) {
	Progress.Print(message)
}

// PrintTimer prints timing information
func PrintTimer(message string) {
	Timer.Print(message)
}

// PrintThinking prints thinking indicator
func PrintThinking(message string) {
	Thinking.Print(message)
}

// PrintWorking prints working indicator
func PrintWorking(message string) {
	Working.Print(message)
}

// PrintDone prints completion message
func PrintDone(message string) {
	Done.Println("✅ " + message)
}

// PrintCancelled prints cancellation message
func PrintCancelled(message string) {
	Cancelled.Println("⏹️  " + message)
}

// PrintStreamingHeader prints the streaming header
func PrintStreamingHeader() {
	Prompt.Print("🤖 LLM: ")
}

// PrintStreamingChunk prints a streaming chunk
func PrintStreamingChunk(chunk string) {
	Response.Print(chunk)
}

// PrintStreamingProgress prints streaming progress
func PrintStreamingProgress() {
	Progress.Print(".")
}

// PrintStreamingTimer prints streaming timer
func PrintStreamingTimer(duration string) {
	Timer.Printf("⏱️  Response time: %s\n", duration)
}

// PrintStreamingCancelled prints streaming cancellation
func PrintStreamingCancelled() {
	Cancelled.Println("\n\n⚠️  Cancelling... Press Ctrl+C again to force quit.")
}

// PrintStreamingComplete prints streaming completion
func PrintStreamingComplete() {
	Done.Println("\n✨ Response complete!")
}

// PrintChatHeader prints chat session header
func PrintChatHeader(model string, streaming bool) {
	Info.Printf("Starting chat session with model: %s\n", model)
	if streaming {
		Success.Println("✨ Streaming mode enabled - responses will appear in real-time")
	}
	Info.Println("Type 'quit' to exit")
	Warning.Println("Press Ctrl+C to cancel ongoing responses")
	Info.Println("----------------------------------------")
}

// PrintUserPrompt prints user prompt
func PrintUserPrompt() {
	Prompt.Print("You: ")
}

// PrintModelResponse prints model response header
func PrintModelResponse() {
	Prompt.Print("🤖 LLM: ")
}

// PrintResponseTime prints response time
func PrintResponseTime(duration string) {
	Timer.Printf("(Response time: %s)\n", duration)
}

// PrintGoodbye prints goodbye message
func PrintGoodbye() {
	Success.Println("👋 Goodbye!")
}

// PrintErrorWithDetails prints error with details
func PrintErrorWithDetails(err error, details string) {
	Error.Printf("❌ Error: %v\n", err)
	if details != "" {
		Info.Printf("💡 %s\n", details)
	}
}

// PrintModelWarning prints model warning
func PrintModelWarning(model string) {
	Warning.Printf("⚠️  Warning: Unknown model '%s'. Using default model.\n", model)
}

// PrintStreamingInfo prints streaming information
func PrintStreamingInfo() {
	Info.Println("🔄 Streaming response...")
}

// PrintStreamingEnd prints streaming end
func PrintStreamingEnd() {
	Done.Println("✨ Streaming complete!")
}

// PrintCancellationInfo prints cancellation information
func PrintCancellationInfo() {
	Warning.Println("⏹️  Operation cancelled by user")
}

// PrintProgressBar prints a progress bar
func PrintProgressBar(current, total int, message string) {
	percentage := float64(current) / float64(total)
	filled := int(percentage * 50)

	bar := ""
	for i := 0; i < 50; i++ {
		if i < filled {
			bar += "█"
		} else {
			bar += "░"
		}
	}

	Progress.Printf("\r[%s] %d/%d (%d%%) %s", bar, current, total, int(percentage*100), message)
}

// PrintSpinner prints a spinner frame
func PrintSpinner(frame string, message string) {
	SpinnerUI.Printf("\r%s %s", frame, message)
}

// PrintTableHeader prints a table header
func PrintTableHeader(headers ...string) {
	for i, header := range headers {
		if i > 0 {
			Info.Print(" | ")
		}
		Primary.Print(header)
	}
	Info.Println()
	Info.Println(strings.Repeat("-", len(headers)*20))
}

// PrintTableRow prints a table row
func PrintTableRow(cells ...string) {
	for i, cell := range cells {
		if i > 0 {
			Info.Print(" | ")
		}
		Response.Print(cell)
	}
	Info.Println()
}

// PrintSeparator prints a separator line
func PrintSeparator() {
	Info.Println("----------------------------------------")
}

// PrintNewLine prints a new line
func PrintNewLine() {
	Info.Println()
}

// PrintBold prints bold text
func PrintBold(text string) {
	Primary.Print(text)
}

// PrintItalic prints italic text
func PrintItalic(text string) {
	Secondary.Print(text)
}

// PrintUnderline prints underlined text
func PrintUnderline(text string) {
	Info.Print(text)
}

// PrintHighlighted prints highlighted text
func PrintHighlighted(text string) {
	Code.Print(text)
}

// PrintMuted prints muted text
func PrintMuted(text string) {
	Timer.Print(text)
}

// PrintEmphasis prints emphasized text
func PrintEmphasis(text string) {
	Success.Print(text)
}

// PrintCaution prints caution text
func PrintCaution(text string) {
	Warning.Print(text)
}

// PrintCritical prints critical text
func PrintCritical(text string) {
	Error.Print(text)
}
