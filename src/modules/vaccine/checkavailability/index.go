package checkavailability

import (
	"encoding/json"
	"errors"
	"log"
	"strconv"
	"time"

	"github.com/asaskevich/govalidator"

	cowin "vaccine-bot-lamda-aws/src/thirdparty/cowin"
	utils "vaccine-bot-lamda-aws/src/utils"
)

func CheckAvailability(pincode string) (msg *[]string, err error) {
	// Validates input
	errMsg, err := validates(pincode)
	if err != nil {
		return errMsg, err
	}

	// Forms request
	reqParams, headers := formRequest(pincode)
	log.Println("Request details - ", map[string]interface{}{"params": reqParams, "headers": headers})

	// Hits GET request
	statusCode, res, err := utils.HitsWebGETRequest(cowin.Host+cowin.EndPoint, reqParams, headers)
	if err != nil {
		return nil, err
	}
	log.Println("Web response recieved - ", map[string]interface{}{"code": statusCode, "res": res, "err": err})

	// Forms api response
	response, err := formAPIResponse(statusCode, res)
	if err != nil {
		return nil, err
	}

	// Forms response to show user side
	centers := formResponse(response)

	// Returns
	return &centers, nil
}

func validates(pincode string) (errMsg *[]string, err error) {
	// Checks for numeric
	if !govalidator.IsNumeric(pincode) {
		return &[]string{"Numeric Input Required"}, ErrNumRequired
	}

	// Checks for length
	if len(pincode) < 6 {
		return &[]string{"Invalid Pincode"}, ErrPinInvalid
	}
	if len(pincode) > 6 {
		return &[]string{"Pincode not more than 6 digits"}, ErrPinLenInvalid
	}

	// Returns
	return nil, nil
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

func formResponse(res *cowin.APIResponse) (centers []string) {
	centers = []string{}
	if res.Centers == nil || len(res.Centers) == 0 {
		msg := "No Vaccination Center is available for booking"
		return []string{msg}
	}

	for key, center := range res.Centers {
		cname := center.Name
		addr := center.Address
		info := ""
		if len(center.Sessions) > 0 {
			for i, sess := range center.Sessions {
				info = info + "\t\tInfo - " + strconv.Itoa(i+1) + "\n" + "\t\t---------------" + "\n"
				d := sess.Date
				a := sess.AvailableCapacity
				minAge := sess.MinAgeLimit
				v := sess.Vaccine

				info = info + "\t\t\t| Date: " + d + "\n" + "\t\t\t| Available: " + strconv.Itoa(a) + "\n" + "\t\t\t| Vaccine: " + v + "\n" + "\t\t\t| MinAge: " + strconv.Itoa(minAge) + "\n" + "\t\t---------------" + "\n"
			}
		}
		msg := "Center - " + strconv.Itoa(key+1) + "\n" + "\t| Center: " + cname + "\n" + "\t| Address: " + addr + "\n" + info
		centers = append(centers, msg)
	}

	// Returns
	return centers
}

func formAPIResponse(statusCode *int, res map[string]interface{}) (response *cowin.APIResponse, err error) {
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
