package p2p

import (
	"errors"
	"net"
)

/*
ErrInvalidHandshakeError is returned when a connection
between local and remote nodes is not established
*/
var ErrInvalidHandshakeError = errors.New("invalid handshake")

// HandShakeFunc is a function that is called when a new connection is established
type HandShakeFunc func(Peer) error

func DefaultHandShakeFunc(conn net.Conn) Peer {
	// For now,  a simple handshake where the peer's address is used as the ID.
	peer := NewTcpPeer(conn, true)
	return peer
}

func NOPHandShakeFunc(Peer) error { return nil }
