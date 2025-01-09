package turnout

type TDSServerID []byte

func (this TDSServerID) isBlank() bool {
	return len([]byte(this)) == 0
}
