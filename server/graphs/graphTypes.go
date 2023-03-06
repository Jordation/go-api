package graphs

import "strings"

// helper bc i dont want to work on my frontend any more
func MapQueryFilters(f *GlobalFilters) map[string][]string {
	filters := make(map[string][]string)
	filters["agent"] = strings.Split(f.Agents, ", ")
	filters["map_name"] = strings.Split(f.Mapnames, ", ")
	filters["team"] = strings.Split(f.Teams, ", ")
	filters["player_name"] = strings.Split(f.Players, ", ")
	filters["side"] = []string{f.Side}
	for k, v := range filters {
		if v[0] == "" {
			filters[k] = nil
		}
	}
	return filters

}

type GroupedBarResponse struct {
	Groups []string
	Data   map[string][]float64
}
type GlobalFilters struct {
	Agents   string `json:"agents"`
	Mapnames string `json:"mapnames"`
	Players  string `json:"players"`
	Teams    string `json:"teams"`
	Side     string `json:"side"`
}
type GroupedBarRequest struct {
	Filters_IS       *GlobalFilters `json:"filters"`
	Filters_NOT      *GlobalFilters `json:"filters_NOT"`
	XTarget          string         `json:"x_target"`
	XGroupsTarget    string         `json:"x_groups_target"`
	YTarget          string         `json:"y_target"`
	AverageResults   bool           `json:"average_results"`
	MinDatasetSize   int            `json:"min_dataset_size,string"`
	MaxDatasetAmount int            `json:"max_dataset_amount,string"`
}
