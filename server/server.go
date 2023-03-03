package server

import (
	"encoding/json"
	"fmt"
	"go-api/server/graphs"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/gorilla/mux"
)

func ReadQuery() (graphs.GroupedBarRequest, error) {
	var query graphs.GroupedBarRequest
	data, err := os.ReadFile("query.json")
	if err != nil {
		return query, err
	}

	json.Unmarshal(data, &query)

	return query, nil
}

func getGroupedBar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	q, err := ReadQuery()
	if err != nil {
		fmt.Println(err.Error())
	}

	data := graphs.GetGroupedBarData(q)

	//res, err := json.Marshal(data)
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	json.NewEncoder(w).Encode(data)
}

func testing(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	query := mux.Vars(r)["q"]
	fmt.Println(query)
	json.NewEncoder(w).Encode("it worked!")
}

func StartServer(wg *sync.WaitGroup) {

	r := mux.NewRouter()

	r.HandleFunc("/graphs/groupedBar", getGroupedBar).Methods("GET")
	r.HandleFunc("/testing", testing).Methods("POST")

	fmt.Println("Starting server on port 8000")
	if err := http.ListenAndServe(":8000", r); err != nil {
		log.Fatal(err)
	}
}
