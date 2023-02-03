package processQuery

// 1. Get all rows matching global filters
// 2. Group rows by first X target
// 3. Refine initial groups by x split target
// 4. Function to handle data processes
//     4.1. Average Values over group
//     4.2. Highest single result in group
// 5. Sort rows and adjust size
//     5.1. Max dataset width trim
//     5.2.
// 6. Generate Labels
// 7. Shape data for chartjs

import (
	"encoding/json"
	"fmt"
	"go-api/initial/api"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DataGroup [2]string

func ProcessQuery() error {

	db, err := gorm.Open(sqlite.Open("my_db/test.db"), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	// read query from json
	query, err := GetQuery()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	// get sql statment from query
	stmt, args, err := query.Global_Filters.MakeSQLStmt(false)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	var results []map[string]interface{}

	// get rows from db
	db.Raw(stmt, args...).Find(&results)

	//groups, err := query.Graph_Params.Make_Data_Groups(rows)
	println("Hello")
	return nil
}

func GetQuery() (api.QueryForm, error) {
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
