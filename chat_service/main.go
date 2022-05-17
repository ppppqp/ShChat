package main

import (
	"log"
	"net"
)

func main() {
	s := newServer()
	go s.run()

	// start the tcp server
	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("unable to start server: %s", err.Error())
	}
	defer listener.Close()
	log.Printf("started server on :8888")

	for {
		// infinite loop

		conn, err := listener.Accept()
		// on receive a new tcp connection
		if err != nil {
			log.Printf("unable to accept connection: %s", err.Error())
			continue
		}
		c := s.newClient(conn)
		go c.readInput()
	}
}
