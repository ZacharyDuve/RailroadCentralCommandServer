package turnout

import "github.com/ZacharyDuve/RailroadCentralCommandServer/src/storage"

type TurnoutManager struct {
	idGen          TurnoutIDGenerator
	turnoutStorage storage.Storage[TurnoutID, *Turnout]
}

func NewTurnoutManager(idGen TurnoutIDGenerator, tStore storage.Storage[TurnoutID, *Turnout]) (*TurnoutManager, error)
