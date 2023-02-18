package GetMyData

import db "go-api/initial/my_db"

func InsertWithURL(url string) {
	data := Scrape(url)
	database_data := MakeORMstruct(data)
	db.InsertData(database_data)
}
