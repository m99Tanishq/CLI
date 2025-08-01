package utils

import (
	"crypto/rand"
	"fmt"
	"regexp"
	"strings"
	"time"
)

func GenerateID() string {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		return fmt.Sprintf("%x", time.Now().UnixNano())
	}
	return fmt.Sprintf("%x", b)
}

func TruncateString(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen-3] + "..."
}

func FormatTimestamp(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func GenerateSessionTitle(content string) string {
	content = strings.TrimSpace(content)
	content = strings.ReplaceAll(content, "\n", " ")

	if len(content) > 50 {
		content = TruncateString(content, 50)
	}

	if content == "" {
		return "Untitled Session"
	}

	return content
}

func IsValidModel(model string) bool {
	validModels := map[string]bool{
		"Rzork-4":       true,
		"Rzork-3-turbo": true,
		"Rzork-4v":      true,
		"cogview-3":     true,
	}
	return validModels[model]
}

func SanitizeInput(input string) string {
	var result strings.Builder
	for _, r := range input {
		if r == '\n' || r == '\t' || (r >= 32 && r != 127) {
			result.WriteRune(r)
		}
	}
	return result.String()
}

func FormatDuration(d time.Duration) string {
	if d < time.Second {
		return fmt.Sprintf("%dms", d.Milliseconds())
	}
	if d < time.Minute {
		return fmt.Sprintf("%.1fs", d.Seconds())
	}
	if d < time.Hour {
		return fmt.Sprintf("%dm %ds", int(d.Minutes()), int(d.Seconds())%60)
	}
	return fmt.Sprintf("%dh %dm", int(d.Hours()), int(d.Minutes())%60)
}

// CleanResponse removes think tags and other unwanted content from AI responses
func CleanResponse(content string) string {
	// Remove <think>...</think> tags and their content
	thinkPattern := regexp.MustCompile(`(?s)<think>.*?</think>`)
	content = thinkPattern.ReplaceAllString(content, "")

	// Remove any remaining think tags without content
	content = strings.ReplaceAll(content, "<think>", "")
	content = strings.ReplaceAll(content, "</think>", "")

	// Clean up extra whitespace but preserve spaces between words
	content = strings.TrimSpace(content)

	// Remove multiple consecutive newlines
	newlinePattern := regexp.MustCompile(`\n\s*\n\s*\n`)
	content = newlinePattern.ReplaceAllString(content, "\n\n")

	return content
}
