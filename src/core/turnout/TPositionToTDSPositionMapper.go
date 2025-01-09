package turnout

import (
	"errors"
	"fmt"
	"strings"
)

const (
	mapping_already_exists_prefix string = "Mapping_Already_Exists"
	mapping_does_not_exist_prefix string = "Mapping_Does_NOT_Exist"
)

/*
Mapper to manage the mappings between a logical turnout vs the actual positions in a TDS to set.
We need to be able to have a two way relationship. Need TurnoutPosition to TDSPosition for throwing turnouts while also needing the
reverse for being able to tell when a physical driver has changed its position to update the logical position of the logical turnout.
*/

type turnoutToTDSPostionMapper struct {
}

/*
Create a relationship between the passed in TurnoutPosition and TurnoutPosition
Throws Mapping_Already_Exists error if there already is a mapping between the two
*/
func (this *turnoutToTDSPostionMapper) MakeRelationship(tP TurnoutPosition, tdsP *TDSPosition) error {
	return errors.ErrUnsupported
}

func (this *turnoutToTDSPostionMapper) BreakRelationship(tP TurnoutPosition, tdsP *TDSPosition) error {
	return errors.ErrUnsupported
}

func NewMappingAlreadyExistsErr(tP TurnoutPosition, tdsP *TDSPosition) error {
	return fmt.Errorf("%s for %s and %s", mapping_already_exists_prefix, tP, tdsP.String())
}

func IsNewMappingAlreadyExistsErr(err error) bool {
	return strings.HasPrefix(err.Error(), mapping_already_exists_prefix)
}

func NewMappingDoesNotExistErr(tP TurnoutPosition, tdsP *TDSPosition) error {
	return fmt.Errorf("%s for %s and %s", mapping_does_not_exist_prefix, tP, tdsP.String())
}

func IsNewMappingDoesNotExistErr(err error) bool {
	return strings.HasPrefix(err.Error(), mapping_does_not_exist_prefix)
}
