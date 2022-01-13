package configuration

import "fmt"

type SensorConfig struct {
	MqttConf    MqttConf
	MeasureType string `json:"measureType"`
	AirportId   string `json:"airportId"`
	ClientId    int    `json:"clientId"`
}

func (s SensorConfig) String() string {
	return fmt.Sprintf("{MqttConf=%s;MeasureType=%s; AirportId=%s;ClientId=%d}",
		s.MqttConf.String(), s.MeasureType, s.AirportId, s.ClientId)
}
