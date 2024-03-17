package p2p_test

import (
	"fmt"
	"net"
	"testing"

	"github.com/Rishi-Mishra0704/NodeNexus/p2p"
	"github.com/stretchr/testify/assert"
)

func TestPeerConnect_Success(t *testing.T) {
	// Mock peer data
	mockPeer := &p2p.Peer{
		ID:   "peer1",
		Addr: "127.0.0.1:8080",
	}

	// Mock TCP connection
	mockConn, err := net.Dial("tcp", "localhost:8080")
	assert.NoError(t, err)
	defer mockConn.Close()

	// Attempt to connect to the mock peer
	conn, err := mockPeer.Connect()

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, conn)
}
func TestPeerConnect_Failure(t *testing.T) {
	// Mock peer data with invalid address
	mockPeer := &p2p.Peer{
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
