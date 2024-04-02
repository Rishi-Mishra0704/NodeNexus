package main

import (
	"log"

	"github.com/Rishi-Mishra0704/NodeNexus/p2p"
)

func main() {
	// Define the TCP transport options
	tcpOpts := p2p.TcpTransportOpts{
		ListenAddr:    ":8080", // Listening on port 8080
		HandShakeFunc: p2p.DefaultHandShakeFunc,
	}

	// Create a new TCP transport instance
	tcpTransport := p2p.NewTCPTransport(tcpOpts)

	// Start listening for incoming connections
	if err := tcpTransport.ListenAndAccept(); err != nil {
		log.Fatalf("Failed to start TCP transport: %v", err)
	}

	// Log a message indicating that the TCP transport is listening
	log.Println("TCP transport started and listening for incoming connections on port 8080")

	// Keep the main goroutine alive
	select {}
}
