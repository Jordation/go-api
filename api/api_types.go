package api

import (
	"encoding/json"
	"fmt"
	"strings"
)

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
func filter_process(f *GlobalQueryFilters) (map[string][]string, string) {
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

// Single usecase get string from interface to satisfy type checker lol
func StringFromInterface(i interface{}) string {
	newstr := `"` + i.(string) + `"`
	return newstr
}

// makes sql statement from query filters
func (f *GlobalQueryFilters) MakeSQLStmt(JUST_STMT bool) (string, []interface{}, error) {
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

	if JUST_STMT {
		for _, arg := range args {
			stmt = strings.Replace(stmt, "?", StringFromInterface(arg), 1)
		}
		return stmt, nil, nil
	}

	return stmt, args, nil
}

type GlobalQueryFilters struct {
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
	Global_Filters *GlobalQueryFilters
	Data_Params    *DataParams
	Graph_Params   *GraphParams
}

type PlayerStatsResult struct {
	ID      int
	Player  string
	Agent   string
	Team    string
	Acs     int
	K       int
	D       int
	A       int
	Kast    int
	Adr     int
	Hsp     int
	Fb      int
	Fd      int
	Map_id  int
	Mapname string
}

type ColumnResult struct {
}

type ListPlayerStatsFilters struct {
	Unique  bool
	Columns []string
}
