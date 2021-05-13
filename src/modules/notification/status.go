package notification

import (
	"errors"
	"vaccine-bot-lamda-aws/src/dal"
)

func getStatus(chatID int64) (res *[]string, err error) {
	// Calls dal
	record, err := dal.Get(chatID)
	if err != nil {
		return nil, errors.New("GET_STATUS.DAL.GET_ERR")
	}

	// Checks if user did not set any notification
	if record == nil && err == nil {
		return &[]string{"You haven't set notification for any pincode"}, errors.New("NO_RECORD_FOUND")
	}

	// Otherwise returns record
	return &[]string{"Your notification is enable on\nPincode: " + record.Pincode}, nil
}
