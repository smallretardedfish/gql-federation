package book_storage

import (
	"errors"
	"github.com/smallretardedfish/gql-federation/book-service/graph/model"
	"golang.org/x/exp/maps"
	"strconv"
	"sync"
)

type InMemoryBookStorage struct {
	mu sync.Mutex
	db map[string]*model.Book
}

func NewInMemoryBookStorage() *InMemoryBookStorage {
	return &InMemoryBookStorage{db: make(map[string]*model.Book)}
}

func (s *InMemoryBookStorage) Create(b model.Book) (*model.Book, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	id := strconv.Itoa(len(s.db) + 1)
	b.ID = id
	s.db[id] = &b

	return &b, nil
}

func (s *InMemoryBookStorage) Get(id string) (*model.Book, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	res, ok := s.db[id]
	if !ok {
		return nil, errors.New("book not found")
	}

	return res, nil
}

func (s *InMemoryBookStorage) GetAll() ([]*model.Book, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return maps.Values(s.db), nil
}
