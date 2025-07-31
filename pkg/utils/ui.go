package utils

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
)

// ModernUI provides a professional UI system inspired by Qwen3-Coder
type ModernUI struct {
	theme *Theme
}

// Theme represents a modern color theme
type Theme struct {
	Primary    *color.Color
	Secondary  *color.Color
	Success    *color.Color
	Warning    *color.Color
	Error      *color.Color
	Info       *color.Color
	Muted      *color.Color
	Background *color.Color
	Accent     *color.Color
	Code       *color.Color
	Prompt     *color.Color
	Response   *color.Color
}

// NewModernUI creates a new modern UI instance
func NewModernUI() *ModernUI {
	return &ModernUI{
		theme: &Theme{
			Primary:    color.New(color.FgHiCyan, color.Bold),
			Secondary:  color.New(color.FgHiBlue),
			Success:    color.New(color.FgHiGreen, color.Bold),
			Warning:    color.New(color.FgHiYellow, color.Bold),
			Error:      color.New(color.FgHiRed, color.Bold),
			Info:       color.New(color.FgHiMagenta),
			Muted:      color.New(color.FgHiBlack),
			Background: color.New(color.BgHiBlack),
			Accent:     color.New(color.FgHiCyan),
			Code:       color.New(color.FgHiYellow),
			Prompt:     color.New(color.FgHiCyan, color.Bold),
			Response:   color.New(color.FgHiWhite),
		},
	}
}

// PrintHeader prints a modern header
func (ui *ModernUI) PrintHeader(title string) {
	fmt.Println()
	ui.theme.Primary.Printf("╭─ %s ─", strings.Repeat("─", len(title)+2))
	fmt.Println()
	ui.theme.Primary.Printf("│ %s", title)
	fmt.Println()
	ui.theme.Primary.Printf("╰─%s─", strings.Repeat("─", len(title)+4))
	fmt.Println()
}

// PrintSection prints a section header
func (ui *ModernUI) PrintSection(title string) {
	fmt.Println()
	ui.theme.Secondary.Printf("▸ %s", title)
	fmt.Println()
}

// PrintFeature prints a feature with icon
func (ui *ModernUI) PrintFeature(icon, text string) {
	ui.theme.Success.Printf("  %s %s\n", icon, text)
}

// PrintCodeBlock prints a code block with syntax highlighting
func (ui *ModernUI) PrintCodeBlock(code, language string) {
	fmt.Println()
	ui.theme.Muted.Printf("```%s\n", language)
	ui.theme.Code.Print(code)
	fmt.Println()
	ui.theme.Muted.Print("```")
	fmt.Println()
}

// PrintCommand prints a command with styling
func (ui *ModernUI) PrintCommand(cmd string) {
	ui.theme.Prompt.Printf("$ %s\n", cmd)
}

// PrintPrompt prints a prompt
func (ui *ModernUI) PrintPrompt(prompt string) {
	ui.theme.Prompt.Print(prompt)
	// Force flush to ensure prompt is displayed immediately
	os.Stdout.Sync()
}

// PrintOutput prints command output
func (ui *ModernUI) PrintOutput(output string) {
	ui.theme.Response.Print(output)
}

// PrintStatus prints a status message
func (ui *ModernUI) PrintStatus(icon, message string) {
	ui.theme.Info.Printf("%s %s\n", icon, message)
}

// PrintProgress prints a modern progress bar
func (ui *ModernUI) PrintProgress(current, total int, message string) {
	percentage := float64(current) / float64(total)
	width := 30
	filled := int(float64(width) * percentage)

	bar := ""
	for i := 0; i < width; i++ {
		if i < filled {
			bar += "█"
		} else {
			bar += "░"
		}
	}

	ui.theme.Accent.Printf("\r[%s] %d/%d (%d%%) %s", bar, current, total, int(percentage*100), message)
}

// PrintSpinner prints a modern spinner
func (ui *ModernUI) PrintSpinner(frame, message string) {
	ui.theme.Accent.Printf("\r%s %s", frame, message)
}

// PrintTable prints a modern table
func (ui *ModernUI) PrintTable(headers []string, rows [][]string) {
	if len(rows) == 0 {
		ui.PrintInfo("No items found")
		return
	}

	// Simple table for now
	fmt.Println()
	ui.theme.Primary.Printf("┌─ %s ─┐\n", strings.Join(headers, " ─┬─ "))

	for _, row := range rows {
		ui.theme.Response.Printf("│ %s │\n", strings.Join(row, " │ "))
	}

	ui.theme.Primary.Printf("└─%s─┘\n", strings.Repeat("───┴─", len(headers)-1)+"───")
	fmt.Println()
}

// PrintCard prints a modern card layout
func (ui *ModernUI) PrintCard(title, content string) {
	fmt.Println()
	ui.theme.Primary.Printf("┌─ %s ─", strings.Repeat("─", len(title)))
	fmt.Println()
	ui.theme.Primary.Printf("│ %s", title)
	fmt.Println()
	ui.theme.Primary.Print("├─")
	fmt.Println()
	ui.theme.Response.Print(content)
	fmt.Println()
	ui.theme.Primary.Print("└─")
	fmt.Println()
}

// PrintList prints a modern list
func (ui *ModernUI) PrintList(items []string, icon string) {
	for _, item := range items {
		ui.theme.Info.Printf("  %s %s\n", icon, item)
	}
}

// PrintDivider prints a modern divider
func (ui *ModernUI) PrintDivider() {
	ui.theme.Muted.Println("─" + strings.Repeat("─", 60) + "─")
}

// PrintBanner prints a modern banner
func (ui *ModernUI) PrintBanner(text string) {
	fmt.Println()
	ui.theme.Primary.Printf("╔═ %s ═", strings.Repeat("═", len(text)+2))
	fmt.Println()
	ui.theme.Primary.Printf("║ %s", text)
	fmt.Println()
	ui.theme.Primary.Printf("╚═%s═", strings.Repeat("═", len(text)+4))
	fmt.Println()
}

// PrintAlert prints a modern alert
func (ui *ModernUI) PrintAlert(level, message string) {
	var icon *color.Color

	switch level {
	case "success":
		icon = ui.theme.Success
	case "warning":
		icon = ui.theme.Warning
	case "error":
		icon = ui.theme.Error
	case "info":
		icon = ui.theme.Info
	default:
		icon = ui.theme.Info
	}

	icon.Printf("● %s\n", message)
}

// PrintChatMessage prints a modern chat message
func (ui *ModernUI) PrintChatMessage(role, content string) {
	if role == "user" {
		ui.theme.Prompt.Printf("👤 You: %s\n", content)
	} else {
		ui.theme.Response.Printf("🤖 AI: %s\n", content)
	}
}

// PrintStreamingMessage prints a streaming message
func (ui *ModernUI) PrintStreamingMessage(content string) {
	ui.theme.Response.Print(content)
	// Force flush to ensure immediate display
	os.Stdout.Sync()
}

// PrintStreamingComplete prints streaming completion
func (ui *ModernUI) PrintStreamingComplete(duration time.Duration) {
	ui.theme.Success.Printf("✨ Response complete in %s\n", FormatDuration(duration))
}

// PrintStreamingCancelled prints streaming cancellation
func (ui *ModernUI) PrintStreamingCancelled() {
	fmt.Println()
	ui.theme.Warning.Println("⏹️  Response cancelled by user")
}

// PrintLoading prints a loading message
func (ui *ModernUI) PrintLoading(message string) {
	ui.theme.Info.Printf("⏳ %s\n", message)
}

// PrintSuccess prints a success message
func (ui *ModernUI) PrintSuccess(message string) {
	ui.theme.Success.Printf("✅ %s\n", message)
}

// PrintError prints an error message
func (ui *ModernUI) PrintError(message string) {
	ui.theme.Error.Printf("❌ %s\n", message)
}

// PrintWarning prints a warning message
func (ui *ModernUI) PrintWarning(message string) {
	ui.theme.Warning.Printf("⚠️  %s\n", message)
}

// PrintInfo prints an info message
func (ui *ModernUI) PrintInfo(message string) {
	ui.theme.Info.Printf("ℹ️  %s\n", message)
}

// PrintModelInfo prints model information
func (ui *ModernUI) PrintModelInfo(model, version string) {
	ui.theme.Primary.Printf("🤖 Model: %s (v%s)\n", model, version)
}

// PrintCapabilities prints model capabilities
func (ui *ModernUI) PrintCapabilities(capabilities []string) {
	ui.PrintSection("Capabilities")
	for _, capability := range capabilities {
		ui.PrintFeature("✨", capability)
	}
}

// PrintUsage prints usage information
func (ui *ModernUI) PrintUsage(usage string) {
	ui.PrintSection("Usage")
	ui.theme.Response.Println(usage)
}

// PrintExamples prints usage examples
func (ui *ModernUI) PrintExamples(examples []string) {
	ui.PrintSection("Examples")
	for i, example := range examples {
		ui.theme.Code.Printf("  %d. %s\n", i+1, example)
	}
}

// PrintFooter prints a modern footer
func (ui *ModernUI) PrintFooter() {
	fmt.Println()
	ui.theme.Muted.Println("Made with ❤️  by m99tanq")
	ui.theme.Muted.Println("Powered by Open Source LLM models")
}
