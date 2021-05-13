package telegram

import (
	"strings"
	"vaccine-bot-lamda-aws/src/modules/notification"
	vaccine "vaccine-bot-lamda-aws/src/modules/vaccine/checkavailability"
)

func getMessages(request string, chatID int64, name string) (response *[]string, err error) {
	// Checks cmd msgs
	reqMsg := strings.ToLower(request)

	// Case of greet msg
	if reqMsg == "/start" || reqMsg == "hi" || reqMsg == "hello" || reqMsg == "hey" {
		greetMsg := "Welcome!\n\nBot helps to check vaccine availability\nplease enter valid pincode to check\n\nDeveloper: Lokesh Chandra"
		return &[]string{greetMsg}, nil
	}

	// Case of help
	if reqMsg == "/help" || reqMsg == "help" {
		helpMsg := "Welcome :-)\n\nBot helps to check vaccine availability\n1. Simply you can enter the pincode and gets the Vaccine availibility details.\n\n2. You can turn ON or OFF the notification for particular pincode so once the vaccine is available in your area bot will automatically notify you. [ Only 18+ ]\n\nCommands to set notification\n\na. For Trun On Notification\n/notify on YOUR_PIN_CODE\n\nb. For Turn Off Notification\n/notify off\n\nc. Check your notification status\n/notify status\n\nTake vaccine and be safe :-)\nLokesh Chandra :-)"
		return &[]string{helpMsg}, nil
	}

	// Case of set notification
	if strings.HasPrefix(reqMsg, "/notify") {
		// Checks and Sets notification
		res, err := notification.CheckAndSet(reqMsg, chatID, name)
		// Returns
		return res, err
	}

	// By default Calls vaccine module
	centers, err := vaccine.CheckAvailability(reqMsg)

	// Returns
	return centers, err
}
