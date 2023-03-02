package server

import (
	"encoding/json"
	"fmt"
	"go-api/api"
	"go-api/server/graphs"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/gorilla/mux"
)

type testRequest struct {
	BodyOfReq string `json:"q"`
}

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
	var data testRequest
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(data)
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
