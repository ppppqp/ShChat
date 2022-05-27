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
	log.Printf("starting server on http://localhost:8080/")
	fileServer := http.FileServer(http.Dir("./static/dist")) // New code
	http.Handle("/", fileServer)                             // New code
	http.ListenAndServe(":8080", nil)
}
