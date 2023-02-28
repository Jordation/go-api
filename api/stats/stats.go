package stats

import (
	"fmt"
	"go-api/api"
	"go-api/orm"
	"strings"
)

func MakeInnerQuery(f api.ListStatsFilter) string {
	query, args := MakeQuery(f)

	for _, v := range args {
		str := fmt.Sprintf("\"%v\"", v)
		query = strings.Replace(query, "?", str, 1)
	}

	return "(" + query + ")"
}
func MakeQuery(f api.ListStatsFilter) (string, []interface{}) {
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
	query := strings.Replace(orm.GetPlayerQueries()["listAVG"], "players", iq, 1)

	for _, v := range g1 {
		x_titles = append(x_titles, v)
		for _, v2 := range g2 {
			f := api.ListStatsFilter{
				Query:   query,
				Target:  "kills",
				Filters: api.ArgClauseMap{xt: []string{v}, xg: []string{v2}},
			}
			newres := ListAverageStat(f)
			groups[v2] = append(groups[v2], newres)
		}
	}

	fmt.Println(groups)
}

func GetGroupedBarData(q api.QueryForm) {

	filters := api.MapQueryFilters(q.Global_Filters)
	x_target := q.Graph_Params.X_target
	x_grouping := q.Graph_Params.X2_target

	f1 := api.ListStatsFilter{
		Query:   orm.GetPlayerQueries()["listDistinct"],
		Filters: filters,
		Target:  x_target,
	}
	f2 := api.ListStatsFilter{
		Query:   orm.GetPlayerQueries()["listDistinct"],
		Filters: filters,
		Target:  x_grouping,
	}

	res1 := ListUniqueStats(f1)
	res2 := ListUniqueStats(f2)

	f3 := api.ListStatsFilter{
		Query:   orm.GetPlayerQueries()["list"],
		Filters: filters,
		Target:  "",
	}

	innerQuery := MakeInnerQuery(f3)

	GroupIterate(res1, res2, innerQuery, x_target, x_grouping)

	fmt.Println(innerQuery)
	_, _ = res1, res2
}
func ListUniqueStats(f api.ListStatsFilter) (res []string) {

	query, args := MakeQuery(f)

	db, err := orm.GetDB()
	if err != nil {
		panic(err)
	}

	db.Raw(query, args...).Pluck(f.Target, &res)
	return res
}

func ListAverageStat(f api.ListStatsFilter) (res float64) {

	query, args := MakeQuery(f)

	db, err := orm.GetDB()
	if err != nil {
		panic(err)
	}

	db.Raw(query, args...).Scan(&res)
	return res
}

func ListStats(f api.ListStatsFilter) (res api.ListStatsResponse) {
	query, args := MakeQuery(f)

	db, err := orm.GetDB()
	if err != nil {
		panic(err)
	}

	db.Raw(query, args...).Scan(&res.Stats)
	return res
}
