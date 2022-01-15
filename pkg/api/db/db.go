package pkg

import (
	. "Project/pkg/api/models"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"strings"
	"time"
)

// RedisConnect connects to a default redis server at port 6379
func RedisConnect() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return rdb
}

// init seeds some ridiculous initial data
func init() {
	CreateSensorData(SensorData{
		Id:        1,
		AirportId: "PTP",
		Measure:   "Temperature",
		Value:     28.35,
	})
	CreateSensorData(SensorData{
		Id:        2,
		AirportId: "PTP",
		Measure:   "Atmospheric pressure",
		Value:     1200,
	})
	CreateSensorData(SensorData{
		Id:        3,
		AirportId: "PTP",
		Measure:   "Wind speed",
		Value:     25,
	})
	CreateSensorData(SensorData{
		Id:        4,
		AirportId: "NTE",
		Measure:   "Temperature",
		Value:     18.80,
	})
	CreateSensorData(SensorData{
		Id:        5,
		AirportId: "NTE",
		Measure:   "Atmospheric pressure",
		Value:     1220,
	})
	CreateSensorData(SensorData{
		Id:        6,
		AirportId: "NTE",
		Measure:   "Wind speed",
		Value:     75,
	})
	CreateSensorData(SensorData{
		Id:        7,
		AirportId: "ORY",
		Measure:   "Temperature",
		Value:     15.20,
	})
	CreateSensorData(SensorData{
		Id:        8,
		AirportId: "ORY",
		Measure:   "Atmospheric pressure",
		Value:     1240,
	})
	CreateSensorData(SensorData{
		Id:        9,
		AirportId: "ORY",
		Measure:   "Wind speed",
		Value:     65,
	})
}

// FindAll return all data
//func FindAll() SensorDatas {
//
//	c := RedisConnect()
//	defer c.Close()
//
//	keys, err := c.Do("KEYS", "sensorData:*")
//	HandleError(err)
//
//	var sensorDatas SensorDatas
//
//	for _, k := range keys.([]interface{}) {
//		var sensorData SensorData
//
//		reply, err := c.Do("GET", k.([]byte))
//		HandleError(err)
//		if err := json.Unmarshal(reply.([]byte), &sensorData); err != nil {
//			panic(err)
//		}
//
//		sensorDatas = append(sensorDatas, sensorData)
//	}
//	return sensorDatas
//}

// FindSensorData returns the data from the sensor whose id was passed as a parameter
func FindSensorData(id int) SensorData {
	var sensorData SensorData

	ctx := context.Background()
	c := RedisConnect()

	iter := c.Scan(ctx, 0, "prefix:9", 0).Iterator()
	for iter.Next(ctx) {
		fmt.Println("keys", iter.Val())
	}
	if err := iter.Err(); err != nil {
		panic(err)
	}

	//val, err := c.Get(ctx,"9_ORY_Jan_14_15:25:20").Result()
	//fmt.Println(val)
	//if err != nil {
	//	panic(err)
	//}

	return sensorData
}

// FindSensorDataByIata returns the data from the sensor whose iata code was passed as a parameter
//func FindSensorDataByIata(iata string) SensorDatas {
//	c := RedisConnect()
//	defer c.Close()
//
//	keys, err := c.Do("KEYS", "sensorData:*")
//	HandleError(err)
//
//	var sensorDatas SensorDatas
//
//	for _, k := range keys.([]interface{}) {
//		var sensorData SensorData
//
//		reply, err := c.Do("GET", k.([]byte))
//		HandleError(err)
//		if err := json.Unmarshal(reply.([]byte), &sensorData); err != nil {
//			panic(err)
//		}
//
//		if sensorData.AirportId == iata {
//			sensorDatas = append(sensorDatas, sensorData)
//		}
//	}
//	return sensorDatas
//}

// CreateSensorData creates a sensor data.
func CreateSensorData(data SensorData) error {

	if data.AirportId == "PTP" {
		data.Timestamp = time.Now().Add(24 * time.Hour)
	} else if data.AirportId == "ORY" {
		data.Timestamp = time.Now().Add(-24 * time.Hour)
	} else {
		data.Timestamp = time.Now()
	}

	ctx := context.Background()
	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("Error while unmarshalling the data, %w", err)
	}
	c := RedisConnect()
	fmt.Println(c.Set(ctx, fmt.Sprintf("%d_%s_%s", data.Id, data.AirportId, strings.Replace(data.Timestamp.Format(time.Stamp), " ", "_", -1)), jsonData, 0))

	return nil
}

// SensorByMeasure returns all the sensors of a type of measure
//func SensorByMeasure(measure string) SensorDatas {
//	c := RedisConnect()
//	defer c.Close()
//
//	keys, err := c.Do("KEYS", "sensorData:*")
//	HandleError(err)
//
//	var sensorDatas SensorDatas
//
//	for _, k := range keys.([]interface{}) {
//		var sensorData SensorData
//
//		reply, err := c.Do("GET", k.([]byte))
//		HandleError(err)
//		if err := json.Unmarshal(reply.([]byte), &sensorData); err != nil {
//			panic(err)
//		}
//
//		if sensorData.Measure == measure {
//			sensorDatas = append(sensorDatas, sensorData)
//		}
//	}
//	return sensorDatas
//}

// SensorByTime returns all the sensors of a type of measurement and which are included between two dates
//func SensorByTime(measure string, timebefore time.Time, timeafter time.Time) SensorDatas {
//
//	c := RedisConnect()
//	defer c.Close()
//
//	keys, err := c.Do("KEYS", "sensorData:*")
//	HandleError(err)
//
//	var sensorDatas SensorDatas
//
//	for _, k := range keys.([]interface{}) {
//		var sensorData SensorData
//
//		reply, err := c.Do("GET", k.([]byte))
//		HandleError(err)
//		if err := json.Unmarshal(reply.([]byte), &sensorData); err != nil {
//			panic(err)
//		}
//
//		if sensorData.Timestamp.Before(timeafter) && sensorData.Timestamp.After(timebefore) && sensorData.Measure == measure {
//			sensorDatas = append(sensorDatas, sensorData)
//		}
//	}
//	return sensorDatas
//}

// SensorAverages returns the average of each type of sensor for a given day
//func SensorAverages(date time.Time) SensorDataAverage {
//
//	c := RedisConnect()
//	defer c.Close()
//
//	keys, err := c.Do("KEYS", "sensorData:*")
//	HandleError(err)
//
//	var sensorDataAverage SensorDataAverage
//	compteurWind := 0.0
//	compteurTemp := 0.0
//	compteurPressure := 0.0
//	sumWind := 0.0
//	sumTemp := 0.0
//	sumPressure := 0.0
//
//	for _, k := range keys.([]interface{}) {
//		var sensorData SensorData
//
//		reply, err := c.Do("GET", k.([]byte))
//		HandleError(err)
//		if err := json.Unmarshal(reply.([]byte), &sensorData); err != nil {
//			panic(err)
//		}
//
//		if sensorData.Timestamp.Before(date.Add(24*time.Hour)) && sensorData.Timestamp.After(date) {
//			if sensorData.Measure == "Wind speed" {
//				compteurWind += 1
//				sumWind += sensorData.Value
//			}
//			if sensorData.Measure == "Temperature" {
//				compteurTemp += 1
//				sumTemp += sensorData.Value
//			}
//			if sensorData.Measure == "Atmospheric pressure" {
//				compteurPressure += 1
//				sumPressure += sensorData.Value
//			}
//		}
//	}
//
//	sensorDataAverage.AverageWind = sumWind / compteurWind
//	sensorDataAverage.AverageTemp = sumTemp / compteurTemp
//	sensorDataAverage.AveragePressure = sumPressure / compteurPressure
//
//	return sensorDataAverage
//}
