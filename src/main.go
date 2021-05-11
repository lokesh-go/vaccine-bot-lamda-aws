package main

import (
	"log"

	"vaccine-bot-lamda-aws/src/clients/telegram"
)

func main() {
	// Initializes telegram bot connection
	err := telegram.Initialize()
	if err != nil {
		log.Panic(err)
	}
}
