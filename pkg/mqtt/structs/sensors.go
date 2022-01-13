package structs

import (
	"fmt"
	"time"
)

// Data sent back by a sensor
type SensorData struct {
	Id        int
	AirportId int
	Measure   Measure
	Value     float64
	Timestamp time.Time
}

func (s SensorData) String() string {
	return fmt.Sprintf("{Id=%d; AirportId=%d; Measure=%s; Value=%f, TimeStamp=%s}", s.Id, s.AirportId, s.Measure, s.Value, s.Timestamp.String())
}
