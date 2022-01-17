package configuration

import "fmt"

type Config struct {
	Dao      DaoConfig
	MqttConf MqttConf
}
type DaoConfig struct {
	DaoType   string
	FilePath  string
	RedisAddr string
	RedisPwd  string
	RedisDb   int
}

type MqttConf struct {
	BrokerAddr string `json:"brokerAddr"`
	BrokerPort int    `json:"brokerPort"`
	QosLevel   int    `json:"qosLevel"`
	ClientName string `json:"clientName"`
}

func (m MqttConf) String() string {
	return fmt.Sprintf("{BrokerAddr=%s;BrokerPort=%d;QosLevel=%d}", m.BrokerAddr, m.BrokerPort, m.QosLevel)
}
