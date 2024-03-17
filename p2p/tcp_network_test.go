package p2p

import (
	"log"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setup() *TCPNetwork {
	tn := NewTCPNetwork()
	err := tn.Start()
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

func TestTCPNetwork_Start(t *testing.T) {
	tn := setup()
	defer teardown(tn)

	assert.NotNil(t, tn.Listener)
}

func TestTCPNetwork_Close(t *testing.T) {
	tn := setup()
	defer teardown(tn)

	err := tn.Close()
	assert.NoError(t, err)
}

func TestTCPNetwork_Connect(t *testing.T) {
	// Create a new TCPNetwork instance
	tn := NewTCPNetwork()

	// Find an available port
	addr, err := findAvailablePort()
	assert.NoError(t, err)

	// Mock peer data
	mockPeer := &Peer{
		ID:   "peer1",
		Addr: addr,
	}

	// Start a mock server to accept connections
	mockServer, err := net.Listen("tcp", addr)
	assert.NoError(t, err)
	defer mockServer.Close()

	// Start a goroutine to handle incoming connections
	go func() {
		conn, err := mockServer.Accept()
		assert.NoError(t, err)
		defer conn.Close()
	}()

	// Attempt to connect to the mock peer
	err = tn.Connect(mockPeer)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, tn.peers[mockPeer.ID])
}
