package server

import (
	"encoding/json"
	"fmt"
	"go-api/api"
	"go-api/api/stats"
	"go-api/orm"
	"go-api/server/graphs"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type ListUniqueResp struct {
	Data []string `json:"data"`
}

func ReadQuery() (graphs.GroupedBarRequest, error) {
	var query graphs.GroupedBarRequest
	data, err := os.ReadFile("query.json")
	if err != nil {
		return query, err
	}

	json.Unmarshal(data, &query)

	return query, nil
}

func getGroupedBarData(w http.ResponseWriter, r *http.Request) {

}

func HandleListUniqueStatsRequest(w http.ResponseWriter, r *http.Request) {
	t := mux.Vars(r)["target"]
	fmt.Println(t)
	q := orm.GetStatQueries()[orm.PlayersListDistinct]

	f := api.ListStatsFilter{Target: t, Query: q}
	list := stats.ListUniqueStats(f)

	w.Header().Set("Content-Type", "application/json")
	res := ListUniqueResp{Data: list}

	json.NewEncoder(w).Encode(res)
}

func StartServer(wg *sync.WaitGroup) {

	r := mux.NewRouter()
	c := cors.Default()
	r.HandleFunc("/graphs/groupedBar", getGroupedBarData).Methods("GET")
	r.HandleFunc("/ListUniqueStats/{target}", HandleListUniqueStatsRequest).Methods("GET")
	fmt.Println("Starting server on port 8000")
	if err := http.ListenAndServe(":8000", c.Handler(r)); err != nil {
		log.Fatal(err)
	}
	wg.Done()
}
