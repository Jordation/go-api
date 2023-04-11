package stats

import (
	"encoding/json"
	"os"
	"testing"

	log "github.com/sirupsen/logrus"
)

func TestQuery(t *testing.T) {
	q, err := ReadQuery()
	if err != nil {
		t.Fatalf(err.Error())
	}
	MappedFilters := MakeMappedStatsFilter(q)
	log.Println("Mapped Filters: ", MappedFilters)
	log.Println("Query: ", q)
}

func TestMakeStmt(t *testing.T) {
	q, err := ReadBarQuery()
	if err != nil {
		t.Fatalf(err.Error())
	}
	base := StatsListQuery
	stmt, args := GetStatsQuery(q, base, nil)
	log.Println("Stmt is: ", stmt)
	log.Println("Args are: ", args)
}

func ReadBarQuery() (*ListStatsRequest, error) {
	var query ListStatsRequest
	data, err := os.ReadFile("../bar_query.json")
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &query)
	return &query, nil
}
func ReadQuery() (*ListStatsRequest, error) {
	var query ListStatsRequest
	data, err := os.ReadFile("../query.json")
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &query)
	return &query, nil
}
