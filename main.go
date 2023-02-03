package main

import (
	"encoding/json"
	"fmt"
	a "go-api/initial/api"
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
	var (
		filters a.ListPlayerStatsFilters
	)
	query, err := ReadQuery()
	if err != nil {
		fmt.Println(err.Error())
	}

	filters.Columns = []string{query.Graph_Params.X_target, query.Graph_Params.X2_target}
	filters.Unique = true
	a.ListPlayerStats(filters, *query.Global_Filters)

	//processQuery.ProcessQuery()
}
