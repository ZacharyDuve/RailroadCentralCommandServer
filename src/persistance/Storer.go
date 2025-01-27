package persistance

type Storer[T Persistable[K], K comparable] interface {
	StoreNew([]T) error
}
