package graphs

import (
	"encoding/json"
	"fmt"
	"refactor/db"
	"refactor/stats"
	"time"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

type GraphAPI struct {
}

func GetGraphAPI() *GraphAPI {
	return &GraphAPI{}
}

type AvgResultStr bool

func (a *AvgResultStr) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	if s == "true" {
		*a = true
	} else {
		*a = false
	}
	return nil
}

type TsrResponse struct {
	Datasets []TsrResDataset `json:"datasets"`
}

type TsrResDataset struct {
	Label     string         `json:"label"`
	Data      []TsrDataPoint `json:"data"`
	BgCol     string         `json:"backgroundColor"`
	BorderCol string         `json:"borderColor"`
}
type TsrDataPoint struct {
	X string  `json:"x"`
	Y float64 `json:"y"`
	R float64 `json:"r"`
}

func (t *TimeRange) ParseInputs() error {
	if (len(t.Patches) == 0) && (len(t.ExtDates) == 0) {
		return nil
	}
	var start, end db.GameVersion
	db := db.Getdb().DB
	if t.Patches != nil {
		if len(t.Patches) != 2 {
			return fmt.Errorf("err: can only range between 2 values")
		}
		if err := db.Find(&start, "patch = ?", t.Patches[0]).Error; err != nil {
			return fmt.Errorf("err: failed to find patch matching input")
		}
		if err := db.Find(&end, "patch = ?", t.Patches[1]).Error; err != nil {
			return fmt.Errorf("err: failed to find patch matching input")
		}
		t.IntDates = []time.Time{start.ReleaseDate, end.ReleaseDate}
		log.Info(fmt.Sprintf("Parsed %v, %v", start.ReleaseDate, end.ReleaseDate))
		return nil
	} else {
		if len(t.ExtDates) != 2 {
			return fmt.Errorf("err: can only range between 2 values")
		}
		sParsed, err := time.Parse("2006-01-02", t.ExtDates[0])
		if err != nil {
			return err
		}
		eParsed, err := time.Parse("2006-01-02", t.ExtDates[1])
		if err != nil {
			return err
		}
		t.IntDates = []time.Time{sParsed, eParsed}
		log.Info(fmt.Sprintf("Parsed %v, %v", sParsed, eParsed))
		return nil
	}
}

type TimeRange struct {
	Patches  []float64 `json:"Patches,omitempty"`
	ExtDates []string  `json:"Dates,omitempty"`
	IntDates []time.Time
}

type TsrRequest struct {
	INfilters      *stats.CompsFilter `json:"IS_filters"`
	NOTINfilters   *stats.CompsFilter `json:"NOT_filters"`
	SearchRange    TimeRange          `json:"Range"`
	Y_Target       string             `json:"y_target"`
	Radius_Scale   string             `json:"r_target"`
	MinDatasetSize int                `json:"min_dataset_size,string"`
	DateSplits     int                `json:"date_splits,string"`
	MaxResults     int                `json:"max_results,string"`
	RequestID      uuid.UUID
}

type DgbRequest struct {
	DoesContainFilters *stats.StatsFilter `json:"IS_filters"`
	NotContainFilters  *stats.StatsFilter `json:"NOT_filters"`
	Side               string             `json:"side"`
	X_Target           string             `json:"x_target"`
	X_Grouping         string             `json:"x_groups_target"`
	Y_Target           string             `json:"y_target"`
	MinDatasetSize     int                `json:"min_dataset_size,string"`
	MaxDatasetsCount   int                `json:"max_dataset_amount,string"`
	AverageResult      AvgResultStr       `json:"average_results"`
	RequestID          uuid.UUID
}

type DgbResponse struct {
	Labels    []string        `json:"labels"`
	Datasets  []DgbResDataset `json:"datasets"`
	RequestID uuid.UUID
}

type DgbQueryGroups struct {
	X_Targets []string
	GroupVals []string
}

// Contains the data result from SQL statement
type DgbReqDataset struct {
	X_Position string
	Data       map[string]float64
	Ind        int
}

type DgbResDataset struct {
	Id              int           `json:"id"`
	Label           string        `json:"label"`
	Data            []interface{} `json:"data"`
	BackgroundColor string        `json:"backgroundColor"`
}
