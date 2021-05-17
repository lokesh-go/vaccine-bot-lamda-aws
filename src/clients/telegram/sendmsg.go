package telegram

import (
	"strings"
	"vaccine-bot-lamda-aws/src/dal"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (tBot *telegramBot) sendMessage(chatID int64, msg string) {
	// Forms new msg
	sendMsg := tgbotapi.NewMessage(chatID, msg)

	// Sends
	tBot.conn.Send(sendMsg)
}

func (tBot *telegramBot) sendMsgToAll(msg string, adminChatID int) {
	// Admin msg template
	msgTemplate := "########  ADMIN  HERE  ########\n\nHi there !\nHope you're doing well.\n\n<adminMsgToAllUsersQ46HiK9>\n\nThanks for using notification service :-)\n\nService here to help you:-\n1. Pincode Service\nSimply input pincode and check availability\n\n2. Notification Service\n\na. For Trun off notification\n/notify off\n\nb. For checking status\n/notify status\n\nc. For turn on notification\n/notify on YOUR_PIN_CODE\n\nPlease give feedback on\n/feedback YOUR_MSG\n\nKeep Safe :-)"
	if msg == "feedback" {
		msg = "Please leave your valuable feedback that will very helpful for me\n\nFor Feedback please input\n/feedback YOUR_MSG"
	}
	sendmsgToUsers := strings.ReplaceAll(msgTemplate, "<adminMsgToAllUsersQ46HiK9>", msg)

	// Calls dal
	users, err := dal.GetAll()
	if err != nil {
		sendMsg := tgbotapi.NewMessage(int64(adminChatID), "SendMsgToAll - "+err.Error())
		tBot.conn.Send(sendMsg)
	}

	if users != nil {
		for _, user := range *users {
			// Forms new msg
			sendMsg := tgbotapi.NewMessage(user.ChatID, sendmsgToUsers)
			// Sends
			tBot.conn.Send(sendMsg)
		}
	}
}
