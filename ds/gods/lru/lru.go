package lru

import "container/list"

type LRU struct {
	List *list.List
	Map  map[string]*list.Element
}

func NewLRU() LRU {
	return LRU{
		List: list.New(),
		Map:  make(map[string]*list.Element),
	}
}

func (lru *LRU) Add(key string, val interface{}) {
	ele := lru.List.PushBack(val)
	lru.Map[key] = ele
}

func (lru *LRU) Get(key string) interface{} {
	ele := lru.Map[key]
	if ele == nil {
		return nil
	}
	lru.List.MoveToBack(ele)

	return ele.Value
}

func (lru *LRU) Delete(key string) {
	ele := lru.Map[key]
	if ele == nil {
		return
	}
	lru.List.Remove(ele)
	delete(lru.Map, key)
}
