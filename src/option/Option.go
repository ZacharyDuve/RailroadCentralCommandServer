package option

type Option[T any] struct {
	t T
}

func (this *Option[T]) IsSome() bool {
	if this == nil {
		return false
	}

	return true
}

func (this *Option[T]) Value() T {
	if !this.IsSome() {
		panic("Getting a value from an option that doesn't have any")
	}

	return this.t
}
