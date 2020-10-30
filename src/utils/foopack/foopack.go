package foopack

import "sync"

type SafeCounter struct {
	V   map[string]int
	mux sync.Mutex
}

func (s *SafeCounter) Inc(key string) {
	s.mux.Lock()
	s.V[key]++
	s.mux.Unlock()
}

func (s *SafeCounter) Get(key string) int {
	s.mux.Lock()
	defer s.mux.Unlock()
	return s.V[key]
}
