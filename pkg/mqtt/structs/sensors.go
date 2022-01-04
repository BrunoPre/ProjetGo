package structs

import (
	"fmt"
	"time"

	"github.com/aquilax/go-perlin"
)

type Sensor struct {
	Id        int
	AirportId int
	Measure   Measure
}

// Data sent back by a sensor
type SensorData struct {
	Id        int
	AirportId int
	Measure   Measure
	Value     float64
	Timestamp time.Time
}

func (s Sensor) GenerateData(currentTime time.Time) SensorData {
	return SensorData{s.Id, s.AirportId, s.Measure, perlin.NewPerlin(2, 2, 2, 451142).Noise1D(float64(currentTime.Unix())), currentTime}
}

func (s SensorData) String() string {
	return fmt.Sprintf("{Id=%d; AirportId=%d; Measure=%s; Value=%f, TimeStamp=%s}", s.Id, s.AirportId, s.Measure, s.Value, s.Timestamp.String())
}
