package my_db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Event struct {
	ID        uint64
	EventName string `gorm:"unique"`
	Maps      []Map
}

type Map struct {
	ID         uint64
	EventID    uint64
	MatchUUID  string
	MatchName  string
	Team1      string
	Team2      string
	Team1Comp  string
	Team2Comp  string
	Winner     string
	AtkRndsWon uint64
	DefRndsWon uint64
	MapName    string
	Players    []Player
}
type Player struct {
	ID         uint64
	MapID      uint64
	PlayerName string
	Team       string
	MapName    string
	Agent      string
	Rating     float64
	ACS        uint64
	Kills      uint64
	Deaths     uint64
	Assists    uint64
	KAST       uint64
	ADR        uint64
	HSP        uint64
	FK         uint64
	FD         uint64
	Side       string
}

func MigrateDB() {
	db, err := gorm.Open(sqlite.Open("dev.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Event{}, &Map{}, &Player{})
}
