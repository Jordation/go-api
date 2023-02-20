package GetMyData

import (
	"errors"
	gormdb "go-api/initial/my_db"
	"log"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

func getStatUint(s string) uint64 {
	i, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		log.Println(s, err)
		return 0
	}
	return i
}
func getStatFloat(s string) float64 {
	i, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Println(s, err)
		return 0
	}
	return i
}

func makeCombinedStruct(atkP gormdb.Player, defP gormdb.Player) gormdb.Player {
	var comStat gormdb.Player
	comStat.Rating = math.Round(((atkP.Rating+defP.Rating)/2)*100) / 100
	comStat.ACS = (atkP.ACS + defP.ACS) / 2
	comStat.Kills = (atkP.Kills + defP.Kills) / 2
	comStat.Deaths = (atkP.Deaths + defP.Deaths) / 2
	comStat.Assists = (atkP.Assists + defP.Assists) / 2
	comStat.KAST = (atkP.KAST + defP.KAST) / 2
	comStat.ADR = (atkP.ADR + defP.ADR) / 2
	comStat.HSP = (atkP.HSP + defP.HSP) / 2
	comStat.FK = (atkP.FK + defP.FK) / 2
	comStat.FD = (atkP.FD + defP.FD) / 2
	comStat.Side = "C"
	comStat.PlayerName = atkP.PlayerName
	comStat.MapName = atkP.MapName
	comStat.Agent = atkP.Agent
	comStat.Team = atkP.Team
	return comStat
}
func makePlayerStruct(p gormdb.Player, md mappedPlayerData, side string) gormdb.Player {
	p.Rating = getStatFloat(md["Rating"])
	p.ACS = getStatUint(md["ACS"])
	p.Kills = getStatUint(md["Kills"])
	p.Deaths = getStatUint(md["Deaths"])
	p.Assists = getStatUint(md["Assists"])
	p.KAST = getStatUint(md["KAST"])
	p.ADR = getStatUint(md["ADR"])
	p.HSP = getStatUint(md["HSP"])
	p.FK = getStatUint(md["FK"])
	p.FD = getStatUint(md["FD"])
	p.Side = side
	return p
}

// takes single row of data from scraper, returns struct for atk, def and combined statline
func formatPlayerData(d playerData, team string, mapname string) []gormdb.Player {
	var pStats gormdb.Player
	pStats.Agent = d.agent
	pStats.PlayerName = d.player
	pStats.Team = team
	pStats.MapName = mapname

	atkStat := makePlayerStruct(pStats, d.statsT, "A")
	defStat := makePlayerStruct(pStats, d.statsCT, "D")
	comStat := makeCombinedStruct(atkStat, defStat)
	return []gormdb.Player{atkStat, defStat, comStat}
}

// takes 5 rows of data from scraper along with the matching team
func handleTeam(players []playerData, team string, mapname string) []gormdb.Player {
	var rows []gormdb.Player
	for _, v := range players {
		rows = append(rows, formatPlayerData(v, team, mapname)...)
	}
	return rows
}

func findWinner(data []string) string {
	t1s := getStatUint(data[2]) + getStatUint(data[4])
	t2s := getStatUint(data[3]) + getStatUint(data[5])
	if t1s > t2s {
		return data[0]
	} else {
		return data[1]
	}
}

func findComp(rows []playerData) string {
	var comp []string
	for _, v := range rows {
		comp = append(comp, v.agent)
	}

	sort.Strings(comp)
	return strings.Join(comp, ",")
}

func handleMap(gd gameData) gormdb.Map {
	var mapData gormdb.Map
	mapData.Team1 = gd.data[0]
	mapData.Team2 = gd.data[1]

	mapData.Team1Comp = findComp(gd.players[:5])
	mapData.Team2Comp = findComp(gd.players[5:])

	mapData.DefRndsWon = getStatUint(gd.data[2]) + getStatUint(gd.data[3])
	mapData.AtkRndsWon = getStatUint(gd.data[4]) + getStatUint(gd.data[5])

	mapData.Winner = findWinner(gd.data)
	mapData.MapName = gd.mapname

	return mapData
}

func formatMapData(gd gameData, mID string, matchName string) gormdb.Map {
	var playerMapData []gormdb.Player
	mapData := handleMap(gd)

	playerMapData = append(playerMapData, handleTeam(gd.players[:5], mapData.Team1, mapData.MapName)...)
	playerMapData = append(playerMapData, handleTeam(gd.players[5:], mapData.Team2, mapData.MapName)...)
	mapData.Players = playerMapData

	mapData.MatchName = matchName
	mapData.MatchUUID = mID

	return mapData
}

func handleMaps(gd []gameData, mID string, matchName string) []gormdb.Map {
	var maps []gormdb.Map

	for _, v := range gd {
		maps = append(maps, formatMapData(v, mID, matchName))
	}

	return maps
}

func ValidateScrapedData(d matchData) error {
	if len(d.matchInfo) != 6 {
		return errors.New("malformed data")
	}

	for _, v := range d.gameData {
		if len(v.data) != 6 {
			return errors.New("malformed data")
		}

		if len(v.players) != 10 {
			return errors.New("malformed data")
		}
		for _, v2 := range v.players {
			if len(v2.statsCT)&len(v2.statsT) != 10 {
				return errors.New("malformed data")
			}
		}
	}

	return nil
}

func MakeORMstruct(d matchData) gormdb.Event {
	var e gormdb.Event
	e.EventName = d.matchInfo[0]
	matchName := d.matchInfo[1]
	matchID := uuid.New().String()
	e.Maps = handleMaps(d.gameData, matchID, matchName)
	return e
}
