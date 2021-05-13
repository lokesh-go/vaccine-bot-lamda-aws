package checking

import (
	"encoding/json"
	"errors"
	"log"
	"strconv"
	"time"
	"vaccine-bot-lamda-aws/src/thirdparty/cowin"
	"vaccine-bot-lamda-aws/src/utils"
)

func Availability(pincode string) (res *cowin.APIResponse, err error) {
	// Forms request
	reqParams, headers := formRequest(pincode)

	// Hits GET request
	statusCode, resp, err := utils.HitsWebGETRequest(cowin.Host+cowin.EndPoint, reqParams, headers)
	if err != nil {
		return nil, err
	}
	log.Println("Web response recieved - ", map[string]interface{}{"code": statusCode, "res": res, "err": err})

	// Forms api response
	res, err = formResponse(statusCode, resp)
	if err != nil {
		return nil, err
	}

	// Returns
	return res, nil
}

func formRequest(pincode string) (queryParams map[string]interface{}, headers map[string]string) {
	// Gets current date
	// DD-MM-YYYY
	date := time.Now().Format("02-01-2006")

	// Forms params
	queryParams = map[string]interface{}{
		"pincode": pincode,
		"date":    date,
	}

	// Returns
	return queryParams, nil
}

func formResponse(statusCode *int, res map[string]interface{}) (response *cowin.APIResponse, err error) {
	// Checks status code
	if *statusCode != 200 {
		resbytes, _ := json.Marshal(res)
		resString := string(resbytes)

		// Returns
		return nil, errors.New("HIT_WEB_RES_ERR" + "\n" + "Code: " + strconv.Itoa(*statusCode) + "\n" + "Msg: " + resString)
	}

	// Forms bytes
	resbytes, _ := json.Marshal(res)
	if err != nil {
		return nil, err
	}

	// Unmarshals
	err = json.Unmarshal(resbytes, &response)
	if err != nil {
		return nil, err
	}

	// Returns
	return response, nil
}
