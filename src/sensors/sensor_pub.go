package sensors

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"time"
)

func InitAndPublishAllSensors(client mqtt.Client) {
	
	tempSensor := NewSensor(0, 111, "temp")
	windSensor := NewSensor(1, 111, "wind")
	pressureSensor := NewSensor(2, 111, "pressure")

	sensors := [...]*Sensor{tempSensor, windSensor, pressureSensor}

	for {
		durationSleepInt := 1
		durationSleepTime := time.Duration(durationSleepInt) * time.Second
		for i := 0; i < len(sensors); i++ {
			sensors[i].publishOnce(client)
		}
		time.Sleep(durationSleepTime) // pub data every second according to `durationSleepInt` value
	}

	client.Disconnect(250)

}
