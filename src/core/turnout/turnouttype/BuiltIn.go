package turnouttype

import "github.com/hashicorp/go-set/v3"

const (
	TURNOUT_TYPE_NAME_LEFT_HANDED  string = "Left Handed"
	TURNOUT_TYPE_NAME_RIGHT_HANDED string = "Right Handed"
	TURNOUT_TYPE_NAME_WYE          string = "Wye"
	POSITION_NAME_STRAIGHT_THROUGH string = "Straight Through"
	POSITION_NAME_DIVERGE_LEFT     string = "Diverge Left"
	POSITION_NAME_DIVERGE_RIGHT    string = "Diverge Right"
)

// Adds all of the builtin turnout types to the specified library
func addBuiltInTypes(ttl *set.TreeSet[TurnoutType]) {
	panicOnFailedInsertBuiltInTurnoutType(ttl.Insert(builtinLeftHandTurnoutType()))
	panicOnFailedInsertBuiltInTurnoutType(ttl.Insert(builtinRightHandTurnoutType()))
	panicOnFailedInsertBuiltInTurnoutType(ttl.Insert(builtinWyeTurnoutType()))
}

func panicOnFailedInsertBuiltInTurnoutType(inserted bool) {
	if !inserted {
		panic("failed to insert builtin turnout type")
	}
}

func builtinLeftHandTurnoutType() TurnoutType {
	return &turnoutTypeImpl{name: TURNOUT_TYPE_NAME_LEFT_HANDED, positionNames: []string{POSITION_NAME_DIVERGE_LEFT, POSITION_NAME_STRAIGHT_THROUGH}}
}

func builtinRightHandTurnoutType() TurnoutType {
	return &turnoutTypeImpl{name: TURNOUT_TYPE_NAME_RIGHT_HANDED, positionNames: []string{POSITION_NAME_STRAIGHT_THROUGH, POSITION_NAME_DIVERGE_RIGHT}}
}

func builtinWyeTurnoutType() TurnoutType {
	return &turnoutTypeImpl{name: TURNOUT_TYPE_NAME_WYE, positionNames: []string{POSITION_NAME_DIVERGE_LEFT, POSITION_NAME_DIVERGE_RIGHT}}
}
