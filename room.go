package main

type room struct {
	name    string
	members map[string]*client
}

func (r *room) broadcast(sender *client, msg string) {
	for _, member := range r.members {
		if sender.conn.RemoteAddr() != member.conn.RemoteAddr() {
			member.message(msg)
		}
	}
}

func (r *room) broadcastServerMessage(msg string) {
	for _, member := range r.members {
		member.srvmessage(msg)
	}
}
