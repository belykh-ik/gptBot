package main

import (
	"fmt"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {

	botToken := os.Getenv("TOKEN")
	if botToken == "" {
		log.Fatalf("BOT TOKEN not set")
	}

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Fatalf("Ошибка при создании бота: %v", err)
	}

	bot.Debug = true
	fmt.Printf("Бот %s успешно запущен\n", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	output, err := getToken()
	if err != nil {
		log.Fatalf("Ошибка при получении ответа: %v", err)
	}

	for update := range updates {
		if update.Message != nil && update.Message.Text != "" {
			chatID := update.Message.Chat.ID
			command := update.Message.Text

			if command == "/start" {
				msg := tgbotapi.NewMessage(chatID, "Привет! Задай свой вопрос.")
				bot.Send(msg)
			} else {
				content, err := getQuery(output, command)
				if err != nil {
					output, err = getToken()
					if err != nil {
						fmt.Printf("Ошибка при получении ответа: %v", err)
						break
					}
					msg := tgbotapi.NewMessage(chatID, "Попробуй задать вопрос еще раз.\n")
					bot.Send(msg)
					continue
				}
				msg := tgbotapi.NewMessage(chatID, fmt.Sprintf("Результат команды:\n%s", content))
				bot.Send(msg)
			}
		}
	}
}
