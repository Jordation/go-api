package stats

// i.e. StatsFilterMap["IS"]["agents"][args]
// i.e. StatsFilterMap["NOT"]["agents"][args]
type StatsFilterMap map[string]map[string][]string
type HavingClause struct {
	Targets []string
	Count   int
}

type StatsFilter struct {
	Agents  []string `json:"agents,omitempty"`
	Maps    []string `json:"maps,omitempty"`
	Players []string `json:"players,omitempty"`
	Teams   []string `json:"teams,omitempty"`
}
type CompsFilter struct {
	Teams  []string `json:"teams,omitempty"`
	Maps   []string `json:"maps,omitempty"`
	Agents []string `json:"agents,omitempty"`
}
type MapsFilter struct {
	Teams []string `json:"teams,omitempty"`
	Maps  []string `json:"maps,omitempty"`
}

// JSON tags for API, manually filled in graph service
type ListStatsRequest struct {
	DoesContainFilters *StatsFilter `json:"IS_filters"`
	NotContainFilters  *StatsFilter `json:"NOT_filters"`
	Side               string       `json:"side"`
	TargetRow          string       `json:"y_target"`
	ExtraRows          []string
}

type ListMapsRequest struct {
	// start and end MUST be formatted datetimes to YY:MM:DD HH:MM:SS
	Start    string
	End      string
	EventIDs []uint

	// is: team1: [vals]
	// not: team2: [other vals]
	INfilters    *MapsFilter
	NOTINfilters *MapsFilter
}

type ListCompsRequest struct {
	MapIds         []uint
	MinDatasetSize int `json:"min_dataset_size,string"`
	Hc             *HavingClause
	INfilters      *CompsFilter `json:"IS_filters"`
	NOTINfilters   *CompsFilter `json:"NOT_filters"`
}
