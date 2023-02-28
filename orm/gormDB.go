package orm

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type HitLink struct {
	ID   uint64
	Link string `gorm:"unique"`
}

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

func GetDB() (*gorm.DB, error) {
	return gorm.Open(sqlite.Open("dev.db"), &gorm.Config{})
}

func MigrateDB() {
	db, err := GetDB()
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Event{}, &Map{}, &Player{}, &HitLink{})
}

func CheckPlayerRow(r Player) bool {
	if r.MapName == "" {
		return false
	}
	if r.PlayerName == "" {
		return false
	}
	if r.Kills == 0 && r.Deaths == 0 && r.Assists == 0 && r.ACS == 0 {
		return false
	}
	return true
}

func CleanDB() {
	db, err := GetDB()
	if err != nil {
		panic("failed to connect database")
	}

	var players []Player
	result := db.Find(&players)
	if result.Error != nil {
		fmt.Println(result.Error.Error())
	}

	for _, p := range players {
		if !CheckPlayerRow(p) {
			fmt.Println("deleting", p)
			db.Delete(&p)
		}
	}
}
