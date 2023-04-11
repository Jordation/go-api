package graphs

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/Jordation/go-api/server/helpers"

	log "github.com/sirupsen/logrus"
)

func TestGetTimeScaleGraph(t *testing.T) {
	var q TsrRequest
	fpath := helpers.GetQueryDir() + "\\TsrReq1.json"

	data, err := os.ReadFile(fpath)
	if err != nil {
		t.Fatal(err)
	}

	json.Unmarshal(data, &q)
	if err := q.SearchRange.ParseInputs(); err != nil {
		t.Fatal(err)
	}
	GetTsrQueries(&q)
	fmt.Printf("%+v", q)
}

func TestPremadeQuaz(t *testing.T) {
	rType := "TsrReq"
	n := "1"
	dir := helpers.GetQueryDir()
	fpath := filepath.Join(dir, rType+n+".json")
	data, err := os.ReadFile(fpath)
	if err != nil {
		t.Fatal(err)
	}
	var query TsrRequest
	json.Unmarshal(data, &query)
	fmt.Printf("%+v", query)
}

func TestGetGroupedBarData(t *testing.T) {
	start := time.Now()
	q, err := ReadBarQuery()
	if err != nil {
		log.Fatal(err)
	}

	LogGroupedBarRequest(&q)

	data, err := StartGroupedBarResponse(&q)
	if err != nil {
		log.Fatal(err)
	}
	log.Info("Took", time.Since(start))
	_ = data
}

func TestGetBQ(t *testing.T) {
	resp, err := http.Get("http://localhost:8000/GetGroupedBar/2")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(resp)
}
