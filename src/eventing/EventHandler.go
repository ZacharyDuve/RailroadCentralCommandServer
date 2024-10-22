package eventing

type EventHandler[E any] interface {
	HandleEvent(E)
}
