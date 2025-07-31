package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// Config represents the application configuration
type Config struct {
	Model       string  `json:"model"`
	APIKey      string  `json:"api_key"`
	BaseURL     string  `json:"base_url"`
	MaxHistory  int     `json:"max_history"`
	MaxTokens   int     `json:"max_tokens"`
	Temperature float64 `json:"temperature"`
}

// DefaultConfig returns the default configuration
func DefaultConfig() *Config {
	return &Config{
		Model:       "microsoft/DialoGPT-medium",
		APIKey:      "",
		BaseURL:     "https://router.huggingface.co/v1/chat/completions",
		MaxHistory:  100,
		MaxTokens:   1000,
		Temperature: 0.7,
	}
}

// Load loads the configuration from file
func Load() (*Config, error) {
	configPath, err := getConfigPath()
	if err != nil {
		return nil, err
	}

	// If config file doesn't exist, create default config
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		config := DefaultConfig()
		if err := Save(config); err != nil {
			return nil, err
		}
		return config, nil
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	// Set defaults for missing fields
	if config.Model == "" {
		config.Model = "microsoft/DialoGPT-medium"
	}
	if config.MaxHistory == 0 {
		config.MaxHistory = 100
	}
	if config.MaxTokens == 0 {
		config.MaxTokens = 1000
	}
	if config.Temperature == 0 {
		config.Temperature = 0.7
	}
	if config.BaseURL == "" {
		config.BaseURL = "https://router.huggingface.co/v1/chat/completions"
	}

	return &config, nil
}

// Save saves the configuration to file
func Save(config *Config) error {
	configPath, err := getConfigPath()
	if err != nil {
		return err
	}

	// Create config directory if it doesn't exist
	configDir := filepath.Dir(configPath)
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}

	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	if err := os.WriteFile(configPath, data, 0600); err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	return nil
}

// getConfigPath returns the path to the configuration file
func getConfigPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get home directory: %w", err)
	}

	return filepath.Join(homeDir, ".glm-cli", "config.json"), nil
}

// Set sets a configuration value
func Set(key, value string) error {
	config, err := Load()
	if err != nil {
		return err
	}

	switch key {
	case "model":
		config.Model = value
	case "api_key":
		config.APIKey = value
	case "base_url":
		config.BaseURL = value
	case "max_tokens":
		// Parse as integer
		var maxTokens int
		if _, err := fmt.Sscanf(value, "%d", &maxTokens); err != nil {
			return fmt.Errorf("max_tokens must be a number: %w", err)
		}
		config.MaxTokens = maxTokens
	case "temperature":
		// Parse as float
		var temp float64
		if _, err := fmt.Sscanf(value, "%f", &temp); err != nil {
			return fmt.Errorf("temperature must be a number: %w", err)
		}
		config.Temperature = temp
	default:
		return fmt.Errorf("unknown config key: %s", key)
	}

	return Save(config)
}

// Get gets a configuration value
func Get(key string) (string, error) {
	config, err := Load()
	if err != nil {
		return "", err
	}

	switch key {
	case "model":
		return config.Model, nil
	case "api_key":
		return config.APIKey, nil
	case "base_url":
		return config.BaseURL, nil
	case "max_tokens":
		return fmt.Sprintf("%d", config.MaxTokens), nil
	case "temperature":
		return fmt.Sprintf("%.2f", config.Temperature), nil
	default:
		return "", fmt.Errorf("unknown config key: %s", key)
	}
}
