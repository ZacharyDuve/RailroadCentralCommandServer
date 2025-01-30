package turnout

import "github.com/ZacharyDuve/RailroadCentralCommandServer/src/core/turnout/turnouttype"

// TurnoutID is the Id of a specific turnout on the railroad
type TurnoutID uint32

// TurnoutIDGenerator is something that is capable of generating unique turnout ids
type TurnoutIDGenerator interface {
	// NewID generates a new unique TurnoutID could return an error
	NewID() (TurnoutID, error)
}

// Turnout is a device that allows for a collection of switch machines to be set together
// Also allows for easier human interaction with physical turnout
type Turnout interface {
	// ID is the id of this turnout
	ID() TurnoutID
	// Name is the human readable name of it
	Name() string
	// Type is the type of turnout that this is
	Type() turnouttype.TurnoutType
	// Positions allow for one to get a list of all available positions that the turnout has
	Positions() []TurnoutPosition
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

// TurnoutPosition is a container attributes about a specific position
type TurnoutPosition interface {
	// Name of the position
	// All positions on a single turnout should have unique names
	Name() string
	// State is the current state of the given position
	CurrentState() TurnoutPositionState
}
