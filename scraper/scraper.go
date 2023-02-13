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

type statRow struct {
	statsT  []string
	statsCT []string
	player  string
	agent   string
}
type mapData struct {
	data    []string
	mapname string
	stats   []statRow
}
type matchData struct {
	data []string
}

// RawMatchData struct for fields to be scraped into
type RawMatchData struct {
	MapNames  []string
	MapCount  int
	MapData   []mapData
	matchInfo []string
}

// instantiate collector
var c = colly.NewCollector(
	colly.AllowedDomains("www.vlr.gg"),
	//colly.Debugger(&debug.LogDebugger{}),
)

// new line killer
func strip(str string) string {
	return s.ReplaceAll(s.ReplaceAll(str, "\t", ""), "\n", "")
}

// Scrape - get stats for match of provided url
func Scrape(url string) {
	var rd RawMatchData

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	//get map names
	c.OnHTML("div.map>div>span", func(e *colly.HTMLElement) {
		rd.MapNames = append(rd.MapNames, strip(s.Split(e.Text, "PICK")[0]))
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
			// group scores, mapnames, teamnames
			var mapdata mapData
			e.ForEach("div.vm-stats-game-header>div.team>div>div.team-name", func(_ int, e2 *colly.HTMLElement) {
				mapdata.data = append(mapdata.data, strip(e2.Text))
			})
			e.ForEach("div.vm-stats-game-header>div.team>div>span.mod-ct", func(_ int, e2 *colly.HTMLElement) {
				mapdata.data = append(mapdata.data, strip(e2.Text))
			})
			e.ForEach("div.vm-stats-game-header>div.team>div>span.mod-t", func(_ int, e2 *colly.HTMLElement) {
				mapdata.data = append(mapdata.data, strip(e2.Text))
			})
			mapdata.mapname = rd.MapNames[rd.MapCount]

			// group rows of player stats
			e.ForEach("div>div>table>tbody>tr", func(_ int, e2 *colly.HTMLElement) {
				var stats statRow
				e2.ForEach("td>span>span.mod-ct", func(_ int, e3 *colly.HTMLElement) {
					stats.statsCT = append(stats.statsCT, strip(e3.Text))
				})
				e2.ForEach("td>span>span.mod-t", func(_ int, e3 *colly.HTMLElement) {
					stats.statsT = append(stats.statsT, strip(e3.Text))
				})
				e2.ForEach("td>div>a>div.text-of", func(_ int, e3 *colly.HTMLElement) {
					stats.player = strip(e3.Text)
				})
				e2.ForEach("td>div>span>img", func(_ int, e3 *colly.HTMLElement) {
					stats.agent = e3.Attr("title")
				})
				mapdata.stats = append(mapdata.stats, stats)
			})
			rd.MapData = append(rd.MapData, mapdata)
			rd.MapCount++
		}
	})

	c.Visit(url)
	// SCRAPED

	// fix validation so it works on BO1 games
	// test old stats pages before rating existed

	ProcessRawData(rd)

	fmt.Println("hello")
}
