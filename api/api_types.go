package api

import (
	"go-api/orm"
)

type ArgClauseMap map[string][]string

type ListStatsResponse struct {
	Stats []orm.Player
}

type ListStatsResponse2 struct {
	Stats []Stats
}

type Stats struct {
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

type ListStatsFilter struct {
	Filters_IS         ArgClauseMap
	Filters_NOT        ArgClauseMap
	MinimumDatasetSize int
	Query              string
	Target             string `json:"target"`
}
type GlobalFilters struct {
	Agents   []string `json:"agents"`
	Mapnames []string `json:"mapnames"`
	Players  []string `json:"players"`
	Teams    []string `json:"teams"`
	Side     string   `json:"side"`
}
type ListStatsRequest struct {
	Filters_IS  GlobalFilters `json:"filters"`
	Filters_NOT GlobalFilters `json:"filters_NOT"`
}
