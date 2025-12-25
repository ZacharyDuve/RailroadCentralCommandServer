package compare

// Equality is an interface that when implemented allows an item to be compared for equality.
// Does not compare whether less or greater than.
type Equality[T any] interface {
	Equal(T) bool
}
