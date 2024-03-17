package p2p

import (
	"log"
	"net"
)

func setup() *TCPNetwork {
	tn := NewTCPNetwork()
	err := tn.Start("8080")
	if err != nil {
		log.Fatalf("Failed to start TCP network: %v", err)
	}
	return tn
}

func teardown(tn *TCPNetwork) {
	err := tn.Close()
	if err != nil {
		log.Fatalf("Failed to close TCP network: %v", err)
	}
}

func findAvailablePort() (string, error) {
	// Listen on a random port to find an available one
	listener, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		return "", err
	}
	defer listener.Close()

	// Get the address of the listener
	addr := listener.Addr().String()
	return addr, nil
}
