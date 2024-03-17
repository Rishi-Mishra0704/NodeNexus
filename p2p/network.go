package p2p

// Network represents the interface for interacting with the peer-to-peer network.
type Network interface {
	Start() error
	Connect(peer *Peer) error
	Send(peer *Peer, message []byte) error
	Receive() ([]byte, *Peer, error)
	Close() error
}
