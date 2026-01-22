package turnout

type TurnoutIDGenerator interface {
	Next() (TurnoutID, error)
}
