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

	url := `https://www.vlr.gg/168037/95x-esports-vs-built-for-greatness-challengers-league-oceania-split-1-w3/?game=113776&tab=overview`
	db.MigrateDB()
	get.InsertWithURL(url)
}
