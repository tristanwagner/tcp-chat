package main

import (
	"flag"
	"fmt"
	"log"
	"net"
)

func main() {

	port := flag.Int("p", 1234, "the port number on which the server runs")

	flag.Parse()

	// init server
	s := NewServer()

	// run in go routine
	go s.run()

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("Unable to start server: %s", err.Error())
	}

	fmt.Printf("Started listening on port %d\n", *port)

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
