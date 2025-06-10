package turnout

// TurnoutID is the Id of a specific turnout on the railroad
type TurnoutID uint32

// TurnoutIDGenerator is something that is capable of generating unique turnout ids
type TurnoutIDGenerator interface {
	// NewID generates a new unique TurnoutID could return an error
	NewID() (TurnoutID, error)
}

type TurnoutType uint8

const (
	TypeUnSet TurnoutType = iota
	TypeLeftHanded
	TypeRightHanded
	TypeWye
	TypeThreeWay
	TypeSingleSlip
	TypeDoubleSlip
	TypeCrossover
)

// func CompTurnoutType(a, b TurnoutType) int {
// 	if a < b {
// 		return -1
// 	} else if a > b {
// 		return 1
// 	}
// 	return 0
// }

type TurnoutPosition uint8

const (
	PositionUnSet TurnoutPosition = iota
	PositionThrough
	PositionLeft
	PositionRight
	PositionTrackA
	PositionTrackB
	PositionDivergeTrackALeft
	PositionDivergeTrackBLeft
	PositionCrossover
)

// Turnout is a device that allows for a collection of switch machines to be set together
// Also allows for easier human interaction with physical turnout
type Turnout interface {
	// ID is the id of this turnout
	ID() TurnoutID
	// Name is the human readable name of it
	Name() string
	// Type is the type of turnout that this is
	Type() TurnoutType
	// Positions allow for one to get a list of all available positions that the turnout has
	Positions() []TurnoutPosition
}

func CompTurnoutID(a, b TurnoutID) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}

func TurnoutIDFromTurnout(t Turnout) TurnoutID {
	if t == nil {
		panic("turnout cannot be nil")
	}
	return t.ID()
}

// TurnoutPositionState is an enum of the possible states that a position could be in
type TurnoutPositionState uint8

const (
	// UnSet is a position that the turnout is not currently set to
	UnSet TurnoutPositionState = iota
	// Setting is a position that the turnout out is trying to Set but hasn't become Set yet
	Setting
	// Set is where the turnout is currently routing through.
	// Only one position could be set at a given time
	Set
)

// // TurnoutPosition is a container attributes about a specific position
// type TurnoutPosition interface {
// 	// Name of the position
// 	// All positions on a single turnout should have unique names
// 	Name() string
// 	// CurrentState is the current state of the given position
// 	CurrentState() TurnoutPositionState
// 	// LinkedTDSPositionIDs are the ids of the device positions that need to be set for this Turnout to be set to this position
// 	LinkedTDSPositionIDs() []TDSDevicePositionID
// }
