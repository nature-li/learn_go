package common

import "sync"

type SynchronizedMap struct {
	rw *sync.RWMutex
	data map[interface{}]interface{}
}

func (sm *SynchronizedMap) Put (k, v interface{}) {
	sm.rw.Lock()
	defer sm.rw.Unlock()

	sm.data[k] = v
}

func (sm *SynchronizedMap) Get(k interface{}) interface{} {
	sm.rw.RLock()
	defer sm.rw.RUnlock()

	return sm.data[k]
}

func (sm *SynchronizedMap) Delete(k interface{}) {
	sm.rw.Lock()
	defer sm.rw.Unlock()

	delete(sm.data, k)
}

func (sm *SynchronizedMap) Each (cb func(interface{}, interface{})) {
	sm.rw.RLock()
	defer sm.rw.RUnlock()

	for k, v := range sm.data {
		cb(k, v)
	}
}

func NewSynchronizedMap() *SynchronizedMap {
	return &SynchronizedMap{
		data : make(map[interface{}]interface{}),
	}
}