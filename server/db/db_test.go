package db

import (
	"fmt"
	"testing"
	"time"

	log "github.com/sirupsen/logrus"
)

func TestMakePatches(t *testing.T) {
	MigrateNewDb()

	Dates := []string{"2021-12-07", "2022-01-11", "2022-01-19", "2022-01-02", "2022-02-15", "2022-03-01", "2022-03-22", "2022-04-12", "2022-04-27", "2022-05-10", "2022-05-24", "2022-06-07", "2022-06-22", "2022-07-12", "2022-08-09", "2022-08-23", "2022-09-07", "2022-09-20", "2022-10-04", "2022-10-18", "2022-11-01", "2022-11-15", "2022-12-06", "2023-01-10", "2023-01-18", "2023-02-07", "2023-02-14", "2023-03-07", "2023-03-14", "2023-03-28"}
	Versions := []float64{3.12, 4, 4.01, 4.02, 4.03, 4.04, 4.05, 4.07, 4.08, 4.09, 4.10, 4.11, 5.0, 5.01, 5.03, 5.04, 5.05, 5.06, 5.07, 5.08, 5.09, 5.10, 5.12, 6.0, 6.01, 6.02, 6.03, 6.04, 6.05, 6.06}
	if len(Dates) != len(Versions) {
		t.Fatal(fmt.Errorf("oh shit"))
	}
	db := Getdb().DB

	for i, v := range Dates {
		var p GameVersion
		d, err := time.Parse("2006-01-02", v)
		if err != nil {
			t.Fatal(err)
		}
		p.Patch = Versions[i]
		p.ReleaseDate = d
		if err := db.Create(&p).Error; err != nil {
			log.Info(err)
		}
	}
}

/* func TestGetVCTMatches(t *testing.T) {
	c := GetCollector()
	urls := GetVCTmatches(c)
	for i := range urls {
		log.Println(urls[i])
	}
}
func TestGetDatasetChan(t *testing.T) {
	urlC := GetCollector()
	dataC := GetCollector()
	urls := GetVCTmatches(urlC)
	dchan := GetCleanDataChan(urls, dataC)
	for ds := range dchan {
		log.Println(ds.Shared.Teams)
	}
}
func TestButActuallyInsertToDb(t *testing.T) {
	urlC := GetCollector()
	dataC := GetCollector()
	urls := GetVCTmatches(urlC)
	furls := filterUrls(urls)
	dchan := GetCleanDataChan(furls, dataC)
	db := GetNewDB()
	for ds := range dchan {
		if err := CreateDbEntries(ds, db); err != nil {
			log.Info(err)
		}
	}
}

func TestScrapeOne(t *testing.T) {
	c := GetCollector()
	d, err := Scrape("https://www.vlr.gg/183774", c)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v", d)
}

func TestSniff(t *testing.T) {
	//url := "https://www.vlr.gg/183708"
	url := "https://www.vlr.gg/183708"
	c := colly.NewCollector(
		colly.AllowedDomains("vlr.gg"),
	)
	isValid := SniffMatch(url, c)
	log.Info(isValid)
	log.Info(url == isValid)
}
*/
