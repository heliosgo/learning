package main

import (
	"sync"
)

type Player struct {
	name  string
	table *Table
	lobby *Lobby
}

var (
	players sync.Map
)

func newPlayer(name string, lobby *Lobby) *Player {
	return &Player{
		name:  name,
		lobby: lobby,
	}
}

func (p *Player) sitDown(id string) {
	lobbyMutex.Lock()
	defer lobbyMutex.Unlock()

	table, ok := p.lobby.findTable(id)
	if !ok {
		table = newTable(id)
		p.lobby.create <- createEvent{table: table}
	}

	table.addPlayer(p.name)
	p.table = table
}

func (p *Player) standUp() {
	if p.table == nil {
		return
	}
	lobbyMutex.Lock()
	defer lobbyMutex.Unlock()

	p.table.deletePlayer(p.name)
	if len(p.table.Player) == 0 {
		p.lobby.delete <- deleteEvent{id: p.table.ID}
	}
	p.table = nil
}

func (p *Player) drop(x, y int) (bool, bool) {
	table := p.table
	table.Mutex.Lock()
	defer table.Mutex.Unlock()

	if !table.IsStart || table.LastPlayer == p.name ||
		table.White != p.name && table.Black != p.name ||
		x < 0 || x >= sx || y < 0 || y >= sy ||
		table.CurState[x][y] != '+' {
		return false, false
	}

	var chess byte
	switch p.name {
	case table.White:
		chess = 'o'
	case table.Black:
		chess = '*'
	}

	table.CurState[x][y] = chess
	table.LastPlayer = p.name

	isWin := isWin(table.CurState, x, y)
	if isWin {
		table.IsStart = false
	}

	return isWin, true
}

var (
	dxy = [][][]int{
		{{-1, -1}, {1, 1}},
		{{0, -1}, {0, 1}},
		{{-1, 0}, {1, 0}},
		{{-1, 1}, {1, -1}},
	}
)

func isWin(state [sx][sy]byte, x, y int) bool {
	t := state[x][y]
	for _, v := range dxy {
		count := 1
		for _, d := range v {
			cx, cy := x, y
			for {
				tx, ty := cx+d[0], cy+d[1]
				if tx < 0 || tx >= sx || ty < 0 || ty >= sy {
					break
				}
				cx, cy = tx, ty
				if state[cx][cy] == t {
					count++
					if count >= 5 {
						return true
					}
				}
			}
		}
	}

	return false
}
