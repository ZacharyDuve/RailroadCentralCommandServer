package routing

import "github.com/google/uuid"

/*
Routing for trains on a layout
Layout is made up of tracks.
Tracks can be though of as
*/

//type RoutePathID uint64

type RoutePath struct {
	fromNode *RouteNode
	toNode   *RouteNode
	length   float32
}

type RouteNodeType uint8

const (
	Terminus RouteNodeType = iota
	Interconnect
)

type RouteNode struct {
	name          string
	id            uuid.UUID
	nodeType      RouteNodeType
	activePath    *RoutePath
	possiblePaths []RoutePath
}

func (this *RouteNode) Name() string {
	return this.name
}

func (this *RouteNode) SetName(newName string) {
	this.name = newName
}

func (this *RouteNode) Type() RouteNodeType {
	if len(this.possiblePaths) == 0 {
		return Terminus
	}
	return Interconnect
}

type RouteGraph struct {
	nodes []RouteNode
}

func NewRouteGraph() RouteGraph {
	return RouteGraph{nodes: make([]RouteNode, 0)}
}

func (this *RouteGraph) NewPoint(name string) *RouteNode {
	return &RouteNode{id: uuid.New(), name: name, possiblePaths: make([]RoutePath, 0)}
}
