package checkavailability

import (
	"fmt"
	"strconv"

	"github.com/asaskevich/govalidator"

	cowin "vaccine-bot-lamda-aws/src/thirdparty/cowin"
	cowinCheck "vaccine-bot-lamda-aws/src/thirdparty/cowin/checking"
)

func CheckAvailability(pincode string) (msg *[]string, err error) {
	// Validates input
	errMsg, err := validates(pincode)
	if err != nil {
		return errMsg, err
	}

	// Calls cowin
	res, err := cowinCheck.Availability(pincode)
	if err != nil {
		return nil, err
	}

	// Forms response to show user side
	centers := formResponse(res)

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

				info = info + "\t\t\t| Date: " + d + "\n" + "\t\t\t| Available: " + fmt.Sprintf("%v", a) + "\n" + "\t\t\t| Vaccine: " + v + "\n" + "\t\t\t| MinAge: " + strconv.Itoa(minAge) + "\n" + "\t\t---------------" + "\n"
			}
		}
		msg := "Center - " + strconv.Itoa(key+1) + "\n" + "\t| Center: " + cname + "\n" + "\t| Address: " + addr + "\n" + info
		centers = append(centers, msg)
	}

	// Returns
	return centers
}
