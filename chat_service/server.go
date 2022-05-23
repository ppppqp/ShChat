package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

type server struct {
	rooms    map[string]*room
	commands chan command
	ticker   *time.Ticker
}

func newServer() *server {
	s := &server{
		rooms:    make(map[string]*room),
		commands: make(chan command),
		ticker:   time.NewTicker(60 * time.Second),
	}
	go func() {
		for {
			select {
			case <-s.ticker.C:
				// on every tick
				for k, r := range s.rooms {
					// update the heartbeat of each room
					if !r.tick() {
						// if the room is nolonger considered active
						// close the room
						delete(s.rooms, k)
					}
				}
			}
		}
	}()
	return s
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (s *server) newClient(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Client Connected")
	err = ws.WriteMessage(1, []byte("Greetings from the chat-room-tender! You are now connected. Join or create a room."))
	if err != nil {
		log.Println(err)
		return
	}
	c := client{
		conn:     ws,
		nick:     "anonymous",
		commands: s.commands,
	}
	c.readInput()
}

func (s *server) setupRoutes() {
	http.HandleFunc("/connect", s.newClient)
}

func (s *server) run() {
	for cmd := range s.commands {
		switch cmd.id {
		case CMD_NICK:
			s.nick(cmd.client, cmd.args, cmd.options)
		case CMD_JOIN:
			s.join(cmd.client, cmd.args, cmd.options)
		case CMD_ROOMS:
			s.listRooms(cmd.client, cmd.args, cmd.options)
		case CMD_MSG:
			s.msg(cmd.client, cmd.args, cmd.options)
		case CMD_QUIT:
			s.quit(cmd.client, cmd.args, cmd.options)
		}
	}
}

func (s *server) nick(c *client, args []string, options map[string]bool) {
	c.nick = args[1]
	c.msg(fmt.Sprintf("all right, I will call you %s", c.nick))
}

func (s *server) join(c *client, args []string, options map[string]bool) {
	roomName := args[1]
	r, ok := s.rooms[roomName]
	if !ok {
		// room doesn't exists
		r = &room{
			name:    roomName,
			members: make(map[net.Addr]*client),
		}
		s.rooms[roomName] = r
	}
	r.members[c.conn.RemoteAddr()] = c
	r.heartbeat = 60
	s.quitCurrentRoom(c)
	c.room = r
	r.broadcast(c, fmt.Sprintf("%s has joined the room", c.nick))
	c.msg(fmt.Sprintf("welcome to %s", r.name))
}

func (s *server) listRooms(c *client, args []string, options map[string]bool) {
	var rooms []string
	for name := range s.rooms {
		rooms = append(rooms, name)
	}
	if len(rooms) == 0 {
		c.msg(fmt.Sprintf("There is no available room currently. Create one!"))
	} else {
		if _, ok := options["-v"]; ok {
			str := ""
			str += fmt.Sprintf("+%s+\n", strings.Repeat("-", 65))
			str += fmt.Sprintf("| %-20s | %-20s/60 | %-20s \n", "Room Name", "Active Level", "#Members")
			str += fmt.Sprintf("+%s+\n", strings.Repeat("-", 65))
			for k, r := range s.rooms {
				str += fmt.Sprintf("- %-20s   %-20d   %-20d \n", k, r.heartbeat, len(r.members))
			}
			c.msg(str)
		} else {
			c.msg(fmt.Sprintf("available rooms are: %s", strings.Join(rooms, ",")))
		}

	}
}

func (s *server) msg(c *client, args []string, options map[string]bool) {
	if c.room == nil {
		c.err(errors.New("you must join the room first"))
		return
	}
	if len(args) < 2 {
		c.msg("message is required, usage: /msg MSG")
		return
	}
	msg := strings.Join(args[1:], " ")
	c.room.broadcast(c, c.nick+": "+msg)
}

func (s *server) quit(c *client, args []string, options map[string]bool) {
	log.Printf("client has disconnected: %s", c.conn.RemoteAddr().String())
	delete(c.room.members, c.conn.RemoteAddr())
	c.msg("Bye!")
	c.conn.Close()
}

func (s *server) quitCurrentRoom(c *client) {
	if c.room != nil {
		delete(c.room.members, c.conn.RemoteAddr())
		c.room.broadcast(c, fmt.Sprintf("%s has left the room", c.nick))
	}
}
