package p2p

import (
	"errors"
	"fmt"
	"net"
)

var _ Transport = (*TCPTransport)(nil)

type TCPTransport struct {
	// ListenAddress and ListenPort can be used for listening for incoming connections
	ListenAddress string
	ListenPort    int
}

func (t *TCPTransport) Connect(peer Peer, protocol Protocol) error {
	// Ensure protocol is TCP
	if protocol != TCP {
		return errors.New("invalid protocol for TCP transport")
	}

	// Validate peer's address and port
	if peer.Address == "" {
		return errors.New("peer address cannot be empty")
	}
	if peer.Port <= 0 || peer.Port > 65535 {
		return errors.New("invalid peer port")
	}

	// Construct peer address
	address := fmt.Sprintf("%s:%d", peer.Address, peer.Port)

	// Establish TCP connection
	conn, err := net.Dial(string(protocol), address)
	if err != nil {
		return fmt.Errorf("failed to connect to peer %s: %w", peer.ID, err)
	}
	defer conn.Close()

	// Connection successful, handle further logic if needed

	return nil
}

// Disconnect closes the connection with a peer.
func (t *TCPTransport) Disconnect(peer Peer) error {
	return nil
}

// SendData sends data to a peer over the established connection.
func (t *TCPTransport) SendData(peer Peer, data []byte) error {
	return nil
}

// ReceiveData receives data from a peer over the established connection.
func (t *TCPTransport) ReceiveData() ([]byte, Peer, error) {
	return nil, Peer{}, nil
}
