package subs

import (
	"ClientPaho"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"sensors"
)

func SubTest(client mqtt.Client) {
	topic := "sensors/wind"

	rangesData := sensors.InitSensorData()
	_, _ = sensors.GenerateRandomData("wind", rangesData)

	token := client.Subscribe(topic, 0, ClientPaho.RespondToPub)
	token.Wait()
	fmt.Printf("Subscribed to topic: %s\n", topic)
}

func sub(client mqtt.Client) {
	topic := "sensors/"
	nb_sensors := 3
	topics := [...]string{"temp", "wind", "pressure"}
	for i := 0; i < nb_sensors; i++ {
		topics[i] = topic + topics[i]
	}
}
