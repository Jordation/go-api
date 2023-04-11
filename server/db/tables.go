package db

import (
	"path/filepath"
	"runtime"
	"time"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetNewDB() *gorm.DB {
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	filePath := filepath.Join(dir, "new_DEV.db")
	db, err := gorm.Open(sqlite.Open(filePath), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func MigrateNewDb() {
	db := GetNewDB()
	db.AutoMigrate(&HitLink{})
	db.AutoMigrate(&PlayerStat{})
	db.AutoMigrate(&AgentComp{})
	db.AutoMigrate(&Map{})
	db.AutoMigrate(&Event{})
	db.AutoMigrate(&GameVersion{})
}

type GameVersion struct {
	Patch           float64 `gorm:"unique"`
	ReleaseDate     time.Time
	RelevantChanges string
}

type HitLink struct {
	ID   uint
	Link string `gorm:"unique"`
}

type Event struct {
	ID        uint
	EventName string `gorm:"unique"`
	Maps      []Map
}

type Map struct {
	ID         uint
	EventID    uint
	MatchUUID  uuid.UUID
	MatchDate  time.Time
	Team1      string
	Team2      string
	Winner     string
	AtkRndsWon uint
	DefRndsWon uint
	Map        string
	Players    []PlayerStat
	Comps      []AgentComp
}

type AgentComp struct {
	ID       uint
	MapID    uint
	Map      string
	Comp     string
	PickedBy string
	Won      bool
}

type PlayerStat struct {
	ID      uint
	MapID   uint
	Player  string
	Team    string
	Map     string
	Agent   string
	Rating  float64
	ACS     uint
	Kills   uint
	Deaths  uint
	Assists uint
	KAST    uint
	ADR     uint
	HSP     uint
	FK      uint
	FD      uint
	Side    string
}
