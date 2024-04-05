package p2p

type CommandID int8

const (
	CMD_NAME CommandID = iota
	CMD_HELP
	CMD_JOIN
	CMD_ROOMS
	CMD_MSG
	CMD_QUIT
)

type Command struct {
	ID            CommandID
	Client        *Client
	Args          []string
	RecipientName string
	Message       string
}
