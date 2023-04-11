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

func (s *PlayerStat) LoadValues(Player, Agent, Team, Map, Side string, IntStats []uint, Rating float64) {
	s.Player = Player
	s.Agent = Agent
	s.Team = Team
	s.Map = Map
	s.Side = Side
	s.Rating = Rating
	s.ACS = IntStats[0]
	s.Kills = IntStats[1]
	s.Deaths = IntStats[2]
	s.Assists = IntStats[3]
	s.KAST = IntStats[4]
	s.ADR = IntStats[5]
	s.HSP = IntStats[6]
	s.FK = IntStats[7]
	s.FD = IntStats[8]
}

func (m *Map) LinkChildren(Players []PlayerStat, Comps []AgentComp) {
	m.Players = Players
	m.Comps = Comps
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
