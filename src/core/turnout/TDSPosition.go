package turnout

import (
	"errors"
	"fmt"
)

type TDSPosition struct {
	serverID         TDSServerID
	deviceID         TDSDeviceID
	devicePositionID TDSDevicePositionID
}

func NewTDSPosition(serverID TDSServerID, deviceID TDSDeviceID, devicePosID TDSDevicePositionID) (*TDSPosition, error) {
	if serverID.isBlank() {
		return nil, errors.New("unable to create new TDSPosition due to blank serverID")
	} else if deviceID.isBlank() {
		return nil, errors.New("unable to create new TDSPosition due to blank deviceID")
	}

	return &TDSPosition{serverID: serverID, deviceID: deviceID, devicePositionID: devicePosID}, nil
}

func (this *TDSPosition) ServerID() TDSServerID {
	return this.serverID
}

func (this *TDSPosition) DeviceID() TDSDeviceID {
	return this.deviceID
}

func (this *TDSPosition) DevicePositionID() TDSDevicePositionID {
	return this.devicePositionID
}

func (this *TDSPosition) String() string {
	return fmt.Sprintf("%x-%x-%d", this.serverID, this.deviceID, this.devicePositionID)
}
