# CI/CD Setup

This document explains how the CI/CD pipeline works for the `notify` tool.

## GitHub Actions Workflow

The project uses GitHub Actions to automatically build, test, and deploy the tool. The workflow is defined in `.github/workflows/build.yml`.

### Workflow Steps

1. **Checkout Code**: Retrieves the latest code from the repository
2. **Set up Go**: Configures the Go environment
3. **Install Dependencies**: Downloads all required Go modules
4. **Run Tests**: Executes all unit tests
5. **Build**: Compiles the binary for Linux (amd64)
6. **Test Binary**: Verifies the binary works by running it with the `--test` flag
7. **Deploy to Server**: Copies the binary to a configured server via SSH

### Required Secrets

To make the CI/CD pipeline work, you need to configure the following secrets in your GitHub repository:

- `TELEGRAM_BOT_TOKEN`: Your Telegram bot token
- `TELEGRAM_CHAT_ID`: Your Telegram chat ID
- `SSH_PRIVATE_KEY`: The private SSH key for accessing your server
- `SSH_HOST`: The hostname or IP address of your server
- `SSH_USERNAME`: The username to use for SSH connections

### Adding Secrets to GitHub

1. Go to your GitHub repository
2. Click on "Settings"
3. Select "Secrets and variables" > "Actions"
4. Click "New repository secret"
5. Add each secret with its appropriate name and value

## Manual Deployment

If you prefer to deploy manually, you can:

1. Build the binary:
   ```bash
   GOOS=linux GOARCH=amd64 go build -o notify
   ```

2. Copy it to your server:
   ```bash
   scp notify user@your-server:~/
   ```

3. Make it executable:
   ```bash
   ssh user@your-server "chmod +x ~/notify && sudo mv ~/notify /usr/local/bin/"
   ```

## Troubleshooting Deployments

If deployments fail, check:

1. SSH key permissions: The private key should have `600` permissions
2. Server firewall: Ensure port 22 is open for SSH connections
3. GitHub Secrets: Verify all required secrets are properly configured
4. Server disk space: Ensure there's enough space to copy the binary
