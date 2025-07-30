package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

// Config represents the application configuration
type Config struct {
	APIKey     string `json:"api_key" mapstructure:"api_key"`
	Model      string `json:"model" mapstructure:"model"`
	BaseURL    string `json:"base_url" mapstructure:"base_url"`
	MaxHistory int    `json:"max_history" mapstructure:"max_history"`
}

// DefaultConfig returns the default configuration
func DefaultConfig() *Config {
	return &Config{
		Model:      "glm-4",
		BaseURL:    "https://open.bigmodel.cn/api/paas/v4",
		MaxHistory: 100,
	}
}

// Load loads configuration from file and environment variables
func Load() (*Config, error) {
	configPath := getConfigPath()

	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(configPath)

	// Set default values
	viper.SetDefault("model", "glm-4")
	viper.SetDefault("base_url", "https://open.bigmodel.cn/api/paas/v4")
	viper.SetDefault("max_history", 100)

	// Read from environment variables
	viper.AutomaticEnv()
	viper.SetEnvPrefix("GLM")
	if err := viper.BindEnv("api_key", "GLM_API_KEY"); err != nil {
		return nil, fmt.Errorf("failed to bind api_key env: %w", err)
	}
	if err := viper.BindEnv("model", "GLM_MODEL"); err != nil {
		return nil, fmt.Errorf("failed to bind model env: %w", err)
	}
	if err := viper.BindEnv("base_url", "GLM_BASE_URL"); err != nil {
		return nil, fmt.Errorf("failed to bind base_url env: %w", err)
	}
	if err := viper.BindEnv("max_history", "GLM_MAX_HISTORY"); err != nil {
		return nil, fmt.Errorf("failed to bind max_history env: %w", err)
	}

	// Read config file if it exists
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("failed to read config file: %w", err)
		}
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &config, nil
}

// Save saves the configuration to file
func Save(config *Config) error {
	configPath := getConfigPath()

	// Create config directory if it doesn't exist
	if err := os.MkdirAll(configPath, 0755); err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}

	configFile := filepath.Join(configPath, "config.json")

	jsonData, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	if err := os.WriteFile(configFile, jsonData, 0600); err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	return nil
}

// getConfigPath returns the configuration directory path
func getConfigPath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		homeDir = "."
	}
	return filepath.Join(homeDir, ".CLI")
}
