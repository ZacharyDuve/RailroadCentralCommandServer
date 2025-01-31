package turnouttype

import (
	"strings"

	"github.com/hashicorp/go-set/v3"
)

var libary *set.TreeSet[TurnoutType]

func init() {
	libary = set.NewTreeSet(turnoutTypeComp)
	addBuiltInTypes(libary)
}

func DefaultTurnoutTypeLibrary() *set.TreeSet[TurnoutType] {
	return libary
}

// TurnoutType is to group which type of turnout it is
// Examples would be right handed, left handed, double switch
type TurnoutType interface {
	Name() string
	PositionNames() []string
}

type turnoutTypeImpl struct {
	name          string
	positionNames []string
}

func turnoutTypeComp(t0, t1 TurnoutType) int {
	return strings.Compare(t0.Name(), t1.Name())
}

func (this *turnoutTypeImpl) Name() string {
	return this.name
}

func (this *turnoutTypeImpl) PositionNames() []string {
	return this.positionNames
}

func compareTurnoutTypes(t1, t2 TurnoutType) int {
	return strings.Compare(t1.Name(), t2.Name())
}
