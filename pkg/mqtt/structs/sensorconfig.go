package structs

type SensorConfig struct {
	// uppercase on first letter --> export the attrs for json unmarshalling
	ClientId    int    `json:"clientId"`
	BrokerAddr  string `json:"brokerAddr"`
	BrokerPort  int    `json:"brokerPort"`
	QosLevel    int    `json:"qosLevel"`
	MeasureType string `json:"measureType"`
	AirportId   int    `json:"airportId"`
}

type SensorsConfigs struct {
	SensorsConfigs []SensorConfig `json:"sensors-configs"'`
}
