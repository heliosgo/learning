package main

import (
	"strings"
	"sync"
	"time"
)

const (
	sx = 10
	sy = 10
)

var (
	startState [sx][sy]byte = initStartState()
)

func initStartState() [sx][sy]byte {
	var res [sx][sy]byte
	for i := range res {
		for j := range res[i] {
			res[i][j] = '+'
		}
	}

	return res
}

type Table struct {
	ID         string
	Player     map[string]int64
	LastPlayer string
	Time       time.Duration
	X          int
	Y          int
	CurState   [sx][sy]byte
	White      string
	Black      string
	IsStart    bool

	Mutex sync.Mutex
}

func newTable(id string) *Table {
	return &Table{
		ID:       id,
		Player:   make(map[string]int64),
		Time:     30 * time.Second,
		X:        sx,
		Y:        sy,
		CurState: startState,
	}
}

func (t *Table) addPlayer(user string) {
	t.Mutex.Lock()
	t.Player[user] = time.Now().Unix()
	switch {
	case t.Black == "":
		t.Black = user
	case t.White == "":
		t.White = user
	}
	if t.Black != "" && t.White != "" {
		t.IsStart = true
	}
	t.Mutex.Unlock()
}

func (t *Table) deletePlayer(user string) {
	t.Mutex.Lock()
	delete(t.Player, user)
	switch {
	case t.White == user:
		t.White = ""
	case t.Black == user:
		t.Black = ""
	}
	if t.LastPlayer == user {
		t.LastPlayer = ""
	}
	if t.White == "" || t.Black == "" {
		t.IsStart = false
	}
	t.Mutex.Unlock()
}

func (t *Table) getStringState() string {
	t.Mutex.Lock()
	var res strings.Builder
	for _, line := range t.CurState {
		res.Write(line[:])
		res.WriteByte('\n')
	}

	t.Mutex.Unlock()

	return res.String()
}
