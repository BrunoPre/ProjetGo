package pkg

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	. "Project/pkg/api/db"
	. "Project/pkg/api/models"
	mux "github.com/julienschmidt/httprouter"
)

// Index handler handles the index at "/" and writes a welcome message.
func Index(w http.ResponseWriter, r *http.Request, _ mux.Params) {
	fmt.Fprintf(w, "<h1 style=\"font-family: Helvetica;\">Hello, welcome to airport service</h1>")
}

// GetSensors handler handles "/sensors" and show all the sensors data as JSON
func GetSensors(w http.ResponseWriter, r *http.Request, _ mux.Params) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	sensorData := FindAll()
	if err := json.NewEncoder(w).Encode(sensorData); err != nil {
		panic(err)
	}
}

// GetSensor handler shows the sensor at "sensor/id" as JSON.
func GetSensor(w http.ResponseWriter, r *http.Request, ps mux.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	HandleError(err)
	sensorData := FindSensorData(id)
	if err := json.NewEncoder(w).Encode(sensorData); err != nil {
		panic(err)
	}
}

// GetSensorDataByIata handler shows the sensor by IATA at "/sensors-iata/airportId" as JSON.
func GetSensorDataByIata(w http.ResponseWriter, r *http.Request, ps mux.Params) {
	sensorData := FindSensorDataByIata(ps.ByName("airportId"))
	if err := json.NewEncoder(w).Encode(sensorData); err != nil {
		panic(err)
	}
}

// PostSensor creates a new Sensor data
func PostSensor(w http.ResponseWriter, r *http.Request, _ mux.Params) {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	HandleError(err)
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	// Save JSON to SensorData struct
	var sensorData SensorData
	if err := json.Unmarshal(body, &sensorData); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
	}
	fmt.Println(sensorData)

	if err := json.NewEncoder(w).Encode(err); err != nil {
		panic(err)
	}

	CreateSensorData(sensorData)
}

// GetSensorByMeasure handler shows the sensor by measure at "/sensors-measure" as JSON.
func GetSensorByMeasure(w http.ResponseWriter, r *http.Request, _ mux.Params) {
	measure := r.URL.Query().Get("measure")
	measures := SensorByMeasure(measure)
	if err := json.NewEncoder(w).Encode(measures); err != nil {
		panic(err)
	}
}

// GetSensorByTime handler shows the sensor between two time and one type of measure at "/time" as JSON.
func GetSensorByTime(w http.ResponseWriter, r *http.Request, _ mux.Params) {
	layout := "2006-01-02T15:04:05.00"

	str := r.URL.Query().Get("timebefore")
	timebefore, err := time.Parse(layout, str)
	if err != nil {
		panic(err)
	}
	str1 := r.URL.Query().Get("timeafter")
	timeafter, err := time.Parse(layout, str1)
	if err != nil {
		panic(err)
	}

	measure := r.URL.Query().Get("measure")

	times := SensorByTime(measure, timebefore, timeafter)
	if err := json.NewEncoder(w).Encode(times); err != nil {
		panic(err)
	}
}

// GetAverage handler shows the average of each type of sensor for a given day at "/average" as JSON.
func GetAverage(w http.ResponseWriter, r *http.Request, _ mux.Params) {
	layout := "2006-01-02"

	str := r.URL.Query().Get("date")
	timebefore, err := time.Parse(layout, str)
	if err != nil {
		panic(err)
	}

	averages := SensorAverages(timebefore)
	if err := json.NewEncoder(w).Encode(averages); err != nil {
		panic(err)
	}
}
