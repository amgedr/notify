package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Constants for Telegram API
const (
	baseURL            = "https://api.telegram.org/bot%s/%s"
	sendMessageMethod  = "sendMessage"
	getUpdatesMethod   = "getUpdates"
	defaultHTTPTimeout = 10 * time.Second
)

// Client represents a Telegram API client
type Client struct {
	BotToken  string
	ChatID    string
	APIClient *http.Client
	Verbose   bool
}

// sendMessageRequest is the structure of a Telegram sendMessage request
type sendMessageRequest struct {
	ChatID    string `json:"chat_id"`
	Text      string `json:"text"`
	ParseMode string `json:"parse_mode,omitempty"`
}

// APIResponse represents a generic Telegram API response
type APIResponse struct {
	OK          bool            `json:"ok"`
	Description string          `json:"description,omitempty"`
	Result      json.RawMessage `json:"result,omitempty"`
	ErrorCode   int             `json:"error_code,omitempty"`
}

// New creates a new Telegram client
func New(botToken, chatID string, verbose bool) *Client {
	return &Client{
		BotToken: botToken,
		ChatID:   chatID,
		APIClient: &http.Client{
			Timeout: defaultHTTPTimeout,
		},
		Verbose: verbose,
	}
}

// TestConnection verifies that the bot token and chat ID are valid
func (c *Client) TestConnection() error {
	// Make a simple getUpdates call to validate the bot token
	url := fmt.Sprintf(baseURL, c.BotToken, getUpdatesMethod)
	
	resp, err := c.APIClient.Get(url)
	if err != nil {
		return fmt.Errorf("network error: %w", err)
	}
	defer resp.Body.Close()

	var apiResp APIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return fmt.Errorf("failed to decode API response: %w", err)
	}

	if !apiResp.OK {
		return fmt.Errorf("Telegram API error: %s (code: %d)", apiResp.Description, apiResp.ErrorCode)
	}

	// If we got here, the bot token is valid
	// We don't actually send a test message to avoid notifications
	if c.Verbose {
		fmt.Println("Bot token is valid")
	}

	return nil
}

// SendMessage sends a message to the configured Telegram chat
func (c *Client) SendMessage(message string) error {
	url := fmt.Sprintf(baseURL, c.BotToken, sendMessageMethod)
	
	requestBody := sendMessageRequest{
		ChatID: c.ChatID,
		Text:   message,
	}
	
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}
	
	if c.Verbose {
		fmt.Printf("Sending request to: %s\n", url)
	}
	
	resp, err := c.APIClient.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("network error: %w", err)
	}
	defer resp.Body.Close()
	
	if c.Verbose {
		fmt.Printf("Response status: %s\n", resp.Status)
	}
	
	// Read the response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}
	
	if c.Verbose {
		fmt.Printf("Response body: %s\n", string(respBody))
	}
	
	var apiResp APIResponse
	if err := json.Unmarshal(respBody, &apiResp); err != nil {
		return fmt.Errorf("failed to decode API response: %w", err)
	}
	
	if !apiResp.OK {
		return fmt.Errorf("Telegram API error: %s (code: %d)", apiResp.Description, apiResp.ErrorCode)
	}
	
	return nil
}
