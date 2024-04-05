package p2p

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type Client struct {
	Conn     net.Conn
	Name     string
	Room     *Room
	Commands chan<- Command
}

func (c *Client) ReadInput() {
	for {
		msg, err := bufio.NewReader(c.Conn).ReadString('\n')
		if err != nil {
			return
		}

		msg = strings.Trim(msg, "\r\n")

		args := strings.Split(msg, " ")
		cmd := strings.TrimSpace(args[0])

		switch cmd {
		case "/name":
			c.Commands <- Command{
				ID:     CMD_NAME,
				Client: c,
				Args:   args,
			}
		case "/join":
			c.Commands <- Command{
				ID:     CMD_JOIN,
				Client: c,
				Args:   args,
			}
		case "/rooms":
			c.Commands <- Command{
				ID:     CMD_ROOMS,
				Client: c,
			}
		case "/msg":
			c.Commands <- Command{
				ID:     CMD_MSG,
				Client: c,
				Args:   args,
			}
		case "/quit":
			c.Commands <- Command{
				ID:     CMD_QUIT,
				Client: c,
			}
		default:
			c.Err(fmt.Errorf("unknown command: %s", cmd))
		}
	}
}

func (c *Client) Err(err error) {
	c.Conn.Write([]byte("err: " + err.Error() + "\n"))
}

func (c *Client) Msg(msg string) {
	c.Conn.Write([]byte("> " + msg + "\n"))
}
