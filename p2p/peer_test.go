package p2p

import (
	"fmt"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPeerConnect_Success(t *testing.T) {
	addr, err := findAvailablePort()
	assert.NoError(t, err)

	// Create a new peer with the mock address
	mockPeer := NewPeer("peer1", addr)

	// Mock TCP server
	mockServer, err := net.Listen("tcp", addr)
	assert.NoError(t, err)
	defer mockServer.Close()

	// Attempt to connect to the mock peer
	conn, err := mockPeer.Connect()

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, conn)
}

func TestPeerConnect_Failure(t *testing.T) {
	// Create a new peer with an invalid address
	mockPeer := NewPeer("peer2", "invalid_address")

	// Attempt to connect to the mock peer
	conn, err := mockPeer.Connect()

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, conn)
	fmt.Println("Error:", err)
}
