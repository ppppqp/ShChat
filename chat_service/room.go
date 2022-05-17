package main

import "net"

type room struct {
	name    string
	members map[net.Addr]*client
}

func (r *room) broadcast(sender *client, msg string) {
	// send the message to all member in the room
	for addr, m := range r.members {
		if addr != sender.conn.RemoteAddr() {
			m.msg(msg)
		}
	}
}
