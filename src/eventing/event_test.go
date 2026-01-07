package eventing

import (
	"testing"
	"time"
)

//************************** Mocking **************************

func mockTimeStampFunction(t time.Time) func() time.Time {
	return func() time.Time {
		return t
	}
}

//========================== NewEventBuilder ==========================

func TestNewEventBuilder_ReturnsError_WhenMissingTimeStampFunction(t *testing.T) {
	_, err := NewEventBuilder[int, int](nil)

	if err == nil {
		t.Error("expected an error from NewEventBuilder as no TimeStampFunction was passed in")
	} else if err.Error() != ErrMsgUnableNewEventBuilderMissingEventSendTimeFunc {
		t.Error("expected the error message for missing timestamp function")
	}
}

func TestNewEventBuilder_ReturnsNoError_WhenPassedTimeStampFunction(t *testing.T) {
	_, err := NewEventBuilder[int, int](time.Now)

	if err != nil {
		t.Error("expected no error from NewEventBuilder as TimeStampFunction was passed in")
	}
}

//========================== NewEvent ==========================

func TestNewEvent_HasPassedInID(t *testing.T) {
	id := "127"

	eb, _ := NewEventBuilder[string, int](time.Now)

	e := eb.NewEvent(id, 0)

	if e.ID() != id {
		t.Errorf("expected event to have id of %s but id of %s was returned", id, e.ID())
	}
}

func TestNewEvent_HasPassedInPayload(t *testing.T) {
	payload := 42

	eb, _ := NewEventBuilder[string, int](time.Now)

	e := eb.NewEvent("car", payload)

	if e.Payload() != payload {
		t.Errorf("expected event to have payload of %d but payload of %d was returned", payload, e.Payload())
	}
}

func TestNewEvent_TimeStampFrom_EventBuildersTimestampFunction(t *testing.T) {

	timeStamp := time.Now()

	eb, _ := NewEventBuilder[string, int](mockTimeStampFunction(timeStamp))

	e := eb.NewEvent("car", 2)

	if timeStamp.Compare(e.SendTime()) != 0 {
		t.Errorf("expected event to have timestamp of %v but payload of %v was returned", timeStamp, e.SendTime())
	}
}
