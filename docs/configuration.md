# Configuration

The `notify` tool supports multiple configuration methods with a clear precedence order.

## Configuration Files

The tool looks for TOML configuration files in the following order:

1. `~/.notify.conf` (user's home directory)
2. `/etc/notify.conf` (system-wide)

The configuration file should have the following format:

```toml
[telegram]
bot_token = "YOUR_TELEGRAM_BOT_TOKEN"
chat_id = "YOUR_TELEGRAM_CHAT_ID"
