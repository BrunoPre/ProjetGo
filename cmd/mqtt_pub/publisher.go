package main

import (
	mqttClient "Project/pkg/mqtt/client"
	"Project/pkg/mqtt/structs"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		panic("Incorrect arguments lengths. Please provide the path to the config file")
	}
	fmt.Println(os.Args)

	fileByte, err := ioutil.ReadFile(filepath.Clean(os.Args[1]))
	if err != nil {
		panic(fmt.Sprintf("Couldn't open config file \"%w\"", err))
	}

	var sensorsConfigs structs.SensorsConfigs
	err = json.Unmarshal(fileByte, &sensorsConfigs)
	if err != nil {
		panic(fmt.Sprintf("Couldn't parse the file \"%w\"", err))
	}

	// 2 printing ways
	for i := 0; i < len(sensorsConfigs.SensorsConfigs); i++ {
		s := sensorsConfigs.SensorsConfigs[i]
		fmt.Println("clientId: " + strconv.Itoa(s.ClientId))
		fmt.Println("brokenAddr: " + s.BrokerAddr)
		fmt.Println("brokerPort: " + strconv.Itoa(s.BrokerPort))
		fmt.Println("qosLevel: " + strconv.Itoa(s.QosLevel))
		fmt.Println("measureType: " + s.MeasureType)
		fmt.Println("airportId: " + strconv.Itoa(s.AirportId))
	}
	fmt.Println(sensorsConfigs)

	client := mqttClient.Connect("localhost:1883", "golang-sub")
	var sensorData structs.SensorData
	for {
		sensorData = structs.SensorData{rand.Int() % 50, rand.Int() % 50, structs.Pressure, rand.Float64(), time.Now()}
		json, err := json.Marshal(sensorData)
		if err != nil {
			fmt.Println("erreur :( %s", err.Error())
		} else {
			fmt.Println("sending data ", sensorData)
		}

		client.Publish("topic-2", 0, false, json)
		time.Sleep(time.Second)
	}
}
