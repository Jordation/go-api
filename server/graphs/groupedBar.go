package graphs

import (
	"go-api/api"
	"go-api/api/stats"
	"go-api/orm"
	"sort"
	"strings"
)

func checkEmptyGroup(g ...float64) bool {
	var t float64
	for i := range g {
		t += g[i]
	}
	return t == 0
}
func sortEvalGroup(eg map[string]float64) map[string]float64 {
	keys := make([]string, 0, len(eg))

	for k := range eg { // get keys from map
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool { // sort map by key values
		return eg[keys[i]] > eg[keys[j]]
	})

	for i := 0; i < len(eg); i++ { // replace results after top 5 with 0
		if i > 4 {
			k := keys[i]
			eg[k] = 0
		}
	}
	return eg
}
func finalizeGroups(groups map[string][]float64, groupInds []string) map[string][]float64 {
	for i := range groupInds {
		evalGroup := make(map[string]float64)
		for k, v2 := range groups {
			evalGroup[k] = v2[i] // create flatenned map of index i from each group
		}
		evalGroup = sortEvalGroup(evalGroup) // sort eval group
		for k3, v3 := range evalGroup {
			groups[k3][i] = v3 // replace values
		}
	}
	for k, v := range groups {
		if checkEmptyGroup(v...) {
			delete(groups, k)
		}
	}
	return groups
}

func groupIterate(g1 []string, g2 []string, iq string, q GroupedBarRequest) (res GroupedBarResponse) {
	var query string
	if q.AverageResults {
		query = strings.Replace(orm.GetStatQueries()["listAVG"], "players", iq, 1)
	} else {
		query = strings.Replace(orm.GetStatQueries()["listMAX"], "players", iq, 1)
	}

	res.Data = make(map[string][]float64)

	for _, v := range g1 {
		res.Groups = append(res.Groups, v)
		for _, v2 := range g2 {
			f := api.ListStatsFilter{
				Query:              query,
				Target:             q.YTarget,
				Filters_IS:         api.ArgClauseMap{q.XTarget: []string{v}, q.XGroupsTarget: []string{v2}},
				MinimumDatasetSize: q.MinDatasetSize,
			}
			newres := stats.ListAverageStat(f)
			res.Data[v2] = append(res.Data[v2], newres)
		}
	}

	return res
}
func getGroupPreReqs(q GroupedBarRequest) ([]string, []string, string) {
	f_i := MapQueryFilters(q.Filters_IS)
	f_n := MapQueryFilters(q.Filters_NOT)
	f1 := api.ListStatsFilter{
		Query:       orm.GetStatQueries()["listDistinct"],
		Filters_IS:  f_i,
		Filters_NOT: f_n,
		Target:      q.XTarget,
	}
	f2 := api.ListStatsFilter{
		Query:       orm.GetStatQueries()["listDistinct"],
		Filters_IS:  f_i,
		Filters_NOT: f_n,
		Target:      q.XGroupsTarget,
	}
	f3 := api.ListStatsFilter{
		Query:       orm.GetStatQueries()["list"],
		Filters_IS:  f_i,
		Filters_NOT: f_n,
		Target:      "",
	}

	g1 := stats.ListUniqueStats(f1)       // list of unique values matching query results
	g2 := stats.ListUniqueStats(f2)       // secondary list of unique values matching query results
	innerStmt := stats.MakeInnerQuery(f3) // inner stmt used for refining on top of group reqs

	return g1, g2, innerStmt
}
func ProcessGroupedBarQuery(q GroupedBarRequest) (res GroupedBarResponse) {

	g1, g2, innerQuery := getGroupPreReqs(q)

	res = groupIterate(g1, g2, innerQuery, q)

	res.Data = finalizeGroups(res.Data, g1)

	return res
}
