package eventing

import (
	"cmp"
	"errors"
	"time"
)

const (
	ErrMsgUnableNewEventBuilderMissingEventSendTimeFunc string = "unable to create new EventBuilder due to missing required send time function"
)

type Event[I cmp.Ordered, T any] struct {
	sendTime time.Time
	eventID  I
	// event could be received by more than one listener
	payload T
}

type EventBuilder[I cmp.Ordered, T any] struct {
	eventSendTimeFunc func() time.Time
}

func NewEventBuilder[I cmp.Ordered, T any](eventSendTimeFunc func() time.Time) (*EventBuilder[I, T], error) {
	if eventSendTimeFunc == nil {
		return nil, errors.New(ErrMsgUnableNewEventBuilderMissingEventSendTimeFunc)
	}

	return &EventBuilder[I, T]{eventSendTimeFunc: eventSendTimeFunc}, nil
}

func (eb *EventBuilder[I, T]) NewEvent(id I, payload T) *Event[I, T] {
	return &Event[I, T]{
		sendTime: eb.eventSendTimeFunc(),
		eventID:  id,
		payload:  payload,
	}
}
