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
	Filters_IS         ArgClauseMap
	Filters_NOT        ArgClauseMap
	MinimumDatasetSize int
	Query              string
	Target             string `json:"target"`
}
