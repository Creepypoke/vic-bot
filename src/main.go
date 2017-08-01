package main

import (
	"_models"
	"gopkg.in/telegram-bot-api.v4"
	"log"
)

const MY_AWESOME_TOKEN = "448497992:AAHzjuyWkr3OALw1oyUsThHDbzhPiNq1cU0"

func main() {
	
	questionList := []_models.Question

	bot, err := tgbotapi.NewBotAPI(MY_AWESOME_TOKEN)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}
