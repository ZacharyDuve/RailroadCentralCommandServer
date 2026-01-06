package eventing

import "cmp"

type EventSubscriber[I cmp.Ordered, T any] interface {
	HandleEvent(*Event[I, T]) error
}
