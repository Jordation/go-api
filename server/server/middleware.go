package server

import (
	"net/http"
	"refactor/db"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var LoggerDbConnection *gorm.DB

func init() { LoggerDbConnection = db.Getdb().DB }

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		log.Println("Request URI: ", r.RequestURI)
		log.Println("Requestor:", r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}
