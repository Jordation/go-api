package graphs

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

type Request interface{}

type Result struct {
	Value interface{}
	Error error
}

func (A *GraphAPI) HandleGroupedBar(w http.ResponseWriter, r *http.Request) {
	var query DgbRequest
	json.NewDecoder(r.Body).Decode(&query)
	query.RequestID = uuid.New()

	res, err := StartGroupedBarResponse(&query)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	res.RequestID = query.RequestID

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
func StartGroupedBarResponse(req *DgbRequest) (*DgbResponse, error) {
	queries, err := getDgbQueries(req)
	if err != nil {
		return nil, err
	}
	datasetChan, err := executeDgbQueries(req, queries)
	if err != nil {
		return nil, err
	}
	res := makeDgbChartDatasets(datasetChan, len(queries.X_Targets), req.MaxDatasetsCount)

	return &DgbResponse{
		Labels:   queries.X_Targets,
		Datasets: res,
	}, nil
}

func (A *GraphAPI) HandleTimescaleRequest(w http.ResponseWriter, r *http.Request) {
	req := TsrRequest{}

	res, err := StartTimescaleResponse(&req)
	if err != nil {
		http.Error(w, "Err", 404)
	}
	json.NewEncoder(w).Encode(res)
}
func StartTimescaleResponse(req *TsrRequest) (*TsrResponse, error) {
	if err := req.SearchRange.ParseInputs(); err != nil {
		return nil, err
	}
	res, err := GetTsrQueries(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
