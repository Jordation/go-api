package GetGroupedBarData

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

	// config for filters just done here for ease
	filters2.Columns = [2]string{ // the columns to target
		q.Graph_Params.X_target,
		q.Graph_Params.X2_target}

	filters2.Filters = q.Global_Filters                   // redundant but in here for ease of access, creates inner stmt
	filters2.Groups = groups                              // the sets of groups matching initial filters
	filters2.ResultType.Avg = true                        // avg or max (have to change to resulttype.Max if u want to check (might use a const enum)
	filters2.Y_target = q.Graph_Params.Y_target           // self exp.
	filters2.Min_ds_size = q.Data_Params.Min_dataset_size // self exp.

	datasets, err := a.GetRowsAsGroups(filters2)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	_ = datasets
	return nil
}
