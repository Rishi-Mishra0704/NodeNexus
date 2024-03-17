package p2p

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSend(t *testing.T) {
	// Create a new TCPNetwork instance
	tn := setup()
	defer teardown(tn)

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

		// Read message from the connection
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		assert.NoError(t, err)

		// Assertions for received message
		assert.Equal(t, "Hello", string(buffer[:n]))
	}()

	// Establish a connection for the mock peer
	conn, err := net.Dial("tcp", addr)
	assert.NoError(t, err)
	tn.peers[mockPeer.ID] = conn

	// Attempt to send a message to the mock peer
	err = tn.Send(mockPeer, []byte("Hello"))

	// Assertions
	assert.NoError(t, err)
}
