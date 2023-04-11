package ports

import (
	"context"
	"net/http"
)

type DbPort interface {
}

type LoggerPort interface {
	LogToDB(e error, ctx context.Context) error
}

type StatsApiPort interface {
	HandleListStats(w http.ResponseWriter, r *http.Request)
	HandleListUniqueStats(w http.ResponseWriter, r *http.Request)
	HandleListUnique(w http.ResponseWriter, r *http.Request)
}

type GraphsApiPort interface {
	HandleTimescaleRequest(w http.ResponseWriter, r *http.Request)
	HandleGroupedBar(w http.ResponseWriter, r *http.Request)
}
