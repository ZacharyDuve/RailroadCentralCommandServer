package eventing

import (
	"cmp"
	"errors"
)

const (
	ErrMsgUnableAddSubscriber    string = "unable to add subscriber as it is already subscribed"
	ErrMsgUnableRemoveSubscriber string = "unable to remove subscriber as it is not subscribed"
)

type eventChannel[I cmp.Ordered, T any] struct {
	subscribers []EventSubscriber[I, T]
}

func NewEventChannel[I cmp.Ordered, T any]() *eventChannel[I, T] {
	return &eventChannel[I, T]{subscribers: make([]EventSubscriber[I, T], 0)}
}

func (ec *eventChannel[I, T]) AddSubScriber(newSub EventSubscriber[I, T]) error {

	for _, curSubscriber := range ec.subscribers {
		if curSubscriber == newSub {
			return errors.New(ErrMsgUnableAddSubscriber)
		}
	}

	ec.subscribers = append(ec.subscribers, newSub)

	return nil
}

func (ec *eventChannel[I, T]) RemoveSubscriber(subscriber EventSubscriber[I, T]) error {
	for i, curSubscriber := range ec.subscribers {
		if curSubscriber == subscriber {
			ec.subscribers = append(ec.subscribers[:i], ec.subscribers[i+1:]...)
			return nil
		}
	}

	return errors.New(ErrMsgUnableRemoveSubscriber)
}

func (ec *eventChannel[I, T]) SendEvent(e *Event[I, T]) {
	for _, curSubscriber := range ec.subscribers {
		curSubscriber.HandleEvent(e)
	}
}
