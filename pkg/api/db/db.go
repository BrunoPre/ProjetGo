package pkg

import (
	. "Project/pkg/api/models"
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"strconv"
	"time"
)

var (
	currentSensorDatasID int
)

// RedisConnect connects to a default redis server at port 6379
func RedisConnect() redis.Conn {
	c, err := redis.Dial("tcp", ":6379")
	HandleError(err)
	return c
}

//// init seeds some ridiculous initial data
//func init() {
//	CreateSensorData(SensorData{
//		AirportId: "PTP",
//		Measure:   "Temperature",
//		Value:     15.20,
//	})
//	CreateSensorData(SensorData{
//		AirportId: "PTP",
//		Measure:   "Atmospheric pressure",
//		Value:     1220,
//	})
//	CreateSensorData(SensorData{
//		AirportId: "PTP",
//		Measure:   "Wind speed",
//		Value:     65,
//	})
//	CreateSensorData(SensorData{
//		AirportId: "NTE",
//		Measure:   "Temperature",
//		Value:     15.20,
//	})
//	CreateSensorData(SensorData{
//		AirportId: "NTE",
//		Measure:   "Atmospheric pressure",
//		Value:     1220,
//	})
//	CreateSensorData(SensorData{
//		AirportId: "NTE",
//		Measure:   "Wind speed",
//		Value:     65,
//	})
//}

func FindAll() SensorDatas {

	c := RedisConnect()
	defer c.Close()

	keys, err := c.Do("KEYS", "sensorData:*")
	HandleError(err)

	var sensorDatas SensorDatas

	for _, k := range keys.([]interface{}) {
		var sensorData SensorData

		reply, err := c.Do("GET", k.([]byte))
		HandleError(err)
		if err := json.Unmarshal(reply.([]byte), &sensorData); err != nil {
			panic(err)
		}

		sensorDatas = append(sensorDatas, sensorData)
	}
	return sensorDatas
}

func FindSensorData(id int) SensorData {
	var sensorData SensorData

	c := RedisConnect()
	defer c.Close()
	reply, err := c.Do("GET", "sensorData:"+strconv.Itoa(id))
	HandleError(err)
	if err = json.Unmarshal(reply.([]byte), &sensorData); err != nil {
		panic(err)
	}
	return sensorData
}

func FindSensorDataByIata(iata string) SensorDatas {
	c := RedisConnect()
	defer c.Close()

	keys, err := c.Do("KEYS", "sensorData:*")
	HandleError(err)

	var sensorDatas SensorDatas

	for _, k := range keys.([]interface{}) {
		var sensorData SensorData

		reply, err := c.Do("GET", k.([]byte))
		HandleError(err)
		if err := json.Unmarshal(reply.([]byte), &sensorData); err != nil {
			panic(err)
		}

		if sensorData.AirportId == iata {
			sensorDatas = append(sensorDatas, sensorData)
		}
	}
	return sensorDatas
}

// CreateSensorData creates a sensor data.
func CreateSensorData(s SensorData) {
	currentSensorDatasID++

	s.Id = currentSensorDatasID
	s.Timestamp = time.Now()

	c := RedisConnect()
	defer c.Close()

	b, err := json.Marshal(s)
	HandleError(err)

	// Save JSON blob to Redis
	reply, err := c.Do("SET", "sensorData:"+strconv.Itoa(s.Id), b)
	HandleError(err)

	fmt.Println("GET ", reply)
}
