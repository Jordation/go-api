package main

import (
	"context"
	"flag"
	"fmt"

	"os"
	"os/signal"
	"time"

	"github.com/Jordation/go-api/server/db"
	"github.com/Jordation/go-api/server/server"

	log "github.com/sirupsen/logrus"
)

var wait time.Duration

func parseWait(s int, w *time.Duration) {
	flag.DurationVar(
		w,
		"graceful-timeout",
		time.Duration(s)*time.Second,
		"The time the server will wait for existing connections to close before shutting down")
	flag.Parse()
}

func main() {
	parseWait(15, &wait)
	stats_db := db.Getdb()
	server := server.ConfigureServer(stats_db)

	go func() {
		fmt.Println("Server starting on", server.Addr)
		if err := server.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	server.Shutdown(ctx)
	log.Println("shutting down")
	os.Exit(0)
}
