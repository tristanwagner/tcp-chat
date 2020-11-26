package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	// init server
	s := NewServer()

	// run in go routine
	go s.run()

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatalf("Unable to start server: %s", err.Error())
	}

	fmt.Printf("Started listening on port %s\n", "1234")

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("Unable to accept connection: %s", err.Error())
			continue
		}

		go s.newClient(conn)
	}
}
