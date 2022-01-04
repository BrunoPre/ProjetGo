package main

import (
	"ClientPaho"
	"sensors"
	"subs"
)

func main() {
	client := ClientPaho.Connect("tcp://localhost:1883", "myClientId")
	subs.SubTest(client)
	sensors.InitAndPublishAllSensors(client)
}
