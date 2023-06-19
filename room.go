package main

import "sync"

type room struct {
	name    string
	members map[string]*client
	mu      sync.Mutex
}

func (r *room) broadcast(sender *client, msg string) {
	r.mu.Lock()
	for _, member := range r.members {
		if sender.conn.RemoteAddr() != member.conn.RemoteAddr() {
			member.message(msg)
		}
	}
	r.mu.Unlock()
}

func (r *room) broadcastServerMessage(msg string) {
	r.mu.Lock()
	for _, member := range r.members {
		member.srvmessage(msg)
	}
	r.mu.Unlock()
}
