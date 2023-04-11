package graphs

import (
	"time"

	"github.com/Jordation/go-api/server/db"
	"github.com/Jordation/go-api/server/helpers"
	"github.com/Jordation/go-api/server/stats"

	"github.com/jinzhu/copier"
)

/* datasets: [
   {

     label: 'COMP NAME',
     data: [
       { x: new Date('2022-01-01'), y: 23, r: 10 },
       { x: new Date('2022-02-02'), y: 40, r: 120 },
       { x: new Date('2022-03-03'), y: 60, r: 30 },
       { x: new Date('2022-04-04'), y: 10, r: 55 },
     ],
     backgroundColor: 'rgba(255, 99, 132, 0.2)',
     borderColor: 'rgba(255, 99, 132, 1)',
     borderWidth: 1,
     pointRadius: function(context) {
       return context.dataset.data[context.dataIndex].r;
     },
     pointHoverRadius: function(context) {
       return context.dataset.data[context.dataIndex].r * 1.5;
     },
   }, */

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

func ProcessCompGroups(compGroups [][]stats.CompStats, dateGroups [][]string) map[string][]TsrDataPoint {
	res := make(map[string][]TsrDataPoint, 0)
	for i, v := range compGroups {
		total := GetTotalPicks(v)
		for _, v2 := range v {
			if v2.Ws == 0 || v2.Ls == 0 {
				continue
			}
			wr := helpers.GetPercent(float64(v2.Ws), float64(v2.Total()))
			pr := helpers.GetPercent(float64(v2.Total()), float64(total))
			date := dateGroups[i][0]

			res[v2.Comp] = append(res[v2.Comp], TsrDataPoint{
				X: date,
				Y: pr,
				R: wr,
			})
		}
	}
	return res
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
	datasets := ProcessCompGroups(compGroups, ranges)

	res := TsrResponse{
		Datasets: make([]TsrResDataset, 0),
	}
	makeTsrDatasets(datasets, &res)
	return &res, nil
}
