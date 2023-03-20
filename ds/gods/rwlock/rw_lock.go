package rwlock

import "sync"

type RWLock struct {
	Mutex           *sync.Mutex
	ReaderCond      *sync.Cond
	WriterCond      *sync.Cond
	ReaderWaitCount int
	WriterWaitCount int
	Flag            int
}

func New() RWLock {
	mutex := &sync.Mutex{}
	return RWLock{
		Mutex:           mutex,
		ReaderCond:      sync.NewCond(mutex),
		WriterCond:      sync.NewCond(mutex),
		ReaderWaitCount: 0,
		WriterWaitCount: 0,
		Flag:            0,
	}
}

func (rw *RWLock) RLock() {
	rw.Mutex.Lock()
	rw.ReaderWaitCount++
	for rw.Flag == -1 {
		rw.ReaderCond.Wait()
	}
	rw.Flag++
	rw.ReaderWaitCount--
	rw.Mutex.Unlock()
}

func (rw *RWLock) RUnlock() {
	rw.Mutex.Lock()
	if rw.Flag == -1 || rw.Flag == 0 {
		rw.Mutex.Unlock()
		return
	}
	rw.Flag--
	if rw.Flag == 0 && rw.WriterWaitCount > 0 {
		rw.WriterCond.Signal()
	}
	rw.Mutex.Unlock()
}

func (rw *RWLock) WLock() {
	rw.Mutex.Lock()
	rw.WriterWaitCount++
	for rw.Flag != 0 {
		rw.WriterCond.Wait()
	}
	rw.Flag = -1
	rw.WriterWaitCount--
	rw.Mutex.Unlock()
}

func (rw *RWLock) WUnlock() {
	rw.Mutex.Lock()
	if rw.Flag != -1 {
		rw.Mutex.Unlock()
		return
	}
	rw.Flag = 0
	if rw.WriterWaitCount > 0 {
		rw.WriterCond.Signal()
		rw.Mutex.Unlock()
		return
	}
	if rw.ReaderWaitCount > 0 {
		rw.ReaderCond.Broadcast()
		rw.Mutex.Unlock()
		return
	}

	rw.Mutex.Unlock()
}
