package structs

import (
	"fmt"
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
	/*p := perlin.NewPerlin(2, 2, 2, int64(100))
	rand := float64(rand.Intn(5)) / 10
	valRand := math.Abs(p.Noise1D(rand) * 1000)*/
	valRand := 0.0
	return SensorData{s.Id, s.AirportId, s.Measure, valRand, currentTime}
}

func (s SensorData) String() string {
	return fmt.Sprintf("{Id=%d; AirportId=%s; Measure=%s; Value=%f; TimeStamp=%s}", s.Id, s.AirportId, s.Measure, s.Value, s.Timestamp.String())
}
