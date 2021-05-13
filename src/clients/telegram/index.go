package telegram

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	dal "vaccine-bot-lamda-aws/src/dal"
	notificationModule "vaccine-bot-lamda-aws/src/modules/notification"
)

// Initialize ...
func Initialize() (err error) {
	// Connects bot
	bot, updates, err := connect()
	if err != nil {
		return err
	}
	log.Printf("Bot connect successful")

	// Connect mongo db
	err = dal.Initialize()
	if err != nil {
		return err
	}
	log.Printf("mongodb connect successful")

	// Runs in background for sending msg for notification
	if Notification {
		go func(bot *tgbotapi.BotAPI) {
			notification := notificationModule.New(bot)
			notification.Send()
		}(bot)
	}

	// Starts conversation
	tBot := new(bot, updates)
	tBot.startConversation()

	// Returns
	return nil
}
