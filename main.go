package main

import (
	"encoding/json"
	"fmt"
	_ "go-api/initial/GetMyData"
	a "go-api/initial/api"
	_ "go-api/initial/my_db"
	"log"
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
	q, err := ReadQuery()
	if err != nil {
		log.Println(err)
	}
	a.GetGroupedBarData(q)
	//my_db.CleanDB()
	//my_db.MigrateDB()
	//GetMyData.InsertFromTxtList()
}
