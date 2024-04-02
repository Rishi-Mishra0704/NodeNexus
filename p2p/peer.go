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

// ChatService defines the interface for the chat service.
// It has 4 methods: Join, Leave, SendMessage, and ReceiveMessages.
// Join allows a peer to join the chat service and returns a Peer or an error.
// Leave allows a peer to leave the chat service and returns an error.
// SendMessage sends a message to a peer and returns an error.
// ReceiveMessages returns a channel to receive messages or an error.
type ChatService interface {
	Join(username string) (Peer, error)
	Leave(peerID string) error
	SendMessage(message Message, sendChan chan<- Message) error // Added sendChan parameter
	ReceiveMessages() (<-chan Message, error)
}

// PeerDiscoveryService defines the interface for peer discovery.
// It has 3 methods:
// FindPeers: Retrieves a list of peers currently available in the network.
// RegisterPeer: Registers a peer with the discovery service, making its presence known to other peers.
// DeregisterPeer: Deregisters a peer from the discovery service, indicating that it is no longer available for communication.
type PeerDiscoveryService interface {
	FindPeers() ([]Peer, error)
	RegisterPeer(peer Peer) error
	DeregisterPeer(peerID string) error
}

// PeerCommunicationService defines the interface for peer-to-peer communication.
// It has 3 methods:
// Connect: Establishes a connection with a peer or returns an error.
// Disconnect: Closes the connection with a peer or returns an error.
// SendData: Sends data to a peer or returns an error.
type PeerCommunicationService interface {
	Connect(peer Peer) error
	Disconnect(peer Peer) error
	SendData(peer Peer, data []byte) error
}

// EncryptionService defines the interface for encryption and decryption.
// It has 2 methods: Encrypt and Decrypt.
// Encrypt encrypts the data and returns the encrypted data or an error.
// Decrypt decrypts the data and returns the decrypted data or an error.
type EncryptionService interface {
	Encrypt(data []byte) ([]byte, error)
	Decrypt(data []byte) ([]byte, error)
}
