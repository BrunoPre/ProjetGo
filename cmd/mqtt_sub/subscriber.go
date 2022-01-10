package main

import (
	mqttClient "Project/pkg/mqtt/client"
	"Project/pkg/mqtt/controller"
	"Project/pkg/mqtt/storage"
	"Project/pkg/mqtt/structs"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
)

func main() {

	// argument parsing
	if len(os.Args) != 2 {
		panic("Incorrect arguments lengths. Please provide the path to the config file")
	}
	fmt.Println(os.Args)

	fileByte, err := ioutil.ReadFile(filepath.Clean(os.Args[1]))
	if err != nil {
		panic(fmt.Sprintf("Couldn't open config file \"%w\"", err))
	}

	// unmarshalling the JSON config file
	var sensorConfig structs.SensorConfig
	err = json.Unmarshal(fileByte, &sensorConfig)
	if err != nil {
		panic(fmt.Sprintf("Couldn't parse the file \"%w\"", err))
	}

	// instance of sensor config
	s := sensorConfig
	fmt.Println(s.String())
	brokerUri := s.BrokerAddr + ":" + strconv.Itoa(s.BrokerPort)
	clientId := strconv.Itoa(s.ClientId)
	qosLevel := byte(s.QosLevel)
	airportId := s.AirportId

	storage.Init()
	controller.Init()

	client := mqttClient.Connect(brokerUri, clientId+"-sub")
	token := client.Subscribe(airportId, qosLevel, controller.Controller.HandleSensorData)

	token.Wait()
	for {
	}

}
