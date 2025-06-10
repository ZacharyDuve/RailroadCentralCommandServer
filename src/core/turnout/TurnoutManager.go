package turnout

import (
	"github.com/ZacharyDuve/godatacollections"
)

var (
	unsetPositions       = []TurnoutPosition{PositionUnSet}
	leftHandedPositions  = []TurnoutPosition{PositionLeft, PositionThrough}
	rightHandedPositions = []TurnoutPosition{PositionThrough, PositionRight}
	wyePositions         = []TurnoutPosition{PositionLeft, PositionRight}
	threeWayPositions    = []TurnoutPosition{PositionLeft, PositionThrough, PositionRight}
	singleSlipPositions  = []TurnoutPosition{PositionTrackA, PositionTrackB, PositionDivergeTrackALeft}
	doubleSlipPositions  = []TurnoutPosition{PositionTrackA, PositionTrackB, PositionDivergeTrackALeft, PositionDivergeTrackBLeft}
	crossoverPositions   = []TurnoutPosition{PositionThrough, PositionCrossover}
)

// turnoutImpl is the backing implementation of the turnout type this type ties together with an turnout manager
// NOTE: This hidden to prevent creation exeternal of a turnout manager
type turnoutImpl struct {
	// tM is a reference to the turnout manager that manages it
	tM *turnoutManager

	id    TurnoutID
	name  string
	tType TurnoutType
}

func (this *turnoutImpl) ID() TurnoutID {
	return this.id
}

func (this *turnoutImpl) Name() string {
	return this.name
}

func (this *turnoutImpl) Type() TurnoutType {
	return this.tType
}

func (this *turnoutImpl) Positions() []TurnoutPosition {
	switch this.tType {
	case TypeLeftHanded:
		return leftHandedPositions
	case TypeRightHanded:
		return rightHandedPositions
	case TypeWye:
		return wyePositions
	case TypeThreeWay:
		return threeWayPositions
	case TypeSingleSlip:
		return singleSlipPositions
	case TypeDoubleSlip:
		return doubleSlipPositions
	case TypeCrossover:
		return crossoverPositions
	default:
		return unsetPositions
	}
}

type turnoutManager struct {
	idGen    TurnoutIDGenerator
	turnouts godatacollections.Set[TurnoutID, Turnout]
}

func NewTurnoutManager(idGen TurnoutIDGenerator, turnoutStorage godatacollections.Set[TurnoutID, Turnout]) *turnoutManager {

	return &turnoutManager{idGen: idGen, turnouts: turnoutStorage}
}

func (this *turnoutManager) NewTurnout(tType TurnoutType) (Turnout, error) {
	id, err := this.idGen.NewID()

	if err != nil {
		return nil, err
	}

	t := &turnoutImpl{tM: this, id: id, tType: tType}

	return t, nil

}
