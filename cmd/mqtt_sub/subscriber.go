package main

import (
	mqttClient "Project/pkg/mqtt/client"
	"Project/pkg/mqtt/controller"
	"Project/pkg/mqtt/storage"
)

func main() {
	storage.Init()
	controller.Init()

	client := mqttClient.Connect("localhost:1883", "golang-pub")
	token := client.Subscribe("topic-2", 0, controller.Controller.HandleSensorData)

	token.Wait()
	for {
	}
}
