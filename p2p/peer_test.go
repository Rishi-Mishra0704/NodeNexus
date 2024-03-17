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
	// Mock peer data
	mockPeer := &Peer{
		ID:   "peer1",
		Addr: addr,
	}

	// Mock TCP connection
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
	// Mock peer data with invalid address
	mockPeer := &Peer{
		ID:   "peer2",
		Addr: "invalid_address", // Invalid address for testing failure
	}

	// Attempt to connect to the mock peer
	conn, err := mockPeer.Connect()

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, conn)
	fmt.Println("Error:", err)
}
