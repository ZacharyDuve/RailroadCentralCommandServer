package turnout

import "github.com/ZacharyDuve/RailroadCentralCommandServer/src/persistance"

type turnoutManager struct {
	idGen             TurnoutIDGenerator
	turnoutLoadStorer persistance.LoadStorer[*Turnout]
}

func NewTurnoutManager() *turnoutManager {
	return &turnoutManager{}
}
