package p2p

import (
	"fmt"
)

// Send sends a message to a peer in the network.
func (tn *TCPNetwork) Send(peer *Peer, message []byte) error {
	// Check if the peer is connected
	conn, ok := tn.peers[peer.ID]
	if !ok || conn == nil {
		// If not connected or connection is nil, return an error
		return fmt.Errorf("peer not connected or connection is nil: %s", peer.ID)
	}

	// Send the message through the peer's message channel
	select {
	case peer.Message <- message:
		return nil
	default:
		// If the message channel is full, consider it as a write error
		delete(tn.peers, peer.ID)
		return fmt.Errorf("error sending message to peer %s: message channel full", peer.ID)
	}
}

// Receive receives a message from a peer in the network.
func (tn *TCPNetwork) Receive() ([]byte, *Peer, error) {
	for id, conn := range tn.peers {
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			return nil, nil, err
		}

		// Get the peer associated with the connection
		peer := &Peer{ID: id, Addr: conn.RemoteAddr().String()}

		return buffer[:n], peer, nil
	}
	return nil, nil, fmt.Errorf("no messages received")
}
