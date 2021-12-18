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

func (v Optional[T]) Get() T {
	if v.valid {
		return v.value
	} else {
		panic("get a value from none")
	}
}

func (v Optional[T]) GetOr(defaults T) T {
	if v.valid {
		return v.value
	} else {
		return defaults
	}
}

func (v Optional[T]) GetOrFunc(f func() T) T {
	if v.valid {
		return v.value
	} else {
		return f()
	}
}
