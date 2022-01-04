package sensors

import (
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"time"
)

type Sensor struct {
	idSensor    int
	idAirport   int    // IATA format
	typeMeasure string // temp, wind, pressure
	valMeasure  float32
	timestamp   string // format: YYYY-MM-DD-hh-mm-ss
}

func (s Sensor) String() string {
	res, _ := json.Marshal(s)
	return fmt.Sprintf(string(res))
}

func (s Sensor) publishOnce(client mqtt.Client) {

	msg := s.String()
	token := client.Publish("sensors/"+s.typeMeasure, 0, false, msg)
	token.Wait()
}

func NewSensor(idSensor, idAirport int, typeMeasure string) *Sensor {
	currentTime := time.Now()
	timestamp := string(currentTime.Year() + '-' + int(currentTime.Month()) + '-' + currentTime.Day() + '-' + currentTime.Hour() + '-' + currentTime.Minute() + '-' + currentTime.Second())
	return &Sensor{idSensor: idSensor, idAirport: idAirport, typeMeasure: typeMeasure, valMeasure: 0, timestamp: timestamp}
}
