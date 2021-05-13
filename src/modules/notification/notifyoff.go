package notification

import (
	"errors"
	"vaccine-bot-lamda-aws/src/dal"
)

func off(chatID int64) (res *[]string, err error) {
	// Gets
	record, err := dal.Get(chatID)
	if err != nil {
		return nil, errors.New("NOTIFY_OFF.DAL.GET_ERR")
	}

	// Backup
	bReq := map[string]interface{}{
		"chatID":  record.ChatID,
		"name":    record.Name,
		"pincode": record.Pincode,
	}
	err = dal.BackUp(bReq)
	if err != nil {
		return nil, errors.New("NOTIFY_OFF.DAL.BACKUP_ERR")
	}

	// Deletes
	err = dal.Delete(chatID)
	if err != nil {
		if err.Error() == "REC_NOT_FOUND" {
			return &[]string{"You haven't set notification for any pincode !"}, err
		}

		// Returns
		return nil, err
	}

	// Returns
	return &[]string{"You have turn off the notification!"}, nil
}
