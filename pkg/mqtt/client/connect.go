package client

import (
	"fmt"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func createClientOptions(brokerUri string, clientId string) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()

	opts.AddBroker(brokerUri)
	opts.SetClientID(clientId)

	return opts
}

func Connect(brokerUri string, clientId string) mqtt.Client {
	fmt.Println("Trying to connect to (" + brokerUri + ", " + clientId + ")...")

	opts := createClientOptions(brokerUri, clientId)
	client := mqtt.NewClient(opts)
	token := client.Connect()

	for !token.WaitTimeout(3 * time.Second) {

	}

	if err := token.Error(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection successful !")

	return client
}
