package notification

import (
	"errors"
	"strings"
)

func CheckAndSet(reqMsg string, chatID int64, name string) (res *[]string, err error) {
	// Checks cases
	// Case of on notification
	if strings.HasPrefix(reqMsg, "/notify on") {
		res, err = on(reqMsg, chatID, name)
		// Returns
		return res, err
	}

	// Case of off notification
	if reqMsg == "/notify off" {
		res, err = off(chatID)
		// Returns
		return res, err
	}

	// Case of check notification status
	if reqMsg == "/notify status" {
		res, err = getStatus(chatID)
		// Returns
		return res, err
	}

	// Returns
	return &[]string{"Wrong Input for Notification!\nPlease input:\n\t/notify on YOUR_PIN_CODE\n\t/notify off\n\t/notify status"}, errors.New("NOTIFY_INPUT_ERR")
}
