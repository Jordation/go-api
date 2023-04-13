package helpers

import (
	"math"
)

func RoundFloat(f float64) float64 {
	return math.Round(f*100) / 100
}

func GetPercent(part, whole float64) float64 {
	return RoundFloat(part / whole * 100)
}

func GetQueryDir() string {
	return `C:\DEV\go-api\server\graphs\PreMadeQuazzas\`
}
