package GetMyData

import (
	gormdb "go-api/initial/my_db"
	"log"
)

func InsertWithURL(url string) {

	err := gormdb.CheckMatchLink(url)
	if err != nil {
		log.Println("url already hit: ", url)
	}

	data := Scrape(url)

	err = ValidateScrapedData(data)
	if err != nil {
		log.Println("Not inserting URL: ", url, err)
	} else {
		database_data := MakeORMstruct(data)
		gormdb.InsertEvent(database_data)
	}
}
