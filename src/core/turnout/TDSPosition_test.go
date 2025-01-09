package turnout

import (
	"testing"

	"github.com/google/uuid"
)

func TestNewPositionWithoutServerIDReturnsError(t *testing.T) {
	var serverID TDSServerID = nil
	deviceID, _ := uuid.New().MarshalBinary()
	_, err := NewTDSPosition(serverID, deviceID, 0)

	if err == nil {
		t.Fatal("NewTDSPosition should have failed and returned error for missing serverID parameter")
	}
}

func TestNewPositionWithoutDeviceIDReturnsError(t *testing.T) {
	serverID, _ := uuid.New().MarshalBinary()
	var deviceID TDSDeviceID = nil
	_, err := NewTDSPosition(serverID, deviceID, 0)

	if err == nil {
		t.Fatal("NewTDSPosition should have failed and returned error for missing deviceID parameter")
	}
}

func TestNewPositionWithValidParamsReturnsNoError(t *testing.T) {
	serverID, _ := uuid.New().MarshalBinary()
	deviceID, _ := uuid.New().MarshalBinary()

	_, err := NewTDSPosition(serverID, deviceID, 0)

	if err != nil {
		t.Fatal("NewTDSPositions should NOT have returned error with valid parameters")
	}
}

func TestNewPositionWithValidParamsReturnsANewPosition(t *testing.T) {
	serverID, _ := uuid.New().MarshalBinary()
	deviceID, _ := uuid.New().MarshalBinary()

	pos, _ := NewTDSPosition(serverID, deviceID, 0)

	if pos == nil {
		t.Fatal("NewTDSPositions should have returned a new TDSPosition with valid parameters")
	}
}
