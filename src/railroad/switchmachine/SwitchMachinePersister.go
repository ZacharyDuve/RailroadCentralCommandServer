package switchmachine

type SwitchMachinePersister interface {
	UpdateSwitchMachine(s *SwitchMachine) error
	AddSwitchMachine
}
