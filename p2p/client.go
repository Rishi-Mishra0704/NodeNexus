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
	Commands chan<- Command
}

func NewClient(conn net.Conn, commands chan<- Command) *Client {
	return &Client{
		Conn:     conn,
		Commands: commands,
	}
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
		case "/msg":
			if len(args) >= 3 {
				recipientName := args[1]
				message := strings.Join(args[2:], " ")
				c.Commands <- Command{
					ID:            CMD_MSG,
					Client:        c,
					RecipientName: recipientName,
					Message:       message,
				}
			} else {
				c.Err(fmt.Errorf("usage: /msg <recipient_name> <message>"))
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
