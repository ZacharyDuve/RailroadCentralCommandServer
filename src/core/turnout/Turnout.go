package turnout

import (
	"errors"
	"strings"
)

type TurnoutID []byte
type TurnoutTypeName string

const (
	TURNOUT_TYPE_UNSET TurnoutTypeName = "UNSET Turnout Type"
)

var turnoutIDSequence uint8

type Turnout struct {
	id    TurnoutID
	name  string
	notes string
	//Need some way to associate a turnout to positions
	turnoutTypeName TurnoutTypeName
	positions       []Position
}

func createNewTurnoutID() []byte {
	newId := []byte{turnoutIDSequence}
	turnoutIDSequence++

	return newId
}

func createBlankTurnout() *Turnout {
	return &Turnout{id: createNewTurnoutID()}
}

func (this *Turnout) ID() []byte {
	return this.id
}

func (this *Turnout) Name() string {
	return this.name
}

func (this *Turnout) UpdateName(newName string) error {
	if len(strings.TrimSpace(newName)) == 0 {
		return errors.New("Unable to update name to a blank name")
	}

	return nil
}

func (this *Turnout) UpdateNotes(newNotes string) {
	this.notes = newNotes
}

/*
Position is just a position on a turnout, ex: "Diverging", "Through"

There can be duplicates of positions through out the application. Such as if you have multiple Ys then you can have multiple "Left"s
*/
type Position string

/*
TurnoutPosition is a combination of the turnout Id along with the position.

These are supposed to be globally unique due to the nature of being related to the turnout's id which is supposed to be globally unique
*/
type TurnoutPosition struct {
	turnoutID TurnoutID
	position  Position
}
