package main

import (
	"net/http"
)

const (
	// Version is the current version of the application
	Version        = "1.0.0"
	HttpServerPort = ":8080"
)

func main() {
	// Start the web server
	router := http.NewServeMux()

	router.HandleFunc("/", nil)

	err := http.ListenAndServe(HttpServerPort, router)
	if err != nil {
		panic(err.Error())
	}
}
