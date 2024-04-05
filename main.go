package main

import (
	"log"
	"net"

	"github.com/Rishi-Mishra0704/NodeNexus/p2p"
)

func main() {
	s := p2p.NewServer()
	go s.Run()

	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("unable to start server: %s", err.Error())
	}

	defer listener.Close()
	log.Printf("server started on :8888")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("failed to accept connection: %s", err.Error())
			continue
		}

		c := s.NewClient(conn)
		go c.ReadInput()
	}
}
