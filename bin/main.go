package main

import (
	"FCP-games-notifier/src"
	"log"
	"os"
)

func main() {
	if len(os.Getenv("FCP_GAMES_NOTIFIER_RECIPIENT_NUMBER")) == 0 {
		log.Fatal("Recipient number environment variable not found")
	}

	src.SendSMS(os.Getenv("FCP_GAMES_NOTIFIER_RECIPIENT_NUMBER"), "Isto é um assalto!")
}
