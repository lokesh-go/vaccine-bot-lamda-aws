package telegram

import vaccine "vaccine-bot-lamda-aws/src/modules/vaccine/checkavailability"

func getMessages(request string) (response *[]string, err error) {
	// Checks for greet msg
	greetMsg := checkGreetMsg(request)
	if greetMsg != nil {
		return greetMsg, nil
	}

	// Calls vaccine module
	centers, err := vaccine.CheckAvailability(request)
	if err != nil {
		return centers, err
	}

	// Returns
	return centers, nil
}

func checkGreetMsg(msg string) (res *[]string) {
	greetMsg := "Welcome!\n\nBot helps to check vaccine availability\nplease enter valid pincode to check\n\nDeveloper: Lokesh Chandra"
	if msg == "hi" || msg == "hello" || msg == "Hi" || msg == "Hello" || msg == "/start" {
		return &[]string{greetMsg}
	}

	// Returns
	return nil
}
