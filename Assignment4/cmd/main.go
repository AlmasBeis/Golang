package main

import (
	"Assignment4/pkg"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("6205463250:AAHd5GJt1XeZG2zsjAEdb5jbOGk-p7T6Xhs")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	telegramBot := pkg.NewBot(bot)
	if err := telegramBot.Start(); err != nil {
		log.Fatal(err)
	}
}
