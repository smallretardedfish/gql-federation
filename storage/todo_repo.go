package storage

import (
	"github.com/dgryski/trifles/uuid"
	"golang.org/x/exp/maps"
	"sync"
)

type TodoStore interface {
	Save(todo Todo) (*Todo, error)
	GetAll() ([]*Todo, error)
}

type InMemoryStore struct {
	lock sync.Mutex
	m    map[string]*Todo
}

func (s *InMemoryStore) Save(todo Todo) (*Todo, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	todo.ID = uuid.UUIDv4()
	s.m[todo.ID] = &todo

	return &todo, nil
}

func (s *InMemoryStore) GetAll() ([]*Todo, error) {
	return maps.Values(s.m), nil
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{m: make(map[string]*Todo)}
}
