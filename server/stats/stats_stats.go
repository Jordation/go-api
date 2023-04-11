package stats

import (
	"fmt"
	"refactor/db"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// Query with ? MUST offer a target value in query
const (
	StatsListQuery         = "SELECT * FROM player_stats"
	StatsListDistinctQuery = "SELECT DISTINCT ? FROM player_stats"
	StatsListMAXQuery      = "SELECT MAX(?) FROM player_stats"
	StatsListAVGQuery      = "SELECT AVG(?) FROM player_stats"
)

const ()

type StatsAPI struct {
	*db.Adapter
}

func ListStats(q *ListStatsRequest, dbConn *gorm.DB) (res []db.PlayerStat, err error) {
	stmt, args := GetStatsQuery(q, StatsListQuery, nil)
	log.Info("ListStats: ", stmt, args)
	if err = dbConn.Raw(stmt, args...).Scan(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}

func ListUnique(q *ListStatsRequest, dbConn *gorm.DB) (res []string, err error) {
	stmt, args := GetStatsQuery(q, StatsListDistinctQuery, nil)
	log.Info("ListUnique: ", stmt, args)
	if err = dbConn.Raw(stmt, args...).Scan(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}

func ListUniqueStats(q *ListStatsRequest, hq *HavingClause, dbConn *gorm.DB) (res []string, err error) {
	stmt, args := GetStatsQuery(q, StatsListDistinctQuery, hq)
	log.Info("ListUniqueStats: ", stmt, args)
	if err = dbConn.Raw(stmt, args...).Scan(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}

// function needs having clause

func GetAvgStat(q *ListStatsRequest, hq *HavingClause, dbConn *gorm.DB) (map[string]float64, error) {
	stmt, args := GetStatsQuery(q, StatsListAVGQuery, hq)
	log.Info("Getting Stat: ", stmt, args)
	res := make([]map[string]interface{}, 0)
	if err := dbConn.Raw(stmt, args...).Scan(&res).Error; err != nil {
		return nil, err
	}

	resKey := "AVG(" + q.TargetRow + ")"
	resMap := make(map[string]float64)

	// maybe learn gorm properly or the default db drivers so i dont have to deal with this
	for _, v := range res {
		nameData, ok := v[hq.Targets[0]]
		statData, alsoOk := v[resKey]
		if ok && alsoOk {
			dataStr := nameData.(string)
			dataFloat := *statData.(*interface{})
			if dataFloat.(float64) != 0 {
				resMap[dataStr] = dataFloat.(float64)
			}
		} else {
			return nil, fmt.Errorf("**KEY MISMATCH**")
		}
	}
	// SHIT

	return resMap, nil
}

func GetStatsQuery(q *ListStatsRequest, StmtBase string, HClause *HavingClause) (string, []interface{}) {
	var (
		stmt    string
		clauses []string
		args    []interface{}
	)

	// handle cases where a target value is required i.e. MAX, AVG
	stmt = StmtBase
	if strings.Contains(StmtBase, "?") {
		stmt = strings.Replace(stmt, "?", q.TargetRow, 1)

		// handle cases where more than 1 < n < * rows are needs
		if q.ExtraRows != nil {
			insInd := strings.Index(stmt, ")") + 1
			rows := ""
			for i := range q.ExtraRows {
				rows += (", " + q.ExtraRows[i])
			}
			stmt = stmt[:insInd] + rows + stmt[insInd:]
		}
	}

	mf := MakeMappedStatsFilter(q)

	for filter, filterVals := range mf["IS"] {
		if len(filterVals) != 0 {
			clauses = append(clauses, filter+" IN ("+strings.Repeat("?, ", len(filterVals)-1)+"?)")
			for _, v := range filterVals {
				args = append(args, v)
			}
		}
	}

	for filter, filterVals := range mf["NOT"] {
		if len(filterVals) != 0 {
			clauses = append(clauses, filter+" NOT IN ("+strings.Repeat("?, ", len(filterVals)-1)+"?)")
			for _, v := range filterVals {
				args = append(args, v)
			}
		}
	}

	if len(clauses) != 0 {
		stmt += " WHERE " + strings.Join(clauses, " AND ")
	}

	if HClause != nil {
		stmt += AddHavingClause(HClause)
	}

	return stmt, args
}

func AddHavingClause(HClause *HavingClause) string {
	base := "GROUP BY (" + strings.Repeat("? AND ", len(HClause.Targets)-1) + "?) HAVING COUNT(*) > ?"
	for _, v := range HClause.Targets {
		base = strings.Replace(base, "?", v, 1)
	}
	base = strings.Replace(base, "?", strconv.Itoa(HClause.Count), 1)
	return base
}
func MapStatsFilter(f *StatsFilter) map[string][]string {
	m := make(map[string][]string)
	m["agent"] = f.Agents
	m["map"] = f.Maps
	m["player"] = f.Players
	m["team"] = f.Teams
	return m
}

// maybe can refactor to a struct,interface for mapping the filters of any query
func MakeMappedStatsFilter(q *ListStatsRequest) StatsFilterMap {
	MappedFilters := make(StatsFilterMap, 2)

	if q.DoesContainFilters != nil {
		MappedFilters["IS"] = MapStatsFilter(q.DoesContainFilters)
	}

	if q.Side != "" {
		MappedFilters["IS"]["side"] = []string{q.Side}
	}

	if q.NotContainFilters != nil {
		MappedFilters["NOT"] = MapStatsFilter(q.NotContainFilters)
	}

	return MappedFilters
}

func GetStatsAPI(db *db.Adapter) *StatsAPI {
	return &StatsAPI{db}
}
