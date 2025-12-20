package storage

import (
	"errors"

	"github.com/ZacharyDuve/RailroadCentralCommandServer/src/id"
)

// memory is an implementation of the Storage interface that just stores objects into memory
type memory[T id.IDable] struct {
	zeroValue *T
	objects   map[*id.ID]T
}

func NewMemoryStorage[T id.IDable]() Storage[T] {
	return &memory[T]{
		objects: make(map[*id.ID]T),
	}
}

func (m *memory[T]) Save(*T) error {
	return errors.ErrUnsupported
}

func (m *memory[T]) Load(id.ID) (*T, error) {
	return m.zeroValue, errors.ErrUnsupported
}
