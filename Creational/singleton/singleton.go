package singleton

import "sync"

type Singleton interface {
	AddOne() int
}

func GetInstance() *singleton {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		instance = new(singleton)
	}
	return instance
}

var lock = &sync.Mutex{}
var instance *singleton

type singleton struct {
	mu    sync.RWMutex
	count int
}

func (s *singleton) GetCount() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.count
}
func (s *singleton) AddValue(val int) int {
	s.mu.Lock()
	s.count += val
	s.mu.Unlock()
	return s.count
}
