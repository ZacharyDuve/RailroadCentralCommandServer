package tds

import (
	"errors"
	"fmt"
	"time"
)

type TDSEventType uint8

const (
	//We now know of of this TDSPositionID
	Connected TDSEventType = iota
	//We know that we no longer have control of it
	Disconnected
	//Position was requested but it isn't actually set yet
	Requested
	//The position just took as the current one for the driver
	Set
	//The position is no longer the set position of this driver
	Unset

	//Reserving the last one to be invalid so we can do a simple check for invalid types
	invalid
)

type TDSEvent struct {
	eventType           TDSEventType
	tdsDriverPositionID *TurnoutDriverPositionID
	fireTime            time.Time
}

func NewTDSEvent(eType TDSEventType, posID *TurnoutDriverPositionID) (tdse *TDSEvent, err error) {
	if eType >= invalid {
		return nil, fmt.Errorf("eType of %d is invalid, unable to create new TDSEvent", eType)
	} else if posID == nil {
		return nil, errors.New("unable to create new TDSEvent due to posID being nil")
	}

	return &TDSEvent{eventType: eType, tdsDriverPositionID: posID, fireTime: time.Now()}, nil
}

type TDSEventHandler interface {
	HandleTDSEvent(*TDSEvent)
}
