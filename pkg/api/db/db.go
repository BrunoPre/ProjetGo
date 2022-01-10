package pkg

import (
	. "Project/pkg/api/models"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/gomodule/redigo/redis"
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

	if s.AirportId == "PTP" {
		s.Timestamp = time.Now().Add(24 * time.Hour)
	} else {
		s.Timestamp = time.Now().Add(-24 * time.Hour)
	}
	s.Id = 2

	c := RedisConnect()
	defer c.Close()

	b, err := json.Marshal(s)
	HandleError(err)

	// Save JSON blob to Redis
	reply, err := c.Do("SET", "sensorData:"+strconv.Itoa(s.Id), b)
	HandleError(err)

	fmt.Println("GET ", reply)
}

func SensorByMeasure(measure string) SensorDatas {

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

		if sensorData.Measure == measure {
			sensorDatas = append(sensorDatas, sensorData)
		}
	}
	return sensorDatas
}

func SensorByTime(measure string, timebefore time.Time, timeafter time.Time) SensorDatas {

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

		if sensorData.Timestamp.Before(timeafter) && sensorData.Timestamp.After(timebefore) && sensorData.Measure == measure {
			sensorDatas = append(sensorDatas, sensorData)
		}
	}
	return sensorDatas
}

func SensorAverages() SensorDataAverage {

	c := RedisConnect()
	defer c.Close()

	keys, err := c.Do("KEYS", "sensorData:*")
	HandleError(err)

	var sensorDataAverage SensorDataAverage
	compteurWind := 0.0
	compteurTemp := 0.0
	compteurPressure := 0.0
	sumWind := 0.0
	sumTemp := 0.0
	sumPressure := 0.0

	for _, k := range keys.([]interface{}) {
		var sensorData SensorData

		reply, err := c.Do("GET", k.([]byte))
		HandleError(err)
		if err := json.Unmarshal(reply.([]byte), &sensorData); err != nil {
			panic(err)
		}

		if sensorData.Measure == "Wind speed" {
			compteurWind += 1
			sumWind += sensorData.Value
		}
		if sensorData.Measure == "Temperature" {
			compteurTemp += 1
			sumTemp += sensorData.Value
		}
		if sensorData.Measure == "Atmospheric pressure" {
			compteurPressure += 1
			sumPressure += sensorData.Value
		}
	}

	sensorDataAverage.AverageWind = sumWind / compteurWind
	sensorDataAverage.AverageTemp = sumTemp / compteurTemp
	sensorDataAverage.AveragePressure = sumPressure / compteurPressure

	return sensorDataAverage
}
