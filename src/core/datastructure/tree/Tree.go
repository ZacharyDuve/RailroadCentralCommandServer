package tree

type Tree[T any] interface {
	Add(T) bool
	Remove(T) bool
	Clear(T)
}