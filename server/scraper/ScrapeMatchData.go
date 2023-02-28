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
	s "strings"

	"github.com/gocolly/colly"
)

type stats []string
type playerData struct {
	statsT  mappedPlayerData
	statsCT mappedPlayerData
	player  string
	agent   string
}
type mappedPlayerData map[string]string

type gameData struct {
	data    []string
	mapname string
	players []playerData
}

// matchData struct for fields to be scraped into
type matchData struct {
	mapCount  int
	mapNames  []string
	matchInfo []string
	gameData  []gameData
}

var statCategories = []string{"Rating", "ACS", "Kills", "Deaths", "Assists",
	"", "KAST", "ADR", "HSP", "FK", "FD", ""}

func (s stats) MapData() mappedPlayerData {
	ld := make(map[string]string)
	if len(s) != len(statCategories) {
		return ld
	}

	for i, v := range statCategories {
		switch v {
		case "":
		default:
			ld[v] = s[i]
		}
	}
	return ld
}

// instantiate collector
var c = colly.NewCollector(
	colly.AllowedDomains("www.vlr.gg"),
	//colly.Debugger(&debug.LogDebugger{}),
)

// kills formatting
func strip(str string) string {
	return s.ReplaceAll(s.ReplaceAll(s.ReplaceAll(str, "\t", ""), "\n", ""), "%", "")
}

// Scrape - get stats for match of provided url
func Scrape(url string) *matchData {
	var rd matchData
	fmt.Println("Visiting: ", url)
	//get map names
	c.OnHTML("div.map>div>span", func(e *colly.HTMLElement) {
		rd.mapNames = append(rd.mapNames, strip(s.Split(e.Text, "PICK")[0]))
	})

	//get header data
	c.OnHTML("div.wf-card.match-header", func(e *colly.HTMLElement) {
		e.ForEach("div.match-header-super>div>a>div>div", func(_ int, e2 *colly.HTMLElement) {
			rd.matchInfo = append(rd.matchInfo, strip(e2.Text))
		})
		e.ForEach("div.match-header-super>div>div>div.moment-tz-convert", func(_ int, e2 *colly.HTMLElement) {
			rd.matchInfo = append(rd.matchInfo, strip(e2.Text))
		})
		e.ForEach("div.match-header-vs>a>div>div.wf-title-med", func(_ int, e2 *colly.HTMLElement) {
			rd.matchInfo = append(rd.matchInfo, strip(e2.Text))
		})
	})

	// get stats
	c.OnHTML("div.vm-stats-game", func(e *colly.HTMLElement) {
		id := e.Attr("data-game-id")
		if id != "all" {
			// group scores, mapNames, teamnames
			var mapdata gameData
			e.ForEach("div.vm-stats-game-header>div.team>div>div.team-name", func(_ int, e2 *colly.HTMLElement) {
				mapdata.data = append(mapdata.data, strip(e2.Text))
			})
			e.ForEach("div.vm-stats-game-header>div.team>div>span.mod-ct", func(_ int, e2 *colly.HTMLElement) {
				mapdata.data = append(mapdata.data, strip(e2.Text))
			})
			e.ForEach("div.vm-stats-game-header>div.team>div>span.mod-t", func(_ int, e2 *colly.HTMLElement) {
				mapdata.data = append(mapdata.data, strip(e2.Text))
			})
			mapdata.mapname = rd.mapNames[rd.mapCount]

			// group rows of player stats
			e.ForEach("div>div>table>tbody>tr", func(_ int, e2 *colly.HTMLElement) {
				var d playerData
				var statT stats
				var statCT stats
				e2.ForEach("span.mod-ct", func(_ int, e3 *colly.HTMLElement) {
					statCT = append(statCT, strip(e3.Text))
				})
				e2.ForEach("span.mod-t", func(_ int, e3 *colly.HTMLElement) {
					statT = append(statT, strip(e3.Text))
				})
				e2.ForEach("td>div>a>div.text-of", func(_ int, e3 *colly.HTMLElement) {
					d.player = s.TrimRight(strip(e3.Text), " ")
				})
				e2.ForEach("td>div>span>img", func(_ int, e3 *colly.HTMLElement) {
					d.agent = e3.Attr("title")
				})
				d.statsCT = statCT.MapData()
				d.statsT = statT.MapData()
				mapdata.players = append(mapdata.players, d)
			})
			rd.gameData = append(rd.gameData, mapdata)
			rd.mapCount++
		}
	})

	c.Visit(url)

	return &rd
}
