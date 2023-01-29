package main

import (
	"go-api/initial/processQuery"
)

func main() {

	// var result []db.Result
	// var query processQuery.GraphQuery

	// db, err := gorm.Open(sqlite.Open("./db/test.db"), &gorm.Config{})
	// if err != nil {
	// 	panic("failed to connect to db")
	// }

	// db.Raw("SELECT * FROM player_stats_combined;").Scan(&result)
	// fmt.Println(len(result))
	// err = processQuery.GroupRowsByTarget(query, result)

	// if err != nil {
	// 	return
	// }

	processQuery.ProcessQuery()
}
