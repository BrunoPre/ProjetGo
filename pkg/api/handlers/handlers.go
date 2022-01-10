package pkg

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	. "Project/pkg/api/db"
	. "Project/pkg/api/models"
	mux "github.com/julienschmidt/httprouter"
)

// Logger logs the method, URI, header, and the dispatch time of the request.
func Logger(r *http.Request) {
	start := time.Now()
	log.Printf(
		"%s\t%s\t%q\t%s",
		r.Method,
		r.RequestURI,
		r.Header,
		time.Since(start),
	)
}

// Index handler handles the index at "/" and writes a welcome message.
func Index(w http.ResponseWriter, r *http.Request, _ mux.Params) {
	fmt.Fprintf(w, "<h1 style=\"font-family: Helvetica;\">Hello, welcome to airport service</h1>")
}

// PostIndex handler handles "/posts" and show all the blog posts data as JSON
func SensorDataIndex(w http.ResponseWriter, r *http.Request, _ mux.Params) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	sensorData := FindAll()
	if err := json.NewEncoder(w).Encode(sensorData); err != nil {
		panic(err)
	}
}

// PostShow handler shows the post at "posts/id" as JSON.
func SensorDataShow(w http.ResponseWriter, r *http.Request, ps mux.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	HandleError(err)
	sensorData := FindSensorData(id)
	if err := json.NewEncoder(w).Encode(sensorData); err != nil {
		panic(err)
	}
}

// PostShow handler shows the post at "posts/id" as JSON.
func GetSensorDataByIata(w http.ResponseWriter, r *http.Request, ps mux.Params) {
	sensorData := FindSensorDataByIata(ps.ByName("airportId"))
	if err := json.NewEncoder(w).Encode(sensorData); err != nil {
		panic(err)
	}
}

// PostCreate creates a new post data
func SensorDataCreate(w http.ResponseWriter, r *http.Request, _ mux.Params) {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	HandleError(err)
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	// Save JSON to Post struct (should this be a pointer?)
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

func GetSensorByMeasure(w http.ResponseWriter, r *http.Request, _ mux.Params) {
	measure := r.URL.Query().Get("measure")
	measures := SensorByMeasure(measure)
	if err := json.NewEncoder(w).Encode(measures); err != nil {
		panic(err)
	}
}

func GetTime(w http.ResponseWriter, r *http.Request, _ mux.Params) {
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

func GetAverage(w http.ResponseWriter, r *http.Request, _ mux.Params) {
	averages := SensorAverages()
	if err := json.NewEncoder(w).Encode(averages); err != nil {
		panic(err)
	}
}
