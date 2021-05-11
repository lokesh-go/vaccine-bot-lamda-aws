package telegram

import (
	"log"
)

// Initialize ...
func Initialize() (err error) {
	// Connects bot
	bot, updates, err := connect()
	if err != nil {
		return err
	}
	log.Printf("Bot connection successful")

	// Starts conversation
	tBot := new(bot, updates)
	tBot.startConversation()

	// Returns
	return nil
}
