package main

import (
	"encoding/json"
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
			log.Println("Error:", err)
			c.commands <- command{
				id:      CMD_SHUTDOWN,
				client:  c,
				args:    nil,
				options: nil,
			}
			return
		}
		var cmdObj commandObj
		err = json.Unmarshal(msg, &cmdObj)
		if err != nil {
			fmt.Println(err)
		}
		message := cmdObj.Command
		fmt.Println(string(message))
		message = strings.Trim(message, "\r\n")

		// naive parse argument
		args := strings.Split(message, " ")

		// command
		cmd := strings.TrimSpace(args[0])

		// get option flags
		// convert to map for O(1) access
		var options = make(map[string]bool)
		for _, op := range cmdObj.Options {
			options[op] = true
		}
		switch cmd {
		case "/nick":
			c.commands <- command{
				id:      CMD_NICK,
				client:  c,
				args:    args,
				options: options,
			}
		case "/join":
			c.commands <- command{
				id:      CMD_JOIN,
				client:  c,
				args:    args,
				options: options,
			}
		case "/rooms":
			c.commands <- command{
				id:      CMD_ROOMS,
				client:  c,
				args:    args,
				options: options,
			}
		case "/leave":
			c.commands <- command{
				id:      CMD_LEAVE,
				client:  c,
				args:    args,
				options: options,
			}
		case "/msg":
			c.commands <- command{
				id:      CMD_MSG,
				client:  c,
				args:    args,
				options: options,
			}
		case "/quit":
			c.commands <- command{
				id:      CMD_QUIT,
				client:  c,
				args:    args,
				options: options,
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
