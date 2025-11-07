package eventing

import "time"

type EventID []byte

type Event struct {
	topic    string
	sendTime time.Time
	eventID  EventID
	// event could be received by more than one listener
	data []byte
}

type EventBus interface {
	AddSubsciber()
	RemoveSubsciber()
	PublishEvent(Event)
}

type EventSubscriber interface {
	HandleEvent(*Event) error
}
