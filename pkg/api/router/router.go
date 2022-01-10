package pkg

import mux "github.com/julienschmidt/httprouter"

// NewRouter creates a Router object for this application.
// to create this api we used a module named httprouter as well as the following link which explains how to use it
// https://medium.com/code-zen/rest-apis-server-in-go-and-redis-66e9cb80a71b
func NewRouter() *mux.Router {
	router := mux.New()
	for _, route := range routes {
		router.Handle(route.Method, route.Pattern, route.Handle)
	}

	return router
}
