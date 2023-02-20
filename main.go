package main

import (
	"encoding/json"
	"fmt"
	a "go-api/initial/api"

	get "go-api/initial/GetMyData"
	db "go-api/initial/my_db"

	//pq "go-api/initial/processQuery"
	"os"
)

func ReadQuery() (a.QueryForm, error) {
	var query a.QueryForm
	bytes, err := os.ReadFile("query.json")
	if err != nil {
		fmt.Println(err.Error())
		return query, err
	}

	json.Unmarshal(bytes, &query.Global_Filters)
	json.Unmarshal(bytes, &query.Data_Params)
	json.Unmarshal(bytes, &query.Graph_Params)

	return query, nil
}

func main() {

	db.MigrateDB()

	//url := `https://www.vlr.gg/10258/envy-vs-sentinels-champions-tour-north-america-stage-1-challengers-2-gf`
	//data := get.Scrape(url)
	//get.MakeORMstruct(data)

	get.InsertFromTextFile()
}
