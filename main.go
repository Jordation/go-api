package main

import (
	"encoding/json"
	"fmt"
	a "go-api/api"
	"go-api/orm"
	"go-api/server/scraper"
	_ "log"
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
	//q, err := ReadQuery()
	//if err != nil {
	//	log.Println(err)
	//}
	//a.GetGroupedBarData(q)
	orm.MigrateDB()
	scraper.InsertFromTxtList()
	orm.CleanDB()
}
