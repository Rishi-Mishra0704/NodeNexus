package p2p

type Protocol string

// Define constants for TCP and UDP protocols
const (
	TCP Protocol = "tcp"
	UDP Protocol = "udp"
)

// Transport defines the interface for peer-to-peer communication transport.
type Transport interface {
	// Connect establishes a connection with a peer using the specified protocol (TCP/UDP).
	Connect(peer Peer, protocol Protocol) error

	// Disconnect closes the connection with a peer.
	Disconnect(peer Peer) error

	// SendData sends data to a peer.
	SendData(peer Peer, data []byte) error

	// ReceiveData receives data from a peer.
	ReceiveData() ([]byte, Peer, error)
}
