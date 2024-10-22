package eventing

import (
	"errors"
	"fmt"
	"sync"
)

type EventHandlerManager[E any] struct {
	eventHandlers     []EventHandler[E]
	eventHandlersLock sync.Mutex
}

func (this *EventHandlerManager[E]) RegisterEventHandler(eh EventHandler[E]) error {

	if eh == nil {
		return errors.New("unable to register new event handler due eh passed in being nil")
	}

	dup := false

	this.eventHandlersLock.Lock()

	for _, curEH := range this.eventHandlers {
		if curEH == eh {
			dup = true
		}
	}

	if !dup {
		this.eventHandlers = append(this.eventHandlers, eh)
	}

	this.eventHandlersLock.Unlock()

	if dup {
		return fmt.Errorf("unable to register event handler due to duplicate for %v already being registered", eh)
	}

	return nil
}

func (this *EventHandlerManager[E]) PropagateEvent(e E) {
	this.eventHandlersLock.Lock()

	for _, curEH := range this.eventHandlers {
		curEH.HandleEvent(e)
	}

	this.eventHandlersLock.Unlock()
}
