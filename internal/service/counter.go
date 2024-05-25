package service

import (
	"maps"
	"sync"
)

type Counter struct {
	sync.RWMutex
	data map[int]int
}

func NewCounter() *Counter {
	return &Counter{
		data: make(map[int]int),
	}
}

func (m *Counter) Add(key int) {
	m.Lock()
	defer m.Unlock()
	m.data[key]++
}

func (m *Counter) Get(key int) (int, bool) {
	m.RLock()
	defer m.RUnlock()
	value, ok := m.data[key]
	return value, ok
}

func (m *Counter) GetAll() map[int]int {
	m.RLock()
	defer m.RUnlock()
	return maps.Clone(m.data)
}
