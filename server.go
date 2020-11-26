package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

type server struct {
	rooms    map[string]*room
	commands chan command
}

func NewServer() *server {
	return &server{
		rooms:    make(map[string]*room),
		commands: make(chan command),
	}
}

func (s *server) run() {
	// listen on the commands channel
	for cmd := range s.commands {
		switch cmd.id {
		case CMD_NAME:
			s.rename(cmd)
		case CMD_MSG:
			s.message(cmd)
		case CMD_JOIN:
			s.joinRoom(cmd)
		case CMD_LEAVE:
			s.leaveRoom(cmd)
		case CMD_ROOMS:
			s.listRooms(cmd)
		case CMD_ROOM:
			s.displayRoomInfos(cmd)
		case CMD_QUIT:
			s.quit(cmd)
		case CMD_HELP:
			s.displayHelp(cmd)
		}
	}
}

func (s *server) newClient(conn net.Conn) {
	log.Printf("new client has joined: %s\n", conn.RemoteAddr().String())
	c := &client{
		conn:     conn,
		name:     "anon",
		commands: s.commands,
	}

	c.readInput()
}

// rename client and let him know
func (s *server) rename(cmd command) {
	name := strings.Join(cmd.args[1:], " ")
	cmd.client.name = name
	cmd.client.message(fmt.Sprintf("SERVER: Ok now I will call you %v", name))
}

func (s *server) message(cmd command) {
	msg := strings.Join(cmd.args, " ")
	if ok := cmd.client.room; ok != nil && len(cmd.client.room.members) > 1 {
		cmd.client.room.broadcast(cmd.client, cmd.client.name+": "+msg)
	} else {
		cmd.client.message("SERVER: Nobody hears you")
	}
}

func (s *server) joinRoom(cmd command) {
	name := strings.Join(cmd.args[1:], " ")

	//check if room exist
	r, ok := s.rooms[name]
	if !ok {
		//if not we create a room
		r = &room{
			name:    name,
			members: make(map[net.Addr]*client),
		}

		s.rooms[name] = r
	}

	// add client to room members
	r.members[cmd.client.conn.RemoteAddr()] = cmd.client

	//leave previous room if any
	s.leaveRoom(cmd)

	// set new room on client
	cmd.client.room = r

	// broadcast to room
	r.broadcast(cmd.client, fmt.Sprintf("SERVER: %v has joined the room.", cmd.client.name))

	// greet client from joining room
	cmd.client.message(fmt.Sprintf("SERVER: Welcome to \"%v\".", r.name))
}

func (s *server) listRooms(cmd command) {
	var names []string
	for name := range s.rooms {
		names = append(names, name)
	}

	cmd.client.message(fmt.Sprintf("SERVER: Here are the available rooms: %s", strings.Join(names, ", ")))
}

func (s *server) displayRoomInfos(cmd command) {
	var names []string
	for _, client := range cmd.client.room.members {
		names = append(names, client.name)
	}

	cmd.client.message(fmt.Sprintf("SERVER: List of users connected to the room \"%s\": %s", cmd.client.room.name, strings.Join(names, ", ")))
}

func (s *server) leaveRoom(cmd command) {
	if cmd.client.room != nil {
		oldRoom := s.rooms[cmd.client.room.name]
		delete(s.rooms[cmd.client.room.name].members, cmd.client.conn.RemoteAddr())
		cmd.client.room = nil
		oldRoom.broadcast(cmd.client, fmt.Sprintf("SERVER: %v has left the room.", cmd.client.name))
	}
}

func (s *server) displayHelp(cmd command) {
	cmd.client.message("TODO")
}

func (s *server) quit(cmd command) {
	defer log.Printf("a client as left the chat: %s", cmd.client.conn.RemoteAddr().String())
	s.leaveRoom(cmd)
	cmd.client.message("Have a good one, see you soon !")
	cmd.client.conn.Close()
}
