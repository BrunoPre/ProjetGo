package pkg

import (
	"time"
)

type SensorData struct {
	Id        int       `json:"id"`
	AirportId string    `json:"airportId"`
	Measure   Measure   `json:"measure"`
	Value     float64   `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}

type Measure string

// Defining existing measure types
const (
	Temperature Measure = "Temperature"
	Pressure            = "Atmospheric pressure"
	Wind                = "Wind speed"
)

type SensorDatas []SensorData

// HandleError conveniently handles error.
func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}
