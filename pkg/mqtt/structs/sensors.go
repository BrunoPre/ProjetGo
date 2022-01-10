package structs

import (
	"fmt"
	"github.com/aquilax/go-perlin"
	"math"
	"math/rand"
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
	p := perlin.NewPerlinRandSource(2, 2, 2, rand.NewSource(100))
	valRand := math.Abs(p.Noise1D(float64(currentTime.Unix())))
	valRand = valRand / 1000000000000000000000000000000000000 // TODO: find a proper & cleaner way to get a 2-digit integer
	//valRand := GenerateStableRandomNumericalValues(s.Measure)
	return SensorData{s.Id, s.AirportId, s.Measure, float64(valRand), currentTime}
}

func (s SensorData) String() string {
	return fmt.Sprintf("{Id=%d; AirportId=%s; Measure=%s; Value=%f; TimeStamp=%s}", s.Id, s.AirportId, s.Measure, s.Value, s.Timestamp.String())
}
