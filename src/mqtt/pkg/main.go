package main

import (
	"Project/src/mqtt/pkg/structs"
	json "encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	client := connect("localhost:1883", "golang-client")

	token := client.Subscribe("topic-2", 0, callback)
	token.Wait()
	fmt.Println("Subscribed")
	var sensorData structs.SensorData
	for i := 1; i < 10; i++ {
		sensorData = structs.SensorData{rand.Int() % 50, rand.Int() % 50, structs.Pressure, rand.Float64(), time.Now()}
		json, err := json.Marshal(sensorData)
		if err != nil {
			return
		}
		client.Publish("topic-2", 0, false, json)
		time.Sleep(time.Second)
	}
}

func callback(client mqtt.Client, message mqtt.Message) {
	receivedData := &structs.SensorData{}
	if err := json.Unmarshal(message.Payload(), receivedData); err != nil {
		fmt.Printf("Error unmarshalling data %s\n", err.Error())
		return
	}
	fmt.Println("Received new data (" + receivedData.String() + ")")
}
func createClientOptions(brokerUri string, clientId string) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()

	opts.AddBroker(brokerUri)
	opts.SetClientID(clientId)

	return opts
}

func connect(brokerUri string, clientId string) mqtt.Client {
	fmt.Println("Trying to connect to (" + brokerUri + ", " + clientId + ")...")

	opts := createClientOptions(brokerUri, clientId)
	client := mqtt.NewClient(opts)
	token := client.Connect()

	for !token.WaitTimeout(3 * time.Second) {

	}

	if err := token.Error(); err != nil {
		log.Fatal(err)
	}

	return client
}
