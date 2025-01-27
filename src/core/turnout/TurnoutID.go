package turnout

type TurnoutIDGenerator interface {
	NewID() TurnoutID
}
