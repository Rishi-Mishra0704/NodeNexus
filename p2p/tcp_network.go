package p2p

import (
	"errors"
	"fmt"
	"net"
	"strconv"
)

// TCPNetwork represents a TCP-based peer-to-peer network.
type TCPNetwork struct {
	listener net.Listener
	peers    map[string]net.Conn
}

// Start starts the TCP network by listening for incoming connections.
func (tn *TCPNetwork) Start() error {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		return err
	}
	tn.listener = listener
	tn.peers = make(map[string]net.Conn)
	return nil
}

// Connect connects to a peer in the network.
func (tn *TCPNetwork) Connect(peer *Peer) error {
	conn, err := net.Dial("tcp", peer.IP+":"+strconv.Itoa(peer.Port))
	if err != nil {
		return err
	}
	tn.peers[peer.ID] = conn
	return nil
}

// Send sends a message to a peer in the network.
func (tn *TCPNetwork) Send(peer *Peer, message []byte) error {
	conn, ok := tn.peers[peer.ID]
	if !ok {
		return fmt.Errorf("peer not connected: %s", peer.ID)
	}
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
		peer := tn.findPeerByConn(conn)
		return buffer[:n], peer, nil
	}
	return nil, nil, errors.New("no connected peers")
}

// Close closes the TCP network.
func (tn *TCPNetwork) Close() error {
	for _, conn := range tn.peers {
		conn.Close()
	}
	return tn.listener.Close()
}

// Helper function to find peer by connection
func (tn *TCPNetwork) findPeerByConn(conn net.Conn) *Peer {
	for id, c := range tn.peers {
		if c == conn {
			return &Peer{ID: id}
		}
	}
	return nil
}
