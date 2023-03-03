package api

import (
	"go-api/orm"
)

type ArgClauseMap map[string][]string

type ListStatsResponse struct {
	Stats []orm.Player
}

type ListStatsFilter struct {
	//			  arg    : clause
	Filters            ArgClauseMap
	MinimumDatasetSize int
	Query              string
	Target             string
}
