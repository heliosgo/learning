package main

import (
	"log"
	"net/http"
	"sync"

	_ "net/http/pprof"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/ws/{room}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["room"]
		globalRoomMutex.Lock()
		if _, ok := roomMutex[id]; !ok {
			roomMutex[id] = new(sync.Mutex)
		}
		roomMutex[id].Lock()
		globalRoomMutex.Unlock()
		room, ok := house.Load(id)
		var hub *Hub
		if ok {
			hub = room.(*Hub)
		} else {
			hub = newHub(id)
			house.Store(id, hub)
			go hub.run()
		}
		serveWs(hub, w, r)
	})
	go http.ListenAndServe(":6060", nil)

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalf("ListenAndServe: %v", err)
	}
}
