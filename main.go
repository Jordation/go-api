package main

import (
	"encoding/json"
	"fmt"
	a "go-api/initial/api"

	db "go-api/initial/my_db"
	"go-api/initial/scraper"

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
	//query, err := ReadQuery()
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//pq.GetGroupedBarData(query)
	//return
	//url := `https://www.vlr.gg/164480/diamant-esports-vs-enterprise-esports-challengers-league-east-surge-split-1-w5`
	url2 := `https://www.vlr.gg/168037/95x-esports-vs-built-for-greatness-challengers-league-oceania-split-1-w3/?game=113776&tab=overview`
	data := scraper.Scrape(url2)
	ORMdata := scraper.ProcessRawData(data)
	_ = ORMdata
	db.MigrateDB()
}
