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
	args := os.Args
	// argument parsing
	if len(args) != 2 {
		panic("Incorrect arguments lengths. Please provide the path to the config file")
	}
	fmt.Println(args)

	arg := os.Args[1]

	fileByte, err := ioutil.ReadFile(filepath.Clean(arg))
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
	brokerUri := s.BrokerAddr + ":" + strconv.Itoa(s.BrokerPort) // "addr:port"
	//clientId := strconv.Itoa(s.ClientId)
	qosLevel := byte(s.QosLevel)
	//airportId := s.AirportId

	// Init REDIS DB & controller
	storage.Init()
	controller.Init()

	client := mqttClient.Connect(brokerUri, "sub")
	// subscribe to all airports (allowed by '#' wildcard)
	token := client.Subscribe("airport/#", qosLevel, controller.Controller.HandleSensorData)

	token.Wait()
	for {
	}

}
