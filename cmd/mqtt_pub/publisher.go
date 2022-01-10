package main

import (
	mqttClient "Project/pkg/mqtt/client"
	"Project/pkg/mqtt/structs"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"time"
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
	clientId := s.ClientId
	qosLevel := byte(s.QosLevel)
	airportId := s.AirportId
	fmt.Printf("airportId: %s\n", airportId)
	measure := s.MeasureType

	sensor := structs.Sensor{Id: clientId, AirportId: airportId, Measure: structs.Measure(measure)}

	// connecting to MQTT client
	client := mqttClient.Connect(brokerUri,
		// "golang-sub")
		strconv.Itoa(clientId)+"-pub")

	var sensorData structs.SensorData

	for {
		sensorData = sensor.GenerateData(time.Now())

		json, err := json.Marshal(sensorData)
		if err != nil {
			fmt.Println("erreur :( %s", err.Error())
		} else {
			fmt.Println("sending data ", sensorData)
		}

		client.Publish(airportId, qosLevel, false, json)
		time.Sleep(time.Second)
	}
}
