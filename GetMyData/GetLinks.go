package GetMyData

import (
	"bufio"
	"os"

	"github.com/gocolly/colly"
)

func ReadEachLineFromFile() []string {
	//read each line from file
	var links []string
	file, err := os.Open("./events.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		links = append(links, scanner.Text())
	}
	file.Close()
	return links
}

func GetMatchLinksFromEvent(url string) []string {
	var links []string
	c := colly.NewCollector(colly.AllowedDomains("www.vlr.gg"))
	c.OnHTML("a.wf-module-item.match-item", func(e *colly.HTMLElement) {
		url := e.Attr("href")
		links = append(links, "https://www.vlr.gg"+url)
	})
	c.Visit(url)
	return links
}

func InsertFromTextFile() error {
	var matchLinks []string

	eventLinks := ReadEachLineFromFile()
	for _, link := range eventLinks {
		matchLinks = append(matchLinks, GetMatchLinksFromEvent(link)...)
	}

	for _, l := range matchLinks {
		InsertWithURL(l)
	}

	return nil
}
