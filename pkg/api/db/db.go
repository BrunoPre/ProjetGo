package pkg

import (
	. "Project/pkg/api/models"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"strconv"
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

//// init seeds some ridiculous initial data
//func init() {
//	CreateSensorData(SensorData{
//		Id:        1,
//		AirportId: "PTP",
//		Measure:   "Temperature",
//		Value:     28.35,
//	})
//	CreateSensorData(SensorData{
//		Id:        2,
//		AirportId: "PTP",
//		Measure:   "Atmospheric pressure",
//		Value:     1200,
//	})
//	CreateSensorData(SensorData{
//		Id:        3,
//		AirportId: "PTP",
//		Measure:   "Wind speed",
//		Value:     25,
//	})
//}

// FindAll return all data
func FindAll() SensorDatas {
	var sensorDatas SensorDatas

	ctx := context.Background()
	c := RedisConnect()

	val, err := c.Keys(ctx, "*").Result()

	for _, key := range val {
		var sensorData SensorData

		result, err := c.Get(ctx, key).Result()
		if err := json.Unmarshal([]byte(result), &sensorData); err != nil {
			panic(err)
		}
		if err != nil {
			panic(err)
		}
		sensorDatas = append(sensorDatas, sensorData)
	}
	if err != nil {
		panic(err)
	}

	return sensorDatas
}

// FindSensorData returns the data from the sensor whose id was passed as a parameter
func FindSensorData(id int) SensorDatas {
	var sensorDatas SensorDatas

	ctx := context.Background()
	c := RedisConnect()

	val, err := c.Keys(ctx, strconv.Itoa(id)+"*").Result()

	for _, key := range val {
		var sensorData SensorData

		result, err := c.Get(ctx, key).Result()
		if err := json.Unmarshal([]byte(result), &sensorData); err != nil {
			panic(err)
		}
		if err != nil {
			panic(err)
		}
		sensorDatas = append(sensorDatas, sensorData)
	}
	if err != nil {
		panic(err)
	}

	return sensorDatas
}

// FindSensorDataByIata returns the data from the sensor whose iata code was passed as a parameter
func FindSensorDataByIata(iata string) SensorDatas {
	var sensorDatasBis SensorDatas
	var sensorDatas SensorDatas

	sensorDatasBis = FindAll()

	for _, data := range sensorDatasBis {
		if data.AirportId == iata {
			sensorDatas = append(sensorDatas, data)
		}
	}
	return sensorDatas
}

// CreateSensorData creates a sensor data.
func CreateSensorData(data SensorData) error {

	data.Timestamp = time.Now()

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
func SensorByMeasure(measure string) SensorDatas {
	var sensorDatasBis SensorDatas
	var sensorDatas SensorDatas

	sensorDatasBis = FindAll()

	for _, k := range sensorDatasBis {
		if k.Measure == measure {
			sensorDatas = append(sensorDatas, k)
		}
	}
	return sensorDatas
}

//SensorByTime returns all the sensors of a type of measurement and which are included between two dates
func SensorByTime(measure string, timebefore time.Time, timeafter time.Time) SensorDatas {
	var sensorDatasBis SensorDatas
	var sensorDatas SensorDatas

	sensorDatasBis = FindAll()

	for _, data := range sensorDatasBis {
		if data.Timestamp.Before(timeafter) && data.Timestamp.After(timebefore) && data.Measure == measure {
			sensorDatas = append(sensorDatas, data)
		}
	}
	return sensorDatas
}

// SensorAverages returns the average of each type of sensor for a given day
func SensorAverages(date time.Time) SensorDataAverage {
	var sensorDatas SensorDatas
	var sensorDataAverage SensorDataAverage

	sensorDatas = FindAll()

	compteurWind := 0.0
	compteurTemp := 0.0
	compteurPressure := 0.0
	sumWind := 0.0
	sumTemp := 0.0
	sumPressure := 0.0

	for _, data := range sensorDatas {
		if data.Timestamp.Before(date.Add(24*time.Hour)) && data.Timestamp.After(date) {
			if data.Measure == "Wind speed" {
				compteurWind += 1
				sumWind += data.Value
			}
			if data.Measure == "Temperature" {
				compteurTemp += 1
				sumTemp += data.Value
			}
			if data.Measure == "Atmospheric pressure" {
				compteurPressure += 1
				sumPressure += data.Value
			}
		}
	}

	if sumWind != 0.0 && compteurWind != 0.0 {
		sensorDataAverage.AverageWind = sumWind / compteurWind
	} else {
		sensorDataAverage.AverageWind = 0
	}
	if sumTemp != 0.0 && compteurTemp != 0.0 {
		sensorDataAverage.AverageTemp = sumTemp / compteurTemp
	} else {
		sensorDataAverage.AverageTemp = 0
	}
	if sumPressure != 0.0 && compteurPressure != 0.0 {
		sensorDataAverage.AveragePressure = sumPressure / compteurPressure
	} else {
		sensorDataAverage.AveragePressure = 0
	}

	return sensorDataAverage
}
