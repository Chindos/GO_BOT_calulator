package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() { //bot token from bot father
	token := os.Getenv("TELEGRAM_TOKEN")
	if token == "" {
		log.Fatal("Error dont have bot token")
	}

	//iniclaithatsion bots
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("bot: %s", bot.Self.UserName)

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	updates := bot.GetUpdatesChan(updateConfig)

	//whatching stdin
	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("Message %s: %s", update.Message.From.UserName, update.Message.Text)

		//calculating
		result, err := evaluateExpression(update.Message.Text)
		var reply string
		if err != nil {
			reply = "Error: " + err.Error()
		} else {
			reply = fmt.Sprintf("Result: %v", result)
		}

		// giving him result
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
		if _, err := bot.Send(msg); err != nil {
			log.Printf("Error giving result: %v", err)
		}
	}
}

func evaluateExpression(expr string) (float64, error) {

	expr = strings.ReplaceAll(expr, " ", "")

	var operator rune
	opIndex := -1
	for i, ch := range expr {
		if i == 0 && ch == '-' {
			continue
		}
		if ch == '+' || ch == '-' || ch == '*' || ch == '/' {
			operator = ch
			opIndex = i
			break
		}
	}
	if opIndex == -1 {
		return 0, fmt.Errorf("operator not found. Enter an expression like number+number")
	}

	leftPart := expr[:opIndex]
	rightPart := expr[opIndex+1:]

	left, err := strconv.ParseFloat(leftPart, 64)
	if err != nil {
		return 0, fmt.Errorf("Error left side: %v", err)
	}
	right, err := strconv.ParseFloat(rightPart, 64)
	if err != nil {
		return 0, fmt.Errorf("Error right side: %v", err)
	}

	switch operator {
	case '+':
		return left + right, nil
	case '-':
		return left - right, nil
	case '*':
		return left * right, nil
	case '/':
		if right == 0 {
			return 0, fmt.Errorf("Error /0!")
		}
		return left / right, nil
	default:
		return 0, fmt.Errorf("Unknow operator: %c", operator)
	}
}
