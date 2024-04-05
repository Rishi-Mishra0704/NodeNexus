package p2p

import (
	"fmt"
	"log"
	"net"
	"strings"
)

type Server struct {
	Rooms    map[string]*Room
	Commands chan Command
}

func NewServer() *Server {
	return &Server{
		Rooms:    make(map[string]*Room),
		Commands: make(chan Command),
	}
}

func (s *Server) Run() {
	for cmd := range s.Commands {
		switch cmd.ID {
		case CMD_NAME:
			s.Name(cmd.Client, cmd.Args)
		case CMD_JOIN:
			s.Join(cmd.Client, cmd.Args)
		case CMD_ROOMS:
			s.ListRooms(cmd.Client)
		case CMD_MSG:
			s.Msg(cmd.Client, cmd.Args)
		case CMD_QUIT:
			s.Quit(cmd.Client)
		}
	}
}

func (s *Server) NewClient(conn net.Conn) *Client {
	log.Printf("new Client has Joined: %s", conn.RemoteAddr().String())

	return &Client{
		Conn:     conn,
		Name:     "anonymous",
		Commands: s.Commands,
	}
}

func (s *Server) Name(c *Client, Args []string) {
	if len(Args) < 2 {
		c.Msg("nick is required. usage: /nick Name")
		return
	}

	c.Name = Args[1]
	c.Msg(fmt.Sprintf("all right, I will call you %s", c.Name))
}

func (s *Server) Join(c *Client, Args []string) {
	if len(Args) < 2 {
		c.Msg("Room Name is required. usage: /Join Room_Name")
		return
	}

	RoomName := Args[1]

	r, ok := s.Rooms[RoomName]
	if !ok {
		r = &Room{
			Name:    RoomName,
			Members: make(map[net.Addr]*Client),
		}
		s.Rooms[RoomName] = r
	}
	r.Members[c.Conn.RemoteAddr()] = c

	s.QuitCurrentRoom(c)
	c.Room = r

	r.broadcast(c, fmt.Sprintf("%s Joined the Room", c.Name))

	c.Msg(fmt.Sprintf("welcome to %s", RoomName))
}

func (s *Server) ListRooms(c *Client) {
	var Rooms []string
	for Name := range s.Rooms {
		Rooms = append(Rooms, Name)
	}

	c.Msg(fmt.Sprintf("available Rooms: %s", strings.Join(Rooms, ", ")))
}

func (s *Server) Msg(c *Client, Args []string) {
	if len(Args) < 2 {
		c.Msg("message is required, usage: /Msg Msg")
		return
	}

	Msg := strings.Join(Args[1:], " ")
	c.Room.broadcast(c, c.Name+": "+Msg)
}

func (s *Server) Quit(c *Client) {
	log.Printf("Client has left the chat: %s", c.Conn.RemoteAddr().String())

	s.QuitCurrentRoom(c)

	c.Msg("sad to see you go =(")
	c.Conn.Close()
}

func (s *Server) QuitCurrentRoom(c *Client) {
	if c.Room != nil {
		oldRoom := s.Rooms[c.Room.Name]
		delete(s.Rooms[c.Room.Name].Members, c.Conn.RemoteAddr())
		oldRoom.broadcast(c, fmt.Sprintf("%s has left the Room", c.Name))
	}
}
