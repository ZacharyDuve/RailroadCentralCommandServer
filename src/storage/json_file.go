package storage

import (
	"errors"
	"path"

	"github.com/ZacharyDuve/RailroadCentralCommandServer/src/id"
)

const (
	ErrMsgBasePathMissing       = "error, basePath is required to be not empty"
	ErrMsgObjectTypeNameMissing = "error, objectTypeName is required to be not empty"
)

// JSONStorage stores and saves the objects as JSON files.
type JSONStorage[T id.IDable] struct {
	filesPath string
}

// NewJSONStorage implements the Storage interface and allows storing objects as json files
// basePath is the base directory where one intends to store files. Files will be stored in a sub directory
// objectTypeName is the name that one wants to give the object type. This should be unique for all types in the application.
// The sub directory in the basePath will be named this

func NewJSONStorage[T id.IDable](basePath, objectTypeName string) (Storage[T], error) {
	if basePath == "" {
		return nil, errors.New(ErrMsgBasePathMissing)
	}
	if objectTypeName == "" {
		return nil, errors.New(ErrMsgObjectTypeNameMissing)
	}

	filesPath := path.Join(basePath, objectTypeName)
	return &JSONStorage[T]{filesPath: filesPath}, nil
}

func (js *JSONStorage[T]) Save(obj *T) error {
	// 1) Check to ensure that
	return nil
}

func (js *JSONStorage[T]) Load(id id.ID) (*T, error) {
	// Implement the logic to load the object with the specified ID from a JSON file in the specified directory
	// You can use the filesPath field of the JSONStorage struct to determine the directory where the file should be loaded from
	return nil, nil
}
