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
	fmt.Fprintf(w, "<h1 style=\"font-family: Helvetica;\">Hello, welcome to blog service</h1>")
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
	Logger(r)
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

	if err := json.NewEncoder(w).Encode(err); err != nil {
		panic(err)
	}

	CreateSensorData(sensorData)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
}
