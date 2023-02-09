package scraper

// url https://www.vlr.gg/168037/95x-esports-vs-built-for-greatness-challengers-league-oceania-split-1-w3
// class="wf-card  match-header" - Team names, score, event title, match title, date
// class="vm-stats-game" - stats !class="mod-active"
//class="vm-stats-game-header" - Team1 + 2 Scores for atk + def , Mapname
//class="team", "map", "team mod-right"
// class="wf-table-inset mod-overview" - stats table
// "mod-agents", "mod-stat"
// "side mod-side mod-t", "side mod-side mod-ct"

import (
	"fmt"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
)

func Scrape(url string) {
	c := colly.NewCollector(
		colly.AllowedDomains("www.vlr.gg"),
		colly.Debugger(&debug.LogDebugger{}),
	)
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnHTML("td.mod-stat", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
	})

	c.Visit(url)

}
