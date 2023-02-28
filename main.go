package main

import (
	_ "encoding/json"
	_ "fmt"
	_ "go-api/api"
	_ "go-api/orm"
	"go-api/server"
	_ "go-api/server/graphs"
	_ "go-api/server/scraper"
	_ "log"
	_ "os"
)

func main() {
	// q, err := server.ReadQuery()
	// if err != nil {
	// 	log.Println(err)
	// }
	// graphs.GetGroupedBarData(q)
	server.StartServer()
}
