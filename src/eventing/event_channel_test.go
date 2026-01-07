package eventing

import (
	"testing"
	"time"
)

//************************** Mocking **************************

type mockSubscriber struct {
	handleEventValidationFunc func(*Event[int, int]) error
	receivedEvents            []*Event[int, int]
}

func (ms *mockSubscriber) HandleEvent(e *Event[int, int]) error {
	if ms.handleEventValidationFunc != nil {
		return ms.handleEventValidationFunc(e)
	}

	return nil
}

//========================== NewEventChannel ==========================

func TestEventChannel_New_StartsWithNoSubscribers(t *testing.T) {
	ec := NewEventChannel[int, int]()
	if len(ec.subscribers) != 0 {
		t.Error("expected new event channel to have no subscribers")
	}
}

//========================== AddSubScriber ==========================

func TestEventChannel_AddSubscribe_WhenAddingNewSubscriber_AddsSubscriber(t *testing.T) {
	ec := NewEventChannel[int, int]()

	ec.AddSubScriber(&mockSubscriber{})

	if len(ec.subscribers) != 1 {
		t.Error("expected event channel to have one subscriber after adding a single non duplicate")
	}
}

func TestEventChannel_AddSubscriber_WhenAlreadySubscribed_DoesNotAddAndReturnsError(t *testing.T) {
	ec := NewEventChannel[int, int]()

	sub := &mockSubscriber{}

	ec.AddSubScriber(sub)

	err := ec.AddSubScriber(sub)

	if len(ec.subscribers) != 1 {
		t.Error("expected event channel to have only one subscriber after attempting to add duplicate")
	}

	if err == nil {
		t.Error("expected error to have been returned when duplicate subscriber was attempted to be added")
	} else if err.Error() != ErrMsgUnableAddSubscriber {
		t.Error("expected the error of attempting to add duplicate to be add error message")
	}
}

//========================== RemoveSubscriber ==========================

func TestEventChannel_RemoveSubscriber_DoesNotRemoveAndReturnsError_WhenNotSubscribed(t *testing.T) {
	ec := NewEventChannel[int, int]()

	sub := &mockSubscriber{}

	err := ec.RemoveSubscriber(sub)

	if len(ec.subscribers) != 0 {
		t.Error("expected the subscribers to be unmodified")
	}

	if err == nil {
		t.Error("expected error to have been returned when subscriber was attempted to be removed when it never had been added")
	} else if err.Error() != ErrMsgUnableRemoveSubscriber {
		t.Error("expected the error of attempting to remove an un-added subscriber to be remove error message")
	}
}

func TestEventChannel_RemoveSubscriber_DoesRemove_WhenSubscribed(t *testing.T) {
	ec := NewEventChannel[int, int]()

	sub := &mockSubscriber{}

	ec.AddSubScriber(sub)

	if len(ec.subscribers) != 1 {
		t.Error("expected the subscribers to be 1 after adding before remove")
	}

	err := ec.RemoveSubscriber(sub)

	if len(ec.subscribers) != 0 {
		t.Error("expected the subscribers to be back to zero after removing")
	}

	if err != nil {
		t.Error("expected no error to have been returned when subscriber was removed after it had been added")
	}
}

//========================== SendEvent ==========================

func TestEventChannel_SendEvent_WhenSubscribed_HandlerReceiveEvent(t *testing.T) {
	sub := &mockSubscriber{}

	sub.receivedEvents = make([]*Event[int, int], 0)
	sub.handleEventValidationFunc = func(e *Event[int, int]) error {
		sub.receivedEvents = append(sub.receivedEvents, e)
		return nil
	}

	ec := NewEventChannel[int, int]()

	err := ec.AddSubScriber(sub)

	if err != nil {
		t.Errorf("expected to have no error when adding un-added subscriber but got error: %v", err)
	}

	eb, err := NewEventBuilder[int, int](time.Now)

	if err != nil {
		t.Errorf("expected to have no error when creating event builder but got error: %v", err)
	}

	e := eb.NewEvent(0, 1)

	ec.SendEvent(e)

	// if the subscriber never got event we expect this to hang
	if len(sub.receivedEvents) != 1 || sub.receivedEvents[0] != e {
		t.Error("expected for subscriber to event to have received it but none was")
	}
}

func TestEventChannel_SendEvent_WhenNotSubscribed_HandlerDoesNotReceiveEvent(t *testing.T) {
	sub := &mockSubscriber{}

	sub.receivedEvents = make([]*Event[int, int], 0)
	sub.handleEventValidationFunc = func(e *Event[int, int]) error {
		sub.receivedEvents = append(sub.receivedEvents, e)
		return nil
	}

	ec := NewEventChannel[int, int]()

	eb, err := NewEventBuilder[int, int](time.Now)

	if err != nil {
		t.Errorf("expected to have no error when creating event builder but got error: %v", err)
	}

	e := eb.NewEvent(0, 1)

	ec.SendEvent(e)

	// if the subscriber never got event we expect this to hang
	if len(sub.receivedEvents) != 0 {
		t.Error("expected for handler that never subscribed to never get event")
	}
}
