package pkg

import (
	"time"
)

type SensorData struct {
	Id        int       `json:"id"`
	AirportId string    `json:"airportId"`
	Measure   string    `json:"measure"`
	Value     float64   `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}

type Measure string

type SensorDatas []SensorData

type SensorDataAverage struct {
	AverageWind     float64 `json:"averageWind"`
	AveragePressure float64 `json:"averagePressure"`
	AverageTemp     float64 `json:"averageTemp"`
}

// HandleError conveniently handles error.
func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}
