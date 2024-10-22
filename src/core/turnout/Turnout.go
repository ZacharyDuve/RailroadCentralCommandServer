package turnout

import "github.com/google/uuid"

type TurnoutPosition uint16

type Turnout struct {
	dbID         uint
	uuid         uuid.UUID
	name         string
	numPositions uint16
}

//Format TDCUUID_TDID_TDP

//type TurnoutDriverPosition uint16

func CreateTurnoutDriverPositionID(driverUUID uuid.UUID, driverID uint16)
