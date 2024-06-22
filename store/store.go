package store

import (
	"sync"
)

type KeyValueStore struct {
	data sync.Map
}

func NewKeyValueStore() *KeyValueStore {
	return &KeyValueStore{
		data: sync.Map{},
	}
}

func (s *KeyValueStore) Set(key string, value string) {
	s.data.Store(key, value)
}

func (s *KeyValueStore) Get(Key string) (string, bool) {
	value, ok := s.data.Load(Key)
	if !ok {
		return "", false
	}
	return value.(string), true
}

func (s *KeyValueStore) Delete(Key string) {
	s.data.Delete(Key)
}
