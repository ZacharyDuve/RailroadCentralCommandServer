package block

// Segment really is just a join of two connectors and be related to occupancy detectors
type Segment struct {
	// The two Connectors that this segment is connected to. There should always be two
	connectors [2]*Connector
}

type ConnectorType uint8

const (
	// EndConnector is a connector that only connects to a single line segment. It is effectively an end of the line
	EndConnector ConnectorType = iota
	// BlockDivision is a connection that connects two segments and divides them into blocks
	BlockDivisionConnector
	// A JunctionConnector is a connector that connects 3 or more segments together
	JunctionConnector
)

type Connector struct {
	connectorType ConnectorType
}
