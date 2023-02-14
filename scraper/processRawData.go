package scraper

import (
	"fmt"
	gormdb "go-api/initial/my_db"
)

func makeMapStructs() {}

func makeORMStruct(d matchData) gormdb.Event {
	var e gormdb.Event

	return e
}

func ProcessRawData(d matchData) gormdb.Event {
	ee := makeORMStruct(d)
	fmt.Println("Hello")
	return ee
}
