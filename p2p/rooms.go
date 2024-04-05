package p2p

import "net"

type Room struct {
	Name    string
	Members map[net.Addr]*Client
}

func (r *Room) broadcast(sender *Client, msg string) {
	for addr, m := range r.Members {
		if sender.Conn.RemoteAddr() != addr {
			m.Msg(msg)
		}
	}
}
