package railroad

import (
	"container/list"
	"fmt"

	"github.com/ZacharyDuve/RailroadCentralCommandServer/src/persistance"
)

type memoryRailroadStorer struct {
	railroads []*list.List
}

func NewMemoryRailroadStorer() (RailroadStorer, error) {
	mRRS := &memoryRailroadStorer{}
	mRRS.railroads = make([]*railroad, 0)

	return mRRS, nil
}

func (this *memoryRailroadStorer) Add(nRR *railroad) error {
	for _, curRR := range this.railroads {
		if curRR.PrimaryKey() == nRR.PrimaryKey() {
			return fmt.Errorf("unable to add new railroad due to existing already having id %s", nRR.PrimaryKey())
		}
	}

	this.railroads = append(this.railroads, nRR)

	return nil
}

func (this *memoryRailroadStorer) Update(newRRVersion *railroad) error {
	var rrToUpdateIndex int = -1
	for i, curRR := range this.railroads {
		if curRR.PrimaryKey() == newRRVersion.PrimaryKey() {
			rrToUpdateIndex = i
		}
	}

	if rrToUpdateIndex == -1 {
		return fmt.Errorf("Unable to update railroad with primary key of %s due to it not having be previously stored", newRRVersion.PrimaryKey())
	}

	this.railroads[rrToUpdateIndex] = newRRVersion

	return nil
}

func (this *memoryRailroadStorer) Delete(pk persistance.PrimaryKey) error {
	return nil
}

func (this *memoryRailroadStorer) FindByPK(pk persistance.PrimaryKey) (*railroad, error) {
	return nil, nil
}

func (this *memoryRailroadStorer) FindByName(name string) (*railroad, error) {
	return nil, nil
}
