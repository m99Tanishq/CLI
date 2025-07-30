package models

import (
	"time"
)

// ChatRequest represents a request to the GLM API
type ChatRequest struct {
	Model       string    `json:"model"`
	Messages    []Message `json:"messages"`
	Stream      bool      `json:"stream,omitempty"`
	MaxTokens   int       `json:"max_tokens,omitempty"`
	Temperature float64   `json:"temperature,omitempty"`
}

// ChatResponse represents a response from the GLM API
type ChatResponse struct {
	ID      string   `json:"id"`
	Object  string   `json:"object"`
	Created int64    `json:"created"`
	Model   string   `json:"model"`
	Choices []Choice `json:"choices"`
	Usage   Usage    `json:"usage"`
}

// Choice represents a choice in the API response
type Choice struct {
	Index        int     `json:"index"`
	Message      Message `json:"message"`
	FinishReason string  `json:"finish_reason"`
}

// Message represents a chat message
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// Usage represents token usage information
type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

// ChatSession represents a chat session
type ChatSession struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Messages  []Message `json:"messages"`
	Model     string    `json:"model"`
}

// Config represents application configuration
type Config struct {
	APIKey      string  `json:"api_key" mapstructure:"api_key"`
	Model       string  `json:"model" mapstructure:"model"`
	BaseURL     string  `json:"base_url" mapstructure:"base_url"`
	MaxHistory  int     `json:"max_history" mapstructure:"max_history"`
	MaxTokens   int     `json:"max_tokens" mapstructure:"max_tokens"`
	Temperature float64 `json:"temperature" mapstructure:"temperature"`
}

// ChatHistory represents a collection of chat sessions
type ChatHistory struct {
	Sessions []ChatSession `json:"sessions"`
	Total    int           `json:"total"`
}

// ErrorResponse represents an error response from the API
type ErrorResponse struct {
	Error struct {
		Message string `json:"message"`
		Type    string `json:"type"`
		Code    string `json:"code"`
	} `json:"error"`
}

// StreamResponse represents a streaming response chunk
type StreamResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int64  `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Index int `json:"index"`
		Delta struct {
			Content string `json:"content"`
		} `json:"delta"`
		FinishReason *string `json:"finish_reason"`
	} `json:"choices"`
}
