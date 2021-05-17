package telegram

import (
	"errors"
	"strconv"
	"strings"
	"vaccine-bot-lamda-aws/src/modules/notification"
	vaccine "vaccine-bot-lamda-aws/src/modules/vaccine/checkavailability"
)

func (tBot *telegramBot) getMessages(request string, chatID int64, adminChatID int, name string) (response *[]string, err error) {
	// Checks cmd msgs
	reqMsg := strings.ToLower(request)

	// Case of greet msg
	if reqMsg == "/start" || reqMsg == "hi" || reqMsg == "hello" || reqMsg == "hey" {
		// Froms msgs
		greetMsg := "Welcome!\n\nBot helps to check vaccine availability\nplease enter valid pincode to check\n\nYou can also add notification for particular pincode\n\nFor ON\n\t/notify on YOUR_PIN_CODE\n\nFor check status\n\t/notify status\n\nFor OFF\n\t/notify off\n\nNote:- The bot gets data from Co-WIN public APIs.\n\nPlease give feedback on\n/feedback YOUR_MSG\n\nDeveloper: Lokesh Chandra"
		// Returns
		return &[]string{greetMsg}, nil
	} else if reqMsg == "/help" || reqMsg == "help" {
		// Forms msgs
		helpMsg := "Welcome :-)\n\nBot helps to check vaccine availability\n1. Simply you can enter the pincode and gets the Vaccine availibility details.\n\n2. You can turn ON or OFF the notification for particular pincode so once the vaccine is available in your area bot will automatically notify you. [ Only 18+ ]\n\nCommands to set notification\n\na. For Trun On Notification\n/notify on YOUR_PIN_CODE\n\nb. For Turn Off Notification\n/notify off\n\nc. Check your notification status\n/notify status\n\nPlease give feedback on\n/feedback YOUR_MSG\n\nTake vaccine and be safe :-)\nLokesh Chandra :-)\n"
		// Returns
		return &[]string{helpMsg}, nil
	} else if strings.HasPrefix(reqMsg, "/adm") && chatID == int64(adminChatID) {
		// Splits msgs
		s := strings.SplitN(reqMsg, " ", 3)

		// Checks length
		if len(s) != 3 {
			return &[]string{"Wrong Input"}, errors.New("ADMIN_WRONG_INPUT")
		}

		switch s[1] {
		case "eve":
			{
				tBot.sendMsgToAll(s[2], adminChatID)
			}
		default:
			{
				cID, _ := strconv.Atoi(s[1])
				m := s[2]
				tBot.sendMessage(int64(cID), m)
			}
		}
	} else if strings.HasPrefix(reqMsg, "/notify") {
		// Checks and Sets notification
		res, err := notification.CheckAndSet(reqMsg, chatID, name)
		// Returns
		return res, err
	} else if strings.HasPrefix(reqMsg, "/feedback") {
		// Splits msgs
		f := strings.SplitN(reqMsg, " ", 2)
		if len(f) != 2 {
			return &[]string{"Please input\n\n/feedback YOUR_MSG"}, errors.New("FEEDBACK_INPUT_ERR")
		}

		fmsg := "******* Feedback *******\n\nID: " + strconv.Itoa(int(chatID)) + "\nName: " + name + "\n\nFeedback msg:- \n" + f[1]
		// Sends feedback to admin
		tBot.sendMessage(int64(adminChatID), fmsg)
		// Sends thanks to user
		tBot.sendMessage(chatID, "Your feedback received\nThanks :-)")

	} else {
		// Checks for vaccine availability
		response, err = vaccine.CheckAvailability(reqMsg)
	}

	// Returns
	return response, err
}
