package main

import (
	mqttClient "Project/pkg/mqtt/client"
	"Project/pkg/mqtt/structs"
	json "encoding/json"
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	client := mqttClient.Connect("localhost:1883", "golang-pub")
	token := client.Subscribe("topic-2", 0, callback)
	token.Wait()
	for {
	}
}

func callback(client mqtt.Client, message mqtt.Message) {
	receivedData := &structs.SensorData{}
	if err := json.Unmarshal(message.Payload(), receivedData); err != nil {
		fmt.Printf("Error unmarshalling data %s\n", err.Error())
	}
	fmt.Println("Received new data (" + receivedData.String() + ")")
}
