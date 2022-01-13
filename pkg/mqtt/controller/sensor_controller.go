package controller

import (
	"Project/pkg/mqtt/configuration"
	"Project/pkg/mqtt/storage"
	"Project/pkg/mqtt/structs"
	"encoding/json"
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// Structure d'un contrôleur utilisant un DAO pour y écrire les données.
// Le contrôleur à pour but de vérifier la validité des données
type SensorController struct {
	dao storage.SensorDao
}

func (s SensorController) HandleSensorData(_ mqtt.Client, message mqtt.Message) {
	receivedData := &structs.SensorData{}
	if err := json.Unmarshal(message.Payload(), receivedData); err != nil {
		fmt.Printf("Error unmarshalling data %s\n", err.Error())
	}

	if err := s.dao.Write(*receivedData); err != nil {
		fmt.Println(err)
	}
	//s.dao.WriteCSV(*receivedData)
	fmt.Println("Received new data (" + receivedData.String() + ")")
}

func FactoryControllerDao(config configuration.Config) (SensorController, error) {
	var (
		dao storage.SensorDao
		err error
	)

	if dao, err = storage.FactorySensorDao(config.Dao); err != nil {
		return SensorController{}, err
	}

	return SensorController{dao}, nil
}
