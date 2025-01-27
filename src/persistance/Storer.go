package persistance

type Storer[T Persistable] interface {
	Store([]T) error
}
