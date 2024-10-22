package tds

import (
	"errors"
	"fmt"
	"log/slog"
	"sync"

	"github.com/ZacharyDuve/RailroadCentralCommandServer/src/eventing"
	"github.com/ZacharyDuve/apireg"
	"github.com/google/uuid"
)

const (
	API_SERVICE_NAME_TDS        string = "Turnout-Driver-Service"
	API_SERVICE_NAME_TDS_SERVER string = API_SERVICE_NAME_TDS + "-Server"
)

// Manages making connections to Turnout Driver Servers and handling requests
type TDSClientManager struct {
	//activePositionIDsByDevice map[TurnoutDriverID]TurnoutDriverPositionID
	apiRegistrySvc         apireg.ApiRegistry
	tdsEventHandlerManager eventing.EventHandlerManager[*TDSEvent]
	tdsClientsLock         sync.Mutex
	tdsClients             map[uuid.UUID]*tdsClient
	tdsEventsInChan        chan *TDSEvent
}

func NewTDSClientManager(apiReg apireg.ApiRegistry) (*TDSClientManager, error) {
	if apiReg == nil {
		return nil, errors.New("unable to create a new TDSClientManager without an api registry")
	}
	tdsCM := &TDSClientManager{apiRegistrySvc: apiReg, tdsEventsInChan: make(chan *TDSEvent)}
	tdsCM.apiRegistrySvc.AddEventListener(tdsCM)

	return tdsCM, nil
}

func (this *TDSClientManager) RequestPosition(posID *TurnoutDriverPositionID) error {
	if posID == nil {
		return errors.New("unable to request position for nil TDSPositionID")
	}

	return nil
}

func (this *TDSClientManager) RegisterTDSEventHandler(tdsEH eventing.EventHandler[*TDSEvent]) error {
	return this.tdsEventHandlerManager.RegisterEventHandler(tdsEH)
}

// HandleRegistration implements apievent.RegistrationListener.
func (this *TDSClientManager) HandleRegistration(rE apireg.RegistrationEvent) {

	//NOTE: This could be fired in a separate thread so we need to make sure to be thread safe

	//Need to make sure that it is matching for service
	if rE.Api().Name() != API_SERVICE_NAME_TDS_SERVER {
		//Service that we don't know how to handle therefore lets ignore
		return
	}

	if rE.Type() == apireg.Added {
		//Need to connect to the server with a new client
		this.handleServerAddedApiRegistration(rE)
	} else if rE.Type() == apireg.Removed {
		//Need to figure out what we want to do for removed. Probably close the client if it is still open
		this.handleServerRemovedApiRegistration(rE)
	} else {
		panic(fmt.Errorf("error handling registration as type %v is unknown to us", rE.Type()))
	}
}

func (this *TDSClientManager) handleServerAddedApiRegistration(rE apireg.RegistrationEvent) {
	this.tdsClientsLock.Lock()

	//TODO: PROBLEM. How do we know if a client is already configured until after we do the whole handshake dance at which point we have to stage events

	this.tdsClientsLock.Unlock()
}

func (this *TDSClientManager) handleServerRemovedApiRegistration(rE apireg.RegistrationEvent) {
	this.tdsClientsLock.Lock()

	client, hasClient := this.tdsClients[rE.Api().UUID()]

	//If we found a matching client then lets delete it
	if hasClient {
		err := client.Close()
		//If an error ocurred while trying to close the client then lets log it so we know it happened.
		//TODO: figure out if there is more we can do
		if err != nil {
			slog.Warn("issue ocurred when trying to close client", "err", err)
		}
		delete(this.tdsClients, rE.Api().UUID())
	}

	this.tdsClientsLock.Unlock()
}

func (this *TDSClientManager) Close() error {
	panic(errors.ErrUnsupported)
}
