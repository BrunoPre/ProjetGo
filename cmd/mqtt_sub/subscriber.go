package main

import (
	mqttClient "Project/pkg/mqtt/client"
	"Project/pkg/mqtt/configuration"
	"Project/pkg/mqtt/controller"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"syscall"
)

func main() {
	args := os.Args
	// argument parsing
	if len(args) != 2 {
		panic("Incorrect arguments lengths. Please provide the path to the config file")
	}
	fmt.Println(args)

	arg := os.Args[1]

	fileByte, err := ioutil.ReadFile(filepath.Clean(arg))
	if err != nil {
		panic(fmt.Sprintf("Couldn't open config file \"%w\"", err))
	}

	// unmarshalling the JSON config file
	var sensorConfig configuration.Config
	err = json.Unmarshal(fileByte, &sensorConfig)
	if err != nil {
		panic(fmt.Sprintf("Couldn't parse the file \"%w\"", err))
	}

	brokerUri := sensorConfig.MqttConf.BrokerAddr + ":" + strconv.Itoa(sensorConfig.MqttConf.BrokerPort) // "addr:port"
	qosLevel := byte(sensorConfig.MqttConf.QosLevel)

	var sensorController controller.SensorController
	if sensorController, err = controller.FactoryControllerDao(sensorConfig); err != nil {
		panic(err)
	}
	client := mqttClient.Connect(brokerUri, sensorConfig.MqttConf.ClientName)
	// subscribe to all airports (allowed by '#' wildcard)
	token := client.Subscribe("airport/#", qosLevel, sensorController.HandleSensorData)

	token.Wait()

	// Si écoute les signaux de terminaison pour déconnecter le client en cas d'arrêt du programme.
	// defer ne fonctionne qu'en cas d'arrêt normal du programme (sortie de bloc par exemple)
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	go func() {
		<-sigs
		client.Disconnect(0)
	}()

	for {
	}

}
