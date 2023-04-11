package graphs

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"runtime"
	"sort"

	"github.com/Jordation/go-api/server/db"
	"github.com/Jordation/go-api/server/stats"

	"github.com/jinzhu/copier"
	log "github.com/sirupsen/logrus"
)

// random helpers
func randColor() string {
	return fmt.Sprintf("#%06x", rand.Intn(0xffffff))
}

func LogGroupedBarRequest(r *DgbRequest) {
	q := fmt.Sprintf("%v"+"\n", r)
	f1 := fmt.Sprintf("%v"+"\n", r.DoesContainFilters)
	f2 := fmt.Sprintf("%v", r.NotContainFilters)
	log.Info(log.WithFields(log.Fields{
		"IS Filters":        f1,
		"Groubed Bar Query": q,
		"NOT Filters":       f2,
	}))
}

type kv struct {
	key   string
	value float64
}

func ReadBarQuery() (DgbRequest, error) {
	var query DgbRequest
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	filePath := filepath.Join(dir, "bar_query.json")
	data, err := os.ReadFile(filePath)
	if err != nil {
		return query, err
	}
	json.Unmarshal(data, &query)
	return query, nil
}

//

// main functionality
func extendQueryFilters(XT string, XGT string, XTval string, XGTval []string, preIsF *stats.StatsFilter, preNotF *stats.StatsFilter) (stats.StatsFilter, stats.StatsFilter) {
	var (
		IsF, NotF stats.StatsFilter
	)
	copier.Copy(&IsF, preIsF)
	copier.Copy(&NotF, preNotF)
	switch XT {
	case "agent":
		IsF.Agents = make([]string, 0)
		IsF.Agents = append(IsF.Agents, XTval)
	case "map":
		IsF.Maps = make([]string, 0)
		IsF.Maps = append(IsF.Maps, XTval)
	case "player":
		IsF.Players = make([]string, 0)
		IsF.Players = append(IsF.Players, XTval)
	case "team":
		IsF.Teams = make([]string, 0)
		IsF.Teams = append(IsF.Teams, XTval)
	}

	switch XGT {
	case "agent":
		IsF.Agents = append(IsF.Agents, XGTval...)
	case "map":
		IsF.Maps = append(IsF.Maps, XGTval...)
	case "player":
		IsF.Players = append(IsF.Players, XGTval...)
	case "team":
		IsF.Teams = append(IsF.Teams, XGTval...)
	}

	return IsF, NotF
}

func gbrToStatsRequest(gbr *DgbRequest, XT string, XGTs []string) (stats.ListStatsRequest, stats.HavingClause) {
	IsF, NotF := extendQueryFilters(
		gbr.X_Target,
		gbr.X_Grouping,
		XT,
		XGTs,
		gbr.DoesContainFilters,
		gbr.NotContainFilters,
	)
	req := stats.ListStatsRequest{
		DoesContainFilters: &IsF,
		NotContainFilters:  &NotF,
		Side:               gbr.Side,
		TargetRow:          gbr.Y_Target,
		ExtraRows:          []string{gbr.X_Grouping},
	}
	havingClause := stats.HavingClause{
		Count:   gbr.MinDatasetSize,
		Targets: []string{gbr.X_Grouping},
	}
	return req, havingClause
}

func trimDataset(dataSet map[string]float64, maxDsCount int) map[string]float64 {
	if maxDsCount >= len(dataSet) {
		return dataSet
	}
	var kvs []kv
	for k, v := range dataSet {
		kvs = append(kvs, kv{k, v})
	}

	sort.Slice(kvs, func(i, j int) bool {
		return kvs[i].value > kvs[j].value
	})

	res := make(map[string]float64)
	for i := 0; i < maxDsCount; i++ {
		res[kvs[i].key] = kvs[i].value
	}

	return res
}

func getUniqueValues(baseReq *DgbRequest, TargRow string) ([]string, error) {
	var q stats.ListStatsRequest
	copier.Copy(&q, baseReq)
	db := db.Getdb().DB
	q.TargetRow = TargRow

	hc := stats.HavingClause{
		Targets: []string{TargRow},
		Count:   baseReq.MinDatasetSize,
	}
	res, err := stats.ListUniqueStats(&q, &hc, db)
	if err != nil {
		return nil, err
	}
	sort.Strings(res)
	return res, nil
}
