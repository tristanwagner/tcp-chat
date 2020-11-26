package main

import (
	"bufio"
	// "fmt"
	"net"
	"strings"
)

type client struct {
	conn     net.Conn
	name     string
	room     *room
	commands chan<- command
}

func (c *client) message(msg string) {
	c.conn.Write([]byte("> " + msg + "\n"))
}

func (c *client) err(err error) {
	c.conn.Write([]byte("ERROR: " + err.Error() + "\n"))
}

func (c *client) readInput() {
	for {
		msg, err := bufio.NewReader(c.conn).ReadString('\n')
		if err != nil {
			return
		}

		msg = strings.Trim(msg, "\r\n")
		args := strings.Split(msg, " ")
		cmd := args[0]

		switch cmd {
		case "/name":
			c.commands <- command{
				id:     CMD_NAME,
				client: c,
				args:   args,
			}
		case "/join":
			c.commands <- command{
				id:     CMD_JOIN,
				client: c,
				args:   args,
			}
		case "/leave":
			c.commands <- command{
				id:     CMD_LEAVE,
				client: c,
				args:   args,
			}
		case "/rooms":
			c.commands <- command{
				id:     CMD_ROOMS,
				client: c,
				args:   args,
			}
		case "/room":
			c.commands <- command{
				id:     CMD_ROOM,
				client: c,
				args:   args,
			}
		case "/quit":
			c.commands <- command{
				id:     CMD_QUIT,
				client: c,
				args:   args,
			}
		case "/help":
			c.commands <- command{
				id:     CMD_HELP,
				client: c,
				args:   args,
			}
		default:
			c.commands <- command{
				id:     CMD_MSG,
				client: c,
				args:   args,
			}
			// c.message(fmt.Sprintf("SERVER: unknown command %s", cmd))
		}
	}
}
