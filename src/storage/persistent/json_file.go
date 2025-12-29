package persistent

import (
	"cmp"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"

	"github.com/ZacharyDuve/RailroadCentralCommandServer/src/storage"
)

const (
	ErrMsgBasePathMissing       = "error, basePath is required to be not empty"
	ErrMsgObjectTypeNameMissing = "error, objectTypeName is required to be not empty"

	// Needs to be octal or stuff get wild
	defaultFilePermissions = 0740
)

// FileManager is something that allows for creating, opening, removing files as well as directories
// Abstracting away so dependency of os can be injected

type FileManager interface {
	// Files
	OpenFile(path string, flags int, perm os.FileMode) (*os.File, error)
	Remove(path string) error

	// Directories
	Mkdir(path string, perm os.FileMode) error
	MkdirAll(path string, perm os.FileMode) error
}

type OSFileManager struct {
}

func (o OSFileManager) OpenFile(path string, flags int, perm os.FileMode) (*os.File, error) {
	return os.OpenFile(path, flags, perm)
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
type JSONStorage[I cmp.Ordered, T storage.Storable[I]] struct {
	filesPath   string
	fileManager FileManager
}

// NewJSONStorage implements the Storage interface and allows storing objects as json files
// basePath is the base directory where one intends to store files. Files will be stored in a sub directory
// objectTypeName is the name that one wants to give the object type. This should be unique for all types in the application.
// The sub directory in the basePath will be named this

func NewJSONStorage[I cmp.Ordered, T storage.Storable[I]](path string, fM FileManager) (storage.Storage[I, T], error) {
	if path == "" {
		return nil, errors.New(ErrMsgBasePathMissing)
	}

	// Need to ensure that the directories required are built out
	if err := fM.MkdirAll(path, defaultFilePermissions); err != nil {
		return nil, err
	}

	return &JSONStorage[I, T]{filesPath: path, fileManager: fM}, nil
}

func (js *JSONStorage[I, T]) Save(obj T) error {

	filePath := path.Join(js.filesPath, fmt.Sprintf("%v.json", obj.ID()))

	f, err := js.fileManager.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, defaultFilePermissions)

	if err == nil {
		err = json.NewEncoder(f).Encode(obj)
	}

	return err
}

func (js *JSONStorage[I, T]) Load(id I) (T, error) {
	var t T

	filePath := path.Join(js.filesPath, fmt.Sprintf("%v.json", id))

	f, err := js.fileManager.OpenFile(filePath, os.O_RDONLY, defaultFilePermissions)

	if err == nil {
		err = json.NewDecoder(f).Decode(&t)
	}

	return t, err
}

func (js *JSONStorage[I, T]) Delete(id I) error {
	filePath := path.Join(js.filesPath, fmt.Sprintf("%v.json", id))

	return js.fileManager.Remove(filePath)
}
