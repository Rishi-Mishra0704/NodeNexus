package p2p

import (
	"fmt"
)

// Send sends a message to a peer in the network.
func (tn *TCPNetwork) Send(peer *Peer, message []byte) error {
	// Check if the peer is connected
	conn, ok := tn.peers[peer.ID]
	if !ok {
		// If not connected, return an error
		return fmt.Errorf("peer not connected: %s", peer.ID)
	}

	// Ensure the connection is not nil before writing to it
	if conn == nil {
		return fmt.Errorf("nil connection for peer: %s", peer.ID)
	}

	// Write the message to the peer connection
	_, err := conn.Write(message)
	return err
}

// Receive receives a message from a peer in the network.
func (tn *TCPNetwork) Receive() ([]byte, *Peer, error) {
	for _, conn := range tn.peers {
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			return nil, nil, err
		}
		// For simplicity, assume all messages come from the same peer for now

		// Get the peer associated with the connection
		var peer *Peer
		for id, c := range tn.peers {
			if c == conn {
				peer = &Peer{ID: id, Addr: conn.RemoteAddr().String()}
				break
			}
		}
		if peer == nil {
			return nil, nil, fmt.Errorf("failed to identify peer for connection")
		}

		return buffer[:n], peer, nil
	}
	return nil, nil, fmt.Errorf("no connected peers")
}
