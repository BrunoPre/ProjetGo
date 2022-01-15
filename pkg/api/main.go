package main

import (
	router "Project/pkg/api/router"
	"log"
	"net/http"
)

func main() {
	router := router.NewRouter()
	log.Fatal(http.ListenAndServe(":8081", router))
}
