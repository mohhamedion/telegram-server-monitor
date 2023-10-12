# Go Telegram Server Monitor

This Go Telegram Server Monitor is a program that monitors your server's CPU usage, memory usage, and available disk space, sending periodic updates via Telegram messages to keep you informed about your server's health and resource usage.

## Features

- Monitors CPU usage.
- Monitors memory usage.
- Monitors available disk space.
- Sends updates via Telegram messages.
- Customizable threshold values for alerts.

## Prerequisites

Before you can use this program, ensure you have the following:

- Go installed on your server.
- A Telegram account and a bot token. You can create a bot and obtain the token by talking to the [BotFather](https://core.telegram.org/bots#botfather) on Telegram.

## Installation

1. Clone this repository to your server:

   ```bash
   git clone https://github.com/your-username/go-telegram-server-monitor.git
   cd go-telegram-server-monitor
   ```

2. Build the Go program:

   ```bash
   go build monitor.go
   ```

3. Create a `config.json` file at the root of the directory with the following content:

```json
{
    "ChatID": 0, // your chat id
    "BotToken": "YOUR_BOT_TOKEN"
}
```

Replace `"YOUR_BOT_TOKEN"` with your Telegram bot's token, and ensure you've provided the correct `ChatID` where you want to receive updates.

4. Adjust the threshold values and monitoring frequency in `config.json` as needed.

## Configuration

Edit the `config.json` file to customize the monitor's behavior:

- `telegram_bot_token`: Replace `"YOUR_BOT_TOKEN"` with your Telegram bot's token.
- `telegram_chat_id`: Replace `-4039470109` with the chat ID where you want to receive updates. You can create a group chat and invite your bot, then send a message to the group to obtain the chat ID.
- `cpu_threshold`: Set the CPU usage threshold in percentage. If the CPU usage exceeds this value, you will receive an alert.
- `memory_threshold`: Set the memory usage threshold in percentage. If memory usage exceeds this value, you will receive an alert.
- `disk_threshold`: Set the minimum available disk space threshold in GB. If the available disk space falls below this value, you will receive an alert.
- `check_interval_seconds`: Set the monitoring interval in seconds. The program will check the server's resources at this frequency.

## Usage

Run the program with the following command:

```bash
./monitor
```

The program will start monitoring your server's resources and send Telegram messages when any of the thresholds are exceeded.

## Troubleshooting

If you encounter any issues or have questions, please feel free to open an issue on the [GitHub repository](https://github.com/your-username/go-telegram-server-monitor).

## License

This Go Telegram Server Monitor is open-source and distributed under the MIT License. See the [LICENSE](LICENSE) file for details.

**Note**: Ensure you have the necessary permissions to run the program and access server resources before using this tool.