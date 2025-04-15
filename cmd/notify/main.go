package notify

import (
	"fmt"
	"os"
	"strings"

	"github.com/username/notify/internal/config"
	"github.com/username/notify/pkg/cli"
)

// Execute runs the notify command
func Execute() {
	// Parse command line flags
	flags, args, err := cli.ParseFlags()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing flags: %v\n", err)
		os.Exit(1)
	}

	// Load configuration
	cfg, err := config.Load(flags.Verbose)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading configuration: %v\n", err)
		os.Exit(1)
	}

	// Validate configuration
	if cfg.Telegram.BotToken == "" || cfg.Telegram.ChatID == "" {
		fmt.Fprintln(os.Stderr, "Error: Telegram bot token and chat ID are required.")
		fmt.Fprintln(os.Stderr, "Please set them in ~/.notify.conf, /etc/notify.conf, or via environment variables.")
		os.Exit(1)
	}

	if flags.Verbose {
		// Don't print full token for security
		tokenPreview := ""
		if len(cfg.Telegram.BotToken) > 8 {
			tokenPreview = cfg.Telegram.BotToken[:8] + "..."
		}
		fmt.Printf("Configuration loaded: Bot token: %s, Chat ID: %s\n", tokenPreview, cfg.Telegram.ChatID)
	}

	// Test mode - verify connection without sending a message
	if flags.Test {
		if flags.Verbose {
			fmt.Println("Test mode enabled. Testing connection to Telegram...")
		}
		
		err := cli.TestConnection(cfg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Connection test failed: %v\n", err)
			os.Exit(1)
		}
		
		fmt.Println("Connection test successful! Configuration is valid.")
		os.Exit(0)
	}

	// Regular message sending mode
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "Error: No message provided. Please provide a message to send.")
		fmt.Fprintln(os.Stderr, "Usage: notify \"your message here\"")
		os.Exit(1)
	}

	// Join all arguments as a single message
	message := strings.Join(args, " ")
	
	// Send the message
	err = cli.SendMessage(cfg, message, flags.Verbose)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to send message: %v\n", err)
		os.Exit(1)
	}

	if flags.Verbose {
		fmt.Println("Message sent successfully!")
	}
}
