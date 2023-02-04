package api

import (
	"log"
	"reflect"

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

func ListPlayerStats(f ListPlayerStatsFilters, gf GlobalQueryFilters) (
	[]PlayerStatsResult, // results of rows
	[][]string, // results of columns
	error) {

	// Connect to DB
	db, err := gorm.Open(sqlite.Open("./my_db/test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return nil, nil, err
	}

	var (
		cols       = make(map[string][]PlayerStatsResult)
		results    []PlayerStatsResult
		groups     [][]string
		inner_stmt string
	)

	if gf != (GlobalQueryFilters{}) {
		inner_stmt, _, err = gf.MakeSQLStmt(true)
		if err != nil {
			log.Fatal(err)
			return nil, nil, err
		}
	} else {
		inner_stmt = "SELECT * FROM player_stats_combined"
	}

	// If no filters, return all rows
	if reflect.DeepEqual(f, ListPlayerStatsFilters{}) {
		var result []PlayerStatsResult
		db.Raw(inner_stmt).Scan(&result)
		return results, nil, nil
	}

	// If Reqiest specifies columns, return a grouped list of columns and rows
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
		stmt := "SELECT * FROM (" + inner_stmt + ")"
		db.Raw(stmt).Scan(&results)
		groups = GetWordsForGroups(cols)
		return results, groups, nil
	}
	return nil, nil, nil
}
