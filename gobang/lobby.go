package main

import "sync"

type Lobby struct {
	tables sync.Map
	create chan createEvent
	delete chan deleteEvent
}

type createEvent struct {
	table *Table
	f     func()
}

type deleteEvent struct {
	id string
	f  func()
}

var (
	lobbyMutex sync.Mutex

	lobbyOnce sync.Once
)

func newLobby() *Lobby {
	return &Lobby{
		create: make(chan createEvent, 100),
		delete: make(chan deleteEvent, 100),
		tables: sync.Map{},
	}

}

func (lobby *Lobby) run() {
	for {
		select {
		case event := <-lobby.create:
			lobby.tables.Store(event.table.ID, event.table)
		case event := <-lobby.delete:
			lobby.tables.Delete(event.id)
		}
	}
}

func (lobby *Lobby) findTable(id string) (*Table, bool) {
	table, ok := lobby.tables.Load(id)
	if !ok {
		return nil, false
	}

	return table.(*Table), true
}
