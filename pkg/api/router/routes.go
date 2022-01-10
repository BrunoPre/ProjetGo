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
	// It is a good idea to name the fields when declaring a struct object. SensorDataShow
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"PostIndex",
		"GET",
		"/sensors",
		SensorDataIndex,
	},
	Route{
		"PostCreate",
		"POST",
		"/sensors",
		SensorDataCreate,
	},
	Route{
		"SensorDataShow",
		"GET",
		"/sensors/:id",
		SensorDataShow,
	},
	Route{
		"GetSensorDataByIata",
		"GET",
		"/sensors-iata/:airportId",
		GetSensorDataByIata,
	},
	Route{
		"PostMeasure",
		"GET",
		"/sensors-measure",
		GetSensorByMeasure,
	},
	Route{
		"Time",
		"GET",
		"/time",
		GetTime,
	},
	Route{
		"Average",
		"GET",
		"/average",
		GetAverage,
	},
}
