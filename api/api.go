package api

import (
	"log"
	"reflect"
	"strconv"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetWordsForGroups(r map[string][]PlayerStatsResult) [][]string {

	groups := make([][]string, len(r))
	c := 0
	for k, v := range r {
		switch k {
		case "player":
			for _, v := range v {
				groups[c] = append(groups[c], v.Player)
			}
		case "agent":
			for _, v := range v {
				groups[c] = append(groups[c], v.Agent)
			}
		case "mapname":
			for _, v := range v {
				groups[c] = append(groups[c], v.Mapname)
			}
		case "team":
			for _, v := range v {
				groups[c] = append(groups[c], v.Team)
			}
		}
		c++
	}
	return groups
}

func GetRowsAsGroups(f GetRowsAsGroupsFilters) (map[[2]string]interface{}, error) {
	var (
		groupedResults = make(map[[2]string]interface{})
		prefix         string
	)
	// Connect to DB
	db, err := gorm.Open(sqlite.Open("./my_db/test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	// get innter statement for group query
	inner_stmt, _, err := f.Filters.MakeSQLStmt(true)
	if err != nil {
		log.Fatal(err)
	}

	if f.ResultType.Avg {
		prefix = `AVG(` + f.Y_target + `)`
	}
	if f.ResultType.Max {
		prefix = f.Y_target
	}
	// for each group get matching rows
	for _, grp_vals := range f.Groups {
		var result map[string]interface{}
		stmt := `SELECT ` + prefix + ` FROM 
		(` + inner_stmt + `) WHERE `
		stmt += f.Columns[0] + ` IN ("` + grp_vals[0] + `") 
		AND ` + f.Columns[1] + ` IN ("` + grp_vals[1] + `") `
		if f.ResultType.Max {
			stmt += `GROUP BY ` + f.Y_target +
				` HAVING COUNT(*) >= ` + strconv.Itoa(f.Min_ds_size) +
				` ORDER BY ` + f.Y_target + ` DESC LIMIT 1`
		} else {
			stmt += `HAVING COUNT(*) >= ` + strconv.Itoa(f.Min_ds_size)
		}
		db.Raw(stmt).Scan(&result)
		if result[prefix] != nil {
			if f.ResultType.Avg {
				groupedResults[grp_vals] = result[prefix]
			} else {
				groupedResults[grp_vals] = result[prefix].(int64)
			}
		}
	}

	return groupedResults, nil
}

func ListPlayerStats(f ListPlayerStatsFilters, qf QueryForm) (
	[]PlayerStatsResult, // results of rows
	[][]string, // results of columns
	error) {

	var (
		cols       = make(map[string][]PlayerStatsResult)
		results    []PlayerStatsResult
		groups     [][]string
		inner_stmt string
	)

	// Connect to DB
	db, err := gorm.Open(sqlite.Open("./my_db/test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return nil, nil, err
	}
	// if no filters at all return all rows
	if *qf.Global_Filters != (GlobalQueryFilters{}) {
		inner_stmt, _, err = qf.Global_Filters.MakeSQLStmt(true)
		if err != nil {
			log.Fatal(err)
			return nil, nil, err
		}
	} else {
		inner_stmt = "SELECT * FROM player_stats_combined"
	}

	// If no filters, return all rows
	if reflect.DeepEqual(f, ListPlayerStatsFilters{}) {
		inner_stmt += " ORDER BY \"" + qf.Graph_Params.Y_target + "\" ASC"
		var result []PlayerStatsResult
		db.Raw(inner_stmt).Scan(&result)
		return results, nil, nil
	}

	// If request specifies columns, return a grouped list of columns and rows
	if f.Columns != nil {
		for _, col := range f.Columns {
			var col_result []PlayerStatsResult
			col_stmt := "SELECT "
			if f.Unique {
				col_stmt += "DISTINCT "
			}
			col_stmt += col + " FROM (" + inner_stmt + ")"
			db.Raw(col_stmt).Scan(&col_result)
			cols[col] = col_result
		}
		stmt := `SELECT * FROM ( ` + inner_stmt + `) 
		ORDER BY "` + qf.Graph_Params.Y_target + `" ASC`

		db.Raw(stmt).Scan(&results)
		groups = GetWordsForGroups(cols)
		return results, groups, nil
	}

	return nil, nil, nil
}
