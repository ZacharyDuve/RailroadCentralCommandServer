package turnout

type TurnoutID uint64

type TurnoutType uint8

const (
	Unknown TurnoutType = iota
	TwoWay
)

type Turnout struct {
	id          TurnoutID
	turnoutType TurnoutType
}

func (t *Turnout) ID() TurnoutID {
	return t.id
}
