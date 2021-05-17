package notification

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
	dal "vaccine-bot-lamda-aws/src/dal"
	"vaccine-bot-lamda-aws/src/dal/models"
	"vaccine-bot-lamda-aws/src/thirdparty/cowin"
	"vaccine-bot-lamda-aws/src/thirdparty/cowin/checking"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type telegram struct {
	bot *tgbotapi.BotAPI
}

func New(bot *tgbotapi.BotAPI) *telegram {
	return &telegram{bot}
}

func (t *telegram) Send() {
	count := 12
	adminChatID, _ := strconv.Atoi(os.Getenv("AdminChatID"))
	// Run notification check
	for {
		// Gets All notified users list from db
		uCount := 0
		users, err := dal.GetAll()
		if err != nil {
			// Sends GetAll Err msg to admin
			t.sendMsg(int64(adminChatID), "NOTIFY.SEND.GETALL_ERR - "+err.Error())
		}

		// Removes repeated pincodes from lists
		pincodes := removeRepeatedPincode(users)
		// No of pincodes hits for notification
		log.Println("No of pincodes to hit: - ", len(pincodes))
		// Checks availablity
		msgMap, err := checkAvailability(pincodes)
		if err != nil {
			// Sends GetAll Err msg to admin
			t.sendMsg(int64(adminChatID), err.Error())
		}

		// Checks msg per user and send the msg if available
		if users != nil {
			uCount = len(*users)
			for _, user := range *users {
				// Checks vaccine is available or not
				pinMsg := msgMap[user.Pincode]
				msgSendToUser := pinMsg.(*string)
				if msgSendToUser != nil {
					// Sends msg to the user
					sendMsg := strings.ReplaceAll(*msgSendToUser, "<nameofuserJB7O0>", user.Name)
					// Sends msg to the user
					t.sendMsg(user.ChatID, sendMsg)
					// Sends msg to the admin
					//t.sendMsg(int64(adminChatID), "ChatID: "+strconv.Itoa(int(user.ChatID))+"\n"+sendMsg)
				}

				// Sends processing request every hour to admin
				if count == 12 {
					//log.Println("USER CHECK - ", user.Name)
					// Inform to the admin
					t.sendMsg(int64(adminChatID), "Checked Notification for user\n\n| ChatID: "+strconv.Itoa(int(user.ChatID))+"\n| Name: "+user.Name+"\n| Pincode: "+user.Pincode)
				}
			}
		}

		// Count 0 after every hour
		if count == 12 {
			count = 0
		}
		count++

		log.Println("Execute in every 5 minutes")
		// Inform to the admin
		t.sendMsg(int64(adminChatID), "Execute every 5 minutes for\nUsers: "+strconv.Itoa(uCount)+"\nNo of req per 5 mint hits on Co-Win: "+strconv.Itoa(len(pincodes)))
		// Time interval to check
		time.Sleep(5 * time.Minute)
	}
}

func (t *telegram) sendMsg(chatID int64, msg string) {
	// Forms new msg
	sendMsg := tgbotapi.NewMessage(chatID, msg)

	// Sends
	t.bot.Send(sendMsg)
}

func checkAvailability(pincodes []string) (msgMap map[string]interface{}, err error) {
	msgMap = map[string]interface{}{}
	for _, pincode := range pincodes {
		// Gets vaccine details from cowin
		records, err := checking.Availability(pincode)
		if err != nil {
			return nil, errors.New("NOTIFY.SEND.CHECK_AVAIL_COWIN_ERR - " + pincode + err.Error())
		}

		msg := checkVaccineCount(records)
		msgMap[pincode] = msg
	}

	// Returns
	return msgMap, nil
}

func checkVaccineCount(records *cowin.APIResponse) (msg *string) {
	// Checks vaccine count
	if records != nil && records.Centers != nil && len(records.Centers) > 0 {
		for _, center := range records.Centers {
			if center.Sessions != nil && len(center.Sessions) > 0 {
				for _, session := range center.Sessions {
					if session.AvailableCapacity > 0 && session.MinAgeLimit == 18 {
						// Available vaccine in user area
						userMsg := "******* Got the Vaccine *******\n\nHey <nameofuserJB7O0> !\nVaccines are available in you Area\n\t| Pincode: " + strconv.Itoa(center.Pincode) + "\n\t| Center: " + center.Name + "\n\t| Date: " + session.Date + "\n\t| Available: " + fmt.Sprintf("%v", session.AvailableCapacity) + "\n\t| Vaccine: " + session.Vaccine + "\n\t| Age: " + strconv.Itoa(session.MinAgeLimit) + "\n\nHurry up! Go and book vaccine :-)"

						// Returns
						return &userMsg
					}
				}
			}
		}
	}

	// Returns
	return nil
}

func removeRepeatedPincode(users *[]models.Get) (pincodes []string) {
	pincodes = []string{}
	occured := map[string]bool{}

	if users != nil {
		for _, user := range *users {
			if occured[user.Pincode] != true {
				occured[user.Pincode] = true

				// Append to pincodes
				pincodes = append(pincodes, user.Pincode)
			}
		}
	}

	// Returns
	return pincodes
}
