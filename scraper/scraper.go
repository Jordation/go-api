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
	"log"
	s "strings"

	"github.com/gocolly/colly"
)

// instantiate collector
var c = colly.NewCollector(
	colly.AllowedDomains("www.vlr.gg"),
	//colly.Debugger(&debug.LogDebugger{}),
)

func stripStr(str string) string {
	return s.ReplaceAll(s.ReplaceAll(str, "\t", ""), "\n", "")
}

func Scrape(url string) {
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	// vars
	var (
		cntT, cntCT int

		statsT, statsCT,
		agentsPlayed, playerNames, mapNames,
		headerData []string
	)

	//get header data
	c.OnHTML("div.wf-title-med", func(e *colly.HTMLElement) {
		headerData = append(headerData, stripStr(e.Text))
	})
	c.OnHTML("div.js-spoiler(2)", func(e *colly.HTMLElement) {
		headerData = append(headerData, stripStr(e.Text))
	})
	c.OnHTML("a.match-header-event>div", func(e *colly.HTMLElement) {
		e.ForEach("div", func(_ int, e2 *colly.HTMLElement) {
			headerData = append(headerData, stripStr(e2.Text))
		})
	})
	c.OnHTML("div.match-header-date", func(e *colly.HTMLElement) {
		headerData = append(headerData, stripStr(e.Text))
	})
	c.OnHTML("div.js-spoiler ", func(e *colly.HTMLElement) {
		headerData = append(headerData, stripStr(e.Text))
	})

	// get t stats
	c.OnHTML("span.mod-t.side", func(e *colly.HTMLElement) {
		if cntT < 122 || cntT > 241 { // combined stat panel indicies, dont need
			statsT = append(statsT, stripStr(e.Text))
		}
		cntT++
	})

	// get ct stats
	c.OnHTML("span.mod-ct.side", func(e *colly.HTMLElement) {
		if cntCT < 122 || cntCT > 241 { // combined stat panel indicies, dont need
			statsCT = append(statsCT, stripStr(e.Text))
		}
		cntCT++
	})

	// get agents played
	c.OnHTML("span.mod-agent", func(e *colly.HTMLElement) {
		agentsPlayed = append(agentsPlayed, e.ChildAttr("img", "title"))
	})

	// get player names
	c.OnHTML("td.mod-player", func(e *colly.HTMLElement) {
		f := s.Split(stripStr(e.Text), " ")
		name := s.Join((f[:len(f)-1]), " ")
		playerNames = append(playerNames, name)
	})

	//get map names
	c.OnHTML("div.vm-stats-gamesnav-item", func(e *colly.HTMLElement) {
		mapNames = append(mapNames, stripStr(e.Text))
	})
	c.Visit(url)
	// SCRAPED

	// validate shape of data
	if len(playerNames)%10 != 0 {
		log.Fatal("Malformed Data - Subs used between maps")
	}
	if len(statsT) != len(statsCT) {
		log.Fatal("Malformed HTML - Stats T and CT length do not match")
	}
	if (len(statsT)%120 != 0) || (len(statsCT)%120 != 0) {
		log.Fatal("Malformed HTML - Expected data points not found")
	}

	players := playerNames[:10]
	// instead should be labelling statsT[:120], statsCT[:120], theory works just need to form in the map info parts

	//final_data := labelAllData(statsT[:120], players)

	// fix validation so it works on BO1 games
	// test old stats pages before rating existed

	ProcessRawData(statsT, statsCT, players, headerData, mapNames)

	fmt.Println("hello")
}
