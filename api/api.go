package api

import (
	"fmt"
	"go-api/initial/my_db"
	"log"
	"strings"
)

func MakeQuery(base string, f *GlobalQueryFilters, target string) (string, []interface{}) {

	var (
		clauses []string
		args    []interface{}
	)

	base = strings.Replace(base, "?", target, 1)
	filters := filter_process(f)

	for filter, filterVals := range filters {
		if filterVals != nil {
			clauses = append(clauses, filter+" IN ("+strings.Repeat("?,", len(filterVals)-1)+"?)")
			for _, vals := range filterVals {
				args = append(args, vals)
			}
		}
	}

	if len(clauses) != 0 {
		base += " WHERE " + strings.Join(clauses, " AND ")
	}

	return base, args
}

func ListUniquePlayers(query string, target string, args []interface{}) []string {
	var results []string

	db, err := my_db.GetDB()
	if err != nil {
		log.Fatal(err)
	}

	db.Raw(query, args...).Pluck(target, &results)
	return results
}

func GetGroupedBarData(q QueryForm) {

	base1 := my_db.GetPlayerQueries()[my_db.PlayersListDistinct]

	query1, args1 := MakeQuery(base1, q.Global_Filters, q.Graph_Params.X_target)
	query2, args2 := MakeQuery(base1, q.Global_Filters, q.Graph_Params.X2_target)

	res1 := ListUniquePlayers(query1, q.Graph_Params.X_target, args1)
	res2 := ListUniquePlayers(query2, q.Graph_Params.X2_target, args2)

	fmt.Println("done")
	_, _ = res1, res2
}

func ListPlayers(query string, args []interface{}) []*my_db.Player {
	var results []*my_db.Player
	db, err := my_db.GetDB()
	if err != nil {
		log.Fatal(err)
	}
	db.Raw(query, args...).Scan(&results)
	if len(results) == 0 {
		return nil
	}
	return results
}
