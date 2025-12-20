package id

type ID[T any] interface {
	ID() T
	equal.Equalable
}
