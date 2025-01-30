package turnouttype

const (
	POSITION_NAME_STRAIGHT_THROUGH string = "Straight Through"
	POSITION_NAME_DIVERGE_LEFT     string = "Diverge Left"
	POSITION_NAME_DIVERGE_RIGHT    string = "Diverge Right"
)

// Adds all of the builtin turnout types to the specified library
func addBuiltInTypes(ttl TurnoutTypeLibrary) {
	panicOnError(ttl.AddTurnoutType(builtinLeftHandTurnoutType()))
	panicOnError(ttl.AddTurnoutType(builtinRightHandTurnoutType()))
	panicOnError(ttl.AddTurnoutType(builtinWyeTurnoutType()))
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func builtinLeftHandTurnoutType() TurnoutType {
	return &turnoutTypeImpl{name: "Left Handed Turnout", positionNames: []string{POSITION_NAME_DIVERGE_LEFT, POSITION_NAME_STRAIGHT_THROUGH}}
}

func builtinRightHandTurnoutType() TurnoutType {
	return &turnoutTypeImpl{name: "Right Handed Turnout", positionNames: []string{POSITION_NAME_STRAIGHT_THROUGH, POSITION_NAME_DIVERGE_RIGHT}}
}

func builtinWyeTurnoutType() TurnoutType {
	return &turnoutTypeImpl{name: "Wye Turnout", positionNames: []string{POSITION_NAME_DIVERGE_LEFT, POSITION_NAME_DIVERGE_RIGHT}}
}
