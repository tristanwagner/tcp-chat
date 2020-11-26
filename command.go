package main

type commandId int

const (
	CMD_NAME commandId = iota
	CMD_MSG
	CMD_JOIN
	CMD_LEAVE
	CMD_ROOMS
	CMD_ROOM
	CMD_QUIT
	CMD_HELP
)

type command struct {
	id     commandId
	client *client
	args   []string
}
