package p2p

import (
	"bufio"
	"fmt"
	"net"
)

type Server struct {
	Address string // Server address (e.g., ":8080")
}

// NewServer creates a new Server instance with the specified address.
func NewServer(address string) *Server {
	return &Server{Address: address}
}

// Start starts the TCP server.
func (s *Server) Start() error {
	// Listen for incoming connections
	listener, err := net.Listen("tcp", s.Address)
	if err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}
	defer listener.Close()

	// Accept incoming connections in a loop
	for {
		conn, err := listener.Accept()
		if err != nil {
			return fmt.Errorf("failed to accept incoming connection: %w", err)
		}

		// Handle incoming connection in a separate goroutine
		go s.handleConnection(conn)
	}
}

// handleConnection handles an incoming connection.
func (s *Server) handleConnection(conn net.Conn) {
	defer conn.Close()

	// Read incoming data from the connection
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		// Print received data
		fmt.Printf("Received data: %s\n", scanner.Text())

		// Echo the received data back to the client
		_, err := fmt.Fprintf(conn, "Echo: %s\n", scanner.Text())
		if err != nil {
			fmt.Printf("Error sending data: %v\n", err)
			return
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading data: %v\n", err)
		return
	}
}
