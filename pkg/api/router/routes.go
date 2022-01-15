package pkg

import (
	. "Project/pkg/api/handlers"
	mux "github.com/julienschmidt/httprouter"
)

// Route represents a URL that serves a specific resource.
type Route struct {
	Name    string
	Method  string
	Pattern string
	Handle  mux.Handle
}

// Routes are a list of Routes for this application.
type Routes []Route

var routes = Routes{
	//Route{
	//	"Index",
	//	"GET",
	//	"/",
	//	Index,
	//},
	//Route{
	//	"GetSensors",
	//	"GET",
	//	"/sensors",
	//	GetSensors,
	//},
	Route{
		"PostSensor",
		"POST",
		"/sensors",
		PostSensor,
	},
	Route{
		"GetSensor",
		"GET",
		"/sensors/:id",
		GetSensor,
	},
	//Route{
	//	"GetSensorDataByIata",
	//	"GET",
	//	"/sensors-iata/:airportId",
	//	GetSensorDataByIata,
	//},
	//Route{
	//	"GetSensorByMeasure",
	//	"GET",
	//	"/sensors-measure",
	//	GetSensorByMeasure,
	//},
	//Route{
	//	"GetSensorByTime",
	//	"GET",
	//	"/time",
	//	GetSensorByTime,
	//},
	//Route{
	//	"GetAverage",
	//	"GET",
	//	"/average",
	//	GetAverage,
	//},
}
