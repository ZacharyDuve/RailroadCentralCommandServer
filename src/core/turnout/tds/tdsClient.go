package tds

import (
	"errors"

	"github.com/ZacharyDuve/apireg"
)

type tdsClient struct {
	clientForAPI     apireg.Api
	tdsClientManager *TDSClientManager
}

func newTDSClient(a apireg.Api, tCM *TDSClientManager) {
	if tCM == nil {
		//If we get here then something horribly has gone wrong
		panic("error ocurred while trying to create new tdsClient without a related TDSClientManager")
	}

	if a == nil {
		//Also if we get here something has gone wrong from a coding standpoint
		panic("error ocurred while trying to create new tdsClient without an API")
	}

	//Want to start dialing out to the server to connect.
	//Need to make websocket connection to server first so we get events as they trickle in

	//After websocket connection then we start collecting the actual state of the system

}

func (this *tdsClient) Close() error {
	return errors.ErrUnsupported
}
