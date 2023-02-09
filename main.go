package main

import (
	"encoding/json"
	"fmt"
	a "go-api/initial/api"

	//"go-api/initial/scraper"

	pq "go-api/initial/processQuery"
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
	query, err := ReadQuery()
	if err != nil {
		fmt.Println(err.Error())
	}
	pq.GetGroupedBarData(query)
	return
	//url := `https://www.vlr.gg/168037/95x-esports-vs-built-for-greatness-challengers-league-oceania-split-1-w3`
	//scraper.Scrape(url)
}
