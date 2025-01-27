package persistance

type AsyncStorer[T Persistable] interface {
	StoreAsync(items <-chan T) <-chan error
}
