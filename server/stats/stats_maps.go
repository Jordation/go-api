package stats

import (
	"fmt"
	"strings"
	"time"

	"github.com/Jordation/go-api/server/db"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

const (
	ListMapsStmt   = "SELECT * FROM maps"
	ListMapIdsStmt = "SELECT id FROM maps"
	ListDateRange  = "SELECT * FROM maps maps ORDER BY match_date"
)

func GetMapDateRange(dbConn *gorm.DB) (time.Time, time.Time, error) {
	var firstMap, lastMap db.Map
	log.Info("[GET MAP DATE RANGE]")
	if err := dbConn.Raw(ListDateRange + " ASC LIMIT 1").Scan(&firstMap).Error; err != nil {
		return firstMap.MatchDate, lastMap.MatchDate, err
	}
	if err := dbConn.Raw(ListDateRange + " DESC LIMIT 1").Scan(&lastMap).Error; err != nil {
		return firstMap.MatchDate, lastMap.MatchDate, err
	}
	return firstMap.MatchDate, lastMap.MatchDate, nil
}

func ListMapIDS(f *ListMapsRequest, dbConn *gorm.DB) ([]uint, error) {
	var res []db.Map
	stmt, args := GetMapsQuery(f, ListMapsStmt)
	log.Info("[LIST MAP IDS]: ", stmt, args)
	if err := dbConn.Raw(stmt, args...).Scan(&res).Error; err != nil {
		return nil, err
	}
	ids := make([]uint, 0)
	for _, v := range res {
		ids = append(ids, v.ID)
	}

	return ids, nil
}

func ListMaps(f *ListMapsRequest, dbConn *gorm.DB, res *[]db.Map) error {
	stmt, args := GetMapsQuery(f, ListMapsStmt)
	log.Info("[LIST MAPS]: ", stmt, args)
	if err := dbConn.Raw(stmt, args...).Scan(&res).Error; err != nil {
		return err
	}
	return nil
}

func makeDateTimeClause(start, end string, solo bool) string {
	if solo {
		return fmt.Sprintf(" WHERE match_date BETWEEN date('%v') AND date('%v')", start, end)
	} else {
		return fmt.Sprintf(" AND match_date BETWEEN date('%v') AND date('%v')", start, end)
	}
}

// experimenting with gorm functionality more... bite me in asS later? maybe.
func GetMapsQuery(f *ListMapsRequest, stmt string) (string, []interface{}) {
	var (
		clauses []string
		args    []interface{}
	)

	mf := MakeMappedMapFilter(f)

	if f.INfilters != nil {
		for filter, filterVals := range mf["IS"] {
			if len(filterVals) != 0 {
				if filter == "team" {
					curr_team := filter + "1"
					clauses = append(clauses, curr_team+" IN ("+strings.Repeat("?, ", len(filterVals)-1)+"?)")
					curr_team = filter + "2"
					clauses = append(clauses, curr_team+" IN ("+strings.Repeat("?, ", len(filterVals)-1)+"?)")
					for _, v := range filterVals {
						args = append(args, v, v)
					}
				} else {
					clauses = append(clauses, filter+" IN ("+strings.Repeat("?, ", len(filterVals)-1)+"?)")
					for _, v := range filterVals {
						args = append(args, v)
					}
				}
			}
		}
	}
	if f.NOTINfilters != nil {
		for filter, filterVals := range mf["NOT"] {
			if len(filterVals) != 0 {
				if filter == "team" {
					curr_team := filter + "1"
					clauses = append(clauses, curr_team+" IN ("+strings.Repeat("?, ", len(filterVals)-1)+"?)")
					curr_team = filter + "2"
					clauses = append(clauses, curr_team+" IN ("+strings.Repeat("?, ", len(filterVals)-1)+"?)")
					for _, v := range filterVals {
						args = append(args, v, v)
					}
				} else {
					clauses = append(clauses, filter+" NOT IN ("+strings.Repeat("?, ", len(filterVals)-1)+"?)")
					for _, v := range filterVals {
						args = append(args, v)
					}
				}
			}
		}
	}

	if len(clauses) != 0 {
		stmt += " WHERE " + strings.Join(clauses, " AND ")
		if f.Start != "" && f.End != "" {
			stmt += makeDateTimeClause(f.Start, f.End, false)
		}
	} else {
		if f.Start != "" && f.End != "" {
			stmt += makeDateTimeClause(f.Start, f.End, true)
		}
	}

	log.Info("[GET MAPS Q, ARGS]: ", stmt, args)
	return stmt, args
}

func MapMapFilter(f *MapsFilter) map[string][]string {
	m := make(map[string][]string)
	m["map"] = f.Maps
	m["team"] = f.Teams
	return m
}

func MakeMappedMapFilter(q *ListMapsRequest) StatsFilterMap {
	MappedFilters := make(StatsFilterMap, 2)

	if q.INfilters != nil {
		MappedFilters["IS"] = MapMapFilter(q.INfilters)
	}

	if q.NOTINfilters != nil {
		MappedFilters["NOT"] = MapMapFilter(q.NOTINfilters)
	}

	return MappedFilters
}
