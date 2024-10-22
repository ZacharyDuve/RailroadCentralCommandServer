package railroad

import (
	"errors"

	"github.com/google/uuid"
)

type fileRailroadStorer struct {
	rootPath string
}

func NewFileRailroadStorer(rootPath string) (RailroadStorer, error) {
	if rootPath == "" {
		return nil, errors.New("invalid rootPath for new file railroad storer of blank")
	}

	fRRS := &fileRailroadStorer{}

	return fRRS, nil
}

func (this *fileRailroadStorer) Add(*railroad) error {
	return nil
}

func (this *fileRailroadStorer) Update(*railroad) error {
	return nil
}

func (this *fileRailroadStorer) Delete(id uuid.UUID) error {
	return nil
}

func (this *fileRailroadStorer) FindById(id uuid.UUID) (*railroad, error) {
	return nil, nil
}

func (this *fileRailroadStorer) FindByName(name string) (*railroad, error) {
	return nil, nil
}
