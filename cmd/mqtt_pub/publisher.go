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

	var sensorConfig structs.SensorConfig
	err = json.Unmarshal(fileByte, &sensorConfig)
	if err != nil {
		panic(fmt.Sprintf("Couldn't parse the file \"%w\"", err))
	}

	s := sensorConfig
	fmt.Println(s.String())

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
