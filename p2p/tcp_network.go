package p2p

import (
	"fmt"
	"log"
	"net"
)

// TCPNetwork represents a TCP-based peer-to-peer network.
type TCPNetwork struct {
	Listener net.Listener
	peers    map[string]net.Conn
}

// NewTCPNetwork creates a new TCPNetwork instance.
func NewTCPNetwork() *TCPNetwork {
	return &TCPNetwork{
		peers: make(map[string]net.Conn),
	}
}

// Start starts the TCP network by listening for incoming connections.
func (tn *TCPNetwork) Start(port string) error {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}
	tn.Listener = listener
	log.Printf("TCP network started on port %s\n", port)
	return nil
}

// Connect connects to a peer in the network.
func (tn *TCPNetwork) Connect(peer *Peer) error {
	conn, err := net.Dial("tcp", peer.Addr)
	if err != nil {
		return fmt.Errorf("failed to connect to peer %s: %w", peer.Addr, err)
	}
	tn.peers[peer.ID] = conn
	log.Printf("Connected to peer %s\n", peer.Addr)
	return nil
}

// Close closes the TCP network.
func (tn *TCPNetwork) Close() error {
	if tn.Listener != nil {
		// Close all peer connections
		for _, conn := range tn.peers {
			conn.Close()
		}

		// Close the listener and reset it to nil
		err := tn.Listener.Close()
		tn.Listener = nil // Reset the listener
		return err
	}
	return nil
}
