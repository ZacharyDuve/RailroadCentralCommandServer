package tds

import (
	"fmt"

	"github.com/google/uuid"
)

type TurnoutDriverID uint16
type TurnoutDriverPosition uint16

type TurnoutDriverPositionID struct {
	tdsUUID        uuid.UUID
	driverID       TurnoutDriverID
	driverPosition TurnoutDriverPosition
}

func NewTurnoutDriverPositionID(tdsUUID uuid.UUID, dID TurnoutDriverID, dPos TurnoutDriverPosition) TurnoutDriverPositionID {
	return TurnoutDriverPositionID{tdsUUID: tdsUUID, driverID: dID, driverPosition: dPos}
}

func (this *TurnoutDriverPositionID) TDSUUID() uuid.UUID {
	return this.tdsUUID
}

func (this *TurnoutDriverPositionID) DriverID() TurnoutDriverID {
	return this.driverID
}

func (this *TurnoutDriverPositionID) Position() TurnoutDriverPosition {
	return this.driverPosition
}

func (this *TurnoutDriverPositionID) String() string {
	return fmt.Sprintf("%s.%d.%d", this.tdsUUID.String(), this.driverID, this.driverPosition)
}

type TurnoutDriver interface {
	ID() TurnoutDriverID
	Connected() bool
	NumberPositions() uint16
	CurrentPosition() TurnoutDriverPosition
	RequestPosition() error
}
