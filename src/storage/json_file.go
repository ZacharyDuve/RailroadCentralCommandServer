package storage

import (
	"cmp"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path"
)

const (
	ErrMsgBasePathMissing       = "error, basePath is required to be not empty"
	ErrMsgObjectTypeNameMissing = "error, objectTypeName is required to be not empty"
)

// FileManager is something that allows for creating, opening, removing files as well as directories
// Abstracting away so dependency of os can be injected

type FileManager interface {
	// Files
	Create(string) (*os.File, error)
	Open(string) (*os.File, error)
	Remove(string) error

	// Directories
	Mkdir(string, os.FileMode) error
	MkdirAll(string, os.FileMode) error
}

type OSFileManager struct {
}

type FileOpenFunc func(string) (io.ReadWriteCloser, error)

// JSONStorage stores and saves the objects as JSON files.
type JSONStorage[I cmp.Ordered, T Storable[I]] struct {
	filesPath    string
	fileOpenFunc FileOpenFunc
}

// NewJSONStorage implements the Storage interface and allows storing objects as json files
// basePath is the base directory where one intends to store files. Files will be stored in a sub directory
// objectTypeName is the name that one wants to give the object type. This should be unique for all types in the application.
// The sub directory in the basePath will be named this

func NewJSONStorage[I cmp.Ordered, T Storable[I]](basePath, objectTypeName string, fileOFunction FileOpenFunc) (Storage[I, T], error) {
	if basePath == "" {
		return nil, errors.New(ErrMsgBasePathMissing)
	}
	if objectTypeName == "" {
		return nil, errors.New(ErrMsgObjectTypeNameMissing)
	}

	filesPath := path.Join(basePath, objectTypeName)
	return &JSONStorage[I, T]{filesPath: filesPath}, nil
}

func (js *JSONStorage[I, T]) Save(obj T) error {

	filePath := fmt.Sprintf("%s%q%s", js.filesPath, os.PathSeparator, obj.ID())

	f, err := js.fileOpenFunc(filePath)

	if err == nil {
		err = json.NewEncoder(f).Encode(obj)
	}

	return err
}

func (js *JSONStorage[I, T]) Load(id Storable[I]) (T, error) {
	// Implement the logic to load the object with the specified ID from a JSON file in the specified directory
	// You can use the filesPath field of the JSONStorage struct to determine the directory where the file should be loaded from
	return *new(T), errors.ErrUnsupported
}

func (js *JSONStorage[I, T]) Delete(id Storable[I]) error {
	return errors.ErrUnsupported
}
