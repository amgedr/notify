package cli

import (
	"flag"
	"fmt"

	"github.com/username/notify/internal/config"
	"github.com/username/notify/internal/telegram"
)

// Flags represents command-line flags
type Flags struct {
	Test    bool
	Verbose bool
}

// ParseFlags parses command line flags and returns them along with non-flag arguments
func ParseFlags() (Flags, []string, error) {
	flags := Flags{}
	
	// Define the flags
	flag.BoolVar(&flags.Test, "test", false, "Test the Telegram connection without sending a message")
	flag.BoolVar(&flags.Verbose, "verbose", false, "Enable verbose output for debugging")
	
	// Parse the flags
	flag.Parse()
	
	// Return the flags and the remaining non-flag arguments
	return flags, flag.Args(), nil
}

// TestConnection tests the Telegram connection using the provided configuration
func TestConnection(cfg *config.Config) error {
	client := telegram.New(cfg.Telegram.BotToken, cfg.Telegram.ChatID, true)
	return client.TestConnection()
}

// SendMessage sends a message using the provided configuration
func SendMessage(cfg *config.Config, message string, verbose bool) error {
	client := telegram.New(cfg.Telegram.BotToken, cfg.Telegram.ChatID, verbose)
	
	// Send the message
	if err := client.SendMessage(message); err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}
	
	return nil
}
