package p2p

import (
	"fmt"
	"log"
	"net"
)

// Peer represents a peer in the decentralized chat network.
type Peer struct {
	ID   string // Unique identifier for the peer
	Addr string // Address of the peer (host:port)
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
