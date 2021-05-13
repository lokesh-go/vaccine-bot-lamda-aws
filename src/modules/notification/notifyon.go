package notification

import (
	"errors"
	"strings"

	"vaccine-bot-lamda-aws/src/dal"
	vaccine "vaccine-bot-lamda-aws/src/modules/vaccine/checkavailability"
)

func on(req string, chatID int64, name string) (res *[]string, err error) {
	// Splits request
	msgs := strings.SplitN(req, " ", 3)

	// Gets pincode form msgs
	pincode := strings.Trim(msgs[2], " ")

	// Checks center is available or not for this pincode
	resp, err := vaccine.CheckAvailability(pincode)
	if err != nil {
		return resp, err
	}

	// Checks for center available or not
	if resp != nil && len(*resp) > 0 {
		if (*resp)[0] == "No Vaccination Center is available for booking" {
			return &[]string{"There is no any center available for booking!\n\nSo, you can't set notification on this pincode!\nPlease try another pincode!"}, errors.New("NO_CENTER_AVAIL_TO_SET_NOTIFICATION")
		}
	}

	// Checks in db already user sets notification or not
	getRes, err := dal.Get(chatID)
	if err != nil {
		return nil, errors.New("REC_GET_ERR.ON - " + err.Error())
	}

	// Checks record exists or not
	if getRes != nil {
		return &[]string{"You have already set notification!\n\nPlease off first then set notification for new pincode!"}, errors.New("ALREADY_SET_NOTIFICATION")
	}

	// Sets for notification
	err = dal.Create(chatID, pincode, name)
	if err != nil {
		return nil, errors.New("REC_CREATE_ERR.ON - " + err.Error())
	}

	// Returns
	return &[]string{"You have turn on notification for\nPincode: " + pincode + "\n\nWhenever we got vaccine in your area will notify you immediately\n\nYou can check your notification status on\n/notify status"}, nil
}
