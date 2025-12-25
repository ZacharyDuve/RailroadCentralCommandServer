package compare

type CompareResult int8

const (
	LessThan CompareResult = iota - 1
	Equal
	GreaterThan
)

// Comparable
type Comparable[T any] interface {
	Equality[T]
	Compare(T) CompareResult
}
