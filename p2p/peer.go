package p2p

import (
	"fmt"
	"log"
	"net"
)

// Peer represents a peer in the decentralized chat network.
type Peer struct {
	ID      string      // Unique identifier for the peer
	Addr    string      // Address of the peer (host:port)
	Message chan []byte // Channel for sending/receiving messages
}

// NewPeer creates a new Peer instance with the specified ID and address.
func NewPeer(id, addr string) *Peer {
	return &Peer{
		ID:      id,
		Addr:    addr,
		Message: make(chan []byte),
	}
}

// Connect establishes a connection to the specified peer.
func (p *Peer) Connect() (net.Conn, error) {
	conn, err := net.Dial("tcp", p.Addr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to peer %s: %w", p.Addr, err)
	}
	log.Printf("Connected to peer %s\n", p.Addr)
	return conn, nil
}
