package main

import (
	mqttClient "Project/pkg/mqtt/client"
	"Project/pkg/mqtt/structs"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

func main() {

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