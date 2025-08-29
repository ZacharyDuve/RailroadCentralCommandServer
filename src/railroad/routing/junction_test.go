package routing

import "testing"

func TestjImplImplementsJunction(t *testing.T) {
	var _ Turnout = &jImpl{}
}
