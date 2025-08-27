package routing

type JunctionType string

type JunctionLead struct {
	name            string
	track           *Track
	ownedByJunction Junction
	active          bool
}

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

type jImpl struct {
	leads       []JunctionLead
	routes      []JunctionRoute
	activeRoute *JunctionRoute
}

func newjImpl(nLeads int, nRoutes int) jImpl {
	return jImpl{leads: make([]JunctionLead, 0, nLeads), routes: make([]JunctionRoute, 0, nRoutes)}
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
	panic
}

type LeftHandedTurnout struct {
	jImpl
}

func NewLeftHandedTurnout() *LeftHandedTurnout {
	j := &LeftHandedTurnout{jImpl: newjImpl(3, 2)}
	entranceLead := JunctionLead{name: "entrance", ownedByJunction: j}
	leftLead := JunctionLead{name: "diverge left", ownedByJunction: j}
	throughLead := JunctionLead{name: "through", ownedByJunction: j}

	j.leads = append(j.leads, entranceLead)
	j.leads = append(j.leads, leftLead)
	j.leads = append(j.leads, throughLead)

	divergeRoute := JunctionRoute{connections: [][2]*JunctionLead{[2]*JunctionLead{&entranceLead, &leftLead}}}
	j.routes = append(j.routes, divergeRoute)
	throughRoute := JunctionRoute{connections: [][2]*JunctionLead{[2]*JunctionLead{&entranceLead, &throughLead}}}
	j.routes = append(j.routes, throughRoute)
	j.activeRoute = &throughRoute

	return j
}

func (this *LeftHandedTurnout) JunctionType() JunctionType {
	return JunctionType("Left Handed")
}
