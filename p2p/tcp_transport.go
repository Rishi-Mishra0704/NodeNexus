package p2p

import (
	"fmt"
	"log"
	"net"
)

// TCPPeer represents the remote node over a TCP established connection
type TCPPeer struct {

	// conn is the underlying connection of the peer
	conn net.Conn

	// if we dial and retrieve a conn => outbound == true
	// if we accept and retrieve a conn => outbound == false
	outbound bool
}

type TcpTransportOpts struct {
	ListenAddr    string
	HandShakeFunc func(net.Conn) Peer
	OnPeer        func(Peer) error
}

type TcpTransport struct {
	TcpTransportOpts
	listener net.Listener
	MsgCh    chan Message
}

func NewTcpPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn:     conn,
		outbound: outbound,
	}
}

func (p *TCPPeer) Close() error {
	return p.conn.Close()
}

func NewTCPTransport(opts TcpTransportOpts) *TcpTransport {
	return &TcpTransport{
		TcpTransportOpts: opts,
		MsgCh:            make(chan Message),
	}
}

// Consume implements the transport interface, which will return read-only channel
// for reading the message recieved from another peer in the network
func (t *TcpTransport) Consume() <-chan Message {
	return t.MsgCh
}

func (t *TcpTransport) ListenAndAccept() error {

	var err error
	t.listener, err = net.Listen("tcp", t.ListenAddr)
	if err != nil {
		log.Fatal(err)
	}

	go t.startAcceptLoop()

	return nil
}

func (t *TcpTransport) startAcceptLoop() {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			fmt.Printf("TCP accept error %s\n", err)
		}
		fmt.Printf("New incoming connection %v\n", conn)

		go t.handleConn(conn)

	}
}
func (t *TcpTransport) handleConn(conn net.Conn) {
	var err error
	defer func() {
		fmt.Printf("Dropping peer connection%s", err)
		conn.Close()
	}()

	// Create a new peer
	peer := NewTcpPeer(conn, true)

	// Perform handshake and log the result
	handShakePeer := t.HandShakeFunc(peer.conn)
	if handShakePeer != nil {
		fmt.Printf("Handshake successful for peer: %v\n", handShakePeer)
	} else {
		fmt.Println("Handshake failed")
	}

	// Call OnPeer callback if provided
	if t.OnPeer != nil {
		if err := t.OnPeer(peer); err != nil {
			return
		}
	}

	// Read loop
}
