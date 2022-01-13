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
	/*
		paramValues := map[Measure][3]float64{
			Temperature: {2, 2, 2},
			Pressure:    {2, 2, 2},
			Wind:        {2, 2, 2},
		}
		param := paramValues[s.Measure]
		p := perlin.NewPerlinRandSource(param[0], param[1], int32(param[2]), rand.NewSource(100))
		valRand := math.Abs(p.Noise1D(float64(currentTime.Unix())))
		valRand = valRand / 1000000000000000000000000000000000000 // TODO: find a proper & cleaner way to get a 2-digit integer
		//valRand := GenerateStableRandomNumericalValues(s.Measure)

	*/
	val := GenData(currentTime)

	return SensorData{s.Id, s.AirportId, s.Measure, val, currentTime}
}

func (s SensorData) String() string {
	return fmt.Sprintf("{Id=%d; AirportId=%s; Measure=%s; Value=%f; TimeStamp=%s}", s.Id, s.AirportId, s.Measure, s.Value, s.Timestamp.String())
}

func GenData(currentTime time.Time) float64 {
	firstTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, currentTime.Location())
	x := currentTime.Sub(firstTime)
	return CourbeTemperature(x.Seconds())
}

func CourbeTemperature(x float64) float64 {
	x = x / 60
	return 3.5*math.Cos(math.Pi/12*x+2.7) + 0.02*math.Cos(4*math.Pi*x) + 0.5*math.Cos(0.5*math.Pi*x) + 8
}

func CourbePression(x float64) float64 {
	x = x / 60
	return 0.001*math.Cos(math.Pi/12*x+2.7) + 0.02*math.Cos(4*math.Pi*x) + 0.1*math.Cos(0.5*math.Pi*x) + 1.013
}

func CourbeVent(x float64) float64 {
	x = x / 60
	y := 0.00

	for i := 0; i < 60; i++ {
		y += 4 / math.Pi * (math.Sin(2*float64(i)+1) * 2 * math.Pi * 1 / 12 * x)
	}
	return y

}
