package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (tBot *telegramBot) sendMessage(chatID int64, msg string) {
	// Forms new msg
	sendMsg := tgbotapi.NewMessage(chatID, msg)

	// Sends
	tBot.conn.Send(sendMsg)
}
