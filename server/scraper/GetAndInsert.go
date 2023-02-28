package scraper

import (
	"go-api/orm"
	"log"
)

// wg *sync.WaitGroup
func InsertWithURL(url string) {
	data := *Scrape(url)

	if err := ValidateScrapedData(data); err != nil {
		log.Println("Not inserting URL: ", url, err)
		return
	}

	database_data := MakeORMstruct(data)
	orm.InsertEvent(database_data)
}

func InsertFromTxtList() {
	var (
		matchLinks []string
		//wg         sync.WaitGroup
	)

	events := ReadEachLineFromFile()
	for _, l := range events {
		matchLinks = append(matchLinks, GetMatchLinksFromEvent(l)...)
	}

	for _, l := range matchLinks {
		if err := orm.CheckMatchLink(l); err != nil {
			continue
		}
		//wg.Add(1)
		//go InsertWithURL(l, &wg)
		InsertWithURL(l)
	}
	//wg.Wait()
}
