package storage

import (
	"cmp"
	"encoding/json"
	"errors"
	"fmt"
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

func (o OSFileManager) Create(s string) (*os.File, error) {
	return os.Create(s)
}

func (o OSFileManager) Open(s string) (*os.File, error) {
	return os.Open(s)
}

func (o OSFileManager) Remove(s string) error {
	return os.Remove(s)
}

func (o OSFileManager) Mkdir(s string, p os.FileMode) error {
	return os.Mkdir(s, p)
}

func (o OSFileManager) MkdirAll(s string, p os.FileMode) error {
	return os.MkdirAll(s, p)
}

// JSONStorage stores and saves the objects as JSON files.
type JSONStorage[I cmp.Ordered, T Storable[I]] struct {
	filesPath   string
	fileManager FileManager
}

// NewJSONStorage implements the Storage interface and allows storing objects as json files
// basePath is the base directory where one intends to store files. Files will be stored in a sub directory
// objectTypeName is the name that one wants to give the object type. This should be unique for all types in the application.
// The sub directory in the basePath will be named this

func NewJSONStorage[I cmp.Ordered, T Storable[I]](basePath, objectTypeName string, fM FileManager) (Storage[I, T], error) {
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

func (js *JSONStorage[I, T]) Load(I) (T, error) {
	// Implement the logic to load the object with the specified ID from a JSON file in the specified directory
	// You can use the filesPath field of the JSONStorage struct to determine the directory where the file should be loaded from
	return *new(T), errors.ErrUnsupported
}

func (js *JSONStorage[I, T]) Delete(I) error {
	return errors.ErrUnsupported
}
