package p2p

import "fmt"

// Peer represents a peer in the peer-to-peer network.
type Peer struct {
	ID   string
	IP   string
	Port int
}

// NewPeer creates a new Peer with the given ID, IP, and port.
// It returns a pointer to the new Peer.
func NewPeer(id, ip string, port int) *Peer {
	return &Peer{
		ID:   id,
		IP:   ip,
		Port: port,
	}
}

// String returns a string representation of the Peer.
func (p *Peer) String() string {
	return fmt.Sprintf("Peer(ID: %s, IP: %s, Port: %d)", p.ID, p.IP, p.Port)
}
