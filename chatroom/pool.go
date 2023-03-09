package main

import (
	"sync"

	"github.com/panjf2000/ants/v2"
)

const (
	maxClient = 100000
)

var pool *ants.Pool

var poolMutex sync.Mutex

func initGoroutinePool() (func(), error) {
	var err error
	pool, err = ants.NewPool(maxClient)
	if err != nil {
		return nil, err
	}

	return func() {
		pool.Release()
		ants.Release()
	}, nil
}

func submitToPool(task func()) {
	pool.Submit(task)
}

func getAvailableWorker() int {
	return pool.Free()
}
