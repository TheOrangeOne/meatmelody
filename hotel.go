package main

import (
	"crypto/rand"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

const URL_LEN = 16

type Hotel struct {
	rooms map[string]*Room
}

func newHotel() *Hotel {
	return &Hotel{
		rooms: make(map[string]*Room),
	}
}

// check a guest in to a room
func (h *Hotel) checkinGuest(id string, room *Room, conn *websocket.Conn) {
	guest := &Guest{
		id:   id,
		room: room,
		conn: conn,
		send: make(chan []byte, 256),
	}
	guest.room.register <- guest

	go guest.writeSocket()
	go guest.readSocket()
}

func (h *Hotel) prepareRoom(roomName string) *Room {
	log.Printf("prepared new room %s", roomName)
	room := newRoom(roomName)
	h.rooms[roomName] = room
	go h.reservation(room)
	return room
}

func GenRandBytes(n int) ([]byte, error) {
	bytes := make([]byte, n)

	_, err := rand.Read(bytes)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func GenRandString(n int) (string, error) {
	const alpha = "023456789abcdefghjkmnopqrstvwxyz"
	bytes, err := GenRandBytes(n)

	if err != nil {
		return "", err
	}

	for i, b := range bytes {
		bytes[i] = alpha[b%byte(len(alpha))]
	}

	return string(bytes), nil
}

func (h *Hotel) createRoom() (*Room, error) {
	roomName, err := GenRandString(URL_LEN)
	if err != nil {
		return nil, err
	}
	room := h.prepareRoom(roomName)
	return room, nil
}

func (h *Hotel) reservation(room *Room) {
	room.reserve()
	delete(h.rooms, room.name)
	log.Printf("unreserved room %s", room.name)
}

func (h *Hotel) serveHotel(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		log.Printf("id not given dropping connection")
		conn.Close()
		return
	}

	roomName := r.URL.Path
	room, ok := h.rooms[roomName]
	if !ok {
		room = h.prepareRoom(roomName)
	}

	h.checkinGuest(id, room, conn)
}
