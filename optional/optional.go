package optional

type Optional[T any] struct {
	value T
	valid bool
}

func New[T any](v T) Optional[T] {
	return Optional[T]{
		value: v,
		valid: true,
	}
}

func None[T any]() Optional[T] {
	return Optional[T]{}
}

func (v Optional[T]) Valid() bool {
	return v.valid
}
