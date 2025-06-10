package main

import (
	"log"

	"github.com/ZacharyDuve/RailroadCentralCommandServer/src/core/turnout"
	"github.com/ZacharyDuve/godatacollections/tree"
	_ "github.com/go-sql-driver/mysql"
)

// ...

func main() {
	turnouts, err := tree.NewBST(turnout.CompTurnoutID, turnout.TurnoutIDFromTurnout, nil)
	if err != nil {
		panic(err)
	}

	tm := turnout.NewTurnoutManager(turnout.NewMemTurnoutIDGen(), turnouts)

	t, err := tm.NewTurnout(turnout.TypeLeftHanded)

	if err != nil {
		panic(err)
	}

	log.Println("Created turnout:", t.ID(), t.Name(), t.Type())
}
