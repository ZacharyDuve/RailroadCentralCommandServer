package routing

import "errors"

type JunctionType string

type JunctionLead struct {
	name            string
	track           *Track
	ownedByJunction Junction
	active          bool
}

const (
	JUNCTION_LEAD_NAME_ENTRANCE      string = "entrance"
	JUNCTION_LEAD_NAME_THROUGH       string = "through"
	JUNCTION_LEAD_NAME_DIVERGE_LEFT  string = "diverge left"
	JUNCTION_LEAD_NAME_DIVERGE_RIGHT string = "diverge right"
)

type JunctionRoute struct {
	// Is a list of pairs of points that are connections
	connections [][2]*JunctionLead
}

type Junction interface {
	// External connections of the junction to external
	JunctionType() JunctionType
	JunctionLeads() []*JunctionLead
	ActiveJunctionRoute() *JunctionRoute
	AvailableJunctionRoutes() []*JunctionRoute
	// Potentially could be an error in setting the route for the junction
	SetActiveJunctionRoute(*JunctionRoute) error
}

// A motor that controls the points on a junction
type JunctionPointMotor struct {
}

type jImpl struct {
	jType       JunctionType
	leads       []JunctionLead
	routes      []JunctionRoute
	activeRoute *JunctionRoute
}

// Most Turnouts are simple single entrance with two other leads.
// A route between the entrance and each of the other leads.
// This is to simplify creation
func newjImplWith3LeadsAnd2Routes(outAName, outBName string) *jImpl {
	j := &jImpl{leads: make([]JunctionLead, 0, 3), routes: make([]JunctionRoute, 0, 2)}

	entranceLead := JunctionLead{name: JUNCTION_LEAD_NAME_ENTRANCE, ownedByJunction: j}
	leadA := JunctionLead{name: outAName, ownedByJunction: j}
	leadB := JunctionLead{name: outBName, ownedByJunction: j}

	j.leads = append(j.leads, entranceLead)
	j.leads = append(j.leads, leadA)
	j.leads = append(j.leads, leadB)

	divergeRoute := JunctionRoute{connections: [][2]*JunctionLead{[2]*JunctionLead{&entranceLead, &leadA}}}
	j.routes = append(j.routes, divergeRoute)
	throughRoute := JunctionRoute{connections: [][2]*JunctionLead{[2]*JunctionLead{&entranceLead, &leadB}}}
	j.routes = append(j.routes, throughRoute)

	return j
}

func (this *jImpl) JunctionType() JunctionType {
	return this.jType
}

func (this *jImpl) JunctionLeads() []*JunctionLead {
	leadsRef := make([]*JunctionLead, len(this.leads))
	for i, l := range this.leads {
		leadsRef[i] = &l
	}
	return leadsRef
}

func (this *jImpl) ActiveJunctionRoute() *JunctionRoute {
	return this.activeRoute
}

func (this *jImpl) AvailableJunctionRoutes() []*JunctionRoute {
	routesRef := make([]*JunctionRoute, len(this.routes))
	for i, r := range this.routes {
		routesRef[i] = &r
	}

	return routesRef
}

func (this *jImpl) SetActiveJunctionRoute(r *JunctionRoute) error {

	hasRoute := false

	for _, curR := range this.routes {
		if &curR == r {
			// If we are pointing to the same thing
			hasRoute = true
			break
		}
	}

	if !hasRoute {
		return errors.New("unable to set route as it is not part of this junction")
	}

	this.activeRoute = r

	return nil
}

func NewWyeTurnout() Junction {
	return newjImplWith3LeadsAnd2Routes(JUNCTION_LEAD_NAME_DIVERGE_LEFT, JUNCTION_LEAD_NAME_DIVERGE_RIGHT)
}

func NewLeftHandedTurnout() Junction {
	return newjImplWith3LeadsAnd2Routes(JUNCTION_LEAD_NAME_DIVERGE_LEFT, JUNCTION_LEAD_NAME_THROUGH)
}

func NewRightHandedTurnout() Junction {
	return newjImplWith3LeadsAnd2Routes(JUNCTION_LEAD_NAME_THROUGH, JUNCTION_LEAD_NAME_DIVERGE_RIGHT)
}
