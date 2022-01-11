package main

import (
	mqttClient "Project/pkg/mqtt/client"
	"Project/pkg/mqtt/controller"
	"Project/pkg/mqtt/storage"
	"Project/pkg/mqtt/structs"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"sync"
)

// Subscribes to a pub associated to a sensor given by its config file `fileByteJson`
func sub(wg *sync.WaitGroup, brokerUri string, qos byte) {
	defer wg.Done()

	// Init REDIS DB & controller
	storage.Init()
	controller.Init()

	client := mqttClient.Connect(brokerUri, "sub")
	// subscribe to all airports (allowed by '#' wildcard)
	token := client.Subscribe("airport/#", qos, controller.Controller.HandleSensorData)

	token.Wait()
	for {
	}
}

// Get subscription params from a sensor config file
func getSubParams(fileByte []byte) (string, byte) {

	// unmarshalling the JSON config file
	var sensorConfig structs.SensorConfig
	err := json.Unmarshal(fileByte, &sensorConfig)
	if err != nil {
		panic(fmt.Sprintf("Couldn't parse the file \"%e\"", err))
	}

	// instance of sensor config
	s := sensorConfig
	brokerUri := s.BrokerAddr + ":" + strconv.Itoa(s.BrokerPort) // "addr:port"
	qosLevel := byte(s.QosLevel)

	return brokerUri, qosLevel
}

// Check if an element is in an array of string
func exists(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

// Subscribes to all pubs associated to sensors, given a folder full of sensor config JSON files
func main() {
	// get all the args
	args := os.Args
	// argument parsing
	if len(args) != 2 {
		panic("Incorrect arguments lengths. Please provide the path to the config folder")
	}
	fmt.Println(args)

	// get the interesting one only
	arg := os.Args[1]

	files, err := ioutil.ReadDir(arg)
	if err != nil {
		log.Fatal(err)
	}

	wg := new(sync.WaitGroup)

	// avoid instantiate twice a MQTT client for the same broker URI
	var listBrokerURIs []string

	for _, file := range files {
		// if not a JSON config file
		fileExtension := file.Name()[len(file.Name())-5:]
		if fileExtension != ".json" {
			continue
		}

		// read the JSON file
		fileByte, err := ioutil.ReadFile(arg + file.Name())
		if err != nil {
			panic(fmt.Sprintf("couldn't open config file \"%e\"", err))
		}

		// get broker URI & qos
		brokerUri, qos := getSubParams(fileByte)

		if !exists(listBrokerURIs, brokerUri) {
			listBrokerURIs = append(listBrokerURIs, brokerUri)
			// add 1 to the WaitGroup counter
			wg.Add(1)

			// subscribe to the pub
			go sub(wg, brokerUri, qos)
		}
	}
	wg.Wait()

}