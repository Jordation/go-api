package server

import (
	"encoding/json"
	"fmt"
	"go-api/api"
	"go-api/server/graphs"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func ReadQuery() (api.QueryForm, error) {
	var query api.QueryForm
	bytes, err := os.ReadFile("query.json")
	if err != nil {
		return query, err
	}

	json.Unmarshal(bytes, &query.Global_Filters)
	json.Unmarshal(bytes, &query.Data_Params)
	json.Unmarshal(bytes, &query.Graph_Params)

	return query, nil
}

func getGroupedBar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// query := mux.Vars(r)["query"]
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
func StartServer() {

	r := mux.NewRouter()

	r.HandleFunc("/graphs/groupedBar", getGroupedBar).Methods("GET")

	if err := http.ListenAndServe(":8000", r); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Starting server on port 8000")
}
