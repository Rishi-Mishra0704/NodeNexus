package p2p

import "time"

// Peer represents a peer in the network.
type Peer struct {
	ID       string
	Username string
	Address  string
	Port     int
}

// Message represents a chat message exchanged between peers.
type Message struct {
	From      Peer
	To        Peer
	Content   string
	Timestamp time.Time
}
