package turnout

type memTurnoutIDGen struct {
	next TurnoutID
}

func (this *memTurnoutIDGen) NewID() (TurnoutID, error) {
	id := this.next
	this.next = this.next + 1

	return id, nil
}

func NewMemTurnoutIDGen() TurnoutIDGenerator {
	return &memTurnoutIDGen{}
}
