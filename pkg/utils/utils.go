package utils

import (
	"crypto/rand"
	"fmt"
	"strings"
	"time"
)

// GenerateID generates a random ID for chat sessions
func GenerateID() string {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		// Fallback to timestamp-based ID if crypto/rand fails
		return fmt.Sprintf("%x", time.Now().UnixNano())
	}
	return fmt.Sprintf("%x", b)
}

// TruncateString truncates a string to the specified length
func TruncateString(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen-3] + "..."
}

// FormatTimestamp formats a timestamp for display
func FormatTimestamp(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

// GenerateSessionTitle generates a title for a chat session based on the first message
func GenerateSessionTitle(content string) string {
	// Remove extra whitespace and newlines
	content = strings.TrimSpace(content)
	content = strings.ReplaceAll(content, "\n", " ")

	// Truncate to reasonable length
	if len(content) > 50 {
		content = TruncateString(content, 50)
	}

	if content == "" {
		return "Untitled Session"
	}

	return content
}

// IsValidModel checks if the provided model name is valid
func IsValidModel(model string) bool {
	validModels := []string{
		"glm-4",
		"glm-3-turbo",
		"glm-4v",
		"cogview-3",
		"zai-org/GLM-4.5",
		"zai-org/GLM-4.5:novita",
	}

	for _, validModel := range validModels {
		if model == validModel {
			return true
		}
	}

	return false
}

// SanitizeInput removes potentially dangerous characters from user input
func SanitizeInput(input string) string {
	// Remove control characters except newlines and tabs
	var result strings.Builder
	for _, r := range input {
		if r == '\n' || r == '\t' || (r >= 32 && r != 127) {
			result.WriteRune(r)
		}
	}
	return result.String()
}

// FormatDuration formats a duration in a human-readable way
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
