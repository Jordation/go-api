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
	"go-api/initial/my_db"
	"os"
	"strings"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DataGroup [2]string
type StrToBool bool

// converts any text in bool fields to true, empty str to false
func (f StrToBool) UnmarshalJSON(b []byte) (err error) {
	var s string
	fmt.Print(s)
	if err := json.Unmarshal(b, &s); err != nil {
		fmt.Print(err.Error())
	}

	if s != "" {
		f = true
	} else {
		f = false
	}
	return nil
}

// helper bc i dont want to work on my frontend any more
func filter_process(f *QueryFilters) (map[string][]string, string) {
	filters := make(map[string][]string)
	filters["agent"] = strings.Split(f.Agents, ", ")
	filters["mapname"] = strings.Split(f.Mapnames, ", ")
	filters["team"] = strings.Split(f.Teams, ", ")
	filters["player"] = strings.Split(f.Players, ", ")

	for k, v := range filters {
		if v[0] == "" {
			filters[k] = nil
		}
	}

	side := "player_stats_" + f.Side

	return filters, side
}

// makes sql statement from query filters
func (f *QueryFilters) Make_SQL_Stmt() (string, []interface{}, error) {
	var (
		clauses []string
		args    []interface{}
		stmt    string
	)

	// this should be done on the client, but for now it's here
	filters, side := filter_process(f)

	stmt = "SELECT * FROM " + side

	for key, arg := range filters {
		if len(arg) > 0 {
			clauses = append(clauses, key+" IN ("+strings.Repeat("?, ", len(arg)-1)+"?)")
			for _, value := range arg {
				args = append(args, value)
			}
		}
	}

	if len(clauses) != 0 {
		stmt += " WHERE " + strings.Join(clauses, " AND ")
	}

	return stmt, args, nil
}

type QueryFilters struct {
	Side     string `json:"side"`
	Agents   string `json:"agents"`
	Mapnames string `json:"mapnames"`
	Players  string `json:"players"`
	Teams    string `json:"teams"`
}
type GraphParams struct {
	Query_level int    `json:"query_level,string"`
	Y_target    string `json:"y_target"`
	X_target    string `json:"x_target"`
	X2_target   string `json:"x2_target"`
}
type DataParams struct {
	Average_over_x    StrToBool `json:"average_over_x"`
	Order_by_y_target StrToBool `json:"order_by_y_target"`
	Min_dataset_size  int       `json:"min_dataset_size,string"`
	Max_dataset_width int       `json:"max_dataset_width,string"`
}
type QueryForm struct {
	Global_Filters *QueryFilters
	Data_Params    *DataParams
	Graph_Params   *GraphParams
}

func (p *GraphParams) Make_Data_Groups([]my_db.BetterResult) (map[DataGroup][]my_db.BetterResult, error) {

	return nil, nil
}

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
	stmt, args, err := query.Global_Filters.Make_SQL_Stmt()
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

func GetQuery() (QueryForm, error) {
	var query QueryForm

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
