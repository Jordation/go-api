package GetMyData

import (
	gormdb "go-api/initial/my_db"
	"log"
	"sync"
)

func InsertWithURL(url string, wg *sync.WaitGroup) {
	defer wg.Done()

	data := Scrape(url)

	if err := ValidateScrapedData(data); err != nil {
		log.Println("Not inserting URL: ", url, err)
		return
	}

	database_data := MakeORMstruct(data)
	gormdb.InsertEvent(database_data)
}

func InsertFromTxtList() {
	var (
		matchLinks []string
		wg         sync.WaitGroup
	)

	events := ReadEachLineFromFile()
	for _, l := range events {
		matchLinks = append(matchLinks, GetMatchLinksFromEvent(l)...)
	}

	for _, l := range matchLinks {
		if err := gormdb.CheckMatchLink(l); err != nil {
			continue
		}
		wg.Add(1)
		go InsertWithURL(l, &wg)
	}
}
