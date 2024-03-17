package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPeer(t *testing.T) {
	// Test NewPeer function
	peer := NewPeer("peer1", "127.0.0.1", 8080)
	assert.NotNil(t, peer)
	assert.Equal(t, "peer1", peer.ID)
	assert.Equal(t, "127.0.0.1", peer.IP)
	assert.Equal(t, 8080, peer.Port)
}

func TestPeerString(t *testing.T) {
	// Test String method of Peer
	peer := &Peer{ID: "peer1", IP: "127.0.0.1", Port: 8080}
	expected := "Peer(ID: peer1, IP: 127.0.0.1, Port: 8080)"
	assert.Equal(t, expected, peer.String())
}
