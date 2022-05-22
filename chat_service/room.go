package main

import (
	"net"
)

type room struct {
	name      string
	members   map[net.Addr]*client
	heartbeat int // a minute a heartbeat; maximum 60 minutes
}

func (r *room) broadcast(sender *client, msg string) {
	// send the message to all member in the room
	if r.heartbeat < 60 {
		r.heartbeat++
	}
	for addr, m := range r.members {
		if addr != sender.conn.RemoteAddr() {
			m.msg(msg)
		}
	}
}

func (r *room) tick() bool {
	if len(r.members) > 0 {
		return true
	}
	return (r.heartbeat > 0)
}
