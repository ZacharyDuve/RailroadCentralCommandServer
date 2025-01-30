package turnouttype

import (
	"errors"

	"github.com/hashicorp/go-set/v3"
)

const (
	TURNOUT_TYPE_LIBRARY_ERROR_ALREADY_CONTAINS string = "error, unable to add turnout type to library because one of the same name already is contained"
)

var libary TurnoutTypeLibrary

func init() {
	libary = &turnoutTypeLibaryImpl{tTypes: set.NewTreeSet[TurnoutType](compareTurnoutTypes)}
	addBuiltInTypes(libary)
}

type TurnoutTypeLibrary interface {
	AddTurnoutType(TurnoutType) error
}

func DefaultTurnoutTypeLibrary() TurnoutTypeLibrary {
	return libary
}

type turnoutTypeLibaryImpl struct {
	tTypes *set.TreeSet[TurnoutType]
}

func (this *turnoutTypeLibaryImpl) AddTurnoutType(tt TurnoutType) error {
	if !this.tTypes.Insert(tt) {
		return errors.New(TURNOUT_TYPE_LIBRARY_ERROR_ALREADY_CONTAINS)
	}
	return nil

}
