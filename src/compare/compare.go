package compare

type CompareResult int8

const (
	LessThan CompareResult = iota - 1
	Equal
	GreaterThan
)

// Comparable
type Comparable[T any] interface {
	Equalable[T]
	Compare(T) CompareResult
}
