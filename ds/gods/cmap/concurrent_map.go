package cmap

import (
	"errors"
	"sync"
	"time"
)

type onceChan struct {
	sync.Once
	ch chan struct{}
}

func newOnceChan() *onceChan {
	return &onceChan{
		ch: make(chan struct{}),
	}
}

func (oc *onceChan) close() {
	oc.Once.Do(func() {
		close(oc.ch)
	})
}

type ConcurrentMap struct {
	sync.Mutex
	mp        map[int]int
	keyToChan map[int]*onceChan
}

func NewConcurrentMap() *ConcurrentMap {
	return &ConcurrentMap{
		mp:        make(map[int]int),
		keyToChan: make(map[int]*onceChan),
	}
}

func (cmap *ConcurrentMap) Put(key, val int) {
	cmap.Lock()
	defer cmap.Unlock()

	cmap.mp[key] = val
	ch, ok := cmap.keyToChan[key]
	if !ok {
		return
	}

	ch.close()
}

func (cmap *ConcurrentMap) Get(key int, timeout time.Duration) (int, error) {
	cmap.Lock()
	if val, ok := cmap.mp[key]; !ok {
		cmap.Unlock()
		return val, nil
	}

	ch, ok := cmap.keyToChan[key]
	if !ok {
		ch = newOnceChan()
		cmap.keyToChan[key] = ch
	}
	cmap.Unlock()

	timer := time.NewTicker(timeout)
	select {
	case <-timer.C:
		return 0, errors.New("timeout")
	case <-ch.ch:
	}

	cmap.Lock()
	defer cmap.Unlock()

	return cmap.mp[key], nil
}
