package main

import (
	"log"
)

type Room struct {
	name string

	guests map[*Guest]bool

	broadcast chan []byte

	register chan *Guest

	unregister chan *Guest
}

func newRoom(name string) *Room {
	return &Room{
		name:       name,
		broadcast:  make(chan []byte),
		register:   make(chan *Guest),
		unregister: make(chan *Guest),
		guests:     make(map[*Guest]bool),
	}
}

// a reserved room allows guests to enter and exit freely
func (r *Room) reserve() {
	for {
		select {
		case guest := <-r.register:
			r.guests[guest] = true
			go func() {
				// TODO on connection message
				// r.broadcast <-
			}()
			log.Printf("%s entered room %s", guest.id, r.name)
		case guest := <-r.unregister:
			if _, ok := r.guests[guest]; ok {
				delete(r.guests, guest)
				close(guest.send)
				go func() {
					// TODO on disconnect message
					// r.broadcast
				}()
				log.Printf("%s left room %s", guest.id, r.name)
			}

			// finish if there are no more guests
			if len(r.guests) == 0 {
				// TODO: timeout mechanism?
				return
			}
		case message := <-r.broadcast:
			for guest := range r.guests {
				select {
				case guest.send <- message:
				default:
					close(guest.send)
					delete(r.guests, guest)
				}
			}
		}
	}
}
