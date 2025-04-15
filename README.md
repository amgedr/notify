# Notify

A simple command-line tool that sends notifications to Telegram.

## Installation

### From Source

1. Clone the repository:
   ```bash
   git clone https://github.com/username/notify.git
   cd notify
   ```

2. Build the binary:
   ```bash
   go build -o notify
   ```

3. Move the binary to a location in your PATH:
   ```bash
   sudo mv notify /usr/local/bin/
   ```

### From Releases

Download the latest release from the [Releases page](https://github.com/username/notify/releases).

## Configuration

The tool can be configured in two ways:

1. **TOML Configuration File**:
   Create a file at `~/.notify.conf` or `/etc/notify.conf` with the following content:

   ```toml
   [telegram]
   bot_token = "YOUR_TELEGRAM_BOT_TOKEN"
   chat_id = "YOUR_TELEGRAM_CHAT_ID"
   ```

2. **Environment Variables**:
   Set the following environment variables:

   ```bash
   export TELEGRAM_BOT_TOKEN="YOUR_TELEGRAM_BOT_TOKEN"
   export TELEGRAM_CHAT_ID="YOUR_TELEGRAM_CHAT_ID"
   ```

Environment variables will override values from the configuration files.

## Usage

### Sending a Simple Message

```bash
notify "Your message here"
