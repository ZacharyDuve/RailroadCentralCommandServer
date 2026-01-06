package eventing

import "testing"

type mockSubscriber struct {
	handleEventValidationFunc func(*Event[int, int]) error
}

func (ms *mockSubscriber) HandleEvent(e *Event[int, int]) error {
	if ms.handleEventValidationFunc != nil {
		return ms.handleEventValidationFunc(e)
	}

	return nil
}

func TestEventChannelStartsWithNoSubscribers(t *testing.T) {
	ec := NewEventChannel[int, int]()
	if len(ec.subscribers) != 0 {
		t.Error("expected new event channel to have no subscribers")
	}
}

func TestEventChannelAddingUnsubscribedAdds(t *testing.T) {
	ec := NewEventChannel[int, int]()

	ec.AddSubScriber(&mockSubscriber{})

	if len(ec.subscribers) != 1 {
		t.Error("expected event channel to have one subscriber after adding a single non duplicate")
	}
}

func TestEventChannelAddingSubscribedDoesNotAddAndReturnsError(t *testing.T) {
	ec := NewEventChannel[int, int]()

	sub := &mockSubscriber{}

	ec.AddSubScriber(sub)

	err := ec.AddSubScriber(sub)

	if len(ec.subscribers) != 1 {
		t.Error("expected event channel to have only one subscriber after attempting to add duplicate")
	}

	if err == nil {
		t.Error("expected error to have been returned when duplicate subscriber was attempted to be added")
	} else if err.Error() == ErrMsgUnableAddSubscriber {
		t.Error("expected the error of attempting to add duplicate to be add error message")
	}
}
