package turnout

type TDSDeviceID []byte

func (this TDSDeviceID) isBlank() bool {
	return len([]byte(this)) == 0
}

type TDSDevicePositionID uint16
