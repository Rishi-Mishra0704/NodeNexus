package main

import (
	"fmt"
	"log"

	"github.com/Rishi-Mishra0704/NodeNexus/p2p"
)

func main() {
	// Create three peers
	peer1 := &p2p.Peer{ID: "peer1", Addr: "localhost:8001"}
	peer2 := &p2p.Peer{ID: "peer2", Addr: "localhost:8002"}
	peer3 := &p2p.Peer{ID: "peer3", Addr: "localhost:8003"}

	// Create a network instance
	network := p2p.NewTCPNetwork()

	// Start the network
	err := network.Start("8001")
	if err != nil {
		log.Fatalf("Error starting network: %v", err)
	}
	err1 := network.Start("8002")
	if err1 != nil {
		log.Fatalf("Error starting network: %v", err)
	}
	err2 := network.Start("8003")
	if err2 != nil {
		log.Fatalf("Error starting network: %v", err)
	}

	// Connect peers
	connectPeers(network, peer1, peer2)
	connectPeers(network, peer1, peer3)

	// Send messages
	sendMessage(network, peer1, peer2, []byte("Hello from peer1 to peer2"))
	sendMessage(network, peer1, peer3, []byte("Hello from peer1 to peer3"))

	// Receive messages
	receiveMessage(network, peer2)
	receiveMessage(network, peer3)

	// Close network
	err = network.Close()
	if err != nil {
		log.Fatalf("Error closing network: %v", err)
	}
}

func connectPeers(network *p2p.TCPNetwork, sender, receiver *p2p.Peer) {
	err := network.Connect(receiver)
	if err != nil {
		log.Fatalf("Error connecting peer %s to peer %s: %v", sender.ID, receiver.ID, err)
	}
	fmt.Printf("Connected %s to %s\n", sender.ID, receiver.ID)
}

func sendMessage(network *p2p.TCPNetwork, sender, receiver *p2p.Peer, message []byte) {
	err := network.Send(receiver, message)
	if err != nil {
		log.Fatalf("Error sending message from %s to %s: %v", sender.ID, receiver.ID, err)
	}
	fmt.Printf("Message sent from %s to %s\n", sender.ID, receiver.ID)
}

func receiveMessage(network *p2p.TCPNetwork, peer *p2p.Peer) {
	message, sender, err := network.Receive()
	if err != nil {
		log.Fatalf("Error receiving message for peer %s: %v", peer.ID, err)
	}
	fmt.Printf("Received message for %s from %s: %s\n", peer.ID, sender.ID, string(message))
}
