package sensors

import (
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
	"time"
)

type Sensor struct {
	idSensor    int     `json:"idSensor"`
	idAirport   int     `json:"idAirport"`   // IATA format
	typeMeasure string  `json:"typeMeasure"` // temp, wind, pressure
	valMeasure  float32 `json:"valMeasure"`
	timestamp   string  `json:"timestamp"` // format: YYYY-MM-DD-hh-mm-ss
}

func (s Sensor) String() string { //TODO: correct the stringer
	res, err := json.MarshalIndent(s, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	return string(res)
}

func (s Sensor) publishOnce(client mqtt.Client) {
	msg := s.String()
	fmt.Println(msg)
	token := client.Publish("sensors/"+s.typeMeasure, 0, false, msg)
	token.Wait()
}

func NewSensor(idSensor, idAirport int, typeMeasure string) *Sensor {
	currentTime := time.Now()
	timestamp := string(currentTime.Year() + '-' + int(currentTime.Month()) + '-' + currentTime.Day() + '-' + currentTime.Hour() + '-' + currentTime.Minute() + '-' + currentTime.Second())
	return &Sensor{idSensor: idSensor, idAirport: idAirport, typeMeasure: typeMeasure, valMeasure: 0, timestamp: timestamp}
}
