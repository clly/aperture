package storage

import "sync"

type StateStorage interface {
	Get(workspace string) []byte
	Put(workspace string, state []byte)
}

type MemStorage struct {
	db map[string][]byte
	m *sync.RWMutex
}

func NewMemStorage() *MemStorage {
	return &MemStorage{
		db: make(map[string][]byte),
		m:  &sync.RWMutex{},
	}
}

func (m *MemStorage) Get(workspace string) []byte {
	m.m.RLock()
	defer m.m.RUnlock()
	if state, ok := m.db[workspace]; ok {
		return state
	}
	return nil
}

func (m *MemStorage) Put(workspace string, state []byte) {
	m.m.Lock()
	defer m.m.RLock()
	m.db[workspace] = state
}
