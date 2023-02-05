package GetGroupedBarData

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
	a "go-api/initial/api"
	"log"
)

func CreateGroups(gc [][]string) [][2]string {

	var groups [][2]string

	for _, val := range gc[0] {
		for _, val2 := range gc[1] {
			group := [2]string{val, val2}
			groups = append(groups, group)
		}
	}

	return groups
}

func RowSatisfiesGroup(g [2]string, row a.PlayerStatsResult, targs []string) bool {
	var (
		satisfies [2]bool
	)

	for i, v := range targs {
		switch v {
		case "player":
			if row.Player == g[i] {
				satisfies[i] = true
			} else {
				return false
			}
		case "agent":
			if row.Agent == g[i] {
				satisfies[i] = true
			} else {
				return false
			}
		case "mapname":
			if row.Mapname == g[i] {
				satisfies[i] = true
			} else {
				return false
			}
		case "team":
			if row.Team == g[i] {
				satisfies[i] = true
			} else {
				return false
			}
		default:
			return false
		}
	}
	return satisfies[0] == satisfies[1]
}

func FillGroupsWithRows(grps [][2]string, rows []a.PlayerStatsResult, targs []string) map[[2]string][]a.PlayerStatsResult {
	filledGroups := make(map[[2]string][]a.PlayerStatsResult)

	for _, group := range grps {
		for _, row := range rows {
			if RowSatisfiesGroup(group, row, targs) {
				filledGroups[group] = append(filledGroups[group], row)
			}
		}
	}
	return filledGroups
}

func AverageResultsOverGroup(groups map[[2]string][]a.PlayerStatsResult) map[[2]string]a.PlayerStatsResult {
	var newgroups = make(map[[2]string]a.PlayerStatsResult)

	for k, v := range groups {
		c := 0
		var avg_data a.PlayerStatsResult
		for _, v2 := range v {
			avg_data.JoinResult(v2)
			c++
		}
		avg_data.AvgResult(c)

		newgroups[k] = avg_data
	}

	return newgroups
}

func TakeHighestResult(groups map[[2]string][]a.PlayerStatsResult, q a.GraphParams) map[[2]string]a.PlayerStatsResult {
	var newgroups = make(map[[2]string]a.PlayerStatsResult)
	for k, v := range groups {
		max := 0
		for _, v2 := range v {
			if v2.FindValue(q.Y_target) > max {
				max = v2.FindValue(q.Y_target)
				newgroups[k] = v2
			}
		}
	}
	return newgroups
}

func ProcessGroups(groups map[[2]string][]a.PlayerStatsResult, q a.QueryForm) map[[2]string]a.PlayerStatsResult {
	var newgroups map[[2]string]a.PlayerStatsResult
	for k, v := range groups {
		if len(v) < q.Data_Params.Min_dataset_size {
			delete(groups, k)
		}
	}

	if q.Data_Params.Average_over_groups {
		newgroups = AverageResultsOverGroup(groups)
	} else {
		newgroups = TakeHighestResult(groups, *q.Graph_Params)
	}

	// create chart js labels, formatting
	return newgroups
}

func GetGroupedBarData(q a.QueryForm) error {
	var (
		filters a.ListPlayerStatsFilters
	)

	filters.Columns = []string{q.Graph_Params.X_target, q.Graph_Params.X2_target}
	filters.Unique = true
	rows, cols, err := a.ListPlayerStats(filters, q)

	if err != nil {
		log.Fatal(err)
		return err
	}

	_ = rows

	groups := CreateGroups(cols)
	filledGroups := FillGroupsWithRows(groups, rows, filters.Columns)
	finalisedDatasets := ProcessGroups(filledGroups, q)

	_ = finalisedDatasets
	return nil
}
