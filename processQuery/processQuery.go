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
	"fmt"
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

func GetGroupedBarData(q a.QueryForm) error {

	var (
		filters  a.ListPlayerStatsFilters
		filters2 a.GetRowsAsGroupsFilters
	)
	filters.Columns = []string{q.Graph_Params.X_target, q.Graph_Params.X2_target}
	filters.Unique = true

	_, cols, err := a.ListPlayerStats(filters, q)
	if err != nil {
		log.Fatal(err)
		return err
	}
	groups := CreateGroups(cols)
	filters2.Columns = [2]string{q.Graph_Params.X_target, q.Graph_Params.X2_target}
	filters2.Filters = q.Global_Filters
	filters2.Groups = groups
	filters2.ResultType.Max = true
	filters2.Y_target = q.Graph_Params.Y_target

	datasets, err := a.GetRowsAsGroups(filters2)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	_ = datasets
	return nil
}
