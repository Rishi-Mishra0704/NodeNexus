package p2p

import (
	"net"
)

type Server struct {
	Peers     map[string]*Client // Map of peer addresses to connections
	Commands  chan Command
	Bootstrap *net.TCPListener // Bootstrap server for peer discovery
}

func NewServer() *Server {
	return &Server{
		Peers:    make(map[string]*Client),
		Commands: make(chan Command),
	}
}

func (s *Server) HandleClient(conn net.Conn) {
	client := NewClient(conn, s.Commands)
	s.Peers[conn.RemoteAddr().String()] = client
	go client.ReadInput()
}
