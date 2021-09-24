package storage

import "sync"

type Memory struct {
	mu sync.RWMutex
	data map[string][]byte
}

// NewMemory returns a new memory storage.
func NewMemory() *Memory {
	return &Memory{
		mu:   sync.RWMutex{},
		data: map[string][]byte{},
	}
}

// Get value by a key.
func (s *Memory) Get(key string) ([]byte, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	val, ok := s.data[key]
	return val, ok
}

func (s *Memory) Set(key string, val []byte) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[key] = val
}

func (s *Memory) Delete(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.data, key)
}

func (s *Memory) Count() int {
	return len(s.data)
}