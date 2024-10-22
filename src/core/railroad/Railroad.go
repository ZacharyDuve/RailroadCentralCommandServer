package railroad

import (
	"errors"

	"github.com/ZacharyDuve/RailroadCentralCommandServer/src/persistance"
)

type RailroadStorer interface {
	persistance.Storer[*railroad]
	FindByName(name string) (*railroad, error)
}

type railroad struct {
	name        string
	description string
	persistance.StorableTimeStamp
}

func (this *railroad) TypeName() string {
	return "Railroad"
}

func (this *railroad) PrimaryKey() persistance.PrimaryKey {
	return persistance.PrimaryKey(this.name)
}

func NewRailroad(name, description string) (*railroad, error) {
	if name == "" {
		return nil, errors.New("invalid leaving name for new railroad blank")
	}

	rr := &railroad{name: name, description: description}
	rr.StorableTimeStamp.Init()

	return rr, nil
}
