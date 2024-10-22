package user

import (
	"testing"

	"github.com/ZacharyDuve/RailroadCentralCommandServer/src/persistance"
)

func TestUserImplementsStorable(t *testing.T) {
	var _ persistance.Storable = &User{}
}
