package api

import (
	"encoding/json"
	"fmt"
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

// Single usecase get string and add quotes
func StringFromInterface(i interface{}) string {
	return `"` + i.(string) + `"`
}

// Single usecase get float64
func Float64FromInterface(i interface{}) float64 {
	return i.(float64)
}

// Single usecase get int64
func Int64FromInterface(i interface{}) int64 {
	return i.(int64)
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

func (r *PlayerStatsResult) JoinResult(r2 PlayerStatsResult) {
	r.Player = r2.Player
	r.Team = r2.Team
	r.Mapname = r2.Mapname
	r.Agent = r2.Agent
	r.Acs += r2.Acs
	r.K += r2.K
	r.D += r2.D
	r.A += r2.A
	r.Kast += r2.Kast
	r.Adr += r2.Adr
	r.Hsp += r2.Hsp
	r.Fb += r2.Fb
	r.Fd += r2.Fd
}

func (r *PlayerStatsResult) AvgResult(c float32) {
	r.Acs /= c
	r.K /= c
	r.D /= c
	r.A /= c
	r.Kast /= c
	r.Adr /= c
	r.Hsp /= c
	r.Fb /= c
	r.Fd /= c
}

func (r *PlayerStatsResult) FindValue(k string) float32 {
	switch k {
	case "k":
		return r.K
	case "d":
		return r.D
	case "a":
		return r.A
	case "kast":
		return r.Kast
	case "acs":
		return r.Acs
	case "adr":
		return r.Adr
	case "hsp":
		return r.Hsp
	case "fd":
		return r.Fd
	case "fb":
		return r.Fb
	default:
		return 0
	}
}

type PlayerStatsResult struct {
	ID      int
	Player  string
	Agent   string
	Team    string
	Acs     float32
	K       float32
	D       float32
	A       float32
	Kast    float32
	Adr     float32
	Hsp     float32
	Fb      float32
	Fd      float32
	Map_id  int
	Mapname string
}

type ColumnResult struct {
}

type ListPlayerStatsFilters struct {
	Unique  bool
	Columns []string
}

type GetRowsResultFilter struct {
	Avg bool
	Max bool
	All bool
}

type GetRowsAsGroupsFilters struct {
	Columns     [2]string
	Filters     *GlobalQueryFilters
	Groups      [][2]string
	ResultType  GetRowsResultFilter
	Y_target    string
	Min_ds_size int
}
