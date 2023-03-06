package main

import (
	_ "encoding/json"
	"fmt"
	"sync"

	_ "go-api/api"
	_ "go-api/orm"
	"go-api/server"
	_ "go-api/server/scraper"
	_ "os"
)

func main() {
	var wg sync.WaitGroup
	//q, err := server.ReadQuery()
	//if err != nil {
	//	log.Println(err)
	//}
	//graphs.ProcessGroupedBarQuery(q)
	wg.Add(1)
	go server.StartServer(&wg)

	fmt.Println("I didnt get blocked")
	wg.Wait()
}
