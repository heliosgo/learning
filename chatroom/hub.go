package main

import "sync"

var (
	house = sync.Map{}

	roomMutex       = make(map[string]*sync.Mutex)
	globalRoomMutex = sync.Mutex{}
)

type Hub struct {
	clients    map[*Client]struct{}
	broadcast  chan []byte
	unregister chan *Client
	roomID     string
}

func newHub(id string) *Hub {
	return &Hub{
		roomID:     id,
		clients:    make(map[*Client]struct{}),
		broadcast:  make(chan []byte),
		unregister: make(chan *Client),
	}
}

func (h *Hub) run() {
	defer func() {
		close(h.unregister)
		close(h.broadcast)
	}()
	for {
		select {
		case client := <-h.unregister:
			globalRoomMutex.Lock()
			roomMutex[h.roomID].Lock()
			globalRoomMutex.Unlock()
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
			if len(h.clients) == 0 {
				house.Delete(h.roomID)
				roomMutex[h.roomID].Unlock()
				return
			}
			roomMutex[h.roomID].Unlock()
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}
