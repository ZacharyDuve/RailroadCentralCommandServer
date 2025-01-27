package persistance

type Key[T comparable] interface {
	UnMarshal([]byte) (T, error)
	Marshal(T) ([]byte, error)
	Compare(other Key[T]) int8
}
