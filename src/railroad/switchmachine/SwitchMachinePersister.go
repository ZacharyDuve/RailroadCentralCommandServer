package switchmachine

type SwitchMachinePersister interface {
	SaveOrUpdateSwitchMachineInfo(s *SwitchMachine) error
}
