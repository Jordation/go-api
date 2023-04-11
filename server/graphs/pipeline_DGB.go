package graphs

import (
	"refactor/db"
	"refactor/helpers"
	"refactor/stats"
	"sync"
)

// server concurrency...
// channels for each type of request

func getDgbQueries(req *DgbRequest) (*DgbQueryGroups, error) {
	Xvals, err := getUniqueValues(req, req.X_Target)
	if err != nil {
		return nil, err
	}
	XGvals, err := getUniqueValues(req, req.X_Grouping)
	if err != nil {
		return nil, err
	}
	return &DgbQueryGroups{
		X_Targets: Xvals,
		GroupVals: XGvals,
	}, nil
}

func executeDgbQueries(barReq *DgbRequest, queries *DgbQueryGroups) (<-chan DgbReqDataset, error) {
	resChan := make(chan DgbReqDataset)
	db := db.Getdb().DB
	go func() {
		for i, v := range queries.X_Targets {
			req, hc := gbrToStatsRequest(barReq, v, queries.GroupVals)
			stat, err := stats.GetAvgStat(&req, &hc, db)
			if err != nil {
				return
			}
			resChan <- DgbReqDataset{
				X_Position: v,
				Data:       stat,
				Ind:        i,
			}
		}
		close(resChan)
	}()
	return resChan, nil
}

func insertvals(ds DgbReqDataset, dest *map[string]DgbResDataset, max int, wg *sync.WaitGroup) {
	defer wg.Done()
	if len(ds.Data) == 0 {
		return
	}
	for k, v := range ds.Data {
		i, ok := (*dest)[k]
		if ok {
			i.Data[ds.Ind] = helpers.RoundFloat(v)
		} else {
			vals := make([]interface{}, max)
			vals[ds.Ind] = helpers.RoundFloat(v)
			newDataset := DgbResDataset{
				Id:              len(*dest),
				Label:           k,
				Data:            vals,
				BackgroundColor: randColor(),
			}
			(*dest)[k] = newDataset
		}
	}
}

func makeDgbChartDatasets(datasets <-chan DgbReqDataset, XvalCnt, MaxDsCnt int) []DgbResDataset {
	var (
		chartDatasets = make(map[string]DgbResDataset)
		res           = make([]DgbResDataset, 0)
		wg            sync.WaitGroup
	)
	for ds := range datasets {
		wg.Add(1)
		ds.Data = trimDataset(ds.Data, MaxDsCnt)
		go insertvals(ds, &chartDatasets, XvalCnt, &wg)
	}
	wg.Wait()
	for _, v := range chartDatasets {
		res = append(res, v)
	}
	return res
}
