package api

import (
	"encoding/json"
	"fmt"
	"go-api/orm"
	"strings"
)

type StrToBool bool

// converts any text in bool fields to true, empty str to false
func (f *StrToBool) UnmarshalJSON(b []byte) error {
	var s string
	fmt.Print(s)
	if err := json.Unmarshal(b, &s); err != nil {
		fmt.Print(err.Error())
	}

	if s != "" {
		*f = true
	} else {
		*f = false
	}
	return nil
}

// helper bc i dont want to work on my frontend any more
func MapQueryFilters(f *GlobalQueryFilters) map[string][]string {
	filters := make(map[string][]string)
	filters["agent"] = strings.Split(f.Agents, ", ")
	filters["map_name"] = strings.Split(f.Mapnames, ", ")
	filters["team"] = strings.Split(f.Teams, ", ")
	filters["player_name"] = strings.Split(f.Players, ", ")
	filters["side"] = []string{f.Side}
	for k, v := range filters {
		if v[0] == "" {
			filters[k] = nil
		}
	}
	return filters
}

// collaplse this SHIT
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
	Average_over_groups StrToBool `json:"average_rows_to_groups"`
	Order_by_y_target   StrToBool `json:"order_by_y_target"`
	Min_dataset_size    int       `json:"min_dataset_size,string"`
	Max_dataset_width   int       `json:"max_dataset_width,string"`
}
type QueryForm struct {
	Global_Filters *GlobalQueryFilters
	Data_Params    *DataParams
	Graph_Params   *GraphParams
}

type ArgClauseMap map[string][]string

type ListStatsResponse struct {
	Stats []orm.Player
}

type ListStatsFilter struct {
	//			  arg    : clause
	Filters            ArgClauseMap
	MinimumDatasetSize int
	Query              string
	Target             string
}
