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
	var query graphs.GroupedBarRequest
	json.NewDecoder(r.Body).Decode(&query)
	query.AverageResults = true
	data := graphs.ProcessGroupedBarQuery(query)
	json.NewEncoder(w).Encode(data)
}

func HandleListUniqueStatsRequest(w http.ResponseWriter, r *http.Request) {
	t := mux.Vars(r)["target"]
	fmt.Println(t)
	q := orm.GetStatQueries()[orm.PlayersListDistinct]
	f := api.ListStatsFilter{
		Target: t,
		Query:  q,
	}

	list := stats.ListUniqueStats(f)

	w.Header().Set("Content-Type", "application/json")
	res := ListUniqueResp{Data: list}

	json.NewEncoder(w).Encode(res)
}
func HandleListStatsRequest(w http.ResponseWriter, r *http.Request) {
	var query api.ListStatsRequest
	json.NewDecoder(r.Body).Decode(&query)
	f_i := *MapQueryFilters(query.Filters_IS)
	f := api.ListStatsFilter{
		Query:      orm.GetStatQueries()[orm.PlayersList],
		Filters_IS: f_i,
	}
	fmt.Println(f_i, f, query)
	res := stats.ListStats(f)
	json.NewEncoder(w).Encode(res)
}

func StartServer(wg *sync.WaitGroup) {

	r := mux.NewRouter()
	c := cors.Default()

	r.HandleFunc("/graphs/groupedBar", getGroupedBarData).Methods("POST")

	r.HandleFunc("/ListUniqueStats/{target}", HandleListUniqueStatsRequest).Methods("GET")

	r.HandleFunc("/ListStats", HandleListStatsRequest).Methods("POST")

	fmt.Println("Starting server on port 8000")
	if err := http.ListenAndServe(":8000", c.Handler(r)); err != nil {
		log.Fatal(err)
	}
	wg.Done()
}
