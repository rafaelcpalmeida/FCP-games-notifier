package src

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
	"log"
	"strings"
	"time"
)

type game struct {
	Time	string
	Team	string
}

func PortoPlaysToday() (bool, []game) {
	gamesToday := parsePage()

	if len(gamesToday) == 0 {
		return false, nil
	}

	return true, gamesToday
}

func parsePage() []game {
	var games []game
	c := colly.NewCollector(colly.Debugger(&debug.WebDebugger{}))

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	c.OnHTML("body", func(body *colly.HTMLElement) {
		body.DOM.Find("table.ink-table.ink-table-f365.alternating.all-100").First().Find("tr>td").Parent().Each(func(_ int, s *goquery.Selection) {
			currentDate := time.Now().Local().Format("02-01-2006")
			gameDate := s.Find("td.align-left.hide-large.hide-medium.hide-small.hide-tiny.cell-nowrap").Text()

			if currentDate == gameDate {
				games = append(games, game{
					Time: s.Find("td.align-center.hide-tiny.hide-small").Text(),
					Team: strings.TrimSpace(s.Find("td.align-right").Text()),
				})
			}
		})
	})

	c.Visit("https://www.futebol365.pt/estadio/300/proximos-jogos/")

	c.Wait()

	return games
}
