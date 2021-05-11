package telegram

import (
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type telegramBot struct {
	conn    *tgbotapi.BotAPI
	updates tgbotapi.UpdatesChannel
}

func connect() (bot *tgbotapi.BotAPI, updates tgbotapi.UpdatesChannel, err error) {
	// Connects
	bot, err = tgbotapi.NewBotAPI(os.Getenv("TBotToken"))
	if err != nil {
		return nil, nil, err
	}

	// Enables Debuger
	bot.Debug = Debug

	// Update config
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	// GetUpdatesChan starts and returns a channel for getting updates
	updates, _ = bot.GetUpdatesChan(u)

	// Returns
	return bot, updates, nil
}
