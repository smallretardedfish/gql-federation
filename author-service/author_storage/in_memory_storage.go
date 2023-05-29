package author_storage

import (
	"errors"
	"github.com/smallretardedfish/gql-federation/author-service/graph/model"
	"golang.org/x/exp/maps"
	"strconv"
	"sync"
)

type InMemoryAuthorStorage struct {
	mu sync.Mutex
	db map[string]*model.Author
}

func NewInMemoryAuthorStorage() *InMemoryAuthorStorage {
	return &InMemoryAuthorStorage{db: make(map[string]*model.Author)}
}

func (s *InMemoryAuthorStorage) Create(b model.Author) (*model.Author, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	id := strconv.Itoa(len(s.db) + 1)
	b.ID = id
	s.db[id] = &b

	return &b, nil
}

func (s *InMemoryAuthorStorage) Get(id string) (*model.Author, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	res, ok := s.db[id]
	if !ok {
		return nil, errors.New("book not found")
	}

	return res, nil
}

func (s *InMemoryAuthorStorage) GetAll() ([]*model.Author, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return maps.Values(s.db), nil
}
