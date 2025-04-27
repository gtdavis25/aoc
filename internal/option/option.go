package option

type Option[T any] struct {
	value T
	set   bool
}

func New[T any](value T) Option[T] {
	return Option[T]{
		value: value,
		set:   true,
	}
}

func (o Option[T]) Value() T {
	return o.value
}

func (o Option[T]) Set() bool {
	return o.set
}
