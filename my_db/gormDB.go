package my_db

import (
	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Event struct {
	ID uint
	gorm.Model
	EventName string
	Maps      []Map
}

type Map struct {
	ID         uint
	EventID    uint
	MatchUUID  uuid.UUID
	Team1      string
	Team2      string
	Team1Comp  string
	Team2Comp  string
	Winner     string
	AtkRndsWon uint
	DefRndsWon uint
	MapName    string
	MatchName  string
	Players    []Player
}
type Player struct {
	ID         uint
	MapID      uint
	PlayerName string
	Team       string
	MapName    string
	Agent      string
	Rating     uint
	ACS        uint
	Kills      uint
	Deaths     uint
	Assists    uint
	KAST       uint
	ADR        uint
	HSP        uint
	FK         uint
	FD         uint
	Side       string
}

func MigrateDB() {
	db, err := gorm.Open(sqlite.Open("dev.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Event{}, &Map{}, &Player{})
}
