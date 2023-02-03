package api

import (
	"fmt"
	"log"
	"reflect"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func MakeGroups(r map[string][]PlayerStatsResult) map[string][]string {

	groups := make(map[string][]string)

	for k, v := range r {
		switch k {
		case "player":
			for _, v := range v {
				groups[k] = append(groups[k], v.Player)
			}
		case "agent":
			for _, v := range v {
				groups[k] = append(groups[k], v.Agent)
			}
		case "mapname":
			for _, v := range v {
				groups[k] = append(groups[k], v.Mapname)
			}
		case "team":
			for _, v := range v {
				groups[k] = append(groups[k], v.Team)
			}
		}
	}
	return groups
}

func ListPlayerStats(f ListPlayerStatsFilters, gf GlobalQueryFilters) (
	map[string][]PlayerStatsResult,
	map[string][]string,
	error) {
	// Connect to DB
	db, err := gorm.Open(sqlite.Open("./my_db/test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return nil, nil, err
	}

	var (
		results    = make(map[string][]PlayerStatsResult)
		groups     map[string][]string
		inner_stmt string
	)

	if gf != (GlobalQueryFilters{}) {
		inner_stmt, _, err = gf.MakeSQLStmt(true)
		if err != nil {
			log.Fatal(err)
			return nil, nil, err
		}
	} else {
		fmt.Println("selecting all")
		inner_stmt = "SELECT * FROM player_stats_combined"
	}
	// Base var to store results

	// If no filters, return all rows
	if reflect.DeepEqual(f, ListPlayerStatsFilters{}) {
		var result []PlayerStatsResult
		db.Raw(inner_stmt).Scan(&result)
		results["All"] = result
		return results, nil, nil
	}

	// If Reqiest specifies columns, return a grouped list of columns
	if f.Columns != nil {
		for _, col := range f.Columns {
			var result []PlayerStatsResult
			stmt := "SELECT "
			if f.Unique {
				stmt += "DISTINCT "
			}
			stmt += col + " FROM (" + inner_stmt + ")"
			db.Raw(stmt).Scan(&result)
			results[col] = result
		}
		groups = MakeGroups(results)
		return nil, groups, nil
	}
	return nil, nil, nil
}
