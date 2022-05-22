package main

import (
	"log"
	"net/http"
)

func main() {
	// start the server
	s := newServer()
	s.setupRoutes()
	go s.run()
	log.Printf("starting server on http://localhost:8080/connect")
	http.ListenAndServe(":8080", nil)
}
