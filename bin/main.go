package main

import (
	"github.com/rafaelcpalmeida/FCP-games-notifier/src"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Getenv("FCP_GAMES_NOTIFIER_RECIPIENT_NUMBER")) == 0 {
		log.Fatal("Recipient number environment variable not found")
	}

	if gameDay, games := src.PortoPlaysToday(); gameDay {
		sms := strings.Builder{}
		sms.WriteString("⚠️⚠️⚠️ Beware! ⚠️⚠️⚠️\nToday is game day.\n\nGames today:\n\n")

		for _, game := range games {
			sms.WriteString(game.Team + " at " + game.Time + "\n")
		}

		src.SendSMS(os.Getenv("FCP_GAMES_NOTIFIER_RECIPIENT_NUMBER"), sms.String())
	}
}
