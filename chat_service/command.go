package main

type commandID int

const (
	CMD_NICK commandID = iota
	CMD_JOIN
	CMD_ROOMS
	CMD_MSG
	CMD_LEAVE
	CMD_QUIT
	CMD_SHUTDOWN
)

type command struct {
	id      commandID
	client  *client
	args    []string        // all arguments. used for /nick and /msg
	options map[string]bool // filtered flags
}
