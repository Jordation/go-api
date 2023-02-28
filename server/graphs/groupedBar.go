package graphs

import (
	"go-api/api"
	"go-api/api/stats"
	"go-api/orm"
	"strings"
)

func groupIterate(g1 []string, g2 []string, iq string, xt string, xg string) (res GroupedBarResponse) {
	query := strings.Replace(orm.GetStatQueries()["listAVG"], "players", iq, 1)

	res.Data = make(map[string][]float64)
	res.Groups = []string{"x"}

	for _, v := range g1 {
		res.Groups = append(res.Groups, v)
		for _, v2 := range g2 {
			f := api.ListStatsFilter{
				Query:              query,
				Target:             "kills",
				Filters:            api.ArgClauseMap{xt: []string{v}, xg: []string{v2}},
				MinimumDatasetSize: 5,
			}
			newres := stats.ListAverageStat(f)
			res.Data[v2] = append(res.Data[v2], newres)
		}
	}

	return res
}

func GetGroupedBarData(q api.QueryForm) (res GroupedBarResponse) {

	filters := api.MapQueryFilters(q.Global_Filters)
	x_target := q.Graph_Params.X_target
	x_grouping := q.Graph_Params.X2_target

	f1 := api.ListStatsFilter{
		Query:   orm.GetStatQueries()["listDistinct"],
		Filters: filters,
		Target:  x_target,
	}
	f2 := api.ListStatsFilter{
		Query:   orm.GetStatQueries()["listDistinct"],
		Filters: filters,
		Target:  x_grouping,
	}

	g1 := stats.ListUniqueStats(f1)
	g2 := stats.ListUniqueStats(f2)

	f3 := api.ListStatsFilter{
		Query:   orm.GetStatQueries()["list"],
		Filters: filters,
		Target:  "",
	}

	innerQuery := stats.MakeInnerQuery(f3)

	res = groupIterate(g1, g2, innerQuery, x_target, x_grouping)

	return res
}
