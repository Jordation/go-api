package graphs

import (
	"sort"
	"time"

	"github.com/Jordation/go-api/server/db"
	"github.com/Jordation/go-api/server/helpers"
	"github.com/Jordation/go-api/server/stats"
	log "github.com/sirupsen/logrus"

	"github.com/jinzhu/copier"
)

const (
	PR_MULTI  = 0.4
	WR_MULTI  = 0.4
	APP_MULTI = 0.2
)

type tsrValue struct {
	comp  string
	value float64
}

func makeTsrDatasets(src map[string][]TsrDataPoint, dest *TsrResponse) {
	for k, v := range src {
		col := randColor()
		newds := TsrResDataset{
			Label:     k,
			Data:      v,
			BgCol:     col,
			BorderCol: col,
		}
		dest.Datasets = append(dest.Datasets, newds)
	}
}

func GetTotalPicks(group []stats.CompStats) (res int64) {
	for _, v := range group {
		res += v.Total()
	}
	return res
}

// for each slice of tsrDatapoint
// calculate score
// divide by len of slice
func CalculateDatasetWorth(ds *[]TsrDataPoint, scores *TsrScores) float64 {
	score := 0.0
	// this is troublesome, can't adjust which param is to radius and which is to Y axis.
	// will need to implement a map of some kind to keep track of what the use query specified as i'm using the values in multiple locations
	for _, v := range *ds {
		pr_score := (PR_MULTI * (v.R * scores.prAvg))
		wr_score := (WR_MULTI * (v.Y * scores.wrAvg))
		app_score := (APP_MULTI * (float64(v.Picks) / float64(scores.maxPicks)))
		score += pr_score + wr_score + app_score
	}
	return score
}
func EvaluateDatasetsWorth(betadata *map[string][]TsrDataPoint, maxResults int, scores TsrScores) []tsrValue {
	sortedData := make([]tsrValue, 0)
	for comp, data := range *betadata {
		sortedData = append(sortedData, tsrValue{
			comp:  comp,
			value: CalculateDatasetWorth(&data, &scores),
		})
	}
	sort.Slice(sortedData, func(i, j int) bool {
		return sortedData[i].value > sortedData[j].value
	})
	if maxResults > len(sortedData) {
		return sortedData
	}
	return sortedData[:maxResults]
}
func ProcessCompGroups(compGroups [][]stats.CompStats, dateGroups [][]string, maxResults int) map[string][]TsrDataPoint {
	var (
		unsorted_data             = make(map[string][]TsrDataPoint, 0)
		final_data                = make(map[string][]TsrDataPoint, 0)
		ttlPickrates, ttlWinrates []float64
		maxPicks                  int64
	)

	for i, v := range compGroups {
		total := GetTotalPicks(v)
		for _, v2 := range v {
			if v2.Ws == 0 || v2.Ls == 0 {
				continue
			}
			if v2.Total() > maxPicks {
				maxPicks = v2.Total()
			}
			date := dateGroups[i][0]
			wr := helpers.GetPercent(float64(v2.Ws), float64(v2.Total()))
			ttlWinrates = append(ttlWinrates, wr)
			pr := helpers.GetPercent(float64(v2.Total()), float64(total))
			ttlPickrates = append(ttlPickrates, pr)
			unsorted_data[v2.Comp] = append(unsorted_data[v2.Comp], TsrDataPoint{
				X:     date,
				Y:     wr,
				R:     pr,
				Picks: v2.Total(),
			})
		}
	}
	tsrScores := TsrScores{
		wrAvg:    helpers.SumNumSlice(ttlWinrates) / float64(len(ttlWinrates)),
		prAvg:    helpers.SumNumSlice(ttlPickrates) / float64(len(ttlPickrates)),
		maxPicks: maxPicks,
	}
	sortedData := EvaluateDatasetsWorth(&unsorted_data, maxResults, tsrScores)
	for _, dataset := range sortedData {
		final_data[dataset.comp] = unsorted_data[dataset.comp]
	}

	return final_data
}

func GetCompGroups(groups *[][]uint, compReq *stats.ListCompsRequest) ([][]stats.CompStats, error) {
	compGroups := make([][]stats.CompStats, 0)
	dbc := db.Getdb().DB
	for _, v := range *groups {
		compReq.MapIds = v
		res, err := stats.GetCompsWinratePickrate(dbc, compReq)
		if err != nil {
			return nil, err
		}
		_ = v
		compGroups = append(compGroups, res)
	}
	return compGroups, nil
}
func SplitTimes(divisions int, start, end time.Time) (res [][]string) {
	log.Info(divisions, start, end)
	var (
		duration = end.Sub(start)
		interval = duration / time.Duration(divisions)
	)
	for i := 0; i < divisions; i++ {
		grp := make([]string, 0, 2)
		date1 := start.Add(interval * time.Duration(i)).Format("2006-01-02")
		date2 := start.Add(interval * time.Duration(i+1)).Format("2006-01-02")
		grp = append(grp, date1, date2)
		res = append(res, grp)
	}
	return res
}
func GetDateRanges(times *TimeRange, divisions int) ([][]string, error) {
	db := db.Getdb().DB
	if times.IntDates != nil {
		res := SplitTimes(divisions, times.IntDates[0], times.IntDates[1])
		return res, nil
	}
	first, last, err := stats.GetMapDateRange(db)
	if err != nil {
		return nil, err
	}
	res := SplitTimes(divisions, first, last)
	return res, nil
}

func GetMapsInRange(req *TsrRequest, ranges [][]string, mapReq *stats.ListMapsRequest) ([][]uint, error) {
	dbc := db.Getdb().DB
	res := make([][]uint, 0)

	for _, v := range ranges {
		mapReq.Start = v[0]
		mapReq.End = v[1]
		ids, err := stats.ListMapIDS(mapReq, dbc)
		if err != nil {
			return nil, err
		}
		res = append(res, ids)
	}
	return res, nil
}

func GetTsrQueries(req *TsrRequest) (*TsrResponse, error) {
	var (
		MapReq  stats.ListMapsRequest
		CompReq stats.ListCompsRequest
		res     = TsrResponse{Datasets: make([]TsrResDataset, 0)}
	)
	copier.Copy(&MapReq, req)
	copier.Copy(&CompReq, req)

	ranges, err := GetDateRanges(&req.SearchRange, req.DateSplits)
	if err != nil {
		return nil, err
	}
	mapGroups, err := GetMapsInRange(req, ranges, &MapReq)
	if err != nil {
		return nil, err
	}
	compGroups, err := GetCompGroups(&mapGroups, &CompReq)
	if err != nil {
		return nil, err
	}
	datasets := ProcessCompGroups(compGroups, ranges, req.MaxResults)

	makeTsrDatasets(datasets, &res)
	return &res, nil
}
