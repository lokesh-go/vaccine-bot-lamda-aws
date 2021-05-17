package telegram

import (
	"log"
	"os"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func new(bot *tgbotapi.BotAPI, updates tgbotapi.UpdatesChannel) *telegramBot {
	return &telegramBot{bot, updates}
}

func (tBot *telegramBot) startConversation() {
	// Starts
	for update := range tBot.updates {
		if update.Message == nil {
			continue
		}
		log.Println("Recieved request from telegram - ", update.Message.Chat.ID)

		// Forms admin ID
		adminChatID, _ := strconv.Atoi(os.Getenv("AdminChatID"))

		// Gets msgs to send the users
		logErr := ""
		res, err := tBot.getMessages(update.Message.Text, update.Message.Chat.ID, adminChatID, update.Message.From.FirstName + " " + update.Message.From.LastName)
		if err != nil {
			logErr = err.Error()
			if res == nil {
				res = &[]string{"Something went wrong !\nPlease try later !"}
			}
		}

		// Sends logs to admin
		if update.Message.Chat.ID != int64(adminChatID) || err != nil {
			// Froms log msg
			logMsg := "User Request Logs\n------------------------------\n\t| ChatID: " + strconv.Itoa(int(update.Message.Chat.ID)) + "\n\t| UserName: " + update.Message.From.UserName + "\n\t| Name: " + update.Message.From.FirstName + " " + update.Message.From.LastName + "\n\t| Request: " + update.Message.Text + "\n\t| Error: " + logErr

			// Case of notify
			if strings.Contains(update.Message.Text, "/notify") {
				logMsg = logMsg + "\n\t| Res: " + (*res)[0]
			}

			// Sends logs to admin
			tBot.sendMessage(int64(adminChatID), logMsg)
		}

		// Sends msgs to the client
		if res != nil {
			for _, m := range *res {
				// Sends
				tBot.sendMessage(update.Message.Chat.ID, m)
			}
		}
	}
}
