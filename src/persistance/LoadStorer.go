package persistance

type LoadStorer[T Persistable] interface {
	Loader[T]
	Storer[T]
}
