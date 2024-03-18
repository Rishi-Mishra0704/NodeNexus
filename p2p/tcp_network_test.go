package p2p

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

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

	// Ensure that all peer connections are closed
	for _, conn := range tn.peers {
		assert.Error(t, conn.Close())
	}
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
func TestTCPNetwork_Start_AlreadyStarted(t *testing.T) {
	tn := setup()
	defer teardown(tn)

	// Attempt to start the network again
	err := tn.Start("8080")

	// Assert that an error is returned
	assert.Error(t, err)
}

func TestTCPNetwork_Connect_Error(t *testing.T) {
	tn := setup()
	defer teardown(tn)

	// Create a mock peer with an invalid address
	mockPeer := &Peer{
		ID:   "peer2",
		Addr: "invalid_address", // Invalid address for testing failure
	}

	// Attempt to connect to the mock peer
	err := tn.Connect(mockPeer)

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, tn.peers[mockPeer.ID])
}
func TestTCPNetwork_Close_NoListener(t *testing.T) {
	tn := NewTCPNetwork()

	// Attempt to close the network when no listener is present
	err := tn.Close()

	// Assert that no error is returned
	assert.NoError(t, err)
}

func TestTCPNetwork_Close_Error(t *testing.T) {
	tn := setup()
	defer teardown(tn)

	// Close the listener to force an error when closing the network
	tn.Listener.Close()

	// Attempt to close the network
	err := tn.Close()

	// Assert that an error is returned
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error closing listener")
}
