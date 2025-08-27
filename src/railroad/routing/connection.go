package routing

type ConnectionType string

type Connection interface {
	ConnectionType() ConnectionType
	ConnectedTracks() [2]*Track
}

type RailJointer struct {
	conTracks [2]*Track
}

func (r *RailJointer) ConnectionType() ConnectionType {
	return ConnectionType("Rail Jointer")
}

func (r *RailJointer) ConnectedTracks() [2]*Track {
	return r.conTracks
}
