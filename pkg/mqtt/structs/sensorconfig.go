package structs

import "fmt"

type SensorConfig struct {
	// uppercase on first letter --> export the attrs for json unmarshalling
	ClientId    int    `json:"clientId"`
	BrokerAddr  string `json:"brokerAddr"`
	BrokerPort  int    `json:"brokerPort"`
	QosLevel    int    `json:"qosLevel"`
	MeasureType string `json:"measureType"`
	AirportId   int    `json:"airportId"`
}

func (s SensorConfig) String() string {
	return fmt.Sprintf("{ClientId=%d; BrokerAddr=%s; BrokerPort=%d; QosLevel=%d; MeasureType=%s; AirportId=%d}",
		s.ClientId, s.BrokerAddr, s.BrokerPort, s.QosLevel, s.MeasureType, s.AirportId)
}
