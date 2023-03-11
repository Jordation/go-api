package server

import "go-api/api"

func MapQueryFilters(f api.GlobalFilters) *map[string][]string {
	side := []string{f.Side}
	filters := map[string][]string{
		"agent":       f.Agents,
		"map_name":    f.Mapnames,
		"team":        f.Teams,
		"player_name": f.Players,
		"side":        side,
	}
	for k, v := range filters {
		if v == nil {
			filters[k] = nil
		}
	}
	return &filters
}
