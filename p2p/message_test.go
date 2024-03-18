package p2p

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPNetwork_Send(t *testing.T) {
	// Setup
	tn := setup()
	defer teardown(tn)

	// Find an available port for the mock peer
	mockPeerAddr, err := findAvailablePort()
	assert.NoError(t, err)

	// Create a new TCPNetwork instance for the mock peer
	mockPeerNetwork := NewTCPNetwork()

	// Start the mock peer network on the available port
	err = mockPeerNetwork.Start(mockPeerAddr[strings.LastIndex(mockPeerAddr, ":")+1:])
	assert.NoError(t, err)
	defer mockPeerNetwork.Close()

	// Create a mock peer
	mockPeer := &Peer{
		ID:      "peer1",
		Addr:    mockPeerAddr,
		Message: make(chan []byte, 1), // Buffered channel to avoid blocking
	}

	// Connect the mock peer to the network
	err = tn.Connect(mockPeer)
	assert.NoError(t, err)

	// Define the message to send
	message := []byte("Hello, world!")

	// Attempt to send the message to the mock peer (successful)
	err = tn.Send(mockPeer, message)
	assert.NoError(t, err)

	// Create another mock peer
	mockPeerNotConnected := &Peer{
		ID:      "peer2",
		Addr:    "localhost:9999",     // This port is not used, simulating a peer that is not connected
		Message: make(chan []byte, 1), // Buffered channel to avoid blocking
	}

	// Attempt to send a message to the mock peer that is not connected
	err = tn.Send(mockPeerNotConnected, message)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "peer not connected")

	mockPeerNilConn := &Peer{
		ID:      "peer3",
		Addr:    mockPeerAddr,
		Message: make(chan []byte, 1), // Buffered channel to avoid blocking
	}

	// Attempt to send a message to the mock peer with a full channel
	mockPeerFullChannel := &Peer{
		ID:      "peer4",
		Addr:    mockPeerAddr,
		Message: make(chan []byte), // Unbuffered channel to ensure it's full
	}

	err = tn.Send(mockPeerNilConn, message)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "peer not connected or connection is nil")
	// Attempt to send a message to the mock peer with a full channel
	err = tn.Send(mockPeerFullChannel, message)
	assert.Error(t, err)
	assert.Errorf(t, err, "error sending message to peer %s: message channel full", mockPeerFullChannel.ID)
}
