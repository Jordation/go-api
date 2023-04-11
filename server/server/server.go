package server

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"refactor/db"
	"refactor/graphs"
	"refactor/helpers"
	"refactor/stats"

	"refactor/ports"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
)

type Application struct {
	db        ports.DbPort
	statsAPI  ports.StatsApiPort
	graphsAPI ports.GraphsApiPort
}

func GetApplication(db *db.Adapter) *Application {
	return &Application{
		db:        db,
		statsAPI:  stats.GetStatsAPI(db),
		graphsAPI: graphs.GetGraphAPI(),
	}
}

func handlePremadeDgbReq(w http.ResponseWriter, r *http.Request) {
	qnum, ok := mux.Vars(r)["id"]
	if ok {
		file := helpers.GetQueryDir() + "bq" + qnum + ".json"
		data, err := os.ReadFile(file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		var q graphs.DgbRequest
		json.Unmarshal(data, &q)
		res, err := graphs.StartGroupedBarResponse(&q)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(res)
	} else {
		logrus.Error("No query number provided")
	}
}
func handlePremadeTsrReq(w http.ResponseWriter, r *http.Request) {
	qnum, ok := mux.Vars(r)["id"]
	if ok {
		file := helpers.GetQueryDir() + "TsrReq" + qnum + ".json"
		data, err := os.ReadFile(file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		var q graphs.TsrRequest
		json.Unmarshal(data, &q)
		res, err := graphs.StartTimescaleResponse(&q)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(res)
	} else {
		logrus.Error("No query number provided")
	}
}

func ConfigureServer(db *db.Adapter) *http.Server {

	a := GetApplication(db)
	c := cors.Default()
	r := mux.NewRouter()

	r.HandleFunc("/ListStats", a.statsAPI.HandleListStats).Methods("POST")
	r.HandleFunc("/ListUniqueStats", a.statsAPI.HandleListUniqueStats).Methods("POST")
	r.HandleFunc("/ListUniqueStats/{target}", a.statsAPI.HandleListUnique).Methods("GET")
	r.HandleFunc("/GetGroupedBar", a.graphsAPI.HandleGroupedBar).Methods("POST")
	r.HandleFunc("/GetGroupedBar/{id}", handlePremadeDgbReq).Methods("GET")
	r.HandleFunc("/GetTimescaleGraph", a.graphsAPI.HandleTimescaleRequest).Methods("GET")
	r.HandleFunc("/GetTimescaleGraph/{id}", handlePremadeTsrReq).Methods("GET")
	r.Use(loggingMiddleware)

	return &http.Server{
		Addr:         "0.0.0.0:8000",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      c.Handler(r),
	}
}
