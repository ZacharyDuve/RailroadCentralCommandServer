package turnout

import (
	"errors"

	"github.com/ZacharyDuve/RailroadCentralCommandServer/src/core/turnout/turnouttype"
)

type turnoutManager struct {
	idGen TurnoutIDGenerator
	//turnoutLoadStorer persistance.LoadStorer[*Turnout]
	turnouts []Turnout
}

// turnoutImpl is the backing implementation of the turnout type this type ties together with an turnout manager
// NOTE: This hidden to prevent creation exeternal of a turnout manager
type turnoutImpl struct {
	// tM is a reference to the turnout manager that manages it
	tM *turnoutManager

	id    TurnoutID
	name  string
	tType turnouttype.TurnoutType

	positions []TurnoutPosition
}

func (this *turnoutImpl) ID() TurnoutID {
	return this.id
}

func (this *turnoutImpl) Name() string {
	return this.name
}

func (this *turnoutImpl) Type() turnouttype.TurnoutType {
	return this.tType
}

func (this *turnoutImpl) Positions() []TurnoutPosition {
	panic(errors.ErrUnsupported)
}

type turnoutPositionImpl struct {
	// reference back to the turnout that this position is associated with
	t *turnoutImpl

	name  string
	state TurnoutPositionState
}

func (this *turnoutPositionImpl) Name() string {
	return this.name
}

func (this *turnoutPositionImpl) CurrentState() TurnoutPositionState {
	return this.state
}

func NewTurnoutManager(idGen TurnoutIDGenerator, ttLib turnouttype.TurnoutTypeLibrary) *turnoutManager {
	return &turnoutManager{idGen: idGen, turnouts: make([]Turnout, 0)}
}

func (this *turnoutManager) NewTurnout(tType turnouttype.TurnoutType) (Turnout, error) {
	id, err := this.idGen.NewID()

	if err != nil {
		return nil, err
	}

	t := &turnoutImpl{tM: this, id: id, tType: tType, positions: make([]TurnoutPosition, 0)}

	for _, curPositionName := range tType.PositionNames() {
		t.positions = append(t.positions, &turnoutPositionImpl{t: t, name: curPositionName, state: UnSet})
	}

	return t, nil

}
