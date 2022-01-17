package structs

import (
	"fmt"
	"math"
	"time"
)

type Sensor struct {
	Id        int
	AirportId string
	Measure   Measure
}

// Data sent back by a sensor
type SensorData struct {
	Id        int
	AirportId string
	Measure   Measure
	Value     float64
	Timestamp time.Time
}

func (s Sensor) GenerateData(currentTime time.Time) SensorData {
	val := s.GenData(currentTime)
	return SensorData{s.Id, s.AirportId, s.Measure, val, currentTime}
}

func (s SensorData) String() string {
	return fmt.Sprintf("{Id=%d; AirportId=%s; Measure=%s; Value=%f; TimeStamp=%s}", s.Id, s.AirportId, s.Measure, s.Value, s.Timestamp.String())
}

func (s Sensor) GenData(currentTime time.Time) float64 {
	firstTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, currentTime.Location())
	x := currentTime.Sub(firstTime)
	timeForGenData := x.Seconds()
	switch s.Measure {
	case Temperature:
		return GenTemp(timeForGenData)
	case Pressure:
		return GenPres(timeForGenData)
	case Wind:
		return GenWind(timeForGenData)
	default: // just in case
		return 0
	}
}

func GenTemp(x float64) float64 {
	x = x / 3600
	return 3.5*math.Cos(math.Pi/12*x+2.7) + 0.02*math.Cos(4*math.Pi*x) + 0.5*math.Cos(0.5*math.Pi*x) + 8
}

func GenPres(x float64) float64 {
	x = x / 3600
	return 2.5*math.Cos(math.Pi/12*x+2.7) + 0.02*math.Cos(4*math.Pi*x) + 0.1*math.Cos(0.5*math.Pi*x) + 1013
}

func GenWind(x float64) float64 {
	x = x / 3600
	y := 0.00

	for i := 0; i < 60; i++ {
		y += 4 / math.Pi * (math.Sin(2*float64(i)+1) * 2 * math.Pi * 1 / 12 * x)
	}
	return y

}
