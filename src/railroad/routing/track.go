package routing

type TrackType uint8

const (
	Mainline TrackType = iota
	Branchline
	Yard
	Spur
)

type Track struct {
	tType          TrackType
	endConnections [2]*Connection
}
