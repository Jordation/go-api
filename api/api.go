package api

import (
	"fmt"
	"go-api/initial/my_db"
	"strings"
)

func MakeInnerQuery(f ListPlayersFilter) string {
	query, args := MakeQuery(f)

	for _, v := range args {
		str := fmt.Sprintf("\"%v\"", v)
		query = strings.Replace(query, "?", str, 1)
	}

	return "(" + query + ")"
}
func MakeQuery(f ListPlayersFilter) (string, []interface{}) {
	var (
		clauses []string
		args    []interface{}
	)

	if f.Target != "" {
		f.Query = strings.Replace(f.Query, "?", f.Target, 1)
	}

	for filter, filterVals := range f.Filters {
		if filterVals != nil {
			clauses = append(clauses, filter+" IN ("+strings.Repeat("?,", len(filterVals)-1)+"?)")
			for _, vals := range filterVals {
				args = append(args, vals)
			}
		}
	}

	if len(clauses) != 0 {
		f.Query += " WHERE " + strings.Join(clauses, " AND ")
	}

	return f.Query, args
}

func GroupIterate(g1 []string, g2 []string, iq string, xt string, xg string) {
	groups := make(map[string][]float64)
	x_titles := []string{"x"}
	query := strings.Replace(my_db.GetPlayerQueries()["listAVG"], "players", iq, 1)

	for _, v := range g1 {
		x_titles = append(x_titles, v)
		for _, v2 := range g2 {
			f := ListPlayersFilter{
				Query:   query,
				Target:  "kills",
				Filters: ArgClauseMap{xt: []string{v}, xg: []string{v2}},
			}
			newres := ListAveragePlayers(f)
			groups[v2] = append(groups[v2], newres)
		}
	}

	fmt.Println(groups)
}

func GetGroupedBarData(q QueryForm) {

	filters := mapQueryFilters(q.Global_Filters)
	x_target := q.Graph_Params.X_target
	x_grouping := q.Graph_Params.X2_target

	f1 := ListPlayersFilter{
		Query:   my_db.GetPlayerQueries()["listDistinct"],
		Filters: filters,
		Target:  x_target,
	}
	f2 := ListPlayersFilter{
		Query:   my_db.GetPlayerQueries()["listDistinct"],
		Filters: filters,
		Target:  x_grouping,
	}

	res1 := ListUniquePlayers(f1)
	res2 := ListUniquePlayers(f2)

	f3 := ListPlayersFilter{
		Query:   my_db.GetPlayerQueries()["list"],
		Filters: filters,
		Target:  "",
	}

	innerQuery := MakeInnerQuery(f3)

	GroupIterate(res1, res2, innerQuery, x_target, x_grouping)

	fmt.Println(innerQuery)
	_, _ = res1, res2
}
func ListUniquePlayers(f ListPlayersFilter) []string {
	var res []string

	query, args := MakeQuery(f)

	db, err := my_db.GetDB()
	if err != nil {
		panic(err)
	}

	db.Raw(query, args...).Pluck(f.Target, &res)
	return res
}

func ListAveragePlayers(f ListPlayersFilter) float64 {
	var res float64

	query, args := MakeQuery(f)

	db, err := my_db.GetDB()
	if err != nil {
		panic(err)
	}

	db.Raw(query, args...).Scan(&res)
	return res
}

func ListPlayers(f ListPlayersFilter) ListPlayersResponse {
	var res ListPlayersResponse

	query, args := MakeQuery(f)

	db, err := my_db.GetDB()
	if err != nil {
		panic(err)
	}

	db.Raw(query, args...).Scan(&res.players)
	return res
}
