package main

import (
	"log"

	"github.com/Rishi-Mishra0704/NodeNexus/p2p"
)

func main() {
	// Define the TCP transport options
	tcpOpts := p2p.TcpTransportOpts{
		ListenAddr:    ":8080", // Listening on port 8080
		HandShakeFunc: p2p.DefaultHandShakeFunc,
		Decoder:       p2p.DefaultDecoder{},
	}

	// Create a new TCP transport instance
	tcpTransport := p2p.NewTCPTransport(tcpOpts)

	// Start listening for incoming connections
	if err := tcpTransport.ListenAndAccept(); err != nil {
		log.Fatalf("Failed to start TCP transport: %v", err)
	}

	// Log a message indicating that the TCP transport is listening
	log.Println("TCP transport started and listening for incoming connections on port 8080")
	// Create a goroutine to consume messages from the transport
	go func() {
		for {
			select {
			case msg := <-tcpTransport.Consume():
				log.Printf("Received message: %s\n", msg.Payload)
			}
		}
	}()
	// Send a dummy message to the message channel
	dummyMessage := p2p.Message{
		From:    nil,                               // Set the sender address to nil or specify the actual address
		To:      nil,                               // Set the recipient address to nil or specify the actual address
		Payload: []byte("This is a dummy message"), // Set the payload to any desired value
	}
	log.Printf("Sending dummy message: %v", dummyMessage)
	tcpTransport.MsgCh <- dummyMessage

	// Keep the main goroutine alive
	select {}
}
