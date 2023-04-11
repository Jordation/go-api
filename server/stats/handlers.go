package stats

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func (A *StatsAPI) HandleListStats(w http.ResponseWriter, r *http.Request) {
	var query ListStatsRequest
	json.NewDecoder(r.Body).Decode(&query)
	res, err := ListStats(&query, A.DB)
	if err != nil {
		json.NewEncoder(w).Encode("{'Err': '" + err.Error() + "+'}")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func (A *StatsAPI) HandleListUniqueStats(w http.ResponseWriter, r *http.Request) {
	var query ListStatsRequest
	json.NewDecoder(r.Body).Decode(&query)
	if query.TargetRow == "" {
		json.NewEncoder(w).Encode("{'err': 'Error: This query requires a target'}")
		return
	}
	res, err := ListUniqueStats(&query, nil, A.DB)
	if err != nil {
		json.NewEncoder(w).Encode("{'err': 'Error: " + err.Error() + "'}")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func (A *StatsAPI) HandleListUnique(w http.ResponseWriter, r *http.Request) {
	targ := mux.Vars(r)["target"]
	log.Info("[LIST UNIQUE]: target: ", targ)
	res, err := ListUnique(&ListStatsRequest{TargetRow: targ}, A.DB)
	if err != nil {
		fmt.Println(err)
	}
	log.Info("ListUnique: ", res)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
