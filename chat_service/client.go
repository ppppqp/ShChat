package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/gorilla/websocket"
)

type client struct {
	conn     *websocket.Conn
	nick     string
	room     *room
	commands chan<- command
}

func (c *client) readInput() {
	for {
		// read in a message
		_, msg, err := c.conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		// print out that message for clarity
		message := string(msg)
		fmt.Println(string(message))
		// if err := c.conn.WriteMessage(messageType, p); err != nil {
		// 	log.Println(err)
		// 	return
		// }

		message = strings.Trim(message, "\r\n")
		// parse argument
		args := strings.Split(message, " ")
		cmd := strings.TrimSpace(args[0])
		switch cmd {
		case "/nick":
			c.commands <- command{
				id:     CMD_NICK,
				client: c,
				args:   args,
			}
		case "/join":
			c.commands <- command{
				id:     CMD_JOIN,
				client: c,
				args:   args,
			}
		case "/rooms":
			c.commands <- command{
				id:     CMD_ROOMS,
				client: c,
				args:   args,
			}
		case "/msg":
			c.commands <- command{
				id:     CMD_MSG,
				client: c,
				args:   args,
			}
		case "/quit":
			c.commands <- command{
				id:     CMD_QUIT,
				client: c,
				args:   args,
			}
		default:
			c.err(fmt.Errorf("unknown cmd"))
		}
	}
}

func (c *client) err(err error) {
	c.conn.WriteMessage(1, []byte(err.Error()))
}

func (c *client) msg(msg string) {
	if err := c.conn.WriteMessage(1, []byte(msg)); err != nil {
		log.Println(err)
		return
	}
}
