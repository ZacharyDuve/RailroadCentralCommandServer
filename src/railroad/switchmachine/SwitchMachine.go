package switchmachine

import (
	"errors"
)

type SwitchMachineId string

type SwitchMachine struct {
	//Human readable name for the switch machine. Starts are the FQSN
	name string
	//Human readable string
	description string
	//Switch Machine Server UUID
	id SwitchMachineId

	positionMap []uint8
	driver      SwitchMachineDriver
	persistance SwitchMachinePersister
}

func (this *SwitchMachine) Name() string {
	return this.name
}

func (this *SwitchMachine) UpdateName(string) error {
	panic(errors.ErrUnsupported)
}

func (this *SwitchMachine) Description() string {
	return this.description
}

func (this *SwitchMachine) UpdateDescription(newDescription string) {
	this.description = newDescription
}
