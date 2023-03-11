package stats

import (
	"fmt"
	"go-api/api"
	"go-api/orm"
	"math"
	"strings"
)

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}
func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

func MakeInnerQuery(f api.ListStatsFilter) string {
	query, args := MakeQuery(f)

	for _, v := range args {
		str := fmt.Sprintf("\"%v\"", v)
		query = strings.Replace(query, "?", str, 1)
	}

	return "(" + query + ")"
}
func AddHavingClause(g []string) string {
	clause := "GROUP BY (? AND ?) HAVING COUNT(*) > ?"
	for _, v := range g {
		clause = strings.Replace(clause, "?", v, 1)
	}
	return clause
}

func MakeQuery(f api.ListStatsFilter) (string, []interface{}) {
	var (
		clauses []string
		args    []interface{}
	)

	if f.Target != "" {
		f.Query = strings.Replace(f.Query, "?", f.Target, 1)
	}

	for filter, filterVals := range f.Filters_IS {
		if len(filterVals) != 0 {
			clauses = append(clauses, filter+" IN ("+strings.Repeat("?,", len(filterVals)-1)+"?)")
			for _, vals := range filterVals {
				args = append(args, vals)
			}
		}
	}

	for filter, filterVals := range f.Filters_NOT {
		if len(filterVals) != 0 {
			clauses = append(clauses, filter+" NOT IN ("+strings.Repeat("?,", len(filterVals)-1)+"?)")
			for _, vals := range filterVals {
				args = append(args, vals)
			}
		}
	}

	if len(clauses) != 0 {
		f.Query += " WHERE " + strings.Join(clauses, " AND ")
	}

	if f.MinimumDatasetSize != 0 {
		args = append(args, f.MinimumDatasetSize)
		var gs []string
		if len(f.Filters_IS) == 2 {
			for k := range f.Filters_IS {
				gs = append(gs, k)
			}
		}
		f.Query += AddHavingClause(gs)
	}

	return f.Query, args
}

func ListUniqueStats(f api.ListStatsFilter) (res []string) {

	query, args := MakeQuery(f)

	db, err := orm.GetDB()
	if err != nil {
		panic(err)
	}

	db.Raw(query, args...).Pluck(f.Target, &res)
	return res
}

func ListAverageStat(f api.ListStatsFilter) (res float64) {

	query, args := MakeQuery(f)

	db, err := orm.GetDB()
	if err != nil {
		panic(err)
	}

	db.Raw(query, args...).Scan(&res)
	return toFixed(res, 2)
}

func ListStats(f api.ListStatsFilter) (res api.ListStatsResponse2) {
	query, args := MakeQuery(f)

	db, err := orm.GetDB()
	if err != nil {
		panic(err)
	}

	db.Raw(query, args...).Scan(&res.Stats)
	return res
}
