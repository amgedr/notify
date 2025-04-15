package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml"
)

// Config holds application configuration
type Config struct {
	Telegram TelegramConfig
}

// TelegramConfig holds Telegram-specific configuration
type TelegramConfig struct {
	BotToken string `toml:"bot_token"`
	ChatID   string `toml:"chat_id"`
}

// Load reads configuration from files or environment variables
func Load(verbose bool) (*Config, error) {
	// Default config with empty values
	config := &Config{
		Telegram: TelegramConfig{
			BotToken: "",
			ChatID:   "",
		},
	}

	// Try to load from home directory config file
	homeDir, err := os.UserHomeDir()
	if err == nil {
		homePath := filepath.Join(homeDir, ".notify.conf")
		loaded, err := loadFromFile(homePath, config)
		if err != nil && verbose {
			fmt.Printf("Could not load config from %s: %v\n", homePath, err)
		}
		if loaded && verbose {
			fmt.Printf("Loaded configuration from %s\n", homePath)
		}
	}

	// Try to load from system config file if values are still empty
	if config.Telegram.BotToken == "" || config.Telegram.ChatID == "" {
		sysPath := "/etc/notify.conf"
		loaded, err := loadFromFile(sysPath, config)
		if err != nil && verbose {
			fmt.Printf("Could not load config from %s: %v\n", sysPath, err)
		}
		if loaded && verbose {
			fmt.Printf("Loaded configuration from %s\n", sysPath)
		}
	}

	// Environment variables override file configs
	if envToken := os.Getenv("TELEGRAM_BOT_TOKEN"); envToken != "" {
		if verbose {
			fmt.Println("Using bot token from environment variable")
		}
		config.Telegram.BotToken = envToken
	}

	if envChatID := os.Getenv("TELEGRAM_CHAT_ID"); envChatID != "" {
		if verbose {
			fmt.Println("Using chat ID from environment variable")
		}
		config.Telegram.ChatID = envChatID
	}

	return config, nil
}

// loadFromFile attempts to load configuration from a TOML file
func loadFromFile(path string, config *Config) (bool, error) {
	file, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	defer file.Close()

	decoder := toml.NewDecoder(file)
	if err := decoder.Decode(config); err != nil {
		return false, fmt.Errorf("failed to parse TOML file: %w", err)
	}

	return true, nil
}
