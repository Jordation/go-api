package main

import (
	"encoding/json"
	"fmt"
	"go-api/api"
	_ "go-api/orm"
	"go-api/server/graphs"
	_ "go-api/server/scraper"
	"log"
	"os"
)

func ReadQuery() (api.QueryForm, error) {
	var query api.QueryForm
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
	q, err := ReadQuery()
	if err != nil {
		log.Println(err)
	}
	graphs.GetGroupedBarData(q)
}
