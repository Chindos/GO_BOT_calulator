# GO_BOT_calulator

This project is a simple Telegram bot written in Go that acts as a calculator. The bot accepts messages from users containing simple arithmetic expressions, calculates them and returns the result.

## Features

- **Simplicity:**Supports only simple expressions with two numbers and one operator (e.g. `3+4`, `10-2`, `5*6`, `8/2`).
- **Manual parsing:** The expression is parsed and evaluated without using third-party arithmetic libraries.
- **Error handling:** The bot notifies the user about input errors and division by zero.

## Requirements

- [Go](https://golang.org) (version 1.16 or higher)
- Telegram bot token. It can be obtained via [@BotFather](https://t.me/BotFather).

## Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/yourusername/telegram-calculator-bot.git
   cd telegram-calculator-bot
